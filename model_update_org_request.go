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

// checks if the UpdateOrgRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrgRequest{}

// UpdateOrgRequest struct for UpdateOrgRequest
type UpdateOrgRequest struct {
	// The unique ID of an organization. If an organization with this ID does not exist, this method will throw an exception.
	OrgId *string `json:"org_id,omitempty"`
}

// NewUpdateOrgRequest instantiates a new UpdateOrgRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrgRequest() *UpdateOrgRequest {
	this := UpdateOrgRequest{}
	return &this
}

// NewUpdateOrgRequestWithDefaults instantiates a new UpdateOrgRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgRequestWithDefaults() *UpdateOrgRequest {
	this := UpdateOrgRequest{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetOrgId() string {
	if o == nil || isNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetOrgIdOk() (*string, bool) {
	if o == nil || isNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasOrgId() bool {
	if o != nil && !isNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *UpdateOrgRequest) SetOrgId(v string) {
	o.OrgId = &v
}

func (o UpdateOrgRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrgRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	return toSerialize, nil
}

type NullableUpdateOrgRequest struct {
	value *UpdateOrgRequest
	isSet bool
}

func (v NullableUpdateOrgRequest) Get() *UpdateOrgRequest {
	return v.value
}

func (v *NullableUpdateOrgRequest) Set(val *UpdateOrgRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrgRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrgRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrgRequest(val *UpdateOrgRequest) *NullableUpdateOrgRequest {
	return &NullableUpdateOrgRequest{value: val, isSet: true}
}

func (v NullableUpdateOrgRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrgRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


