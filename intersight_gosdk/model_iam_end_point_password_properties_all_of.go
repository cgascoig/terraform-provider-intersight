/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IamEndPointPasswordPropertiesAllOf Definition of the list of properties defined in 'iam.EndPointPasswordProperties', excluding properties defined in parent classes.
type IamEndPointPasswordPropertiesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables password expiry on the endpoint.
	EnablePasswordExpiry *bool `json:"EnablePasswordExpiry,omitempty"`
	// Enables a strong password policy. Strong password requirements: A. The password must have a minimum of 8 and a maximum of 20 characters. B. The password must not contain the User's Name. C. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &, *, -, _, +, =).
	EnforceStrongPassword *bool `json:"EnforceStrongPassword,omitempty"`
	// User password will always be sent to endpoint device. If the option is not selected, then user password will be sent to endpoint device for new users and if user password is changed for existing users.
	ForceSendPassword *bool `json:"ForceSendPassword,omitempty"`
	// Time period until when you can use the existing password, after it expires.
	GracePeriod *int64 `json:"GracePeriod,omitempty"`
	// The duration after which the password will expire.
	NotificationPeriod *int64 `json:"NotificationPeriod,omitempty"`
	// Set time period for password expiration. Value should be greater than notification period and grace period.
	PasswordExpiryDuration *int64 `json:"PasswordExpiryDuration,omitempty"`
	// Tracks password change history. Specifies in number of instances, that the new password was already used.
	PasswordHistory      *int64 `json:"PasswordHistory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointPasswordPropertiesAllOf IamEndPointPasswordPropertiesAllOf

// NewIamEndPointPasswordPropertiesAllOf instantiates a new IamEndPointPasswordPropertiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointPasswordPropertiesAllOf(classId string, objectType string) *IamEndPointPasswordPropertiesAllOf {
	this := IamEndPointPasswordPropertiesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enablePasswordExpiry bool = false
	this.EnablePasswordExpiry = &enablePasswordExpiry
	var enforceStrongPassword bool = true
	this.EnforceStrongPassword = &enforceStrongPassword
	var forceSendPassword bool = false
	this.ForceSendPassword = &forceSendPassword
	var gracePeriod int64 = 0
	this.GracePeriod = &gracePeriod
	var notificationPeriod int64 = 15
	this.NotificationPeriod = &notificationPeriod
	var passwordExpiryDuration int64 = 90
	this.PasswordExpiryDuration = &passwordExpiryDuration
	var passwordHistory int64 = 5
	this.PasswordHistory = &passwordHistory
	return &this
}

// NewIamEndPointPasswordPropertiesAllOfWithDefaults instantiates a new IamEndPointPasswordPropertiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointPasswordPropertiesAllOfWithDefaults() *IamEndPointPasswordPropertiesAllOf {
	this := IamEndPointPasswordPropertiesAllOf{}
	var classId string = "iam.EndPointPasswordProperties"
	this.ClassId = classId
	var objectType string = "iam.EndPointPasswordProperties"
	this.ObjectType = objectType
	var enablePasswordExpiry bool = false
	this.EnablePasswordExpiry = &enablePasswordExpiry
	var enforceStrongPassword bool = true
	this.EnforceStrongPassword = &enforceStrongPassword
	var forceSendPassword bool = false
	this.ForceSendPassword = &forceSendPassword
	var gracePeriod int64 = 0
	this.GracePeriod = &gracePeriod
	var notificationPeriod int64 = 15
	this.NotificationPeriod = &notificationPeriod
	var passwordExpiryDuration int64 = 90
	this.PasswordExpiryDuration = &passwordExpiryDuration
	var passwordHistory int64 = 5
	this.PasswordHistory = &passwordHistory
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointPasswordPropertiesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointPasswordPropertiesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointPasswordPropertiesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointPasswordPropertiesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnablePasswordExpiry returns the EnablePasswordExpiry field value if set, zero value otherwise.
func (o *IamEndPointPasswordPropertiesAllOf) GetEnablePasswordExpiry() bool {
	if o == nil || o.EnablePasswordExpiry == nil {
		var ret bool
		return ret
	}
	return *o.EnablePasswordExpiry
}

// GetEnablePasswordExpiryOk returns a tuple with the EnablePasswordExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetEnablePasswordExpiryOk() (*bool, bool) {
	if o == nil || o.EnablePasswordExpiry == nil {
		return nil, false
	}
	return o.EnablePasswordExpiry, true
}

// HasEnablePasswordExpiry returns a boolean if a field has been set.
func (o *IamEndPointPasswordPropertiesAllOf) HasEnablePasswordExpiry() bool {
	if o != nil && o.EnablePasswordExpiry != nil {
		return true
	}

	return false
}

// SetEnablePasswordExpiry gets a reference to the given bool and assigns it to the EnablePasswordExpiry field.
func (o *IamEndPointPasswordPropertiesAllOf) SetEnablePasswordExpiry(v bool) {
	o.EnablePasswordExpiry = &v
}

// GetEnforceStrongPassword returns the EnforceStrongPassword field value if set, zero value otherwise.
func (o *IamEndPointPasswordPropertiesAllOf) GetEnforceStrongPassword() bool {
	if o == nil || o.EnforceStrongPassword == nil {
		var ret bool
		return ret
	}
	return *o.EnforceStrongPassword
}

// GetEnforceStrongPasswordOk returns a tuple with the EnforceStrongPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetEnforceStrongPasswordOk() (*bool, bool) {
	if o == nil || o.EnforceStrongPassword == nil {
		return nil, false
	}
	return o.EnforceStrongPassword, true
}

// HasEnforceStrongPassword returns a boolean if a field has been set.
func (o *IamEndPointPasswordPropertiesAllOf) HasEnforceStrongPassword() bool {
	if o != nil && o.EnforceStrongPassword != nil {
		return true
	}

	return false
}

// SetEnforceStrongPassword gets a reference to the given bool and assigns it to the EnforceStrongPassword field.
func (o *IamEndPointPasswordPropertiesAllOf) SetEnforceStrongPassword(v bool) {
	o.EnforceStrongPassword = &v
}

// GetForceSendPassword returns the ForceSendPassword field value if set, zero value otherwise.
func (o *IamEndPointPasswordPropertiesAllOf) GetForceSendPassword() bool {
	if o == nil || o.ForceSendPassword == nil {
		var ret bool
		return ret
	}
	return *o.ForceSendPassword
}

// GetForceSendPasswordOk returns a tuple with the ForceSendPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetForceSendPasswordOk() (*bool, bool) {
	if o == nil || o.ForceSendPassword == nil {
		return nil, false
	}
	return o.ForceSendPassword, true
}

// HasForceSendPassword returns a boolean if a field has been set.
func (o *IamEndPointPasswordPropertiesAllOf) HasForceSendPassword() bool {
	if o != nil && o.ForceSendPassword != nil {
		return true
	}

	return false
}

// SetForceSendPassword gets a reference to the given bool and assigns it to the ForceSendPassword field.
func (o *IamEndPointPasswordPropertiesAllOf) SetForceSendPassword(v bool) {
	o.ForceSendPassword = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *IamEndPointPasswordPropertiesAllOf) GetGracePeriod() int64 {
	if o == nil || o.GracePeriod == nil {
		var ret int64
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetGracePeriodOk() (*int64, bool) {
	if o == nil || o.GracePeriod == nil {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *IamEndPointPasswordPropertiesAllOf) HasGracePeriod() bool {
	if o != nil && o.GracePeriod != nil {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given int64 and assigns it to the GracePeriod field.
func (o *IamEndPointPasswordPropertiesAllOf) SetGracePeriod(v int64) {
	o.GracePeriod = &v
}

// GetNotificationPeriod returns the NotificationPeriod field value if set, zero value otherwise.
func (o *IamEndPointPasswordPropertiesAllOf) GetNotificationPeriod() int64 {
	if o == nil || o.NotificationPeriod == nil {
		var ret int64
		return ret
	}
	return *o.NotificationPeriod
}

// GetNotificationPeriodOk returns a tuple with the NotificationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetNotificationPeriodOk() (*int64, bool) {
	if o == nil || o.NotificationPeriod == nil {
		return nil, false
	}
	return o.NotificationPeriod, true
}

// HasNotificationPeriod returns a boolean if a field has been set.
func (o *IamEndPointPasswordPropertiesAllOf) HasNotificationPeriod() bool {
	if o != nil && o.NotificationPeriod != nil {
		return true
	}

	return false
}

// SetNotificationPeriod gets a reference to the given int64 and assigns it to the NotificationPeriod field.
func (o *IamEndPointPasswordPropertiesAllOf) SetNotificationPeriod(v int64) {
	o.NotificationPeriod = &v
}

// GetPasswordExpiryDuration returns the PasswordExpiryDuration field value if set, zero value otherwise.
func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordExpiryDuration() int64 {
	if o == nil || o.PasswordExpiryDuration == nil {
		var ret int64
		return ret
	}
	return *o.PasswordExpiryDuration
}

// GetPasswordExpiryDurationOk returns a tuple with the PasswordExpiryDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordExpiryDurationOk() (*int64, bool) {
	if o == nil || o.PasswordExpiryDuration == nil {
		return nil, false
	}
	return o.PasswordExpiryDuration, true
}

// HasPasswordExpiryDuration returns a boolean if a field has been set.
func (o *IamEndPointPasswordPropertiesAllOf) HasPasswordExpiryDuration() bool {
	if o != nil && o.PasswordExpiryDuration != nil {
		return true
	}

	return false
}

// SetPasswordExpiryDuration gets a reference to the given int64 and assigns it to the PasswordExpiryDuration field.
func (o *IamEndPointPasswordPropertiesAllOf) SetPasswordExpiryDuration(v int64) {
	o.PasswordExpiryDuration = &v
}

// GetPasswordHistory returns the PasswordHistory field value if set, zero value otherwise.
func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordHistory() int64 {
	if o == nil || o.PasswordHistory == nil {
		var ret int64
		return ret
	}
	return *o.PasswordHistory
}

// GetPasswordHistoryOk returns a tuple with the PasswordHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordHistoryOk() (*int64, bool) {
	if o == nil || o.PasswordHistory == nil {
		return nil, false
	}
	return o.PasswordHistory, true
}

// HasPasswordHistory returns a boolean if a field has been set.
func (o *IamEndPointPasswordPropertiesAllOf) HasPasswordHistory() bool {
	if o != nil && o.PasswordHistory != nil {
		return true
	}

	return false
}

// SetPasswordHistory gets a reference to the given int64 and assigns it to the PasswordHistory field.
func (o *IamEndPointPasswordPropertiesAllOf) SetPasswordHistory(v int64) {
	o.PasswordHistory = &v
}

func (o IamEndPointPasswordPropertiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnablePasswordExpiry != nil {
		toSerialize["EnablePasswordExpiry"] = o.EnablePasswordExpiry
	}
	if o.EnforceStrongPassword != nil {
		toSerialize["EnforceStrongPassword"] = o.EnforceStrongPassword
	}
	if o.ForceSendPassword != nil {
		toSerialize["ForceSendPassword"] = o.ForceSendPassword
	}
	if o.GracePeriod != nil {
		toSerialize["GracePeriod"] = o.GracePeriod
	}
	if o.NotificationPeriod != nil {
		toSerialize["NotificationPeriod"] = o.NotificationPeriod
	}
	if o.PasswordExpiryDuration != nil {
		toSerialize["PasswordExpiryDuration"] = o.PasswordExpiryDuration
	}
	if o.PasswordHistory != nil {
		toSerialize["PasswordHistory"] = o.PasswordHistory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointPasswordPropertiesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamEndPointPasswordPropertiesAllOf := _IamEndPointPasswordPropertiesAllOf{}

	if err = json.Unmarshal(bytes, &varIamEndPointPasswordPropertiesAllOf); err == nil {
		*o = IamEndPointPasswordPropertiesAllOf(varIamEndPointPasswordPropertiesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnablePasswordExpiry")
		delete(additionalProperties, "EnforceStrongPassword")
		delete(additionalProperties, "ForceSendPassword")
		delete(additionalProperties, "GracePeriod")
		delete(additionalProperties, "NotificationPeriod")
		delete(additionalProperties, "PasswordExpiryDuration")
		delete(additionalProperties, "PasswordHistory")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamEndPointPasswordPropertiesAllOf struct {
	value *IamEndPointPasswordPropertiesAllOf
	isSet bool
}

func (v NullableIamEndPointPasswordPropertiesAllOf) Get() *IamEndPointPasswordPropertiesAllOf {
	return v.value
}

func (v *NullableIamEndPointPasswordPropertiesAllOf) Set(val *IamEndPointPasswordPropertiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointPasswordPropertiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointPasswordPropertiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointPasswordPropertiesAllOf(val *IamEndPointPasswordPropertiesAllOf) *NullableIamEndPointPasswordPropertiesAllOf {
	return &NullableIamEndPointPasswordPropertiesAllOf{value: val, isSet: true}
}

func (v NullableIamEndPointPasswordPropertiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointPasswordPropertiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
