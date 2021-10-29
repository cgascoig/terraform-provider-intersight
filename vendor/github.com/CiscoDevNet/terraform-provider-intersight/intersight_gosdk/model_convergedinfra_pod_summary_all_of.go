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

// ConvergedinfraPodSummaryAllOf Definition of the list of properties defined in 'convergedinfra.PodSummary', excluding properties defined in parent classes.
type ConvergedinfraPodSummaryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of Nodes that are powered on and have at least 1 VM associated with the pod.
	ActiveNodes *int64 `json:"ActiveNodes,omitempty"`
	// Number of VMs associated with the pod.
	VmCount              *int64 `json:"VmCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraPodSummaryAllOf ConvergedinfraPodSummaryAllOf

// NewConvergedinfraPodSummaryAllOf instantiates a new ConvergedinfraPodSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraPodSummaryAllOf(classId string, objectType string) *ConvergedinfraPodSummaryAllOf {
	this := ConvergedinfraPodSummaryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraPodSummaryAllOfWithDefaults instantiates a new ConvergedinfraPodSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraPodSummaryAllOfWithDefaults() *ConvergedinfraPodSummaryAllOf {
	this := ConvergedinfraPodSummaryAllOf{}
	var classId string = "convergedinfra.PodSummary"
	this.ClassId = classId
	var objectType string = "convergedinfra.PodSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraPodSummaryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPodSummaryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraPodSummaryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraPodSummaryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPodSummaryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraPodSummaryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveNodes returns the ActiveNodes field value if set, zero value otherwise.
func (o *ConvergedinfraPodSummaryAllOf) GetActiveNodes() int64 {
	if o == nil || o.ActiveNodes == nil {
		var ret int64
		return ret
	}
	return *o.ActiveNodes
}

// GetActiveNodesOk returns a tuple with the ActiveNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPodSummaryAllOf) GetActiveNodesOk() (*int64, bool) {
	if o == nil || o.ActiveNodes == nil {
		return nil, false
	}
	return o.ActiveNodes, true
}

// HasActiveNodes returns a boolean if a field has been set.
func (o *ConvergedinfraPodSummaryAllOf) HasActiveNodes() bool {
	if o != nil && o.ActiveNodes != nil {
		return true
	}

	return false
}

// SetActiveNodes gets a reference to the given int64 and assigns it to the ActiveNodes field.
func (o *ConvergedinfraPodSummaryAllOf) SetActiveNodes(v int64) {
	o.ActiveNodes = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *ConvergedinfraPodSummaryAllOf) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPodSummaryAllOf) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *ConvergedinfraPodSummaryAllOf) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *ConvergedinfraPodSummaryAllOf) SetVmCount(v int64) {
	o.VmCount = &v
}

func (o ConvergedinfraPodSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActiveNodes != nil {
		toSerialize["ActiveNodes"] = o.ActiveNodes
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraPodSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConvergedinfraPodSummaryAllOf := _ConvergedinfraPodSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varConvergedinfraPodSummaryAllOf); err == nil {
		*o = ConvergedinfraPodSummaryAllOf(varConvergedinfraPodSummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveNodes")
		delete(additionalProperties, "VmCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvergedinfraPodSummaryAllOf struct {
	value *ConvergedinfraPodSummaryAllOf
	isSet bool
}

func (v NullableConvergedinfraPodSummaryAllOf) Get() *ConvergedinfraPodSummaryAllOf {
	return v.value
}

func (v *NullableConvergedinfraPodSummaryAllOf) Set(val *ConvergedinfraPodSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraPodSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraPodSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraPodSummaryAllOf(val *ConvergedinfraPodSummaryAllOf) *NullableConvergedinfraPodSummaryAllOf {
	return &NullableConvergedinfraPodSummaryAllOf{value: val, isSet: true}
}

func (v NullableConvergedinfraPodSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraPodSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
