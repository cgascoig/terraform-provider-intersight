/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// FirmwareSwitchUpgradeAllOf Definition of the list of properties defined in 'firmware.SwitchUpgrade', excluding properties defined in parent classes.
type FirmwareSwitchUpgradeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen.
	EnableFabricEvacuation *bool                                `json:"EnableFabricEvacuation,omitempty"`
	Device                 *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements      []NetworkElementRelationship `json:"NetworkElements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareSwitchUpgradeAllOf FirmwareSwitchUpgradeAllOf

// NewFirmwareSwitchUpgradeAllOf instantiates a new FirmwareSwitchUpgradeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareSwitchUpgradeAllOf(classId string, objectType string) *FirmwareSwitchUpgradeAllOf {
	this := FirmwareSwitchUpgradeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	return &this
}

// NewFirmwareSwitchUpgradeAllOfWithDefaults instantiates a new FirmwareSwitchUpgradeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareSwitchUpgradeAllOfWithDefaults() *FirmwareSwitchUpgradeAllOf {
	this := FirmwareSwitchUpgradeAllOf{}
	var classId string = "firmware.SwitchUpgrade"
	this.ClassId = classId
	var objectType string = "firmware.SwitchUpgrade"
	this.ObjectType = objectType
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareSwitchUpgradeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgradeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareSwitchUpgradeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareSwitchUpgradeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgradeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareSwitchUpgradeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableFabricEvacuation returns the EnableFabricEvacuation field value if set, zero value otherwise.
func (o *FirmwareSwitchUpgradeAllOf) GetEnableFabricEvacuation() bool {
	if o == nil || o.EnableFabricEvacuation == nil {
		var ret bool
		return ret
	}
	return *o.EnableFabricEvacuation
}

// GetEnableFabricEvacuationOk returns a tuple with the EnableFabricEvacuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgradeAllOf) GetEnableFabricEvacuationOk() (*bool, bool) {
	if o == nil || o.EnableFabricEvacuation == nil {
		return nil, false
	}
	return o.EnableFabricEvacuation, true
}

// HasEnableFabricEvacuation returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgradeAllOf) HasEnableFabricEvacuation() bool {
	if o != nil && o.EnableFabricEvacuation != nil {
		return true
	}

	return false
}

// SetEnableFabricEvacuation gets a reference to the given bool and assigns it to the EnableFabricEvacuation field.
func (o *FirmwareSwitchUpgradeAllOf) SetEnableFabricEvacuation(v bool) {
	o.EnableFabricEvacuation = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FirmwareSwitchUpgradeAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgradeAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgradeAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareSwitchUpgradeAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareSwitchUpgradeAllOf) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareSwitchUpgradeAllOf) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElements == nil {
		return nil, false
	}
	return &o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgradeAllOf) HasNetworkElements() bool {
	if o != nil && o.NetworkElements != nil {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *FirmwareSwitchUpgradeAllOf) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

func (o FirmwareSwitchUpgradeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnableFabricEvacuation != nil {
		toSerialize["EnableFabricEvacuation"] = o.EnableFabricEvacuation
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareSwitchUpgradeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareSwitchUpgradeAllOf := _FirmwareSwitchUpgradeAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareSwitchUpgradeAllOf); err == nil {
		*o = FirmwareSwitchUpgradeAllOf(varFirmwareSwitchUpgradeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableFabricEvacuation")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "NetworkElements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareSwitchUpgradeAllOf struct {
	value *FirmwareSwitchUpgradeAllOf
	isSet bool
}

func (v NullableFirmwareSwitchUpgradeAllOf) Get() *FirmwareSwitchUpgradeAllOf {
	return v.value
}

func (v *NullableFirmwareSwitchUpgradeAllOf) Set(val *FirmwareSwitchUpgradeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareSwitchUpgradeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareSwitchUpgradeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareSwitchUpgradeAllOf(val *FirmwareSwitchUpgradeAllOf) *NullableFirmwareSwitchUpgradeAllOf {
	return &NullableFirmwareSwitchUpgradeAllOf{value: val, isSet: true}
}

func (v NullableFirmwareSwitchUpgradeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareSwitchUpgradeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
