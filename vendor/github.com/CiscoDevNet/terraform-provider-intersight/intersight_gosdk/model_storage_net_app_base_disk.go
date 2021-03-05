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
	"reflect"
	"strings"
)

// StorageNetAppBaseDisk NetApp base disk is a storage array disk.
type StorageNetAppBaseDisk struct {
	StorageBaseArrayDisk
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// NetApp base disk model.
	BaseDiskModel *string `json:"BaseDiskModel,omitempty"`
	// Supported container type for NetApp disk. * `Unknown` - Container is currently unknown. This is the default setting. * `Aggregate` - Disk is used as a physical disk in an aggregate. * `Broken` - Disk is in broken pool. * `Label Maintenance` - Disk is in online label maintenance list. * `Foreign` - Array LUN has been marked foreign. * `Maintenance` - Disk is in maintenance center. * `Mediator` - A mediator disk is a disk used on non-shared HA systems hosted by an external node which is used to communicate the viability of the storage failover between non-shared HA nodes. * `Shared` - Disk is partitioned or in a storage pool. * `Remote` - Disk belongs to a remote cluster. * `Spare` - Disk is a spare disk. * `Unassigned` - Disk ownership has not been assigned. * `Unsupported` - Disk is not supported.
	ContainerType *string `json:"ContainerType,omitempty"`
	// Type of the NetApp disk. * `Unknown` - Default unknown disk type. * `SSDNVM` - Solid state disk with Non-Volatile Memory Express protocol enabled. * `ATA` - Advanced Technology Attachment is a type of disk drive that integrates the drive controller directly on the drive itself. * `FCAL` - For the FC-AL disk connection type, disk shelves are connected to the controller in a loop * `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, and rotational speed of traditional enterprise-class SATA drives with the fully capable SAS interface typical for classic SAS drives. * `LUN` - Logical Unit Number refers to a logical disk. * `SAS` - Storage disk with serial attached SCSI. * `MSATA` - SATA disk in multi-disk carrier storage shelf. * `SSD` - Storage disk with Solid state disk. * `VMDISK` - Virtual machine Data Disk.
	DiskType *string `json:"DiskType,omitempty"`
	// Current state of the NetApp disk. * `Present` - Storage disk state type is present. * `Copy` - Storage disk state type is copy. * `Broken` - Storage disk state type is broken. * `Maintenance` - Storage disk state type is maintenance. * `Partner` - Storage disk state type is partner. * `Pending` - Storage disk state type is pending. * `Reconstructing` - Storage disk state type is reconstructing. * `Removed` - Storage disk state type is removed. * `Spare` - Storage disk state type is spare. * `Unfail` - Storage disk state type is unfail. * `Zeroing` - Storage disk state type is zeroing.
	State *string `json:"State,omitempty"`
	// Uuid of  NetApp Disk.
	Uuid            *string                           `json:"Uuid,omitempty"`
	Array           *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	ArrayController *StorageNetAppNodeRelationship    `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppAggregate resources.
	DiskPool             []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppBaseDisk StorageNetAppBaseDisk

// NewStorageNetAppBaseDisk instantiates a new StorageNetAppBaseDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppBaseDisk(classId string, objectType string) *StorageNetAppBaseDisk {
	this := StorageNetAppBaseDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	var containerType string = "Unknown"
	this.ContainerType = &containerType
	var diskType string = "Unknown"
	this.DiskType = &diskType
	var state string = "Present"
	this.State = &state
	return &this
}

// NewStorageNetAppBaseDiskWithDefaults instantiates a new StorageNetAppBaseDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppBaseDiskWithDefaults() *StorageNetAppBaseDisk {
	this := StorageNetAppBaseDisk{}
	var classId string = "storage.NetAppBaseDisk"
	this.ClassId = classId
	var objectType string = "storage.NetAppBaseDisk"
	this.ObjectType = objectType
	var containerType string = "Unknown"
	this.ContainerType = &containerType
	var diskType string = "Unknown"
	this.DiskType = &diskType
	var state string = "Present"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppBaseDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppBaseDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppBaseDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppBaseDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBaseDiskModel returns the BaseDiskModel field value if set, zero value otherwise.
func (o *StorageNetAppBaseDisk) GetBaseDiskModel() string {
	if o == nil || o.BaseDiskModel == nil {
		var ret string
		return ret
	}
	return *o.BaseDiskModel
}

// GetBaseDiskModelOk returns a tuple with the BaseDiskModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetBaseDiskModelOk() (*string, bool) {
	if o == nil || o.BaseDiskModel == nil {
		return nil, false
	}
	return o.BaseDiskModel, true
}

// HasBaseDiskModel returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasBaseDiskModel() bool {
	if o != nil && o.BaseDiskModel != nil {
		return true
	}

	return false
}

// SetBaseDiskModel gets a reference to the given string and assigns it to the BaseDiskModel field.
func (o *StorageNetAppBaseDisk) SetBaseDiskModel(v string) {
	o.BaseDiskModel = &v
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *StorageNetAppBaseDisk) GetContainerType() string {
	if o == nil || o.ContainerType == nil {
		var ret string
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetContainerTypeOk() (*string, bool) {
	if o == nil || o.ContainerType == nil {
		return nil, false
	}
	return o.ContainerType, true
}

// HasContainerType returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasContainerType() bool {
	if o != nil && o.ContainerType != nil {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given string and assigns it to the ContainerType field.
func (o *StorageNetAppBaseDisk) SetContainerType(v string) {
	o.ContainerType = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *StorageNetAppBaseDisk) GetDiskType() string {
	if o == nil || o.DiskType == nil {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetDiskTypeOk() (*string, bool) {
	if o == nil || o.DiskType == nil {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasDiskType() bool {
	if o != nil && o.DiskType != nil {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *StorageNetAppBaseDisk) SetDiskType(v string) {
	o.DiskType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppBaseDisk) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppBaseDisk) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseDisk) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppBaseDisk) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppBaseDisk) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppBaseDisk) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppBaseDisk) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseDisk) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppBaseDisk) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetDiskPool returns the DiskPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppBaseDisk) GetDiskPool() []StorageNetAppAggregateRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateRelationship
		return ret
	}
	return o.DiskPool
}

// GetDiskPoolOk returns a tuple with the DiskPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppBaseDisk) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool) {
	if o == nil || o.DiskPool == nil {
		return nil, false
	}
	return &o.DiskPool, true
}

// HasDiskPool returns a boolean if a field has been set.
func (o *StorageNetAppBaseDisk) HasDiskPool() bool {
	if o != nil && o.DiskPool != nil {
		return true
	}

	return false
}

// SetDiskPool gets a reference to the given []StorageNetAppAggregateRelationship and assigns it to the DiskPool field.
func (o *StorageNetAppBaseDisk) SetDiskPool(v []StorageNetAppAggregateRelationship) {
	o.DiskPool = v
}

func (o StorageNetAppBaseDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseArrayDisk, errStorageBaseArrayDisk := json.Marshal(o.StorageBaseArrayDisk)
	if errStorageBaseArrayDisk != nil {
		return []byte{}, errStorageBaseArrayDisk
	}
	errStorageBaseArrayDisk = json.Unmarshal([]byte(serializedStorageBaseArrayDisk), &toSerialize)
	if errStorageBaseArrayDisk != nil {
		return []byte{}, errStorageBaseArrayDisk
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BaseDiskModel != nil {
		toSerialize["BaseDiskModel"] = o.BaseDiskModel
	}
	if o.ContainerType != nil {
		toSerialize["ContainerType"] = o.ContainerType
	}
	if o.DiskType != nil {
		toSerialize["DiskType"] = o.DiskType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.DiskPool != nil {
		toSerialize["DiskPool"] = o.DiskPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppBaseDisk) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppBaseDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// NetApp base disk model.
		BaseDiskModel *string `json:"BaseDiskModel,omitempty"`
		// Supported container type for NetApp disk. * `Unknown` - Container is currently unknown. This is the default setting. * `Aggregate` - Disk is used as a physical disk in an aggregate. * `Broken` - Disk is in broken pool. * `Label Maintenance` - Disk is in online label maintenance list. * `Foreign` - Array LUN has been marked foreign. * `Maintenance` - Disk is in maintenance center. * `Mediator` - A mediator disk is a disk used on non-shared HA systems hosted by an external node which is used to communicate the viability of the storage failover between non-shared HA nodes. * `Shared` - Disk is partitioned or in a storage pool. * `Remote` - Disk belongs to a remote cluster. * `Spare` - Disk is a spare disk. * `Unassigned` - Disk ownership has not been assigned. * `Unsupported` - Disk is not supported.
		ContainerType *string `json:"ContainerType,omitempty"`
		// Type of the NetApp disk. * `Unknown` - Default unknown disk type. * `SSDNVM` - Solid state disk with Non-Volatile Memory Express protocol enabled. * `ATA` - Advanced Technology Attachment is a type of disk drive that integrates the drive controller directly on the drive itself. * `FCAL` - For the FC-AL disk connection type, disk shelves are connected to the controller in a loop * `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, and rotational speed of traditional enterprise-class SATA drives with the fully capable SAS interface typical for classic SAS drives. * `LUN` - Logical Unit Number refers to a logical disk. * `SAS` - Storage disk with serial attached SCSI. * `MSATA` - SATA disk in multi-disk carrier storage shelf. * `SSD` - Storage disk with Solid state disk. * `VMDISK` - Virtual machine Data Disk.
		DiskType *string `json:"DiskType,omitempty"`
		// Current state of the NetApp disk. * `Present` - Storage disk state type is present. * `Copy` - Storage disk state type is copy. * `Broken` - Storage disk state type is broken. * `Maintenance` - Storage disk state type is maintenance. * `Partner` - Storage disk state type is partner. * `Pending` - Storage disk state type is pending. * `Reconstructing` - Storage disk state type is reconstructing. * `Removed` - Storage disk state type is removed. * `Spare` - Storage disk state type is spare. * `Unfail` - Storage disk state type is unfail. * `Zeroing` - Storage disk state type is zeroing.
		State *string `json:"State,omitempty"`
		// Uuid of  NetApp Disk.
		Uuid            *string                           `json:"Uuid,omitempty"`
		Array           *StorageNetAppClusterRelationship `json:"Array,omitempty"`
		ArrayController *StorageNetAppNodeRelationship    `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppAggregate resources.
		DiskPool []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	}

	varStorageNetAppBaseDiskWithoutEmbeddedStruct := StorageNetAppBaseDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseDiskWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppBaseDisk := _StorageNetAppBaseDisk{}
		varStorageNetAppBaseDisk.ClassId = varStorageNetAppBaseDiskWithoutEmbeddedStruct.ClassId
		varStorageNetAppBaseDisk.ObjectType = varStorageNetAppBaseDiskWithoutEmbeddedStruct.ObjectType
		varStorageNetAppBaseDisk.BaseDiskModel = varStorageNetAppBaseDiskWithoutEmbeddedStruct.BaseDiskModel
		varStorageNetAppBaseDisk.ContainerType = varStorageNetAppBaseDiskWithoutEmbeddedStruct.ContainerType
		varStorageNetAppBaseDisk.DiskType = varStorageNetAppBaseDiskWithoutEmbeddedStruct.DiskType
		varStorageNetAppBaseDisk.State = varStorageNetAppBaseDiskWithoutEmbeddedStruct.State
		varStorageNetAppBaseDisk.Uuid = varStorageNetAppBaseDiskWithoutEmbeddedStruct.Uuid
		varStorageNetAppBaseDisk.Array = varStorageNetAppBaseDiskWithoutEmbeddedStruct.Array
		varStorageNetAppBaseDisk.ArrayController = varStorageNetAppBaseDiskWithoutEmbeddedStruct.ArrayController
		varStorageNetAppBaseDisk.DiskPool = varStorageNetAppBaseDiskWithoutEmbeddedStruct.DiskPool
		*o = StorageNetAppBaseDisk(varStorageNetAppBaseDisk)
	} else {
		return err
	}

	varStorageNetAppBaseDisk := _StorageNetAppBaseDisk{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseDisk)
	if err == nil {
		o.StorageBaseArrayDisk = varStorageNetAppBaseDisk.StorageBaseArrayDisk
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseDiskModel")
		delete(additionalProperties, "ContainerType")
		delete(additionalProperties, "DiskType")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "DiskPool")

		// remove fields from embedded structs
		reflectStorageBaseArrayDisk := reflect.ValueOf(o.StorageBaseArrayDisk)
		for i := 0; i < reflectStorageBaseArrayDisk.Type().NumField(); i++ {
			t := reflectStorageBaseArrayDisk.Type().Field(i)

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

type NullableStorageNetAppBaseDisk struct {
	value *StorageNetAppBaseDisk
	isSet bool
}

func (v NullableStorageNetAppBaseDisk) Get() *StorageNetAppBaseDisk {
	return v.value
}

func (v *NullableStorageNetAppBaseDisk) Set(val *StorageNetAppBaseDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppBaseDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppBaseDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppBaseDisk(val *StorageNetAppBaseDisk) *NullableStorageNetAppBaseDisk {
	return &NullableStorageNetAppBaseDisk{value: val, isSet: true}
}

func (v NullableStorageNetAppBaseDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppBaseDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
