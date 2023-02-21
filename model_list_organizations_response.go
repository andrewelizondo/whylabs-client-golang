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

// checks if the ListOrganizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationsResponse{}

// ListOrganizationsResponse Response for listing organization
type ListOrganizationsResponse struct {
	// A list of all known organization metadata
	Items []OrganizationSummary `json:"items"`
	Internal *bool `json:"internal,omitempty"`
}

// NewListOrganizationsResponse instantiates a new ListOrganizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationsResponse(items []OrganizationSummary) *ListOrganizationsResponse {
	this := ListOrganizationsResponse{}
	this.Items = items
	return &this
}

// NewListOrganizationsResponseWithDefaults instantiates a new ListOrganizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationsResponseWithDefaults() *ListOrganizationsResponse {
	this := ListOrganizationsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ListOrganizationsResponse) GetItems() []OrganizationSummary {
	if o == nil {
		var ret []OrganizationSummary
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponse) GetItemsOk() ([]OrganizationSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListOrganizationsResponse) SetItems(v []OrganizationSummary) {
	o.Items = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ListOrganizationsResponse) GetInternal() bool {
	if o == nil || isNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponse) GetInternalOk() (*bool, bool) {
	if o == nil || isNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ListOrganizationsResponse) HasInternal() bool {
	if o != nil && !isNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *ListOrganizationsResponse) SetInternal(v bool) {
	o.Internal = &v
}

func (o ListOrganizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !isNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	return toSerialize, nil
}

type NullableListOrganizationsResponse struct {
	value *ListOrganizationsResponse
	isSet bool
}

func (v NullableListOrganizationsResponse) Get() *ListOrganizationsResponse {
	return v.value
}

func (v *NullableListOrganizationsResponse) Set(val *ListOrganizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationsResponse(val *ListOrganizationsResponse) *NullableListOrganizationsResponse {
	return &NullableListOrganizationsResponse{value: val, isSet: true}
}

func (v NullableListOrganizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


