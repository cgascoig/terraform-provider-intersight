/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ServerBaseProfile A profile specifying configuration settings for a physical server.
type ServerBaseProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform *string `json:"TargetPlatform,omitempty"`
	// UUID address allocation type selected to assign an UUID address for the server. * `NONE` - The user did not assign any UUID address. * `STATIC` - The user assigns a static UUID address. * `POOL` - The user selects a pool from which the address will be leased.
	UuidAddressType      *string                   `json:"UuidAddressType,omitempty"`
	UuidPool             *UuidpoolPoolRelationship `json:"UuidPool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerBaseProfile ServerBaseProfile

// NewServerBaseProfile instantiates a new ServerBaseProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerBaseProfile(classId string, objectType string) *ServerBaseProfile {
	this := ServerBaseProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	var uuidAddressType string = "NONE"
	this.UuidAddressType = &uuidAddressType
	return &this
}

// NewServerBaseProfileWithDefaults instantiates a new ServerBaseProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerBaseProfileWithDefaults() *ServerBaseProfile {
	this := ServerBaseProfile{}
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	var uuidAddressType string = "NONE"
	this.UuidAddressType = &uuidAddressType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerBaseProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerBaseProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerBaseProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerBaseProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerBaseProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerBaseProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *ServerBaseProfile) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBaseProfile) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *ServerBaseProfile) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *ServerBaseProfile) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetUuidAddressType returns the UuidAddressType field value if set, zero value otherwise.
func (o *ServerBaseProfile) GetUuidAddressType() string {
	if o == nil || o.UuidAddressType == nil {
		var ret string
		return ret
	}
	return *o.UuidAddressType
}

// GetUuidAddressTypeOk returns a tuple with the UuidAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBaseProfile) GetUuidAddressTypeOk() (*string, bool) {
	if o == nil || o.UuidAddressType == nil {
		return nil, false
	}
	return o.UuidAddressType, true
}

// HasUuidAddressType returns a boolean if a field has been set.
func (o *ServerBaseProfile) HasUuidAddressType() bool {
	if o != nil && o.UuidAddressType != nil {
		return true
	}

	return false
}

// SetUuidAddressType gets a reference to the given string and assigns it to the UuidAddressType field.
func (o *ServerBaseProfile) SetUuidAddressType(v string) {
	o.UuidAddressType = &v
}

// GetUuidPool returns the UuidPool field value if set, zero value otherwise.
func (o *ServerBaseProfile) GetUuidPool() UuidpoolPoolRelationship {
	if o == nil || o.UuidPool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.UuidPool
}

// GetUuidPoolOk returns a tuple with the UuidPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBaseProfile) GetUuidPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.UuidPool == nil {
		return nil, false
	}
	return o.UuidPool, true
}

// HasUuidPool returns a boolean if a field has been set.
func (o *ServerBaseProfile) HasUuidPool() bool {
	if o != nil && o.UuidPool != nil {
		return true
	}

	return false
}

// SetUuidPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the UuidPool field.
func (o *ServerBaseProfile) SetUuidPool(v UuidpoolPoolRelationship) {
	o.UuidPool = &v
}

func (o ServerBaseProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.UuidAddressType != nil {
		toSerialize["UuidAddressType"] = o.UuidAddressType
	}
	if o.UuidPool != nil {
		toSerialize["UuidPool"] = o.UuidPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerBaseProfile) UnmarshalJSON(bytes []byte) (err error) {
	type ServerBaseProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		TargetPlatform *string `json:"TargetPlatform,omitempty"`
		// UUID address allocation type selected to assign an UUID address for the server. * `NONE` - The user did not assign any UUID address. * `STATIC` - The user assigns a static UUID address. * `POOL` - The user selects a pool from which the address will be leased.
		UuidAddressType *string                   `json:"UuidAddressType,omitempty"`
		UuidPool        *UuidpoolPoolRelationship `json:"UuidPool,omitempty"`
	}

	varServerBaseProfileWithoutEmbeddedStruct := ServerBaseProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varServerBaseProfileWithoutEmbeddedStruct)
	if err == nil {
		varServerBaseProfile := _ServerBaseProfile{}
		varServerBaseProfile.ClassId = varServerBaseProfileWithoutEmbeddedStruct.ClassId
		varServerBaseProfile.ObjectType = varServerBaseProfileWithoutEmbeddedStruct.ObjectType
		varServerBaseProfile.TargetPlatform = varServerBaseProfileWithoutEmbeddedStruct.TargetPlatform
		varServerBaseProfile.UuidAddressType = varServerBaseProfileWithoutEmbeddedStruct.UuidAddressType
		varServerBaseProfile.UuidPool = varServerBaseProfileWithoutEmbeddedStruct.UuidPool
		*o = ServerBaseProfile(varServerBaseProfile)
	} else {
		return err
	}

	varServerBaseProfile := _ServerBaseProfile{}

	err = json.Unmarshal(bytes, &varServerBaseProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varServerBaseProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "UuidAddressType")
		delete(additionalProperties, "UuidPool")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableServerBaseProfile struct {
	value *ServerBaseProfile
	isSet bool
}

func (v NullableServerBaseProfile) Get() *ServerBaseProfile {
	return v.value
}

func (v *NullableServerBaseProfile) Set(val *ServerBaseProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableServerBaseProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableServerBaseProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerBaseProfile(val *ServerBaseProfile) *NullableServerBaseProfile {
	return &NullableServerBaseProfile{value: val, isSet: true}
}

func (v NullableServerBaseProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerBaseProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
