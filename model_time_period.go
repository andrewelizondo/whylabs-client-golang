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
	"fmt"
)

// TimePeriod A TimePeriod represents the bucketing that a dataset has undergone.
type TimePeriod string

// List of TimePeriod
const (
	P1_M TimePeriod = "P1M"
	P1_W TimePeriod = "P1W"
	P1_D TimePeriod = "P1D"
	PT1_H TimePeriod = "PT1H"
	PT5_M TimePeriod = "PT5M"
)

// All allowed values of TimePeriod enum
var AllowedTimePeriodEnumValues = []TimePeriod{
	"P1M",
	"P1W",
	"P1D",
	"PT1H",
	"PT5M",
}

func (v *TimePeriod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimePeriod(value)
	for _, existing := range AllowedTimePeriodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimePeriod", value)
}

// NewTimePeriodFromValue returns a pointer to a valid TimePeriod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimePeriodFromValue(v string) (*TimePeriod, error) {
	ev := TimePeriod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimePeriod: valid values are %v", v, AllowedTimePeriodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimePeriod) IsValid() bool {
	for _, existing := range AllowedTimePeriodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimePeriod value
func (v TimePeriod) Ptr() *TimePeriod {
	return &v
}

type NullableTimePeriod struct {
	value *TimePeriod
	isSet bool
}

func (v NullableTimePeriod) Get() *TimePeriod {
	return v.value
}

func (v *NullableTimePeriod) Set(val *TimePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableTimePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableTimePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimePeriod(val *TimePeriod) *NullableTimePeriod {
	return &NullableTimePeriod{value: val, isSet: true}
}

func (v NullableTimePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

