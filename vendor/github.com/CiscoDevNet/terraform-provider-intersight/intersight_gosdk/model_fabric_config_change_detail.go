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

// FabricConfigChangeDetail This captures details of configuration changes.
type FabricConfigChangeDetail struct {
	PolicyAbstractConfigChangeDetail
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	Profile              *FabricSwitchProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricConfigChangeDetail FabricConfigChangeDetail

// NewFabricConfigChangeDetail instantiates a new FabricConfigChangeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricConfigChangeDetail(classId string, objectType string) *FabricConfigChangeDetail {
	this := FabricConfigChangeDetail{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configChangeFlag string = "Pending-changes"
	this.ConfigChangeFlag = &configChangeFlag
	var modStatus string = "None"
	this.ModStatus = &modStatus
	return &this
}

// NewFabricConfigChangeDetailWithDefaults instantiates a new FabricConfigChangeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricConfigChangeDetailWithDefaults() *FabricConfigChangeDetail {
	this := FabricConfigChangeDetail{}
	var classId string = "fabric.ConfigChangeDetail"
	this.ClassId = classId
	var objectType string = "fabric.ConfigChangeDetail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricConfigChangeDetail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricConfigChangeDetail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricConfigChangeDetail) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricConfigChangeDetail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricConfigChangeDetail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricConfigChangeDetail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *FabricConfigChangeDetail) GetProfile() FabricSwitchProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret FabricSwitchProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricConfigChangeDetail) GetProfileOk() (*FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *FabricConfigChangeDetail) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given FabricSwitchProfileRelationship and assigns it to the Profile field.
func (o *FabricConfigChangeDetail) SetProfile(v FabricSwitchProfileRelationship) {
	o.Profile = &v
}

func (o FabricConfigChangeDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigChangeDetail, errPolicyAbstractConfigChangeDetail := json.Marshal(o.PolicyAbstractConfigChangeDetail)
	if errPolicyAbstractConfigChangeDetail != nil {
		return []byte{}, errPolicyAbstractConfigChangeDetail
	}
	errPolicyAbstractConfigChangeDetail = json.Unmarshal([]byte(serializedPolicyAbstractConfigChangeDetail), &toSerialize)
	if errPolicyAbstractConfigChangeDetail != nil {
		return []byte{}, errPolicyAbstractConfigChangeDetail
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricConfigChangeDetail) UnmarshalJSON(bytes []byte) (err error) {
	type FabricConfigChangeDetailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                           `json:"ObjectType"`
		Profile    *FabricSwitchProfileRelationship `json:"Profile,omitempty"`
	}

	varFabricConfigChangeDetailWithoutEmbeddedStruct := FabricConfigChangeDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricConfigChangeDetailWithoutEmbeddedStruct)
	if err == nil {
		varFabricConfigChangeDetail := _FabricConfigChangeDetail{}
		varFabricConfigChangeDetail.ClassId = varFabricConfigChangeDetailWithoutEmbeddedStruct.ClassId
		varFabricConfigChangeDetail.ObjectType = varFabricConfigChangeDetailWithoutEmbeddedStruct.ObjectType
		varFabricConfigChangeDetail.Profile = varFabricConfigChangeDetailWithoutEmbeddedStruct.Profile
		*o = FabricConfigChangeDetail(varFabricConfigChangeDetail)
	} else {
		return err
	}

	varFabricConfigChangeDetail := _FabricConfigChangeDetail{}

	err = json.Unmarshal(bytes, &varFabricConfigChangeDetail)
	if err == nil {
		o.PolicyAbstractConfigChangeDetail = varFabricConfigChangeDetail.PolicyAbstractConfigChangeDetail
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Profile")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigChangeDetail := reflect.ValueOf(o.PolicyAbstractConfigChangeDetail)
		for i := 0; i < reflectPolicyAbstractConfigChangeDetail.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigChangeDetail.Type().Field(i)

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

type NullableFabricConfigChangeDetail struct {
	value *FabricConfigChangeDetail
	isSet bool
}

func (v NullableFabricConfigChangeDetail) Get() *FabricConfigChangeDetail {
	return v.value
}

func (v *NullableFabricConfigChangeDetail) Set(val *FabricConfigChangeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricConfigChangeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricConfigChangeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricConfigChangeDetail(val *FabricConfigChangeDetail) *NullableFabricConfigChangeDetail {
	return &NullableFabricConfigChangeDetail{value: val, isSet: true}
}

func (v NullableFabricConfigChangeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricConfigChangeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
