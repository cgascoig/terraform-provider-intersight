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

// PoolAbstractPoolMember PoolMember represents a single identifier that is part of a pool.
type PoolAbstractPoolMember struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Boolean to represent whether the ID is assigned or not.
	Assigned             *bool `json:"Assigned,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractPoolMember PoolAbstractPoolMember

// NewPoolAbstractPoolMember instantiates a new PoolAbstractPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractPoolMember(classId string, objectType string) *PoolAbstractPoolMember {
	this := PoolAbstractPoolMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	return &this
}

// NewPoolAbstractPoolMemberWithDefaults instantiates a new PoolAbstractPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractPoolMemberWithDefaults() *PoolAbstractPoolMember {
	this := PoolAbstractPoolMember{}
	var assigned bool = false
	this.Assigned = &assigned
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *PoolAbstractPoolMember) GetAssigned() bool {
	if o == nil || o.Assigned == nil {
		var ret bool
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolMember) GetAssignedOk() (*bool, bool) {
	if o == nil || o.Assigned == nil {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *PoolAbstractPoolMember) HasAssigned() bool {
	if o != nil && o.Assigned != nil {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given bool and assigns it to the Assigned field.
func (o *PoolAbstractPoolMember) SetAssigned(v bool) {
	o.Assigned = &v
}

func (o PoolAbstractPoolMember) MarshalJSON() ([]byte, error) {
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
	if o.Assigned != nil {
		toSerialize["Assigned"] = o.Assigned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractPoolMember) UnmarshalJSON(bytes []byte) (err error) {
	type PoolAbstractPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Boolean to represent whether the ID is assigned or not.
		Assigned *bool `json:"Assigned,omitempty"`
	}

	varPoolAbstractPoolMemberWithoutEmbeddedStruct := PoolAbstractPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPoolAbstractPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varPoolAbstractPoolMember := _PoolAbstractPoolMember{}
		varPoolAbstractPoolMember.ClassId = varPoolAbstractPoolMemberWithoutEmbeddedStruct.ClassId
		varPoolAbstractPoolMember.ObjectType = varPoolAbstractPoolMemberWithoutEmbeddedStruct.ObjectType
		varPoolAbstractPoolMember.Assigned = varPoolAbstractPoolMemberWithoutEmbeddedStruct.Assigned
		*o = PoolAbstractPoolMember(varPoolAbstractPoolMember)
	} else {
		return err
	}

	varPoolAbstractPoolMember := _PoolAbstractPoolMember{}

	err = json.Unmarshal(bytes, &varPoolAbstractPoolMember)
	if err == nil {
		o.MoBaseMo = varPoolAbstractPoolMember.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Assigned")

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

type NullablePoolAbstractPoolMember struct {
	value *PoolAbstractPoolMember
	isSet bool
}

func (v NullablePoolAbstractPoolMember) Get() *PoolAbstractPoolMember {
	return v.value
}

func (v *NullablePoolAbstractPoolMember) Set(val *PoolAbstractPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractPoolMember(val *PoolAbstractPoolMember) *NullablePoolAbstractPoolMember {
	return &NullablePoolAbstractPoolMember{value: val, isSet: true}
}

func (v NullablePoolAbstractPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
