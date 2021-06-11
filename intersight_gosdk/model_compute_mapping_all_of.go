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

// ComputeMappingAllOf Definition of the list of properties defined in 'compute.Mapping', excluding properties defined in parent classes.
type ComputeMappingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Remote image location from where the image is uploaded to server.
	FileLocation *string `json:"FileLocation,omitempty"`
	// The identity assigned to this Virtual Media Image by server.
	Identifier *string `json:"Identifier,omitempty"`
	// Image name of uploaded Virtual Media Image.
	ImageName  *string  `json:"ImageName,omitempty"`
	MediaTypes []string `json:"MediaTypes,omitempty"`
	// Name of Virtual Media mapping assigne by server.
	Name                 *string                              `json:"Name,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Vmedia               *ComputeVmediaRelationship           `json:"Vmedia,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeMappingAllOf ComputeMappingAllOf

// NewComputeMappingAllOf instantiates a new ComputeMappingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeMappingAllOf(classId string, objectType string) *ComputeMappingAllOf {
	this := ComputeMappingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeMappingAllOfWithDefaults instantiates a new ComputeMappingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeMappingAllOfWithDefaults() *ComputeMappingAllOf {
	this := ComputeMappingAllOf{}
	var classId string = "compute.Mapping"
	this.ClassId = classId
	var objectType string = "compute.Mapping"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeMappingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeMappingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeMappingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeMappingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *ComputeMappingAllOf) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *ComputeMappingAllOf) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ComputeMappingAllOf) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *ComputeMappingAllOf) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *ComputeMappingAllOf) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *ComputeMappingAllOf) SetImageName(v string) {
	o.ImageName = &v
}

// GetMediaTypes returns the MediaTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeMappingAllOf) GetMediaTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MediaTypes
}

// GetMediaTypesOk returns a tuple with the MediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeMappingAllOf) GetMediaTypesOk() (*[]string, bool) {
	if o == nil || o.MediaTypes == nil {
		return nil, false
	}
	return &o.MediaTypes, true
}

// HasMediaTypes returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasMediaTypes() bool {
	if o != nil && o.MediaTypes != nil {
		return true
	}

	return false
}

// SetMediaTypes gets a reference to the given []string and assigns it to the MediaTypes field.
func (o *ComputeMappingAllOf) SetMediaTypes(v []string) {
	o.MediaTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeMappingAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeMappingAllOf) SetName(v string) {
	o.Name = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ComputeMappingAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeMappingAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeMappingAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeMappingAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVmedia returns the Vmedia field value if set, zero value otherwise.
func (o *ComputeMappingAllOf) GetVmedia() ComputeVmediaRelationship {
	if o == nil || o.Vmedia == nil {
		var ret ComputeVmediaRelationship
		return ret
	}
	return *o.Vmedia
}

// GetVmediaOk returns a tuple with the Vmedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMappingAllOf) GetVmediaOk() (*ComputeVmediaRelationship, bool) {
	if o == nil || o.Vmedia == nil {
		return nil, false
	}
	return o.Vmedia, true
}

// HasVmedia returns a boolean if a field has been set.
func (o *ComputeMappingAllOf) HasVmedia() bool {
	if o != nil && o.Vmedia != nil {
		return true
	}

	return false
}

// SetVmedia gets a reference to the given ComputeVmediaRelationship and assigns it to the Vmedia field.
func (o *ComputeMappingAllOf) SetVmedia(v ComputeVmediaRelationship) {
	o.Vmedia = &v
}

func (o ComputeMappingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.Identifier != nil {
		toSerialize["Identifier"] = o.Identifier
	}
	if o.ImageName != nil {
		toSerialize["ImageName"] = o.ImageName
	}
	if o.MediaTypes != nil {
		toSerialize["MediaTypes"] = o.MediaTypes
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Vmedia != nil {
		toSerialize["Vmedia"] = o.Vmedia
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeMappingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeMappingAllOf := _ComputeMappingAllOf{}

	if err = json.Unmarshal(bytes, &varComputeMappingAllOf); err == nil {
		*o = ComputeMappingAllOf(varComputeMappingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "Identifier")
		delete(additionalProperties, "ImageName")
		delete(additionalProperties, "MediaTypes")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Vmedia")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeMappingAllOf struct {
	value *ComputeMappingAllOf
	isSet bool
}

func (v NullableComputeMappingAllOf) Get() *ComputeMappingAllOf {
	return v.value
}

func (v *NullableComputeMappingAllOf) Set(val *ComputeMappingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeMappingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeMappingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeMappingAllOf(val *ComputeMappingAllOf) *NullableComputeMappingAllOf {
	return &NullableComputeMappingAllOf{value: val, isSet: true}
}

func (v NullableComputeMappingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeMappingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
