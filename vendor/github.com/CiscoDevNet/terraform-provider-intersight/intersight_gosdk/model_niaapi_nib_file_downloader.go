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

// NiaapiNibFileDownloader Provide a presigned url to download the metadata file from server.
type NiaapiNibFileDownloader struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Filename of this Metadata package file, folder will be handled by api.
	FileName *string `json:"FileName,omitempty"`
	// The presigned URL from server to download this file.
	PresignedUrl         *string `json:"PresignedUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiNibFileDownloader NiaapiNibFileDownloader

// NewNiaapiNibFileDownloader instantiates a new NiaapiNibFileDownloader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNibFileDownloader(classId string, objectType string) *NiaapiNibFileDownloader {
	this := NiaapiNibFileDownloader{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiNibFileDownloaderWithDefaults instantiates a new NiaapiNibFileDownloader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNibFileDownloaderWithDefaults() *NiaapiNibFileDownloader {
	this := NiaapiNibFileDownloader{}
	var classId string = "niaapi.NibFileDownloader"
	this.ClassId = classId
	var objectType string = "niaapi.NibFileDownloader"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiNibFileDownloader) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloader) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiNibFileDownloader) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiNibFileDownloader) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloader) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiNibFileDownloader) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *NiaapiNibFileDownloader) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloader) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *NiaapiNibFileDownloader) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *NiaapiNibFileDownloader) SetFileName(v string) {
	o.FileName = &v
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise.
func (o *NiaapiNibFileDownloader) GetPresignedUrl() string {
	if o == nil || o.PresignedUrl == nil {
		var ret string
		return ret
	}
	return *o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloader) GetPresignedUrlOk() (*string, bool) {
	if o == nil || o.PresignedUrl == nil {
		return nil, false
	}
	return o.PresignedUrl, true
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *NiaapiNibFileDownloader) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl != nil {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given string and assigns it to the PresignedUrl field.
func (o *NiaapiNibFileDownloader) SetPresignedUrl(v string) {
	o.PresignedUrl = &v
}

func (o NiaapiNibFileDownloader) MarshalJSON() ([]byte, error) {
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
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.PresignedUrl != nil {
		toSerialize["PresignedUrl"] = o.PresignedUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiNibFileDownloader) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiNibFileDownloaderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Filename of this Metadata package file, folder will be handled by api.
		FileName *string `json:"FileName,omitempty"`
		// The presigned URL from server to download this file.
		PresignedUrl *string `json:"PresignedUrl,omitempty"`
	}

	varNiaapiNibFileDownloaderWithoutEmbeddedStruct := NiaapiNibFileDownloaderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiNibFileDownloaderWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiNibFileDownloader := _NiaapiNibFileDownloader{}
		varNiaapiNibFileDownloader.ClassId = varNiaapiNibFileDownloaderWithoutEmbeddedStruct.ClassId
		varNiaapiNibFileDownloader.ObjectType = varNiaapiNibFileDownloaderWithoutEmbeddedStruct.ObjectType
		varNiaapiNibFileDownloader.FileName = varNiaapiNibFileDownloaderWithoutEmbeddedStruct.FileName
		varNiaapiNibFileDownloader.PresignedUrl = varNiaapiNibFileDownloaderWithoutEmbeddedStruct.PresignedUrl
		*o = NiaapiNibFileDownloader(varNiaapiNibFileDownloader)
	} else {
		return err
	}

	varNiaapiNibFileDownloader := _NiaapiNibFileDownloader{}

	err = json.Unmarshal(bytes, &varNiaapiNibFileDownloader)
	if err == nil {
		o.MoBaseMo = varNiaapiNibFileDownloader.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "PresignedUrl")

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

type NullableNiaapiNibFileDownloader struct {
	value *NiaapiNibFileDownloader
	isSet bool
}

func (v NullableNiaapiNibFileDownloader) Get() *NiaapiNibFileDownloader {
	return v.value
}

func (v *NullableNiaapiNibFileDownloader) Set(val *NiaapiNibFileDownloader) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNibFileDownloader) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNibFileDownloader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNibFileDownloader(val *NiaapiNibFileDownloader) *NullableNiaapiNibFileDownloader {
	return &NullableNiaapiNibFileDownloader{value: val, isSet: true}
}

func (v NullableNiaapiNibFileDownloader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNibFileDownloader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
