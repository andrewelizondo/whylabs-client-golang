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

// checks if the OrganizationSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSummary{}

// OrganizationSummary Summary about an organization
type OrganizationSummary struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	SubscriptionTier *SubscriptionTier `json:"subscriptionTier,omitempty"`
	Domain NullableString `json:"domain,omitempty"`
	EmailDomains NullableString `json:"emailDomains,omitempty"`
	ObservatoryUrl NullableString `json:"observatoryUrl,omitempty"`
	NotificationEmailAddress NullableString `json:"notificationEmailAddress,omitempty"`
	SlackWebhook NullableString `json:"slackWebhook,omitempty"`
	PagerDutyKey NullableString `json:"pagerDutyKey,omitempty"`
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
}

// NewOrganizationSummary instantiates a new OrganizationSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSummary(id string) *OrganizationSummary {
	this := OrganizationSummary{}
	this.Id = id
	return &this
}

// NewOrganizationSummaryWithDefaults instantiates a new OrganizationSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSummaryWithDefaults() *OrganizationSummary {
	this := OrganizationSummary{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationSummary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationSummary) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationSummary) SetName(v string) {
	o.Name = &v
}

// GetSubscriptionTier returns the SubscriptionTier field value if set, zero value otherwise.
func (o *OrganizationSummary) GetSubscriptionTier() SubscriptionTier {
	if o == nil || isNil(o.SubscriptionTier) {
		var ret SubscriptionTier
		return ret
	}
	return *o.SubscriptionTier
}

// GetSubscriptionTierOk returns a tuple with the SubscriptionTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSummary) GetSubscriptionTierOk() (*SubscriptionTier, bool) {
	if o == nil || isNil(o.SubscriptionTier) {
		return nil, false
	}
	return o.SubscriptionTier, true
}

// HasSubscriptionTier returns a boolean if a field has been set.
func (o *OrganizationSummary) HasSubscriptionTier() bool {
	if o != nil && !isNil(o.SubscriptionTier) {
		return true
	}

	return false
}

// SetSubscriptionTier gets a reference to the given SubscriptionTier and assigns it to the SubscriptionTier field.
func (o *OrganizationSummary) SetSubscriptionTier(v SubscriptionTier) {
	o.SubscriptionTier = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSummary) GetDomain() string {
	if o == nil || isNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSummary) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *OrganizationSummary) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *OrganizationSummary) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *OrganizationSummary) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *OrganizationSummary) UnsetDomain() {
	o.Domain.Unset()
}

// GetEmailDomains returns the EmailDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSummary) GetEmailDomains() string {
	if o == nil || isNil(o.EmailDomains.Get()) {
		var ret string
		return ret
	}
	return *o.EmailDomains.Get()
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSummary) GetEmailDomainsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailDomains.Get(), o.EmailDomains.IsSet()
}

// HasEmailDomains returns a boolean if a field has been set.
func (o *OrganizationSummary) HasEmailDomains() bool {
	if o != nil && o.EmailDomains.IsSet() {
		return true
	}

	return false
}

// SetEmailDomains gets a reference to the given NullableString and assigns it to the EmailDomains field.
func (o *OrganizationSummary) SetEmailDomains(v string) {
	o.EmailDomains.Set(&v)
}
// SetEmailDomainsNil sets the value for EmailDomains to be an explicit nil
func (o *OrganizationSummary) SetEmailDomainsNil() {
	o.EmailDomains.Set(nil)
}

// UnsetEmailDomains ensures that no value is present for EmailDomains, not even an explicit nil
func (o *OrganizationSummary) UnsetEmailDomains() {
	o.EmailDomains.Unset()
}

// GetObservatoryUrl returns the ObservatoryUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSummary) GetObservatoryUrl() string {
	if o == nil || isNil(o.ObservatoryUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ObservatoryUrl.Get()
}

// GetObservatoryUrlOk returns a tuple with the ObservatoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSummary) GetObservatoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObservatoryUrl.Get(), o.ObservatoryUrl.IsSet()
}

// HasObservatoryUrl returns a boolean if a field has been set.
func (o *OrganizationSummary) HasObservatoryUrl() bool {
	if o != nil && o.ObservatoryUrl.IsSet() {
		return true
	}

	return false
}

// SetObservatoryUrl gets a reference to the given NullableString and assigns it to the ObservatoryUrl field.
func (o *OrganizationSummary) SetObservatoryUrl(v string) {
	o.ObservatoryUrl.Set(&v)
}
// SetObservatoryUrlNil sets the value for ObservatoryUrl to be an explicit nil
func (o *OrganizationSummary) SetObservatoryUrlNil() {
	o.ObservatoryUrl.Set(nil)
}

// UnsetObservatoryUrl ensures that no value is present for ObservatoryUrl, not even an explicit nil
func (o *OrganizationSummary) UnsetObservatoryUrl() {
	o.ObservatoryUrl.Unset()
}

// GetNotificationEmailAddress returns the NotificationEmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSummary) GetNotificationEmailAddress() string {
	if o == nil || isNil(o.NotificationEmailAddress.Get()) {
		var ret string
		return ret
	}
	return *o.NotificationEmailAddress.Get()
}

// GetNotificationEmailAddressOk returns a tuple with the NotificationEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSummary) GetNotificationEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationEmailAddress.Get(), o.NotificationEmailAddress.IsSet()
}

// HasNotificationEmailAddress returns a boolean if a field has been set.
func (o *OrganizationSummary) HasNotificationEmailAddress() bool {
	if o != nil && o.NotificationEmailAddress.IsSet() {
		return true
	}

	return false
}

