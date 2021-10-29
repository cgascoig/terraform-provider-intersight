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

// VnicIscsiAuthProfileAllOf Definition of the list of properties defined in 'vnic.IscsiAuthProfile', excluding properties defined in parent classes.
type VnicIscsiAuthProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Password of Initiator/Target Interface. Enter between 12 and 16 characters, including special characters except spaces, tabs, line breaks.
	Password *string `json:"Password,omitempty"`
	// User Id of Initiator/Target Interface. Enter between 1 and 128 characters, spaces, or special characters.
	UserId               *string `json:"UserId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicIscsiAuthProfileAllOf VnicIscsiAuthProfileAllOf

// NewVnicIscsiAuthProfileAllOf instantiates a new VnicIscsiAuthProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicIscsiAuthProfileAllOf(classId string, objectType string) *VnicIscsiAuthProfileAllOf {
	this := VnicIscsiAuthProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicIscsiAuthProfileAllOfWithDefaults instantiates a new VnicIscsiAuthProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicIscsiAuthProfileAllOfWithDefaults() *VnicIscsiAuthProfileAllOf {
	this := VnicIscsiAuthProfileAllOf{}
	var classId string = "vnic.IscsiAuthProfile"
	this.ClassId = classId
	var objectType string = "vnic.IscsiAuthProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicIscsiAuthProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAuthProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicIscsiAuthProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicIscsiAuthProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAuthProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicIscsiAuthProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *VnicIscsiAuthProfileAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAuthProfileAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *VnicIscsiAuthProfileAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *VnicIscsiAuthProfileAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *VnicIscsiAuthProfileAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAuthProfileAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *VnicIscsiAuthProfileAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *VnicIscsiAuthProfileAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *VnicIscsiAuthProfileAllOf) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAuthProfileAllOf) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *VnicIscsiAuthProfileAllOf) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *VnicIscsiAuthProfileAllOf) SetUserId(v string) {
	o.UserId = &v
}

func (o VnicIscsiAuthProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicIscsiAuthProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicIscsiAuthProfileAllOf := _VnicIscsiAuthProfileAllOf{}

	if err = json.Unmarshal(bytes, &varVnicIscsiAuthProfileAllOf); err == nil {
		*o = VnicIscsiAuthProfileAllOf(varVnicIscsiAuthProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "UserId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicIscsiAuthProfileAllOf struct {
	value *VnicIscsiAuthProfileAllOf
	isSet bool
}

func (v NullableVnicIscsiAuthProfileAllOf) Get() *VnicIscsiAuthProfileAllOf {
	return v.value
}

func (v *NullableVnicIscsiAuthProfileAllOf) Set(val *VnicIscsiAuthProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicIscsiAuthProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicIscsiAuthProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicIscsiAuthProfileAllOf(val *VnicIscsiAuthProfileAllOf) *NullableVnicIscsiAuthProfileAllOf {
	return &NullableVnicIscsiAuthProfileAllOf{value: val, isSet: true}
}

func (v NullableVnicIscsiAuthProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicIscsiAuthProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
