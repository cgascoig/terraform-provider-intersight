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

// StorageBaseTenantAllOf Definition of the list of properties defined in 'storage.BaseTenant', excluding properties defined in parent classes.
type StorageBaseTenantAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the tenant in storage array.
	Name *string `json:"Name,omitempty"`
	// The state of this tenant. * `Unknown` - Component state is not available. * `Starting` - Component is being started. * `Running` - Component is currently running. * `Stopping` - Component is being stopped. * `Stopped` - Component has been stopped. * `Deleting` - Component deletion is in progress.
	State *string `json:"State,omitempty"`
	// The uuid of this tenant in storage array.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseTenantAllOf StorageBaseTenantAllOf

// NewStorageBaseTenantAllOf instantiates a new StorageBaseTenantAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseTenantAllOf(classId string, objectType string) *StorageBaseTenantAllOf {
	this := StorageBaseTenantAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "Unknown"
	this.State = &state
	return &this
}

// NewStorageBaseTenantAllOfWithDefaults instantiates a new StorageBaseTenantAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseTenantAllOfWithDefaults() *StorageBaseTenantAllOf {
	this := StorageBaseTenantAllOf{}
	var classId string = "storage.NetAppStorageVm"
	this.ClassId = classId
	var objectType string = "storage.NetAppStorageVm"
	this.ObjectType = objectType
	var state string = "Unknown"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseTenantAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseTenantAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseTenantAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseTenantAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseTenantAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseTenantAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseTenantAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseTenantAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseTenantAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseTenantAllOf) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageBaseTenantAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseTenantAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageBaseTenantAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageBaseTenantAllOf) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageBaseTenantAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseTenantAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageBaseTenantAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageBaseTenantAllOf) SetUuid(v string) {
	o.Uuid = &v
}

func (o StorageBaseTenantAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseTenantAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseTenantAllOf := _StorageBaseTenantAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseTenantAllOf); err == nil {
		*o = StorageBaseTenantAllOf(varStorageBaseTenantAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseTenantAllOf struct {
	value *StorageBaseTenantAllOf
	isSet bool
}

func (v NullableStorageBaseTenantAllOf) Get() *StorageBaseTenantAllOf {
	return v.value
}

func (v *NullableStorageBaseTenantAllOf) Set(val *StorageBaseTenantAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseTenantAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseTenantAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseTenantAllOf(val *StorageBaseTenantAllOf) *NullableStorageBaseTenantAllOf {
	return &NullableStorageBaseTenantAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseTenantAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseTenantAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
