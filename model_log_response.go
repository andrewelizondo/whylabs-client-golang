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

// checks if the LogResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogResponse{}

// LogResponse Response payload for UploadDatasetProfile.
type LogResponse struct {
	Id *string `json:"id,omitempty"`
	ModelId *string `json:"modelId,omitempty"`
	UploadTimestamp *int64 `json:"uploadTimestamp,omitempty"`
}

// NewLogResponse instantiates a new LogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogResponse() *LogResponse {
	this := LogResponse{}
	return &this
}

// NewLogResponseWithDefaults instantiates a new LogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogResponseWithDefaults() *LogResponse {
	this := LogResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogResponse) SetId(v string) {
	o.Id = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *LogResponse) GetModelId() string {
	if o == nil || isNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetModelIdOk() (*string, bool) {
	if o == nil || isNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *LogResponse) HasModelId() bool {
	if o != nil && !isNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *LogResponse) SetModelId(v string) {
	o.ModelId = &v
}

// GetUploadTimestamp returns the UploadTimestamp field value if set, zero value otherwise.
func (o *LogResponse) GetUploadTimestamp() int64 {
	if o == nil || isNil(o.UploadTimestamp) {
		var ret int64
		return ret
	}
	return *o.UploadTimestamp
}

// GetUploadTimestampOk returns a tuple with the UploadTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetUploadTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.UploadTimestamp) {
		return nil, false
	}
	return o.UploadTimestamp, true
}

// HasUploadTimestamp returns a boolean if a field has been set.
func (o *LogResponse) HasUploadTimestamp() bool {
	if o != nil && !isNil(o.UploadTimestamp) {
		return true
	}

	return false
}

// SetUploadTimestamp gets a reference to the given int64 and assigns it to the UploadTimestamp field.
func (o *LogResponse) SetUploadTimestamp(v int64) {
	o.UploadTimestamp = &v
}

func (o LogResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	if !isNil(o.UploadTimestamp) {
		toSerialize["uploadTimestamp"] = o.UploadTimestamp
	}
	return toSerialize, nil
}

type NullableLogResponse struct {
	value *LogResponse
	isSet bool
}

func (v NullableLogResponse) Get() *LogResponse {
	return v.value
}

func (v *NullableLogResponse) Set(val *LogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogResponse(val *LogResponse) *NullableLogResponse {
	return &NullableLogResponse{value: val, isSet: true}
}

func (v NullableLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


