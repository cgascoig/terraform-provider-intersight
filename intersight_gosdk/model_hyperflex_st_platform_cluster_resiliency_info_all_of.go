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

// HyperflexStPlatformClusterResiliencyInfoAllOf Definition of the list of properties defined in 'hyperflex.StPlatformClusterResiliencyInfo', excluding properties defined in parent classes.
type HyperflexStPlatformClusterResiliencyInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of persistent storage device failures tolerable before the storage cluster becomes offline.
	HddFailuresTolerable *int64   `json:"HddFailuresTolerable,omitempty"`
	Messages             []string `json:"Messages,omitempty"`
	// The current message describing the auto-healing process of the cluster.
	MessagesIterator interface{} `json:"MessagesIterator,omitempty"`
	// The number of elements in the messages collection.
	MessagesSize *int64 `json:"MessagesSize,omitempty"`
	// The number of node failures tolerable before the storage cluster becomes offline.
	NodeFailuresTolerable *int64 `json:"NodeFailuresTolerable,omitempty"`
	// The number of caching device failures tolerable before the storage cluster becomes offline.
	SsdFailuresTolerable *int64 `json:"SsdFailuresTolerable,omitempty"`
	// The resiliency state of the cluster. The resiliency status is 'HEALTHY' if there are no failures and the storage cluster is fully operational. The resiliency status is 'WARNING' when the cluster has experienced failures that may adversely affect the cluster. It is 'UNKNOWN' if the cluster is offline or if the state cannot be determined.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexStPlatformClusterResiliencyInfoAllOf HyperflexStPlatformClusterResiliencyInfoAllOf

// NewHyperflexStPlatformClusterResiliencyInfoAllOf instantiates a new HyperflexStPlatformClusterResiliencyInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexStPlatformClusterResiliencyInfoAllOf(classId string, objectType string) *HyperflexStPlatformClusterResiliencyInfoAllOf {
	this := HyperflexStPlatformClusterResiliencyInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults instantiates a new HyperflexStPlatformClusterResiliencyInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults() *HyperflexStPlatformClusterResiliencyInfoAllOf {
	this := HyperflexStPlatformClusterResiliencyInfoAllOf{}
	var classId string = "hyperflex.StPlatformClusterResiliencyInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.StPlatformClusterResiliencyInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHddFailuresTolerable returns the HddFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetHddFailuresTolerable() int64 {
	if o == nil || o.HddFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.HddFailuresTolerable
}

// GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetHddFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.HddFailuresTolerable == nil {
		return nil, false
	}
	return o.HddFailuresTolerable, true
}

// HasHddFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasHddFailuresTolerable() bool {
	if o != nil && o.HddFailuresTolerable != nil {
		return true
	}

	return false
}

// SetHddFailuresTolerable gets a reference to the given int64 and assigns it to the HddFailuresTolerable field.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetHddFailuresTolerable(v int64) {
	o.HddFailuresTolerable = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessages(v []string) {
	o.Messages = v
}

// GetMessagesIterator returns the MessagesIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesIterator() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MessagesIterator
}

// GetMessagesIteratorOk returns a tuple with the MessagesIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesIteratorOk() (*interface{}, bool) {
	if o == nil || o.MessagesIterator == nil {
		return nil, false
	}
	return &o.MessagesIterator, true
}

// HasMessagesIterator returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasMessagesIterator() bool {
	if o != nil && o.MessagesIterator != nil {
		return true
	}

	return false
}

// SetMessagesIterator gets a reference to the given interface{} and assigns it to the MessagesIterator field.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessagesIterator(v interface{}) {
	o.MessagesIterator = v
}

// GetMessagesSize returns the MessagesSize field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesSize() int64 {
	if o == nil || o.MessagesSize == nil {
		var ret int64
		return ret
	}
	return *o.MessagesSize
}

// GetMessagesSizeOk returns a tuple with the MessagesSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesSizeOk() (*int64, bool) {
	if o == nil || o.MessagesSize == nil {
		return nil, false
	}
	return o.MessagesSize, true
}

