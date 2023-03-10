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

// checks if the DTOS3StorageInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DTOS3StorageInfoDTO{}

// DTOS3StorageInfoDTO struct for DTOS3StorageInfoDTO
type DTOS3StorageInfoDTO struct {
	Destination *string `json:"destination,omitempty"`
	Region *string `json:"region,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	EnableEncryption *bool `json:"enableEncryption,omitempty"`
	EncryptionType *string `json:"encryptionType,omitempty"`
	KmsKey *string `json:"kmsKey,omitempty"`
	CannedAcl *string `json:"cannedAcl,omitempty"`
}

// NewDTOS3StorageInfoDTO instantiates a new DTOS3StorageInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDTOS3StorageInfoDTO() *DTOS3StorageInfoDTO {
	this := DTOS3StorageInfoDTO{}
	return &this
}

// NewDTOS3StorageInfoDTOWithDefaults instantiates a new DTOS3StorageInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDTOS3StorageInfoDTOWithDefaults() *DTOS3StorageInfoDTO {
	this := DTOS3StorageInfoDTO{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *DTOS3StorageInfoDTO) GetDestination() string {
	if o == nil || isNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTOS3StorageInfoDTO) GetDestinationOk() (*string, bool) {
	if o == nil || isNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *DTOS3StorageInfoDTO) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *DTOS3StorageInfoDTO) SetDestination(v string) {
	o.Destination = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DTOS3StorageInfoDTO) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTOS3StorageInfoDTO) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DTOS3StorageInfoDTO) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DTOS3StorageInfoDTO) SetRegion(v string) {
	o.Region = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *DTOS3StorageInfoDTO) GetEndpoint() string {
	if o == nil || isNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTOS3StorageInfoDTO) GetEndpointOk() (*string, bool) {
	if o == nil || isNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *DTOS3StorageInfoDTO) HasEndpoint() bool {
	if o != nil && !isNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *DTOS3StorageInfoDTO) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetEnableEncryption returns the EnableEncryption field value if set, zero value otherwise.
func (o *DTOS3StorageInfoDTO) GetEnableEncryption() bool {
	if o == nil || isNil(o.EnableEncryption) {
		var ret bool
		return ret
	}
	return *o.EnableEncryption
}

// GetEnableEncryptionOk returns a tuple with the EnableEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTOS3StorageInfoDTO) GetEnableEncryptionOk() (*bool, bool) {
	if o == nil || isNil(o.EnableEncryption) {
		return nil, false
	}
	return o.EnableEncryption, true
}

// HasEnableEncryption returns a boolean if a field has been set.
func (o *DTOS3StorageInfoDTO) HasEnableEncryption() bool {
	if o != nil && !isNil(o.EnableEncryption) {
		return true
	}

	return false
}

// SetEnableEncryption gets a reference to the given bool and assigns it to the EnableEncryption field.
func (o *DTOS3StorageInfoDTO) SetEnableEncryption(v bool) {
	o.EnableEncryption = &v
}

// GetEncryptionType returns the EncryptionType field value if set, zero value otherwise.
func (o *DTOS3StorageInfoDTO) GetEncryptionType() string {
	if o == nil || isNil(o.EncryptionType) {
		var ret string
		return ret
	}
	return *o.EncryptionType
}

// GetEncryptionTypeOk returns a tuple with the EncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTOS3StorageInfoDTO) GetEncryptionTypeOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionType) {
		return nil, false
	}
	return o.EncryptionType, true
}

// HasEncryptionType returns a boolean if a field has been set.
func (o *DTOS3StorageInfoDTO) HasEncryptionType() bool {
	if o != nil && !isNil(o.EncryptionType) {
		return true
	}

	return false
}

// SetEncryptionType gets a reference to the given string and assigns it to the EncryptionType field.
func (o *DTOS3StorageInfoDTO) SetEncryptionType(v string) {
	o.EncryptionType = &v
}

// GetKmsKey returns the KmsKey field value if set, zero value otherwise.
func (o *DTOS3StorageInfoDTO) GetKmsKey() string {
	if o == nil || isNil(o.KmsKey) {
		var ret string
		return ret
	}
	return *o.KmsKey
}

// GetKmsKeyOk returns a tuple with the KmsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTOS3StorageInfoDTO) GetKmsKeyOk() (*string, bool) {
	if o == nil || isNil(o.KmsKey) {
		return nil, false
	}
	return o.KmsKey, true
}

// HasKmsKey returns a boolean if a field has been set.
func (o *DTOS3StorageInfoDTO) HasKmsKey() bool {
	if o != nil && !isNil(o.KmsKey) {
		return true
	}

	return false
}

// SetKmsKey gets a reference to the given string and assigns it to the KmsKey field.
func (o *DTOS3StorageInfoDTO) SetKmsKey(v string) {
	o.KmsKey = &v
}

// GetCannedAcl returns the CannedAcl field value if set, zero value otherwise.
func (o *DTOS3StorageInfoDTO) GetCannedAcl() string {
	if o == nil || isNil(o.CannedAcl) {
		var ret string
		return ret
	}
	return *o.CannedAcl
}

// GetCannedAclOk returns a tuple with the CannedAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTOS3StorageInfoDTO) GetCannedAclOk() (*string, bool) {
	if o == nil || isNil(o.CannedAcl) {
		return nil, false
	}
	return o.CannedAcl, true
}

// HasCannedAcl returns a boolean if a field has been set.
func (o *DTOS3StorageInfoDTO) HasCannedAcl() bool {
	if o != nil && !isNil(o.CannedAcl) {
		return true
	}

	return false
}

// SetCannedAcl gets a reference to the given string and assigns it to the CannedAcl field.
func (o *DTOS3StorageInfoDTO) SetCannedAcl(v string) {
	o.CannedAcl = &v
}

func (o DTOS3StorageInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DTOS3StorageInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !isNil(o.EnableEncryption) {
		toSerialize["enableEncryption"] = o.EnableEncryption
	}
	if !isNil(o.EncryptionType) {
		toSerialize["encryptionType"] = o.EncryptionType
	}
	if !isNil(o.KmsKey) {
		toSerialize["kmsKey"] = o.KmsKey
	}
	if !isNil(o.CannedAcl) {
		toSerialize["cannedAcl"] = o.CannedAcl
	}
	return toSerialize, nil
}

type NullableDTOS3StorageInfoDTO struct {
	value *DTOS3StorageInfoDTO
	isSet bool
}

func (v NullableDTOS3StorageInfoDTO) Get() *DTOS3StorageInfoDTO {
	return v.value
}

func (v *NullableDTOS3StorageInfoDTO) Set(val *DTOS3StorageInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDTOS3StorageInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDTOS3StorageInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTOS3StorageInfoDTO(val *DTOS3StorageInfoDTO) *NullableDTOS3StorageInfoDTO {
	return &NullableDTOS3StorageInfoDTO{value: val, isSet: true}
}

func (v NullableDTOS3StorageInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTOS3StorageInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


