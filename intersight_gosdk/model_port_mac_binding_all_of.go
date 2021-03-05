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
)

// PortMacBindingAllOf Definition of the list of properties defined in 'port.MacBinding', excluding properties defined in parent classes.
type PortMacBindingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Aggregate Port ID of the local Switch Interface.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// Chassis/FEX device idetifier that is local to a cluster.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// Device ID value that is advertised and available as a part of LLDP TLV.
	DeviceMac *string `json:"DeviceMac,omitempty"`
	// Port ID of the local Switch Interface.
	PortId *int64 `json:"PortId,omitempty"`
	// Port ID value that is advertised and available as a part of LLDP TLV.
	PortMac *string `json:"PortMac,omitempty"`
	// Slot ID of the local Switch slot Interface.
	SlotId *int64 `json:"SlotId,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId             *int64                               `json:"SwitchId,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortMacBindingAllOf PortMacBindingAllOf

// NewPortMacBindingAllOf instantiates a new PortMacBindingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortMacBindingAllOf(classId string, objectType string) *PortMacBindingAllOf {
	this := PortMacBindingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPortMacBindingAllOfWithDefaults instantiates a new PortMacBindingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortMacBindingAllOfWithDefaults() *PortMacBindingAllOf {
	this := PortMacBindingAllOf{}
	var classId string = "port.MacBinding"
	this.ClassId = classId
	var objectType string = "port.MacBinding"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PortMacBindingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PortMacBindingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PortMacBindingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PortMacBindingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetAggregatePortId() int64 {
	if o == nil || o.AggregatePortId == nil {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || o.AggregatePortId == nil {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasAggregatePortId() bool {
	if o != nil && o.AggregatePortId != nil {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *PortMacBindingAllOf) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *PortMacBindingAllOf) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetDeviceMac returns the DeviceMac field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetDeviceMac() string {
	if o == nil || o.DeviceMac == nil {
		var ret string
		return ret
	}
	return *o.DeviceMac
}

// GetDeviceMacOk returns a tuple with the DeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetDeviceMacOk() (*string, bool) {
	if o == nil || o.DeviceMac == nil {
		return nil, false
	}
	return o.DeviceMac, true
}

// HasDeviceMac returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasDeviceMac() bool {
	if o != nil && o.DeviceMac != nil {
		return true
	}

	return false
}

// SetDeviceMac gets a reference to the given string and assigns it to the DeviceMac field.
func (o *PortMacBindingAllOf) SetDeviceMac(v string) {
	o.DeviceMac = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *PortMacBindingAllOf) SetPortId(v int64) {
	o.PortId = &v
}

// GetPortMac returns the PortMac field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetPortMac() string {
	if o == nil || o.PortMac == nil {
		var ret string
		return ret
	}
	return *o.PortMac
}

// GetPortMacOk returns a tuple with the PortMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetPortMacOk() (*string, bool) {
	if o == nil || o.PortMac == nil {
		return nil, false
	}
	return o.PortMac, true
}

// HasPortMac returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasPortMac() bool {
	if o != nil && o.PortMac != nil {
		return true
	}

	return false
}

// SetPortMac gets a reference to the given string and assigns it to the PortMac field.
func (o *PortMacBindingAllOf) SetPortMac(v string) {
	o.PortMac = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *PortMacBindingAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetSwitchId() int64 {
	if o == nil || o.SwitchId == nil {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetSwitchIdOk() (*int64, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *PortMacBindingAllOf) SetSwitchId(v int64) {
	o.SwitchId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *PortMacBindingAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PortMacBindingAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBindingAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PortMacBindingAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PortMacBindingAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o PortMacBindingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AggregatePortId != nil {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.DeviceMac != nil {
		toSerialize["DeviceMac"] = o.DeviceMac
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.PortMac != nil {
		toSerialize["PortMac"] = o.PortMac
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PortMacBindingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPortMacBindingAllOf := _PortMacBindingAllOf{}

	if err = json.Unmarshal(bytes, &varPortMacBindingAllOf); err == nil {
		*o = PortMacBindingAllOf(varPortMacBindingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "DeviceMac")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "PortMac")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortMacBindingAllOf struct {
	value *PortMacBindingAllOf
	isSet bool
}

func (v NullablePortMacBindingAllOf) Get() *PortMacBindingAllOf {
	return v.value
}

func (v *NullablePortMacBindingAllOf) Set(val *PortMacBindingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePortMacBindingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePortMacBindingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortMacBindingAllOf(val *PortMacBindingAllOf) *NullablePortMacBindingAllOf {
	return &NullablePortMacBindingAllOf{value: val, isSet: true}
}

func (v NullablePortMacBindingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortMacBindingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
