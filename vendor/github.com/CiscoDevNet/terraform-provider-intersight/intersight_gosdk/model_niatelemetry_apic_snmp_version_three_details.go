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
	"reflect"
	"strings"
)

// NiatelemetryApicSnmpVersionThreeDetails Object to capture Snmp V3 details in APIC.
type NiatelemetryApicSnmpVersionThreeDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// AuthType of SNMP V3 in APIC.
	AuthType *string `json:"AuthType,omitempty"`
	// Dn of SNMP V3 attribute in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of SNMP V3 attribute in APIC.
	Name *string `json:"Name,omitempty"`
	// PrivType of SNMP V3 in APIC.
	PrivType *string `json:"PrivType,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicSnmpVersionThreeDetails NiatelemetryApicSnmpVersionThreeDetails

// NewNiatelemetryApicSnmpVersionThreeDetails instantiates a new NiatelemetryApicSnmpVersionThreeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSnmpVersionThreeDetails(classId string, objectType string) *NiatelemetryApicSnmpVersionThreeDetails {
	this := NiatelemetryApicSnmpVersionThreeDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSnmpVersionThreeDetailsWithDefaults instantiates a new NiatelemetryApicSnmpVersionThreeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSnmpVersionThreeDetailsWithDefaults() *NiatelemetryApicSnmpVersionThreeDetails {
	this := NiatelemetryApicSnmpVersionThreeDetails{}
	var classId string = "niatelemetry.ApicSnmpVersionThreeDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSnmpVersionThreeDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetAuthType(v string) {
	o.AuthType = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetDn(v string) {
	o.Dn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetName(v string) {
	o.Name = &v
}

// GetPrivType returns the PrivType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetPrivType() string {
	if o == nil || o.PrivType == nil {
		var ret string
		return ret
	}
	return *o.PrivType
}

// GetPrivTypeOk returns a tuple with the PrivType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetPrivTypeOk() (*string, bool) {
	if o == nil || o.PrivType == nil {
		return nil, false
	}
	return o.PrivType, true
}

// HasPrivType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasPrivType() bool {
	if o != nil && o.PrivType != nil {
		return true
	}

	return false
}

// SetPrivType gets a reference to the given string and assigns it to the PrivType field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetPrivType(v string) {
	o.PrivType = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicSnmpVersionThreeDetails) MarshalJSON() ([]byte, error) {
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
	if o.AuthType != nil {
		toSerialize["AuthType"] = o.AuthType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PrivType != nil {
		toSerialize["PrivType"] = o.PrivType
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

func (o *NiatelemetryApicSnmpVersionThreeDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// AuthType of SNMP V3 in APIC.
		AuthType *string `json:"AuthType,omitempty"`
		// Dn of SNMP V3 attribute in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Name of SNMP V3 attribute in APIC.
		Name *string `json:"Name,omitempty"`
		// PrivType of SNMP V3 in APIC.
		PrivType *string `json:"PrivType,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct := NiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicSnmpVersionThreeDetails := _NiatelemetryApicSnmpVersionThreeDetails{}
		varNiatelemetryApicSnmpVersionThreeDetails.ClassId = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicSnmpVersionThreeDetails.ObjectType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicSnmpVersionThreeDetails.AuthType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.AuthType
		varNiatelemetryApicSnmpVersionThreeDetails.Dn = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicSnmpVersionThreeDetails.Name = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.Name
		varNiatelemetryApicSnmpVersionThreeDetails.PrivType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.PrivType
		varNiatelemetryApicSnmpVersionThreeDetails.RecordType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicSnmpVersionThreeDetails.RecordVersion = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicSnmpVersionThreeDetails.SiteName = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicSnmpVersionThreeDetails.RegisteredDevice = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicSnmpVersionThreeDetails(varNiatelemetryApicSnmpVersionThreeDetails)
	} else {
		return err
	}

	varNiatelemetryApicSnmpVersionThreeDetails := _NiatelemetryApicSnmpVersionThreeDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSnmpVersionThreeDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicSnmpVersionThreeDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PrivType")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
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

type NullableNiatelemetryApicSnmpVersionThreeDetails struct {
	value *NiatelemetryApicSnmpVersionThreeDetails
	isSet bool
}

func (v NullableNiatelemetryApicSnmpVersionThreeDetails) Get() *NiatelemetryApicSnmpVersionThreeDetails {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpVersionThreeDetails) Set(val *NiatelemetryApicSnmpVersionThreeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpVersionThreeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpVersionThreeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpVersionThreeDetails(val *NiatelemetryApicSnmpVersionThreeDetails) *NullableNiatelemetryApicSnmpVersionThreeDetails {
	return &NullableNiatelemetryApicSnmpVersionThreeDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpVersionThreeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpVersionThreeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
