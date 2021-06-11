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
	"reflect"
	"strings"
)

// ServerProfileTemplate A profile template specifying configuration settings for a physical server.
type ServerProfileTemplate struct {
	ServerBaseProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The count of the server profiles derived from the template.
	Usage                *int64                                `json:"Usage,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerProfileTemplate ServerProfileTemplate

// NewServerProfileTemplate instantiates a new ServerProfileTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProfileTemplate(classId string, objectType string) *ServerProfileTemplate {
	this := ServerProfileTemplate{}
	this.ClassId = classId
	this.ObjectType = objectType
	var usage int64 = 0
	this.Usage = &usage
	return &this
}

// NewServerProfileTemplateWithDefaults instantiates a new ServerProfileTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerProfileTemplateWithDefaults() *ServerProfileTemplate {
	this := ServerProfileTemplate{}
	var classId string = "server.ProfileTemplate"
	this.ClassId = classId
	var objectType string = "server.ProfileTemplate"
	this.ObjectType = objectType
	var usage int64 = 0
	this.Usage = &usage
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerProfileTemplate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerProfileTemplate) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerProfileTemplate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerProfileTemplate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ServerProfileTemplate) GetUsage() int64 {
	if o == nil || o.Usage == nil {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplate) GetUsageOk() (*int64, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ServerProfileTemplate) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *ServerProfileTemplate) SetUsage(v int64) {
	o.Usage = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ServerProfileTemplate) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplate) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServerProfileTemplate) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ServerProfileTemplate) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ServerProfileTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedServerBaseProfile, errServerBaseProfile := json.Marshal(o.ServerBaseProfile)
	if errServerBaseProfile != nil {
		return []byte{}, errServerBaseProfile
	}
	errServerBaseProfile = json.Unmarshal([]byte(serializedServerBaseProfile), &toSerialize)
	if errServerBaseProfile != nil {
		return []byte{}, errServerBaseProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Usage != nil {
		toSerialize["Usage"] = o.Usage
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerProfileTemplate) UnmarshalJSON(bytes []byte) (err error) {
	type ServerProfileTemplateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The count of the server profiles derived from the template.
		Usage        *int64                                `json:"Usage,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varServerProfileTemplateWithoutEmbeddedStruct := ServerProfileTemplateWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varServerProfileTemplateWithoutEmbeddedStruct)
	if err == nil {
		varServerProfileTemplate := _ServerProfileTemplate{}
		varServerProfileTemplate.ClassId = varServerProfileTemplateWithoutEmbeddedStruct.ClassId
		varServerProfileTemplate.ObjectType = varServerProfileTemplateWithoutEmbeddedStruct.ObjectType
		varServerProfileTemplate.Usage = varServerProfileTemplateWithoutEmbeddedStruct.Usage
		varServerProfileTemplate.Organization = varServerProfileTemplateWithoutEmbeddedStruct.Organization
		*o = ServerProfileTemplate(varServerProfileTemplate)
	} else {
		return err
	}

	varServerProfileTemplate := _ServerProfileTemplate{}

	err = json.Unmarshal(bytes, &varServerProfileTemplate)
	if err == nil {
		o.ServerBaseProfile = varServerProfileTemplate.ServerBaseProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectServerBaseProfile := reflect.ValueOf(o.ServerBaseProfile)
		for i := 0; i < reflectServerBaseProfile.Type().NumField(); i++ {
			t := reflectServerBaseProfile.Type().Field(i)

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

type NullableServerProfileTemplate struct {
	value *ServerProfileTemplate
	isSet bool
}

func (v NullableServerProfileTemplate) Get() *ServerProfileTemplate {
	return v.value
}

func (v *NullableServerProfileTemplate) Set(val *ServerProfileTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProfileTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProfileTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProfileTemplate(val *ServerProfileTemplate) *NullableServerProfileTemplate {
	return &NullableServerProfileTemplate{value: val, isSet: true}
}

func (v NullableServerProfileTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProfileTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
