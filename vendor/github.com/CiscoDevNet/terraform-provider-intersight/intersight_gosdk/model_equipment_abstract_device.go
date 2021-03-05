/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-24T06:47:07Z.
 *
 * API version: 1.0.9-3824
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EquipmentAbstractDevice Common attributes for inventory device in Intersight.
type EquipmentAbstractDevice struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Administrator defined name for the device.
	Name *string `json:"Name,omitempty"`
	// Unique identity of the device.
	Uuid *string `json:"Uuid,omitempty"`
	// Current running software version of the device.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentAbstractDevice EquipmentAbstractDevice

// NewEquipmentAbstractDevice instantiates a new EquipmentAbstractDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentAbstractDevice(classId string, objectType string) *EquipmentAbstractDevice {
	this := EquipmentAbstractDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentAbstractDeviceWithDefaults instantiates a new EquipmentAbstractDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentAbstractDeviceWithDefaults() *EquipmentAbstractDevice {
	this := EquipmentAbstractDevice{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentAbstractDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentAbstractDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentAbstractDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentAbstractDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentAbstractDevice) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *EquipmentAbstractDevice) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EquipmentAbstractDevice) SetVersion(v string) {
	o.Version = &v
}

func (o EquipmentAbstractDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentAbstractDevice) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentAbstractDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Administrator defined name for the device.
		Name *string `json:"Name,omitempty"`
		// Unique identity of the device.
		Uuid *string `json:"Uuid,omitempty"`
		// Current running software version of the device.
		Version *string `json:"Version,omitempty"`
	}

	varEquipmentAbstractDeviceWithoutEmbeddedStruct := EquipmentAbstractDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentAbstractDeviceWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentAbstractDevice := _EquipmentAbstractDevice{}
		varEquipmentAbstractDevice.ClassId = varEquipmentAbstractDeviceWithoutEmbeddedStruct.ClassId
		varEquipmentAbstractDevice.ObjectType = varEquipmentAbstractDeviceWithoutEmbeddedStruct.ObjectType
		varEquipmentAbstractDevice.Name = varEquipmentAbstractDeviceWithoutEmbeddedStruct.Name
		varEquipmentAbstractDevice.Uuid = varEquipmentAbstractDeviceWithoutEmbeddedStruct.Uuid
		varEquipmentAbstractDevice.Version = varEquipmentAbstractDeviceWithoutEmbeddedStruct.Version
		*o = EquipmentAbstractDevice(varEquipmentAbstractDevice)
	} else {
		return err
	}

	varEquipmentAbstractDevice := _EquipmentAbstractDevice{}

	err = json.Unmarshal(bytes, &varEquipmentAbstractDevice)
	if err == nil {
		o.EquipmentBase = varEquipmentAbstractDevice.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentAbstractDevice struct {
	value *EquipmentAbstractDevice
	isSet bool
}

func (v NullableEquipmentAbstractDevice) Get() *EquipmentAbstractDevice {
	return v.value
}

func (v *NullableEquipmentAbstractDevice) Set(val *EquipmentAbstractDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentAbstractDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentAbstractDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentAbstractDevice(val *EquipmentAbstractDevice) *NullableEquipmentAbstractDevice {
	return &NullableEquipmentAbstractDevice{value: val, isSet: true}
}

func (v NullableEquipmentAbstractDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentAbstractDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
