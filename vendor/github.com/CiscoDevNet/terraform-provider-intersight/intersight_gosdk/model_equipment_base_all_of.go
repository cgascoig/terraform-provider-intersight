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
)

// EquipmentBaseAllOf Definition of the list of properties defined in 'equipment.Base', excluding properties defined in parent classes.
type EquipmentBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// This field identifies the model of the given component.
	Model *string `json:"Model,omitempty"`
	// This field identifies the revision of the given component.
	Revision *string `json:"Revision,omitempty"`
	// This field identifies the serial of the given component.
	Serial *string `json:"Serial,omitempty"`
	// This field identifies the vendor of the given component.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentBaseAllOf EquipmentBaseAllOf

// NewEquipmentBaseAllOf instantiates a new EquipmentBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentBaseAllOf(classId string, objectType string) *EquipmentBaseAllOf {
	this := EquipmentBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentBaseAllOfWithDefaults instantiates a new EquipmentBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentBaseAllOfWithDefaults() *EquipmentBaseAllOf {
	this := EquipmentBaseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EquipmentBaseAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EquipmentBaseAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EquipmentBaseAllOf) SetModel(v string) {
	o.Model = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *EquipmentBaseAllOf) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseAllOf) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *EquipmentBaseAllOf) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *EquipmentBaseAllOf) SetRevision(v string) {
	o.Revision = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EquipmentBaseAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EquipmentBaseAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EquipmentBaseAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *EquipmentBaseAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *EquipmentBaseAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *EquipmentBaseAllOf) SetVendor(v string) {
	o.Vendor = &v
}

func (o EquipmentBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentBaseAllOf := _EquipmentBaseAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentBaseAllOf); err == nil {
		*o = EquipmentBaseAllOf(varEquipmentBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Revision")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentBaseAllOf struct {
	value *EquipmentBaseAllOf
	isSet bool
}

func (v NullableEquipmentBaseAllOf) Get() *EquipmentBaseAllOf {
	return v.value
}

func (v *NullableEquipmentBaseAllOf) Set(val *EquipmentBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentBaseAllOf(val *EquipmentBaseAllOf) *NullableEquipmentBaseAllOf {
	return &NullableEquipmentBaseAllOf{value: val, isSet: true}
}

func (v NullableEquipmentBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
