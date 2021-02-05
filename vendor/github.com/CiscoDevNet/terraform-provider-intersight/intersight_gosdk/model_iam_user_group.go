/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IamUserGroup User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.
type IamUserGroup struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the user group which the dynamic user belongs to.
	Name         *string                      `json:"Name,omitempty"`
	Idp          *IamIdpRelationship          `json:"Idp,omitempty"`
	Idpreference *IamIdpReferenceRelationship `json:"Idpreference,omitempty"`
	// An array of relationships to iamPermission resources.
	Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
	Qualifier   *IamQualifierRelationship   `json:"Qualifier,omitempty"`
	// An array of relationships to iamUser resources.
	Users                []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamUserGroup IamUserGroup

// NewIamUserGroup instantiates a new IamUserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserGroup(classId string, objectType string) *IamUserGroup {
	this := IamUserGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamUserGroupWithDefaults instantiates a new IamUserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserGroupWithDefaults() *IamUserGroup {
	this := IamUserGroup{}
	var classId string = "iam.UserGroup"
	this.ClassId = classId
	var objectType string = "iam.UserGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamUserGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamUserGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamUserGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamUserGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamUserGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamUserGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamUserGroup) SetName(v string) {
	o.Name = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamUserGroup) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamUserGroup) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamUserGroup) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetIdpreference returns the Idpreference field value if set, zero value otherwise.
func (o *IamUserGroup) GetIdpreference() IamIdpReferenceRelationship {
	if o == nil || o.Idpreference == nil {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.Idpreference
}

// GetIdpreferenceOk returns a tuple with the Idpreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil || o.Idpreference == nil {
		return nil, false
	}
	return o.Idpreference, true
}

// HasIdpreference returns a boolean if a field has been set.
func (o *IamUserGroup) HasIdpreference() bool {
	if o != nil && o.Idpreference != nil {
		return true
	}

	return false
}

// SetIdpreference gets a reference to the given IamIdpReferenceRelationship and assigns it to the Idpreference field.
func (o *IamUserGroup) SetIdpreference(v IamIdpReferenceRelationship) {
	o.Idpreference = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserGroup) GetPermissions() []IamPermissionRelationship {
	if o == nil {
		var ret []IamPermissionRelationship
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserGroup) GetPermissionsOk() (*[]IamPermissionRelationship, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamUserGroup) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []IamPermissionRelationship and assigns it to the Permissions field.
func (o *IamUserGroup) SetPermissions(v []IamPermissionRelationship) {
	o.Permissions = v
}

// GetQualifier returns the Qualifier field value if set, zero value otherwise.
func (o *IamUserGroup) GetQualifier() IamQualifierRelationship {
	if o == nil || o.Qualifier == nil {
		var ret IamQualifierRelationship
		return ret
	}
	return *o.Qualifier
}

// GetQualifierOk returns a tuple with the Qualifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetQualifierOk() (*IamQualifierRelationship, bool) {
	if o == nil || o.Qualifier == nil {
		return nil, false
	}
	return o.Qualifier, true
}

// HasQualifier returns a boolean if a field has been set.
func (o *IamUserGroup) HasQualifier() bool {
	if o != nil && o.Qualifier != nil {
		return true
	}

	return false
}

// SetQualifier gets a reference to the given IamQualifierRelationship and assigns it to the Qualifier field.
func (o *IamUserGroup) SetQualifier(v IamQualifierRelationship) {
	o.Qualifier = &v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserGroup) GetUsers() []IamUserRelationship {
	if o == nil {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserGroup) GetUsersOk() (*[]IamUserRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamUserGroup) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamUserGroup) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamUserGroup) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.Idpreference != nil {
		toSerialize["Idpreference"] = o.Idpreference
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}
	if o.Qualifier != nil {
		toSerialize["Qualifier"] = o.Qualifier
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamUserGroup) UnmarshalJSON(bytes []byte) (err error) {
	type IamUserGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the user group which the dynamic user belongs to.
		Name         *string                      `json:"Name,omitempty"`
		Idp          *IamIdpRelationship          `json:"Idp,omitempty"`
		Idpreference *IamIdpReferenceRelationship `json:"Idpreference,omitempty"`
		// An array of relationships to iamPermission resources.
		Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
		Qualifier   *IamQualifierRelationship   `json:"Qualifier,omitempty"`
		// An array of relationships to iamUser resources.
		Users []IamUserRelationship `json:"Users,omitempty"`
	}

	varIamUserGroupWithoutEmbeddedStruct := IamUserGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamUserGroupWithoutEmbeddedStruct)
	if err == nil {
		varIamUserGroup := _IamUserGroup{}
		varIamUserGroup.ClassId = varIamUserGroupWithoutEmbeddedStruct.ClassId
		varIamUserGroup.ObjectType = varIamUserGroupWithoutEmbeddedStruct.ObjectType
		varIamUserGroup.Name = varIamUserGroupWithoutEmbeddedStruct.Name
		varIamUserGroup.Idp = varIamUserGroupWithoutEmbeddedStruct.Idp
		varIamUserGroup.Idpreference = varIamUserGroupWithoutEmbeddedStruct.Idpreference
		varIamUserGroup.Permissions = varIamUserGroupWithoutEmbeddedStruct.Permissions
		varIamUserGroup.Qualifier = varIamUserGroupWithoutEmbeddedStruct.Qualifier
		varIamUserGroup.Users = varIamUserGroupWithoutEmbeddedStruct.Users
		*o = IamUserGroup(varIamUserGroup)
	} else {
		return err
	}

	varIamUserGroup := _IamUserGroup{}

	err = json.Unmarshal(bytes, &varIamUserGroup)
	if err == nil {
		o.MoBaseMo = varIamUserGroup.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "Idpreference")
		delete(additionalProperties, "Permissions")
		delete(additionalProperties, "Qualifier")
		delete(additionalProperties, "Users")

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

type NullableIamUserGroup struct {
	value *IamUserGroup
	isSet bool
}

func (v NullableIamUserGroup) Get() *IamUserGroup {
	return v.value
}

func (v *NullableIamUserGroup) Set(val *IamUserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserGroup(val *IamUserGroup) *NullableIamUserGroup {
	return &NullableIamUserGroup{value: val, isSet: true}
}

func (v NullableIamUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
