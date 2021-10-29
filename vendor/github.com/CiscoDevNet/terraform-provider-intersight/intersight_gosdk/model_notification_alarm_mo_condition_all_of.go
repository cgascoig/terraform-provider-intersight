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
)

// NotificationAlarmMoConditionAllOf Definition of the list of properties defined in 'notification.AlarmMoCondition', excluding properties defined in parent classes.
type NotificationAlarmMoConditionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	Severity             []string `json:"Severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationAlarmMoConditionAllOf NotificationAlarmMoConditionAllOf

// NewNotificationAlarmMoConditionAllOf instantiates a new NotificationAlarmMoConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAlarmMoConditionAllOf(classId string, objectType string) *NotificationAlarmMoConditionAllOf {
	this := NotificationAlarmMoConditionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationAlarmMoConditionAllOfWithDefaults instantiates a new NotificationAlarmMoConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAlarmMoConditionAllOfWithDefaults() *NotificationAlarmMoConditionAllOf {
	this := NotificationAlarmMoConditionAllOf{}
	var classId string = "notification.AlarmMoCondition"
	this.ClassId = classId
	var objectType string = "notification.AlarmMoCondition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationAlarmMoConditionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationAlarmMoConditionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationAlarmMoConditionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationAlarmMoConditionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationAlarmMoConditionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationAlarmMoConditionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationAlarmMoConditionAllOf) GetSeverity() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationAlarmMoConditionAllOf) GetSeverityOk() (*[]string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return &o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *NotificationAlarmMoConditionAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given []string and assigns it to the Severity field.
func (o *NotificationAlarmMoConditionAllOf) SetSeverity(v []string) {
	o.Severity = v
}

func (o NotificationAlarmMoConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationAlarmMoConditionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationAlarmMoConditionAllOf := _NotificationAlarmMoConditionAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationAlarmMoConditionAllOf); err == nil {
		*o = NotificationAlarmMoConditionAllOf(varNotificationAlarmMoConditionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Severity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationAlarmMoConditionAllOf struct {
	value *NotificationAlarmMoConditionAllOf
	isSet bool
}

func (v NullableNotificationAlarmMoConditionAllOf) Get() *NotificationAlarmMoConditionAllOf {
	return v.value
}

func (v *NullableNotificationAlarmMoConditionAllOf) Set(val *NotificationAlarmMoConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAlarmMoConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAlarmMoConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAlarmMoConditionAllOf(val *NotificationAlarmMoConditionAllOf) *NullableNotificationAlarmMoConditionAllOf {
	return &NullableNotificationAlarmMoConditionAllOf{value: val, isSet: true}
}

func (v NullableNotificationAlarmMoConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAlarmMoConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
