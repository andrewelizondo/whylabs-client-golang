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

// checks if the UniqueValuesMonitorRequestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueValuesMonitorRequestConfig{}

// UniqueValuesMonitorRequestConfig struct for UniqueValuesMonitorRequestConfig
type UniqueValuesMonitorRequestConfig struct {
	Enable NullableBool `json:"enable,omitempty"`
	MinRecordCount NullableInt32 `json:"minRecordCount,omitempty"`
	MinThreshold NullableFloat64 `json:"minThreshold,omitempty"`
	MaxThreshold NullableFloat64 `json:"maxThreshold,omitempty"`
}

// NewUniqueValuesMonitorRequestConfig instantiates a new UniqueValuesMonitorRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueValuesMonitorRequestConfig() *UniqueValuesMonitorRequestConfig {
	this := UniqueValuesMonitorRequestConfig{}
	return &this
}

// NewUniqueValuesMonitorRequestConfigWithDefaults instantiates a new UniqueValuesMonitorRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueValuesMonitorRequestConfigWithDefaults() *UniqueValuesMonitorRequestConfig {
	this := UniqueValuesMonitorRequestConfig{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UniqueValuesMonitorRequestConfig) GetEnable() bool {
	if o == nil || isNil(o.Enable.Get()) {
		var ret bool
		return ret
	}
	return *o.Enable.Get()
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UniqueValuesMonitorRequestConfig) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enable.Get(), o.Enable.IsSet()
}

// HasEnable returns a boolean if a field has been set.
func (o *UniqueValuesMonitorRequestConfig) HasEnable() bool {
	if o != nil && o.Enable.IsSet() {
		return true
	}

	return false
}

// SetEnable gets a reference to the given NullableBool and assigns it to the Enable field.
func (o *UniqueValuesMonitorRequestConfig) SetEnable(v bool) {
	o.Enable.Set(&v)
}
// SetEnableNil sets the value for Enable to be an explicit nil
func (o *UniqueValuesMonitorRequestConfig) SetEnableNil() {
	o.Enable.Set(nil)
}

// UnsetEnable ensures that no value is present for Enable, not even an explicit nil
func (o *UniqueValuesMonitorRequestConfig) UnsetEnable() {
	o.Enable.Unset()
}

// GetMinRecordCount returns the MinRecordCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UniqueValuesMonitorRequestConfig) GetMinRecordCount() int32 {
	if o == nil || isNil(o.MinRecordCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MinRecordCount.Get()
}

// GetMinRecordCountOk returns a tuple with the MinRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UniqueValuesMonitorRequestConfig) GetMinRecordCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinRecordCount.Get(), o.MinRecordCount.IsSet()
}

// HasMinRecordCount returns a boolean if a field has been set.
func (o *UniqueValuesMonitorRequestConfig) HasMinRecordCount() bool {
	if o != nil && o.MinRecordCount.IsSet() {
		return true
	}

	return false
}

// SetMinRecordCount gets a reference to the given NullableInt32 and assigns it to the MinRecordCount field.
func (o *UniqueValuesMonitorRequestConfig) SetMinRecordCount(v int32) {
	o.MinRecordCount.Set(&v)
}
// SetMinRecordCountNil sets the value for MinRecordCount to be an explicit nil
func (o *UniqueValuesMonitorRequestConfig) SetMinRecordCountNil() {
	o.MinRecordCount.Set(nil)
}

// UnsetMinRecordCount ensures that no value is present for MinRecordCount, not even an explicit nil
func (o *UniqueValuesMonitorRequestConfig) UnsetMinRecordCount() {
	o.MinRecordCount.Unset()
}

// GetMinThreshold returns the MinThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UniqueValuesMonitorRequestConfig) GetMinThreshold() float64 {
	if o == nil || isNil(o.MinThreshold.Get()) {
		var ret float64
		return ret
	}
	return *o.MinThreshold.Get()
}

