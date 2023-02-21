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

// checks if the LogAsyncRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogAsyncRequest{}

// LogAsyncRequest Request payload for LogAsync.
type LogAsyncRequest struct {
	DatasetTimestamp *int64 `json:"datasetTimestamp,omitempty"`
	SegmentTags []SegmentTag `json:"segmentTags,omitempty"`
}

// NewLogAsyncRequest instantiates a new LogAsyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogAsyncRequest() *LogAsyncRequest {
	this := LogAsyncRequest{}
	return &this
}

// NewLogAsyncRequestWithDefaults instantiates a new LogAsyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogAsyncRequestWithDefaults() *LogAsyncRequest {
	this := LogAsyncRequest{}
	return &this
}

// GetDatasetTimestamp returns the DatasetTimestamp field value if set, zero value otherwise.
func (o *LogAsyncRequest) GetDatasetTimestamp() int64 {
	if o == nil || isNil(o.DatasetTimestamp) {
		var ret int64
		return ret
	}
	return *o.DatasetTimestamp
}

// GetDatasetTimestampOk returns a tuple with the DatasetTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAsyncRequest) GetDatasetTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.DatasetTimestamp) {
		return nil, false
	}
	return o.DatasetTimestamp, true
}

// HasDatasetTimestamp returns a boolean if a field has been set.
func (o *LogAsyncRequest) HasDatasetTimestamp() bool {
	if o != nil && !isNil(o.DatasetTimestamp) {
		return true
	}

	return false
}

// SetDatasetTimestamp gets a reference to the given int64 and assigns it to the DatasetTimestamp field.
func (o *LogAsyncRequest) SetDatasetTimestamp(v int64) {
	o.DatasetTimestamp = &v
}

// GetSegmentTags returns the SegmentTags field value if set, zero value otherwise.
func (o *LogAsyncRequest) GetSegmentTags() []SegmentTag {
	if o == nil || isNil(o.SegmentTags) {
		var ret []SegmentTag
		return ret
	}
	return o.SegmentTags
}

// GetSegmentTagsOk returns a tuple with the SegmentTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAsyncRequest) GetSegmentTagsOk() ([]SegmentTag, bool) {
	if o == nil || isNil(o.SegmentTags) {
		return nil, false
	}
	return o.SegmentTags, true
}

// HasSegmentTags returns a boolean if a field has been set.
func (o *LogAsyncRequest) HasSegmentTags() bool {
	if o != nil && !isNil(o.SegmentTags) {
		return true
	}

	return false
}

// SetSegmentTags gets a reference to the given []SegmentTag and assigns it to the SegmentTags field.
func (o *LogAsyncRequest) SetSegmentTags(v []SegmentTag) {
	o.SegmentTags = v
}

func (o LogAsyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogAsyncRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DatasetTimestamp) {
		toSerialize["datasetTimestamp"] = o.DatasetTimestamp
	}
	if !isNil(o.SegmentTags) {
		toSerialize["segmentTags"] = o.SegmentTags
	}
	return toSerialize, nil
}

type NullableLogAsyncRequest struct {
	value *LogAsyncRequest
	isSet bool
}

func (v NullableLogAsyncRequest) Get() *LogAsyncRequest {
	return v.value
}

func (v *NullableLogAsyncRequest) Set(val *LogAsyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLogAsyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLogAsyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogAsyncRequest(val *LogAsyncRequest) *NullableLogAsyncRequest {
	return &NullableLogAsyncRequest{value: val, isSet: true}
}

func (v NullableLogAsyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogAsyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


