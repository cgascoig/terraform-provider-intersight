/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// CapabilitySwitchSystemLimits Object combines and lists the storage-config limits for each of the Fabric/Switch platforms.
type CapabilitySwitchSystemLimits struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect.
	MaximumChassisCount *int64 `json:"MaximumChassisCount,omitempty"`
	// Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect.
	MaximumFexPerDomain *int64 `json:"MaximumFexPerDomain,omitempty"`
	// Maximum UCS servers per Switch/Fabric-Interconnect.
	MaximumServersPerDomain *int64 `json:"MaximumServersPerDomain,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _CapabilitySwitchSystemLimits CapabilitySwitchSystemLimits

// NewCapabilitySwitchSystemLimits instantiates a new CapabilitySwitchSystemLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchSystemLimits(classId string, objectType string) *CapabilitySwitchSystemLimits {
	this := CapabilitySwitchSystemLimits{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilitySwitchSystemLimitsWithDefaults instantiates a new CapabilitySwitchSystemLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchSystemLimitsWithDefaults() *CapabilitySwitchSystemLimits {
	this := CapabilitySwitchSystemLimits{}
	var classId string = "capability.SwitchSystemLimits"
	this.ClassId = classId
	var objectType string = "capability.SwitchSystemLimits"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchSystemLimits) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimits) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchSystemLimits) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchSystemLimits) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimits) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchSystemLimits) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMaximumChassisCount returns the MaximumChassisCount field value if set, zero value otherwise.
func (o *CapabilitySwitchSystemLimits) GetMaximumChassisCount() int64 {
	if o == nil || o.MaximumChassisCount == nil {
		var ret int64
		return ret
	}
	return *o.MaximumChassisCount
}

// GetMaximumChassisCountOk returns a tuple with the MaximumChassisCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimits) GetMaximumChassisCountOk() (*int64, bool) {
	if o == nil || o.MaximumChassisCount == nil {
		return nil, false
	}
	return o.MaximumChassisCount, true
}

// HasMaximumChassisCount returns a boolean if a field has been set.
func (o *CapabilitySwitchSystemLimits) HasMaximumChassisCount() bool {
	if o != nil && o.MaximumChassisCount != nil {
		return true
	}

	return false
}

// SetMaximumChassisCount gets a reference to the given int64 and assigns it to the MaximumChassisCount field.
func (o *CapabilitySwitchSystemLimits) SetMaximumChassisCount(v int64) {
	o.MaximumChassisCount = &v
}

// GetMaximumFexPerDomain returns the MaximumFexPerDomain field value if set, zero value otherwise.
func (o *CapabilitySwitchSystemLimits) GetMaximumFexPerDomain() int64 {
	if o == nil || o.MaximumFexPerDomain == nil {
		var ret int64
		return ret
	}
	return *o.MaximumFexPerDomain
}

// GetMaximumFexPerDomainOk returns a tuple with the MaximumFexPerDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimits) GetMaximumFexPerDomainOk() (*int64, bool) {
	if o == nil || o.MaximumFexPerDomain == nil {
		return nil, false
	}
	return o.MaximumFexPerDomain, true
}

// HasMaximumFexPerDomain returns a boolean if a field has been set.
func (o *CapabilitySwitchSystemLimits) HasMaximumFexPerDomain() bool {
	if o != nil && o.MaximumFexPerDomain != nil {
		return true
	}

	return false
}

// SetMaximumFexPerDomain gets a reference to the given int64 and assigns it to the MaximumFexPerDomain field.
func (o *CapabilitySwitchSystemLimits) SetMaximumFexPerDomain(v int64) {
	o.MaximumFexPerDomain = &v
}

// GetMaximumServersPerDomain returns the MaximumServersPerDomain field value if set, zero value otherwise.
func (o *CapabilitySwitchSystemLimits) GetMaximumServersPerDomain() int64 {
	if o == nil || o.MaximumServersPerDomain == nil {
		var ret int64
		return ret
	}
	return *o.MaximumServersPerDomain
}

// GetMaximumServersPerDomainOk returns a tuple with the MaximumServersPerDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimits) GetMaximumServersPerDomainOk() (*int64, bool) {
	if o == nil || o.MaximumServersPerDomain == nil {
		return nil, false
	}
	return o.MaximumServersPerDomain, true
}

// HasMaximumServersPerDomain returns a boolean if a field has been set.
func (o *CapabilitySwitchSystemLimits) HasMaximumServersPerDomain() bool {
	if o != nil && o.MaximumServersPerDomain != nil {
		return true
	}

	return false
}

// SetMaximumServersPerDomain gets a reference to the given int64 and assigns it to the MaximumServersPerDomain field.
func (o *CapabilitySwitchSystemLimits) SetMaximumServersPerDomain(v int64) {
	o.MaximumServersPerDomain = &v
}

func (o CapabilitySwitchSystemLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MaximumChassisCount != nil {
		toSerialize["MaximumChassisCount"] = o.MaximumChassisCount
	}
	if o.MaximumFexPerDomain != nil {
		toSerialize["MaximumFexPerDomain"] = o.MaximumFexPerDomain
	}
	if o.MaximumServersPerDomain != nil {
		toSerialize["MaximumServersPerDomain"] = o.MaximumServersPerDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchSystemLimits) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySwitchSystemLimitsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect.
		MaximumChassisCount *int64 `json:"MaximumChassisCount,omitempty"`
		// Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect.
		MaximumFexPerDomain *int64 `json:"MaximumFexPerDomain,omitempty"`
		// Maximum UCS servers per Switch/Fabric-Interconnect.
		MaximumServersPerDomain *int64 `json:"MaximumServersPerDomain,omitempty"`
	}

	varCapabilitySwitchSystemLimitsWithoutEmbeddedStruct := CapabilitySwitchSystemLimitsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchSystemLimitsWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchSystemLimits := _CapabilitySwitchSystemLimits{}
		varCapabilitySwitchSystemLimits.ClassId = varCapabilitySwitchSystemLimitsWithoutEmbeddedStruct.ClassId
		varCapabilitySwitchSystemLimits.ObjectType = varCapabilitySwitchSystemLimitsWithoutEmbeddedStruct.ObjectType
		varCapabilitySwitchSystemLimits.MaximumChassisCount = varCapabilitySwitchSystemLimitsWithoutEmbeddedStruct.MaximumChassisCount
		varCapabilitySwitchSystemLimits.MaximumFexPerDomain = varCapabilitySwitchSystemLimitsWithoutEmbeddedStruct.MaximumFexPerDomain
		varCapabilitySwitchSystemLimits.MaximumServersPerDomain = varCapabilitySwitchSystemLimitsWithoutEmbeddedStruct.MaximumServersPerDomain
		*o = CapabilitySwitchSystemLimits(varCapabilitySwitchSystemLimits)
	} else {
		return err
	}

	varCapabilitySwitchSystemLimits := _CapabilitySwitchSystemLimits{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchSystemLimits)
	if err == nil {
		o.MoBaseComplexType = varCapabilitySwitchSystemLimits.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaximumChassisCount")
		delete(additionalProperties, "MaximumFexPerDomain")
		delete(additionalProperties, "MaximumServersPerDomain")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySwitchSystemLimits struct {
	value *CapabilitySwitchSystemLimits
	isSet bool
}

func (v NullableCapabilitySwitchSystemLimits) Get() *CapabilitySwitchSystemLimits {
	return v.value
}

func (v *NullableCapabilitySwitchSystemLimits) Set(val *CapabilitySwitchSystemLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchSystemLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchSystemLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchSystemLimits(val *CapabilitySwitchSystemLimits) *NullableCapabilitySwitchSystemLimits {
	return &NullableCapabilitySwitchSystemLimits{value: val, isSet: true}
}

func (v NullableCapabilitySwitchSystemLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchSystemLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