// GetMinThresholdOk returns a tuple with the MinThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UniqueValuesMonitorRequestConfig) GetMinThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinThreshold.Get(), o.MinThreshold.IsSet()
}

// HasMinThreshold returns a boolean if a field has been set.
func (o *UniqueValuesMonitorRequestConfig) HasMinThreshold() bool {
	if o != nil && o.MinThreshold.IsSet() {
		return true
	}

	return false
}

// SetMinThreshold gets a reference to the given NullableFloat64 and assigns it to the MinThreshold field.
func (o *UniqueValuesMonitorRequestConfig) SetMinThreshold(v float64) {
	o.MinThreshold.Set(&v)
}
// SetMinThresholdNil sets the value for MinThreshold to be an explicit nil
func (o *UniqueValuesMonitorRequestConfig) SetMinThresholdNil() {
	o.MinThreshold.Set(nil)
}

// UnsetMinThreshold ensures that no value is present for MinThreshold, not even an explicit nil
func (o *UniqueValuesMonitorRequestConfig) UnsetMinThreshold() {
	o.MinThreshold.Unset()
}

// GetMaxThreshold returns the MaxThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UniqueValuesMonitorRequestConfig) GetMaxThreshold() float64 {
	if o == nil || isNil(o.MaxThreshold.Get()) {
		var ret float64
		return ret
	}
	return *o.MaxThreshold.Get()
}

// GetMaxThresholdOk returns a tuple with the MaxThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UniqueValuesMonitorRequestConfig) GetMaxThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxThreshold.Get(), o.MaxThreshold.IsSet()
}

// HasMaxThreshold returns a boolean if a field has been set.
func (o *UniqueValuesMonitorRequestConfig) HasMaxThreshold() bool {
	if o != nil && o.MaxThreshold.IsSet() {
		return true
	}

	return false
}

// SetMaxThreshold gets a reference to the given NullableFloat64 and assigns it to the MaxThreshold field.
func (o *UniqueValuesMonitorRequestConfig) SetMaxThreshold(v float64) {
	o.MaxThreshold.Set(&v)
}
// SetMaxThresholdNil sets the value for MaxThreshold to be an explicit nil
func (o *UniqueValuesMonitorRequestConfig) SetMaxThresholdNil() {
	o.MaxThreshold.Set(nil)
}

// UnsetMaxThreshold ensures that no value is present for MaxThreshold, not even an explicit nil
func (o *UniqueValuesMonitorRequestConfig) UnsetMaxThreshold() {
	o.MaxThreshold.Unset()
}

func (o UniqueValuesMonitorRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueValuesMonitorRequestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable.IsSet() {
		toSerialize["enable"] = o.Enable.Get()
	}
	if o.MinRecordCount.IsSet() {
		toSerialize["minRecordCount"] = o.MinRecordCount.Get()
	}
	if o.MinThreshold.IsSet() {
		toSerialize["minThreshold"] = o.MinThreshold.Get()
	}
	if o.MaxThreshold.IsSet() {
		toSerialize["maxThreshold"] = o.MaxThreshold.Get()
	}
	return toSerialize, nil
}

type NullableUniqueValuesMonitorRequestConfig struct {
	value *UniqueValuesMonitorRequestConfig
	isSet bool
}

func (v NullableUniqueValuesMonitorRequestConfig) Get() *UniqueValuesMonitorRequestConfig {
	return v.value
}

func (v *NullableUniqueValuesMonitorRequestConfig) Set(val *UniqueValuesMonitorRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueValuesMonitorRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueValuesMonitorRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueValuesMonitorRequestConfig(val *UniqueValuesMonitorRequestConfig) *NullableUniqueValuesMonitorRequestConfig {
	return &NullableUniqueValuesMonitorRequestConfig{value: val, isSet: true}
}

func (v NullableUniqueValuesMonitorRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueValuesMonitorRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


