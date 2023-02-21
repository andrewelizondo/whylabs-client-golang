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

// checks if the EventsPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsPath{}

// EventsPath Metadata for an event path (might contain multiple events)
type EventsPath struct {
	Id *string `json:"id,omitempty"`
	OrgId *string `json:"orgId,omitempty"`
	ModelId *string `json:"modelId,omitempty"`
	Segment *Segment `json:"segment,omitempty"`
	CreationTimestamp *int64 `json:"creationTimestamp,omitempty"`
	DatasetTimestamp *int64 `json:"datasetTimestamp,omitempty"`
	S3Path *string `json:"s3Path,omitempty"`
	Etag *string `json:"etag,omitempty"`
	Version NullableString `json:"version,omitempty"`
}

// NewEventsPath instantiates a new EventsPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsPath() *EventsPath {
	this := EventsPath{}
	return &this
}

// NewEventsPathWithDefaults instantiates a new EventsPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsPathWithDefaults() *EventsPath {
	this := EventsPath{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventsPath) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventsPath) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventsPath) SetId(v string) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *EventsPath) GetOrgId() string {
	if o == nil || isNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetOrgIdOk() (*string, bool) {
	if o == nil || isNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *EventsPath) HasOrgId() bool {
	if o != nil && !isNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *EventsPath) SetOrgId(v string) {
	o.OrgId = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *EventsPath) GetModelId() string {
	if o == nil || isNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetModelIdOk() (*string, bool) {
	if o == nil || isNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *EventsPath) HasModelId() bool {
	if o != nil && !isNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *EventsPath) SetModelId(v string) {
	o.ModelId = &v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *EventsPath) GetSegment() Segment {
	if o == nil || isNil(o.Segment) {
		var ret Segment
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetSegmentOk() (*Segment, bool) {
	if o == nil || isNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *EventsPath) HasSegment() bool {
	if o != nil && !isNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given Segment and assigns it to the Segment field.
func (o *EventsPath) SetSegment(v Segment) {
	o.Segment = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *EventsPath) GetCreationTimestamp() int64 {
	if o == nil || isNil(o.CreationTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetCreationTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.CreationTimestamp) {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *EventsPath) HasCreationTimestamp() bool {
	if o != nil && !isNil(o.CreationTimestamp) {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given int64 and assigns it to the CreationTimestamp field.
func (o *EventsPath) SetCreationTimestamp(v int64) {
	o.CreationTimestamp = &v
}

// GetDatasetTimestamp returns the DatasetTimestamp field value if set, zero value otherwise.
func (o *EventsPath) GetDatasetTimestamp() int64 {
	if o == nil || isNil(o.DatasetTimestamp) {
		var ret int64
		return ret
	}
	return *o.DatasetTimestamp
}

// GetDatasetTimestampOk returns a tuple with the DatasetTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetDatasetTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.DatasetTimestamp) {
		return nil, false
	}
	return o.DatasetTimestamp, true
}

// HasDatasetTimestamp returns a boolean if a field has been set.
func (o *EventsPath) HasDatasetTimestamp() bool {
	if o != nil && !isNil(o.DatasetTimestamp) {
		return true
	}

	return false
}

// SetDatasetTimestamp gets a reference to the given int64 and assigns it to the DatasetTimestamp field.
func (o *EventsPath) SetDatasetTimestamp(v int64) {
	o.DatasetTimestamp = &v
}

// GetS3Path returns the S3Path field value if set, zero value otherwise.
func (o *EventsPath) GetS3Path() string {
	if o == nil || isNil(o.S3Path) {
		var ret string
		return ret
	}
	return *o.S3Path
}

// GetS3PathOk returns a tuple with the S3Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetS3PathOk() (*string, bool) {
	if o == nil || isNil(o.S3Path) {
		return nil, false
	}
	return o.S3Path, true
}

// HasS3Path returns a boolean if a field has been set.
func (o *EventsPath) HasS3Path() bool {
	if o != nil && !isNil(o.S3Path) {
		return true
	}

	return false
}

// SetS3Path gets a reference to the given string and assigns it to the S3Path field.
func (o *EventsPath) SetS3Path(v string) {
	o.S3Path = &v
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *EventsPath) GetEtag() string {
	if o == nil || isNil(o.Etag) {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsPath) GetEtagOk() (*string, bool) {
	if o == nil || isNil(o.Etag) {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *EventsPath) HasEtag() bool {
	if o != nil && !isNil(o.Etag) {
		return true
	}

	return false
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *EventsPath) SetEtag(v string) {
	o.Etag = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsPath) GetVersion() string {
	if o == nil || isNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsPath) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *EventsPath) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *EventsPath) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *EventsPath) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *EventsPath) UnsetVersion() {
	o.Version.Unset()
}

func (o EventsPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !isNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	if !isNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !isNil(o.CreationTimestamp) {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	if !isNil(o.DatasetTimestamp) {
		toSerialize["datasetTimestamp"] = o.DatasetTimestamp
	}
	if !isNil(o.S3Path) {
		toSerialize["s3Path"] = o.S3Path
	}
	if !isNil(o.Etag) {
		toSerialize["etag"] = o.Etag
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return toSerialize, nil
}

type NullableEventsPath struct {
	value *EventsPath
	isSet bool
}

func (v NullableEventsPath) Get() *EventsPath {
	return v.value
}

func (v *NullableEventsPath) Set(val *EventsPath) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsPath) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsPath(val *EventsPath) *NullableEventsPath {
	return &NullableEventsPath{value: val, isSet: true}
}

func (v NullableEventsPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


