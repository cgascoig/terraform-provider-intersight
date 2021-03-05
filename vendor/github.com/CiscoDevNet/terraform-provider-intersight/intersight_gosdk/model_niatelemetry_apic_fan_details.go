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
	"reflect"
	"strings"
)

// NiatelemetryApicFanDetails Object to capture the fan details in APIC.
type NiatelemetryApicFanDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn for the fan in the inventory.
	Dn *string `json:"Dn,omitempty"`
	// Model number of the fan in APIC.
	ModelNumber *string `json:"ModelNumber,omitempty"`
	// Node id for the fan in the inventory.
	NodeId *string `json:"NodeId,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the fan in APIC.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Vendor name of the fan in APIC.
	VendorName           *string                              `json:"VendorName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicFanDetails NiatelemetryApicFanDetails

// NewNiatelemetryApicFanDetails instantiates a new NiatelemetryApicFanDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicFanDetails(classId string, objectType string) *NiatelemetryApicFanDetails {
	this := NiatelemetryApicFanDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicFanDetailsWithDefaults instantiates a new NiatelemetryApicFanDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicFanDetailsWithDefaults() *NiatelemetryApicFanDetails {
	this := NiatelemetryApicFanDetails{}
	var classId string = "niatelemetry.ApicFanDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicFanDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicFanDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicFanDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicFanDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicFanDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicFanDetails) SetDn(v string) {
	o.Dn = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *NiatelemetryApicFanDetails) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *NiatelemetryApicFanDetails) SetNodeId(v string) {
	o.NodeId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicFanDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicFanDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryApicFanDetails) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicFanDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetVendorName() string {
	if o == nil || o.VendorName == nil {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetVendorNameOk() (*string, bool) {
	if o == nil || o.VendorName == nil {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasVendorName() bool {
	if o != nil && o.VendorName != nil {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *NiatelemetryApicFanDetails) SetVendorName(v string) {
	o.VendorName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicFanDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFanDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicFanDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicFanDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicFanDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.ModelNumber != nil {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.VendorName != nil {
		toSerialize["VendorName"] = o.VendorName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicFanDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicFanDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn for the fan in the inventory.
		Dn *string `json:"Dn,omitempty"`
		// Model number of the fan in APIC.
		ModelNumber *string `json:"ModelNumber,omitempty"`
		// Node id for the fan in the inventory.
		NodeId *string `json:"NodeId,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Serial number of the fan in APIC.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// Vendor name of the fan in APIC.
		VendorName       *string                              `json:"VendorName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicFanDetailsWithoutEmbeddedStruct := NiatelemetryApicFanDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicFanDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicFanDetails := _NiatelemetryApicFanDetails{}
		varNiatelemetryApicFanDetails.ClassId = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicFanDetails.ObjectType = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicFanDetails.Dn = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicFanDetails.ModelNumber = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.ModelNumber
		varNiatelemetryApicFanDetails.NodeId = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.NodeId
		varNiatelemetryApicFanDetails.RecordType = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicFanDetails.RecordVersion = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicFanDetails.SerialNumber = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.SerialNumber
		varNiatelemetryApicFanDetails.SiteName = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicFanDetails.VendorName = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.VendorName
		varNiatelemetryApicFanDetails.RegisteredDevice = varNiatelemetryApicFanDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicFanDetails(varNiatelemetryApicFanDetails)
	} else {
		return err
	}

	varNiatelemetryApicFanDetails := _NiatelemetryApicFanDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicFanDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicFanDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "ModelNumber")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "VendorName")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicFanDetails struct {
	value *NiatelemetryApicFanDetails
	isSet bool
}

func (v NullableNiatelemetryApicFanDetails) Get() *NiatelemetryApicFanDetails {
	return v.value
}

func (v *NullableNiatelemetryApicFanDetails) Set(val *NiatelemetryApicFanDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicFanDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicFanDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicFanDetails(val *NiatelemetryApicFanDetails) *NullableNiatelemetryApicFanDetails {
	return &NullableNiatelemetryApicFanDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicFanDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicFanDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
