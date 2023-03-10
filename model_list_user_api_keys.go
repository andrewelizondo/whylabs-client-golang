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

// checks if the ListUserApiKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserApiKeys{}

// ListUserApiKeys Response for listing API key metadata
type ListUserApiKeys struct {
	// A list of all known API key metadata
	Items []UserApiKey `json:"items"`
}

// NewListUserApiKeys instantiates a new ListUserApiKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserApiKeys(items []UserApiKey) *ListUserApiKeys {
	this := ListUserApiKeys{}
	this.Items = items
	return &this
}

// NewListUserApiKeysWithDefaults instantiates a new ListUserApiKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserApiKeysWithDefaults() *ListUserApiKeys {
	this := ListUserApiKeys{}
	return &this
}

// GetItems returns the Items field value
func (o *ListUserApiKeys) GetItems() []UserApiKey {
	if o == nil {
		var ret []UserApiKey
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListUserApiKeys) GetItemsOk() ([]UserApiKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListUserApiKeys) SetItems(v []UserApiKey) {
	o.Items = v
}

func (o ListUserApiKeys) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserApiKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListUserApiKeys struct {
	value *ListUserApiKeys
	isSet bool
}

func (v NullableListUserApiKeys) Get() *ListUserApiKeys {
	return v.value
}

func (v *NullableListUserApiKeys) Set(val *ListUserApiKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserApiKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserApiKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserApiKeys(val *ListUserApiKeys) *NullableListUserApiKeys {
	return &NullableListUserApiKeys{value: val, isSet: true}
}

func (v NullableListUserApiKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserApiKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


