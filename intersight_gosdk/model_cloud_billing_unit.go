/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// CloudBillingUnit Details of the paying account.
type CloudBillingUnit struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The ID of the paying account.
	BillingId *string `json:"BillingId,omitempty"`
	// The name of the paying account.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBillingUnit CloudBillingUnit

// NewCloudBillingUnit instantiates a new CloudBillingUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBillingUnit(classId string, objectType string) *CloudBillingUnit {
	this := CloudBillingUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudBillingUnitWithDefaults instantiates a new CloudBillingUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBillingUnitWithDefaults() *CloudBillingUnit {
	this := CloudBillingUnit{}
	var classId string = "cloud.BillingUnit"
	this.ClassId = classId
	var objectType string = "cloud.BillingUnit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBillingUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBillingUnit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBillingUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudBillingUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBillingUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBillingUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingId returns the BillingId field value if set, zero value otherwise.
func (o *CloudBillingUnit) GetBillingId() string {
	if o == nil || o.BillingId == nil {
		var ret string
		return ret
	}
	return *o.BillingId
}

// GetBillingIdOk returns a tuple with the BillingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBillingUnit) GetBillingIdOk() (*string, bool) {
	if o == nil || o.BillingId == nil {
		return nil, false
	}
	return o.BillingId, true
}

// HasBillingId returns a boolean if a field has been set.
func (o *CloudBillingUnit) HasBillingId() bool {
	if o != nil && o.BillingId != nil {
		return true
	}

	return false
}

// SetBillingId gets a reference to the given string and assigns it to the BillingId field.
func (o *CloudBillingUnit) SetBillingId(v string) {
	o.BillingId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudBillingUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBillingUnit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudBillingUnit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudBillingUnit) SetName(v string) {
	o.Name = &v
}

func (o CloudBillingUnit) MarshalJSON() ([]byte, error) {
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
	if o.BillingId != nil {
		toSerialize["BillingId"] = o.BillingId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudBillingUnit) UnmarshalJSON(bytes []byte) (err error) {
	type CloudBillingUnitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The ID of the paying account.
		BillingId *string `json:"BillingId,omitempty"`
		// The name of the paying account.
		Name *string `json:"Name,omitempty"`
	}

	varCloudBillingUnitWithoutEmbeddedStruct := CloudBillingUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudBillingUnitWithoutEmbeddedStruct)
	if err == nil {
		varCloudBillingUnit := _CloudBillingUnit{}
		varCloudBillingUnit.ClassId = varCloudBillingUnitWithoutEmbeddedStruct.ClassId
		varCloudBillingUnit.ObjectType = varCloudBillingUnitWithoutEmbeddedStruct.ObjectType
		varCloudBillingUnit.BillingId = varCloudBillingUnitWithoutEmbeddedStruct.BillingId
		varCloudBillingUnit.Name = varCloudBillingUnitWithoutEmbeddedStruct.Name
		*o = CloudBillingUnit(varCloudBillingUnit)
	} else {
		return err
	}

	varCloudBillingUnit := _CloudBillingUnit{}

	err = json.Unmarshal(bytes, &varCloudBillingUnit)
	if err == nil {
		o.MoBaseComplexType = varCloudBillingUnit.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingId")
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

type NullableCloudBillingUnit struct {
	value *CloudBillingUnit
	isSet bool
}

func (v NullableCloudBillingUnit) Get() *CloudBillingUnit {
	return v.value
}

func (v *NullableCloudBillingUnit) Set(val *CloudBillingUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBillingUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBillingUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBillingUnit(val *CloudBillingUnit) *NullableCloudBillingUnit {
	return &NullableCloudBillingUnit{value: val, isSet: true}
}

func (v NullableCloudBillingUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBillingUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
