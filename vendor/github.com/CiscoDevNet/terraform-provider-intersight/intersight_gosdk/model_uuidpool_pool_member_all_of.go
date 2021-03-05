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

// UuidpoolPoolMemberAllOf Definition of the list of properties defined in 'uuidpool.PoolMember', excluding properties defined in parent classes.
type UuidpoolPoolMemberAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UUID Prefix+Suffix of this PoolMember.
	Uuid                 *string                        `json:"Uuid,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	BlockHead            *UuidpoolBlockRelationship     `json:"BlockHead,omitempty"`
	Peer                 *UuidpoolUuidLeaseRelationship `json:"Peer,omitempty"`
	Pool                 *UuidpoolPoolRelationship      `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolPoolMemberAllOf UuidpoolPoolMemberAllOf

// NewUuidpoolPoolMemberAllOf instantiates a new UuidpoolPoolMemberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolPoolMemberAllOf(classId string, objectType string) *UuidpoolPoolMemberAllOf {
	this := UuidpoolPoolMemberAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUuidpoolPoolMemberAllOfWithDefaults instantiates a new UuidpoolPoolMemberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolPoolMemberAllOfWithDefaults() *UuidpoolPoolMemberAllOf {
	this := UuidpoolPoolMemberAllOf{}
	var classId string = "uuidpool.PoolMember"
	this.ClassId = classId
	var objectType string = "uuidpool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolPoolMemberAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMemberAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolPoolMemberAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolPoolMemberAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMemberAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolPoolMemberAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UuidpoolPoolMemberAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMemberAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UuidpoolPoolMemberAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UuidpoolPoolMemberAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *UuidpoolPoolMemberAllOf) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMemberAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *UuidpoolPoolMemberAllOf) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *UuidpoolPoolMemberAllOf) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise.
func (o *UuidpoolPoolMemberAllOf) GetBlockHead() UuidpoolBlockRelationship {
	if o == nil || o.BlockHead == nil {
		var ret UuidpoolBlockRelationship
		return ret
	}
	return *o.BlockHead
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMemberAllOf) GetBlockHeadOk() (*UuidpoolBlockRelationship, bool) {
	if o == nil || o.BlockHead == nil {
		return nil, false
	}
	return o.BlockHead, true
}

// HasBlockHead returns a boolean if a field has been set.
func (o *UuidpoolPoolMemberAllOf) HasBlockHead() bool {
	if o != nil && o.BlockHead != nil {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given UuidpoolBlockRelationship and assigns it to the BlockHead field.
func (o *UuidpoolPoolMemberAllOf) SetBlockHead(v UuidpoolBlockRelationship) {
	o.BlockHead = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *UuidpoolPoolMemberAllOf) GetPeer() UuidpoolUuidLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret UuidpoolUuidLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMemberAllOf) GetPeerOk() (*UuidpoolUuidLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *UuidpoolPoolMemberAllOf) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given UuidpoolUuidLeaseRelationship and assigns it to the Peer field.
func (o *UuidpoolPoolMemberAllOf) SetPeer(v UuidpoolUuidLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *UuidpoolPoolMemberAllOf) GetPool() UuidpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMemberAllOf) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolPoolMemberAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolPoolMemberAllOf) SetPool(v UuidpoolPoolRelationship) {
	o.Pool = &v
}

func (o UuidpoolPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.BlockHead != nil {
		toSerialize["BlockHead"] = o.BlockHead
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UuidpoolPoolMemberAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varUuidpoolPoolMemberAllOf := _UuidpoolPoolMemberAllOf{}

	if err = json.Unmarshal(bytes, &varUuidpoolPoolMemberAllOf); err == nil {
		*o = UuidpoolPoolMemberAllOf(varUuidpoolPoolMemberAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUuidpoolPoolMemberAllOf struct {
	value *UuidpoolPoolMemberAllOf
	isSet bool
}

func (v NullableUuidpoolPoolMemberAllOf) Get() *UuidpoolPoolMemberAllOf {
	return v.value
}

func (v *NullableUuidpoolPoolMemberAllOf) Set(val *UuidpoolPoolMemberAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolPoolMemberAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolPoolMemberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolPoolMemberAllOf(val *UuidpoolPoolMemberAllOf) *NullableUuidpoolPoolMemberAllOf {
	return &NullableUuidpoolPoolMemberAllOf{value: val, isSet: true}
}

func (v NullableUuidpoolPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolPoolMemberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
