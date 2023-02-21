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

// checks if the RequestMonitorConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestMonitorConfig{}

// RequestMonitorConfig A configuration for the whylabs monitor for a single model or segment.
type RequestMonitorConfig struct {
	SchemaVersion *string `json:"schemaVersion,omitempty"`
	NumStdDev NullableFloat64 `json:"numStdDev,omitempty"`
	Reference *MonitorRequestReference `json:"reference,omitempty"`
	MissingRecentData *MissingRecentDataRequestConfig `json:"missingRecentData,omitempty"`
	MissingRecentProfiles *MissingRecentProfilesRequestConfig `json:"missingRecentProfiles,omitempty"`
	Distribution *DistributionMonitorRequestConfig `json:"distribution,omitempty"`
	MissingValues *MissingValuesMonitorRequestConfig `json:"missingValues,omitempty"`
	UniqueValues *UniqueValuesMonitorRequestConfig `json:"uniqueValues,omitempty"`
	Datatype *DatatypeMonitorRequestConfig `json:"datatype,omitempty"`
	SeasonalARIMA *SeasonalARIMARequestConfig `json:"seasonalARIMA,omitempty"`
}

// NewRequestMonitorConfig instantiates a new RequestMonitorConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMonitorConfig() *RequestMonitorConfig {
	this := RequestMonitorConfig{}
	return &this
}

// NewRequestMonitorConfigWithDefaults instantiates a new RequestMonitorConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMonitorConfigWithDefaults() *RequestMonitorConfig {
	this := RequestMonitorConfig{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetSchemaVersion() string {
	if o == nil || isNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetSchemaVersionOk() (*string, bool) {
	if o == nil || isNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasSchemaVersion() bool {
	if o != nil && !isNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *RequestMonitorConfig) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetNumStdDev returns the NumStdDev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestMonitorConfig) GetNumStdDev() float64 {
	if o == nil || isNil(o.NumStdDev.Get()) {
		var ret float64
		return ret
	}
	return *o.NumStdDev.Get()
}

// GetNumStdDevOk returns a tuple with the NumStdDev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMonitorConfig) GetNumStdDevOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumStdDev.Get(), o.NumStdDev.IsSet()
}

// HasNumStdDev returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasNumStdDev() bool {
	if o != nil && o.NumStdDev.IsSet() {
		return true
	}

	return false
}

// SetNumStdDev gets a reference to the given NullableFloat64 and assigns it to the NumStdDev field.
func (o *RequestMonitorConfig) SetNumStdDev(v float64) {
	o.NumStdDev.Set(&v)
}
// SetNumStdDevNil sets the value for NumStdDev to be an explicit nil
func (o *RequestMonitorConfig) SetNumStdDevNil() {
	o.NumStdDev.Set(nil)
}

