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
	"reflect"
	"strings"
)

// MemoryPersistentMemoryRegion Persistent Memory Region configured on the Persistent Memory Modules on a server.
type MemoryPersistentMemoryRegion struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Free capacity in GiB of the Persistent Memory Region.
	FreeCapacity *string `json:"FreeCapacity,omitempty"`
	// Health state of the Persistent Memory Region.
	HealthState *string `json:"HealthState,omitempty"`
	// ID of the Interleaved Set formed for this Persistent Memory Region.
	InterleavedSetId *string `json:"InterleavedSetId,omitempty"`
	// Set of locator IDs that are included in the Persistent Memory Region.
	LocaterIds *string `json:"LocaterIds,omitempty"`
	// Persistent Memory type of the Persistent Memory Region.
	PersistentMemoryType *string `json:"PersistentMemoryType,omitempty"`
	// ID of the Persistent Memory Region.
	RegionId *string `json:"RegionId,omitempty"`
	// Socket ID of the Persistent Memory Region.
	SocketId *string `json:"SocketId,omitempty"`
	// Socket Memory ID of the Persistent Memory Region.
	SocketMemoryId *string `json:"SocketMemoryId,omitempty"`
	// Total capacity in GiB of the Persistent Memory Region.
	TotalCapacity                       *string                                          `json:"TotalCapacity,omitempty"`
	InventoryDeviceInfo                 *InventoryDeviceInfoRelationship                 `json:"InventoryDeviceInfo,omitempty"`
	MemoryPersistentMemoryConfiguration *MemoryPersistentMemoryConfigurationRelationship `json:"MemoryPersistentMemoryConfiguration,omitempty"`
	// An array of relationships to memoryPersistentMemoryNamespace resources.
	PersistentMemoryNamespaces []MemoryPersistentMemoryNamespaceRelationship `json:"PersistentMemoryNamespaces,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship          `json:"RegisteredDevice,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _MemoryPersistentMemoryRegion MemoryPersistentMemoryRegion

// NewMemoryPersistentMemoryRegion instantiates a new MemoryPersistentMemoryRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryRegion(classId string, objectType string) *MemoryPersistentMemoryRegion {
	this := MemoryPersistentMemoryRegion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryRegionWithDefaults instantiates a new MemoryPersistentMemoryRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryRegionWithDefaults() *MemoryPersistentMemoryRegion {
	this := MemoryPersistentMemoryRegion{}
	var classId string = "memory.PersistentMemoryRegion"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryRegion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryRegion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryRegion) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryRegion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryRegion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFreeCapacity returns the FreeCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetFreeCapacity() string {
	if o == nil || o.FreeCapacity == nil {
		var ret string
		return ret
	}
	return *o.FreeCapacity
}

// GetFreeCapacityOk returns a tuple with the FreeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetFreeCapacityOk() (*string, bool) {
	if o == nil || o.FreeCapacity == nil {
		return nil, false
	}
	return o.FreeCapacity, true
}

// HasFreeCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasFreeCapacity() bool {
	if o != nil && o.FreeCapacity != nil {
		return true
	}

	return false
}

// SetFreeCapacity gets a reference to the given string and assigns it to the FreeCapacity field.
func (o *MemoryPersistentMemoryRegion) SetFreeCapacity(v string) {
	o.FreeCapacity = &v
}

// GetHealthState returns the HealthState field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetHealthState() string {
	if o == nil || o.HealthState == nil {
		var ret string
		return ret
	}
	return *o.HealthState
}

// GetHealthStateOk returns a tuple with the HealthState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetHealthStateOk() (*string, bool) {
	if o == nil || o.HealthState == nil {
		return nil, false
	}
	return o.HealthState, true
}

// HasHealthState returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasHealthState() bool {
	if o != nil && o.HealthState != nil {
		return true
	}

	return false
}

// SetHealthState gets a reference to the given string and assigns it to the HealthState field.
func (o *MemoryPersistentMemoryRegion) SetHealthState(v string) {
	o.HealthState = &v
}

// GetInterleavedSetId returns the InterleavedSetId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetInterleavedSetId() string {
	if o == nil || o.InterleavedSetId == nil {
		var ret string
		return ret
	}
	return *o.InterleavedSetId
}

// GetInterleavedSetIdOk returns a tuple with the InterleavedSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetInterleavedSetIdOk() (*string, bool) {
	if o == nil || o.InterleavedSetId == nil {
		return nil, false
	}
	return o.InterleavedSetId, true
}

// HasInterleavedSetId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasInterleavedSetId() bool {
	if o != nil && o.InterleavedSetId != nil {
		return true
	}

	return false
}

// SetInterleavedSetId gets a reference to the given string and assigns it to the InterleavedSetId field.
func (o *MemoryPersistentMemoryRegion) SetInterleavedSetId(v string) {
	o.InterleavedSetId = &v
}

// GetLocaterIds returns the LocaterIds field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetLocaterIds() string {
	if o == nil || o.LocaterIds == nil {
		var ret string
		return ret
	}
	return *o.LocaterIds
}

