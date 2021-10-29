/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ChassisConfigChangeDetailAllOf Definition of the list of properties defined in 'chassis.ConfigChangeDetail', excluding properties defined in parent classes.
type ChassisConfigChangeDetailAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                      `json:"ObjectType"`
	Profile              *ChassisProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisConfigChangeDetailAllOf ChassisConfigChangeDetailAllOf

// NewChassisConfigChangeDetailAllOf instantiates a new ChassisConfigChangeDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisConfigChangeDetailAllOf(classId string, objectType string) *ChassisConfigChangeDetailAllOf {
	this := ChassisConfigChangeDetailAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewChassisConfigChangeDetailAllOfWithDefaults instantiates a new ChassisConfigChangeDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisConfigChangeDetailAllOfWithDefaults() *ChassisConfigChangeDetailAllOf {
	this := ChassisConfigChangeDetailAllOf{}
	var classId string = "chassis.ConfigChangeDetail"
	this.ClassId = classId
	var objectType string = "chassis.ConfigChangeDetail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisConfigChangeDetailAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigChangeDetailAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisConfigChangeDetailAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisConfigChangeDetailAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigChangeDetailAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisConfigChangeDetailAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ChassisConfigChangeDetailAllOf) GetProfile() ChassisProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret ChassisProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigChangeDetailAllOf) GetProfileOk() (*ChassisProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ChassisConfigChangeDetailAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ChassisProfileRelationship and assigns it to the Profile field.
func (o *ChassisConfigChangeDetailAllOf) SetProfile(v ChassisProfileRelationship) {
	o.Profile = &v
}

func (o ChassisConfigChangeDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisConfigChangeDetailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varChassisConfigChangeDetailAllOf := _ChassisConfigChangeDetailAllOf{}

	if err = json.Unmarshal(bytes, &varChassisConfigChangeDetailAllOf); err == nil {
		*o = ChassisConfigChangeDetailAllOf(varChassisConfigChangeDetailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChassisConfigChangeDetailAllOf struct {
	value *ChassisConfigChangeDetailAllOf
	isSet bool
}

func (v NullableChassisConfigChangeDetailAllOf) Get() *ChassisConfigChangeDetailAllOf {
	return v.value
}

func (v *NullableChassisConfigChangeDetailAllOf) Set(val *ChassisConfigChangeDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisConfigChangeDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisConfigChangeDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisConfigChangeDetailAllOf(val *ChassisConfigChangeDetailAllOf) *NullableChassisConfigChangeDetailAllOf {
	return &NullableChassisConfigChangeDetailAllOf{value: val, isSet: true}
}

func (v NullableChassisConfigChangeDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisConfigChangeDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
