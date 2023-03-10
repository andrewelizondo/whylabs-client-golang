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

// checks if the EntityWeightRecordMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityWeightRecordMetadata{}

// EntityWeightRecordMetadata struct for EntityWeightRecordMetadata
type EntityWeightRecordMetadata struct {
	Version *int64 `json:"version,omitempty"`
	UpdatedTimestamp *int64 `json:"updatedTimestamp,omitempty"`
	Author *string `json:"author,omitempty"`
}

// NewEntityWeightRecordMetadata instantiates a new EntityWeightRecordMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityWeightRecordMetadata() *EntityWeightRecordMetadata {
	this := EntityWeightRecordMetadata{}
	return &this
}

// NewEntityWeightRecordMetadataWithDefaults instantiates a new EntityWeightRecordMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWeightRecordMetadataWithDefaults() *EntityWeightRecordMetadata {
	this := EntityWeightRecordMetadata{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EntityWeightRecordMetadata) GetVersion() int64 {
	if o == nil || isNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityWeightRecordMetadata) GetVersionOk() (*int64, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EntityWeightRecordMetadata) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *EntityWeightRecordMetadata) SetVersion(v int64) {
	o.Version = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *EntityWeightRecordMetadata) GetUpdatedTimestamp() int64 {
	if o == nil || isNil(o.UpdatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityWeightRecordMetadata) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *EntityWeightRecordMetadata) HasUpdatedTimestamp() bool {
	if o != nil && !isNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int64 and assigns it to the UpdatedTimestamp field.
func (o *EntityWeightRecordMetadata) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *EntityWeightRecordMetadata) GetAuthor() string {
	if o == nil || isNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityWeightRecordMetadata) GetAuthorOk() (*string, bool) {
	if o == nil || isNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *EntityWeightRecordMetadata) HasAuthor() bool {
	if o != nil && !isNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *EntityWeightRecordMetadata) SetAuthor(v string) {
	o.Author = &v
}

func (o EntityWeightRecordMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityWeightRecordMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.UpdatedTimestamp) {
		toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	}
	if !isNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	return toSerialize, nil
}

type NullableEntityWeightRecordMetadata struct {
	value *EntityWeightRecordMetadata
	isSet bool
}

func (v NullableEntityWeightRecordMetadata) Get() *EntityWeightRecordMetadata {
	return v.value
}

func (v *NullableEntityWeightRecordMetadata) Set(val *EntityWeightRecordMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityWeightRecordMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityWeightRecordMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityWeightRecordMetadata(val *EntityWeightRecordMetadata) *NullableEntityWeightRecordMetadata {
	return &NullableEntityWeightRecordMetadata{value: val, isSet: true}
}

func (v NullableEntityWeightRecordMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityWeightRecordMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


