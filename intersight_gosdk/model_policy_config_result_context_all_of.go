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

// PolicyConfigResultContextAllOf Definition of the list of properties defined in 'policy.ConfigResultContext', excluding properties defined in parent classes.
type PolicyConfigResultContextAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The data of the object present in config result context.
	EntityData interface{} `json:"EntityData,omitempty"`
	// The Moid of the object present in config result context.
	EntityMoid *string `json:"EntityMoid,omitempty"`
	// The name of the object present in config result context.
	EntityName *string `json:"EntityName,omitempty"`
	// The type of the object present in config result context.
	EntityType           *string `json:"EntityType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyConfigResultContextAllOf PolicyConfigResultContextAllOf

// NewPolicyConfigResultContextAllOf instantiates a new PolicyConfigResultContextAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigResultContextAllOf(classId string, objectType string) *PolicyConfigResultContextAllOf {
	this := PolicyConfigResultContextAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyConfigResultContextAllOfWithDefaults instantiates a new PolicyConfigResultContextAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigResultContextAllOfWithDefaults() *PolicyConfigResultContextAllOf {
	this := PolicyConfigResultContextAllOf{}
	var classId string = "policy.ConfigResultContext"
	this.ClassId = classId
	var objectType string = "policy.ConfigResultContext"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyConfigResultContextAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContextAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyConfigResultContextAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyConfigResultContextAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContextAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyConfigResultContextAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEntityData returns the EntityData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigResultContextAllOf) GetEntityData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigResultContextAllOf) GetEntityDataOk() (*interface{}, bool) {
	if o == nil || o.EntityData == nil {
		return nil, false
	}
	return &o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *PolicyConfigResultContextAllOf) HasEntityData() bool {
	if o != nil && o.EntityData != nil {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given interface{} and assigns it to the EntityData field.
func (o *PolicyConfigResultContextAllOf) SetEntityData(v interface{}) {
	o.EntityData = v
}

// GetEntityMoid returns the EntityMoid field value if set, zero value otherwise.
func (o *PolicyConfigResultContextAllOf) GetEntityMoid() string {
	if o == nil || o.EntityMoid == nil {
		var ret string
		return ret
	}
	return *o.EntityMoid
}

// GetEntityMoidOk returns a tuple with the EntityMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContextAllOf) GetEntityMoidOk() (*string, bool) {
	if o == nil || o.EntityMoid == nil {
		return nil, false
	}
	return o.EntityMoid, true
}

// HasEntityMoid returns a boolean if a field has been set.
func (o *PolicyConfigResultContextAllOf) HasEntityMoid() bool {
	if o != nil && o.EntityMoid != nil {
		return true
	}

	return false
}

// SetEntityMoid gets a reference to the given string and assigns it to the EntityMoid field.
func (o *PolicyConfigResultContextAllOf) SetEntityMoid(v string) {
	o.EntityMoid = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *PolicyConfigResultContextAllOf) GetEntityName() string {
	if o == nil || o.EntityName == nil {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContextAllOf) GetEntityNameOk() (*string, bool) {
	if o == nil || o.EntityName == nil {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *PolicyConfigResultContextAllOf) HasEntityName() bool {
	if o != nil && o.EntityName != nil {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *PolicyConfigResultContextAllOf) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *PolicyConfigResultContextAllOf) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContextAllOf) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *PolicyConfigResultContextAllOf) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *PolicyConfigResultContextAllOf) SetEntityType(v string) {
	o.EntityType = &v
}

func (o PolicyConfigResultContextAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EntityData != nil {
		toSerialize["EntityData"] = o.EntityData
	}
	if o.EntityMoid != nil {
		toSerialize["EntityMoid"] = o.EntityMoid
	}
	if o.EntityName != nil {
		toSerialize["EntityName"] = o.EntityName
	}
	if o.EntityType != nil {
		toSerialize["EntityType"] = o.EntityType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyConfigResultContextAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyConfigResultContextAllOf := _PolicyConfigResultContextAllOf{}

	if err = json.Unmarshal(bytes, &varPolicyConfigResultContextAllOf); err == nil {
		*o = PolicyConfigResultContextAllOf(varPolicyConfigResultContextAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EntityData")
		delete(additionalProperties, "EntityMoid")
		delete(additionalProperties, "EntityName")
		delete(additionalProperties, "EntityType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyConfigResultContextAllOf struct {
	value *PolicyConfigResultContextAllOf
	isSet bool
}

func (v NullablePolicyConfigResultContextAllOf) Get() *PolicyConfigResultContextAllOf {
	return v.value
}

func (v *NullablePolicyConfigResultContextAllOf) Set(val *PolicyConfigResultContextAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigResultContextAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigResultContextAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigResultContextAllOf(val *PolicyConfigResultContextAllOf) *NullablePolicyConfigResultContextAllOf {
	return &NullablePolicyConfigResultContextAllOf{value: val, isSet: true}
}

func (v NullablePolicyConfigResultContextAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigResultContextAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
