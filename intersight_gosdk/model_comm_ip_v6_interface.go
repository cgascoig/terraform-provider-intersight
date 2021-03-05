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
	"reflect"
	"strings"
)

// CommIpV6Interface The configuration data of a single IPv6 interface, including IP address, IPv6 prefix and default gateway.
type CommIpV6Interface struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The IPv6 address of the default gateway.
	Gateway *string `json:"Gateway,omitempty"`
	// The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons.
	IpAddress *string `json:"IpAddress,omitempty"`
	// The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons.
	Prefix               *string `json:"Prefix,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommIpV6Interface CommIpV6Interface

// NewCommIpV6Interface instantiates a new CommIpV6Interface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommIpV6Interface(classId string, objectType string) *CommIpV6Interface {
	this := CommIpV6Interface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCommIpV6InterfaceWithDefaults instantiates a new CommIpV6Interface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommIpV6InterfaceWithDefaults() *CommIpV6Interface {
	this := CommIpV6Interface{}
	var classId string = "comm.IpV6Interface"
	this.ClassId = classId
	var objectType string = "comm.IpV6Interface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommIpV6Interface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommIpV6Interface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommIpV6Interface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CommIpV6Interface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommIpV6Interface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommIpV6Interface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *CommIpV6Interface) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV6Interface) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *CommIpV6Interface) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CommIpV6Interface) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *CommIpV6Interface) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV6Interface) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CommIpV6Interface) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *CommIpV6Interface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CommIpV6Interface) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV6Interface) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CommIpV6Interface) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CommIpV6Interface) SetPrefix(v string) {
	o.Prefix = &v
}

func (o CommIpV6Interface) MarshalJSON() ([]byte, error) {
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
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommIpV6Interface) UnmarshalJSON(bytes []byte) (err error) {
	type CommIpV6InterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The IPv6 address of the default gateway.
		Gateway *string `json:"Gateway,omitempty"`
		// The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons.
		IpAddress *string `json:"IpAddress,omitempty"`
		// The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons.
		Prefix *string `json:"Prefix,omitempty"`
	}

	varCommIpV6InterfaceWithoutEmbeddedStruct := CommIpV6InterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCommIpV6InterfaceWithoutEmbeddedStruct)
	if err == nil {
		varCommIpV6Interface := _CommIpV6Interface{}
		varCommIpV6Interface.ClassId = varCommIpV6InterfaceWithoutEmbeddedStruct.ClassId
		varCommIpV6Interface.ObjectType = varCommIpV6InterfaceWithoutEmbeddedStruct.ObjectType
		varCommIpV6Interface.Gateway = varCommIpV6InterfaceWithoutEmbeddedStruct.Gateway
		varCommIpV6Interface.IpAddress = varCommIpV6InterfaceWithoutEmbeddedStruct.IpAddress
		varCommIpV6Interface.Prefix = varCommIpV6InterfaceWithoutEmbeddedStruct.Prefix
		*o = CommIpV6Interface(varCommIpV6Interface)
	} else {
		return err
	}

	varCommIpV6Interface := _CommIpV6Interface{}

	err = json.Unmarshal(bytes, &varCommIpV6Interface)
	if err == nil {
		o.MoBaseComplexType = varCommIpV6Interface.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Prefix")

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

type NullableCommIpV6Interface struct {
	value *CommIpV6Interface
	isSet bool
}

func (v NullableCommIpV6Interface) Get() *CommIpV6Interface {
	return v.value
}

func (v *NullableCommIpV6Interface) Set(val *CommIpV6Interface) {
	v.value = val
	v.isSet = true
}

func (v NullableCommIpV6Interface) IsSet() bool {
	return v.isSet
}

func (v *NullableCommIpV6Interface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommIpV6Interface(val *CommIpV6Interface) *NullableCommIpV6Interface {
	return &NullableCommIpV6Interface{value: val, isSet: true}
}

func (v NullableCommIpV6Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommIpV6Interface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