// SetNotificationEmailAddress gets a reference to the given NullableString and assigns it to the NotificationEmailAddress field.
func (o *OrganizationSummary) SetNotificationEmailAddress(v string) {
	o.NotificationEmailAddress.Set(&v)
}
// SetNotificationEmailAddressNil sets the value for NotificationEmailAddress to be an explicit nil
func (o *OrganizationSummary) SetNotificationEmailAddressNil() {
	o.NotificationEmailAddress.Set(nil)
}

// UnsetNotificationEmailAddress ensures that no value is present for NotificationEmailAddress, not even an explicit nil
func (o *OrganizationSummary) UnsetNotificationEmailAddress() {
	o.NotificationEmailAddress.Unset()
}

// GetSlackWebhook returns the SlackWebhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSummary) GetSlackWebhook() string {
	if o == nil || isNil(o.SlackWebhook.Get()) {
		var ret string
		return ret
	}
	return *o.SlackWebhook.Get()
}

// GetSlackWebhookOk returns a tuple with the SlackWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSummary) GetSlackWebhookOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackWebhook.Get(), o.SlackWebhook.IsSet()
}

// HasSlackWebhook returns a boolean if a field has been set.
func (o *OrganizationSummary) HasSlackWebhook() bool {
	if o != nil && o.SlackWebhook.IsSet() {
		return true
	}

	return false
}

// SetSlackWebhook gets a reference to the given NullableString and assigns it to the SlackWebhook field.
func (o *OrganizationSummary) SetSlackWebhook(v string) {
	o.SlackWebhook.Set(&v)
}
// SetSlackWebhookNil sets the value for SlackWebhook to be an explicit nil
func (o *OrganizationSummary) SetSlackWebhookNil() {
	o.SlackWebhook.Set(nil)
}

// UnsetSlackWebhook ensures that no value is present for SlackWebhook, not even an explicit nil
func (o *OrganizationSummary) UnsetSlackWebhook() {
	o.SlackWebhook.Unset()
}

// GetPagerDutyKey returns the PagerDutyKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSummary) GetPagerDutyKey() string {
	if o == nil || isNil(o.PagerDutyKey.Get()) {
		var ret string
		return ret
	}
	return *o.PagerDutyKey.Get()
}

// GetPagerDutyKeyOk returns a tuple with the PagerDutyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSummary) GetPagerDutyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PagerDutyKey.Get(), o.PagerDutyKey.IsSet()
}

// HasPagerDutyKey returns a boolean if a field has been set.
func (o *OrganizationSummary) HasPagerDutyKey() bool {
	if o != nil && o.PagerDutyKey.IsSet() {
		return true
	}

	return false
}

// SetPagerDutyKey gets a reference to the given NullableString and assigns it to the PagerDutyKey field.
func (o *OrganizationSummary) SetPagerDutyKey(v string) {
	o.PagerDutyKey.Set(&v)
}
// SetPagerDutyKeyNil sets the value for PagerDutyKey to be an explicit nil
func (o *OrganizationSummary) SetPagerDutyKeyNil() {
	o.PagerDutyKey.Set(nil)
}

// UnsetPagerDutyKey ensures that no value is present for PagerDutyKey, not even an explicit nil
func (o *OrganizationSummary) UnsetPagerDutyKey() {
	o.PagerDutyKey.Unset()
}

// GetNotificationSettings returns the NotificationSettings field value if set, zero value otherwise.
func (o *OrganizationSummary) GetNotificationSettings() NotificationSettings {
	if o == nil || isNil(o.NotificationSettings) {
		var ret NotificationSettings
		return ret
	}
	return *o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSummary) GetNotificationSettingsOk() (*NotificationSettings, bool) {
	if o == nil || isNil(o.NotificationSettings) {
		return nil, false
	}
	return o.NotificationSettings, true
}

// HasNotificationSettings returns a boolean if a field has been set.
func (o *OrganizationSummary) HasNotificationSettings() bool {
	if o != nil && !isNil(o.NotificationSettings) {
		return true
	}

	return false
}

// SetNotificationSettings gets a reference to the given NotificationSettings and assigns it to the NotificationSettings field.
func (o *OrganizationSummary) SetNotificationSettings(v NotificationSettings) {
	o.NotificationSettings = &v
}

func (o OrganizationSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SubscriptionTier) {
		toSerialize["subscriptionTier"] = o.SubscriptionTier
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.EmailDomains.IsSet() {
		toSerialize["emailDomains"] = o.EmailDomains.Get()
	}
	if o.ObservatoryUrl.IsSet() {
		toSerialize["observatoryUrl"] = o.ObservatoryUrl.Get()
	}
	if o.NotificationEmailAddress.IsSet() {
		toSerialize["notificationEmailAddress"] = o.NotificationEmailAddress.Get()
	}
	if o.SlackWebhook.IsSet() {
		toSerialize["slackWebhook"] = o.SlackWebhook.Get()
	}
	if o.PagerDutyKey.IsSet() {
		toSerialize["pagerDutyKey"] = o.PagerDutyKey.Get()
	}
	if !isNil(o.NotificationSettings) {
		toSerialize["notificationSettings"] = o.NotificationSettings
	}
	return toSerialize, nil
}

type NullableOrganizationSummary struct {
	value *OrganizationSummary
	isSet bool
}

func (v NullableOrganizationSummary) Get() *OrganizationSummary {
	return v.value
}

func (v *NullableOrganizationSummary) Set(val *OrganizationSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSummary(val *OrganizationSummary) *NullableOrganizationSummary {
	return &NullableOrganizationSummary{value: val, isSet: true}
}

func (v NullableOrganizationSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