// HasMessagesSize returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasMessagesSize() bool {
	if o != nil && o.MessagesSize != nil {
		return true
	}

	return false
}

// SetMessagesSize gets a reference to the given int64 and assigns it to the MessagesSize field.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessagesSize(v int64) {
	o.MessagesSize = &v
}

// GetNodeFailuresTolerable returns the NodeFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetNodeFailuresTolerable() int64 {
	if o == nil || o.NodeFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.NodeFailuresTolerable
}

// GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetNodeFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.NodeFailuresTolerable == nil {
		return nil, false
	}
	return o.NodeFailuresTolerable, true
}

// HasNodeFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasNodeFailuresTolerable() bool {
	if o != nil && o.NodeFailuresTolerable != nil {
		return true
	}

	return false
}

// SetNodeFailuresTolerable gets a reference to the given int64 and assigns it to the NodeFailuresTolerable field.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetNodeFailuresTolerable(v int64) {
	o.NodeFailuresTolerable = &v
}

// GetSsdFailuresTolerable returns the SsdFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetSsdFailuresTolerable() int64 {
	if o == nil || o.SsdFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.SsdFailuresTolerable
}

// GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetSsdFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.SsdFailuresTolerable == nil {
		return nil, false
	}
	return o.SsdFailuresTolerable, true
}

// HasSsdFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasSsdFailuresTolerable() bool {
	if o != nil && o.SsdFailuresTolerable != nil {
		return true
	}

	return false
}

// SetSsdFailuresTolerable gets a reference to the given int64 and assigns it to the SsdFailuresTolerable field.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetSsdFailuresTolerable(v int64) {
	o.SsdFailuresTolerable = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetState(v string) {
	o.State = &v
}

func (o HyperflexStPlatformClusterResiliencyInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HddFailuresTolerable != nil {
		toSerialize["HddFailuresTolerable"] = o.HddFailuresTolerable
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.MessagesIterator != nil {
		toSerialize["MessagesIterator"] = o.MessagesIterator
	}
	if o.MessagesSize != nil {
		toSerialize["MessagesSize"] = o.MessagesSize
	}
	if o.NodeFailuresTolerable != nil {
		toSerialize["NodeFailuresTolerable"] = o.NodeFailuresTolerable
	}
	if o.SsdFailuresTolerable != nil {
		toSerialize["SsdFailuresTolerable"] = o.SsdFailuresTolerable
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexStPlatformClusterResiliencyInfoAllOf := _HyperflexStPlatformClusterResiliencyInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexStPlatformClusterResiliencyInfoAllOf); err == nil {
		*o = HyperflexStPlatformClusterResiliencyInfoAllOf(varHyperflexStPlatformClusterResiliencyInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HddFailuresTolerable")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "MessagesIterator")
		delete(additionalProperties, "MessagesSize")
		delete(additionalProperties, "NodeFailuresTolerable")
		delete(additionalProperties, "SsdFailuresTolerable")
		delete(additionalProperties, "State")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexStPlatformClusterResiliencyInfoAllOf struct {
	value *HyperflexStPlatformClusterResiliencyInfoAllOf
	isSet bool
}

func (v NullableHyperflexStPlatformClusterResiliencyInfoAllOf) Get() *HyperflexStPlatformClusterResiliencyInfoAllOf {
	return v.value
}

func (v *NullableHyperflexStPlatformClusterResiliencyInfoAllOf) Set(val *HyperflexStPlatformClusterResiliencyInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexStPlatformClusterResiliencyInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexStPlatformClusterResiliencyInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexStPlatformClusterResiliencyInfoAllOf(val *HyperflexStPlatformClusterResiliencyInfoAllOf) *NullableHyperflexStPlatformClusterResiliencyInfoAllOf {
	return &NullableHyperflexStPlatformClusterResiliencyInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexStPlatformClusterResiliencyInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexStPlatformClusterResiliencyInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
