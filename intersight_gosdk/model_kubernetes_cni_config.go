/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// KubernetesCniConfig Abstract class of any CNI configuration.
type KubernetesCniConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// CNI Image registry location.
	Registry *string `json:"Registry,omitempty"`
	// Container newtork interface plugin configuration.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesCniConfig KubernetesCniConfig

// NewKubernetesCniConfig instantiates a new KubernetesCniConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCniConfig(classId string, objectType string) *KubernetesCniConfig {
	this := KubernetesCniConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesCniConfigWithDefaults instantiates a new KubernetesCniConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCniConfigWithDefaults() *KubernetesCniConfig {
	this := KubernetesCniConfig{}
	var classId string = "kubernetes.CalicoConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.CalicoConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesCniConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesCniConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesCniConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesCniConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *KubernetesCniConfig) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfig) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *KubernetesCniConfig) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *KubernetesCniConfig) SetRegistry(v string) {
	o.Registry = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KubernetesCniConfig) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCniConfig) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesCniConfig) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KubernetesCniConfig) SetVersion(v string) {
	o.Version = &v
}

func (o KubernetesCniConfig) MarshalJSON() ([]byte, error) {
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
	if o.Registry != nil {
		toSerialize["Registry"] = o.Registry
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesCniConfig) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesCniConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// CNI Image registry location.
		Registry *string `json:"Registry,omitempty"`
		// Container newtork interface plugin configuration.
		Version *string `json:"Version,omitempty"`
	}

	varKubernetesCniConfigWithoutEmbeddedStruct := KubernetesCniConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesCniConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesCniConfig := _KubernetesCniConfig{}
		varKubernetesCniConfig.ClassId = varKubernetesCniConfigWithoutEmbeddedStruct.ClassId
		varKubernetesCniConfig.ObjectType = varKubernetesCniConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesCniConfig.Registry = varKubernetesCniConfigWithoutEmbeddedStruct.Registry
		varKubernetesCniConfig.Version = varKubernetesCniConfigWithoutEmbeddedStruct.Version
		*o = KubernetesCniConfig(varKubernetesCniConfig)
	} else {
		return err
	}

	varKubernetesCniConfig := _KubernetesCniConfig{}

	err = json.Unmarshal(bytes, &varKubernetesCniConfig)
	if err == nil {
		o.MoBaseComplexType = varKubernetesCniConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Registry")
		delete(additionalProperties, "Version")

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

type NullableKubernetesCniConfig struct {
	value *KubernetesCniConfig
	isSet bool
}

func (v NullableKubernetesCniConfig) Get() *KubernetesCniConfig {
	return v.value
}

func (v *NullableKubernetesCniConfig) Set(val *KubernetesCniConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCniConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCniConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCniConfig(val *KubernetesCniConfig) *NullableKubernetesCniConfig {
	return &NullableKubernetesCniConfig{value: val, isSet: true}
}

func (v NullableKubernetesCniConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCniConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
