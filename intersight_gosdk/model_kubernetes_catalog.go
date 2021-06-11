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

// KubernetesCatalog A catalog to hold the Kubernetes related items such as versions and addons for Intersight Kubernetes Services.
type KubernetesCatalog struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the catalog. The names are populated and predefined during MO creation.
	Name                 *string                               `json:"Name,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	System               *IamSystemRelationship                `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesCatalog KubernetesCatalog

// NewKubernetesCatalog instantiates a new KubernetesCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCatalog(classId string, objectType string) *KubernetesCatalog {
	this := KubernetesCatalog{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesCatalogWithDefaults instantiates a new KubernetesCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCatalogWithDefaults() *KubernetesCatalog {
	this := KubernetesCatalog{}
	var classId string = "kubernetes.Catalog"
	this.ClassId = classId
	var objectType string = "kubernetes.Catalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesCatalog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesCatalog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesCatalog) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesCatalog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesCatalog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesCatalog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesCatalog) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCatalog) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesCatalog) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesCatalog) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesCatalog) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCatalog) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesCatalog) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesCatalog) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *KubernetesCatalog) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCatalog) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *KubernetesCatalog) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *KubernetesCatalog) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o KubernetesCatalog) MarshalJSON() ([]byte, error) {
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
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesCatalog) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesCatalogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the catalog. The names are populated and predefined during MO creation.
		Name         *string                               `json:"Name,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		System       *IamSystemRelationship                `json:"System,omitempty"`
	}

	varKubernetesCatalogWithoutEmbeddedStruct := KubernetesCatalogWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesCatalogWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesCatalog := _KubernetesCatalog{}
		varKubernetesCatalog.ClassId = varKubernetesCatalogWithoutEmbeddedStruct.ClassId
		varKubernetesCatalog.ObjectType = varKubernetesCatalogWithoutEmbeddedStruct.ObjectType
		varKubernetesCatalog.Name = varKubernetesCatalogWithoutEmbeddedStruct.Name
		varKubernetesCatalog.Organization = varKubernetesCatalogWithoutEmbeddedStruct.Organization
		varKubernetesCatalog.System = varKubernetesCatalogWithoutEmbeddedStruct.System
		*o = KubernetesCatalog(varKubernetesCatalog)
	} else {
		return err
	}

	varKubernetesCatalog := _KubernetesCatalog{}

	err = json.Unmarshal(bytes, &varKubernetesCatalog)
	if err == nil {
		o.MoBaseMo = varKubernetesCatalog.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "System")

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

type NullableKubernetesCatalog struct {
	value *KubernetesCatalog
	isSet bool
}

func (v NullableKubernetesCatalog) Get() *KubernetesCatalog {
	return v.value
}

func (v *NullableKubernetesCatalog) Set(val *KubernetesCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCatalog(val *KubernetesCatalog) *NullableKubernetesCatalog {
	return &NullableKubernetesCatalog{value: val, isSet: true}
}

func (v NullableKubernetesCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
