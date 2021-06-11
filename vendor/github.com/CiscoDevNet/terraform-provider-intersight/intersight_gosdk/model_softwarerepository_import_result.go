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

// SoftwarerepositoryImportResult The status of the import operation.
type SoftwarerepositoryImportResult struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The reason for the failure of an import operation, if applicable.
	ErrorMessage *string `json:"ErrorMessage,omitempty"`
	// The progress percentage for the import operation.
	Progress             *int64 `json:"Progress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryImportResult SoftwarerepositoryImportResult

// NewSoftwarerepositoryImportResult instantiates a new SoftwarerepositoryImportResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryImportResult(classId string, objectType string) *SoftwarerepositoryImportResult {
	this := SoftwarerepositoryImportResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryImportResultWithDefaults instantiates a new SoftwarerepositoryImportResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryImportResultWithDefaults() *SoftwarerepositoryImportResult {
	this := SoftwarerepositoryImportResult{}
	var classId string = "softwarerepository.ImportResult"
	this.ClassId = classId
	var objectType string = "softwarerepository.ImportResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryImportResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryImportResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryImportResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryImportResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SoftwarerepositoryImportResult) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SoftwarerepositoryImportResult) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SoftwarerepositoryImportResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *SoftwarerepositoryImportResult) GetProgress() int64 {
	if o == nil || o.Progress == nil {
		var ret int64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResult) GetProgressOk() (*int64, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *SoftwarerepositoryImportResult) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int64 and assigns it to the Progress field.
func (o *SoftwarerepositoryImportResult) SetProgress(v int64) {
	o.Progress = &v
}

func (o SoftwarerepositoryImportResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ErrorMessage != nil {
		toSerialize["ErrorMessage"] = o.ErrorMessage
	}
	if o.Progress != nil {
		toSerialize["Progress"] = o.Progress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryImportResult) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryImportResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The reason for the failure of an import operation, if applicable.
		ErrorMessage *string `json:"ErrorMessage,omitempty"`
		// The progress percentage for the import operation.
		Progress *int64 `json:"Progress,omitempty"`
	}

	varSoftwarerepositoryImportResultWithoutEmbeddedStruct := SoftwarerepositoryImportResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryImportResultWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryImportResult := _SoftwarerepositoryImportResult{}
		varSoftwarerepositoryImportResult.ClassId = varSoftwarerepositoryImportResultWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryImportResult.ObjectType = varSoftwarerepositoryImportResultWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryImportResult.ErrorMessage = varSoftwarerepositoryImportResultWithoutEmbeddedStruct.ErrorMessage
		varSoftwarerepositoryImportResult.Progress = varSoftwarerepositoryImportResultWithoutEmbeddedStruct.Progress
		*o = SoftwarerepositoryImportResult(varSoftwarerepositoryImportResult)
	} else {
		return err
	}

	varSoftwarerepositoryImportResult := _SoftwarerepositoryImportResult{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryImportResult)
	if err == nil {
		o.MoBaseComplexType = varSoftwarerepositoryImportResult.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorMessage")
		delete(additionalProperties, "Progress")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableSoftwarerepositoryImportResult struct {
	value *SoftwarerepositoryImportResult
	isSet bool
}

func (v NullableSoftwarerepositoryImportResult) Get() *SoftwarerepositoryImportResult {
	return v.value
}

func (v *NullableSoftwarerepositoryImportResult) Set(val *SoftwarerepositoryImportResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryImportResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryImportResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryImportResult(val *SoftwarerepositoryImportResult) *NullableSoftwarerepositoryImportResult {
	return &NullableSoftwarerepositoryImportResult{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryImportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryImportResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
