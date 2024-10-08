// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/tenling100/shiharaikun/internal/domain"
)

// MockInvoiceUseCase is a mock of InvoiceUseCase interface.
type MockInvoiceUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceUseCaseMockRecorder
}

// MockInvoiceUseCaseMockRecorder is the mock recorder for MockInvoiceUseCase.
type MockInvoiceUseCaseMockRecorder struct {
	mock *MockInvoiceUseCase
}

// NewMockInvoiceUseCase creates a new mock instance.
func NewMockInvoiceUseCase(ctrl *gomock.Controller) *MockInvoiceUseCase {
	mock := &MockInvoiceUseCase{ctrl: ctrl}
	mock.recorder = &MockInvoiceUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceUseCase) EXPECT() *MockInvoiceUseCaseMockRecorder {
	return m.recorder
}

// CreateInvoice mocks base method.
func (m *MockInvoiceUseCase) CreateInvoice(invoice *domain.InvoiceData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoice", invoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInvoice indicates an expected call of CreateInvoice.
func (mr *MockInvoiceUseCaseMockRecorder) CreateInvoice(invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockInvoiceUseCase)(nil).CreateInvoice), invoice)
}

// GetInvoicesByDateRange mocks base method.
func (m *MockInvoiceUseCase) GetInvoicesByDateRange(startDate, endDate string) ([]domain.InvoiceData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoicesByDateRange", startDate, endDate)
	ret0, _ := ret[0].([]domain.InvoiceData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoicesByDateRange indicates an expected call of GetInvoicesByDateRange.
func (mr *MockInvoiceUseCaseMockRecorder) GetInvoicesByDateRange(startDate, endDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoicesByDateRange", reflect.TypeOf((*MockInvoiceUseCase)(nil).GetInvoicesByDateRange), startDate, endDate)
}

// MockCompanyUsecase is a mock of CompanyUsecase interface.
type MockCompanyUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyUsecaseMockRecorder
}

// MockCompanyUsecaseMockRecorder is the mock recorder for MockCompanyUsecase.
type MockCompanyUsecaseMockRecorder struct {
	mock *MockCompanyUsecase
}

// NewMockCompanyUsecase creates a new mock instance.
func NewMockCompanyUsecase(ctrl *gomock.Controller) *MockCompanyUsecase {
	mock := &MockCompanyUsecase{ctrl: ctrl}
	mock.recorder = &MockCompanyUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyUsecase) EXPECT() *MockCompanyUsecaseMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockCompanyUsecase) CreateCompany(company *domain.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", company)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockCompanyUsecaseMockRecorder) CreateCompany(company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockCompanyUsecase)(nil).CreateCompany), company)
}

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserUsecase) CreateUser(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUsecaseMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUsecase)(nil).CreateUser), user)
}
