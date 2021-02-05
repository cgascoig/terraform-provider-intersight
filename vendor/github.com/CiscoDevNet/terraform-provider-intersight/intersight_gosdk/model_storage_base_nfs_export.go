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

// StorageBaseNfsExport Generic NFS Export object.
type StorageBaseNfsExport struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the NFS export in storage array.
	Name *string `json:"Name,omitempty"`
	// The uuid of this NFS export.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseNfsExport StorageBaseNfsExport

// NewStorageBaseNfsExport instantiates a new StorageBaseNfsExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseNfsExport(classId string, objectType string) *StorageBaseNfsExport {
	this := StorageBaseNfsExport{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseNfsExportWithDefaults instantiates a new StorageBaseNfsExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseNfsExportWithDefaults() *StorageBaseNfsExport {
	this := StorageBaseNfsExport{}
	var classId string = "storage.NetAppExportPolicy"
	this.ClassId = classId
	var objectType string = "storage.NetAppExportPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseNfsExport) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseNfsExport) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseNfsExport) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseNfsExport) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseNfsExport) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseNfsExport) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseNfsExport) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseNfsExport) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseNfsExport) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseNfsExport) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageBaseNfsExport) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseNfsExport) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageBaseNfsExport) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageBaseNfsExport) SetUuid(v string) {
	o.Uuid = &v
}

func (o StorageBaseNfsExport) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseNfsExport) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseNfsExportWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Name of the NFS export in storage array.
		Name *string `json:"Name,omitempty"`
		// The uuid of this NFS export.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varStorageBaseNfsExportWithoutEmbeddedStruct := StorageBaseNfsExportWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseNfsExportWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseNfsExport := _StorageBaseNfsExport{}
		varStorageBaseNfsExport.ClassId = varStorageBaseNfsExportWithoutEmbeddedStruct.ClassId
		varStorageBaseNfsExport.ObjectType = varStorageBaseNfsExportWithoutEmbeddedStruct.ObjectType
		varStorageBaseNfsExport.Name = varStorageBaseNfsExportWithoutEmbeddedStruct.Name
		varStorageBaseNfsExport.Uuid = varStorageBaseNfsExportWithoutEmbeddedStruct.Uuid
		*o = StorageBaseNfsExport(varStorageBaseNfsExport)
	} else {
		return err
	}

	varStorageBaseNfsExport := _StorageBaseNfsExport{}

	err = json.Unmarshal(bytes, &varStorageBaseNfsExport)
	if err == nil {
		o.MoBaseMo = varStorageBaseNfsExport.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Uuid")

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

type NullableStorageBaseNfsExport struct {
	value *StorageBaseNfsExport
	isSet bool
}

func (v NullableStorageBaseNfsExport) Get() *StorageBaseNfsExport {
	return v.value
}

func (v *NullableStorageBaseNfsExport) Set(val *StorageBaseNfsExport) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseNfsExport) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseNfsExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseNfsExport(val *StorageBaseNfsExport) *NullableStorageBaseNfsExport {
	return &NullableStorageBaseNfsExport{value: val, isSet: true}
}

func (v NullableStorageBaseNfsExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseNfsExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
