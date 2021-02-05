/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// FirmwareServerUpgradeImpactAllOf Definition of the list of properties defined in 'firmware.ServerUpgradeImpact', excluding properties defined in parent classes.
type FirmwareServerUpgradeImpactAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                    `json:"ObjectType"`
	ImpactDetail []FirmwareComponentImpact `json:"ImpactDetail,omitempty"`
	// Name of the server impacted by the upgrade.
	Name *string `json:"Name,omitempty"`
	// Details about the server which will be impacted by the upgrade.
	UserLabel            *string `json:"UserLabel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareServerUpgradeImpactAllOf FirmwareServerUpgradeImpactAllOf

// NewFirmwareServerUpgradeImpactAllOf instantiates a new FirmwareServerUpgradeImpactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareServerUpgradeImpactAllOf(classId string, objectType string) *FirmwareServerUpgradeImpactAllOf {
	this := FirmwareServerUpgradeImpactAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareServerUpgradeImpactAllOfWithDefaults instantiates a new FirmwareServerUpgradeImpactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareServerUpgradeImpactAllOfWithDefaults() *FirmwareServerUpgradeImpactAllOf {
	this := FirmwareServerUpgradeImpactAllOf{}
	var classId string = "firmware.ServerUpgradeImpact"
	this.ClassId = classId
	var objectType string = "firmware.ServerUpgradeImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareServerUpgradeImpactAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareServerUpgradeImpactAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareServerUpgradeImpactAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareServerUpgradeImpactAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareServerUpgradeImpactAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareServerUpgradeImpactAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetImpactDetail returns the ImpactDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareServerUpgradeImpactAllOf) GetImpactDetail() []FirmwareComponentImpact {
	if o == nil {
		var ret []FirmwareComponentImpact
		return ret
	}
	return o.ImpactDetail
}

// GetImpactDetailOk returns a tuple with the ImpactDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareServerUpgradeImpactAllOf) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool) {
	if o == nil || o.ImpactDetail == nil {
		return nil, false
	}
	return &o.ImpactDetail, true
}

// HasImpactDetail returns a boolean if a field has been set.
func (o *FirmwareServerUpgradeImpactAllOf) HasImpactDetail() bool {
	if o != nil && o.ImpactDetail != nil {
		return true
	}

	return false
}

// SetImpactDetail gets a reference to the given []FirmwareComponentImpact and assigns it to the ImpactDetail field.
func (o *FirmwareServerUpgradeImpactAllOf) SetImpactDetail(v []FirmwareComponentImpact) {
	o.ImpactDetail = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirmwareServerUpgradeImpactAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareServerUpgradeImpactAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirmwareServerUpgradeImpactAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirmwareServerUpgradeImpactAllOf) SetName(v string) {
	o.Name = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *FirmwareServerUpgradeImpactAllOf) GetUserLabel() string {
	if o == nil || o.UserLabel == nil {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareServerUpgradeImpactAllOf) GetUserLabelOk() (*string, bool) {
	if o == nil || o.UserLabel == nil {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *FirmwareServerUpgradeImpactAllOf) HasUserLabel() bool {
	if o != nil && o.UserLabel != nil {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *FirmwareServerUpgradeImpactAllOf) SetUserLabel(v string) {
	o.UserLabel = &v
}

func (o FirmwareServerUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ImpactDetail != nil {
		toSerialize["ImpactDetail"] = o.ImpactDetail
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.UserLabel != nil {
		toSerialize["UserLabel"] = o.UserLabel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareServerUpgradeImpactAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareServerUpgradeImpactAllOf := _FirmwareServerUpgradeImpactAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareServerUpgradeImpactAllOf); err == nil {
		*o = FirmwareServerUpgradeImpactAllOf(varFirmwareServerUpgradeImpactAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ImpactDetail")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "UserLabel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareServerUpgradeImpactAllOf struct {
	value *FirmwareServerUpgradeImpactAllOf
	isSet bool
}

func (v NullableFirmwareServerUpgradeImpactAllOf) Get() *FirmwareServerUpgradeImpactAllOf {
	return v.value
}

func (v *NullableFirmwareServerUpgradeImpactAllOf) Set(val *FirmwareServerUpgradeImpactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareServerUpgradeImpactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareServerUpgradeImpactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareServerUpgradeImpactAllOf(val *FirmwareServerUpgradeImpactAllOf) *NullableFirmwareServerUpgradeImpactAllOf {
	return &NullableFirmwareServerUpgradeImpactAllOf{value: val, isSet: true}
}

func (v NullableFirmwareServerUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareServerUpgradeImpactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
