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

// checks if the UpdateConnectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateConnectionRequest{}

// UpdateConnectionRequest struct for UpdateConnectionRequest
type UpdateConnectionRequest struct {
	WorkspaceId NullableString `json:"workspaceId,omitempty"`
	OrgId NullableString `json:"orgId,omitempty"`
	Changes *UpdateConnectionChanges `json:"changes,omitempty"`
}

// NewUpdateConnectionRequest instantiates a new UpdateConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConnectionRequest() *UpdateConnectionRequest {
	this := UpdateConnectionRequest{}
	return &this
}

// NewUpdateConnectionRequestWithDefaults instantiates a new UpdateConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConnectionRequestWithDefaults() *UpdateConnectionRequest {
	this := UpdateConnectionRequest{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateConnectionRequest) GetWorkspaceId() string {
	if o == nil || isNil(o.WorkspaceId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkspaceId.Get()
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateConnectionRequest) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceId.Get(), o.WorkspaceId.IsSet()
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *UpdateConnectionRequest) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId.IsSet() {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given NullableString and assigns it to the WorkspaceId field.
func (o *UpdateConnectionRequest) SetWorkspaceId(v string) {
	o.WorkspaceId.Set(&v)
}
// SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil
func (o *UpdateConnectionRequest) SetWorkspaceIdNil() {
	o.WorkspaceId.Set(nil)
}

// UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
func (o *UpdateConnectionRequest) UnsetWorkspaceId() {
	o.WorkspaceId.Unset()
}

// GetOrgId returns the OrgId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateConnectionRequest) GetOrgId() string {
	if o == nil || isNil(o.OrgId.Get()) {
		var ret string
		return ret
	}
	return *o.OrgId.Get()
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateConnectionRequest) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgId.Get(), o.OrgId.IsSet()
}

// HasOrgId returns a boolean if a field has been set.
func (o *UpdateConnectionRequest) HasOrgId() bool {
	if o != nil && o.OrgId.IsSet() {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given NullableString and assigns it to the OrgId field.
func (o *UpdateConnectionRequest) SetOrgId(v string) {
	o.OrgId.Set(&v)
}
// SetOrgIdNil sets the value for OrgId to be an explicit nil
func (o *UpdateConnectionRequest) SetOrgIdNil() {
	o.OrgId.Set(nil)
}

// UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
func (o *UpdateConnectionRequest) UnsetOrgId() {
	o.OrgId.Unset()
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *UpdateConnectionRequest) GetChanges() UpdateConnectionChanges {
	if o == nil || isNil(o.Changes) {
		var ret UpdateConnectionChanges
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequest) GetChangesOk() (*UpdateConnectionChanges, bool) {
	if o == nil || isNil(o.Changes) {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *UpdateConnectionRequest) HasChanges() bool {
	if o != nil && !isNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given UpdateConnectionChanges and assigns it to the Changes field.
func (o *UpdateConnectionRequest) SetChanges(v UpdateConnectionChanges) {
	o.Changes = &v
}

func (o UpdateConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateConnectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WorkspaceId.IsSet() {
		toSerialize["workspaceId"] = o.WorkspaceId.Get()
	}
	if o.OrgId.IsSet() {
		toSerialize["orgId"] = o.OrgId.Get()
	}
	if !isNil(o.Changes) {
		toSerialize["changes"] = o.Changes
	}
	return toSerialize, nil
}

type NullableUpdateConnectionRequest struct {
	value *UpdateConnectionRequest
	isSet bool
}

func (v NullableUpdateConnectionRequest) Get() *UpdateConnectionRequest {
	return v.value
}

func (v *NullableUpdateConnectionRequest) Set(val *UpdateConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConnectionRequest(val *UpdateConnectionRequest) *NullableUpdateConnectionRequest {
	return &NullableUpdateConnectionRequest{value: val, isSet: true}
}

func (v NullableUpdateConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


