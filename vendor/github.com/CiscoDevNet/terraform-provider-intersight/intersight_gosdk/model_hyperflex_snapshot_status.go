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
	"reflect"
	"strings"
)

// HyperflexSnapshotStatus The current state of this virtual machine's snapshot.
type HyperflexSnapshotStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of this Snapshot Point.
	Description *string                     `json:"Description,omitempty"`
	Error       NullableHyperflexErrorStack `json:"Error,omitempty"`
	// Completion percentage for this snapshot.
	PctComplete *int64 `json:"PctComplete,omitempty"`
	// Current snapshot state for this snapshot. * `SUCCESS` - This snapshot status code is success. * `FAILED` - This snapshot status code is failed. * `IN_PROGRESS` - This snapshot status code is in progress. * `DELETING` - This snapshot status code is deleting. * `DELETED` - This snapshot status code is deleted. * `NONE` - This snapshot status code is none.
	Status *string `json:"Status,omitempty"`
	// Timestamp at which the Snapshot is taken.
	Timestamp *int64 `json:"Timestamp,omitempty"`
	// Space Used by this Snapshot Point.
	UsedSpace            *int64 `json:"UsedSpace,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexSnapshotStatus HyperflexSnapshotStatus

// NewHyperflexSnapshotStatus instantiates a new HyperflexSnapshotStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSnapshotStatus(classId string, objectType string) *HyperflexSnapshotStatus {
	this := HyperflexSnapshotStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "SUCCESS"
	this.Status = &status
	return &this
}

// NewHyperflexSnapshotStatusWithDefaults instantiates a new HyperflexSnapshotStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSnapshotStatusWithDefaults() *HyperflexSnapshotStatus {
	this := HyperflexSnapshotStatus{}
	var classId string = "hyperflex.SnapshotStatus"
	this.ClassId = classId
	var objectType string = "hyperflex.SnapshotStatus"
	this.ObjectType = objectType
	var status string = "SUCCESS"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSnapshotStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSnapshotStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSnapshotStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSnapshotStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexSnapshotStatus) SetDescription(v string) {
	o.Description = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotStatus) GetError() HyperflexErrorStack {
	if o == nil || o.Error.Get() == nil {
		var ret HyperflexErrorStack
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotStatus) GetErrorOk() (*HyperflexErrorStack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableHyperflexErrorStack and assigns it to the Error field.
func (o *HyperflexSnapshotStatus) SetError(v HyperflexErrorStack) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *HyperflexSnapshotStatus) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *HyperflexSnapshotStatus) UnsetError() {
	o.Error.Unset()
}

// GetPctComplete returns the PctComplete field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetPctComplete() int64 {
	if o == nil || o.PctComplete == nil {
		var ret int64
		return ret
	}
	return *o.PctComplete
}

// GetPctCompleteOk returns a tuple with the PctComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetPctCompleteOk() (*int64, bool) {
	if o == nil || o.PctComplete == nil {
		return nil, false
	}
	return o.PctComplete, true
}

// HasPctComplete returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasPctComplete() bool {
	if o != nil && o.PctComplete != nil {
		return true
	}

	return false
}

// SetPctComplete gets a reference to the given int64 and assigns it to the PctComplete field.
func (o *HyperflexSnapshotStatus) SetPctComplete(v int64) {
	o.PctComplete = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexSnapshotStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *HyperflexSnapshotStatus) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetUsedSpace() int64 {
	if o == nil || o.UsedSpace == nil {
		var ret int64
		return ret
	}
	return *o.UsedSpace
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetUsedSpaceOk() (*int64, bool) {
	if o == nil || o.UsedSpace == nil {
		return nil, false
	}
	return o.UsedSpace, true
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasUsedSpace() bool {
	if o != nil && o.UsedSpace != nil {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given int64 and assigns it to the UsedSpace field.
func (o *HyperflexSnapshotStatus) SetUsedSpace(v int64) {
	o.UsedSpace = &v
}

func (o HyperflexSnapshotStatus) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Error.IsSet() {
		toSerialize["Error"] = o.Error.Get()
	}
	if o.PctComplete != nil {
		toSerialize["PctComplete"] = o.PctComplete
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.UsedSpace != nil {
		toSerialize["UsedSpace"] = o.UsedSpace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSnapshotStatus) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexSnapshotStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of this Snapshot Point.
		Description *string                     `json:"Description,omitempty"`
		Error       NullableHyperflexErrorStack `json:"Error,omitempty"`
		// Completion percentage for this snapshot.
		PctComplete *int64 `json:"PctComplete,omitempty"`
		// Current snapshot state for this snapshot. * `SUCCESS` - This snapshot status code is success. * `FAILED` - This snapshot status code is failed. * `IN_PROGRESS` - This snapshot status code is in progress. * `DELETING` - This snapshot status code is deleting. * `DELETED` - This snapshot status code is deleted. * `NONE` - This snapshot status code is none.
		Status *string `json:"Status,omitempty"`
		// Timestamp at which the Snapshot is taken.
		Timestamp *int64 `json:"Timestamp,omitempty"`
		// Space Used by this Snapshot Point.
		UsedSpace *int64 `json:"UsedSpace,omitempty"`
	}

	varHyperflexSnapshotStatusWithoutEmbeddedStruct := HyperflexSnapshotStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexSnapshotStatusWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSnapshotStatus := _HyperflexSnapshotStatus{}
		varHyperflexSnapshotStatus.ClassId = varHyperflexSnapshotStatusWithoutEmbeddedStruct.ClassId
		varHyperflexSnapshotStatus.ObjectType = varHyperflexSnapshotStatusWithoutEmbeddedStruct.ObjectType
		varHyperflexSnapshotStatus.Description = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Description
		varHyperflexSnapshotStatus.Error = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Error
		varHyperflexSnapshotStatus.PctComplete = varHyperflexSnapshotStatusWithoutEmbeddedStruct.PctComplete
		varHyperflexSnapshotStatus.Status = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Status
		varHyperflexSnapshotStatus.Timestamp = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Timestamp
		varHyperflexSnapshotStatus.UsedSpace = varHyperflexSnapshotStatusWithoutEmbeddedStruct.UsedSpace
		*o = HyperflexSnapshotStatus(varHyperflexSnapshotStatus)
	} else {
		return err
	}

	varHyperflexSnapshotStatus := _HyperflexSnapshotStatus{}

	err = json.Unmarshal(bytes, &varHyperflexSnapshotStatus)
	if err == nil {
		o.MoBaseComplexType = varHyperflexSnapshotStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "PctComplete")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "UsedSpace")

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

type NullableHyperflexSnapshotStatus struct {
	value *HyperflexSnapshotStatus
	isSet bool
}

func (v NullableHyperflexSnapshotStatus) Get() *HyperflexSnapshotStatus {
	return v.value
}

func (v *NullableHyperflexSnapshotStatus) Set(val *HyperflexSnapshotStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSnapshotStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSnapshotStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSnapshotStatus(val *HyperflexSnapshotStatus) *NullableHyperflexSnapshotStatus {
	return &NullableHyperflexSnapshotStatus{value: val, isSet: true}
}

func (v NullableHyperflexSnapshotStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSnapshotStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
