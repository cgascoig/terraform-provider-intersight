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

// AssetSudiInfoAllOf Definition of the list of properties defined in 'asset.SudiInfo', excluding properties defined in parent classes.
type AssetSudiInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The device model (PID) extracted from the X.509 SUDI Leaf Certificate.
	Pid *string `json:"Pid,omitempty"`
	// The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate.
	Signature *string `json:"Signature,omitempty"`
	// The validation status of the device. * `DeviceStatusUnknown` - SUDI validation is done on the establishment of a connection. Before a device connects or after it disconnects, the SUDI validation status is set to this value. * `Verified` - The device returned a valid PID, Serial Number, Status and X.509 Leaf Certificate. The certificate signing chain was validated. * `CertificateValidationFailed` - Validation of the certificate signing chain failed. * `UnsupportedFirmware` - The firmware version of the Cisco IMC that is installed does not contain the SUDI APIs needed to perform validation. * `UnsupportedHardware` - The device is a model that does not contain a Trust Anchor Module (TAM) and thus cannot be validated. * `DeviceNotResponding` - An request was sent to the device, but no response was received.
	Status               *string                 `json:"Status,omitempty"`
	SudiCertificate      NullableX509Certificate `json:"SudiCertificate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetSudiInfoAllOf AssetSudiInfoAllOf

// NewAssetSudiInfoAllOf instantiates a new AssetSudiInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSudiInfoAllOf(classId string, objectType string) *AssetSudiInfoAllOf {
	this := AssetSudiInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "DeviceStatusUnknown"
	this.Status = &status
	return &this
}

// NewAssetSudiInfoAllOfWithDefaults instantiates a new AssetSudiInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSudiInfoAllOfWithDefaults() *AssetSudiInfoAllOf {
	this := AssetSudiInfoAllOf{}
	var classId string = "asset.SudiInfo"
	this.ClassId = classId
	var objectType string = "asset.SudiInfo"
	this.ObjectType = objectType
	var status string = "DeviceStatusUnknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSudiInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSudiInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSudiInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetSudiInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSudiInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSudiInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AssetSudiInfoAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfoAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AssetSudiInfoAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AssetSudiInfoAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *AssetSudiInfoAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfoAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *AssetSudiInfoAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *AssetSudiInfoAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *AssetSudiInfoAllOf) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfoAllOf) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *AssetSudiInfoAllOf) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *AssetSudiInfoAllOf) SetSignature(v string) {
	o.Signature = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetSudiInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetSudiInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetSudiInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetSudiCertificate returns the SudiCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSudiInfoAllOf) GetSudiCertificate() X509Certificate {
	if o == nil || o.SudiCertificate.Get() == nil {
		var ret X509Certificate
		return ret
	}
	return *o.SudiCertificate.Get()
}

// GetSudiCertificateOk returns a tuple with the SudiCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSudiInfoAllOf) GetSudiCertificateOk() (*X509Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.SudiCertificate.Get(), o.SudiCertificate.IsSet()
}

// HasSudiCertificate returns a boolean if a field has been set.
func (o *AssetSudiInfoAllOf) HasSudiCertificate() bool {
	if o != nil && o.SudiCertificate.IsSet() {
		return true
	}

	return false
}

// SetSudiCertificate gets a reference to the given NullableX509Certificate and assigns it to the SudiCertificate field.
func (o *AssetSudiInfoAllOf) SetSudiCertificate(v X509Certificate) {
	o.SudiCertificate.Set(&v)
}

// SetSudiCertificateNil sets the value for SudiCertificate to be an explicit nil
func (o *AssetSudiInfoAllOf) SetSudiCertificateNil() {
	o.SudiCertificate.Set(nil)
}

// UnsetSudiCertificate ensures that no value is present for SudiCertificate, not even an explicit nil
func (o *AssetSudiInfoAllOf) UnsetSudiCertificate() {
	o.SudiCertificate.Unset()
}

func (o AssetSudiInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Signature != nil {
		toSerialize["Signature"] = o.Signature
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.SudiCertificate.IsSet() {
		toSerialize["SudiCertificate"] = o.SudiCertificate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetSudiInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetSudiInfoAllOf := _AssetSudiInfoAllOf{}

	if err = json.Unmarshal(bytes, &varAssetSudiInfoAllOf); err == nil {
		*o = AssetSudiInfoAllOf(varAssetSudiInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Signature")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SudiCertificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetSudiInfoAllOf struct {
	value *AssetSudiInfoAllOf
	isSet bool
}

func (v NullableAssetSudiInfoAllOf) Get() *AssetSudiInfoAllOf {
	return v.value
}

func (v *NullableAssetSudiInfoAllOf) Set(val *AssetSudiInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSudiInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSudiInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSudiInfoAllOf(val *AssetSudiInfoAllOf) *NullableAssetSudiInfoAllOf {
	return &NullableAssetSudiInfoAllOf{value: val, isSet: true}
}

func (v NullableAssetSudiInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSudiInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
