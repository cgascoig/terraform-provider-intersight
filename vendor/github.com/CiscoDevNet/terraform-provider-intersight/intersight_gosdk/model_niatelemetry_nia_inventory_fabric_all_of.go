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

// NiatelemetryNiaInventoryFabricAllOf Definition of the list of properties defined in 'niatelemetry.NiaInventoryFabric', excluding properties defined in parent classes.
type NiatelemetryNiaInventoryFabricAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the aycast gateway mac.
	AnycastGwMac *string `json:"AnycastGwMac,omitempty"`
	// Returns the value of the dcnmtrackerEnabled field.
	DcnmtrackerEnabled *bool `json:"DcnmtrackerEnabled,omitempty"`
	// Uniquely identifies a fabric.
	FabricId *string `json:"FabricId,omitempty"`
	// Returns the value of the Name of a fabric.
	FabricName *string `json:"FabricName,omitempty"`
	// Returns if ngoam is enabled.
	IsNgoamEnabled *string `json:"IsNgoamEnabled,omitempty"`
	// Returns if the scheduled backup is enabled.
	IsScheduledBackUpEnabled *string `json:"IsScheduledBackUpEnabled,omitempty"`
	// Returns total number of leafs in the fabric.
	LeafCount    *int64                    `json:"LeafCount,omitempty"`
	LogicalLinks []NiatelemetryLogicalLink `json:"LogicalLinks,omitempty"`
	// Returns the value of the nxosVrfCount field.
	NxosVrfCount *int64 `json:"NxosVrfCount,omitempty"`
	// Returns total number of spines in the fabric.
	SpineCount           *int64                               `json:"SpineCount,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNiaInventoryFabricAllOf NiatelemetryNiaInventoryFabricAllOf

// NewNiatelemetryNiaInventoryFabricAllOf instantiates a new NiatelemetryNiaInventoryFabricAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaInventoryFabricAllOf(classId string, objectType string) *NiatelemetryNiaInventoryFabricAllOf {
	this := NiatelemetryNiaInventoryFabricAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNiaInventoryFabricAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryFabricAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaInventoryFabricAllOfWithDefaults() *NiatelemetryNiaInventoryFabricAllOf {
	this := NiatelemetryNiaInventoryFabricAllOf{}
	var classId string = "niatelemetry.NiaInventoryFabric"
	this.ClassId = classId
	var objectType string = "niatelemetry.NiaInventoryFabric"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNiaInventoryFabricAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNiaInventoryFabricAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnycastGwMac returns the AnycastGwMac field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMac() string {
	if o == nil || o.AnycastGwMac == nil {
		var ret string
		return ret
	}
	return *o.AnycastGwMac
}

// GetAnycastGwMacOk returns a tuple with the AnycastGwMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMacOk() (*string, bool) {
	if o == nil || o.AnycastGwMac == nil {
		return nil, false
	}
	return o.AnycastGwMac, true
}

// HasAnycastGwMac returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasAnycastGwMac() bool {
	if o != nil && o.AnycastGwMac != nil {
		return true
	}

	return false
}

// SetAnycastGwMac gets a reference to the given string and assigns it to the AnycastGwMac field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetAnycastGwMac(v string) {
	o.AnycastGwMac = &v
}

// GetDcnmtrackerEnabled returns the DcnmtrackerEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabled() bool {
	if o == nil || o.DcnmtrackerEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DcnmtrackerEnabled
}

// GetDcnmtrackerEnabledOk returns a tuple with the DcnmtrackerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabledOk() (*bool, bool) {
	if o == nil || o.DcnmtrackerEnabled == nil {
		return nil, false
	}
	return o.DcnmtrackerEnabled, true
}

// HasDcnmtrackerEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasDcnmtrackerEnabled() bool {
	if o != nil && o.DcnmtrackerEnabled != nil {
		return true
	}

	return false
}

// SetDcnmtrackerEnabled gets a reference to the given bool and assigns it to the DcnmtrackerEnabled field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetDcnmtrackerEnabled(v bool) {
	o.DcnmtrackerEnabled = &v
}

// GetFabricId returns the FabricId field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricId() string {
	if o == nil || o.FabricId == nil {
		var ret string
		return ret
	}
	return *o.FabricId
}

// GetFabricIdOk returns a tuple with the FabricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricIdOk() (*string, bool) {
	if o == nil || o.FabricId == nil {
		return nil, false
	}
	return o.FabricId, true
}

// HasFabricId returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricId() bool {
	if o != nil && o.FabricId != nil {
		return true
	}

	return false
}

// SetFabricId gets a reference to the given string and assigns it to the FabricId field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricId(v string) {
	o.FabricId = &v
}

// GetFabricName returns the FabricName field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricName() string {
	if o == nil || o.FabricName == nil {
		var ret string
		return ret
	}
	return *o.FabricName
}

// GetFabricNameOk returns a tuple with the FabricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricNameOk() (*string, bool) {
	if o == nil || o.FabricName == nil {
		return nil, false
	}
	return o.FabricName, true
}

// HasFabricName returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricName() bool {
	if o != nil && o.FabricName != nil {
		return true
	}

	return false
}

// SetFabricName gets a reference to the given string and assigns it to the FabricName field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricName(v string) {
	o.FabricName = &v
}

// GetIsNgoamEnabled returns the IsNgoamEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabled() string {
	if o == nil || o.IsNgoamEnabled == nil {
		var ret string
		return ret
	}
	return *o.IsNgoamEnabled
}

// GetIsNgoamEnabledOk returns a tuple with the IsNgoamEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabledOk() (*string, bool) {
	if o == nil || o.IsNgoamEnabled == nil {
		return nil, false
	}
	return o.IsNgoamEnabled, true
}

// HasIsNgoamEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsNgoamEnabled() bool {
	if o != nil && o.IsNgoamEnabled != nil {
		return true
	}

	return false
}

// SetIsNgoamEnabled gets a reference to the given string and assigns it to the IsNgoamEnabled field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsNgoamEnabled(v string) {
	o.IsNgoamEnabled = &v
}

// GetIsScheduledBackUpEnabled returns the IsScheduledBackUpEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabled() string {
	if o == nil || o.IsScheduledBackUpEnabled == nil {
		var ret string
		return ret
	}
	return *o.IsScheduledBackUpEnabled
}

// GetIsScheduledBackUpEnabledOk returns a tuple with the IsScheduledBackUpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabledOk() (*string, bool) {
	if o == nil || o.IsScheduledBackUpEnabled == nil {
		return nil, false
	}
	return o.IsScheduledBackUpEnabled, true
}

// HasIsScheduledBackUpEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsScheduledBackUpEnabled() bool {
	if o != nil && o.IsScheduledBackUpEnabled != nil {
		return true
	}

	return false
}

// SetIsScheduledBackUpEnabled gets a reference to the given string and assigns it to the IsScheduledBackUpEnabled field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsScheduledBackUpEnabled(v string) {
	o.IsScheduledBackUpEnabled = &v
}

// GetLeafCount returns the LeafCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCount() int64 {
	if o == nil || o.LeafCount == nil {
		var ret int64
		return ret
	}
	return *o.LeafCount
}

// GetLeafCountOk returns a tuple with the LeafCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCountOk() (*int64, bool) {
	if o == nil || o.LeafCount == nil {
		return nil, false
	}
	return o.LeafCount, true
}

// HasLeafCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasLeafCount() bool {
	if o != nil && o.LeafCount != nil {
		return true
	}

	return false
}

// SetLeafCount gets a reference to the given int64 and assigns it to the LeafCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetLeafCount(v int64) {
	o.LeafCount = &v
}

// GetLogicalLinks returns the LogicalLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinks() []NiatelemetryLogicalLink {
	if o == nil {
		var ret []NiatelemetryLogicalLink
		return ret
	}
	return o.LogicalLinks
}

// GetLogicalLinksOk returns a tuple with the LogicalLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinksOk() (*[]NiatelemetryLogicalLink, bool) {
	if o == nil || o.LogicalLinks == nil {
		return nil, false
	}
	return &o.LogicalLinks, true
}

// HasLogicalLinks returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasLogicalLinks() bool {
	if o != nil && o.LogicalLinks != nil {
		return true
	}

	return false
}

// SetLogicalLinks gets a reference to the given []NiatelemetryLogicalLink and assigns it to the LogicalLinks field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetLogicalLinks(v []NiatelemetryLogicalLink) {
	o.LogicalLinks = v
}

// GetNxosVrfCount returns the NxosVrfCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCount() int64 {
	if o == nil || o.NxosVrfCount == nil {
		var ret int64
		return ret
	}
	return *o.NxosVrfCount
}

// GetNxosVrfCountOk returns a tuple with the NxosVrfCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCountOk() (*int64, bool) {
	if o == nil || o.NxosVrfCount == nil {
		return nil, false
	}
	return o.NxosVrfCount, true
}

// HasNxosVrfCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVrfCount() bool {
	if o != nil && o.NxosVrfCount != nil {
		return true
	}

	return false
}

// SetNxosVrfCount gets a reference to the given int64 and assigns it to the NxosVrfCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVrfCount(v int64) {
	o.NxosVrfCount = &v
}

// GetSpineCount returns the SpineCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCount() int64 {
	if o == nil || o.SpineCount == nil {
		var ret int64
		return ret
	}
	return *o.SpineCount
}

// GetSpineCountOk returns a tuple with the SpineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCountOk() (*int64, bool) {
	if o == nil || o.SpineCount == nil {
		return nil, false
	}
	return o.SpineCount, true
}

// HasSpineCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasSpineCount() bool {
	if o != nil && o.SpineCount != nil {
		return true
	}

	return false
}

// SetSpineCount gets a reference to the given int64 and assigns it to the SpineCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetSpineCount(v int64) {
	o.SpineCount = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNiaInventoryFabricAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AnycastGwMac != nil {
		toSerialize["AnycastGwMac"] = o.AnycastGwMac
	}
	if o.DcnmtrackerEnabled != nil {
		toSerialize["DcnmtrackerEnabled"] = o.DcnmtrackerEnabled
	}
	if o.FabricId != nil {
		toSerialize["FabricId"] = o.FabricId
	}
	if o.FabricName != nil {
		toSerialize["FabricName"] = o.FabricName
	}
	if o.IsNgoamEnabled != nil {
		toSerialize["IsNgoamEnabled"] = o.IsNgoamEnabled
	}
	if o.IsScheduledBackUpEnabled != nil {
		toSerialize["IsScheduledBackUpEnabled"] = o.IsScheduledBackUpEnabled
	}
	if o.LeafCount != nil {
		toSerialize["LeafCount"] = o.LeafCount
	}
	if o.LogicalLinks != nil {
		toSerialize["LogicalLinks"] = o.LogicalLinks
	}
	if o.NxosVrfCount != nil {
		toSerialize["NxosVrfCount"] = o.NxosVrfCount
	}
	if o.SpineCount != nil {
		toSerialize["SpineCount"] = o.SpineCount
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNiaInventoryFabricAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNiaInventoryFabricAllOf := _NiatelemetryNiaInventoryFabricAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNiaInventoryFabricAllOf); err == nil {
		*o = NiatelemetryNiaInventoryFabricAllOf(varNiatelemetryNiaInventoryFabricAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AnycastGwMac")
		delete(additionalProperties, "DcnmtrackerEnabled")
		delete(additionalProperties, "FabricId")
		delete(additionalProperties, "FabricName")
		delete(additionalProperties, "IsNgoamEnabled")
		delete(additionalProperties, "IsScheduledBackUpEnabled")
		delete(additionalProperties, "LeafCount")
		delete(additionalProperties, "LogicalLinks")
		delete(additionalProperties, "NxosVrfCount")
		delete(additionalProperties, "SpineCount")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNiaInventoryFabricAllOf struct {
	value *NiatelemetryNiaInventoryFabricAllOf
	isSet bool
}

func (v NullableNiatelemetryNiaInventoryFabricAllOf) Get() *NiatelemetryNiaInventoryFabricAllOf {
	return v.value
}

func (v *NullableNiatelemetryNiaInventoryFabricAllOf) Set(val *NiatelemetryNiaInventoryFabricAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaInventoryFabricAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaInventoryFabricAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaInventoryFabricAllOf(val *NiatelemetryNiaInventoryFabricAllOf) *NullableNiatelemetryNiaInventoryFabricAllOf {
	return &NullableNiatelemetryNiaInventoryFabricAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaInventoryFabricAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaInventoryFabricAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
