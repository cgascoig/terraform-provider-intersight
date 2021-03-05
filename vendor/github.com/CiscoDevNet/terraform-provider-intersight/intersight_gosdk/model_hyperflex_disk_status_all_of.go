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

// HyperflexDiskStatusAllOf Definition of the list of properties defined in 'hyperflex.DiskStatus', excluding properties defined in parent classes.
type HyperflexDiskStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Percentage of download completed.
	DownloadPercentage *string `json:"DownloadPercentage,omitempty"`
	// Current state of the virtual disk. * `Unknown` - No details available on the disk state. * `Succeeded` - Last operation on the disk has been successful. * `ImportInProgress` - Import operation on the disk is in progress. * `ImportFailed` - Import operation on the disk has failed. * `CloneInProgress` - Disk clone operation on the disk is in progress. * `CloneFailed` - Clone operation on the disk has failed. * `CloneScheduled` - Clone operation on the disk has been scheduled. * `ImportScheduled` - Import operation on the disk has been scheduled. * `Pending` - Submitted operation on the disk is currently pending.
	State *string `json:"State,omitempty"`
	// Identity of the Volume associated with virtual machine disk.
	VolumeHandle         *string `json:"VolumeHandle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDiskStatusAllOf HyperflexDiskStatusAllOf

// NewHyperflexDiskStatusAllOf instantiates a new HyperflexDiskStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDiskStatusAllOf(classId string, objectType string) *HyperflexDiskStatusAllOf {
	this := HyperflexDiskStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "Unknown"
	this.State = &state
	return &this
}

// NewHyperflexDiskStatusAllOfWithDefaults instantiates a new HyperflexDiskStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDiskStatusAllOfWithDefaults() *HyperflexDiskStatusAllOf {
	this := HyperflexDiskStatusAllOf{}
	var classId string = "hyperflex.DiskStatus"
	this.ClassId = classId
	var objectType string = "hyperflex.DiskStatus"
	this.ObjectType = objectType
	var state string = "Unknown"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDiskStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDiskStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDiskStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDiskStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDownloadPercentage returns the DownloadPercentage field value if set, zero value otherwise.
func (o *HyperflexDiskStatusAllOf) GetDownloadPercentage() string {
	if o == nil || o.DownloadPercentage == nil {
		var ret string
		return ret
	}
	return *o.DownloadPercentage
}

// GetDownloadPercentageOk returns a tuple with the DownloadPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatusAllOf) GetDownloadPercentageOk() (*string, bool) {
	if o == nil || o.DownloadPercentage == nil {
		return nil, false
	}
	return o.DownloadPercentage, true
}

// HasDownloadPercentage returns a boolean if a field has been set.
func (o *HyperflexDiskStatusAllOf) HasDownloadPercentage() bool {
	if o != nil && o.DownloadPercentage != nil {
		return true
	}

	return false
}

// SetDownloadPercentage gets a reference to the given string and assigns it to the DownloadPercentage field.
func (o *HyperflexDiskStatusAllOf) SetDownloadPercentage(v string) {
	o.DownloadPercentage = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexDiskStatusAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatusAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexDiskStatusAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexDiskStatusAllOf) SetState(v string) {
	o.State = &v
}

// GetVolumeHandle returns the VolumeHandle field value if set, zero value otherwise.
func (o *HyperflexDiskStatusAllOf) GetVolumeHandle() string {
	if o == nil || o.VolumeHandle == nil {
		var ret string
		return ret
	}
	return *o.VolumeHandle
}

// GetVolumeHandleOk returns a tuple with the VolumeHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatusAllOf) GetVolumeHandleOk() (*string, bool) {
	if o == nil || o.VolumeHandle == nil {
		return nil, false
	}
	return o.VolumeHandle, true
}

// HasVolumeHandle returns a boolean if a field has been set.
func (o *HyperflexDiskStatusAllOf) HasVolumeHandle() bool {
	if o != nil && o.VolumeHandle != nil {
		return true
	}

	return false
}

// SetVolumeHandle gets a reference to the given string and assigns it to the VolumeHandle field.
func (o *HyperflexDiskStatusAllOf) SetVolumeHandle(v string) {
	o.VolumeHandle = &v
}

func (o HyperflexDiskStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DownloadPercentage != nil {
		toSerialize["DownloadPercentage"] = o.DownloadPercentage
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VolumeHandle != nil {
		toSerialize["VolumeHandle"] = o.VolumeHandle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDiskStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexDiskStatusAllOf := _HyperflexDiskStatusAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexDiskStatusAllOf); err == nil {
		*o = HyperflexDiskStatusAllOf(varHyperflexDiskStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DownloadPercentage")
		delete(additionalProperties, "State")
		delete(additionalProperties, "VolumeHandle")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexDiskStatusAllOf struct {
	value *HyperflexDiskStatusAllOf
	isSet bool
}

func (v NullableHyperflexDiskStatusAllOf) Get() *HyperflexDiskStatusAllOf {
	return v.value
}

func (v *NullableHyperflexDiskStatusAllOf) Set(val *HyperflexDiskStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDiskStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDiskStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDiskStatusAllOf(val *HyperflexDiskStatusAllOf) *NullableHyperflexDiskStatusAllOf {
	return &NullableHyperflexDiskStatusAllOf{value: val, isSet: true}
}

func (v NullableHyperflexDiskStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDiskStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
