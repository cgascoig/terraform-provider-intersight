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

// FabricServerRole Configuration object sent by user to create a server port.
type FabricServerRole struct {
	FabricPortRole
	AdditionalProperties map[string]interface{}
}

type _FabricServerRole FabricServerRole

// NewFabricServerRole instantiates a new FabricServerRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricServerRole() *FabricServerRole {
	this := FabricServerRole{}
	return &this
}

// NewFabricServerRoleWithDefaults instantiates a new FabricServerRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricServerRoleWithDefaults() *FabricServerRole {
	this := FabricServerRole{}
	return &this
}

func (o FabricServerRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortRole, errFabricPortRole := json.Marshal(o.FabricPortRole)
	if errFabricPortRole != nil {
		return []byte{}, errFabricPortRole
	}
	errFabricPortRole = json.Unmarshal([]byte(serializedFabricPortRole), &toSerialize)
	if errFabricPortRole != nil {
		return []byte{}, errFabricPortRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricServerRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricServerRoleWithoutEmbeddedStruct struct {
	}

	varFabricServerRoleWithoutEmbeddedStruct := FabricServerRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricServerRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricServerRole := _FabricServerRole{}
		*o = FabricServerRole(varFabricServerRole)
	} else {
		return err
	}

	varFabricServerRole := _FabricServerRole{}

	err = json.Unmarshal(bytes, &varFabricServerRole)
	if err == nil {
		o.FabricPortRole = varFabricServerRole.FabricPortRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectFabricPortRole := reflect.ValueOf(o.FabricPortRole)
		for i := 0; i < reflectFabricPortRole.Type().NumField(); i++ {
			t := reflectFabricPortRole.Type().Field(i)

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

type NullableFabricServerRole struct {
	value *FabricServerRole
	isSet bool
}

func (v NullableFabricServerRole) Get() *FabricServerRole {
	return v.value
}

func (v *NullableFabricServerRole) Set(val *FabricServerRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricServerRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricServerRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricServerRole(val *FabricServerRole) *NullableFabricServerRole {
	return &NullableFabricServerRole{value: val, isSet: true}
}

func (v NullableFabricServerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricServerRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
