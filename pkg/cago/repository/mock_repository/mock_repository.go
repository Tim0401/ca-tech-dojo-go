// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "ca-tech-dojo-go/pkg/cago/model"
	repository "ca-tech-dojo-go/pkg/cago/repository"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// MustConnection mocks base method.
func (m *MockRepository) MustConnection() repository.Connection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustConnection")
	ret0, _ := ret[0].(repository.Connection)
	return ret0
}

// MustConnection indicates an expected call of MustConnection.
func (mr *MockRepositoryMockRecorder) MustConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustConnection", reflect.TypeOf((*MockRepository)(nil).MustConnection))
}

// NewConnection mocks base method.
func (m *MockRepository) NewConnection() (repository.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConnection")
	ret0, _ := ret[0].(repository.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewConnection indicates an expected call of NewConnection.
func (mr *MockRepositoryMockRecorder) NewConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConnection", reflect.TypeOf((*MockRepository)(nil).NewConnection))
}

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Chara mocks base method.
func (m *MockConnection) Chara() repository.CharaQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chara")
	ret0, _ := ret[0].(repository.CharaQuery)
	return ret0
}

// Chara indicates an expected call of Chara.
func (mr *MockConnectionMockRecorder) Chara() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chara", reflect.TypeOf((*MockConnection)(nil).Chara))
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// Gacha mocks base method.
func (m *MockConnection) Gacha() repository.GachaQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gacha")
	ret0, _ := ret[0].(repository.GachaQuery)
	return ret0
}

// Gacha indicates an expected call of Gacha.
func (mr *MockConnectionMockRecorder) Gacha() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gacha", reflect.TypeOf((*MockConnection)(nil).Gacha))
}

// GachaProbabilityGroup mocks base method.
func (m *MockConnection) GachaProbabilityGroup() repository.GachaProbabilityGroupQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GachaProbabilityGroup")
	ret0, _ := ret[0].(repository.GachaProbabilityGroupQuery)
	return ret0
}

// GachaProbabilityGroup indicates an expected call of GachaProbabilityGroup.
func (mr *MockConnectionMockRecorder) GachaProbabilityGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GachaProbabilityGroup", reflect.TypeOf((*MockConnection)(nil).GachaProbabilityGroup))
}

// RunTransaction mocks base method.
func (m *MockConnection) RunTransaction(arg0 func(repository.Transaction) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction.
func (mr *MockConnectionMockRecorder) RunTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockConnection)(nil).RunTransaction), arg0)
}

// User mocks base method.
func (m *MockConnection) User() repository.UserQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(repository.UserQuery)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockConnectionMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockConnection)(nil).User))
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Chara mocks base method.
func (m *MockTransaction) Chara() repository.CharaCommand {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chara")
	ret0, _ := ret[0].(repository.CharaCommand)
	return ret0
}

// Chara indicates an expected call of Chara.
func (mr *MockTransactionMockRecorder) Chara() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chara", reflect.TypeOf((*MockTransaction)(nil).Chara))
}

// User mocks base method.
func (m *MockTransaction) User() repository.UserCommand {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(repository.UserCommand)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockTransactionMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockTransaction)(nil).User))
}

// MockUserQuery is a mock of UserQuery interface.
type MockUserQuery struct {
	ctrl     *gomock.Controller
	recorder *MockUserQueryMockRecorder
}

// MockUserQueryMockRecorder is the mock recorder for MockUserQuery.
type MockUserQueryMockRecorder struct {
	mock *MockUserQuery
}

// NewMockUserQuery creates a new mock instance.
func NewMockUserQuery(ctrl *gomock.Controller) *MockUserQuery {
	mock := &MockUserQuery{ctrl: ctrl}
	mock.recorder = &MockUserQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserQuery) EXPECT() *MockUserQueryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockUserQuery) Find(id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserQueryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserQuery)(nil).Find), id)
}

// FindByIDs mocks base method.
func (m *MockUserQuery) FindByIDs(IDs []int) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDs", IDs)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs.
func (mr *MockUserQueryMockRecorder) FindByIDs(IDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockUserQuery)(nil).FindByIDs), IDs)
}

