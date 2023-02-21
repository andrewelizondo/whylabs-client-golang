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

// checks if the RegisterDatabricksConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterDatabricksConnectionResponse{}

// RegisterDatabricksConnectionResponse struct for RegisterDatabricksConnectionResponse
type RegisterDatabricksConnectionResponse struct {
	Id string `json:"id"`
}

// NewRegisterDatabricksConnectionResponse instantiates a new RegisterDatabricksConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDatabricksConnectionResponse(id string) *RegisterDatabricksConnectionResponse {
	this := RegisterDatabricksConnectionResponse{}
	this.Id = id
	return &this
}

// NewRegisterDatabricksConnectionResponseWithDefaults instantiates a new RegisterDatabricksConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDatabricksConnectionResponseWithDefaults() *RegisterDatabricksConnectionResponse {
	this := RegisterDatabricksConnectionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *RegisterDatabricksConnectionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegisterDatabricksConnectionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegisterDatabricksConnectionResponse) SetId(v string) {
	o.Id = v
}

func (o RegisterDatabricksConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterDatabricksConnectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableRegisterDatabricksConnectionResponse struct {
	value *RegisterDatabricksConnectionResponse
	isSet bool
}

func (v NullableRegisterDatabricksConnectionResponse) Get() *RegisterDatabricksConnectionResponse {
	return v.value
}

func (v *NullableRegisterDatabricksConnectionResponse) Set(val *RegisterDatabricksConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDatabricksConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDatabricksConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDatabricksConnectionResponse(val *RegisterDatabricksConnectionResponse) *NullableRegisterDatabricksConnectionResponse {
	return &NullableRegisterDatabricksConnectionResponse{value: val, isSet: true}
}

func (v NullableRegisterDatabricksConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDatabricksConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


