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

// TamTextFsmTemplateDataSourceAllOf Definition of the list of properties defined in 'tam.TextFsmTemplateDataSource', excluding properties defined in parent classes.
type TamTextFsmTemplateDataSourceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Command used to gather data needed to evaluate field notice applicability.
	Cmd                  *string `json:"Cmd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamTextFsmTemplateDataSourceAllOf TamTextFsmTemplateDataSourceAllOf

// NewTamTextFsmTemplateDataSourceAllOf instantiates a new TamTextFsmTemplateDataSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamTextFsmTemplateDataSourceAllOf(classId string, objectType string) *TamTextFsmTemplateDataSourceAllOf {
	this := TamTextFsmTemplateDataSourceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamTextFsmTemplateDataSourceAllOfWithDefaults instantiates a new TamTextFsmTemplateDataSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamTextFsmTemplateDataSourceAllOfWithDefaults() *TamTextFsmTemplateDataSourceAllOf {
	this := TamTextFsmTemplateDataSourceAllOf{}
	var classId string = "tam.TextFsmTemplateDataSource"
	this.ClassId = classId
	var objectType string = "tam.TextFsmTemplateDataSource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamTextFsmTemplateDataSourceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamTextFsmTemplateDataSourceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamTextFsmTemplateDataSourceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamTextFsmTemplateDataSourceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamTextFsmTemplateDataSourceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamTextFsmTemplateDataSourceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *TamTextFsmTemplateDataSourceAllOf) GetCmd() string {
	if o == nil || o.Cmd == nil {
		var ret string
		return ret
	}
	return *o.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamTextFsmTemplateDataSourceAllOf) GetCmdOk() (*string, bool) {
	if o == nil || o.Cmd == nil {
		return nil, false
	}
	return o.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *TamTextFsmTemplateDataSourceAllOf) HasCmd() bool {
	if o != nil && o.Cmd != nil {
		return true
	}

	return false
}

// SetCmd gets a reference to the given string and assigns it to the Cmd field.
func (o *TamTextFsmTemplateDataSourceAllOf) SetCmd(v string) {
	o.Cmd = &v
}

func (o TamTextFsmTemplateDataSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cmd != nil {
		toSerialize["Cmd"] = o.Cmd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamTextFsmTemplateDataSourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamTextFsmTemplateDataSourceAllOf := _TamTextFsmTemplateDataSourceAllOf{}

	if err = json.Unmarshal(bytes, &varTamTextFsmTemplateDataSourceAllOf); err == nil {
		*o = TamTextFsmTemplateDataSourceAllOf(varTamTextFsmTemplateDataSourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cmd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamTextFsmTemplateDataSourceAllOf struct {
	value *TamTextFsmTemplateDataSourceAllOf
	isSet bool
}

func (v NullableTamTextFsmTemplateDataSourceAllOf) Get() *TamTextFsmTemplateDataSourceAllOf {
	return v.value
}

func (v *NullableTamTextFsmTemplateDataSourceAllOf) Set(val *TamTextFsmTemplateDataSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamTextFsmTemplateDataSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamTextFsmTemplateDataSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamTextFsmTemplateDataSourceAllOf(val *TamTextFsmTemplateDataSourceAllOf) *NullableTamTextFsmTemplateDataSourceAllOf {
	return &NullableTamTextFsmTemplateDataSourceAllOf{value: val, isSet: true}
}

func (v NullableTamTextFsmTemplateDataSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamTextFsmTemplateDataSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
