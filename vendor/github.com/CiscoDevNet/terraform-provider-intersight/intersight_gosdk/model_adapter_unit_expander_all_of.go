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

// AdapterUnitExpanderAllOf Definition of the list of properties defined in 'adapter.UnitExpander', excluding properties defined in parent classes.
type AdapterUnitExpanderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the partNumber of the given component.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the presence state of the given component.
	Presence *string `json:"Presence,omitempty"`
	// This field identifies the virtual id of the given component.
	Vid                  *string                              `json:"Vid,omitempty"`
	AdapterUnit          *AdapterUnitRelationship             `json:"AdapterUnit,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterUnitExpanderAllOf AdapterUnitExpanderAllOf

// NewAdapterUnitExpanderAllOf instantiates a new AdapterUnitExpanderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterUnitExpanderAllOf(classId string, objectType string) *AdapterUnitExpanderAllOf {
	this := AdapterUnitExpanderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterUnitExpanderAllOfWithDefaults instantiates a new AdapterUnitExpanderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterUnitExpanderAllOfWithDefaults() *AdapterUnitExpanderAllOf {
	this := AdapterUnitExpanderAllOf{}
	var classId string = "adapter.UnitExpander"
	this.ClassId = classId
	var objectType string = "adapter.UnitExpander"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterUnitExpanderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterUnitExpanderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterUnitExpanderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterUnitExpanderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterUnitExpanderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterUnitExpanderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *AdapterUnitExpanderAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitExpanderAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *AdapterUnitExpanderAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *AdapterUnitExpanderAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *AdapterUnitExpanderAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitExpanderAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *AdapterUnitExpanderAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *AdapterUnitExpanderAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *AdapterUnitExpanderAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitExpanderAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *AdapterUnitExpanderAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *AdapterUnitExpanderAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetAdapterUnit returns the AdapterUnit field value if set, zero value otherwise.
func (o *AdapterUnitExpanderAllOf) GetAdapterUnit() AdapterUnitRelationship {
	if o == nil || o.AdapterUnit == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.AdapterUnit
}

// GetAdapterUnitOk returns a tuple with the AdapterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitExpanderAllOf) GetAdapterUnitOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.AdapterUnit == nil {
		return nil, false
	}
	return o.AdapterUnit, true
}

// HasAdapterUnit returns a boolean if a field has been set.
func (o *AdapterUnitExpanderAllOf) HasAdapterUnit() bool {
	if o != nil && o.AdapterUnit != nil {
		return true
	}

	return false
}

// SetAdapterUnit gets a reference to the given AdapterUnitRelationship and assigns it to the AdapterUnit field.
func (o *AdapterUnitExpanderAllOf) SetAdapterUnit(v AdapterUnitRelationship) {
	o.AdapterUnit = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AdapterUnitExpanderAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitExpanderAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AdapterUnitExpanderAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AdapterUnitExpanderAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AdapterUnitExpanderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.AdapterUnit != nil {
		toSerialize["AdapterUnit"] = o.AdapterUnit
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterUnitExpanderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAdapterUnitExpanderAllOf := _AdapterUnitExpanderAllOf{}

	if err = json.Unmarshal(bytes, &varAdapterUnitExpanderAllOf); err == nil {
		*o = AdapterUnitExpanderAllOf(varAdapterUnitExpanderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "AdapterUnit")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdapterUnitExpanderAllOf struct {
	value *AdapterUnitExpanderAllOf
	isSet bool
}

func (v NullableAdapterUnitExpanderAllOf) Get() *AdapterUnitExpanderAllOf {
	return v.value
}

func (v *NullableAdapterUnitExpanderAllOf) Set(val *AdapterUnitExpanderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterUnitExpanderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterUnitExpanderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterUnitExpanderAllOf(val *AdapterUnitExpanderAllOf) *NullableAdapterUnitExpanderAllOf {
	return &NullableAdapterUnitExpanderAllOf{value: val, isSet: true}
}

func (v NullableAdapterUnitExpanderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterUnitExpanderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