// UnsetNumStdDev ensures that no value is present for NumStdDev, not even an explicit nil
func (o *RequestMonitorConfig) UnsetNumStdDev() {
	o.NumStdDev.Unset()
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetReference() MonitorRequestReference {
	if o == nil || isNil(o.Reference) {
		var ret MonitorRequestReference
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetReferenceOk() (*MonitorRequestReference, bool) {
	if o == nil || isNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasReference() bool {
	if o != nil && !isNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given MonitorRequestReference and assigns it to the Reference field.
func (o *RequestMonitorConfig) SetReference(v MonitorRequestReference) {
	o.Reference = &v
}

// GetMissingRecentData returns the MissingRecentData field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetMissingRecentData() MissingRecentDataRequestConfig {
	if o == nil || isNil(o.MissingRecentData) {
		var ret MissingRecentDataRequestConfig
		return ret
	}
	return *o.MissingRecentData
}

// GetMissingRecentDataOk returns a tuple with the MissingRecentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetMissingRecentDataOk() (*MissingRecentDataRequestConfig, bool) {
	if o == nil || isNil(o.MissingRecentData) {
		return nil, false
	}
	return o.MissingRecentData, true
}

// HasMissingRecentData returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasMissingRecentData() bool {
	if o != nil && !isNil(o.MissingRecentData) {
		return true
	}

	return false
}

// SetMissingRecentData gets a reference to the given MissingRecentDataRequestConfig and assigns it to the MissingRecentData field.
func (o *RequestMonitorConfig) SetMissingRecentData(v MissingRecentDataRequestConfig) {
	o.MissingRecentData = &v
}

// GetMissingRecentProfiles returns the MissingRecentProfiles field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetMissingRecentProfiles() MissingRecentProfilesRequestConfig {
	if o == nil || isNil(o.MissingRecentProfiles) {
		var ret MissingRecentProfilesRequestConfig
		return ret
	}
	return *o.MissingRecentProfiles
}

// GetMissingRecentProfilesOk returns a tuple with the MissingRecentProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetMissingRecentProfilesOk() (*MissingRecentProfilesRequestConfig, bool) {
	if o == nil || isNil(o.MissingRecentProfiles) {
		return nil, false
	}
	return o.MissingRecentProfiles, true
}

// HasMissingRecentProfiles returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasMissingRecentProfiles() bool {
	if o != nil && !isNil(o.MissingRecentProfiles) {
		return true
	}

	return false
}

// SetMissingRecentProfiles gets a reference to the given MissingRecentProfilesRequestConfig and assigns it to the MissingRecentProfiles field.
func (o *RequestMonitorConfig) SetMissingRecentProfiles(v MissingRecentProfilesRequestConfig) {
	o.MissingRecentProfiles = &v
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetDistribution() DistributionMonitorRequestConfig {
	if o == nil || isNil(o.Distribution) {
		var ret DistributionMonitorRequestConfig
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetDistributionOk() (*DistributionMonitorRequestConfig, bool) {
	if o == nil || isNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasDistribution() bool {
	if o != nil && !isNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given DistributionMonitorRequestConfig and assigns it to the Distribution field.
func (o *RequestMonitorConfig) SetDistribution(v DistributionMonitorRequestConfig) {
	o.Distribution = &v
}

// GetMissingValues returns the MissingValues field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetMissingValues() MissingValuesMonitorRequestConfig {
	if o == nil || isNil(o.MissingValues) {
		var ret MissingValuesMonitorRequestConfig
		return ret
	}
	return *o.MissingValues
}

// GetMissingValuesOk returns a tuple with the MissingValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetMissingValuesOk() (*MissingValuesMonitorRequestConfig, bool) {
	if o == nil || isNil(o.MissingValues) {
		return nil, false
	}
	return o.MissingValues, true
}

// HasMissingValues returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasMissingValues() bool {
	if o != nil && !isNil(o.MissingValues) {
		return true
	}

	return false
}

// SetMissingValues gets a reference to the given MissingValuesMonitorRequestConfig and assigns it to the MissingValues field.
func (o *RequestMonitorConfig) SetMissingValues(v MissingValuesMonitorRequestConfig) {
	o.MissingValues = &v
}

// GetUniqueValues returns the UniqueValues field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetUniqueValues() UniqueValuesMonitorRequestConfig {
	if o == nil || isNil(o.UniqueValues) {
		var ret UniqueValuesMonitorRequestConfig
		return ret
	}
	return *o.UniqueValues
}

// GetUniqueValuesOk returns a tuple with the UniqueValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetUniqueValuesOk() (*UniqueValuesMonitorRequestConfig, bool) {
	if o == nil || isNil(o.UniqueValues) {
		return nil, false
	}
	return o.UniqueValues, true
}

// HasUniqueValues returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasUniqueValues() bool {
	if o != nil && !isNil(o.UniqueValues) {
		return true
	}

	return false
}

