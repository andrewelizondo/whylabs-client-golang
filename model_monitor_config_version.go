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

// checks if the MonitorConfigVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorConfigVersion{}

// MonitorConfigVersion struct for MonitorConfigVersion
type MonitorConfigVersion struct {
	Path *string `json:"path,omitempty"`
	VersionId *string `json:"versionId,omitempty"`
	LastModified *int64 `json:"lastModified,omitempty"`
}

// NewMonitorConfigVersion instantiates a new MonitorConfigVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorConfigVersion() *MonitorConfigVersion {
	this := MonitorConfigVersion{}
	return &this
}

// NewMonitorConfigVersionWithDefaults instantiates a new MonitorConfigVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorConfigVersionWithDefaults() *MonitorConfigVersion {
	this := MonitorConfigVersion{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MonitorConfigVersion) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorConfigVersion) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MonitorConfigVersion) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MonitorConfigVersion) SetPath(v string) {
	o.Path = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *MonitorConfigVersion) GetVersionId() string {
	if o == nil || isNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorConfigVersion) GetVersionIdOk() (*string, bool) {
	if o == nil || isNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *MonitorConfigVersion) HasVersionId() bool {
	if o != nil && !isNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *MonitorConfigVersion) SetVersionId(v string) {
	o.VersionId = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *MonitorConfigVersion) GetLastModified() int64 {
	if o == nil || isNil(o.LastModified) {
		var ret int64
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorConfigVersion) GetLastModifiedOk() (*int64, bool) {
	if o == nil || isNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *MonitorConfigVersion) HasLastModified() bool {
	if o != nil && !isNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given int64 and assigns it to the LastModified field.
func (o *MonitorConfigVersion) SetLastModified(v int64) {
	o.LastModified = &v
}

func (o MonitorConfigVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorConfigVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.VersionId) {
		toSerialize["versionId"] = o.VersionId
	}
	if !isNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableMonitorConfigVersion struct {
	value *MonitorConfigVersion
	isSet bool
}

func (v NullableMonitorConfigVersion) Get() *MonitorConfigVersion {
	return v.value
}

func (v *NullableMonitorConfigVersion) Set(val *MonitorConfigVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorConfigVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorConfigVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorConfigVersion(val *MonitorConfigVersion) *NullableMonitorConfigVersion {
	return &NullableMonitorConfigVersion{value: val, isSet: true}
}

func (v NullableMonitorConfigVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorConfigVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


