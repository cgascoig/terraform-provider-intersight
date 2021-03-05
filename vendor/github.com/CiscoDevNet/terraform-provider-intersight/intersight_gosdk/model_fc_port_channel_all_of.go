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

// FcPortChannelAllOf Definition of the list of properties defined in 'fc.PortChannel', excluding properties defined in parent classes.
type FcPortChannelAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrator configured Speed applied on the port channel.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Administratively configured state (enabled/disabled) for this portchannel.
	AdminState *string `json:"AdminState,omitempty"`
	// Mode information N_proxy, F or E associated to the Fibre Channel portchannel.
	Mode *string `json:"Mode,omitempty"`
	// Operational speed of this port-channel.
	OperSpeed *string `json:"OperSpeed,omitempty"`
	// Operational state of this port-channel.
	OperState *string `json:"OperState,omitempty"`
	// Reason for this port-channel's Operational state.
	OperStateQual *string `json:"OperStateQual,omitempty"`
	// Unique identifier for this port-channel on the FI.
	PortChannelId *int64 `json:"PortChannelId,omitempty"`
	// This port-channel's configured role (fcUplink, fcStorage, etc.).
	Role *string `json:"Role,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId *string `json:"SwitchId,omitempty"`
	// Virtual San that is associated to the port-channel.
	Vsan                 *int64                               `json:"Vsan,omitempty"`
	EquipmentSwitchCard  *EquipmentSwitchCardRelationship     `json:"EquipmentSwitchCard,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcPortChannelAllOf FcPortChannelAllOf

// NewFcPortChannelAllOf instantiates a new FcPortChannelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcPortChannelAllOf(classId string, objectType string) *FcPortChannelAllOf {
	this := FcPortChannelAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcPortChannelAllOfWithDefaults instantiates a new FcPortChannelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcPortChannelAllOfWithDefaults() *FcPortChannelAllOf {
	this := FcPortChannelAllOf{}
	var classId string = "fc.PortChannel"
	this.ClassId = classId
	var objectType string = "fc.PortChannel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcPortChannelAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcPortChannelAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcPortChannelAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcPortChannelAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FcPortChannelAllOf) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FcPortChannelAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *FcPortChannelAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetOperSpeed returns the OperSpeed field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetOperSpeed() string {
	if o == nil || o.OperSpeed == nil {
		var ret string
		return ret
	}
	return *o.OperSpeed
}

// GetOperSpeedOk returns a tuple with the OperSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetOperSpeedOk() (*string, bool) {
	if o == nil || o.OperSpeed == nil {
		return nil, false
	}
	return o.OperSpeed, true
}

// HasOperSpeed returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasOperSpeed() bool {
	if o != nil && o.OperSpeed != nil {
		return true
	}

	return false
}

// SetOperSpeed gets a reference to the given string and assigns it to the OperSpeed field.
func (o *FcPortChannelAllOf) SetOperSpeed(v string) {
	o.OperSpeed = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *FcPortChannelAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOperStateQual returns the OperStateQual field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetOperStateQual() string {
	if o == nil || o.OperStateQual == nil {
		var ret string
		return ret
	}
	return *o.OperStateQual
}

// GetOperStateQualOk returns a tuple with the OperStateQual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetOperStateQualOk() (*string, bool) {
	if o == nil || o.OperStateQual == nil {
		return nil, false
	}
	return o.OperStateQual, true
}

// HasOperStateQual returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasOperStateQual() bool {
	if o != nil && o.OperStateQual != nil {
		return true
	}

	return false
}

// SetOperStateQual gets a reference to the given string and assigns it to the OperStateQual field.
func (o *FcPortChannelAllOf) SetOperStateQual(v string) {
	o.OperStateQual = &v
}

// GetPortChannelId returns the PortChannelId field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetPortChannelId() int64 {
	if o == nil || o.PortChannelId == nil {
		var ret int64
		return ret
	}
	return *o.PortChannelId
}

// GetPortChannelIdOk returns a tuple with the PortChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetPortChannelIdOk() (*int64, bool) {
	if o == nil || o.PortChannelId == nil {
		return nil, false
	}
	return o.PortChannelId, true
}

// HasPortChannelId returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasPortChannelId() bool {
	if o != nil && o.PortChannelId != nil {
		return true
	}

	return false
}

// SetPortChannelId gets a reference to the given int64 and assigns it to the PortChannelId field.
func (o *FcPortChannelAllOf) SetPortChannelId(v int64) {
	o.PortChannelId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *FcPortChannelAllOf) SetRole(v string) {
	o.Role = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *FcPortChannelAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetVsan returns the Vsan field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetVsan() int64 {
	if o == nil || o.Vsan == nil {
		var ret int64
		return ret
	}
	return *o.Vsan
}

// GetVsanOk returns a tuple with the Vsan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetVsanOk() (*int64, bool) {
	if o == nil || o.Vsan == nil {
		return nil, false
	}
	return o.Vsan, true
}

// HasVsan returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasVsan() bool {
	if o != nil && o.Vsan != nil {
		return true
	}

	return false
}

// SetVsan gets a reference to the given int64 and assigns it to the Vsan field.
func (o *FcPortChannelAllOf) SetVsan(v int64) {
	o.Vsan = &v
}

// GetEquipmentSwitchCard returns the EquipmentSwitchCard field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship {
	if o == nil || o.EquipmentSwitchCard == nil {
		var ret EquipmentSwitchCardRelationship
		return ret
	}
	return *o.EquipmentSwitchCard
}

// GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool) {
	if o == nil || o.EquipmentSwitchCard == nil {
		return nil, false
	}
	return o.EquipmentSwitchCard, true
}

// HasEquipmentSwitchCard returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasEquipmentSwitchCard() bool {
	if o != nil && o.EquipmentSwitchCard != nil {
		return true
	}

	return false
}

// SetEquipmentSwitchCard gets a reference to the given EquipmentSwitchCardRelationship and assigns it to the EquipmentSwitchCard field.
func (o *FcPortChannelAllOf) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship) {
	o.EquipmentSwitchCard = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *FcPortChannelAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPortChannelAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *FcPortChannelAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *FcPortChannelAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o FcPortChannelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.OperSpeed != nil {
		toSerialize["OperSpeed"] = o.OperSpeed
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.OperStateQual != nil {
		toSerialize["OperStateQual"] = o.OperStateQual
	}
	if o.PortChannelId != nil {
		toSerialize["PortChannelId"] = o.PortChannelId
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Vsan != nil {
		toSerialize["Vsan"] = o.Vsan
	}
	if o.EquipmentSwitchCard != nil {
		toSerialize["EquipmentSwitchCard"] = o.EquipmentSwitchCard
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcPortChannelAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcPortChannelAllOf := _FcPortChannelAllOf{}

	if err = json.Unmarshal(bytes, &varFcPortChannelAllOf); err == nil {
		*o = FcPortChannelAllOf(varFcPortChannelAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "OperSpeed")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "OperStateQual")
		delete(additionalProperties, "PortChannelId")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Vsan")
		delete(additionalProperties, "EquipmentSwitchCard")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcPortChannelAllOf struct {
	value *FcPortChannelAllOf
	isSet bool
}

func (v NullableFcPortChannelAllOf) Get() *FcPortChannelAllOf {
	return v.value
}

func (v *NullableFcPortChannelAllOf) Set(val *FcPortChannelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcPortChannelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcPortChannelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcPortChannelAllOf(val *FcPortChannelAllOf) *NullableFcPortChannelAllOf {
	return &NullableFcPortChannelAllOf{value: val, isSet: true}
}

func (v NullableFcPortChannelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcPortChannelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
