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

// StorageNetAppEthernetPortAllOf Definition of the list of properties defined in 'storage.NetAppEthernetPort', excluding properties defined in parent classes.
type StorageNetAppEthernetPortAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the broadcast domain, scoped to its IPspace.
	BroadcastDomainName *string `json:"BroadcastDomainName,omitempty"`
	// Status of port to determine if its enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// MAC address of the port available in storage array.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Maximum transmission unit of the physical port available in storage array.
	Mtu *string `json:"Mtu,omitempty"`
	// Name of the port available in storage array.
	Name                   *string                               `json:"Name,omitempty"`
	NetAppEthernetPortLag  NullableStorageNetAppEthernetPortLag  `json:"NetAppEthernetPortLag,omitempty"`
	NetAppEthernetPortVlan NullableStorageNetAppEthernetPortVlan `json:"NetAppEthernetPortVlan,omitempty"`
	// Operational speed of port measured.
	Speed *int64 `json:"Speed,omitempty"`
	// State of the port available in storage array. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	State *string `json:"State,omitempty"`
	// Type of the port available in storage array. * `LAG` - Storage port of type lag. * `physical` - LIFs can be configured directly on physical ports. * `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.
	Type *string `json:"Type,omitempty"`
	// Universally unique identifier of the physical port.
	Uuid                 *string                        `json:"Uuid,omitempty"`
	ArrayController      *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppEthernetPortAllOf StorageNetAppEthernetPortAllOf

// NewStorageNetAppEthernetPortAllOf instantiates a new StorageNetAppEthernetPortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppEthernetPortAllOf(classId string, objectType string) *StorageNetAppEthernetPortAllOf {
	this := StorageNetAppEthernetPortAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppEthernetPortAllOfWithDefaults instantiates a new StorageNetAppEthernetPortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppEthernetPortAllOfWithDefaults() *StorageNetAppEthernetPortAllOf {
	this := StorageNetAppEthernetPortAllOf{}
	var classId string = "storage.NetAppEthernetPort"
	this.ClassId = classId
	var objectType string = "storage.NetAppEthernetPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppEthernetPortAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppEthernetPortAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppEthernetPortAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppEthernetPortAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBroadcastDomainName returns the BroadcastDomainName field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetBroadcastDomainName() string {
	if o == nil || o.BroadcastDomainName == nil {
		var ret string
		return ret
	}
	return *o.BroadcastDomainName
}

// GetBroadcastDomainNameOk returns a tuple with the BroadcastDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetBroadcastDomainNameOk() (*string, bool) {
	if o == nil || o.BroadcastDomainName == nil {
		return nil, false
	}
	return o.BroadcastDomainName, true
}

// HasBroadcastDomainName returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasBroadcastDomainName() bool {
	if o != nil && o.BroadcastDomainName != nil {
		return true
	}

	return false
}

// SetBroadcastDomainName gets a reference to the given string and assigns it to the BroadcastDomainName field.
func (o *StorageNetAppEthernetPortAllOf) SetBroadcastDomainName(v string) {
	o.BroadcastDomainName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppEthernetPortAllOf) SetEnabled(v string) {
	o.Enabled = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *StorageNetAppEthernetPortAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetMtu() string {
	if o == nil || o.Mtu == nil {
		var ret string
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetMtuOk() (*string, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given string and assigns it to the Mtu field.
func (o *StorageNetAppEthernetPortAllOf) SetMtu(v string) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppEthernetPortAllOf) SetName(v string) {
	o.Name = &v
}

// GetNetAppEthernetPortLag returns the NetAppEthernetPortLag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortLag() StorageNetAppEthernetPortLag {
	if o == nil || o.NetAppEthernetPortLag.Get() == nil {
		var ret StorageNetAppEthernetPortLag
		return ret
	}
	return *o.NetAppEthernetPortLag.Get()
}

// GetNetAppEthernetPortLagOk returns a tuple with the NetAppEthernetPortLag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortLagOk() (*StorageNetAppEthernetPortLag, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetAppEthernetPortLag.Get(), o.NetAppEthernetPortLag.IsSet()
}

// HasNetAppEthernetPortLag returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasNetAppEthernetPortLag() bool {
	if o != nil && o.NetAppEthernetPortLag.IsSet() {
		return true
	}

	return false
}

