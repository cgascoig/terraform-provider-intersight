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

// NiatelemetryHttpsAclFilterDetailsAllOf Definition of the list of properties defined in 'niatelemetry.HttpsAclFilterDetails', excluding properties defined in parent classes.
type NiatelemetryHttpsAclFilterDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Destination From Port HTTPS ACL EPGs filter MO for APIC.
	DestFromPort *string `json:"DestFromPort,omitempty"`
	// Destination To Port HTTPS ACL EPGs filter MO for APIC.
	DestToPort *string `json:"DestToPort,omitempty"`
	// Dn of the HTTPS ACL EPGs filter MO for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of HTTPS ACL filter for APIC.
	FilterName *string `json:"FilterName,omitempty"`
	// Prot of the HTTPS ACL EPGs filter MO for APIC.
	Prot *string `json:"Prot,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Source From Port HTTPS ACL EPGs filter MO for APIC.
	SrcFromPort *string `json:"SrcFromPort,omitempty"`
	// Source To Port HTTPS ACL EPGs filter MO for APIC.
	SrcToPort            *string                              `json:"SrcToPort,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclFilterDetailsAllOf NiatelemetryHttpsAclFilterDetailsAllOf

// NewNiatelemetryHttpsAclFilterDetailsAllOf instantiates a new NiatelemetryHttpsAclFilterDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclFilterDetailsAllOf(classId string, objectType string) *NiatelemetryHttpsAclFilterDetailsAllOf {
	this := NiatelemetryHttpsAclFilterDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclFilterDetailsAllOfWithDefaults instantiates a new NiatelemetryHttpsAclFilterDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclFilterDetailsAllOfWithDefaults() *NiatelemetryHttpsAclFilterDetailsAllOf {
	this := NiatelemetryHttpsAclFilterDetailsAllOf{}
	var classId string = "niatelemetry.HttpsAclFilterDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclFilterDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDestFromPort returns the DestFromPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestFromPort() string {
	if o == nil || o.DestFromPort == nil {
		var ret string
		return ret
	}
	return *o.DestFromPort
}

// GetDestFromPortOk returns a tuple with the DestFromPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestFromPortOk() (*string, bool) {
	if o == nil || o.DestFromPort == nil {
		return nil, false
	}
	return o.DestFromPort, true
}

// HasDestFromPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasDestFromPort() bool {
	if o != nil && o.DestFromPort != nil {
		return true
	}

	return false
}

// SetDestFromPort gets a reference to the given string and assigns it to the DestFromPort field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetDestFromPort(v string) {
	o.DestFromPort = &v
}

// GetDestToPort returns the DestToPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestToPort() string {
	if o == nil || o.DestToPort == nil {
		var ret string
		return ret
	}
	return *o.DestToPort
}

// GetDestToPortOk returns a tuple with the DestToPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestToPortOk() (*string, bool) {
	if o == nil || o.DestToPort == nil {
		return nil, false
	}
	return o.DestToPort, true
}

// HasDestToPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasDestToPort() bool {
	if o != nil && o.DestToPort != nil {
		return true
	}

	return false
}

// SetDestToPort gets a reference to the given string and assigns it to the DestToPort field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetDestToPort(v string) {
	o.DestToPort = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetFilterName returns the FilterName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetFilterName() string {
	if o == nil || o.FilterName == nil {
		var ret string
		return ret
	}
	return *o.FilterName
}

// GetFilterNameOk returns a tuple with the FilterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetFilterNameOk() (*string, bool) {
	if o == nil || o.FilterName == nil {
		return nil, false
	}
	return o.FilterName, true
}

// HasFilterName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasFilterName() bool {
	if o != nil && o.FilterName != nil {
		return true
	}

	return false
}

// SetFilterName gets a reference to the given string and assigns it to the FilterName field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetFilterName(v string) {
	o.FilterName = &v
}

// GetProt returns the Prot field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetProt() string {
	if o == nil || o.Prot == nil {
		var ret string
		return ret
	}
	return *o.Prot
}

// GetProtOk returns a tuple with the Prot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetProtOk() (*string, bool) {
	if o == nil || o.Prot == nil {
		return nil, false
	}
	return o.Prot, true
}

// HasProt returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasProt() bool {
	if o != nil && o.Prot != nil {
		return true
	}

	return false
}

// SetProt gets a reference to the given string and assigns it to the Prot field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetProt(v string) {
	o.Prot = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSrcFromPort returns the SrcFromPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcFromPort() string {
	if o == nil || o.SrcFromPort == nil {
		var ret string
		return ret
	}
	return *o.SrcFromPort
}

// GetSrcFromPortOk returns a tuple with the SrcFromPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcFromPortOk() (*string, bool) {
	if o == nil || o.SrcFromPort == nil {
		return nil, false
	}
	return o.SrcFromPort, true
}

// HasSrcFromPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasSrcFromPort() bool {
	if o != nil && o.SrcFromPort != nil {
		return true
	}

	return false
}

// SetSrcFromPort gets a reference to the given string and assigns it to the SrcFromPort field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetSrcFromPort(v string) {
	o.SrcFromPort = &v
}

// GetSrcToPort returns the SrcToPort field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcToPort() string {
	if o == nil || o.SrcToPort == nil {
		var ret string
		return ret
	}
	return *o.SrcToPort
}

// GetSrcToPortOk returns a tuple with the SrcToPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcToPortOk() (*string, bool) {
	if o == nil || o.SrcToPort == nil {
		return nil, false
	}
	return o.SrcToPort, true
}

// HasSrcToPort returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasSrcToPort() bool {
	if o != nil && o.SrcToPort != nil {
		return true
	}

	return false
}

// SetSrcToPort gets a reference to the given string and assigns it to the SrcToPort field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetSrcToPort(v string) {
	o.SrcToPort = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHttpsAclFilterDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DestFromPort != nil {
		toSerialize["DestFromPort"] = o.DestFromPort
	}
	if o.DestToPort != nil {
		toSerialize["DestToPort"] = o.DestToPort
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.FilterName != nil {
		toSerialize["FilterName"] = o.FilterName
	}
	if o.Prot != nil {
		toSerialize["Prot"] = o.Prot
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
	if o.SrcFromPort != nil {
		toSerialize["SrcFromPort"] = o.SrcFromPort
	}
	if o.SrcToPort != nil {
		toSerialize["SrcToPort"] = o.SrcToPort
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryHttpsAclFilterDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryHttpsAclFilterDetailsAllOf := _NiatelemetryHttpsAclFilterDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclFilterDetailsAllOf); err == nil {
		*o = NiatelemetryHttpsAclFilterDetailsAllOf(varNiatelemetryHttpsAclFilterDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DestFromPort")
		delete(additionalProperties, "DestToPort")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FilterName")
		delete(additionalProperties, "Prot")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SrcFromPort")
		delete(additionalProperties, "SrcToPort")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryHttpsAclFilterDetailsAllOf struct {
	value *NiatelemetryHttpsAclFilterDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryHttpsAclFilterDetailsAllOf) Get() *NiatelemetryHttpsAclFilterDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclFilterDetailsAllOf) Set(val *NiatelemetryHttpsAclFilterDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclFilterDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclFilterDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclFilterDetailsAllOf(val *NiatelemetryHttpsAclFilterDetailsAllOf) *NullableNiatelemetryHttpsAclFilterDetailsAllOf {
	return &NullableNiatelemetryHttpsAclFilterDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclFilterDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclFilterDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
