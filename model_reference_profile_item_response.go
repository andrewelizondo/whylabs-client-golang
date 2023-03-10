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

// checks if the ReferenceProfileItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceProfileItemResponse{}

// ReferenceProfileItemResponse A single reference item response.
type ReferenceProfileItemResponse struct {
	OrgId *string `json:"orgId,omitempty"`
	ModelId *string `json:"modelId,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Id *string `json:"id,omitempty"`
	UploadTimestamp *int64 `json:"uploadTimestamp,omitempty"`
	DatasetTimestamp NullableInt64 `json:"datasetTimestamp,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewReferenceProfileItemResponse instantiates a new ReferenceProfileItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceProfileItemResponse() *ReferenceProfileItemResponse {
	this := ReferenceProfileItemResponse{}
	return &this
}

// NewReferenceProfileItemResponseWithDefaults instantiates a new ReferenceProfileItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceProfileItemResponseWithDefaults() *ReferenceProfileItemResponse {
	this := ReferenceProfileItemResponse{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ReferenceProfileItemResponse) GetOrgId() string {
	if o == nil || isNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProfileItemResponse) GetOrgIdOk() (*string, bool) {
	if o == nil || isNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ReferenceProfileItemResponse) HasOrgId() bool {
	if o != nil && !isNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ReferenceProfileItemResponse) SetOrgId(v string) {
	o.OrgId = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *ReferenceProfileItemResponse) GetModelId() string {
	if o == nil || isNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProfileItemResponse) GetModelIdOk() (*string, bool) {
	if o == nil || isNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *ReferenceProfileItemResponse) HasModelId() bool {
	if o != nil && !isNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *ReferenceProfileItemResponse) SetModelId(v string) {
	o.ModelId = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ReferenceProfileItemResponse) GetAlias() string {
	if o == nil || isNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProfileItemResponse) GetAliasOk() (*string, bool) {
	if o == nil || isNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ReferenceProfileItemResponse) HasAlias() bool {
	if o != nil && !isNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ReferenceProfileItemResponse) SetAlias(v string) {
	o.Alias = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReferenceProfileItemResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProfileItemResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReferenceProfileItemResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReferenceProfileItemResponse) SetId(v string) {
	o.Id = &v
}

// GetUploadTimestamp returns the UploadTimestamp field value if set, zero value otherwise.
func (o *ReferenceProfileItemResponse) GetUploadTimestamp() int64 {
	if o == nil || isNil(o.UploadTimestamp) {
		var ret int64
		return ret
	}
	return *o.UploadTimestamp
}

// GetUploadTimestampOk returns a tuple with the UploadTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProfileItemResponse) GetUploadTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.UploadTimestamp) {
		return nil, false
	}
	return o.UploadTimestamp, true
}

// HasUploadTimestamp returns a boolean if a field has been set.
func (o *ReferenceProfileItemResponse) HasUploadTimestamp() bool {
	if o != nil && !isNil(o.UploadTimestamp) {
		return true
	}

	return false
}

// SetUploadTimestamp gets a reference to the given int64 and assigns it to the UploadTimestamp field.
func (o *ReferenceProfileItemResponse) SetUploadTimestamp(v int64) {
	o.UploadTimestamp = &v
}

// GetDatasetTimestamp returns the DatasetTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferenceProfileItemResponse) GetDatasetTimestamp() int64 {
	if o == nil || isNil(o.DatasetTimestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.DatasetTimestamp.Get()
}

// GetDatasetTimestampOk returns a tuple with the DatasetTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferenceProfileItemResponse) GetDatasetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatasetTimestamp.Get(), o.DatasetTimestamp.IsSet()
}

// HasDatasetTimestamp returns a boolean if a field has been set.
func (o *ReferenceProfileItemResponse) HasDatasetTimestamp() bool {
	if o != nil && o.DatasetTimestamp.IsSet() {
		return true
	}

	return false
}

// SetDatasetTimestamp gets a reference to the given NullableInt64 and assigns it to the DatasetTimestamp field.
func (o *ReferenceProfileItemResponse) SetDatasetTimestamp(v int64) {
	o.DatasetTimestamp.Set(&v)
}
// SetDatasetTimestampNil sets the value for DatasetTimestamp to be an explicit nil
func (o *ReferenceProfileItemResponse) SetDatasetTimestampNil() {
	o.DatasetTimestamp.Set(nil)
}

// UnsetDatasetTimestamp ensures that no value is present for DatasetTimestamp, not even an explicit nil
func (o *ReferenceProfileItemResponse) UnsetDatasetTimestamp() {
	o.DatasetTimestamp.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ReferenceProfileItemResponse) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProfileItemResponse) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ReferenceProfileItemResponse) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ReferenceProfileItemResponse) SetPath(v string) {
	o.Path = &v
}

func (o ReferenceProfileItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceProfileItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !isNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	if !isNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.UploadTimestamp) {
		toSerialize["uploadTimestamp"] = o.UploadTimestamp
	}
	if o.DatasetTimestamp.IsSet() {
		toSerialize["datasetTimestamp"] = o.DatasetTimestamp.Get()
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableReferenceProfileItemResponse struct {
	value *ReferenceProfileItemResponse
	isSet bool
}

func (v NullableReferenceProfileItemResponse) Get() *ReferenceProfileItemResponse {
	return v.value
}

func (v *NullableReferenceProfileItemResponse) Set(val *ReferenceProfileItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceProfileItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceProfileItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceProfileItemResponse(val *ReferenceProfileItemResponse) *NullableReferenceProfileItemResponse {
	return &NullableReferenceProfileItemResponse{value: val, isSet: true}
}

func (v NullableReferenceProfileItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceProfileItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


