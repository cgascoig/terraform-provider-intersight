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

// HyperflexVmInterface Virtual machine interfaces.
type HyperflexVmInterface struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Virtual machine brige name of interface.
	Bridge    *string  `json:"Bridge,omitempty"`
	IpAddress []string `json:"IpAddress,omitempty"`
	// Virtual machine interface mac address.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Virtual machine interface model.
	Model                *string `json:"Model,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVmInterface HyperflexVmInterface

// NewHyperflexVmInterface instantiates a new HyperflexVmInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmInterface(classId string, objectType string) *HyperflexVmInterface {
	this := HyperflexVmInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVmInterfaceWithDefaults instantiates a new HyperflexVmInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmInterfaceWithDefaults() *HyperflexVmInterface {
	this := HyperflexVmInterface{}
	var classId string = "hyperflex.VmInterface"
	this.ClassId = classId
	var objectType string = "hyperflex.VmInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVmInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVmInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVmInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVmInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBridge returns the Bridge field value if set, zero value otherwise.
func (o *HyperflexVmInterface) GetBridge() string {
	if o == nil || o.Bridge == nil {
		var ret string
		return ret
	}
	return *o.Bridge
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmInterface) GetBridgeOk() (*string, bool) {
	if o == nil || o.Bridge == nil {
		return nil, false
	}
	return o.Bridge, true
}

// HasBridge returns a boolean if a field has been set.
func (o *HyperflexVmInterface) HasBridge() bool {
	if o != nil && o.Bridge != nil {
		return true
	}

	return false
}

// SetBridge gets a reference to the given string and assigns it to the Bridge field.
func (o *HyperflexVmInterface) SetBridge(v string) {
	o.Bridge = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmInterface) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmInterface) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *HyperflexVmInterface) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *HyperflexVmInterface) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *HyperflexVmInterface) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmInterface) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *HyperflexVmInterface) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *HyperflexVmInterface) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *HyperflexVmInterface) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmInterface) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *HyperflexVmInterface) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *HyperflexVmInterface) SetModel(v string) {
	o.Model = &v
}

func (o HyperflexVmInterface) MarshalJSON() ([]byte, error) {
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
	if o.Bridge != nil {
		toSerialize["Bridge"] = o.Bridge
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVmInterface) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexVmInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Virtual machine brige name of interface.
		Bridge    *string  `json:"Bridge,omitempty"`
		IpAddress []string `json:"IpAddress,omitempty"`
		// Virtual machine interface mac address.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Virtual machine interface model.
		Model *string `json:"Model,omitempty"`
	}

	varHyperflexVmInterfaceWithoutEmbeddedStruct := HyperflexVmInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexVmInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexVmInterface := _HyperflexVmInterface{}
		varHyperflexVmInterface.ClassId = varHyperflexVmInterfaceWithoutEmbeddedStruct.ClassId
		varHyperflexVmInterface.ObjectType = varHyperflexVmInterfaceWithoutEmbeddedStruct.ObjectType
		varHyperflexVmInterface.Bridge = varHyperflexVmInterfaceWithoutEmbeddedStruct.Bridge
		varHyperflexVmInterface.IpAddress = varHyperflexVmInterfaceWithoutEmbeddedStruct.IpAddress
		varHyperflexVmInterface.MacAddress = varHyperflexVmInterfaceWithoutEmbeddedStruct.MacAddress
		varHyperflexVmInterface.Model = varHyperflexVmInterfaceWithoutEmbeddedStruct.Model
		*o = HyperflexVmInterface(varHyperflexVmInterface)
	} else {
		return err
	}

	varHyperflexVmInterface := _HyperflexVmInterface{}

	err = json.Unmarshal(bytes, &varHyperflexVmInterface)
	if err == nil {
		o.MoBaseComplexType = varHyperflexVmInterface.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bridge")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Model")

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

type NullableHyperflexVmInterface struct {
	value *HyperflexVmInterface
	isSet bool
}

func (v NullableHyperflexVmInterface) Get() *HyperflexVmInterface {
	return v.value
}

func (v *NullableHyperflexVmInterface) Set(val *HyperflexVmInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmInterface(val *HyperflexVmInterface) *NullableHyperflexVmInterface {
	return &NullableHyperflexVmInterface{value: val, isSet: true}
}

func (v NullableHyperflexVmInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
