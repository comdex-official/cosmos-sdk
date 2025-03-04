package runtime

import (
	"fmt"
	"os"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/comet"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/genesis"
	"cosmossdk.io/core/header"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

type appModule struct {
	app *App
}

func (m appModule) RegisterServices(configurator module.Configurator) {
	err := m.app.registerRuntimeServices(configurator)
	if err != nil {
		panic(err)
	}
}

func (m appModule) IsOnePerModuleType() {}
func (m appModule) IsAppModule()        {}

var (
	_ appmodule.AppModule = appModule{}
	_ module.HasServices  = appModule{}
)

// BaseAppOption is a depinject.AutoGroupType which can be used to pass
// BaseApp options into the depinject. It should be used carefully.
type BaseAppOption func(*baseapp.BaseApp)

// IsManyPerContainerType indicates that this is a depinject.ManyPerContainerType.
func (b BaseAppOption) IsManyPerContainerType() {}

func init() {
	appmodule.Register(&runtimev1alpha1.Module{},
		appmodule.Provide(
			ProvideApp,
			ProvideKVStoreKey,
			ProvideTransientStoreKey,
			ProvideMemoryStoreKey,
			ProvideGenesisTxHandler,
			ProvideKVStoreService,
			ProvideMemoryStoreService,
			ProvideTransientStoreService,
			ProvideEventService,
			ProvideHeaderInfoService,
			ProvideCometInfoService,
			ProvideBasicManager,
		),
		appmodule.Invoke(SetupAppBuilder),
	)
}

func ProvideApp() (
	codectypes.InterfaceRegistry,
	codec.Codec,
	*codec.LegacyAmino,
	*AppBuilder,
	codec.ProtoCodecMarshaler,
	*baseapp.MsgServiceRouter,
	appmodule.AppModule,
	protodesc.Resolver,
	protoregistry.MessageTypeResolver,
	error,
) {
	protoFiles, err := proto.MergedRegistry()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	protoTypes := protoregistry.GlobalTypes

	// At startup, check that all proto annotations are correct.
	if err := msgservice.ValidateProtoAnnotations(protoFiles); err != nil {
		// Once we switch to using protoreflect-based antehandlers, we might
		// want to panic here instead of logging a warning.
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}

	interfaceRegistry, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles:            protoFiles,
		AddressCodec:          globalAccAddressCodec{},
		ValidatorAddressCodec: globalValAddressCodec{},
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	// validate the signing context to make sure that messages are properly configured
	// with cosmos.msg.v1.signer
	if err := interfaceRegistry.SigningContext().Validate(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	amino := codec.NewLegacyAmino()

	std.RegisterInterfaces(interfaceRegistry)
	std.RegisterLegacyAminoCodec(amino)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	msgServiceRouter := baseapp.NewMsgServiceRouter()
	app := &App{
		storeKeys:         nil,
		interfaceRegistry: interfaceRegistry,
		cdc:               cdc,
		amino:             amino,
		basicManager:      module.BasicManager{},
		msgServiceRouter:  msgServiceRouter,
	}
	appBuilder := &AppBuilder{app}

	return interfaceRegistry, cdc, amino, appBuilder, cdc, msgServiceRouter, appModule{app}, protoFiles, protoTypes, nil
}

type AppInputs struct {
	depinject.In

	AppConfig          *appv1alpha1.Config
	Config             *runtimev1alpha1.Module
	AppBuilder         *AppBuilder
	Modules            map[string]appmodule.AppModule
	CustomModuleBasics map[string]module.AppModuleBasic `optional:"true"`
	BaseAppOptions     []BaseAppOption
	InterfaceRegistry  codectypes.InterfaceRegistry
	LegacyAmino        *codec.LegacyAmino
	Logger             log.Logger
}

func SetupAppBuilder(inputs AppInputs) {
	app := inputs.AppBuilder.app
	app.baseAppOptions = inputs.BaseAppOptions
	app.config = inputs.Config
	app.appConfig = inputs.AppConfig
	app.logger = inputs.Logger
	app.ModuleManager = module.NewManagerFromMap(inputs.Modules)

	for name, mod := range inputs.Modules {
		if customBasicMod, ok := inputs.CustomModuleBasics[name]; ok {
			app.basicManager[name] = customBasicMod
			customBasicMod.RegisterInterfaces(inputs.InterfaceRegistry)
			customBasicMod.RegisterLegacyAminoCodec(inputs.LegacyAmino)
			continue
		}

		coreAppModuleBasic := module.CoreAppModuleBasicAdaptor(name, mod)
		app.basicManager[name] = coreAppModuleBasic
		coreAppModuleBasic.RegisterInterfaces(inputs.InterfaceRegistry)
		coreAppModuleBasic.RegisterLegacyAminoCodec(inputs.LegacyAmino)
	}
}

func registerStoreKey(wrapper *AppBuilder, key storetypes.StoreKey) {
	wrapper.app.storeKeys = append(wrapper.app.storeKeys, key)
}

func storeKeyOverride(config *runtimev1alpha1.Module, moduleName string) *runtimev1alpha1.StoreKeyConfig {
	for _, cfg := range config.OverrideStoreKeys {
		if cfg.ModuleName == moduleName {
			return cfg
		}
	}
	return nil
}

func ProvideKVStoreKey(config *runtimev1alpha1.Module, key depinject.ModuleKey, app *AppBuilder) *storetypes.KVStoreKey {
	override := storeKeyOverride(config, key.Name())

	var storeKeyName string
	if override != nil {
		storeKeyName = override.KvStoreKey
	} else {
		storeKeyName = key.Name()
	}

	storeKey := storetypes.NewKVStoreKey(storeKeyName)
	registerStoreKey(app, storeKey)
	return storeKey
}

func ProvideTransientStoreKey(key depinject.ModuleKey, app *AppBuilder) *storetypes.TransientStoreKey {
	storeKey := storetypes.NewTransientStoreKey(fmt.Sprintf("transient:%s", key.Name()))
	registerStoreKey(app, storeKey)
	return storeKey
}

func ProvideMemoryStoreKey(key depinject.ModuleKey, app *AppBuilder) *storetypes.MemoryStoreKey {
	storeKey := storetypes.NewMemoryStoreKey(fmt.Sprintf("memory:%s", key.Name()))
	registerStoreKey(app, storeKey)
	return storeKey
}

func ProvideGenesisTxHandler(appBuilder *AppBuilder) genesis.TxHandler {
	return appBuilder.app
}

func ProvideKVStoreService(config *runtimev1alpha1.Module, key depinject.ModuleKey, app *AppBuilder) store.KVStoreService {
	storeKey := ProvideKVStoreKey(config, key, app)
	return kvStoreService{key: storeKey}
}

func ProvideMemoryStoreService(key depinject.ModuleKey, app *AppBuilder) store.MemoryStoreService {
	storeKey := ProvideMemoryStoreKey(key, app)
	return memStoreService{key: storeKey}
}

func ProvideTransientStoreService(key depinject.ModuleKey, app *AppBuilder) store.TransientStoreService {
	storeKey := ProvideTransientStoreKey(key, app)
	return transientStoreService{key: storeKey}
}

func ProvideEventService() event.Service {
	return EventService{}
}

func ProvideCometInfoService() comet.BlockInfoService {
	return cometInfoService{}
}

func ProvideHeaderInfoService(app *AppBuilder) header.Service {
	return headerInfoService{}
}

func ProvideBasicManager(app *AppBuilder) module.BasicManager {
	return app.app.basicManager
}

// globalAccAddressCodec is a temporary address codec that we will use until we
// can populate it with the correct bech32 prefixes without depending on the global.
type globalAccAddressCodec struct{}

func (g globalAccAddressCodec) StringToBytes(text string) ([]byte, error) {
	if text == "" {
		return nil, nil
	}
	return sdk.AccAddressFromBech32(text)
}

func (g globalAccAddressCodec) BytesToString(bz []byte) (string, error) {
	if bz == nil {
		return "", nil
	}
	return sdk.AccAddress(bz).String(), nil
}

// globalValAddressCodec is a temporary address codec that we will use until we
// can populate it with the correct bech32 prefixes without depending on the global.
type globalValAddressCodec struct{}

func (g globalValAddressCodec) StringToBytes(text string) ([]byte, error) {
	return sdk.ValAddressFromBech32(text)
}

func (g globalValAddressCodec) BytesToString(bz []byte) (string, error) {
	return sdk.ValAddress(bz).String(), nil
}
