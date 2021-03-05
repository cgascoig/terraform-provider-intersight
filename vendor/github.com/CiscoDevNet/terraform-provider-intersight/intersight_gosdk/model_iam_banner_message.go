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

// IamBannerMessage Configuration for the banner message, including title, contents, and display toggle.
type IamBannerMessage struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Contents of the banner message.
	BannerContents *string `json:"BannerContents,omitempty"`
	// Whether or not to display the banner message.
	BannerDisplay *bool `json:"BannerDisplay,omitempty"`
	// Title of the banner message.
	BannerTitle          *string                 `json:"BannerTitle,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamBannerMessage IamBannerMessage

// NewIamBannerMessage instantiates a new IamBannerMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamBannerMessage(classId string, objectType string) *IamBannerMessage {
	this := IamBannerMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamBannerMessageWithDefaults instantiates a new IamBannerMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamBannerMessageWithDefaults() *IamBannerMessage {
	this := IamBannerMessage{}
	var classId string = "iam.BannerMessage"
	this.ClassId = classId
	var objectType string = "iam.BannerMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamBannerMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamBannerMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamBannerMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamBannerMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBannerContents returns the BannerContents field value if set, zero value otherwise.
func (o *IamBannerMessage) GetBannerContents() string {
	if o == nil || o.BannerContents == nil {
		var ret string
		return ret
	}
	return *o.BannerContents
}

// GetBannerContentsOk returns a tuple with the BannerContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetBannerContentsOk() (*string, bool) {
	if o == nil || o.BannerContents == nil {
		return nil, false
	}
	return o.BannerContents, true
}

// HasBannerContents returns a boolean if a field has been set.
func (o *IamBannerMessage) HasBannerContents() bool {
	if o != nil && o.BannerContents != nil {
		return true
	}

	return false
}

// SetBannerContents gets a reference to the given string and assigns it to the BannerContents field.
func (o *IamBannerMessage) SetBannerContents(v string) {
	o.BannerContents = &v
}

// GetBannerDisplay returns the BannerDisplay field value if set, zero value otherwise.
func (o *IamBannerMessage) GetBannerDisplay() bool {
	if o == nil || o.BannerDisplay == nil {
		var ret bool
		return ret
	}
	return *o.BannerDisplay
}

// GetBannerDisplayOk returns a tuple with the BannerDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetBannerDisplayOk() (*bool, bool) {
	if o == nil || o.BannerDisplay == nil {
		return nil, false
	}
	return o.BannerDisplay, true
}

// HasBannerDisplay returns a boolean if a field has been set.
func (o *IamBannerMessage) HasBannerDisplay() bool {
	if o != nil && o.BannerDisplay != nil {
		return true
	}

	return false
}

// SetBannerDisplay gets a reference to the given bool and assigns it to the BannerDisplay field.
func (o *IamBannerMessage) SetBannerDisplay(v bool) {
	o.BannerDisplay = &v
}

// GetBannerTitle returns the BannerTitle field value if set, zero value otherwise.
func (o *IamBannerMessage) GetBannerTitle() string {
	if o == nil || o.BannerTitle == nil {
		var ret string
		return ret
	}
	return *o.BannerTitle
}

// GetBannerTitleOk returns a tuple with the BannerTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetBannerTitleOk() (*string, bool) {
	if o == nil || o.BannerTitle == nil {
		return nil, false
	}
	return o.BannerTitle, true
}

// HasBannerTitle returns a boolean if a field has been set.
func (o *IamBannerMessage) HasBannerTitle() bool {
	if o != nil && o.BannerTitle != nil {
		return true
	}

	return false
}

// SetBannerTitle gets a reference to the given string and assigns it to the BannerTitle field.
func (o *IamBannerMessage) SetBannerTitle(v string) {
	o.BannerTitle = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamBannerMessage) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamBannerMessage) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamBannerMessage) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamBannerMessage) MarshalJSON() ([]byte, error) {
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
	if o.BannerContents != nil {
		toSerialize["BannerContents"] = o.BannerContents
	}
	if o.BannerDisplay != nil {
		toSerialize["BannerDisplay"] = o.BannerDisplay
	}
	if o.BannerTitle != nil {
		toSerialize["BannerTitle"] = o.BannerTitle
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamBannerMessage) UnmarshalJSON(bytes []byte) (err error) {
	type IamBannerMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Contents of the banner message.
		BannerContents *string `json:"BannerContents,omitempty"`
		// Whether or not to display the banner message.
		BannerDisplay *bool `json:"BannerDisplay,omitempty"`
		// Title of the banner message.
		BannerTitle *string                 `json:"BannerTitle,omitempty"`
		Account     *IamAccountRelationship `json:"Account,omitempty"`
	}

	varIamBannerMessageWithoutEmbeddedStruct := IamBannerMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamBannerMessageWithoutEmbeddedStruct)
	if err == nil {
		varIamBannerMessage := _IamBannerMessage{}
		varIamBannerMessage.ClassId = varIamBannerMessageWithoutEmbeddedStruct.ClassId
		varIamBannerMessage.ObjectType = varIamBannerMessageWithoutEmbeddedStruct.ObjectType
		varIamBannerMessage.BannerContents = varIamBannerMessageWithoutEmbeddedStruct.BannerContents
		varIamBannerMessage.BannerDisplay = varIamBannerMessageWithoutEmbeddedStruct.BannerDisplay
		varIamBannerMessage.BannerTitle = varIamBannerMessageWithoutEmbeddedStruct.BannerTitle
		varIamBannerMessage.Account = varIamBannerMessageWithoutEmbeddedStruct.Account
		*o = IamBannerMessage(varIamBannerMessage)
	} else {
		return err
	}

	varIamBannerMessage := _IamBannerMessage{}

	err = json.Unmarshal(bytes, &varIamBannerMessage)
	if err == nil {
		o.MoBaseMo = varIamBannerMessage.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BannerContents")
		delete(additionalProperties, "BannerDisplay")
		delete(additionalProperties, "BannerTitle")
		delete(additionalProperties, "Account")

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

type NullableIamBannerMessage struct {
	value *IamBannerMessage
	isSet bool
}

func (v NullableIamBannerMessage) Get() *IamBannerMessage {
	return v.value
}

func (v *NullableIamBannerMessage) Set(val *IamBannerMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableIamBannerMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableIamBannerMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamBannerMessage(val *IamBannerMessage) *NullableIamBannerMessage {
	return &NullableIamBannerMessage{value: val, isSet: true}
}

func (v NullableIamBannerMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamBannerMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
