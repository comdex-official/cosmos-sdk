// Code generated by MockGen. DO NOT EDIT.
// Source: x/evidence/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"
	time "time"

	comet "cosmossdk.io/core/comet"
	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/crypto/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
)

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// GetParams mocks base method.
func (m *MockStakingKeeper) GetParams(ctx types0.Context) types1.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types1.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockStakingKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockStakingKeeper)(nil).GetParams), ctx)
}

// ValidatorByConsAddr mocks base method.
func (m *MockStakingKeeper) ValidatorByConsAddr(arg0 types0.Context, arg1 types0.ConsAddress) types1.ValidatorI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(types1.ValidatorI)
	return ret0
}

// ValidatorByConsAddr indicates an expected call of ValidatorByConsAddr.
func (mr *MockStakingKeeperMockRecorder) ValidatorByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorByConsAddr", reflect.TypeOf((*MockStakingKeeper)(nil).ValidatorByConsAddr), arg0, arg1)
}

// MockSlashingKeeper is a mock of SlashingKeeper interface.
type MockSlashingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSlashingKeeperMockRecorder
}

// MockSlashingKeeperMockRecorder is the mock recorder for MockSlashingKeeper.
type MockSlashingKeeperMockRecorder struct {
	mock *MockSlashingKeeper
}

// NewMockSlashingKeeper creates a new mock instance.
func NewMockSlashingKeeper(ctrl *gomock.Controller) *MockSlashingKeeper {
	mock := &MockSlashingKeeper{ctrl: ctrl}
	mock.recorder = &MockSlashingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlashingKeeper) EXPECT() *MockSlashingKeeperMockRecorder {
	return m.recorder
}

// GetPubkey mocks base method.
func (m *MockSlashingKeeper) GetPubkey(arg0 types0.Context, arg1 types.Address) (types.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubkey", arg0, arg1)
	ret0, _ := ret[0].(types.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubkey indicates an expected call of GetPubkey.
func (mr *MockSlashingKeeperMockRecorder) GetPubkey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubkey", reflect.TypeOf((*MockSlashingKeeper)(nil).GetPubkey), arg0, arg1)
}

// HasValidatorSigningInfo mocks base method.
func (m *MockSlashingKeeper) HasValidatorSigningInfo(arg0 types0.Context, arg1 types0.ConsAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasValidatorSigningInfo", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasValidatorSigningInfo indicates an expected call of HasValidatorSigningInfo.
func (mr *MockSlashingKeeperMockRecorder) HasValidatorSigningInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasValidatorSigningInfo", reflect.TypeOf((*MockSlashingKeeper)(nil).HasValidatorSigningInfo), arg0, arg1)
}

// IsTombstoned mocks base method.
func (m *MockSlashingKeeper) IsTombstoned(arg0 types0.Context, arg1 types0.ConsAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTombstoned", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTombstoned indicates an expected call of IsTombstoned.
func (mr *MockSlashingKeeperMockRecorder) IsTombstoned(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTombstoned", reflect.TypeOf((*MockSlashingKeeper)(nil).IsTombstoned), arg0, arg1)
}

// Jail mocks base method.
func (m *MockSlashingKeeper) Jail(arg0 types0.Context, arg1 types0.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Jail", arg0, arg1)
}

// Jail indicates an expected call of Jail.
func (mr *MockSlashingKeeperMockRecorder) Jail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jail", reflect.TypeOf((*MockSlashingKeeper)(nil).Jail), arg0, arg1)
}

// JailUntil mocks base method.
func (m *MockSlashingKeeper) JailUntil(arg0 types0.Context, arg1 types0.ConsAddress, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JailUntil", arg0, arg1, arg2)
}

// JailUntil indicates an expected call of JailUntil.
func (mr *MockSlashingKeeperMockRecorder) JailUntil(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JailUntil", reflect.TypeOf((*MockSlashingKeeper)(nil).JailUntil), arg0, arg1, arg2)
}

// Slash mocks base method.
func (m *MockSlashingKeeper) Slash(arg0 types0.Context, arg1 types0.ConsAddress, arg2 math.LegacyDec, arg3, arg4 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Slash", arg0, arg1, arg2, arg3, arg4)
}

// Slash indicates an expected call of Slash.
func (mr *MockSlashingKeeperMockRecorder) Slash(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slash", reflect.TypeOf((*MockSlashingKeeper)(nil).Slash), arg0, arg1, arg2, arg3, arg4)
}

// SlashFractionDoubleSign mocks base method.
func (m *MockSlashingKeeper) SlashFractionDoubleSign(arg0 types0.Context) math.LegacyDec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashFractionDoubleSign", arg0)
	ret0, _ := ret[0].(math.LegacyDec)
	return ret0
}

// SlashFractionDoubleSign indicates an expected call of SlashFractionDoubleSign.
func (mr *MockSlashingKeeperMockRecorder) SlashFractionDoubleSign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashFractionDoubleSign", reflect.TypeOf((*MockSlashingKeeper)(nil).SlashFractionDoubleSign), arg0)
}

// SlashWithInfractionReason mocks base method.
func (m *MockSlashingKeeper) SlashWithInfractionReason(arg0 types0.Context, arg1 types0.ConsAddress, arg2 math.LegacyDec, arg3, arg4 int64, arg5 types1.Infraction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SlashWithInfractionReason", arg0, arg1, arg2, arg3, arg4, arg5)
}

// SlashWithInfractionReason indicates an expected call of SlashWithInfractionReason.
func (mr *MockSlashingKeeperMockRecorder) SlashWithInfractionReason(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashWithInfractionReason", reflect.TypeOf((*MockSlashingKeeper)(nil).SlashWithInfractionReason), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Tombstone mocks base method.
func (m *MockSlashingKeeper) Tombstone(arg0 types0.Context, arg1 types0.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Tombstone", arg0, arg1)
}

// Tombstone indicates an expected call of Tombstone.
func (mr *MockSlashingKeeperMockRecorder) Tombstone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tombstone", reflect.TypeOf((*MockSlashingKeeper)(nil).Tombstone), arg0, arg1)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx context.Context, acc types0.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx types0.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx types0.Context, moduleName string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types0.Context, senderModule string, recipientAddr types0.AccAddress, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockCometinfo is a mock of Cometinfo interface.
type MockCometinfo struct {
	ctrl     *gomock.Controller
	recorder *MockCometinfoMockRecorder
}

// MockCometinfoMockRecorder is the mock recorder for MockCometinfo.
type MockCometinfoMockRecorder struct {
	mock *MockCometinfo
}

// NewMockCometinfo creates a new mock instance.
func NewMockCometinfo(ctrl *gomock.Controller) *MockCometinfo {
	mock := &MockCometinfo{ctrl: ctrl}
	mock.recorder = &MockCometinfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCometinfo) EXPECT() *MockCometinfoMockRecorder {
	return m.recorder
}

// GetCometBlockInfo mocks base method.
func (m *MockCometinfo) GetCometBlockInfo(arg0 context.Context) comet.BlockInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCometBlockInfo", arg0)
	ret0, _ := ret[0].(comet.BlockInfo)
	return ret0
}

// GetCometBlockInfo indicates an expected call of GetCometBlockInfo.
func (mr *MockCometinfoMockRecorder) GetCometBlockInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCometBlockInfo", reflect.TypeOf((*MockCometinfo)(nil).GetCometBlockInfo), arg0)
}
