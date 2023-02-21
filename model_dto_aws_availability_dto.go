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

// DTOAwsAvailabilityDTO the model 'DTOAwsAvailabilityDTO'
type DTOAwsAvailabilityDTO string

// List of DTO.AwsAvailabilityDTO
const (
	SPOT DTOAwsAvailabilityDTO = "SPOT"
	ON_DEMAND DTOAwsAvailabilityDTO = "ON_DEMAND"
	SPOT_WITH_FALLBACK DTOAwsAvailabilityDTO = "SPOT_WITH_FALLBACK"
)

// All allowed values of DTOAwsAvailabilityDTO enum
var AllowedDTOAwsAvailabilityDTOEnumValues = []DTOAwsAvailabilityDTO{
	"SPOT",
	"ON_DEMAND",
	"SPOT_WITH_FALLBACK",
}

func (v *DTOAwsAvailabilityDTO) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DTOAwsAvailabilityDTO(value)
	for _, existing := range AllowedDTOAwsAvailabilityDTOEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DTOAwsAvailabilityDTO", value)
}

// NewDTOAwsAvailabilityDTOFromValue returns a pointer to a valid DTOAwsAvailabilityDTO
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDTOAwsAvailabilityDTOFromValue(v string) (*DTOAwsAvailabilityDTO, error) {
	ev := DTOAwsAvailabilityDTO(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DTOAwsAvailabilityDTO: valid values are %v", v, AllowedDTOAwsAvailabilityDTOEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DTOAwsAvailabilityDTO) IsValid() bool {
	for _, existing := range AllowedDTOAwsAvailabilityDTOEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DTO.AwsAvailabilityDTO value
func (v DTOAwsAvailabilityDTO) Ptr() *DTOAwsAvailabilityDTO {
	return &v
}

type NullableDTOAwsAvailabilityDTO struct {
	value *DTOAwsAvailabilityDTO
	isSet bool
}

func (v NullableDTOAwsAvailabilityDTO) Get() *DTOAwsAvailabilityDTO {
	return v.value
}

func (v *NullableDTOAwsAvailabilityDTO) Set(val *DTOAwsAvailabilityDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDTOAwsAvailabilityDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDTOAwsAvailabilityDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTOAwsAvailabilityDTO(val *DTOAwsAvailabilityDTO) *NullableDTOAwsAvailabilityDTO {
	return &NullableDTOAwsAvailabilityDTO{value: val, isSet: true}
}

func (v NullableDTOAwsAvailabilityDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTOAwsAvailabilityDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

