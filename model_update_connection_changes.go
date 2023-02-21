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

// checks if the UpdateConnectionChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateConnectionChanges{}

// UpdateConnectionChanges struct for UpdateConnectionChanges
type UpdateConnectionChanges struct {
	OrgId NullableString `json:"orgId,omitempty"`
	Connected NullableBool `json:"connected,omitempty"`
	Demo NullableBool `json:"demo,omitempty"`
}

// NewUpdateConnectionChanges instantiates a new UpdateConnectionChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConnectionChanges() *UpdateConnectionChanges {
	this := UpdateConnectionChanges{}
	return &this
}

// NewUpdateConnectionChangesWithDefaults instantiates a new UpdateConnectionChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConnectionChangesWithDefaults() *UpdateConnectionChanges {
	this := UpdateConnectionChanges{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateConnectionChanges) GetOrgId() string {
	if o == nil || isNil(o.OrgId.Get()) {
		var ret string
		return ret
	}
	return *o.OrgId.Get()
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateConnectionChanges) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgId.Get(), o.OrgId.IsSet()
}

// HasOrgId returns a boolean if a field has been set.
func (o *UpdateConnectionChanges) HasOrgId() bool {
	if o != nil && o.OrgId.IsSet() {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given NullableString and assigns it to the OrgId field.
func (o *UpdateConnectionChanges) SetOrgId(v string) {
	o.OrgId.Set(&v)
}
// SetOrgIdNil sets the value for OrgId to be an explicit nil
func (o *UpdateConnectionChanges) SetOrgIdNil() {
	o.OrgId.Set(nil)
}

// UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
func (o *UpdateConnectionChanges) UnsetOrgId() {
	o.OrgId.Unset()
}

// GetConnected returns the Connected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateConnectionChanges) GetConnected() bool {
	if o == nil || isNil(o.Connected.Get()) {
		var ret bool
		return ret
	}
	return *o.Connected.Get()
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateConnectionChanges) GetConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connected.Get(), o.Connected.IsSet()
}

// HasConnected returns a boolean if a field has been set.
func (o *UpdateConnectionChanges) HasConnected() bool {
	if o != nil && o.Connected.IsSet() {
		return true
	}

	return false
}

// SetConnected gets a reference to the given NullableBool and assigns it to the Connected field.
func (o *UpdateConnectionChanges) SetConnected(v bool) {
	o.Connected.Set(&v)
}
// SetConnectedNil sets the value for Connected to be an explicit nil
func (o *UpdateConnectionChanges) SetConnectedNil() {
	o.Connected.Set(nil)
}

// UnsetConnected ensures that no value is present for Connected, not even an explicit nil
func (o *UpdateConnectionChanges) UnsetConnected() {
	o.Connected.Unset()
}

// GetDemo returns the Demo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateConnectionChanges) GetDemo() bool {
	if o == nil || isNil(o.Demo.Get()) {
		var ret bool
		return ret
	}
	return *o.Demo.Get()
}

// GetDemoOk returns a tuple with the Demo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateConnectionChanges) GetDemoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Demo.Get(), o.Demo.IsSet()
}

// HasDemo returns a boolean if a field has been set.
func (o *UpdateConnectionChanges) HasDemo() bool {
	if o != nil && o.Demo.IsSet() {
		return true
	}

	return false
}

// SetDemo gets a reference to the given NullableBool and assigns it to the Demo field.
func (o *UpdateConnectionChanges) SetDemo(v bool) {
	o.Demo.Set(&v)
}
// SetDemoNil sets the value for Demo to be an explicit nil
func (o *UpdateConnectionChanges) SetDemoNil() {
	o.Demo.Set(nil)
}

// UnsetDemo ensures that no value is present for Demo, not even an explicit nil
func (o *UpdateConnectionChanges) UnsetDemo() {
	o.Demo.Unset()
}

func (o UpdateConnectionChanges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateConnectionChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgId.IsSet() {
		toSerialize["orgId"] = o.OrgId.Get()
	}
	if o.Connected.IsSet() {
		toSerialize["connected"] = o.Connected.Get()
	}
	if o.Demo.IsSet() {
		toSerialize["demo"] = o.Demo.Get()
	}
	return toSerialize, nil
}

type NullableUpdateConnectionChanges struct {
	value *UpdateConnectionChanges
	isSet bool
}

func (v NullableUpdateConnectionChanges) Get() *UpdateConnectionChanges {
	return v.value
}

func (v *NullableUpdateConnectionChanges) Set(val *UpdateConnectionChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConnectionChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConnectionChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConnectionChanges(val *UpdateConnectionChanges) *NullableUpdateConnectionChanges {
	return &NullableUpdateConnectionChanges{value: val, isSet: true}
}

func (v NullableUpdateConnectionChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConnectionChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


