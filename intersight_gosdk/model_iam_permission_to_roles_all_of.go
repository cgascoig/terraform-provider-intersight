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

// IamPermissionToRolesAllOf Definition of the list of properties defined in 'iam.PermissionToRoles', excluding properties defined in parent classes.
type IamPermissionToRolesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string    `json:"ObjectType"`
	Permission           *MoMoRef  `json:"Permission,omitempty"`
	Roles                []MoMoRef `json:"Roles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPermissionToRolesAllOf IamPermissionToRolesAllOf

// NewIamPermissionToRolesAllOf instantiates a new IamPermissionToRolesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermissionToRolesAllOf(classId string, objectType string) *IamPermissionToRolesAllOf {
	this := IamPermissionToRolesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPermissionToRolesAllOfWithDefaults instantiates a new IamPermissionToRolesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionToRolesAllOfWithDefaults() *IamPermissionToRolesAllOf {
	this := IamPermissionToRolesAllOf{}
	var classId string = "iam.PermissionToRoles"
	this.ClassId = classId
	var objectType string = "iam.PermissionToRoles"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPermissionToRolesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPermissionToRolesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPermissionToRolesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPermissionToRolesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPermissionToRolesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPermissionToRolesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamPermissionToRolesAllOf) GetPermission() MoMoRef {
	if o == nil || o.Permission == nil {
		var ret MoMoRef
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionToRolesAllOf) GetPermissionOk() (*MoMoRef, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamPermissionToRolesAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given MoMoRef and assigns it to the Permission field.
func (o *IamPermissionToRolesAllOf) SetPermission(v MoMoRef) {
	o.Permission = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionToRolesAllOf) GetRoles() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionToRolesAllOf) GetRolesOk() (*[]MoMoRef, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamPermissionToRolesAllOf) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []MoMoRef and assigns it to the Roles field.
func (o *IamPermissionToRolesAllOf) SetRoles(v []MoMoRef) {
	o.Roles = v
}

func (o IamPermissionToRolesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPermissionToRolesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamPermissionToRolesAllOf := _IamPermissionToRolesAllOf{}

	if err = json.Unmarshal(bytes, &varIamPermissionToRolesAllOf); err == nil {
		*o = IamPermissionToRolesAllOf(varIamPermissionToRolesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "Roles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamPermissionToRolesAllOf struct {
	value *IamPermissionToRolesAllOf
	isSet bool
}

func (v NullableIamPermissionToRolesAllOf) Get() *IamPermissionToRolesAllOf {
	return v.value
}

func (v *NullableIamPermissionToRolesAllOf) Set(val *IamPermissionToRolesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPermissionToRolesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPermissionToRolesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPermissionToRolesAllOf(val *IamPermissionToRolesAllOf) *NullableIamPermissionToRolesAllOf {
	return &NullableIamPermissionToRolesAllOf{value: val, isSet: true}
}

func (v NullableIamPermissionToRolesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPermissionToRolesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
