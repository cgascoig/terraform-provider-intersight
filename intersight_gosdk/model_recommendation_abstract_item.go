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
	"reflect"
	"strings"
)

// RecommendationAbstractItem Abstract object representing the recommended physical item.
type RecommendationAbstractItem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the physical device recommended.
	Name *string `json:"Name,omitempty"`
	// The type of physical device recommended. * `Disk` - The Enum value Disk represents that the item type recommended is a Physical Disk. * `Node` - The Enum value Node represents that the item type recommended is a Storage Node. * `Cluster` - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationAbstractItem RecommendationAbstractItem

// NewRecommendationAbstractItem instantiates a new RecommendationAbstractItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationAbstractItem(classId string, objectType string) *RecommendationAbstractItem {
	this := RecommendationAbstractItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationAbstractItemWithDefaults instantiates a new RecommendationAbstractItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationAbstractItemWithDefaults() *RecommendationAbstractItem {
	this := RecommendationAbstractItem{}
	var classId string = "recommendation.PhysicalItem"
	this.ClassId = classId
	var objectType string = "recommendation.PhysicalItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationAbstractItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationAbstractItem) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationAbstractItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationAbstractItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecommendationAbstractItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecommendationAbstractItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecommendationAbstractItem) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RecommendationAbstractItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RecommendationAbstractItem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RecommendationAbstractItem) SetType(v string) {
	o.Type = &v
}

func (o RecommendationAbstractItem) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationAbstractItem) UnmarshalJSON(bytes []byte) (err error) {
	type RecommendationAbstractItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The name of the physical device recommended.
		Name *string `json:"Name,omitempty"`
		// The type of physical device recommended. * `Disk` - The Enum value Disk represents that the item type recommended is a Physical Disk. * `Node` - The Enum value Node represents that the item type recommended is a Storage Node. * `Cluster` - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster.
		Type *string `json:"Type,omitempty"`
	}

	varRecommendationAbstractItemWithoutEmbeddedStruct := RecommendationAbstractItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecommendationAbstractItemWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationAbstractItem := _RecommendationAbstractItem{}
		varRecommendationAbstractItem.ClassId = varRecommendationAbstractItemWithoutEmbeddedStruct.ClassId
		varRecommendationAbstractItem.ObjectType = varRecommendationAbstractItemWithoutEmbeddedStruct.ObjectType
		varRecommendationAbstractItem.Name = varRecommendationAbstractItemWithoutEmbeddedStruct.Name
		varRecommendationAbstractItem.Type = varRecommendationAbstractItemWithoutEmbeddedStruct.Type
		*o = RecommendationAbstractItem(varRecommendationAbstractItem)
	} else {
		return err
	}

	varRecommendationAbstractItem := _RecommendationAbstractItem{}

	err = json.Unmarshal(bytes, &varRecommendationAbstractItem)
	if err == nil {
		o.MoBaseMo = varRecommendationAbstractItem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")

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

type NullableRecommendationAbstractItem struct {
	value *RecommendationAbstractItem
	isSet bool
}

func (v NullableRecommendationAbstractItem) Get() *RecommendationAbstractItem {
	return v.value
}

func (v *NullableRecommendationAbstractItem) Set(val *RecommendationAbstractItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationAbstractItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationAbstractItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationAbstractItem(val *RecommendationAbstractItem) *NullableRecommendationAbstractItem {
	return &NullableRecommendationAbstractItem{value: val, isSet: true}
}

func (v NullableRecommendationAbstractItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationAbstractItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