// FindByToken mocks base method.
func (m *MockUserQuery) FindByToken(token string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", token)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken.
func (mr *MockUserQueryMockRecorder) FindByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockUserQuery)(nil).FindByToken), token)
}

// GetAllUserScore mocks base method.
func (m *MockUserQuery) GetAllUserScore() ([]model.UserRanking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserScore")
	ret0, _ := ret[0].([]model.UserRanking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserScore indicates an expected call of GetAllUserScore.
func (mr *MockUserQueryMockRecorder) GetAllUserScore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserScore", reflect.TypeOf((*MockUserQuery)(nil).GetAllUserScore))
}

// MockUserCommand is a mock of UserCommand interface.
type MockUserCommand struct {
	ctrl     *gomock.Controller
	recorder *MockUserCommandMockRecorder
}

// MockUserCommandMockRecorder is the mock recorder for MockUserCommand.
type MockUserCommandMockRecorder struct {
	mock *MockUserCommand
}

// NewMockUserCommand creates a new mock instance.
func NewMockUserCommand(ctrl *gomock.Controller) *MockUserCommand {
	mock := &MockUserCommand{ctrl: ctrl}
	mock.recorder = &MockUserCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCommand) EXPECT() *MockUserCommandMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserCommand) Create(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserCommandMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserCommand)(nil).Create), user)
}

// Find mocks base method.
func (m *MockUserCommand) Find(id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserCommandMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserCommand)(nil).Find), id)
}

// FindByIDs mocks base method.
func (m *MockUserCommand) FindByIDs(IDs []int) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDs", IDs)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs.
func (mr *MockUserCommandMockRecorder) FindByIDs(IDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockUserCommand)(nil).FindByIDs), IDs)
}

// FindByToken mocks base method.
func (m *MockUserCommand) FindByToken(token string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", token)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken.
func (mr *MockUserCommandMockRecorder) FindByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockUserCommand)(nil).FindByToken), token)
}

// GetAllUserScore mocks base method.
func (m *MockUserCommand) GetAllUserScore() ([]model.UserRanking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserScore")
	ret0, _ := ret[0].([]model.UserRanking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserScore indicates an expected call of GetAllUserScore.
func (mr *MockUserCommandMockRecorder) GetAllUserScore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserScore", reflect.TypeOf((*MockUserCommand)(nil).GetAllUserScore))
}

// UpdateName mocks base method.
func (m *MockUserCommand) UpdateName(name string, UpdatedAt time.Time, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateName", name, UpdatedAt, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateName indicates an expected call of UpdateName.
func (mr *MockUserCommandMockRecorder) UpdateName(name, UpdatedAt, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockUserCommand)(nil).UpdateName), name, UpdatedAt, id)
}

// MockGachaQuery is a mock of GachaQuery interface.
type MockGachaQuery struct {
	ctrl     *gomock.Controller
	recorder *MockGachaQueryMockRecorder
}

// MockGachaQueryMockRecorder is the mock recorder for MockGachaQuery.
type MockGachaQueryMockRecorder struct {
	mock *MockGachaQuery
}

// NewMockGachaQuery creates a new mock instance.
func NewMockGachaQuery(ctrl *gomock.Controller) *MockGachaQuery {
	mock := &MockGachaQuery{ctrl: ctrl}
	mock.recorder = &MockGachaQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGachaQuery) EXPECT() *MockGachaQueryMockRecorder {
	return m.recorder
}

