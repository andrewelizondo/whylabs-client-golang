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

// checks if the CloseSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloseSessionResponse{}

// CloseSessionResponse Response for closing a session
type CloseSessionResponse struct {
	// URL where the session can be viewed.
	WhylabsUrl string `json:"whylabsUrl"`
}

// NewCloseSessionResponse instantiates a new CloseSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseSessionResponse(whylabsUrl string) *CloseSessionResponse {
	this := CloseSessionResponse{}
	this.WhylabsUrl = whylabsUrl
	return &this
}

// NewCloseSessionResponseWithDefaults instantiates a new CloseSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseSessionResponseWithDefaults() *CloseSessionResponse {
	this := CloseSessionResponse{}
	return &this
}

// GetWhylabsUrl returns the WhylabsUrl field value
func (o *CloseSessionResponse) GetWhylabsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WhylabsUrl
}

// GetWhylabsUrlOk returns a tuple with the WhylabsUrl field value
// and a boolean to check if the value has been set.
func (o *CloseSessionResponse) GetWhylabsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhylabsUrl, true
}

// SetWhylabsUrl sets field value
func (o *CloseSessionResponse) SetWhylabsUrl(v string) {
	o.WhylabsUrl = v
}

func (o CloseSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloseSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["whylabsUrl"] = o.WhylabsUrl
	return toSerialize, nil
}

type NullableCloseSessionResponse struct {
	value *CloseSessionResponse
	isSet bool
}

func (v NullableCloseSessionResponse) Get() *CloseSessionResponse {
	return v.value
}

func (v *NullableCloseSessionResponse) Set(val *CloseSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseSessionResponse(val *CloseSessionResponse) *NullableCloseSessionResponse {
	return &NullableCloseSessionResponse{value: val, isSet: true}
}

func (v NullableCloseSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


