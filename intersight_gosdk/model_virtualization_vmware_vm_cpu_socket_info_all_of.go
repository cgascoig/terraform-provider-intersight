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

// VirtualizationVmwareVmCpuSocketInfoAllOf Definition of the list of properties defined in 'virtualization.VmwareVmCpuSocketInfo', excluding properties defined in parent classes.
type VirtualizationVmwareVmCpuSocketInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of core per CPU socket (value may be empty).
	CoresPerSocket *int64 `json:"CoresPerSocket,omitempty"`
	// Number of CPUs allocated to this VM.
	NumCpus *int64 `json:"NumCpus,omitempty"`
	// The number of CPU sockets allocated.
	NumSockets           *int64 `json:"NumSockets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVmCpuSocketInfoAllOf VirtualizationVmwareVmCpuSocketInfoAllOf

// NewVirtualizationVmwareVmCpuSocketInfoAllOf instantiates a new VirtualizationVmwareVmCpuSocketInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmCpuSocketInfoAllOf(classId string, objectType string) *VirtualizationVmwareVmCpuSocketInfoAllOf {
	this := VirtualizationVmwareVmCpuSocketInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVmCpuSocketInfoAllOfWithDefaults instantiates a new VirtualizationVmwareVmCpuSocketInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmCpuSocketInfoAllOfWithDefaults() *VirtualizationVmwareVmCpuSocketInfoAllOf {
	this := VirtualizationVmwareVmCpuSocketInfoAllOf{}
	var classId string = "virtualization.VmwareVmCpuSocketInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVmCpuSocketInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCoresPerSocket returns the CoresPerSocket field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetCoresPerSocket() int64 {
	if o == nil || o.CoresPerSocket == nil {
		var ret int64
		return ret
	}
	return *o.CoresPerSocket
}

// GetCoresPerSocketOk returns a tuple with the CoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetCoresPerSocketOk() (*int64, bool) {
	if o == nil || o.CoresPerSocket == nil {
		return nil, false
	}
	return o.CoresPerSocket, true
}

// HasCoresPerSocket returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) HasCoresPerSocket() bool {
	if o != nil && o.CoresPerSocket != nil {
		return true
	}

	return false
}

// SetCoresPerSocket gets a reference to the given int64 and assigns it to the CoresPerSocket field.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetCoresPerSocket(v int64) {
	o.CoresPerSocket = &v
}

// GetNumCpus returns the NumCpus field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumCpus() int64 {
	if o == nil || o.NumCpus == nil {
		var ret int64
		return ret
	}
	return *o.NumCpus
}

// GetNumCpusOk returns a tuple with the NumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumCpusOk() (*int64, bool) {
	if o == nil || o.NumCpus == nil {
		return nil, false
	}
	return o.NumCpus, true
}

// HasNumCpus returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) HasNumCpus() bool {
	if o != nil && o.NumCpus != nil {
		return true
	}

	return false
}

// SetNumCpus gets a reference to the given int64 and assigns it to the NumCpus field.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetNumCpus(v int64) {
	o.NumCpus = &v
}

// GetNumSockets returns the NumSockets field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumSockets() int64 {
	if o == nil || o.NumSockets == nil {
		var ret int64
		return ret
	}
	return *o.NumSockets
}

// GetNumSocketsOk returns a tuple with the NumSockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumSocketsOk() (*int64, bool) {
	if o == nil || o.NumSockets == nil {
		return nil, false
	}
	return o.NumSockets, true
}

// HasNumSockets returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) HasNumSockets() bool {
	if o != nil && o.NumSockets != nil {
		return true
	}

	return false
}

// SetNumSockets gets a reference to the given int64 and assigns it to the NumSockets field.
func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetNumSockets(v int64) {
	o.NumSockets = &v
}

func (o VirtualizationVmwareVmCpuSocketInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CoresPerSocket != nil {
		toSerialize["CoresPerSocket"] = o.CoresPerSocket
	}
	if o.NumCpus != nil {
		toSerialize["NumCpus"] = o.NumCpus
	}
	if o.NumSockets != nil {
		toSerialize["NumSockets"] = o.NumSockets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareVmCpuSocketInfoAllOf := _VirtualizationVmwareVmCpuSocketInfoAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareVmCpuSocketInfoAllOf); err == nil {
		*o = VirtualizationVmwareVmCpuSocketInfoAllOf(varVirtualizationVmwareVmCpuSocketInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CoresPerSocket")
		delete(additionalProperties, "NumCpus")
		delete(additionalProperties, "NumSockets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareVmCpuSocketInfoAllOf struct {
	value *VirtualizationVmwareVmCpuSocketInfoAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareVmCpuSocketInfoAllOf) Get() *VirtualizationVmwareVmCpuSocketInfoAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareVmCpuSocketInfoAllOf) Set(val *VirtualizationVmwareVmCpuSocketInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmCpuSocketInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmCpuSocketInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmCpuSocketInfoAllOf(val *VirtualizationVmwareVmCpuSocketInfoAllOf) *NullableVirtualizationVmwareVmCpuSocketInfoAllOf {
	return &NullableVirtualizationVmwareVmCpuSocketInfoAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmCpuSocketInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmCpuSocketInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
