/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-05-25T18:18:54Z.
 *
 * API version: 1.0.9-4305
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexHxapDvswitchAllOf Definition of the list of properties defined in 'hyperflex.HxapDvswitch', excluding properties defined in parent classes.
type HyperflexHxapDvswitchAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the dvUplink referenced by this dvswitch.
	DvUplink *string `json:"DvUplink,omitempty"`
	// The last host that update this object.
	LastHostname *string                           `json:"LastHostname,omitempty"`
	Cluster      *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to hyperflexHxapHost resources.
	MemberHosts []HyperflexHxapHostRelationship `json:"MemberHosts,omitempty"`
	// An array of relationships to hyperflexHxapHostVswitch resources.
	MemberVswitches      []HyperflexHxapHostVswitchRelationship `json:"MemberVswitches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapDvswitchAllOf HyperflexHxapDvswitchAllOf

// NewHyperflexHxapDvswitchAllOf instantiates a new HyperflexHxapDvswitchAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapDvswitchAllOf(classId string, objectType string) *HyperflexHxapDvswitchAllOf {
	this := HyperflexHxapDvswitchAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxapDvswitchAllOfWithDefaults instantiates a new HyperflexHxapDvswitchAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapDvswitchAllOfWithDefaults() *HyperflexHxapDvswitchAllOf {
	this := HyperflexHxapDvswitchAllOf{}
	var classId string = "hyperflex.HxapDvswitch"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapDvswitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapDvswitchAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitchAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapDvswitchAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapDvswitchAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitchAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapDvswitchAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDvUplink returns the DvUplink field value if set, zero value otherwise.
func (o *HyperflexHxapDvswitchAllOf) GetDvUplink() string {
	if o == nil || o.DvUplink == nil {
		var ret string
		return ret
	}
	return *o.DvUplink
}

// GetDvUplinkOk returns a tuple with the DvUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitchAllOf) GetDvUplinkOk() (*string, bool) {
	if o == nil || o.DvUplink == nil {
		return nil, false
	}
	return o.DvUplink, true
}

// HasDvUplink returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitchAllOf) HasDvUplink() bool {
	if o != nil && o.DvUplink != nil {
		return true
	}

	return false
}

// SetDvUplink gets a reference to the given string and assigns it to the DvUplink field.
func (o *HyperflexHxapDvswitchAllOf) SetDvUplink(v string) {
	o.DvUplink = &v
}

// GetLastHostname returns the LastHostname field value if set, zero value otherwise.
func (o *HyperflexHxapDvswitchAllOf) GetLastHostname() string {
	if o == nil || o.LastHostname == nil {
		var ret string
		return ret
	}
	return *o.LastHostname
}

// GetLastHostnameOk returns a tuple with the LastHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitchAllOf) GetLastHostnameOk() (*string, bool) {
	if o == nil || o.LastHostname == nil {
		return nil, false
	}
	return o.LastHostname, true
}

// HasLastHostname returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitchAllOf) HasLastHostname() bool {
	if o != nil && o.LastHostname != nil {
		return true
	}

	return false
}

// SetLastHostname gets a reference to the given string and assigns it to the LastHostname field.
func (o *HyperflexHxapDvswitchAllOf) SetLastHostname(v string) {
	o.LastHostname = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapDvswitchAllOf) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitchAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitchAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapDvswitchAllOf) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetMemberHosts returns the MemberHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapDvswitchAllOf) GetMemberHosts() []HyperflexHxapHostRelationship {
	if o == nil {
		var ret []HyperflexHxapHostRelationship
		return ret
	}
	return o.MemberHosts
}

// GetMemberHostsOk returns a tuple with the MemberHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapDvswitchAllOf) GetMemberHostsOk() (*[]HyperflexHxapHostRelationship, bool) {
	if o == nil || o.MemberHosts == nil {
		return nil, false
	}
	return &o.MemberHosts, true
}

// HasMemberHosts returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitchAllOf) HasMemberHosts() bool {
	if o != nil && o.MemberHosts != nil {
		return true
	}

	return false
}

// SetMemberHosts gets a reference to the given []HyperflexHxapHostRelationship and assigns it to the MemberHosts field.
func (o *HyperflexHxapDvswitchAllOf) SetMemberHosts(v []HyperflexHxapHostRelationship) {
	o.MemberHosts = v
}

// GetMemberVswitches returns the MemberVswitches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapDvswitchAllOf) GetMemberVswitches() []HyperflexHxapHostVswitchRelationship {
	if o == nil {
		var ret []HyperflexHxapHostVswitchRelationship
		return ret
	}
	return o.MemberVswitches
}

// GetMemberVswitchesOk returns a tuple with the MemberVswitches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapDvswitchAllOf) GetMemberVswitchesOk() (*[]HyperflexHxapHostVswitchRelationship, bool) {
	if o == nil || o.MemberVswitches == nil {
		return nil, false
	}
	return &o.MemberVswitches, true
}

// HasMemberVswitches returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitchAllOf) HasMemberVswitches() bool {
	if o != nil && o.MemberVswitches != nil {
		return true
	}

	return false
}

// SetMemberVswitches gets a reference to the given []HyperflexHxapHostVswitchRelationship and assigns it to the MemberVswitches field.
func (o *HyperflexHxapDvswitchAllOf) SetMemberVswitches(v []HyperflexHxapHostVswitchRelationship) {
	o.MemberVswitches = v
}

func (o HyperflexHxapDvswitchAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DvUplink != nil {
		toSerialize["DvUplink"] = o.DvUplink
	}
	if o.LastHostname != nil {
		toSerialize["LastHostname"] = o.LastHostname
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.MemberHosts != nil {
		toSerialize["MemberHosts"] = o.MemberHosts
	}
	if o.MemberVswitches != nil {
		toSerialize["MemberVswitches"] = o.MemberVswitches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapDvswitchAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxapDvswitchAllOf := _HyperflexHxapDvswitchAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxapDvswitchAllOf); err == nil {
		*o = HyperflexHxapDvswitchAllOf(varHyperflexHxapDvswitchAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DvUplink")
		delete(additionalProperties, "LastHostname")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "MemberHosts")
		delete(additionalProperties, "MemberVswitches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxapDvswitchAllOf struct {
	value *HyperflexHxapDvswitchAllOf
	isSet bool
}

func (v NullableHyperflexHxapDvswitchAllOf) Get() *HyperflexHxapDvswitchAllOf {
	return v.value
}

func (v *NullableHyperflexHxapDvswitchAllOf) Set(val *HyperflexHxapDvswitchAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapDvswitchAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapDvswitchAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapDvswitchAllOf(val *HyperflexHxapDvswitchAllOf) *NullableHyperflexHxapDvswitchAllOf {
	return &NullableHyperflexHxapDvswitchAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxapDvswitchAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapDvswitchAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
