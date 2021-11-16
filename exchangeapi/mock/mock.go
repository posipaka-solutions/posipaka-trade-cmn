// Code generated by MockGen. DO NOT EDIT.
// Source: exchangeapi/connector_structs.go

// Package mock_exchangeapi is a generated GoMock package.
package mock_exchangeapi

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	order "github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	symbol "github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

// MockApiConnector is a mock of ApiConnector interface.
type MockApiConnector struct {
	ctrl     *gomock.Controller
	recorder *MockApiConnectorMockRecorder
}

// MockApiConnectorMockRecorder is the mock recorder for MockApiConnector.
type MockApiConnectorMockRecorder struct {
	mock *MockApiConnector
}

// NewMockApiConnector creates a new mock instance.
func NewMockApiConnector(ctrl *gomock.Controller) *MockApiConnector {
	mock := &MockApiConnector{ctrl: ctrl}
	mock.recorder = &MockApiConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiConnector) EXPECT() *MockApiConnectorMockRecorder {
	return m.recorder
}

// GetAssetBalance mocks base method.
func (m *MockApiConnector) GetAssetBalance(asset string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetBalance", asset)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetBalance indicates an expected call of GetAssetBalance.
func (mr *MockApiConnectorMockRecorder) GetAssetBalance(asset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetBalance", reflect.TypeOf((*MockApiConnector)(nil).GetAssetBalance), asset)
}

// GetCurrentPrice mocks base method.
func (m *MockApiConnector) GetCurrentPrice(assets symbol.Assets) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPrice", assets)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentPrice indicates an expected call of GetCurrentPrice.
func (mr *MockApiConnectorMockRecorder) GetCurrentPrice(assets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPrice", reflect.TypeOf((*MockApiConnector)(nil).GetCurrentPrice), assets)
}

// GetOrdersList mocks base method.
func (m *MockApiConnector) GetOrdersList(assets symbol.Assets) ([]order.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersList", assets)
	ret0, _ := ret[0].([]order.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersList indicates an expected call of GetOrdersList.
func (mr *MockApiConnectorMockRecorder) GetOrdersList(assets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersList", reflect.TypeOf((*MockApiConnector)(nil).GetOrdersList), assets)
}

// GetServerTime mocks base method.
func (m *MockApiConnector) GetServerTime() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerTime")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerTime indicates an expected call of GetServerTime.
func (mr *MockApiConnectorMockRecorder) GetServerTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerTime", reflect.TypeOf((*MockApiConnector)(nil).GetServerTime))
}

// GetSymbolsLimits mocks base method.
func (m *MockApiConnector) GetSymbolsLimits() ([]symbol.Limits, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSymbolsLimits")
	ret0, _ := ret[0].([]symbol.Limits)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSymbolsLimits indicates an expected call of GetSymbolsLimits.
func (mr *MockApiConnectorMockRecorder) GetSymbolsLimits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSymbolsLimits", reflect.TypeOf((*MockApiConnector)(nil).GetSymbolsLimits))
}

// GetSymbolsList mocks base method.
func (m *MockApiConnector) GetSymbolsList() []symbol.Assets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSymbolsList")
	ret0, _ := ret[0].([]symbol.Assets)
	return ret0
}

// GetSymbolsList indicates an expected call of GetSymbolsList.
func (mr *MockApiConnectorMockRecorder) GetSymbolsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSymbolsList", reflect.TypeOf((*MockApiConnector)(nil).GetSymbolsList))
}

// SetOrder mocks base method.
func (m *MockApiConnector) SetOrder(parameters order.Parameters) (order.OrderInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrder", parameters)
	ret0, _ := ret[0].(order.OrderInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOrder indicates an expected call of SetOrder.
func (mr *MockApiConnectorMockRecorder) SetOrder(parameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrder", reflect.TypeOf((*MockApiConnector)(nil).SetOrder), parameters)
}

// StoreSymbolsLimits mocks base method.
func (m *MockApiConnector) StoreSymbolsLimits(arg0 []symbol.Limits) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StoreSymbolsLimits", arg0)
}

// StoreSymbolsLimits indicates an expected call of StoreSymbolsLimits.
func (mr *MockApiConnectorMockRecorder) StoreSymbolsLimits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSymbolsLimits", reflect.TypeOf((*MockApiConnector)(nil).StoreSymbolsLimits), arg0)
}
