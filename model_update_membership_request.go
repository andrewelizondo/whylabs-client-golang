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

// checks if the UpdateMembershipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMembershipRequest{}

// UpdateMembershipRequest struct for UpdateMembershipRequest
type UpdateMembershipRequest struct {
	OrgId string `json:"orgId"`
	Email string `json:"email"`
	Role Role `json:"role"`
}

// NewUpdateMembershipRequest instantiates a new UpdateMembershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMembershipRequest(orgId string, email string, role Role) *UpdateMembershipRequest {
	this := UpdateMembershipRequest{}
	this.OrgId = orgId
	this.Email = email
	this.Role = role
	return &this
}

// NewUpdateMembershipRequestWithDefaults instantiates a new UpdateMembershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMembershipRequestWithDefaults() *UpdateMembershipRequest {
	this := UpdateMembershipRequest{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *UpdateMembershipRequest) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UpdateMembershipRequest) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *UpdateMembershipRequest) SetOrgId(v string) {
	o.OrgId = v
}

// GetEmail returns the Email field value
func (o *UpdateMembershipRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UpdateMembershipRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UpdateMembershipRequest) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *UpdateMembershipRequest) GetRole() Role {
	if o == nil {
		var ret Role
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UpdateMembershipRequest) GetRoleOk() (*Role, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UpdateMembershipRequest) SetRole(v Role) {
	o.Role = v
}

func (o UpdateMembershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMembershipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orgId"] = o.OrgId
	toSerialize["email"] = o.Email
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableUpdateMembershipRequest struct {
	value *UpdateMembershipRequest
	isSet bool
}

func (v NullableUpdateMembershipRequest) Get() *UpdateMembershipRequest {
	return v.value
}

func (v *NullableUpdateMembershipRequest) Set(val *UpdateMembershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMembershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMembershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMembershipRequest(val *UpdateMembershipRequest) *NullableUpdateMembershipRequest {
	return &NullableUpdateMembershipRequest{value: val, isSet: true}
}

func (v NullableUpdateMembershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMembershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


