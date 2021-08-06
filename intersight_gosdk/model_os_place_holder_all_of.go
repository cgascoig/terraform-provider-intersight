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

// OsPlaceHolderAllOf Definition of the list of properties defined in 'os.PlaceHolder', excluding properties defined in parent classes.
type OsPlaceHolderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flag to indicate if value is set. Value will be used to check if any edit.
	IsValueSet *bool                      `json:"IsValueSet,omitempty"`
	Type       *WorkflowPrimitiveDataType `json:"Type,omitempty"`
	// Value for placeholder provided by user.
	Value                interface{} `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsPlaceHolderAllOf OsPlaceHolderAllOf

// NewOsPlaceHolderAllOf instantiates a new OsPlaceHolderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsPlaceHolderAllOf(classId string, objectType string) *OsPlaceHolderAllOf {
	this := OsPlaceHolderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isValueSet bool = true
	this.IsValueSet = &isValueSet
	return &this
}

// NewOsPlaceHolderAllOfWithDefaults instantiates a new OsPlaceHolderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsPlaceHolderAllOfWithDefaults() *OsPlaceHolderAllOf {
	this := OsPlaceHolderAllOf{}
	var classId string = "os.PlaceHolder"
	this.ClassId = classId
	var objectType string = "os.PlaceHolder"
	this.ObjectType = objectType
	var isValueSet bool = true
	this.IsValueSet = &isValueSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsPlaceHolderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsPlaceHolderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsPlaceHolderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsPlaceHolderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsPlaceHolderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsPlaceHolderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsValueSet returns the IsValueSet field value if set, zero value otherwise.
func (o *OsPlaceHolderAllOf) GetIsValueSet() bool {
	if o == nil || o.IsValueSet == nil {
		var ret bool
		return ret
	}
	return *o.IsValueSet
}

// GetIsValueSetOk returns a tuple with the IsValueSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPlaceHolderAllOf) GetIsValueSetOk() (*bool, bool) {
	if o == nil || o.IsValueSet == nil {
		return nil, false
	}
	return o.IsValueSet, true
}

// HasIsValueSet returns a boolean if a field has been set.
func (o *OsPlaceHolderAllOf) HasIsValueSet() bool {
	if o != nil && o.IsValueSet != nil {
		return true
	}

	return false
}

// SetIsValueSet gets a reference to the given bool and assigns it to the IsValueSet field.
func (o *OsPlaceHolderAllOf) SetIsValueSet(v bool) {
	o.IsValueSet = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OsPlaceHolderAllOf) GetType() WorkflowPrimitiveDataType {
	if o == nil || o.Type == nil {
		var ret WorkflowPrimitiveDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPlaceHolderAllOf) GetTypeOk() (*WorkflowPrimitiveDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OsPlaceHolderAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given WorkflowPrimitiveDataType and assigns it to the Type field.
func (o *OsPlaceHolderAllOf) SetType(v WorkflowPrimitiveDataType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsPlaceHolderAllOf) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsPlaceHolderAllOf) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OsPlaceHolderAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *OsPlaceHolderAllOf) SetValue(v interface{}) {
	o.Value = v
}

func (o OsPlaceHolderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsValueSet != nil {
		toSerialize["IsValueSet"] = o.IsValueSet
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsPlaceHolderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsPlaceHolderAllOf := _OsPlaceHolderAllOf{}

	if err = json.Unmarshal(bytes, &varOsPlaceHolderAllOf); err == nil {
		*o = OsPlaceHolderAllOf(varOsPlaceHolderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsValueSet")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsPlaceHolderAllOf struct {
	value *OsPlaceHolderAllOf
	isSet bool
}

func (v NullableOsPlaceHolderAllOf) Get() *OsPlaceHolderAllOf {
	return v.value
}

func (v *NullableOsPlaceHolderAllOf) Set(val *OsPlaceHolderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsPlaceHolderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsPlaceHolderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsPlaceHolderAllOf(val *OsPlaceHolderAllOf) *NullableOsPlaceHolderAllOf {
	return &NullableOsPlaceHolderAllOf{value: val, isSet: true}
}

func (v NullableOsPlaceHolderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsPlaceHolderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
