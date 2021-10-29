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

// HyperflexDataProtectionPeerAllOf Definition of the list of properties defined in 'hyperflex.DataProtectionPeer', excluding properties defined in parent classes.
type HyperflexDataProtectionPeerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Er                   NullableHyperflexEntityReference     `json:"Er,omitempty"`
	PeerInfo             NullableHyperflexReplicationPeerInfo `json:"PeerInfo,omitempty"`
	SrcCluster           *HyperflexClusterRelationship        `json:"SrcCluster,omitempty"`
	TgtCluster           *HyperflexClusterRelationship        `json:"TgtCluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDataProtectionPeerAllOf HyperflexDataProtectionPeerAllOf

// NewHyperflexDataProtectionPeerAllOf instantiates a new HyperflexDataProtectionPeerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDataProtectionPeerAllOf(classId string, objectType string) *HyperflexDataProtectionPeerAllOf {
	this := HyperflexDataProtectionPeerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexDataProtectionPeerAllOfWithDefaults instantiates a new HyperflexDataProtectionPeerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDataProtectionPeerAllOfWithDefaults() *HyperflexDataProtectionPeerAllOf {
	this := HyperflexDataProtectionPeerAllOf{}
	var classId string = "hyperflex.DataProtectionPeer"
	this.ClassId = classId
	var objectType string = "hyperflex.DataProtectionPeer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDataProtectionPeerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDataProtectionPeerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDataProtectionPeerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDataProtectionPeerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEr returns the Er field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexDataProtectionPeerAllOf) GetEr() HyperflexEntityReference {
	if o == nil || o.Er.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.Er.Get()
}

// GetErOk returns a tuple with the Er field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexDataProtectionPeerAllOf) GetErOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Er.Get(), o.Er.IsSet()
}

// HasEr returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeerAllOf) HasEr() bool {
	if o != nil && o.Er.IsSet() {
		return true
	}

	return false
}

// SetEr gets a reference to the given NullableHyperflexEntityReference and assigns it to the Er field.
func (o *HyperflexDataProtectionPeerAllOf) SetEr(v HyperflexEntityReference) {
	o.Er.Set(&v)
}

// SetErNil sets the value for Er to be an explicit nil
func (o *HyperflexDataProtectionPeerAllOf) SetErNil() {
	o.Er.Set(nil)
}

// UnsetEr ensures that no value is present for Er, not even an explicit nil
func (o *HyperflexDataProtectionPeerAllOf) UnsetEr() {
	o.Er.Unset()
}

// GetPeerInfo returns the PeerInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexDataProtectionPeerAllOf) GetPeerInfo() HyperflexReplicationPeerInfo {
	if o == nil || o.PeerInfo.Get() == nil {
		var ret HyperflexReplicationPeerInfo
		return ret
	}
	return *o.PeerInfo.Get()
}

// GetPeerInfoOk returns a tuple with the PeerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexDataProtectionPeerAllOf) GetPeerInfoOk() (*HyperflexReplicationPeerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeerInfo.Get(), o.PeerInfo.IsSet()
}

// HasPeerInfo returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeerAllOf) HasPeerInfo() bool {
	if o != nil && o.PeerInfo.IsSet() {
		return true
	}

	return false
}

// SetPeerInfo gets a reference to the given NullableHyperflexReplicationPeerInfo and assigns it to the PeerInfo field.
func (o *HyperflexDataProtectionPeerAllOf) SetPeerInfo(v HyperflexReplicationPeerInfo) {
	o.PeerInfo.Set(&v)
}

// SetPeerInfoNil sets the value for PeerInfo to be an explicit nil
func (o *HyperflexDataProtectionPeerAllOf) SetPeerInfoNil() {
	o.PeerInfo.Set(nil)
}

// UnsetPeerInfo ensures that no value is present for PeerInfo, not even an explicit nil
func (o *HyperflexDataProtectionPeerAllOf) UnsetPeerInfo() {
	o.PeerInfo.Unset()
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise.
func (o *HyperflexDataProtectionPeerAllOf) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || o.SrcCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeerAllOf) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SrcCluster == nil {
		return nil, false
	}
	return o.SrcCluster, true
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeerAllOf) HasSrcCluster() bool {
	if o != nil && o.SrcCluster != nil {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexDataProtectionPeerAllOf) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster = &v
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise.
func (o *HyperflexDataProtectionPeerAllOf) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || o.TgtCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDataProtectionPeerAllOf) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.TgtCluster == nil {
		return nil, false
	}
	return o.TgtCluster, true
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexDataProtectionPeerAllOf) HasTgtCluster() bool {
	if o != nil && o.TgtCluster != nil {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexDataProtectionPeerAllOf) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster = &v
}

func (o HyperflexDataProtectionPeerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Er.IsSet() {
		toSerialize["Er"] = o.Er.Get()
	}
	if o.PeerInfo.IsSet() {
		toSerialize["PeerInfo"] = o.PeerInfo.Get()
	}
	if o.SrcCluster != nil {
		toSerialize["SrcCluster"] = o.SrcCluster
	}
	if o.TgtCluster != nil {
		toSerialize["TgtCluster"] = o.TgtCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDataProtectionPeerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexDataProtectionPeerAllOf := _HyperflexDataProtectionPeerAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexDataProtectionPeerAllOf); err == nil {
		*o = HyperflexDataProtectionPeerAllOf(varHyperflexDataProtectionPeerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Er")
		delete(additionalProperties, "PeerInfo")
		delete(additionalProperties, "SrcCluster")
		delete(additionalProperties, "TgtCluster")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexDataProtectionPeerAllOf struct {
	value *HyperflexDataProtectionPeerAllOf
	isSet bool
}

func (v NullableHyperflexDataProtectionPeerAllOf) Get() *HyperflexDataProtectionPeerAllOf {
	return v.value
}

func (v *NullableHyperflexDataProtectionPeerAllOf) Set(val *HyperflexDataProtectionPeerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDataProtectionPeerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDataProtectionPeerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDataProtectionPeerAllOf(val *HyperflexDataProtectionPeerAllOf) *NullableHyperflexDataProtectionPeerAllOf {
	return &NullableHyperflexDataProtectionPeerAllOf{value: val, isSet: true}
}

func (v NullableHyperflexDataProtectionPeerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDataProtectionPeerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
