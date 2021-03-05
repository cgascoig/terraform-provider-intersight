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
)

// HyperflexReplicationScheduleAllOf Definition of the list of properties defined in 'hyperflex.ReplicationSchedule', excluding properties defined in parent classes.
type HyperflexReplicationScheduleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Time interval between two copies in minutes.
	BackupInterval       *int64 `json:"BackupInterval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReplicationScheduleAllOf HyperflexReplicationScheduleAllOf

// NewHyperflexReplicationScheduleAllOf instantiates a new HyperflexReplicationScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationScheduleAllOf(classId string, objectType string) *HyperflexReplicationScheduleAllOf {
	this := HyperflexReplicationScheduleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var backupInterval int64 = 240
	this.BackupInterval = &backupInterval
	return &this
}

// NewHyperflexReplicationScheduleAllOfWithDefaults instantiates a new HyperflexReplicationScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationScheduleAllOfWithDefaults() *HyperflexReplicationScheduleAllOf {
	this := HyperflexReplicationScheduleAllOf{}
	var classId string = "hyperflex.ReplicationSchedule"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationSchedule"
	this.ObjectType = objectType
	var backupInterval int64 = 240
	this.BackupInterval = &backupInterval
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationScheduleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationScheduleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationScheduleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationScheduleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationScheduleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationScheduleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupInterval returns the BackupInterval field value if set, zero value otherwise.
func (o *HyperflexReplicationScheduleAllOf) GetBackupInterval() int64 {
	if o == nil || o.BackupInterval == nil {
		var ret int64
		return ret
	}
	return *o.BackupInterval
}

// GetBackupIntervalOk returns a tuple with the BackupInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationScheduleAllOf) GetBackupIntervalOk() (*int64, bool) {
	if o == nil || o.BackupInterval == nil {
		return nil, false
	}
	return o.BackupInterval, true
}

// HasBackupInterval returns a boolean if a field has been set.
func (o *HyperflexReplicationScheduleAllOf) HasBackupInterval() bool {
	if o != nil && o.BackupInterval != nil {
		return true
	}

	return false
}

// SetBackupInterval gets a reference to the given int64 and assigns it to the BackupInterval field.
func (o *HyperflexReplicationScheduleAllOf) SetBackupInterval(v int64) {
	o.BackupInterval = &v
}

func (o HyperflexReplicationScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackupInterval != nil {
		toSerialize["BackupInterval"] = o.BackupInterval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexReplicationScheduleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexReplicationScheduleAllOf := _HyperflexReplicationScheduleAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexReplicationScheduleAllOf); err == nil {
		*o = HyperflexReplicationScheduleAllOf(varHyperflexReplicationScheduleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupInterval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexReplicationScheduleAllOf struct {
	value *HyperflexReplicationScheduleAllOf
	isSet bool
}

func (v NullableHyperflexReplicationScheduleAllOf) Get() *HyperflexReplicationScheduleAllOf {
	return v.value
}

func (v *NullableHyperflexReplicationScheduleAllOf) Set(val *HyperflexReplicationScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationScheduleAllOf(val *HyperflexReplicationScheduleAllOf) *NullableHyperflexReplicationScheduleAllOf {
	return &NullableHyperflexReplicationScheduleAllOf{value: val, isSet: true}
}

func (v NullableHyperflexReplicationScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
