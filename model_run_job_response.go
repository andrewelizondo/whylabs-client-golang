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

// checks if the RunJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunJobResponse{}

// RunJobResponse struct for RunJobResponse
type RunJobResponse struct {
	RunId int64 `json:"runId"`
}

// NewRunJobResponse instantiates a new RunJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunJobResponse(runId int64) *RunJobResponse {
	this := RunJobResponse{}
	this.RunId = runId
	return &this
}

// NewRunJobResponseWithDefaults instantiates a new RunJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunJobResponseWithDefaults() *RunJobResponse {
	this := RunJobResponse{}
	return &this
}

// GetRunId returns the RunId field value
func (o *RunJobResponse) GetRunId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *RunJobResponse) GetRunIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *RunJobResponse) SetRunId(v int64) {
	o.RunId = v
}

func (o RunJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["runId"] = o.RunId
	return toSerialize, nil
}

type NullableRunJobResponse struct {
	value *RunJobResponse
	isSet bool
}

func (v NullableRunJobResponse) Get() *RunJobResponse {
	return v.value
}

func (v *NullableRunJobResponse) Set(val *RunJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRunJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRunJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunJobResponse(val *RunJobResponse) *NullableRunJobResponse {
	return &NullableRunJobResponse{value: val, isSet: true}
}

func (v NullableRunJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


