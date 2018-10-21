// Code generated by MockGen. DO NOT EDIT.
// Source: hub.go

// Package mock_homehub is a generated GoMock package.
package cmd

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	homehub_client "github.com/jamesnetherton/homehub-client"
)

// MockHub is a mock of Hub interface
type MockHub struct {
	ctrl     *gomock.Controller
	recorder *MockHubMockRecorder
}

// MockHubMockRecorder is the mock recorder for MockHub
type MockHubMockRecorder struct {
	mock *MockHub
}

// NewMockHub creates a new mock instance
func NewMockHub(ctrl *gomock.Controller) *MockHub {
	mock := &MockHub{ctrl: ctrl}
	mock.recorder = &MockHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHub) EXPECT() *MockHubMockRecorder {
	return m.recorder
}

// BandwidthMonitor mocks base method
func (m *MockHub) BandwidthMonitor() (*homehub_client.BandwidthLog, error) {
	ret := m.ctrl.Call(m, "BandwidthMonitor")
	ret0, _ := ret[0].(*homehub_client.BandwidthLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BandwidthMonitor indicates an expected call of BandwidthMonitor
func (mr *MockHubMockRecorder) BandwidthMonitor() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BandwidthMonitor", reflect.TypeOf((*MockHub)(nil).BandwidthMonitor))
}

// BroadbandProductType mocks base method
func (m *MockHub) BroadbandProductType() (string, error) {
	ret := m.ctrl.Call(m, "BroadbandProductType")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadbandProductType indicates an expected call of BroadbandProductType
func (mr *MockHubMockRecorder) BroadbandProductType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadbandProductType", reflect.TypeOf((*MockHub)(nil).BroadbandProductType))
}

// ConnectedDevices mocks base method
func (m *MockHub) ConnectedDevices() ([]homehub_client.DeviceDetail, error) {
	ret := m.ctrl.Call(m, "ConnectedDevices")
	ret0, _ := ret[0].([]homehub_client.DeviceDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectedDevices indicates an expected call of ConnectedDevices
func (mr *MockHubMockRecorder) ConnectedDevices() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedDevices", reflect.TypeOf((*MockHub)(nil).ConnectedDevices))
}

// DataPumpVersion mocks base method
func (m *MockHub) DataPumpVersion() (string, error) {
	ret := m.ctrl.Call(m, "DataPumpVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataPumpVersion indicates an expected call of DataPumpVersion
func (mr *MockHubMockRecorder) DataPumpVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataPumpVersion", reflect.TypeOf((*MockHub)(nil).DataPumpVersion))
}

// DataReceived mocks base method
func (m *MockHub) DataReceived() (int64, error) {
	ret := m.ctrl.Call(m, "DataReceived")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataReceived indicates an expected call of DataReceived
func (mr *MockHubMockRecorder) DataReceived() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataReceived", reflect.TypeOf((*MockHub)(nil).DataReceived))
}

// DataSent mocks base method
func (m *MockHub) DataSent() (int64, error) {
	ret := m.ctrl.Call(m, "DataSent")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataSent indicates an expected call of DataSent
func (mr *MockHubMockRecorder) DataSent() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataSent", reflect.TypeOf((*MockHub)(nil).DataSent))
}

// DeviceInfo mocks base method
func (m *MockHub) DeviceInfo(id int) (*homehub_client.DeviceDetail, error) {
	ret := m.ctrl.Call(m, "DeviceInfo", id)
	ret0, _ := ret[0].(*homehub_client.DeviceDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeviceInfo indicates an expected call of DeviceInfo
func (mr *MockHubMockRecorder) DeviceInfo(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeviceInfo", reflect.TypeOf((*MockHub)(nil).DeviceInfo), id)
}

// DhcpAuthoritative mocks base method
func (m *MockHub) DhcpAuthoritative() (bool, error) {
	ret := m.ctrl.Call(m, "DhcpAuthoritative")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DhcpAuthoritative indicates an expected call of DhcpAuthoritative
func (mr *MockHubMockRecorder) DhcpAuthoritative() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DhcpAuthoritative", reflect.TypeOf((*MockHub)(nil).DhcpAuthoritative))
}

// DhcpPoolStart mocks base method
func (m *MockHub) DhcpPoolStart() (string, error) {
	ret := m.ctrl.Call(m, "DhcpPoolStart")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DhcpPoolStart indicates an expected call of DhcpPoolStart
func (mr *MockHubMockRecorder) DhcpPoolStart() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DhcpPoolStart", reflect.TypeOf((*MockHub)(nil).DhcpPoolStart))
}

// DhcpPoolEnd mocks base method
func (m *MockHub) DhcpPoolEnd() (string, error) {
	ret := m.ctrl.Call(m, "DhcpPoolEnd")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DhcpPoolEnd indicates an expected call of DhcpPoolEnd
func (mr *MockHubMockRecorder) DhcpPoolEnd() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DhcpPoolEnd", reflect.TypeOf((*MockHub)(nil).DhcpPoolEnd))
}

// DhcpSubnetMask mocks base method
func (m *MockHub) DhcpSubnetMask() (string, error) {
	ret := m.ctrl.Call(m, "DhcpSubnetMask")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DhcpSubnetMask indicates an expected call of DhcpSubnetMask
func (mr *MockHubMockRecorder) DhcpSubnetMask() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DhcpSubnetMask", reflect.TypeOf((*MockHub)(nil).DhcpSubnetMask))
}

// DownstreamSyncSpeed mocks base method
func (m *MockHub) DownstreamSyncSpeed() (int, error) {
	ret := m.ctrl.Call(m, "DownstreamSyncSpeed")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownstreamSyncSpeed indicates an expected call of DownstreamSyncSpeed
func (mr *MockHubMockRecorder) DownstreamSyncSpeed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownstreamSyncSpeed", reflect.TypeOf((*MockHub)(nil).DownstreamSyncSpeed))
}

// EnableDebug mocks base method
func (m *MockHub) EnableDebug(enable bool) {
	m.ctrl.Call(m, "EnableDebug", enable)
}

// EnableDebug indicates an expected call of EnableDebug
func (mr *MockHubMockRecorder) EnableDebug(enable interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDebug", reflect.TypeOf((*MockHub)(nil).EnableDebug), enable)
}

// EnableDhcpAuthoritative mocks base method
func (m *MockHub) EnableDhcpAuthoritative(enable bool) error {
	ret := m.ctrl.Call(m, "EnableDhcpAuthoritative", enable)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableDhcpAuthoritative indicates an expected call of EnableDhcpAuthoritative
func (mr *MockHubMockRecorder) EnableDhcpAuthoritative(enable interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDhcpAuthoritative", reflect.TypeOf((*MockHub)(nil).EnableDhcpAuthoritative), enable)
}

// EventLog mocks base method
func (m *MockHub) EventLog() (*homehub_client.EventLog, error) {
	ret := m.ctrl.Call(m, "EventLog")
	ret0, _ := ret[0].(*homehub_client.EventLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventLog indicates an expected call of EventLog
func (mr *MockHubMockRecorder) EventLog() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventLog", reflect.TypeOf((*MockHub)(nil).EventLog))
}

// HardwareVersion mocks base method
func (m *MockHub) HardwareVersion() (string, error) {
	ret := m.ctrl.Call(m, "HardwareVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardwareVersion indicates an expected call of HardwareVersion
func (mr *MockHubMockRecorder) HardwareVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareVersion", reflect.TypeOf((*MockHub)(nil).HardwareVersion))
}

// InternetConnectionStatus mocks base method
func (m *MockHub) InternetConnectionStatus() (string, error) {
	ret := m.ctrl.Call(m, "InternetConnectionStatus")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InternetConnectionStatus indicates an expected call of InternetConnectionStatus
func (mr *MockHubMockRecorder) InternetConnectionStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternetConnectionStatus", reflect.TypeOf((*MockHub)(nil).InternetConnectionStatus))
}

// LightBrightness mocks base method
func (m *MockHub) LightBrightness() (int, error) {
	ret := m.ctrl.Call(m, "LightBrightness")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LightBrightness indicates an expected call of LightBrightness
func (mr *MockHubMockRecorder) LightBrightness() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LightBrightness", reflect.TypeOf((*MockHub)(nil).LightBrightness))
}

// LightBrightnessSet mocks base method
func (m *MockHub) LightBrightnessSet(brightness int) error {
	ret := m.ctrl.Call(m, "LightBrightnessSet", brightness)
	ret0, _ := ret[0].(error)
	return ret0
}

// LightBrightnessSet indicates an expected call of LightBrightnessSet
func (mr *MockHubMockRecorder) LightBrightnessSet(brightness interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LightBrightnessSet", reflect.TypeOf((*MockHub)(nil).LightBrightnessSet), brightness)
}

// LightEnable mocks base method
func (m *MockHub) LightEnable(enable bool) error {
	ret := m.ctrl.Call(m, "LightEnable", enable)
	ret0, _ := ret[0].(error)
	return ret0
}

// LightEnable indicates an expected call of LightEnable
func (mr *MockHubMockRecorder) LightEnable(enable interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LightEnable", reflect.TypeOf((*MockHub)(nil).LightEnable), enable)
}

// LightStatus mocks base method
func (m *MockHub) LightStatus() (string, error) {
	ret := m.ctrl.Call(m, "LightStatus")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LightStatus indicates an expected call of LightStatus
func (mr *MockHubMockRecorder) LightStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LightStatus", reflect.TypeOf((*MockHub)(nil).LightStatus))
}

// LocalTime mocks base method
func (m *MockHub) LocalTime() (string, error) {
	ret := m.ctrl.Call(m, "LocalTime")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocalTime indicates an expected call of LocalTime
func (mr *MockHubMockRecorder) LocalTime() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalTime", reflect.TypeOf((*MockHub)(nil).LocalTime))
}

// Login mocks base method
func (m *MockHub) Login() (bool, error) {
	ret := m.ctrl.Call(m, "Login")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockHubMockRecorder) Login() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockHub)(nil).Login))
}

// MaintenanceFirmwareVersion mocks base method
func (m *MockHub) MaintenanceFirmwareVersion() (string, error) {
	ret := m.ctrl.Call(m, "MaintenanceFirmwareVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceFirmwareVersion indicates an expected call of MaintenanceFirmwareVersion
func (mr *MockHubMockRecorder) MaintenanceFirmwareVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceFirmwareVersion", reflect.TypeOf((*MockHub)(nil).MaintenanceFirmwareVersion))
}

// NatRules mocks base method
func (m *MockHub) NatRules() ([]homehub_client.NatRule, error) {
	ret := m.ctrl.Call(m, "NatRules")
	ret0, _ := ret[0].([]homehub_client.NatRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NatRules indicates an expected call of NatRules
func (mr *MockHubMockRecorder) NatRules() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatRules", reflect.TypeOf((*MockHub)(nil).NatRules))
}

// NatRule mocks base method
func (m *MockHub) NatRule(id int) (*homehub_client.NatRule, error) {
	ret := m.ctrl.Call(m, "NatRule", id)
	ret0, _ := ret[0].(*homehub_client.NatRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NatRule indicates an expected call of NatRule
func (mr *MockHubMockRecorder) NatRule(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatRule", reflect.TypeOf((*MockHub)(nil).NatRule), id)
}

// NatRuleCreate mocks base method
func (m *MockHub) NatRuleCreate(natRule *homehub_client.NatRule) error {
	ret := m.ctrl.Call(m, "NatRuleCreate", natRule)
	ret0, _ := ret[0].(error)
	return ret0
}

// NatRuleCreate indicates an expected call of NatRuleCreate
func (mr *MockHubMockRecorder) NatRuleCreate(natRule interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatRuleCreate", reflect.TypeOf((*MockHub)(nil).NatRuleCreate), natRule)
}

// NatRuleDelete mocks base method
func (m *MockHub) NatRuleDelete(id int) error {
	ret := m.ctrl.Call(m, "NatRuleDelete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// NatRuleDelete indicates an expected call of NatRuleDelete
func (mr *MockHubMockRecorder) NatRuleDelete(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatRuleDelete", reflect.TypeOf((*MockHub)(nil).NatRuleDelete), id)
}

// NatRuleUpdate mocks base method
func (m *MockHub) NatRuleUpdate(natRule homehub_client.NatRule) error {
	ret := m.ctrl.Call(m, "NatRuleUpdate", natRule)
	ret0, _ := ret[0].(error)
	return ret0
}

// NatRuleUpdate indicates an expected call of NatRuleUpdate
func (mr *MockHubMockRecorder) NatRuleUpdate(natRule interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatRuleUpdate", reflect.TypeOf((*MockHub)(nil).NatRuleUpdate), natRule)
}

// PublicIPAddress mocks base method
func (m *MockHub) PublicIPAddress() (string, error) {
	ret := m.ctrl.Call(m, "PublicIPAddress")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicIPAddress indicates an expected call of PublicIPAddress
func (mr *MockHubMockRecorder) PublicIPAddress() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicIPAddress", reflect.TypeOf((*MockHub)(nil).PublicIPAddress))
}

// PublicSubnetMask mocks base method
func (m *MockHub) PublicSubnetMask() (string, error) {
	ret := m.ctrl.Call(m, "PublicSubnetMask")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicSubnetMask indicates an expected call of PublicSubnetMask
func (mr *MockHubMockRecorder) PublicSubnetMask() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicSubnetMask", reflect.TypeOf((*MockHub)(nil).PublicSubnetMask))
}

// Reboot mocks base method
func (m *MockHub) Reboot() error {
	ret := m.ctrl.Call(m, "Reboot")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reboot indicates an expected call of Reboot
func (mr *MockHubMockRecorder) Reboot() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reboot", reflect.TypeOf((*MockHub)(nil).Reboot))
}

// SambaIP mocks base method
func (m *MockHub) SambaIP() (string, error) {
	ret := m.ctrl.Call(m, "SambaIP")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SambaIP indicates an expected call of SambaIP
func (mr *MockHubMockRecorder) SambaIP() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SambaIP", reflect.TypeOf((*MockHub)(nil).SambaIP))
}

// SambaHost mocks base method
func (m *MockHub) SambaHost() (string, error) {
	ret := m.ctrl.Call(m, "SambaHost")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SambaHost indicates an expected call of SambaHost
func (mr *MockHubMockRecorder) SambaHost() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SambaHost", reflect.TypeOf((*MockHub)(nil).SambaHost))
}

// SerialNumber mocks base method
func (m *MockHub) SerialNumber() (string, error) {
	ret := m.ctrl.Call(m, "SerialNumber")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SerialNumber indicates an expected call of SerialNumber
func (mr *MockHubMockRecorder) SerialNumber() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SerialNumber", reflect.TypeOf((*MockHub)(nil).SerialNumber))
}

// SoftwareVersion mocks base method
func (m *MockHub) SoftwareVersion() (string, error) {
	ret := m.ctrl.Call(m, "SoftwareVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SoftwareVersion indicates an expected call of SoftwareVersion
func (mr *MockHubMockRecorder) SoftwareVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftwareVersion", reflect.TypeOf((*MockHub)(nil).SoftwareVersion))
}

// UpstreamSyncSpeed mocks base method
func (m *MockHub) UpstreamSyncSpeed() (int, error) {
	ret := m.ctrl.Call(m, "UpstreamSyncSpeed")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpstreamSyncSpeed indicates an expected call of UpstreamSyncSpeed
func (mr *MockHubMockRecorder) UpstreamSyncSpeed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpstreamSyncSpeed", reflect.TypeOf((*MockHub)(nil).UpstreamSyncSpeed))
}

// Version mocks base method
func (m *MockHub) Version() (string, error) {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockHubMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockHub)(nil).Version))
}

// WiFiFrequency24Ghz mocks base method
func (m *MockHub) WiFiFrequency24Ghz() (*homehub_client.WiFiFrequency, error) {
	ret := m.ctrl.Call(m, "WiFiFrequency24Ghz")
	ret0, _ := ret[0].(*homehub_client.WiFiFrequency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WiFiFrequency24Ghz indicates an expected call of WiFiFrequency24Ghz
func (mr *MockHubMockRecorder) WiFiFrequency24Ghz() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WiFiFrequency24Ghz", reflect.TypeOf((*MockHub)(nil).WiFiFrequency24Ghz))
}

// WiFiFrequency24GhzChannelSet mocks base method
func (m *MockHub) WiFiFrequency24GhzChannelSet(channel int) error {
	ret := m.ctrl.Call(m, "WiFiFrequency24GhzChannelSet", channel)
	ret0, _ := ret[0].(error)
	return ret0
}

// WiFiFrequency24GhzChannelSet indicates an expected call of WiFiFrequency24GhzChannelSet
func (mr *MockHubMockRecorder) WiFiFrequency24GhzChannelSet(channel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WiFiFrequency24GhzChannelSet", reflect.TypeOf((*MockHub)(nil).WiFiFrequency24GhzChannelSet), channel)
}

// WiFiFrequency5Ghz mocks base method
func (m *MockHub) WiFiFrequency5Ghz() (*homehub_client.WiFiFrequency, error) {
	ret := m.ctrl.Call(m, "WiFiFrequency5Ghz")
	ret0, _ := ret[0].(*homehub_client.WiFiFrequency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WiFiFrequency5Ghz indicates an expected call of WiFiFrequency5Ghz
func (mr *MockHubMockRecorder) WiFiFrequency5Ghz() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WiFiFrequency5Ghz", reflect.TypeOf((*MockHub)(nil).WiFiFrequency5Ghz))
}

// WiFiFrequency5GhzChannelSet mocks base method
func (m *MockHub) WiFiFrequency5GhzChannelSet(channel int) error {
	ret := m.ctrl.Call(m, "WiFiFrequency5GhzChannelSet", channel)
	ret0, _ := ret[0].(error)
	return ret0
}

// WiFiFrequency5GhzChannelSet indicates an expected call of WiFiFrequency5GhzChannelSet
func (mr *MockHubMockRecorder) WiFiFrequency5GhzChannelSet(channel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WiFiFrequency5GhzChannelSet", reflect.TypeOf((*MockHub)(nil).WiFiFrequency5GhzChannelSet), channel)
}

// WiFiSecurityMode mocks base method
func (m *MockHub) WiFiSecurityMode() (string, error) {
	ret := m.ctrl.Call(m, "WiFiSecurityMode")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WiFiSecurityMode indicates an expected call of WiFiSecurityMode
func (mr *MockHubMockRecorder) WiFiSecurityMode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WiFiSecurityMode", reflect.TypeOf((*MockHub)(nil).WiFiSecurityMode))
}

// WiFiSSID mocks base method
func (m *MockHub) WiFiSSID() (string, error) {
	ret := m.ctrl.Call(m, "WiFiSSID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WiFiSSID indicates an expected call of WiFiSSID
func (mr *MockHubMockRecorder) WiFiSSID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WiFiSSID", reflect.TypeOf((*MockHub)(nil).WiFiSSID))
}