// SetNetAppEthernetPortLag gets a reference to the given NullableStorageNetAppEthernetPortLag and assigns it to the NetAppEthernetPortLag field.
func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortLag(v StorageNetAppEthernetPortLag) {
	o.NetAppEthernetPortLag.Set(&v)
}

// SetNetAppEthernetPortLagNil sets the value for NetAppEthernetPortLag to be an explicit nil
func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortLagNil() {
	o.NetAppEthernetPortLag.Set(nil)
}

// UnsetNetAppEthernetPortLag ensures that no value is present for NetAppEthernetPortLag, not even an explicit nil
func (o *StorageNetAppEthernetPortAllOf) UnsetNetAppEthernetPortLag() {
	o.NetAppEthernetPortLag.Unset()
}

// GetNetAppEthernetPortVlan returns the NetAppEthernetPortVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortVlan() StorageNetAppEthernetPortVlan {
	if o == nil || o.NetAppEthernetPortVlan.Get() == nil {
		var ret StorageNetAppEthernetPortVlan
		return ret
	}
	return *o.NetAppEthernetPortVlan.Get()
}

// GetNetAppEthernetPortVlanOk returns a tuple with the NetAppEthernetPortVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortVlanOk() (*StorageNetAppEthernetPortVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetAppEthernetPortVlan.Get(), o.NetAppEthernetPortVlan.IsSet()
}

// HasNetAppEthernetPortVlan returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasNetAppEthernetPortVlan() bool {
	if o != nil && o.NetAppEthernetPortVlan.IsSet() {
		return true
	}

	return false
}

// SetNetAppEthernetPortVlan gets a reference to the given NullableStorageNetAppEthernetPortVlan and assigns it to the NetAppEthernetPortVlan field.
func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortVlan(v StorageNetAppEthernetPortVlan) {
	o.NetAppEthernetPortVlan.Set(&v)
}

// SetNetAppEthernetPortVlanNil sets the value for NetAppEthernetPortVlan to be an explicit nil
func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortVlanNil() {
	o.NetAppEthernetPortVlan.Set(nil)
}

// UnsetNetAppEthernetPortVlan ensures that no value is present for NetAppEthernetPortVlan, not even an explicit nil
func (o *StorageNetAppEthernetPortAllOf) UnsetNetAppEthernetPortVlan() {
	o.NetAppEthernetPortVlan.Unset()
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *StorageNetAppEthernetPortAllOf) SetSpeed(v int64) {
	o.Speed = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppEthernetPortAllOf) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageNetAppEthernetPortAllOf) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppEthernetPortAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppEthernetPortAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

func (o StorageNetAppEthernetPortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BroadcastDomainName != nil {
		toSerialize["BroadcastDomainName"] = o.BroadcastDomainName
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NetAppEthernetPortLag.IsSet() {
		toSerialize["NetAppEthernetPortLag"] = o.NetAppEthernetPortLag.Get()
	}
	if o.NetAppEthernetPortVlan.IsSet() {
		toSerialize["NetAppEthernetPortVlan"] = o.NetAppEthernetPortVlan.Get()
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppEthernetPortAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppEthernetPortAllOf := _StorageNetAppEthernetPortAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppEthernetPortAllOf); err == nil {
		*o = StorageNetAppEthernetPortAllOf(varStorageNetAppEthernetPortAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BroadcastDomainName")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetAppEthernetPortLag")
		delete(additionalProperties, "NetAppEthernetPortVlan")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppEthernetPortAllOf struct {
	value *StorageNetAppEthernetPortAllOf
	isSet bool
}

func (v NullableStorageNetAppEthernetPortAllOf) Get() *StorageNetAppEthernetPortAllOf {
	return v.value
}

func (v *NullableStorageNetAppEthernetPortAllOf) Set(val *StorageNetAppEthernetPortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppEthernetPortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppEthernetPortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppEthernetPortAllOf(val *StorageNetAppEthernetPortAllOf) *NullableStorageNetAppEthernetPortAllOf {
	return &NullableStorageNetAppEthernetPortAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppEthernetPortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppEthernetPortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
