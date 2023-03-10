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

// checks if the GetConnectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConnectionRequest{}

// GetConnectionRequest struct for GetConnectionRequest
type GetConnectionRequest struct {
	// Look up a connection by the org that is connected to.
	OrgId NullableString `json:"orgId,omitempty"`
	// Look up a connection by the workspace that it originates from.
	WorkspaceId NullableString `json:"workspaceId,omitempty"`
}

// NewGetConnectionRequest instantiates a new GetConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectionRequest() *GetConnectionRequest {
	this := GetConnectionRequest{}
	return &this
}

// NewGetConnectionRequestWithDefaults instantiates a new GetConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectionRequestWithDefaults() *GetConnectionRequest {
	this := GetConnectionRequest{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConnectionRequest) GetOrgId() string {
	if o == nil || isNil(o.OrgId.Get()) {
		var ret string
		return ret
	}
	return *o.OrgId.Get()
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConnectionRequest) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgId.Get(), o.OrgId.IsSet()
}

// HasOrgId returns a boolean if a field has been set.
func (o *GetConnectionRequest) HasOrgId() bool {
	if o != nil && o.OrgId.IsSet() {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given NullableString and assigns it to the OrgId field.
func (o *GetConnectionRequest) SetOrgId(v string) {
	o.OrgId.Set(&v)
}
// SetOrgIdNil sets the value for OrgId to be an explicit nil
func (o *GetConnectionRequest) SetOrgIdNil() {
	o.OrgId.Set(nil)
}

// UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
func (o *GetConnectionRequest) UnsetOrgId() {
	o.OrgId.Unset()
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConnectionRequest) GetWorkspaceId() string {
	if o == nil || isNil(o.WorkspaceId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkspaceId.Get()
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConnectionRequest) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceId.Get(), o.WorkspaceId.IsSet()
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *GetConnectionRequest) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId.IsSet() {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given NullableString and assigns it to the WorkspaceId field.
func (o *GetConnectionRequest) SetWorkspaceId(v string) {
	o.WorkspaceId.Set(&v)
}
// SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil
func (o *GetConnectionRequest) SetWorkspaceIdNil() {
	o.WorkspaceId.Set(nil)
}

// UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
func (o *GetConnectionRequest) UnsetWorkspaceId() {
	o.WorkspaceId.Unset()
}

func (o GetConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConnectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgId.IsSet() {
		toSerialize["orgId"] = o.OrgId.Get()
	}
	if o.WorkspaceId.IsSet() {
		toSerialize["workspaceId"] = o.WorkspaceId.Get()
	}
	return toSerialize, nil
}

type NullableGetConnectionRequest struct {
	value *GetConnectionRequest
	isSet bool
}

func (v NullableGetConnectionRequest) Get() *GetConnectionRequest {
	return v.value
}

func (v *NullableGetConnectionRequest) Set(val *GetConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectionRequest(val *GetConnectionRequest) *NullableGetConnectionRequest {
	return &NullableGetConnectionRequest{value: val, isSet: true}
}

func (v NullableGetConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


