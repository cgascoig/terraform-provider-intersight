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
)

// CapabilitySwitchingModeCapabilityAllOf Definition of the list of properties defined in 'capability.SwitchingModeCapability', excluding properties defined in parent classes.
type CapabilitySwitchingModeCapabilityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Switching mode type (endhost, switch) of the switch. * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
	SwitchingMode *string `json:"SwitchingMode,omitempty"`
	// VP Compression support on this switch.
	VpCompressionSupported *bool `json:"VpCompressionSupported,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _CapabilitySwitchingModeCapabilityAllOf CapabilitySwitchingModeCapabilityAllOf

// NewCapabilitySwitchingModeCapabilityAllOf instantiates a new CapabilitySwitchingModeCapabilityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchingModeCapabilityAllOf(classId string, objectType string) *CapabilitySwitchingModeCapabilityAllOf {
	this := CapabilitySwitchingModeCapabilityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var switchingMode string = "end-host"
	this.SwitchingMode = &switchingMode
	return &this
}

// NewCapabilitySwitchingModeCapabilityAllOfWithDefaults instantiates a new CapabilitySwitchingModeCapabilityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchingModeCapabilityAllOfWithDefaults() *CapabilitySwitchingModeCapabilityAllOf {
	this := CapabilitySwitchingModeCapabilityAllOf{}
	var classId string = "capability.SwitchingModeCapability"
	this.ClassId = classId
	var objectType string = "capability.SwitchingModeCapability"
	this.ObjectType = objectType
	var switchingMode string = "end-host"
	this.SwitchingMode = &switchingMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchingModeCapabilityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapabilityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchingModeCapabilityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchingModeCapabilityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapabilityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchingModeCapabilityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSwitchingMode returns the SwitchingMode field value if set, zero value otherwise.
func (o *CapabilitySwitchingModeCapabilityAllOf) GetSwitchingMode() string {
	if o == nil || o.SwitchingMode == nil {
		var ret string
		return ret
	}
	return *o.SwitchingMode
}

// GetSwitchingModeOk returns a tuple with the SwitchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapabilityAllOf) GetSwitchingModeOk() (*string, bool) {
	if o == nil || o.SwitchingMode == nil {
		return nil, false
	}
	return o.SwitchingMode, true
}

// HasSwitchingMode returns a boolean if a field has been set.
func (o *CapabilitySwitchingModeCapabilityAllOf) HasSwitchingMode() bool {
	if o != nil && o.SwitchingMode != nil {
		return true
	}

	return false
}

// SetSwitchingMode gets a reference to the given string and assigns it to the SwitchingMode field.
func (o *CapabilitySwitchingModeCapabilityAllOf) SetSwitchingMode(v string) {
	o.SwitchingMode = &v
}

// GetVpCompressionSupported returns the VpCompressionSupported field value if set, zero value otherwise.
func (o *CapabilitySwitchingModeCapabilityAllOf) GetVpCompressionSupported() bool {
	if o == nil || o.VpCompressionSupported == nil {
		var ret bool
		return ret
	}
	return *o.VpCompressionSupported
}

// GetVpCompressionSupportedOk returns a tuple with the VpCompressionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapabilityAllOf) GetVpCompressionSupportedOk() (*bool, bool) {
	if o == nil || o.VpCompressionSupported == nil {
		return nil, false
	}
	return o.VpCompressionSupported, true
}

// HasVpCompressionSupported returns a boolean if a field has been set.
func (o *CapabilitySwitchingModeCapabilityAllOf) HasVpCompressionSupported() bool {
	if o != nil && o.VpCompressionSupported != nil {
		return true
	}

	return false
}

// SetVpCompressionSupported gets a reference to the given bool and assigns it to the VpCompressionSupported field.
func (o *CapabilitySwitchingModeCapabilityAllOf) SetVpCompressionSupported(v bool) {
	o.VpCompressionSupported = &v
}

func (o CapabilitySwitchingModeCapabilityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SwitchingMode != nil {
		toSerialize["SwitchingMode"] = o.SwitchingMode
	}
	if o.VpCompressionSupported != nil {
		toSerialize["VpCompressionSupported"] = o.VpCompressionSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchingModeCapabilityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitySwitchingModeCapabilityAllOf := _CapabilitySwitchingModeCapabilityAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilitySwitchingModeCapabilityAllOf); err == nil {
		*o = CapabilitySwitchingModeCapabilityAllOf(varCapabilitySwitchingModeCapabilityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SwitchingMode")
		delete(additionalProperties, "VpCompressionSupported")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySwitchingModeCapabilityAllOf struct {
	value *CapabilitySwitchingModeCapabilityAllOf
	isSet bool
}

func (v NullableCapabilitySwitchingModeCapabilityAllOf) Get() *CapabilitySwitchingModeCapabilityAllOf {
	return v.value
}

func (v *NullableCapabilitySwitchingModeCapabilityAllOf) Set(val *CapabilitySwitchingModeCapabilityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchingModeCapabilityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchingModeCapabilityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchingModeCapabilityAllOf(val *CapabilitySwitchingModeCapabilityAllOf) *NullableCapabilitySwitchingModeCapabilityAllOf {
	return &NullableCapabilitySwitchingModeCapabilityAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySwitchingModeCapabilityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchingModeCapabilityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