// GetLocaterIdsOk returns a tuple with the LocaterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetLocaterIdsOk() (*string, bool) {
	if o == nil || o.LocaterIds == nil {
		return nil, false
	}
	return o.LocaterIds, true
}

// HasLocaterIds returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasLocaterIds() bool {
	if o != nil && o.LocaterIds != nil {
		return true
	}

	return false
}

// SetLocaterIds gets a reference to the given string and assigns it to the LocaterIds field.
func (o *MemoryPersistentMemoryRegion) SetLocaterIds(v string) {
	o.LocaterIds = &v
}

// GetPersistentMemoryType returns the PersistentMemoryType field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryType() string {
	if o == nil || o.PersistentMemoryType == nil {
		var ret string
		return ret
	}
	return *o.PersistentMemoryType
}

// GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryTypeOk() (*string, bool) {
	if o == nil || o.PersistentMemoryType == nil {
		return nil, false
	}
	return o.PersistentMemoryType, true
}

// HasPersistentMemoryType returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasPersistentMemoryType() bool {
	if o != nil && o.PersistentMemoryType != nil {
		return true
	}

	return false
}

// SetPersistentMemoryType gets a reference to the given string and assigns it to the PersistentMemoryType field.
func (o *MemoryPersistentMemoryRegion) SetPersistentMemoryType(v string) {
	o.PersistentMemoryType = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *MemoryPersistentMemoryRegion) SetRegionId(v string) {
	o.RegionId = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryRegion) SetSocketId(v string) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetSocketMemoryId() string {
	if o == nil || o.SocketMemoryId == nil {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || o.SocketMemoryId == nil {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasSocketMemoryId() bool {
	if o != nil && o.SocketMemoryId != nil {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *MemoryPersistentMemoryRegion) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetTotalCapacity() string {
	if o == nil || o.TotalCapacity == nil {
		var ret string
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetTotalCapacityOk() (*string, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given string and assigns it to the TotalCapacity field.
func (o *MemoryPersistentMemoryRegion) SetTotalCapacity(v string) {
	o.TotalCapacity = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryRegion) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship {
	if o == nil || o.MemoryPersistentMemoryConfiguration == nil {
		var ret MemoryPersistentMemoryConfigurationRelationship
		return ret
	}
	return *o.MemoryPersistentMemoryConfiguration
}

// GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool) {
	if o == nil || o.MemoryPersistentMemoryConfiguration == nil {
		return nil, false
	}
	return o.MemoryPersistentMemoryConfiguration, true
}

// HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasMemoryPersistentMemoryConfiguration() bool {
	if o != nil && o.MemoryPersistentMemoryConfiguration != nil {
		return true
	}

	return false
}

// SetMemoryPersistentMemoryConfiguration gets a reference to the given MemoryPersistentMemoryConfigurationRelationship and assigns it to the MemoryPersistentMemoryConfiguration field.
func (o *MemoryPersistentMemoryRegion) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship) {
	o.MemoryPersistentMemoryConfiguration = &v
}

// GetPersistentMemoryNamespaces returns the PersistentMemoryNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryNamespaces() []MemoryPersistentMemoryNamespaceRelationship {
	if o == nil {
		var ret []MemoryPersistentMemoryNamespaceRelationship
		return ret
	}
	return o.PersistentMemoryNamespaces
}

// GetPersistentMemoryNamespacesOk returns a tuple with the PersistentMemoryNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryNamespacesOk() (*[]MemoryPersistentMemoryNamespaceRelationship, bool) {
	if o == nil || o.PersistentMemoryNamespaces == nil {
		return nil, false
	}
	return &o.PersistentMemoryNamespaces, true
}

// HasPersistentMemoryNamespaces returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasPersistentMemoryNamespaces() bool {
	if o != nil && o.PersistentMemoryNamespaces != nil {
		return true
	}

	return false
}

// SetPersistentMemoryNamespaces gets a reference to the given []MemoryPersistentMemoryNamespaceRelationship and assigns it to the PersistentMemoryNamespaces field.
func (o *MemoryPersistentMemoryRegion) SetPersistentMemoryNamespaces(v []MemoryPersistentMemoryNamespaceRelationship) {
	o.PersistentMemoryNamespaces = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegion) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegion) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegion) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryRegion) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o MemoryPersistentMemoryRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FreeCapacity != nil {
		toSerialize["FreeCapacity"] = o.FreeCapacity
	}
	if o.HealthState != nil {
		toSerialize["HealthState"] = o.HealthState
	}
	if o.InterleavedSetId != nil {
		toSerialize["InterleavedSetId"] = o.InterleavedSetId
	}
	if o.LocaterIds != nil {
		toSerialize["LocaterIds"] = o.LocaterIds
	}
	if o.PersistentMemoryType != nil {
		toSerialize["PersistentMemoryType"] = o.PersistentMemoryType
	}
	if o.RegionId != nil {
		toSerialize["RegionId"] = o.RegionId
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	if o.SocketMemoryId != nil {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}
	if o.TotalCapacity != nil {
		toSerialize["TotalCapacity"] = o.TotalCapacity
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryPersistentMemoryConfiguration != nil {
		toSerialize["MemoryPersistentMemoryConfiguration"] = o.MemoryPersistentMemoryConfiguration
	}
	if o.PersistentMemoryNamespaces != nil {
		toSerialize["PersistentMemoryNamespaces"] = o.PersistentMemoryNamespaces
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryRegion) UnmarshalJSON(bytes []byte) (err error) {
	type MemoryPersistentMemoryRegionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Free capacity in GiB of the Persistent Memory Region.
		FreeCapacity *string `json:"FreeCapacity,omitempty"`
		// Health state of the Persistent Memory Region.
		HealthState *string `json:"HealthState,omitempty"`
		// ID of the Interleaved Set formed for this Persistent Memory Region.
		InterleavedSetId *string `json:"InterleavedSetId,omitempty"`
		// Set of locator IDs that are included in the Persistent Memory Region.
		LocaterIds *string `json:"LocaterIds,omitempty"`
		// Persistent Memory type of the Persistent Memory Region.
		PersistentMemoryType *string `json:"PersistentMemoryType,omitempty"`
		// ID of the Persistent Memory Region.
		RegionId *string `json:"RegionId,omitempty"`
		// Socket ID of the Persistent Memory Region.
		SocketId *string `json:"SocketId,omitempty"`
		// Socket Memory ID of the Persistent Memory Region.
		SocketMemoryId *string `json:"SocketMemoryId,omitempty"`
		// Total capacity in GiB of the Persistent Memory Region.
		TotalCapacity                       *string                                          `json:"TotalCapacity,omitempty"`
		InventoryDeviceInfo                 *InventoryDeviceInfoRelationship                 `json:"InventoryDeviceInfo,omitempty"`
		MemoryPersistentMemoryConfiguration *MemoryPersistentMemoryConfigurationRelationship `json:"MemoryPersistentMemoryConfiguration,omitempty"`
		// An array of relationships to memoryPersistentMemoryNamespace resources.
		PersistentMemoryNamespaces []MemoryPersistentMemoryNamespaceRelationship `json:"PersistentMemoryNamespaces,omitempty"`
		RegisteredDevice           *AssetDeviceRegistrationRelationship          `json:"RegisteredDevice,omitempty"`
	}

	varMemoryPersistentMemoryRegionWithoutEmbeddedStruct := MemoryPersistentMemoryRegionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryRegionWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryRegion := _MemoryPersistentMemoryRegion{}
		varMemoryPersistentMemoryRegion.ClassId = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryRegion.ObjectType = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryRegion.FreeCapacity = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.FreeCapacity
		varMemoryPersistentMemoryRegion.HealthState = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.HealthState
		varMemoryPersistentMemoryRegion.InterleavedSetId = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.InterleavedSetId
		varMemoryPersistentMemoryRegion.LocaterIds = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.LocaterIds
		varMemoryPersistentMemoryRegion.PersistentMemoryType = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.PersistentMemoryType
		varMemoryPersistentMemoryRegion.RegionId = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.RegionId
		varMemoryPersistentMemoryRegion.SocketId = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.SocketId
		varMemoryPersistentMemoryRegion.SocketMemoryId = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.SocketMemoryId
		varMemoryPersistentMemoryRegion.TotalCapacity = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.TotalCapacity
		varMemoryPersistentMemoryRegion.InventoryDeviceInfo = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryPersistentMemoryRegion.MemoryPersistentMemoryConfiguration = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.MemoryPersistentMemoryConfiguration
		varMemoryPersistentMemoryRegion.PersistentMemoryNamespaces = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.PersistentMemoryNamespaces
		varMemoryPersistentMemoryRegion.RegisteredDevice = varMemoryPersistentMemoryRegionWithoutEmbeddedStruct.RegisteredDevice
		*o = MemoryPersistentMemoryRegion(varMemoryPersistentMemoryRegion)
	} else {
		return err
	}

	varMemoryPersistentMemoryRegion := _MemoryPersistentMemoryRegion{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryRegion)
	if err == nil {
		o.InventoryBase = varMemoryPersistentMemoryRegion.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FreeCapacity")
		delete(additionalProperties, "HealthState")
		delete(additionalProperties, "InterleavedSetId")
		delete(additionalProperties, "LocaterIds")
		delete(additionalProperties, "PersistentMemoryType")
		delete(additionalProperties, "RegionId")
		delete(additionalProperties, "SocketId")
		delete(additionalProperties, "SocketMemoryId")
		delete(additionalProperties, "TotalCapacity")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryPersistentMemoryConfiguration")
		delete(additionalProperties, "PersistentMemoryNamespaces")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryPersistentMemoryRegion struct {
	value *MemoryPersistentMemoryRegion
	isSet bool
}

func (v NullableMemoryPersistentMemoryRegion) Get() *MemoryPersistentMemoryRegion {
	return v.value
}

func (v *NullableMemoryPersistentMemoryRegion) Set(val *MemoryPersistentMemoryRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryRegion(val *MemoryPersistentMemoryRegion) *NullableMemoryPersistentMemoryRegion {
	return &NullableMemoryPersistentMemoryRegion{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
