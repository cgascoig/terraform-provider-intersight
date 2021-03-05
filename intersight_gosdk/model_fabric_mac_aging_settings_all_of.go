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

// FabricMacAgingSettingsAllOf Definition of the list of properties defined in 'fabric.MacAgingSettings', excluding properties defined in parent classes.
type FabricMacAgingSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This specifies one of the option to configure the MAC address aging time. * `Default` - This option sets the default MAC address aging time to 14500 seconds for End Host mode. * `Custom` - This option allows the the user to configure the MAC address aging time on the switch. For Switch Model UCS-FI-6454 or higher, the valid range is 120 to 918000 seconds and the switch will set the lower multiple of 5 of the given time. * `Never` - This option disables the MAC address aging process and never allows the MAC address entries to get removed from the table.
	MacAgingOption *string `json:"MacAgingOption,omitempty"`
	// Define the MAC address aging time in seconds. This field is valid when the \"Custom\" MAC address aging option is selected.
	MacAgingTime         *int32 `json:"MacAgingTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricMacAgingSettingsAllOf FabricMacAgingSettingsAllOf

// NewFabricMacAgingSettingsAllOf instantiates a new FabricMacAgingSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricMacAgingSettingsAllOf(classId string, objectType string) *FabricMacAgingSettingsAllOf {
	this := FabricMacAgingSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var macAgingOption string = "Default"
	this.MacAgingOption = &macAgingOption
	return &this
}

// NewFabricMacAgingSettingsAllOfWithDefaults instantiates a new FabricMacAgingSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricMacAgingSettingsAllOfWithDefaults() *FabricMacAgingSettingsAllOf {
	this := FabricMacAgingSettingsAllOf{}
	var classId string = "fabric.MacAgingSettings"
	this.ClassId = classId
	var objectType string = "fabric.MacAgingSettings"
	this.ObjectType = objectType
	var macAgingOption string = "Default"
	this.MacAgingOption = &macAgingOption
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricMacAgingSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricMacAgingSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricMacAgingSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricMacAgingSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricMacAgingSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricMacAgingSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacAgingOption returns the MacAgingOption field value if set, zero value otherwise.
func (o *FabricMacAgingSettingsAllOf) GetMacAgingOption() string {
	if o == nil || o.MacAgingOption == nil {
		var ret string
		return ret
	}
	return *o.MacAgingOption
}

// GetMacAgingOptionOk returns a tuple with the MacAgingOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMacAgingSettingsAllOf) GetMacAgingOptionOk() (*string, bool) {
	if o == nil || o.MacAgingOption == nil {
		return nil, false
	}
	return o.MacAgingOption, true
}

// HasMacAgingOption returns a boolean if a field has been set.
func (o *FabricMacAgingSettingsAllOf) HasMacAgingOption() bool {
	if o != nil && o.MacAgingOption != nil {
		return true
	}

	return false
}

// SetMacAgingOption gets a reference to the given string and assigns it to the MacAgingOption field.
func (o *FabricMacAgingSettingsAllOf) SetMacAgingOption(v string) {
	o.MacAgingOption = &v
}

// GetMacAgingTime returns the MacAgingTime field value if set, zero value otherwise.
func (o *FabricMacAgingSettingsAllOf) GetMacAgingTime() int32 {
	if o == nil || o.MacAgingTime == nil {
		var ret int32
		return ret
	}
	return *o.MacAgingTime
}

// GetMacAgingTimeOk returns a tuple with the MacAgingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMacAgingSettingsAllOf) GetMacAgingTimeOk() (*int32, bool) {
	if o == nil || o.MacAgingTime == nil {
		return nil, false
	}
	return o.MacAgingTime, true
}

// HasMacAgingTime returns a boolean if a field has been set.
func (o *FabricMacAgingSettingsAllOf) HasMacAgingTime() bool {
	if o != nil && o.MacAgingTime != nil {
		return true
	}

	return false
}

// SetMacAgingTime gets a reference to the given int32 and assigns it to the MacAgingTime field.
func (o *FabricMacAgingSettingsAllOf) SetMacAgingTime(v int32) {
	o.MacAgingTime = &v
}

func (o FabricMacAgingSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MacAgingOption != nil {
		toSerialize["MacAgingOption"] = o.MacAgingOption
	}
	if o.MacAgingTime != nil {
		toSerialize["MacAgingTime"] = o.MacAgingTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricMacAgingSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricMacAgingSettingsAllOf := _FabricMacAgingSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varFabricMacAgingSettingsAllOf); err == nil {
		*o = FabricMacAgingSettingsAllOf(varFabricMacAgingSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacAgingOption")
		delete(additionalProperties, "MacAgingTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricMacAgingSettingsAllOf struct {
	value *FabricMacAgingSettingsAllOf
	isSet bool
}

func (v NullableFabricMacAgingSettingsAllOf) Get() *FabricMacAgingSettingsAllOf {
	return v.value
}

func (v *NullableFabricMacAgingSettingsAllOf) Set(val *FabricMacAgingSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricMacAgingSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricMacAgingSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricMacAgingSettingsAllOf(val *FabricMacAgingSettingsAllOf) *NullableFabricMacAgingSettingsAllOf {
	return &NullableFabricMacAgingSettingsAllOf{value: val, isSet: true}
}

func (v NullableFabricMacAgingSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricMacAgingSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
