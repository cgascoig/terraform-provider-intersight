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

// NiatelemetryApicNtpAuthAllOf Definition of the list of properties defined in 'niatelemetry.ApicNtpAuth', excluding properties defined in parent classes.
type NiatelemetryApicNtpAuthAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Authentication state of NTP policy in APIC.
	AuthSt *string `json:"AuthSt,omitempty"`
	// Dn of the NTP Auth in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicNtpAuthAllOf NiatelemetryApicNtpAuthAllOf

// NewNiatelemetryApicNtpAuthAllOf instantiates a new NiatelemetryApicNtpAuthAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicNtpAuthAllOf(classId string, objectType string) *NiatelemetryApicNtpAuthAllOf {
	this := NiatelemetryApicNtpAuthAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicNtpAuthAllOfWithDefaults instantiates a new NiatelemetryApicNtpAuthAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicNtpAuthAllOfWithDefaults() *NiatelemetryApicNtpAuthAllOf {
	this := NiatelemetryApicNtpAuthAllOf{}
	var classId string = "niatelemetry.ApicNtpAuth"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicNtpAuth"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicNtpAuthAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicNtpAuthAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicNtpAuthAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicNtpAuthAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthSt returns the AuthSt field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuthAllOf) GetAuthSt() string {
	if o == nil || o.AuthSt == nil {
		var ret string
		return ret
	}
	return *o.AuthSt
}

// GetAuthStOk returns a tuple with the AuthSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetAuthStOk() (*string, bool) {
	if o == nil || o.AuthSt == nil {
		return nil, false
	}
	return o.AuthSt, true
}

// HasAuthSt returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuthAllOf) HasAuthSt() bool {
	if o != nil && o.AuthSt != nil {
		return true
	}

	return false
}

// SetAuthSt gets a reference to the given string and assigns it to the AuthSt field.
func (o *NiatelemetryApicNtpAuthAllOf) SetAuthSt(v string) {
	o.AuthSt = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuthAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuthAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicNtpAuthAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuthAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuthAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicNtpAuthAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuthAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuthAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicNtpAuthAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuthAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuthAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicNtpAuthAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuthAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuthAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuthAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicNtpAuthAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicNtpAuthAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthSt != nil {
		toSerialize["AuthSt"] = o.AuthSt
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicNtpAuthAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryApicNtpAuthAllOf := _NiatelemetryApicNtpAuthAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryApicNtpAuthAllOf); err == nil {
		*o = NiatelemetryApicNtpAuthAllOf(varNiatelemetryApicNtpAuthAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthSt")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicNtpAuthAllOf struct {
	value *NiatelemetryApicNtpAuthAllOf
	isSet bool
}

func (v NullableNiatelemetryApicNtpAuthAllOf) Get() *NiatelemetryApicNtpAuthAllOf {
	return v.value
}

func (v *NullableNiatelemetryApicNtpAuthAllOf) Set(val *NiatelemetryApicNtpAuthAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicNtpAuthAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicNtpAuthAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicNtpAuthAllOf(val *NiatelemetryApicNtpAuthAllOf) *NullableNiatelemetryApicNtpAuthAllOf {
	return &NullableNiatelemetryApicNtpAuthAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryApicNtpAuthAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicNtpAuthAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
