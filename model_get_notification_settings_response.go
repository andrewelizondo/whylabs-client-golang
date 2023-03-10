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

// checks if the GetNotificationSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNotificationSettingsResponse{}

// GetNotificationSettingsResponse Response for getting notification settings
type GetNotificationSettingsResponse struct {
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
}

// NewGetNotificationSettingsResponse instantiates a new GetNotificationSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNotificationSettingsResponse() *GetNotificationSettingsResponse {
	this := GetNotificationSettingsResponse{}
	return &this
}

// NewGetNotificationSettingsResponseWithDefaults instantiates a new GetNotificationSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNotificationSettingsResponseWithDefaults() *GetNotificationSettingsResponse {
	this := GetNotificationSettingsResponse{}
	return &this
}

// GetNotificationSettings returns the NotificationSettings field value if set, zero value otherwise.
func (o *GetNotificationSettingsResponse) GetNotificationSettings() NotificationSettings {
	if o == nil || isNil(o.NotificationSettings) {
		var ret NotificationSettings
		return ret
	}
	return *o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNotificationSettingsResponse) GetNotificationSettingsOk() (*NotificationSettings, bool) {
	if o == nil || isNil(o.NotificationSettings) {
		return nil, false
	}
	return o.NotificationSettings, true
}

// HasNotificationSettings returns a boolean if a field has been set.
func (o *GetNotificationSettingsResponse) HasNotificationSettings() bool {
	if o != nil && !isNil(o.NotificationSettings) {
		return true
	}

	return false
}

// SetNotificationSettings gets a reference to the given NotificationSettings and assigns it to the NotificationSettings field.
func (o *GetNotificationSettingsResponse) SetNotificationSettings(v NotificationSettings) {
	o.NotificationSettings = &v
}

func (o GetNotificationSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNotificationSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NotificationSettings) {
		toSerialize["notificationSettings"] = o.NotificationSettings
	}
	return toSerialize, nil
}

type NullableGetNotificationSettingsResponse struct {
	value *GetNotificationSettingsResponse
	isSet bool
}

func (v NullableGetNotificationSettingsResponse) Get() *GetNotificationSettingsResponse {
	return v.value
}

func (v *NullableGetNotificationSettingsResponse) Set(val *GetNotificationSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNotificationSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNotificationSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNotificationSettingsResponse(val *GetNotificationSettingsResponse) *NullableGetNotificationSettingsResponse {
	return &NullableGetNotificationSettingsResponse{value: val, isSet: true}
}

func (v NullableGetNotificationSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNotificationSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