// SetUniqueValues gets a reference to the given UniqueValuesMonitorRequestConfig and assigns it to the UniqueValues field.
func (o *RequestMonitorConfig) SetUniqueValues(v UniqueValuesMonitorRequestConfig) {
	o.UniqueValues = &v
}

// GetDatatype returns the Datatype field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetDatatype() DatatypeMonitorRequestConfig {
	if o == nil || isNil(o.Datatype) {
		var ret DatatypeMonitorRequestConfig
		return ret
	}
	return *o.Datatype
}

// GetDatatypeOk returns a tuple with the Datatype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetDatatypeOk() (*DatatypeMonitorRequestConfig, bool) {
	if o == nil || isNil(o.Datatype) {
		return nil, false
	}
	return o.Datatype, true
}

// HasDatatype returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasDatatype() bool {
	if o != nil && !isNil(o.Datatype) {
		return true
	}

	return false
}

// SetDatatype gets a reference to the given DatatypeMonitorRequestConfig and assigns it to the Datatype field.
func (o *RequestMonitorConfig) SetDatatype(v DatatypeMonitorRequestConfig) {
	o.Datatype = &v
}

// GetSeasonalARIMA returns the SeasonalARIMA field value if set, zero value otherwise.
func (o *RequestMonitorConfig) GetSeasonalARIMA() SeasonalARIMARequestConfig {
	if o == nil || isNil(o.SeasonalARIMA) {
		var ret SeasonalARIMARequestConfig
		return ret
	}
	return *o.SeasonalARIMA
}

// GetSeasonalARIMAOk returns a tuple with the SeasonalARIMA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMonitorConfig) GetSeasonalARIMAOk() (*SeasonalARIMARequestConfig, bool) {
	if o == nil || isNil(o.SeasonalARIMA) {
		return nil, false
	}
	return o.SeasonalARIMA, true
}

// HasSeasonalARIMA returns a boolean if a field has been set.
func (o *RequestMonitorConfig) HasSeasonalARIMA() bool {
	if o != nil && !isNil(o.SeasonalARIMA) {
		return true
	}

	return false
}

// SetSeasonalARIMA gets a reference to the given SeasonalARIMARequestConfig and assigns it to the SeasonalARIMA field.
func (o *RequestMonitorConfig) SetSeasonalARIMA(v SeasonalARIMARequestConfig) {
	o.SeasonalARIMA = &v
}

func (o RequestMonitorConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestMonitorConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SchemaVersion) {
		toSerialize["schemaVersion"] = o.SchemaVersion
	}
	if o.NumStdDev.IsSet() {
		toSerialize["numStdDev"] = o.NumStdDev.Get()
	}
	if !isNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !isNil(o.MissingRecentData) {
		toSerialize["missingRecentData"] = o.MissingRecentData
	}
	if !isNil(o.MissingRecentProfiles) {
		toSerialize["missingRecentProfiles"] = o.MissingRecentProfiles
	}
	if !isNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	if !isNil(o.MissingValues) {
		toSerialize["missingValues"] = o.MissingValues
	}
	if !isNil(o.UniqueValues) {
		toSerialize["uniqueValues"] = o.UniqueValues
	}
	if !isNil(o.Datatype) {
		toSerialize["datatype"] = o.Datatype
	}
	if !isNil(o.SeasonalARIMA) {
		toSerialize["seasonalARIMA"] = o.SeasonalARIMA
	}
	return toSerialize, nil
}

type NullableRequestMonitorConfig struct {
	value *RequestMonitorConfig
	isSet bool
}

func (v NullableRequestMonitorConfig) Get() *RequestMonitorConfig {
	return v.value
}

func (v *NullableRequestMonitorConfig) Set(val *RequestMonitorConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMonitorConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMonitorConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMonitorConfig(val *RequestMonitorConfig) *NullableRequestMonitorConfig {
	return &NullableRequestMonitorConfig{value: val, isSet: true}
}

func (v NullableRequestMonitorConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMonitorConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


