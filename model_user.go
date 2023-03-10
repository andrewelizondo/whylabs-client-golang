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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User User metadata
type User struct {
	// The id of the user.
	UserId string `json:"userId"`
	// The user's email address.
	Email string `json:"email"`
	// The user's JSON serialized preferences. Schema defined in Dashbird.
	Preferences NullableString `json:"preferences,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(userId string, email string) *User {
	this := User{}
	this.UserId = userId
	this.Email = email
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetUserId returns the UserId field value
func (o *User) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *User) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *User) SetUserId(v string) {
	o.UserId = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPreferences() string {
	if o == nil || isNil(o.Preferences.Get()) {
		var ret string
		return ret
	}
	return *o.Preferences.Get()
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPreferencesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preferences.Get(), o.Preferences.IsSet()
}

// HasPreferences returns a boolean if a field has been set.
func (o *User) HasPreferences() bool {
	if o != nil && o.Preferences.IsSet() {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given NullableString and assigns it to the Preferences field.
func (o *User) SetPreferences(v string) {
	o.Preferences.Set(&v)
}
// SetPreferencesNil sets the value for Preferences to be an explicit nil
func (o *User) SetPreferencesNil() {
	o.Preferences.Set(nil)
}

// UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
func (o *User) UnsetPreferences() {
	o.Preferences.Unset()
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["email"] = o.Email
	if o.Preferences.IsSet() {
		toSerialize["preferences"] = o.Preferences.Get()
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


