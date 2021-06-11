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

// AssetCustomerInformation Type for saving contract information.
type AssetCustomerInformation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                          `json:"ObjectType"`
	Address    NullableAssetAddressInformation `json:"Address,omitempty"`
	// Unique identifier for an end customer. This identifier is allocated by Cisco.
	Id *string `json:"Id,omitempty"`
	// Name as per the information provided by the user.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetCustomerInformation AssetCustomerInformation

// NewAssetCustomerInformation instantiates a new AssetCustomerInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCustomerInformation(classId string, objectType string) *AssetCustomerInformation {
	this := AssetCustomerInformation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetCustomerInformationWithDefaults instantiates a new AssetCustomerInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCustomerInformationWithDefaults() *AssetCustomerInformation {
	this := AssetCustomerInformation{}
	var classId string = "asset.CustomerInformation"
	this.ClassId = classId
	var objectType string = "asset.CustomerInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetCustomerInformation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetCustomerInformation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetCustomerInformation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetCustomerInformation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetCustomerInformation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetCustomerInformation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetCustomerInformation) GetAddress() AssetAddressInformation {
	if o == nil || o.Address.Get() == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetCustomerInformation) GetAddressOk() (*AssetAddressInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *AssetCustomerInformation) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableAssetAddressInformation and assigns it to the Address field.
func (o *AssetCustomerInformation) SetAddress(v AssetAddressInformation) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *AssetCustomerInformation) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *AssetCustomerInformation) UnsetAddress() {
	o.Address.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetCustomerInformation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCustomerInformation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetCustomerInformation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetCustomerInformation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetCustomerInformation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCustomerInformation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetCustomerInformation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetCustomerInformation) SetName(v string) {
	o.Name = &v
}

func (o AssetCustomerInformation) MarshalJSON() ([]byte, error) {
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
	if o.Address.IsSet() {
		toSerialize["Address"] = o.Address.Get()
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetCustomerInformation) UnmarshalJSON(bytes []byte) (err error) {
	type AssetCustomerInformationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                          `json:"ObjectType"`
		Address    NullableAssetAddressInformation `json:"Address,omitempty"`
		// Unique identifier for an end customer. This identifier is allocated by Cisco.
		Id *string `json:"Id,omitempty"`
		// Name as per the information provided by the user.
		Name *string `json:"Name,omitempty"`
	}

	varAssetCustomerInformationWithoutEmbeddedStruct := AssetCustomerInformationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetCustomerInformationWithoutEmbeddedStruct)
	if err == nil {
		varAssetCustomerInformation := _AssetCustomerInformation{}
		varAssetCustomerInformation.ClassId = varAssetCustomerInformationWithoutEmbeddedStruct.ClassId
		varAssetCustomerInformation.ObjectType = varAssetCustomerInformationWithoutEmbeddedStruct.ObjectType
		varAssetCustomerInformation.Address = varAssetCustomerInformationWithoutEmbeddedStruct.Address
		varAssetCustomerInformation.Id = varAssetCustomerInformationWithoutEmbeddedStruct.Id
		varAssetCustomerInformation.Name = varAssetCustomerInformationWithoutEmbeddedStruct.Name
		*o = AssetCustomerInformation(varAssetCustomerInformation)
	} else {
		return err
	}

	varAssetCustomerInformation := _AssetCustomerInformation{}

	err = json.Unmarshal(bytes, &varAssetCustomerInformation)
	if err == nil {
		o.MoBaseComplexType = varAssetCustomerInformation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")

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

type NullableAssetCustomerInformation struct {
	value *AssetCustomerInformation
	isSet bool
}

func (v NullableAssetCustomerInformation) Get() *AssetCustomerInformation {
	return v.value
}

func (v *NullableAssetCustomerInformation) Set(val *AssetCustomerInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCustomerInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCustomerInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCustomerInformation(val *AssetCustomerInformation) *NullableAssetCustomerInformation {
	return &NullableAssetCustomerInformation{value: val, isSet: true}
}

func (v NullableAssetCustomerInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCustomerInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
