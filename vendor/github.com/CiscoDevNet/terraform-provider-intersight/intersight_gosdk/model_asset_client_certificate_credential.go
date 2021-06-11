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
	"reflect"
	"strings"
)

// AssetClientCertificateCredential A credential which performs authentication with TLS client certificates.
type AssetClientCertificateCredential struct {
	AssetCredential
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// PEM encoded x509 client certificate used to authenticate with the target.
	ClientCertificate *string `json:"ClientCertificate,omitempty"`
	// PEM encoded private key used to authenticate with the target.
	ClientKey *string `json:"ClientKey,omitempty"`
	// Indicates whether the value of the 'clientKey' property has been set.
	IsClientKeySet       *bool `json:"IsClientKeySet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetClientCertificateCredential AssetClientCertificateCredential

// NewAssetClientCertificateCredential instantiates a new AssetClientCertificateCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetClientCertificateCredential(classId string, objectType string) *AssetClientCertificateCredential {
	this := AssetClientCertificateCredential{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isClientKeySet bool = false
	this.IsClientKeySet = &isClientKeySet
	return &this
}

// NewAssetClientCertificateCredentialWithDefaults instantiates a new AssetClientCertificateCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetClientCertificateCredentialWithDefaults() *AssetClientCertificateCredential {
	this := AssetClientCertificateCredential{}
	var classId string = "asset.ClientCertificateCredential"
	this.ClassId = classId
	var objectType string = "asset.ClientCertificateCredential"
	this.ObjectType = objectType
	var isClientKeySet bool = false
	this.IsClientKeySet = &isClientKeySet
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetClientCertificateCredential) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetClientCertificateCredential) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetClientCertificateCredential) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetClientCertificateCredential) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetClientCertificateCredential) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetClientCertificateCredential) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *AssetClientCertificateCredential) GetClientCertificate() string {
	if o == nil || o.ClientCertificate == nil {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClientCertificateCredential) GetClientCertificateOk() (*string, bool) {
	if o == nil || o.ClientCertificate == nil {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *AssetClientCertificateCredential) HasClientCertificate() bool {
	if o != nil && o.ClientCertificate != nil {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *AssetClientCertificateCredential) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise.
func (o *AssetClientCertificateCredential) GetClientKey() string {
	if o == nil || o.ClientKey == nil {
		var ret string
		return ret
	}
	return *o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClientCertificateCredential) GetClientKeyOk() (*string, bool) {
	if o == nil || o.ClientKey == nil {
		return nil, false
	}
	return o.ClientKey, true
}

// HasClientKey returns a boolean if a field has been set.
func (o *AssetClientCertificateCredential) HasClientKey() bool {
	if o != nil && o.ClientKey != nil {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given string and assigns it to the ClientKey field.
func (o *AssetClientCertificateCredential) SetClientKey(v string) {
	o.ClientKey = &v
}

// GetIsClientKeySet returns the IsClientKeySet field value if set, zero value otherwise.
func (o *AssetClientCertificateCredential) GetIsClientKeySet() bool {
	if o == nil || o.IsClientKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsClientKeySet
}

// GetIsClientKeySetOk returns a tuple with the IsClientKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClientCertificateCredential) GetIsClientKeySetOk() (*bool, bool) {
	if o == nil || o.IsClientKeySet == nil {
		return nil, false
	}
	return o.IsClientKeySet, true
}

// HasIsClientKeySet returns a boolean if a field has been set.
func (o *AssetClientCertificateCredential) HasIsClientKeySet() bool {
	if o != nil && o.IsClientKeySet != nil {
		return true
	}

	return false
}

// SetIsClientKeySet gets a reference to the given bool and assigns it to the IsClientKeySet field.
func (o *AssetClientCertificateCredential) SetIsClientKeySet(v bool) {
	o.IsClientKeySet = &v
}

func (o AssetClientCertificateCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetCredential, errAssetCredential := json.Marshal(o.AssetCredential)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	errAssetCredential = json.Unmarshal([]byte(serializedAssetCredential), &toSerialize)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClientCertificate != nil {
		toSerialize["ClientCertificate"] = o.ClientCertificate
	}
	if o.ClientKey != nil {
		toSerialize["ClientKey"] = o.ClientKey
	}
	if o.IsClientKeySet != nil {
		toSerialize["IsClientKeySet"] = o.IsClientKeySet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetClientCertificateCredential) UnmarshalJSON(bytes []byte) (err error) {
	type AssetClientCertificateCredentialWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// PEM encoded x509 client certificate used to authenticate with the target.
		ClientCertificate *string `json:"ClientCertificate,omitempty"`
		// PEM encoded private key used to authenticate with the target.
		ClientKey *string `json:"ClientKey,omitempty"`
		// Indicates whether the value of the 'clientKey' property has been set.
		IsClientKeySet *bool `json:"IsClientKeySet,omitempty"`
	}

	varAssetClientCertificateCredentialWithoutEmbeddedStruct := AssetClientCertificateCredentialWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetClientCertificateCredentialWithoutEmbeddedStruct)
	if err == nil {
		varAssetClientCertificateCredential := _AssetClientCertificateCredential{}
		varAssetClientCertificateCredential.ClassId = varAssetClientCertificateCredentialWithoutEmbeddedStruct.ClassId
		varAssetClientCertificateCredential.ObjectType = varAssetClientCertificateCredentialWithoutEmbeddedStruct.ObjectType
		varAssetClientCertificateCredential.ClientCertificate = varAssetClientCertificateCredentialWithoutEmbeddedStruct.ClientCertificate
		varAssetClientCertificateCredential.ClientKey = varAssetClientCertificateCredentialWithoutEmbeddedStruct.ClientKey
		varAssetClientCertificateCredential.IsClientKeySet = varAssetClientCertificateCredentialWithoutEmbeddedStruct.IsClientKeySet
		*o = AssetClientCertificateCredential(varAssetClientCertificateCredential)
	} else {
		return err
	}

	varAssetClientCertificateCredential := _AssetClientCertificateCredential{}

	err = json.Unmarshal(bytes, &varAssetClientCertificateCredential)
	if err == nil {
		o.AssetCredential = varAssetClientCertificateCredential.AssetCredential
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientCertificate")
		delete(additionalProperties, "ClientKey")
		delete(additionalProperties, "IsClientKeySet")

		// remove fields from embedded structs
		reflectAssetCredential := reflect.ValueOf(o.AssetCredential)
		for i := 0; i < reflectAssetCredential.Type().NumField(); i++ {
			t := reflectAssetCredential.Type().Field(i)

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

type NullableAssetClientCertificateCredential struct {
	value *AssetClientCertificateCredential
	isSet bool
}

func (v NullableAssetClientCertificateCredential) Get() *AssetClientCertificateCredential {
	return v.value
}

func (v *NullableAssetClientCertificateCredential) Set(val *AssetClientCertificateCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetClientCertificateCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetClientCertificateCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetClientCertificateCredential(val *AssetClientCertificateCredential) *NullableAssetClientCertificateCredential {
	return &NullableAssetClientCertificateCredential{value: val, isSet: true}
}

func (v NullableAssetClientCertificateCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetClientCertificateCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
