/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageDriveGroupAllOf Definition of the list of properties defined in 'storage.DriveGroup', excluding properties defined in parent classes.
type StorageDriveGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType          string                             `json:"ObjectType"`
	AutomaticDriveGroup NullableStorageAutomaticDriveGroup `json:"AutomaticDriveGroup,omitempty"`
	ManualDriveGroup    NullableStorageManualDriveGroup    `json:"ManualDriveGroup,omitempty"`
	// The name of the drive group. The name can be between 1 and 64 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
	Name *string `json:"Name,omitempty"`
	// The supported RAID level for the disk group. * `Raid0` - RAID 0 Stripe Raid Level. * `Raid1` - RAID 1 Mirror Raid Level. * `Raid5` - RAID 5 Mirror Raid Level. * `Raid6` - RAID 6 Mirror Raid Level. * `Raid10` - RAID 10 Mirror Raid Level. * `Raid50` - RAID 50 Mirror Raid Level. * `Raid60` - RAID 60 Mirror Raid Level.
	RaidLevel *string `json:"RaidLevel,omitempty"`
	// Type of drive selection to be used for this drive group. * `0` - Drives are selected manually by the user. * `1` - Drives are selected automatically based on the RAID and virtual drive configuration.
	Type                 *int32                             `json:"Type,omitempty"`
	VirtualDrives        []StorageVirtualDriveConfiguration `json:"VirtualDrives,omitempty"`
	StoragePolicy        *StorageStoragePolicyRelationship  `json:"StoragePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageDriveGroupAllOf StorageDriveGroupAllOf

// NewStorageDriveGroupAllOf instantiates a new StorageDriveGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDriveGroupAllOf(classId string, objectType string) *StorageDriveGroupAllOf {
	this := StorageDriveGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// NewStorageDriveGroupAllOfWithDefaults instantiates a new StorageDriveGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDriveGroupAllOfWithDefaults() *StorageDriveGroupAllOf {
	this := StorageDriveGroupAllOf{}
	var classId string = "storage.DriveGroup"
	this.ClassId = classId
	var objectType string = "storage.DriveGroup"
	this.ObjectType = objectType
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageDriveGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageDriveGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageDriveGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageDriveGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageDriveGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageDriveGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutomaticDriveGroup returns the AutomaticDriveGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDriveGroupAllOf) GetAutomaticDriveGroup() StorageAutomaticDriveGroup {
	if o == nil || o.AutomaticDriveGroup.Get() == nil {
		var ret StorageAutomaticDriveGroup
		return ret
	}
	return *o.AutomaticDriveGroup.Get()
}

// GetAutomaticDriveGroupOk returns a tuple with the AutomaticDriveGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDriveGroupAllOf) GetAutomaticDriveGroupOk() (*StorageAutomaticDriveGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutomaticDriveGroup.Get(), o.AutomaticDriveGroup.IsSet()
}

// HasAutomaticDriveGroup returns a boolean if a field has been set.
func (o *StorageDriveGroupAllOf) HasAutomaticDriveGroup() bool {
	if o != nil && o.AutomaticDriveGroup.IsSet() {
		return true
	}

	return false
}

// SetAutomaticDriveGroup gets a reference to the given NullableStorageAutomaticDriveGroup and assigns it to the AutomaticDriveGroup field.
func (o *StorageDriveGroupAllOf) SetAutomaticDriveGroup(v StorageAutomaticDriveGroup) {
	o.AutomaticDriveGroup.Set(&v)
}

// SetAutomaticDriveGroupNil sets the value for AutomaticDriveGroup to be an explicit nil
func (o *StorageDriveGroupAllOf) SetAutomaticDriveGroupNil() {
	o.AutomaticDriveGroup.Set(nil)
}

// UnsetAutomaticDriveGroup ensures that no value is present for AutomaticDriveGroup, not even an explicit nil
func (o *StorageDriveGroupAllOf) UnsetAutomaticDriveGroup() {
	o.AutomaticDriveGroup.Unset()
}

// GetManualDriveGroup returns the ManualDriveGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDriveGroupAllOf) GetManualDriveGroup() StorageManualDriveGroup {
	if o == nil || o.ManualDriveGroup.Get() == nil {
		var ret StorageManualDriveGroup
		return ret
	}
	return *o.ManualDriveGroup.Get()
}

// GetManualDriveGroupOk returns a tuple with the ManualDriveGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDriveGroupAllOf) GetManualDriveGroupOk() (*StorageManualDriveGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManualDriveGroup.Get(), o.ManualDriveGroup.IsSet()
}

// HasManualDriveGroup returns a boolean if a field has been set.
func (o *StorageDriveGroupAllOf) HasManualDriveGroup() bool {
	if o != nil && o.ManualDriveGroup.IsSet() {
		return true
	}

	return false
}

// SetManualDriveGroup gets a reference to the given NullableStorageManualDriveGroup and assigns it to the ManualDriveGroup field.
func (o *StorageDriveGroupAllOf) SetManualDriveGroup(v StorageManualDriveGroup) {
	o.ManualDriveGroup.Set(&v)
}

// SetManualDriveGroupNil sets the value for ManualDriveGroup to be an explicit nil
func (o *StorageDriveGroupAllOf) SetManualDriveGroupNil() {
	o.ManualDriveGroup.Set(nil)
}

// UnsetManualDriveGroup ensures that no value is present for ManualDriveGroup, not even an explicit nil
func (o *StorageDriveGroupAllOf) UnsetManualDriveGroup() {
	o.ManualDriveGroup.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageDriveGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageDriveGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageDriveGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetRaidLevel returns the RaidLevel field value if set, zero value otherwise.
func (o *StorageDriveGroupAllOf) GetRaidLevel() string {
	if o == nil || o.RaidLevel == nil {
		var ret string
		return ret
	}
	return *o.RaidLevel
}

// GetRaidLevelOk returns a tuple with the RaidLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroupAllOf) GetRaidLevelOk() (*string, bool) {
	if o == nil || o.RaidLevel == nil {
		return nil, false
	}
	return o.RaidLevel, true
}

// HasRaidLevel returns a boolean if a field has been set.
func (o *StorageDriveGroupAllOf) HasRaidLevel() bool {
	if o != nil && o.RaidLevel != nil {
		return true
	}

	return false
}

// SetRaidLevel gets a reference to the given string and assigns it to the RaidLevel field.
func (o *StorageDriveGroupAllOf) SetRaidLevel(v string) {
	o.RaidLevel = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageDriveGroupAllOf) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroupAllOf) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageDriveGroupAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *StorageDriveGroupAllOf) SetType(v int32) {
	o.Type = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDriveGroupAllOf) GetVirtualDrives() []StorageVirtualDriveConfiguration {
	if o == nil {
		var ret []StorageVirtualDriveConfiguration
		return ret
	}
	return o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDriveGroupAllOf) GetVirtualDrivesOk() (*[]StorageVirtualDriveConfiguration, bool) {
	if o == nil || o.VirtualDrives == nil {
		return nil, false
	}
	return &o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *StorageDriveGroupAllOf) HasVirtualDrives() bool {
	if o != nil && o.VirtualDrives != nil {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []StorageVirtualDriveConfiguration and assigns it to the VirtualDrives field.
func (o *StorageDriveGroupAllOf) SetVirtualDrives(v []StorageVirtualDriveConfiguration) {
	o.VirtualDrives = v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise.
func (o *StorageDriveGroupAllOf) GetStoragePolicy() StorageStoragePolicyRelationship {
	if o == nil || o.StoragePolicy == nil {
		var ret StorageStoragePolicyRelationship
		return ret
	}
	return *o.StoragePolicy
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDriveGroupAllOf) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool) {
	if o == nil || o.StoragePolicy == nil {
		return nil, false
	}
	return o.StoragePolicy, true
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *StorageDriveGroupAllOf) HasStoragePolicy() bool {
	if o != nil && o.StoragePolicy != nil {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given StorageStoragePolicyRelationship and assigns it to the StoragePolicy field.
func (o *StorageDriveGroupAllOf) SetStoragePolicy(v StorageStoragePolicyRelationship) {
	o.StoragePolicy = &v
}

func (o StorageDriveGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutomaticDriveGroup.IsSet() {
		toSerialize["AutomaticDriveGroup"] = o.AutomaticDriveGroup.Get()
	}
	if o.ManualDriveGroup.IsSet() {
		toSerialize["ManualDriveGroup"] = o.ManualDriveGroup.Get()
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RaidLevel != nil {
		toSerialize["RaidLevel"] = o.RaidLevel
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDrives != nil {
		toSerialize["VirtualDrives"] = o.VirtualDrives
	}
	if o.StoragePolicy != nil {
		toSerialize["StoragePolicy"] = o.StoragePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageDriveGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageDriveGroupAllOf := _StorageDriveGroupAllOf{}

	if err = json.Unmarshal(bytes, &varStorageDriveGroupAllOf); err == nil {
		*o = StorageDriveGroupAllOf(varStorageDriveGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutomaticDriveGroup")
		delete(additionalProperties, "ManualDriveGroup")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RaidLevel")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VirtualDrives")
		delete(additionalProperties, "StoragePolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageDriveGroupAllOf struct {
	value *StorageDriveGroupAllOf
	isSet bool
}

func (v NullableStorageDriveGroupAllOf) Get() *StorageDriveGroupAllOf {
	return v.value
}

func (v *NullableStorageDriveGroupAllOf) Set(val *StorageDriveGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDriveGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDriveGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDriveGroupAllOf(val *StorageDriveGroupAllOf) *NullableStorageDriveGroupAllOf {
	return &NullableStorageDriveGroupAllOf{value: val, isSet: true}
}

func (v NullableStorageDriveGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDriveGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