// FindByGroupIDs mocks base method.
func (m *MockGachaQuery) FindByGroupIDs(groupIDs []string) ([]model.GachaProbability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGroupIDs", groupIDs)
	ret0, _ := ret[0].([]model.GachaProbability)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGroupIDs indicates an expected call of FindByGroupIDs.
func (mr *MockGachaQueryMockRecorder) FindByGroupIDs(groupIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGroupIDs", reflect.TypeOf((*MockGachaQuery)(nil).FindByGroupIDs), groupIDs)
}

// MockCharaQuery is a mock of CharaQuery interface.
type MockCharaQuery struct {
	ctrl     *gomock.Controller
	recorder *MockCharaQueryMockRecorder
}

// MockCharaQueryMockRecorder is the mock recorder for MockCharaQuery.
type MockCharaQueryMockRecorder struct {
	mock *MockCharaQuery
}

// NewMockCharaQuery creates a new mock instance.
func NewMockCharaQuery(ctrl *gomock.Controller) *MockCharaQuery {
	mock := &MockCharaQuery{ctrl: ctrl}
	mock.recorder = &MockCharaQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharaQuery) EXPECT() *MockCharaQueryMockRecorder {
	return m.recorder
}

// FindByIDs mocks base method.
func (m *MockCharaQuery) FindByIDs(IDs []int) ([]model.Chara, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDs", IDs)
	ret0, _ := ret[0].([]model.Chara)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs.
func (mr *MockCharaQueryMockRecorder) FindByIDs(IDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockCharaQuery)(nil).FindByIDs), IDs)
}

// FindUserCharaByUserID mocks base method.
func (m *MockCharaQuery) FindUserCharaByUserID(UserID int) ([]model.CharaUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserCharaByUserID", UserID)
	ret0, _ := ret[0].([]model.CharaUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserCharaByUserID indicates an expected call of FindUserCharaByUserID.
func (mr *MockCharaQueryMockRecorder) FindUserCharaByUserID(UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserCharaByUserID", reflect.TypeOf((*MockCharaQuery)(nil).FindUserCharaByUserID), UserID)
}

// MockCharaCommand is a mock of CharaCommand interface.
type MockCharaCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCharaCommandMockRecorder
}

// MockCharaCommandMockRecorder is the mock recorder for MockCharaCommand.
type MockCharaCommandMockRecorder struct {
	mock *MockCharaCommand
}

// NewMockCharaCommand creates a new mock instance.
func NewMockCharaCommand(ctrl *gomock.Controller) *MockCharaCommand {
	mock := &MockCharaCommand{ctrl: ctrl}
	mock.recorder = &MockCharaCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharaCommand) EXPECT() *MockCharaCommandMockRecorder {
	return m.recorder
}

// AddUserChara mocks base method.
func (m *MockCharaCommand) AddUserChara(charaIDs []int, CreatedAt time.Time, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserChara", charaIDs, CreatedAt, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserChara indicates an expected call of AddUserChara.
func (mr *MockCharaCommandMockRecorder) AddUserChara(charaIDs, CreatedAt, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserChara", reflect.TypeOf((*MockCharaCommand)(nil).AddUserChara), charaIDs, CreatedAt, userID)
}

// MockGachaProbabilityGroupQuery is a mock of GachaProbabilityGroupQuery interface.
type MockGachaProbabilityGroupQuery struct {
	ctrl     *gomock.Controller
	recorder *MockGachaProbabilityGroupQueryMockRecorder
}

// MockGachaProbabilityGroupQueryMockRecorder is the mock recorder for MockGachaProbabilityGroupQuery.
type MockGachaProbabilityGroupQueryMockRecorder struct {
	mock *MockGachaProbabilityGroupQuery
}

// NewMockGachaProbabilityGroupQuery creates a new mock instance.
func NewMockGachaProbabilityGroupQuery(ctrl *gomock.Controller) *MockGachaProbabilityGroupQuery {
	mock := &MockGachaProbabilityGroupQuery{ctrl: ctrl}
	mock.recorder = &MockGachaProbabilityGroupQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGachaProbabilityGroupQuery) EXPECT() *MockGachaProbabilityGroupQueryMockRecorder {
	return m.recorder
}

// FindByGachaType mocks base method.
func (m *MockGachaProbabilityGroupQuery) FindByGachaType(gachaTypeID int) ([]model.GachaProbabilityGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGachaType", gachaTypeID)
	ret0, _ := ret[0].([]model.GachaProbabilityGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGachaType indicates an expected call of FindByGachaType.
func (mr *MockGachaProbabilityGroupQueryMockRecorder) FindByGachaType(gachaTypeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGachaType", reflect.TypeOf((*MockGachaProbabilityGroupQuery)(nil).FindByGachaType), gachaTypeID)
}
