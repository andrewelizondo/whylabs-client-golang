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

// DTOEbsVolumeTypeDTO the model 'DTOEbsVolumeTypeDTO'
type DTOEbsVolumeTypeDTO string

// List of DTO.EbsVolumeTypeDTO
const (
	GENERAL_PURPOSE_SSD DTOEbsVolumeTypeDTO = "GENERAL_PURPOSE_SSD"
	THROUGHPUT_OPTIMIZED_HDD DTOEbsVolumeTypeDTO = "THROUGHPUT_OPTIMIZED_HDD"
)

// All allowed values of DTOEbsVolumeTypeDTO enum
var AllowedDTOEbsVolumeTypeDTOEnumValues = []DTOEbsVolumeTypeDTO{
	"GENERAL_PURPOSE_SSD",
	"THROUGHPUT_OPTIMIZED_HDD",
}

func (v *DTOEbsVolumeTypeDTO) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DTOEbsVolumeTypeDTO(value)
	for _, existing := range AllowedDTOEbsVolumeTypeDTOEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DTOEbsVolumeTypeDTO", value)
}

// NewDTOEbsVolumeTypeDTOFromValue returns a pointer to a valid DTOEbsVolumeTypeDTO
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDTOEbsVolumeTypeDTOFromValue(v string) (*DTOEbsVolumeTypeDTO, error) {
	ev := DTOEbsVolumeTypeDTO(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DTOEbsVolumeTypeDTO: valid values are %v", v, AllowedDTOEbsVolumeTypeDTOEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DTOEbsVolumeTypeDTO) IsValid() bool {
	for _, existing := range AllowedDTOEbsVolumeTypeDTOEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DTO.EbsVolumeTypeDTO value
func (v DTOEbsVolumeTypeDTO) Ptr() *DTOEbsVolumeTypeDTO {
	return &v
}

type NullableDTOEbsVolumeTypeDTO struct {
	value *DTOEbsVolumeTypeDTO
	isSet bool
}

func (v NullableDTOEbsVolumeTypeDTO) Get() *DTOEbsVolumeTypeDTO {
	return v.value
}

func (v *NullableDTOEbsVolumeTypeDTO) Set(val *DTOEbsVolumeTypeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDTOEbsVolumeTypeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDTOEbsVolumeTypeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTOEbsVolumeTypeDTO(val *DTOEbsVolumeTypeDTO) *NullableDTOEbsVolumeTypeDTO {
	return &NullableDTOEbsVolumeTypeDTO{value: val, isSet: true}
}

func (v NullableDTOEbsVolumeTypeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTOEbsVolumeTypeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

