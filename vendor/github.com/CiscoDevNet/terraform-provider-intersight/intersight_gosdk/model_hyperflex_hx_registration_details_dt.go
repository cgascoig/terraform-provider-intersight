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

// HyperflexHxRegistrationDetailsDt License Registration Details of the cluster
type HyperflexHxRegistrationDetailsDt struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Initial Registration Date
	InitialRegistrationDate *string `json:"InitialRegistrationDate,omitempty"`
	// Last License Renewal attempted Date
	LastRenewalAttemptDate *string `json:"LastRenewalAttemptDate,omitempty"`
	// Next Attempt Date for License Renewal
	NextRenewalAttemptDate *string `json:"NextRenewalAttemptDate,omitempty"`
	// Out of compliance start date
	OutOfComplianceStartDate *string `json:"OutOfComplianceStartDate,omitempty"`
	// Date when the registration will expire
	RegistrationExpireDate *string `json:"RegistrationExpireDate,omitempty"`
	// License registration success or failure reason
	RegistrationFailedReason *string `json:"RegistrationFailedReason,omitempty"`
	// Smart Account mapped to cluster
	SmartAccount *string `json:"SmartAccount,omitempty"`
	// Registration Status
	Status *string `json:"Status,omitempty"`
	// Virtual Account mapped to cluster
	VirtualAccount       *string `json:"VirtualAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxRegistrationDetailsDt HyperflexHxRegistrationDetailsDt

// NewHyperflexHxRegistrationDetailsDt instantiates a new HyperflexHxRegistrationDetailsDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxRegistrationDetailsDt(classId string, objectType string) *HyperflexHxRegistrationDetailsDt {
	this := HyperflexHxRegistrationDetailsDt{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxRegistrationDetailsDtWithDefaults instantiates a new HyperflexHxRegistrationDetailsDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxRegistrationDetailsDtWithDefaults() *HyperflexHxRegistrationDetailsDt {
	this := HyperflexHxRegistrationDetailsDt{}
	var classId string = "hyperflex.HxRegistrationDetailsDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxRegistrationDetailsDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxRegistrationDetailsDt) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxRegistrationDetailsDt) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxRegistrationDetailsDt) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxRegistrationDetailsDt) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInitialRegistrationDate returns the InitialRegistrationDate field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetInitialRegistrationDate() string {
	if o == nil || o.InitialRegistrationDate == nil {
		var ret string
		return ret
	}
	return *o.InitialRegistrationDate
}

// GetInitialRegistrationDateOk returns a tuple with the InitialRegistrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetInitialRegistrationDateOk() (*string, bool) {
	if o == nil || o.InitialRegistrationDate == nil {
		return nil, false
	}
	return o.InitialRegistrationDate, true
}

// HasInitialRegistrationDate returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasInitialRegistrationDate() bool {
	if o != nil && o.InitialRegistrationDate != nil {
		return true
	}

	return false
}

// SetInitialRegistrationDate gets a reference to the given string and assigns it to the InitialRegistrationDate field.
func (o *HyperflexHxRegistrationDetailsDt) SetInitialRegistrationDate(v string) {
	o.InitialRegistrationDate = &v
}

// GetLastRenewalAttemptDate returns the LastRenewalAttemptDate field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetLastRenewalAttemptDate() string {
	if o == nil || o.LastRenewalAttemptDate == nil {
		var ret string
		return ret
	}
	return *o.LastRenewalAttemptDate
}

// GetLastRenewalAttemptDateOk returns a tuple with the LastRenewalAttemptDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetLastRenewalAttemptDateOk() (*string, bool) {
	if o == nil || o.LastRenewalAttemptDate == nil {
		return nil, false
	}
	return o.LastRenewalAttemptDate, true
}

// HasLastRenewalAttemptDate returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasLastRenewalAttemptDate() bool {
	if o != nil && o.LastRenewalAttemptDate != nil {
		return true
	}

	return false
}

// SetLastRenewalAttemptDate gets a reference to the given string and assigns it to the LastRenewalAttemptDate field.
func (o *HyperflexHxRegistrationDetailsDt) SetLastRenewalAttemptDate(v string) {
	o.LastRenewalAttemptDate = &v
}

// GetNextRenewalAttemptDate returns the NextRenewalAttemptDate field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetNextRenewalAttemptDate() string {
	if o == nil || o.NextRenewalAttemptDate == nil {
		var ret string
		return ret
	}
	return *o.NextRenewalAttemptDate
}

// GetNextRenewalAttemptDateOk returns a tuple with the NextRenewalAttemptDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetNextRenewalAttemptDateOk() (*string, bool) {
	if o == nil || o.NextRenewalAttemptDate == nil {
		return nil, false
	}
	return o.NextRenewalAttemptDate, true
}

// HasNextRenewalAttemptDate returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasNextRenewalAttemptDate() bool {
	if o != nil && o.NextRenewalAttemptDate != nil {
		return true
	}

	return false
}

// SetNextRenewalAttemptDate gets a reference to the given string and assigns it to the NextRenewalAttemptDate field.
func (o *HyperflexHxRegistrationDetailsDt) SetNextRenewalAttemptDate(v string) {
	o.NextRenewalAttemptDate = &v
}

// GetOutOfComplianceStartDate returns the OutOfComplianceStartDate field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetOutOfComplianceStartDate() string {
	if o == nil || o.OutOfComplianceStartDate == nil {
		var ret string
		return ret
	}
	return *o.OutOfComplianceStartDate
}

// GetOutOfComplianceStartDateOk returns a tuple with the OutOfComplianceStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetOutOfComplianceStartDateOk() (*string, bool) {
	if o == nil || o.OutOfComplianceStartDate == nil {
		return nil, false
	}
	return o.OutOfComplianceStartDate, true
}

// HasOutOfComplianceStartDate returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasOutOfComplianceStartDate() bool {
	if o != nil && o.OutOfComplianceStartDate != nil {
		return true
	}

	return false
}

// SetOutOfComplianceStartDate gets a reference to the given string and assigns it to the OutOfComplianceStartDate field.
func (o *HyperflexHxRegistrationDetailsDt) SetOutOfComplianceStartDate(v string) {
	o.OutOfComplianceStartDate = &v
}

// GetRegistrationExpireDate returns the RegistrationExpireDate field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationExpireDate() string {
	if o == nil || o.RegistrationExpireDate == nil {
		var ret string
		return ret
	}
	return *o.RegistrationExpireDate
}

// GetRegistrationExpireDateOk returns a tuple with the RegistrationExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationExpireDateOk() (*string, bool) {
	if o == nil || o.RegistrationExpireDate == nil {
		return nil, false
	}
	return o.RegistrationExpireDate, true
}

// HasRegistrationExpireDate returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasRegistrationExpireDate() bool {
	if o != nil && o.RegistrationExpireDate != nil {
		return true
	}

	return false
}

// SetRegistrationExpireDate gets a reference to the given string and assigns it to the RegistrationExpireDate field.
func (o *HyperflexHxRegistrationDetailsDt) SetRegistrationExpireDate(v string) {
	o.RegistrationExpireDate = &v
}

// GetRegistrationFailedReason returns the RegistrationFailedReason field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationFailedReason() string {
	if o == nil || o.RegistrationFailedReason == nil {
		var ret string
		return ret
	}
	return *o.RegistrationFailedReason
}

// GetRegistrationFailedReasonOk returns a tuple with the RegistrationFailedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetRegistrationFailedReasonOk() (*string, bool) {
	if o == nil || o.RegistrationFailedReason == nil {
		return nil, false
	}
	return o.RegistrationFailedReason, true
}

// HasRegistrationFailedReason returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasRegistrationFailedReason() bool {
	if o != nil && o.RegistrationFailedReason != nil {
		return true
	}

	return false
}

// SetRegistrationFailedReason gets a reference to the given string and assigns it to the RegistrationFailedReason field.
func (o *HyperflexHxRegistrationDetailsDt) SetRegistrationFailedReason(v string) {
	o.RegistrationFailedReason = &v
}

// GetSmartAccount returns the SmartAccount field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetSmartAccount() string {
	if o == nil || o.SmartAccount == nil {
		var ret string
		return ret
	}
	return *o.SmartAccount
}

// GetSmartAccountOk returns a tuple with the SmartAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetSmartAccountOk() (*string, bool) {
	if o == nil || o.SmartAccount == nil {
		return nil, false
	}
	return o.SmartAccount, true
}

// HasSmartAccount returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasSmartAccount() bool {
	if o != nil && o.SmartAccount != nil {
		return true
	}

	return false
}

// SetSmartAccount gets a reference to the given string and assigns it to the SmartAccount field.
func (o *HyperflexHxRegistrationDetailsDt) SetSmartAccount(v string) {
	o.SmartAccount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexHxRegistrationDetailsDt) SetStatus(v string) {
	o.Status = &v
}

// GetVirtualAccount returns the VirtualAccount field value if set, zero value otherwise.
func (o *HyperflexHxRegistrationDetailsDt) GetVirtualAccount() string {
	if o == nil || o.VirtualAccount == nil {
		var ret string
		return ret
	}
	return *o.VirtualAccount
}

// GetVirtualAccountOk returns a tuple with the VirtualAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxRegistrationDetailsDt) GetVirtualAccountOk() (*string, bool) {
	if o == nil || o.VirtualAccount == nil {
		return nil, false
	}
	return o.VirtualAccount, true
}

// HasVirtualAccount returns a boolean if a field has been set.
func (o *HyperflexHxRegistrationDetailsDt) HasVirtualAccount() bool {
	if o != nil && o.VirtualAccount != nil {
		return true
	}

	return false
}

// SetVirtualAccount gets a reference to the given string and assigns it to the VirtualAccount field.
func (o *HyperflexHxRegistrationDetailsDt) SetVirtualAccount(v string) {
	o.VirtualAccount = &v
}

func (o HyperflexHxRegistrationDetailsDt) MarshalJSON() ([]byte, error) {
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
	if o.InitialRegistrationDate != nil {
		toSerialize["InitialRegistrationDate"] = o.InitialRegistrationDate
	}
	if o.LastRenewalAttemptDate != nil {
		toSerialize["LastRenewalAttemptDate"] = o.LastRenewalAttemptDate
	}
	if o.NextRenewalAttemptDate != nil {
		toSerialize["NextRenewalAttemptDate"] = o.NextRenewalAttemptDate
	}
	if o.OutOfComplianceStartDate != nil {
		toSerialize["OutOfComplianceStartDate"] = o.OutOfComplianceStartDate
	}
	if o.RegistrationExpireDate != nil {
		toSerialize["RegistrationExpireDate"] = o.RegistrationExpireDate
	}
	if o.RegistrationFailedReason != nil {
		toSerialize["RegistrationFailedReason"] = o.RegistrationFailedReason
	}
	if o.SmartAccount != nil {
		toSerialize["SmartAccount"] = o.SmartAccount
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.VirtualAccount != nil {
		toSerialize["VirtualAccount"] = o.VirtualAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxRegistrationDetailsDt) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Initial Registration Date
		InitialRegistrationDate *string `json:"InitialRegistrationDate,omitempty"`
		// Last License Renewal attempted Date
		LastRenewalAttemptDate *string `json:"LastRenewalAttemptDate,omitempty"`
		// Next Attempt Date for License Renewal
		NextRenewalAttemptDate *string `json:"NextRenewalAttemptDate,omitempty"`
		// Out of compliance start date
		OutOfComplianceStartDate *string `json:"OutOfComplianceStartDate,omitempty"`
		// Date when the registration will expire
		RegistrationExpireDate *string `json:"RegistrationExpireDate,omitempty"`
		// License registration success or failure reason
		RegistrationFailedReason *string `json:"RegistrationFailedReason,omitempty"`
		// Smart Account mapped to cluster
		SmartAccount *string `json:"SmartAccount,omitempty"`
		// Registration Status
		Status *string `json:"Status,omitempty"`
		// Virtual Account mapped to cluster
		VirtualAccount *string `json:"VirtualAccount,omitempty"`
	}

	varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct := HyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxRegistrationDetailsDt := _HyperflexHxRegistrationDetailsDt{}
		varHyperflexHxRegistrationDetailsDt.ClassId = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.ClassId
		varHyperflexHxRegistrationDetailsDt.ObjectType = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.ObjectType
		varHyperflexHxRegistrationDetailsDt.InitialRegistrationDate = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.InitialRegistrationDate
		varHyperflexHxRegistrationDetailsDt.LastRenewalAttemptDate = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.LastRenewalAttemptDate
		varHyperflexHxRegistrationDetailsDt.NextRenewalAttemptDate = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.NextRenewalAttemptDate
		varHyperflexHxRegistrationDetailsDt.OutOfComplianceStartDate = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.OutOfComplianceStartDate
		varHyperflexHxRegistrationDetailsDt.RegistrationExpireDate = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.RegistrationExpireDate
		varHyperflexHxRegistrationDetailsDt.RegistrationFailedReason = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.RegistrationFailedReason
		varHyperflexHxRegistrationDetailsDt.SmartAccount = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.SmartAccount
		varHyperflexHxRegistrationDetailsDt.Status = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.Status
		varHyperflexHxRegistrationDetailsDt.VirtualAccount = varHyperflexHxRegistrationDetailsDtWithoutEmbeddedStruct.VirtualAccount
		*o = HyperflexHxRegistrationDetailsDt(varHyperflexHxRegistrationDetailsDt)
	} else {
		return err
	}

	varHyperflexHxRegistrationDetailsDt := _HyperflexHxRegistrationDetailsDt{}

	err = json.Unmarshal(bytes, &varHyperflexHxRegistrationDetailsDt)
	if err == nil {
		o.MoBaseComplexType = varHyperflexHxRegistrationDetailsDt.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitialRegistrationDate")
		delete(additionalProperties, "LastRenewalAttemptDate")
		delete(additionalProperties, "NextRenewalAttemptDate")
		delete(additionalProperties, "OutOfComplianceStartDate")
		delete(additionalProperties, "RegistrationExpireDate")
		delete(additionalProperties, "RegistrationFailedReason")
		delete(additionalProperties, "SmartAccount")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "VirtualAccount")

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

type NullableHyperflexHxRegistrationDetailsDt struct {
	value *HyperflexHxRegistrationDetailsDt
	isSet bool
}

func (v NullableHyperflexHxRegistrationDetailsDt) Get() *HyperflexHxRegistrationDetailsDt {
	return v.value
}

func (v *NullableHyperflexHxRegistrationDetailsDt) Set(val *HyperflexHxRegistrationDetailsDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxRegistrationDetailsDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxRegistrationDetailsDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxRegistrationDetailsDt(val *HyperflexHxRegistrationDetailsDt) *NullableHyperflexHxRegistrationDetailsDt {
	return &NullableHyperflexHxRegistrationDetailsDt{value: val, isSet: true}
}

func (v NullableHyperflexHxRegistrationDetailsDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxRegistrationDetailsDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
