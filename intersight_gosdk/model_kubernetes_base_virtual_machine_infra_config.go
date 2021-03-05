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

// KubernetesBaseVirtualMachineInfraConfig Infrastructure provider allocation configuration for virtual machine based Kubernetes nodes.
type KubernetesBaseVirtualMachineInfraConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType           string   `json:"ObjectType"`
	Interfaces           []string `json:"Interfaces,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesBaseVirtualMachineInfraConfig KubernetesBaseVirtualMachineInfraConfig

// NewKubernetesBaseVirtualMachineInfraConfig instantiates a new KubernetesBaseVirtualMachineInfraConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesBaseVirtualMachineInfraConfig(classId string, objectType string) *KubernetesBaseVirtualMachineInfraConfig {
	this := KubernetesBaseVirtualMachineInfraConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesBaseVirtualMachineInfraConfigWithDefaults instantiates a new KubernetesBaseVirtualMachineInfraConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesBaseVirtualMachineInfraConfigWithDefaults() *KubernetesBaseVirtualMachineInfraConfig {
	this := KubernetesBaseVirtualMachineInfraConfig{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesBaseVirtualMachineInfraConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseVirtualMachineInfraConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesBaseVirtualMachineInfraConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesBaseVirtualMachineInfraConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseVirtualMachineInfraConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesBaseVirtualMachineInfraConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesBaseVirtualMachineInfraConfig) GetInterfaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesBaseVirtualMachineInfraConfig) GetInterfacesOk() (*[]string, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return &o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *KubernetesBaseVirtualMachineInfraConfig) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *KubernetesBaseVirtualMachineInfraConfig) SetInterfaces(v []string) {
	o.Interfaces = v
}

func (o KubernetesBaseVirtualMachineInfraConfig) MarshalJSON() ([]byte, error) {
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
	if o.Interfaces != nil {
		toSerialize["Interfaces"] = o.Interfaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesBaseVirtualMachineInfraConfig) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesBaseVirtualMachineInfraConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string   `json:"ObjectType"`
		Interfaces []string `json:"Interfaces,omitempty"`
	}

	varKubernetesBaseVirtualMachineInfraConfigWithoutEmbeddedStruct := KubernetesBaseVirtualMachineInfraConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesBaseVirtualMachineInfraConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesBaseVirtualMachineInfraConfig := _KubernetesBaseVirtualMachineInfraConfig{}
		varKubernetesBaseVirtualMachineInfraConfig.ClassId = varKubernetesBaseVirtualMachineInfraConfigWithoutEmbeddedStruct.ClassId
		varKubernetesBaseVirtualMachineInfraConfig.ObjectType = varKubernetesBaseVirtualMachineInfraConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesBaseVirtualMachineInfraConfig.Interfaces = varKubernetesBaseVirtualMachineInfraConfigWithoutEmbeddedStruct.Interfaces
		*o = KubernetesBaseVirtualMachineInfraConfig(varKubernetesBaseVirtualMachineInfraConfig)
	} else {
		return err
	}

	varKubernetesBaseVirtualMachineInfraConfig := _KubernetesBaseVirtualMachineInfraConfig{}

	err = json.Unmarshal(bytes, &varKubernetesBaseVirtualMachineInfraConfig)
	if err == nil {
		o.MoBaseComplexType = varKubernetesBaseVirtualMachineInfraConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Interfaces")

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

type NullableKubernetesBaseVirtualMachineInfraConfig struct {
	value *KubernetesBaseVirtualMachineInfraConfig
	isSet bool
}

func (v NullableKubernetesBaseVirtualMachineInfraConfig) Get() *KubernetesBaseVirtualMachineInfraConfig {
	return v.value
}

func (v *NullableKubernetesBaseVirtualMachineInfraConfig) Set(val *KubernetesBaseVirtualMachineInfraConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesBaseVirtualMachineInfraConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesBaseVirtualMachineInfraConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesBaseVirtualMachineInfraConfig(val *KubernetesBaseVirtualMachineInfraConfig) *NullableKubernetesBaseVirtualMachineInfraConfig {
	return &NullableKubernetesBaseVirtualMachineInfraConfig{value: val, isSet: true}
}

func (v NullableKubernetesBaseVirtualMachineInfraConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesBaseVirtualMachineInfraConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
