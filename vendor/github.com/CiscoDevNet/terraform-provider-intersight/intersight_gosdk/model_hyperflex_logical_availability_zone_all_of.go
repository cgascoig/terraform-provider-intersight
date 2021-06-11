/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-05-25T18:18:54Z.
 *
 * API version: 1.0.9-4305
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexLogicalAvailabilityZoneAllOf Definition of the list of properties defined in 'hyperflex.LogicalAvailabilityZone', excluding properties defined in parent classes.
type HyperflexLogicalAvailabilityZoneAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or disable Logical Availability Zones (LAZ). If enabled, HyperFlex Data Platform automatically selects and groups nodes into different availability zones. For HyperFlex Data Platform versions prior to 3.0 release, this setting does not apply. For HyperFlex Data Platform versions 3.0 or higher, this setting is only applicable to Fabric Interconnect attached HyperFlex systems with 8 or more converged nodes.
	AutoConfig           *bool `json:"AutoConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexLogicalAvailabilityZoneAllOf HyperflexLogicalAvailabilityZoneAllOf

// NewHyperflexLogicalAvailabilityZoneAllOf instantiates a new HyperflexLogicalAvailabilityZoneAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexLogicalAvailabilityZoneAllOf(classId string, objectType string) *HyperflexLogicalAvailabilityZoneAllOf {
	this := HyperflexLogicalAvailabilityZoneAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var autoConfig bool = false
	this.AutoConfig = &autoConfig
	return &this
}

// NewHyperflexLogicalAvailabilityZoneAllOfWithDefaults instantiates a new HyperflexLogicalAvailabilityZoneAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexLogicalAvailabilityZoneAllOfWithDefaults() *HyperflexLogicalAvailabilityZoneAllOf {
	this := HyperflexLogicalAvailabilityZoneAllOf{}
	var classId string = "hyperflex.LogicalAvailabilityZone"
	this.ClassId = classId
	var objectType string = "hyperflex.LogicalAvailabilityZone"
	this.ObjectType = objectType
	var autoConfig bool = false
	this.AutoConfig = &autoConfig
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexLogicalAvailabilityZoneAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexLogicalAvailabilityZoneAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexLogicalAvailabilityZoneAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexLogicalAvailabilityZoneAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexLogicalAvailabilityZoneAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexLogicalAvailabilityZoneAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoConfig returns the AutoConfig field value if set, zero value otherwise.
func (o *HyperflexLogicalAvailabilityZoneAllOf) GetAutoConfig() bool {
	if o == nil || o.AutoConfig == nil {
		var ret bool
		return ret
	}
	return *o.AutoConfig
}

// GetAutoConfigOk returns a tuple with the AutoConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLogicalAvailabilityZoneAllOf) GetAutoConfigOk() (*bool, bool) {
	if o == nil || o.AutoConfig == nil {
		return nil, false
	}
	return o.AutoConfig, true
}

// HasAutoConfig returns a boolean if a field has been set.
func (o *HyperflexLogicalAvailabilityZoneAllOf) HasAutoConfig() bool {
	if o != nil && o.AutoConfig != nil {
		return true
	}

	return false
}

// SetAutoConfig gets a reference to the given bool and assigns it to the AutoConfig field.
func (o *HyperflexLogicalAvailabilityZoneAllOf) SetAutoConfig(v bool) {
	o.AutoConfig = &v
}

func (o HyperflexLogicalAvailabilityZoneAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutoConfig != nil {
		toSerialize["AutoConfig"] = o.AutoConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexLogicalAvailabilityZoneAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexLogicalAvailabilityZoneAllOf := _HyperflexLogicalAvailabilityZoneAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexLogicalAvailabilityZoneAllOf); err == nil {
		*o = HyperflexLogicalAvailabilityZoneAllOf(varHyperflexLogicalAvailabilityZoneAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexLogicalAvailabilityZoneAllOf struct {
	value *HyperflexLogicalAvailabilityZoneAllOf
	isSet bool
}

func (v NullableHyperflexLogicalAvailabilityZoneAllOf) Get() *HyperflexLogicalAvailabilityZoneAllOf {
	return v.value
}

func (v *NullableHyperflexLogicalAvailabilityZoneAllOf) Set(val *HyperflexLogicalAvailabilityZoneAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexLogicalAvailabilityZoneAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexLogicalAvailabilityZoneAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexLogicalAvailabilityZoneAllOf(val *HyperflexLogicalAvailabilityZoneAllOf) *NullableHyperflexLogicalAvailabilityZoneAllOf {
	return &NullableHyperflexLogicalAvailabilityZoneAllOf{value: val, isSet: true}
}

func (v NullableHyperflexLogicalAvailabilityZoneAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexLogicalAvailabilityZoneAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
