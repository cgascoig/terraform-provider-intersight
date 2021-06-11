/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-05-25T18:18:54Z.
 *
 * API version: 1.0.9-4305
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// PoolAbstractLease Concrete Children of Abstract Lease represent a single ID that is part of an ID universe. Concrete Children of this class will implement ID request action.
type PoolAbstractLease struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Type of the lease allocation either static or dynamic (i.e via pool). * `dynamic` - Identifiers to be allocated by system. * `static` - Identifiers are assigned by the user.
	AllocationType       *string `json:"AllocationType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractLease PoolAbstractLease

// NewPoolAbstractLease instantiates a new PoolAbstractLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractLease(classId string, objectType string) *PoolAbstractLease {
	this := PoolAbstractLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	return &this
}

// NewPoolAbstractLeaseWithDefaults instantiates a new PoolAbstractLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractLeaseWithDefaults() *PoolAbstractLease {
	this := PoolAbstractLease{}
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractLease) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllocationType returns the AllocationType field value if set, zero value otherwise.
func (o *PoolAbstractLease) GetAllocationType() string {
	if o == nil || o.AllocationType == nil {
		var ret string
		return ret
	}
	return *o.AllocationType
}

// GetAllocationTypeOk returns a tuple with the AllocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractLease) GetAllocationTypeOk() (*string, bool) {
	if o == nil || o.AllocationType == nil {
		return nil, false
	}
	return o.AllocationType, true
}

// HasAllocationType returns a boolean if a field has been set.
func (o *PoolAbstractLease) HasAllocationType() bool {
	if o != nil && o.AllocationType != nil {
		return true
	}

	return false
}

// SetAllocationType gets a reference to the given string and assigns it to the AllocationType field.
func (o *PoolAbstractLease) SetAllocationType(v string) {
	o.AllocationType = &v
}

func (o PoolAbstractLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllocationType != nil {
		toSerialize["AllocationType"] = o.AllocationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractLease) UnmarshalJSON(bytes []byte) (err error) {
	type PoolAbstractLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Type of the lease allocation either static or dynamic (i.e via pool). * `dynamic` - Identifiers to be allocated by system. * `static` - Identifiers are assigned by the user.
		AllocationType *string `json:"AllocationType,omitempty"`
	}

	varPoolAbstractLeaseWithoutEmbeddedStruct := PoolAbstractLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPoolAbstractLeaseWithoutEmbeddedStruct)
	if err == nil {
		varPoolAbstractLease := _PoolAbstractLease{}
		varPoolAbstractLease.ClassId = varPoolAbstractLeaseWithoutEmbeddedStruct.ClassId
		varPoolAbstractLease.ObjectType = varPoolAbstractLeaseWithoutEmbeddedStruct.ObjectType
		varPoolAbstractLease.AllocationType = varPoolAbstractLeaseWithoutEmbeddedStruct.AllocationType
		*o = PoolAbstractLease(varPoolAbstractLease)
	} else {
		return err
	}

	varPoolAbstractLease := _PoolAbstractLease{}

	err = json.Unmarshal(bytes, &varPoolAbstractLease)
	if err == nil {
		o.MoBaseMo = varPoolAbstractLease.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllocationType")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullablePoolAbstractLease struct {
	value *PoolAbstractLease
	isSet bool
}

func (v NullablePoolAbstractLease) Get() *PoolAbstractLease {
	return v.value
}

func (v *NullablePoolAbstractLease) Set(val *PoolAbstractLease) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractLease) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractLease(val *PoolAbstractLease) *NullablePoolAbstractLease {
	return &NullablePoolAbstractLease{value: val, isSet: true}
}

func (v NullablePoolAbstractLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
