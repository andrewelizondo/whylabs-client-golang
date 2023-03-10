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

// checks if the GetMonitorConfigV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMonitorConfigV2Response{}

// GetMonitorConfigV2Response struct for GetMonitorConfigV2Response
type GetMonitorConfigV2Response struct {
	Config *DatasetRequestMonitorConfig `json:"config,omitempty"`
}

// NewGetMonitorConfigV2Response instantiates a new GetMonitorConfigV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMonitorConfigV2Response() *GetMonitorConfigV2Response {
	this := GetMonitorConfigV2Response{}
	return &this
}

// NewGetMonitorConfigV2ResponseWithDefaults instantiates a new GetMonitorConfigV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMonitorConfigV2ResponseWithDefaults() *GetMonitorConfigV2Response {
	this := GetMonitorConfigV2Response{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GetMonitorConfigV2Response) GetConfig() DatasetRequestMonitorConfig {
	if o == nil || isNil(o.Config) {
		var ret DatasetRequestMonitorConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMonitorConfigV2Response) GetConfigOk() (*DatasetRequestMonitorConfig, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *GetMonitorConfigV2Response) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DatasetRequestMonitorConfig and assigns it to the Config field.
func (o *GetMonitorConfigV2Response) SetConfig(v DatasetRequestMonitorConfig) {
	o.Config = &v
}

func (o GetMonitorConfigV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMonitorConfigV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableGetMonitorConfigV2Response struct {
	value *GetMonitorConfigV2Response
	isSet bool
}

func (v NullableGetMonitorConfigV2Response) Get() *GetMonitorConfigV2Response {
	return v.value
}

func (v *NullableGetMonitorConfigV2Response) Set(val *GetMonitorConfigV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMonitorConfigV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMonitorConfigV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMonitorConfigV2Response(val *GetMonitorConfigV2Response) *NullableGetMonitorConfigV2Response {
	return &NullableGetMonitorConfigV2Response{value: val, isSet: true}
}

func (v NullableGetMonitorConfigV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMonitorConfigV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


