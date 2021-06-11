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

// AssetTerraformIntegrationTerraformAgentOptions Cloud options for Terrraform Agent platform type.
type AssetTerraformIntegrationTerraformAgentOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string   `json:"ObjectType"`
	ManagedHosts []string `json:"ManagedHosts,omitempty"`
	// Agent pool name for Terraform Agent platform type.
	TerraformAgentPoolName *string `json:"TerraformAgentPoolName,omitempty"`
	// Organization for Terraform Agent platform type.
	TerraformOrganization *string `json:"TerraformOrganization,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _AssetTerraformIntegrationTerraformAgentOptions AssetTerraformIntegrationTerraformAgentOptions

// NewAssetTerraformIntegrationTerraformAgentOptions instantiates a new AssetTerraformIntegrationTerraformAgentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTerraformIntegrationTerraformAgentOptions(classId string, objectType string) *AssetTerraformIntegrationTerraformAgentOptions {
	this := AssetTerraformIntegrationTerraformAgentOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetTerraformIntegrationTerraformAgentOptionsWithDefaults instantiates a new AssetTerraformIntegrationTerraformAgentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTerraformIntegrationTerraformAgentOptionsWithDefaults() *AssetTerraformIntegrationTerraformAgentOptions {
	this := AssetTerraformIntegrationTerraformAgentOptions{}
	var classId string = "asset.TerraformIntegrationTerraformAgentOptions"
	this.ClassId = classId
	var objectType string = "asset.TerraformIntegrationTerraformAgentOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTerraformIntegrationTerraformAgentOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTerraformIntegrationTerraformAgentOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetManagedHosts returns the ManagedHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetManagedHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ManagedHosts
}

// GetManagedHostsOk returns a tuple with the ManagedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetManagedHostsOk() (*[]string, bool) {
	if o == nil || o.ManagedHosts == nil {
		return nil, false
	}
	return &o.ManagedHosts, true
}

// HasManagedHosts returns a boolean if a field has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptions) HasManagedHosts() bool {
	if o != nil && o.ManagedHosts != nil {
		return true
	}

	return false
}

// SetManagedHosts gets a reference to the given []string and assigns it to the ManagedHosts field.
func (o *AssetTerraformIntegrationTerraformAgentOptions) SetManagedHosts(v []string) {
	o.ManagedHosts = v
}

// GetTerraformAgentPoolName returns the TerraformAgentPoolName field value if set, zero value otherwise.
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetTerraformAgentPoolName() string {
	if o == nil || o.TerraformAgentPoolName == nil {
		var ret string
		return ret
	}
	return *o.TerraformAgentPoolName
}

// GetTerraformAgentPoolNameOk returns a tuple with the TerraformAgentPoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetTerraformAgentPoolNameOk() (*string, bool) {
	if o == nil || o.TerraformAgentPoolName == nil {
		return nil, false
	}
	return o.TerraformAgentPoolName, true
}

// HasTerraformAgentPoolName returns a boolean if a field has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptions) HasTerraformAgentPoolName() bool {
	if o != nil && o.TerraformAgentPoolName != nil {
		return true
	}

	return false
}

// SetTerraformAgentPoolName gets a reference to the given string and assigns it to the TerraformAgentPoolName field.
func (o *AssetTerraformIntegrationTerraformAgentOptions) SetTerraformAgentPoolName(v string) {
	o.TerraformAgentPoolName = &v
}

// GetTerraformOrganization returns the TerraformOrganization field value if set, zero value otherwise.
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetTerraformOrganization() string {
	if o == nil || o.TerraformOrganization == nil {
		var ret string
		return ret
	}
	return *o.TerraformOrganization
}

// GetTerraformOrganizationOk returns a tuple with the TerraformOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptions) GetTerraformOrganizationOk() (*string, bool) {
	if o == nil || o.TerraformOrganization == nil {
		return nil, false
	}
	return o.TerraformOrganization, true
}

// HasTerraformOrganization returns a boolean if a field has been set.
func (o *AssetTerraformIntegrationTerraformAgentOptions) HasTerraformOrganization() bool {
	if o != nil && o.TerraformOrganization != nil {
		return true
	}

	return false
}

// SetTerraformOrganization gets a reference to the given string and assigns it to the TerraformOrganization field.
func (o *AssetTerraformIntegrationTerraformAgentOptions) SetTerraformOrganization(v string) {
	o.TerraformOrganization = &v
}

func (o AssetTerraformIntegrationTerraformAgentOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ManagedHosts != nil {
		toSerialize["ManagedHosts"] = o.ManagedHosts
	}
	if o.TerraformAgentPoolName != nil {
		toSerialize["TerraformAgentPoolName"] = o.TerraformAgentPoolName
	}
	if o.TerraformOrganization != nil {
		toSerialize["TerraformOrganization"] = o.TerraformOrganization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTerraformIntegrationTerraformAgentOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string   `json:"ObjectType"`
		ManagedHosts []string `json:"ManagedHosts,omitempty"`
		// Agent pool name for Terraform Agent platform type.
		TerraformAgentPoolName *string `json:"TerraformAgentPoolName,omitempty"`
		// Organization for Terraform Agent platform type.
		TerraformOrganization *string `json:"TerraformOrganization,omitempty"`
	}

	varAssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct := AssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetTerraformIntegrationTerraformAgentOptions := _AssetTerraformIntegrationTerraformAgentOptions{}
		varAssetTerraformIntegrationTerraformAgentOptions.ClassId = varAssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct.ClassId
		varAssetTerraformIntegrationTerraformAgentOptions.ObjectType = varAssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct.ObjectType
		varAssetTerraformIntegrationTerraformAgentOptions.ManagedHosts = varAssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct.ManagedHosts
		varAssetTerraformIntegrationTerraformAgentOptions.TerraformAgentPoolName = varAssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct.TerraformAgentPoolName
		varAssetTerraformIntegrationTerraformAgentOptions.TerraformOrganization = varAssetTerraformIntegrationTerraformAgentOptionsWithoutEmbeddedStruct.TerraformOrganization
		*o = AssetTerraformIntegrationTerraformAgentOptions(varAssetTerraformIntegrationTerraformAgentOptions)
	} else {
		return err
	}

	varAssetTerraformIntegrationTerraformAgentOptions := _AssetTerraformIntegrationTerraformAgentOptions{}

	err = json.Unmarshal(bytes, &varAssetTerraformIntegrationTerraformAgentOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetTerraformIntegrationTerraformAgentOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ManagedHosts")
		delete(additionalProperties, "TerraformAgentPoolName")
		delete(additionalProperties, "TerraformOrganization")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetTerraformIntegrationTerraformAgentOptions struct {
	value *AssetTerraformIntegrationTerraformAgentOptions
	isSet bool
}

func (v NullableAssetTerraformIntegrationTerraformAgentOptions) Get() *AssetTerraformIntegrationTerraformAgentOptions {
	return v.value
}

func (v *NullableAssetTerraformIntegrationTerraformAgentOptions) Set(val *AssetTerraformIntegrationTerraformAgentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTerraformIntegrationTerraformAgentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTerraformIntegrationTerraformAgentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTerraformIntegrationTerraformAgentOptions(val *AssetTerraformIntegrationTerraformAgentOptions) *NullableAssetTerraformIntegrationTerraformAgentOptions {
	return &NullableAssetTerraformIntegrationTerraformAgentOptions{value: val, isSet: true}
}

func (v NullableAssetTerraformIntegrationTerraformAgentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTerraformIntegrationTerraformAgentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
