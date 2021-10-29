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

// CloudAwsKeyPair Key pair object in AWS inventory. A key pair, consisting of a private key and a public key, is a set of security credentials that are used to prove identity when connecting to an EC2 instance.
type CloudAwsKeyPair struct {
	CloudBaseEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Either the SHA-1 digest of the DER encoded private key or  MD5 public key fingerprint.
	FingerPrint *string `json:"FingerPrint,omitempty"`
	// Used in authenticating to the virtual machine .
	PublicKey            *string                          `json:"PublicKey,omitempty"`
	AwsBillingUnit       *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsKeyPair CloudAwsKeyPair

// NewCloudAwsKeyPair instantiates a new CloudAwsKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsKeyPair(classId string, objectType string) *CloudAwsKeyPair {
	this := CloudAwsKeyPair{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsKeyPairWithDefaults instantiates a new CloudAwsKeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsKeyPairWithDefaults() *CloudAwsKeyPair {
	this := CloudAwsKeyPair{}
	var classId string = "cloud.AwsKeyPair"
	this.ClassId = classId
	var objectType string = "cloud.AwsKeyPair"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsKeyPair) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsKeyPair) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsKeyPair) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsKeyPair) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsKeyPair) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsKeyPair) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFingerPrint returns the FingerPrint field value if set, zero value otherwise.
func (o *CloudAwsKeyPair) GetFingerPrint() string {
	if o == nil || o.FingerPrint == nil {
		var ret string
		return ret
	}
	return *o.FingerPrint
}

// GetFingerPrintOk returns a tuple with the FingerPrint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsKeyPair) GetFingerPrintOk() (*string, bool) {
	if o == nil || o.FingerPrint == nil {
		return nil, false
	}
	return o.FingerPrint, true
}

// HasFingerPrint returns a boolean if a field has been set.
func (o *CloudAwsKeyPair) HasFingerPrint() bool {
	if o != nil && o.FingerPrint != nil {
		return true
	}

	return false
}

// SetFingerPrint gets a reference to the given string and assigns it to the FingerPrint field.
func (o *CloudAwsKeyPair) SetFingerPrint(v string) {
	o.FingerPrint = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CloudAwsKeyPair) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsKeyPair) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CloudAwsKeyPair) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CloudAwsKeyPair) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetAwsBillingUnit returns the AwsBillingUnit field value if set, zero value otherwise.
func (o *CloudAwsKeyPair) GetAwsBillingUnit() CloudAwsBillingUnitRelationship {
	if o == nil || o.AwsBillingUnit == nil {
		var ret CloudAwsBillingUnitRelationship
		return ret
	}
	return *o.AwsBillingUnit
}

// GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsKeyPair) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool) {
	if o == nil || o.AwsBillingUnit == nil {
		return nil, false
	}
	return o.AwsBillingUnit, true
}

// HasAwsBillingUnit returns a boolean if a field has been set.
func (o *CloudAwsKeyPair) HasAwsBillingUnit() bool {
	if o != nil && o.AwsBillingUnit != nil {
		return true
	}

	return false
}

// SetAwsBillingUnit gets a reference to the given CloudAwsBillingUnitRelationship and assigns it to the AwsBillingUnit field.
func (o *CloudAwsKeyPair) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship) {
	o.AwsBillingUnit = &v
}

func (o CloudAwsKeyPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseEntity, errCloudBaseEntity := json.Marshal(o.CloudBaseEntity)
	if errCloudBaseEntity != nil {
		return []byte{}, errCloudBaseEntity
	}
	errCloudBaseEntity = json.Unmarshal([]byte(serializedCloudBaseEntity), &toSerialize)
	if errCloudBaseEntity != nil {
		return []byte{}, errCloudBaseEntity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FingerPrint != nil {
		toSerialize["FingerPrint"] = o.FingerPrint
	}
	if o.PublicKey != nil {
		toSerialize["PublicKey"] = o.PublicKey
	}
	if o.AwsBillingUnit != nil {
		toSerialize["AwsBillingUnit"] = o.AwsBillingUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsKeyPair) UnmarshalJSON(bytes []byte) (err error) {
	type CloudAwsKeyPairWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Either the SHA-1 digest of the DER encoded private key or  MD5 public key fingerprint.
		FingerPrint *string `json:"FingerPrint,omitempty"`
		// Used in authenticating to the virtual machine .
		PublicKey      *string                          `json:"PublicKey,omitempty"`
		AwsBillingUnit *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	}

	varCloudAwsKeyPairWithoutEmbeddedStruct := CloudAwsKeyPairWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudAwsKeyPairWithoutEmbeddedStruct)
	if err == nil {
		varCloudAwsKeyPair := _CloudAwsKeyPair{}
		varCloudAwsKeyPair.ClassId = varCloudAwsKeyPairWithoutEmbeddedStruct.ClassId
		varCloudAwsKeyPair.ObjectType = varCloudAwsKeyPairWithoutEmbeddedStruct.ObjectType
		varCloudAwsKeyPair.FingerPrint = varCloudAwsKeyPairWithoutEmbeddedStruct.FingerPrint
		varCloudAwsKeyPair.PublicKey = varCloudAwsKeyPairWithoutEmbeddedStruct.PublicKey
		varCloudAwsKeyPair.AwsBillingUnit = varCloudAwsKeyPairWithoutEmbeddedStruct.AwsBillingUnit
		*o = CloudAwsKeyPair(varCloudAwsKeyPair)
	} else {
		return err
	}

	varCloudAwsKeyPair := _CloudAwsKeyPair{}

	err = json.Unmarshal(bytes, &varCloudAwsKeyPair)
	if err == nil {
		o.CloudBaseEntity = varCloudAwsKeyPair.CloudBaseEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FingerPrint")
		delete(additionalProperties, "PublicKey")
		delete(additionalProperties, "AwsBillingUnit")

		// remove fields from embedded structs
		reflectCloudBaseEntity := reflect.ValueOf(o.CloudBaseEntity)
		for i := 0; i < reflectCloudBaseEntity.Type().NumField(); i++ {
			t := reflectCloudBaseEntity.Type().Field(i)

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

type NullableCloudAwsKeyPair struct {
	value *CloudAwsKeyPair
	isSet bool
}

func (v NullableCloudAwsKeyPair) Get() *CloudAwsKeyPair {
	return v.value
}

func (v *NullableCloudAwsKeyPair) Set(val *CloudAwsKeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsKeyPair(val *CloudAwsKeyPair) *NullableCloudAwsKeyPair {
	return &NullableCloudAwsKeyPair{value: val, isSet: true}
}

func (v NullableCloudAwsKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
