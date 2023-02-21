/*
WhyLabs Songbird

WhyLabs API that enables end-to-end AI observability

API version: 0.1
Contact: support@whylabs.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RegisterDatabricksConnectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterDatabricksConnectionRequest{}

// RegisterDatabricksConnectionRequest struct for RegisterDatabricksConnectionRequest
type RegisterDatabricksConnectionRequest struct {
	Email string `json:"email"`
	ConnectionEstablished *bool `json:"connectionEstablished,omitempty"`
	AccessToken string `json:"accessToken"`
	Hostname string `json:"hostname"`
	Port int32 `json:"port"`
	WorkspaceUrl string `json:"workspaceUrl"`
	ConnectionId string `json:"connectionId"`
	WorkspaceId string `json:"workspaceId"`
	Demo bool `json:"demo"`
	CloudProvider string `json:"cloudProvider"`
	FreeTrial *bool `json:"freeTrial,omitempty"`
}

// NewRegisterDatabricksConnectionRequest instantiates a new RegisterDatabricksConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDatabricksConnectionRequest(email string, accessToken string, hostname string, port int32, workspaceUrl string, connectionId string, workspaceId string, demo bool, cloudProvider string) *RegisterDatabricksConnectionRequest {
	this := RegisterDatabricksConnectionRequest{}
	this.Email = email
	this.AccessToken = accessToken
	this.Hostname = hostname
	this.Port = port
	this.WorkspaceUrl = workspaceUrl
	this.ConnectionId = connectionId
	this.WorkspaceId = workspaceId
	this.Demo = demo
	this.CloudProvider = cloudProvider
	return &this
}

// NewRegisterDatabricksConnectionRequestWithDefaults instantiates a new RegisterDatabricksConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDatabricksConnectionRequestWithDefaults() *RegisterDatabricksConnectionRequest {
	this := RegisterDatabricksConnectionRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *RegisterDatabricksConnectionRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegisterDatabricksConnectionRequest) SetEmail(v string) {
	o.Email = v
}

// GetConnectionEstablished returns the ConnectionEstablished field value if set, zero value otherwise.
func (o *RegisterDatabricksConnectionRequest) GetConnectionEstablished() bool {
	if o == nil || isNil(o.ConnectionEstablished) {
		var ret bool
		return ret
	}
	return *o.ConnectionEstablished
}

// GetConnectionEstablishedOk returns a tuple with the ConnectionEstablished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetConnectionEstablishedOk() (*bool, bool) {
	if o == nil || isNil(o.ConnectionEstablished) {
		return nil, false
	}
	return o.ConnectionEstablished, true
}

// HasConnectionEstablished returns a boolean if a field has been set.
func (o *RegisterDatabricksConnectionRequest) HasConnectionEstablished() bool {
	if o != nil && !isNil(o.ConnectionEstablished) {
		return true
	}

	return false
}

// SetConnectionEstablished gets a reference to the given bool and assigns it to the ConnectionEstablished field.
func (o *RegisterDatabricksConnectionRequest) SetConnectionEstablished(v bool) {
	o.ConnectionEstablished = &v
}

// GetAccessToken returns the AccessToken field value
func (o *RegisterDatabricksConnectionRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *RegisterDatabricksConnectionRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetHostname returns the Hostname field value
func (o *RegisterDatabricksConnectionRequest) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *RegisterDatabricksConnectionRequest) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *RegisterDatabricksConnectionRequest) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *RegisterDatabricksConnectionRequest) SetPort(v int32) {
	o.Port = v
}

// GetWorkspaceUrl returns the WorkspaceUrl field value
func (o *RegisterDatabricksConnectionRequest) GetWorkspaceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceUrl
}

// GetWorkspaceUrlOk returns a tuple with the WorkspaceUrl field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetWorkspaceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceUrl, true
}

// SetWorkspaceUrl sets field value
func (o *RegisterDatabricksConnectionRequest) SetWorkspaceUrl(v string) {
	o.WorkspaceUrl = v
}

// GetConnectionId returns the ConnectionId field value
func (o *RegisterDatabricksConnectionRequest) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *RegisterDatabricksConnectionRequest) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *RegisterDatabricksConnectionRequest) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *RegisterDatabricksConnectionRequest) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetDemo returns the Demo field value
func (o *RegisterDatabricksConnectionRequest) GetDemo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Demo
}

// GetDemoOk returns a tuple with the Demo field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetDemoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Demo, true
}

// SetDemo sets field value
func (o *RegisterDatabricksConnectionRequest) SetDemo(v bool) {
	o.Demo = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *RegisterDatabricksConnectionRequest) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *RegisterDatabricksConnectionRequest) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetFreeTrial returns the FreeTrial field value if set, zero value otherwise.
func (o *RegisterDatabricksConnectionRequest) GetFreeTrial() bool {
	if o == nil || isNil(o.FreeTrial) {
		var ret bool
		return ret
	}
	return *o.FreeTrial
}

// GetFreeTrialOk returns a tuple with the FreeTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionRequest) GetFreeTrialOk() (*bool, bool) {
	if o == nil || isNil(o.FreeTrial) {
		return nil, false
	}
	return o.FreeTrial, true
}

// HasFreeTrial returns a boolean if a field has been set.
func (o *RegisterDatabricksConnectionRequest) HasFreeTrial() bool {
	if o != nil && !isNil(o.FreeTrial) {
		return true
	}

	return false
}

// SetFreeTrial gets a reference to the given bool and assigns it to the FreeTrial field.
func (o *RegisterDatabricksConnectionRequest) SetFreeTrial(v bool) {
	o.FreeTrial = &v
}

func (o RegisterDatabricksConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterDatabricksConnectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !isNil(o.ConnectionEstablished) {
		toSerialize["connectionEstablished"] = o.ConnectionEstablished
	}
	toSerialize["accessToken"] = o.AccessToken
	toSerialize["hostname"] = o.Hostname
	toSerialize["port"] = o.Port
	toSerialize["workspaceUrl"] = o.WorkspaceUrl
	toSerialize["connectionId"] = o.ConnectionId
	toSerialize["workspaceId"] = o.WorkspaceId
	toSerialize["demo"] = o.Demo
	toSerialize["cloudProvider"] = o.CloudProvider
	if !isNil(o.FreeTrial) {
		toSerialize["freeTrial"] = o.FreeTrial
	}
	return toSerialize, nil
}

type NullableRegisterDatabricksConnectionRequest struct {
	value *RegisterDatabricksConnectionRequest
	isSet bool
}

func (v NullableRegisterDatabricksConnectionRequest) Get() *RegisterDatabricksConnectionRequest {
	return v.value
}

func (v *NullableRegisterDatabricksConnectionRequest) Set(val *RegisterDatabricksConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDatabricksConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDatabricksConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDatabricksConnectionRequest(val *RegisterDatabricksConnectionRequest) *NullableRegisterDatabricksConnectionRequest {
	return &NullableRegisterDatabricksConnectionRequest{value: val, isSet: true}
}

func (v NullableRegisterDatabricksConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDatabricksConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


