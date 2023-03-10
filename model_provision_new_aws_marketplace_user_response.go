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

// checks if the ProvisionNewAWSMarketplaceUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionNewAWSMarketplaceUserResponse{}

// ProvisionNewAWSMarketplaceUserResponse struct for ProvisionNewAWSMarketplaceUserResponse
type ProvisionNewAWSMarketplaceUserResponse struct {
	UserId string `json:"userId"`
	OrgId string `json:"orgId"`
	ModelId string `json:"modelId"`
	CustomerId string `json:"customerId"`
}

// NewProvisionNewAWSMarketplaceUserResponse instantiates a new ProvisionNewAWSMarketplaceUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionNewAWSMarketplaceUserResponse(userId string, orgId string, modelId string, customerId string) *ProvisionNewAWSMarketplaceUserResponse {
	this := ProvisionNewAWSMarketplaceUserResponse{}
	this.UserId = userId
	this.OrgId = orgId
	this.ModelId = modelId
	this.CustomerId = customerId
	return &this
}

// NewProvisionNewAWSMarketplaceUserResponseWithDefaults instantiates a new ProvisionNewAWSMarketplaceUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionNewAWSMarketplaceUserResponseWithDefaults() *ProvisionNewAWSMarketplaceUserResponse {
	this := ProvisionNewAWSMarketplaceUserResponse{}
	return &this
}

// GetUserId returns the UserId field value
func (o *ProvisionNewAWSMarketplaceUserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ProvisionNewAWSMarketplaceUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ProvisionNewAWSMarketplaceUserResponse) SetUserId(v string) {
	o.UserId = v
}

// GetOrgId returns the OrgId field value
func (o *ProvisionNewAWSMarketplaceUserResponse) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ProvisionNewAWSMarketplaceUserResponse) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ProvisionNewAWSMarketplaceUserResponse) SetOrgId(v string) {
	o.OrgId = v
}

// GetModelId returns the ModelId field value
func (o *ProvisionNewAWSMarketplaceUserResponse) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *ProvisionNewAWSMarketplaceUserResponse) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *ProvisionNewAWSMarketplaceUserResponse) SetModelId(v string) {
	o.ModelId = v
}

// GetCustomerId returns the CustomerId field value
func (o *ProvisionNewAWSMarketplaceUserResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ProvisionNewAWSMarketplaceUserResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ProvisionNewAWSMarketplaceUserResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o ProvisionNewAWSMarketplaceUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionNewAWSMarketplaceUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["orgId"] = o.OrgId
	toSerialize["modelId"] = o.ModelId
	toSerialize["customerId"] = o.CustomerId
	return toSerialize, nil
}

type NullableProvisionNewAWSMarketplaceUserResponse struct {
	value *ProvisionNewAWSMarketplaceUserResponse
	isSet bool
}

func (v NullableProvisionNewAWSMarketplaceUserResponse) Get() *ProvisionNewAWSMarketplaceUserResponse {
	return v.value
}

func (v *NullableProvisionNewAWSMarketplaceUserResponse) Set(val *ProvisionNewAWSMarketplaceUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionNewAWSMarketplaceUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionNewAWSMarketplaceUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionNewAWSMarketplaceUserResponse(val *ProvisionNewAWSMarketplaceUserResponse) *NullableProvisionNewAWSMarketplaceUserResponse {
	return &NullableProvisionNewAWSMarketplaceUserResponse{value: val, isSet: true}
}

func (v NullableProvisionNewAWSMarketplaceUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionNewAWSMarketplaceUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


