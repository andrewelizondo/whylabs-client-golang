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

// checks if the DeleteAnalyzerResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAnalyzerResultsResponse{}

// DeleteAnalyzerResultsResponse struct for DeleteAnalyzerResultsResponse
type DeleteAnalyzerResultsResponse struct {
	Id *string `json:"id,omitempty"`
}

// NewDeleteAnalyzerResultsResponse instantiates a new DeleteAnalyzerResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAnalyzerResultsResponse() *DeleteAnalyzerResultsResponse {
	this := DeleteAnalyzerResultsResponse{}
	return &this
}

// NewDeleteAnalyzerResultsResponseWithDefaults instantiates a new DeleteAnalyzerResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAnalyzerResultsResponseWithDefaults() *DeleteAnalyzerResultsResponse {
	this := DeleteAnalyzerResultsResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteAnalyzerResultsResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAnalyzerResultsResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteAnalyzerResultsResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteAnalyzerResultsResponse) SetId(v string) {
	o.Id = &v
}

func (o DeleteAnalyzerResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAnalyzerResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeleteAnalyzerResultsResponse struct {
	value *DeleteAnalyzerResultsResponse
	isSet bool
}

func (v NullableDeleteAnalyzerResultsResponse) Get() *DeleteAnalyzerResultsResponse {
	return v.value
}

func (v *NullableDeleteAnalyzerResultsResponse) Set(val *DeleteAnalyzerResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAnalyzerResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAnalyzerResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAnalyzerResultsResponse(val *DeleteAnalyzerResultsResponse) *NullableDeleteAnalyzerResultsResponse {
	return &NullableDeleteAnalyzerResultsResponse{value: val, isSet: true}
}

func (v NullableDeleteAnalyzerResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAnalyzerResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


