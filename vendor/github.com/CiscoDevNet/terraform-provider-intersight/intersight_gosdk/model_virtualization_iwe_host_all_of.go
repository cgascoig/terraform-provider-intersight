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

// VirtualizationIweHostAllOf Definition of the list of properties defined in 'virtualization.IweHost', excluding properties defined in parent classes.
type VirtualizationIweHostAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Denotes age or life time of the Host in nano seconds.
	Age *string `json:"Age,omitempty"`
	// Chassis version of the Host.
	ChassisVersion *string `json:"ChassisVersion,omitempty"`
	// The UUID of the cluster to which this Host belongs to.
	ClusterUuid   *string                             `json:"ClusterUuid,omitempty"`
	CpuAllocation NullableVirtualizationCpuAllocation `json:"CpuAllocation,omitempty"`
	// Reason of the failure when host is in failed state.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Is the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweringOn` - The entity is powering on. * `PoweredOn` - The entity is powered on. * `PoweringOff` - The entity is powering off. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state. * `Rebooting` - The entity reboot is in progress. * `` - The entity's power state is not available.
	HwPowerState *string `json:"HwPowerState,omitempty"`
	// Internal IP Address of the Host.
	InternalIpAddress *string `json:"InternalIpAddress,omitempty"`
	// Management IP Address of the Host.
	ManagementIpAddress *string `json:"ManagementIpAddress,omitempty"`
	// Is the role of this host is master in the cluster? true or false.
	MasterRole       *bool                                  `json:"MasterRole,omitempty"`
	MemoryAllocation NullableVirtualizationMemoryAllocation `json:"MemoryAllocation,omitempty"`
	StorageCapacity  NullableVirtualizationStorageCapacity  `json:"StorageCapacity,omitempty"`
	// Is the Storage Controller VM on the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweringOn` - The entity is powering on. * `PoweredOn` - The entity is powered on. * `PoweringOff` - The entity is powering off. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state. * `Rebooting` - The entity reboot is in progress. * `` - The entity's power state is not available.
	StorageVmPowerState *string `json:"StorageVmPowerState,omitempty"`
	// Product version of the Host.
	Version              *string                               `json:"Version,omitempty"`
	Cluster              *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	ClusterMember        *AssetClusterMemberRelationship       `json:"ClusterMember,omitempty"`
	PhysicalServer       *ComputePhysicalRelationship          `json:"PhysicalServer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweHostAllOf VirtualizationIweHostAllOf

// NewVirtualizationIweHostAllOf instantiates a new VirtualizationIweHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweHostAllOf(classId string, objectType string) *VirtualizationIweHostAllOf {
	this := VirtualizationIweHostAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	var storageVmPowerState string = "Unknown"
	this.StorageVmPowerState = &storageVmPowerState
	return &this
}

// NewVirtualizationIweHostAllOfWithDefaults instantiates a new VirtualizationIweHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweHostAllOfWithDefaults() *VirtualizationIweHostAllOf {
	this := VirtualizationIweHostAllOf{}
	var classId string = "virtualization.IweHost"
	this.ClassId = classId
	var objectType string = "virtualization.IweHost"
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	var storageVmPowerState string = "Unknown"
	this.StorageVmPowerState = &storageVmPowerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweHostAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweHostAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweHostAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweHostAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetAge() string {
	if o == nil || o.Age == nil {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetAgeOk() (*string, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *VirtualizationIweHostAllOf) SetAge(v string) {
	o.Age = &v
}

// GetChassisVersion returns the ChassisVersion field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetChassisVersion() string {
	if o == nil || o.ChassisVersion == nil {
		var ret string
		return ret
	}
	return *o.ChassisVersion
}

// GetChassisVersionOk returns a tuple with the ChassisVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetChassisVersionOk() (*string, bool) {
	if o == nil || o.ChassisVersion == nil {
		return nil, false
	}
	return o.ChassisVersion, true
}

// HasChassisVersion returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasChassisVersion() bool {
	if o != nil && o.ChassisVersion != nil {
		return true
	}

	return false
}

// SetChassisVersion gets a reference to the given string and assigns it to the ChassisVersion field.
func (o *VirtualizationIweHostAllOf) SetChassisVersion(v string) {
	o.ChassisVersion = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *VirtualizationIweHostAllOf) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetCpuAllocation returns the CpuAllocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweHostAllOf) GetCpuAllocation() VirtualizationCpuAllocation {
	if o == nil || o.CpuAllocation.Get() == nil {
		var ret VirtualizationCpuAllocation
		return ret
	}
	return *o.CpuAllocation.Get()
}

// GetCpuAllocationOk returns a tuple with the CpuAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweHostAllOf) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuAllocation.Get(), o.CpuAllocation.IsSet()
}

// HasCpuAllocation returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasCpuAllocation() bool {
	if o != nil && o.CpuAllocation.IsSet() {
		return true
	}

	return false
}

// SetCpuAllocation gets a reference to the given NullableVirtualizationCpuAllocation and assigns it to the CpuAllocation field.
func (o *VirtualizationIweHostAllOf) SetCpuAllocation(v VirtualizationCpuAllocation) {
	o.CpuAllocation.Set(&v)
}

// SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil
func (o *VirtualizationIweHostAllOf) SetCpuAllocationNil() {
	o.CpuAllocation.Set(nil)
}

// UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
func (o *VirtualizationIweHostAllOf) UnsetCpuAllocation() {
	o.CpuAllocation.Unset()
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *VirtualizationIweHostAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetHwPowerState returns the HwPowerState field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetHwPowerState() string {
	if o == nil || o.HwPowerState == nil {
		var ret string
		return ret
	}
	return *o.HwPowerState
}

// GetHwPowerStateOk returns a tuple with the HwPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetHwPowerStateOk() (*string, bool) {
	if o == nil || o.HwPowerState == nil {
		return nil, false
	}
	return o.HwPowerState, true
}

// HasHwPowerState returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasHwPowerState() bool {
	if o != nil && o.HwPowerState != nil {
		return true
	}

	return false
}

// SetHwPowerState gets a reference to the given string and assigns it to the HwPowerState field.
func (o *VirtualizationIweHostAllOf) SetHwPowerState(v string) {
	o.HwPowerState = &v
}

// GetInternalIpAddress returns the InternalIpAddress field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetInternalIpAddress() string {
	if o == nil || o.InternalIpAddress == nil {
		var ret string
		return ret
	}
	return *o.InternalIpAddress
}

// GetInternalIpAddressOk returns a tuple with the InternalIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetInternalIpAddressOk() (*string, bool) {
	if o == nil || o.InternalIpAddress == nil {
		return nil, false
	}
	return o.InternalIpAddress, true
}

// HasInternalIpAddress returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasInternalIpAddress() bool {
	if o != nil && o.InternalIpAddress != nil {
		return true
	}

	return false
}

// SetInternalIpAddress gets a reference to the given string and assigns it to the InternalIpAddress field.
func (o *VirtualizationIweHostAllOf) SetInternalIpAddress(v string) {
	o.InternalIpAddress = &v
}

// GetManagementIpAddress returns the ManagementIpAddress field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetManagementIpAddress() string {
	if o == nil || o.ManagementIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementIpAddress
}

// GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetManagementIpAddressOk() (*string, bool) {
	if o == nil || o.ManagementIpAddress == nil {
		return nil, false
	}
	return o.ManagementIpAddress, true
}

// HasManagementIpAddress returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasManagementIpAddress() bool {
	if o != nil && o.ManagementIpAddress != nil {
		return true
	}

	return false
}

// SetManagementIpAddress gets a reference to the given string and assigns it to the ManagementIpAddress field.
func (o *VirtualizationIweHostAllOf) SetManagementIpAddress(v string) {
	o.ManagementIpAddress = &v
}

// GetMasterRole returns the MasterRole field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetMasterRole() bool {
	if o == nil || o.MasterRole == nil {
		var ret bool
		return ret
	}
	return *o.MasterRole
}

// GetMasterRoleOk returns a tuple with the MasterRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetMasterRoleOk() (*bool, bool) {
	if o == nil || o.MasterRole == nil {
		return nil, false
	}
	return o.MasterRole, true
}

// HasMasterRole returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasMasterRole() bool {
	if o != nil && o.MasterRole != nil {
		return true
	}

	return false
}

// SetMasterRole gets a reference to the given bool and assigns it to the MasterRole field.
func (o *VirtualizationIweHostAllOf) SetMasterRole(v bool) {
	o.MasterRole = &v
}

// GetMemoryAllocation returns the MemoryAllocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweHostAllOf) GetMemoryAllocation() VirtualizationMemoryAllocation {
	if o == nil || o.MemoryAllocation.Get() == nil {
		var ret VirtualizationMemoryAllocation
		return ret
	}
	return *o.MemoryAllocation.Get()
}

// GetMemoryAllocationOk returns a tuple with the MemoryAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweHostAllOf) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryAllocation.Get(), o.MemoryAllocation.IsSet()
}

// HasMemoryAllocation returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasMemoryAllocation() bool {
	if o != nil && o.MemoryAllocation.IsSet() {
		return true
	}

	return false
}

// SetMemoryAllocation gets a reference to the given NullableVirtualizationMemoryAllocation and assigns it to the MemoryAllocation field.
func (o *VirtualizationIweHostAllOf) SetMemoryAllocation(v VirtualizationMemoryAllocation) {
	o.MemoryAllocation.Set(&v)
}

// SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil
func (o *VirtualizationIweHostAllOf) SetMemoryAllocationNil() {
	o.MemoryAllocation.Set(nil)
}

// UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
func (o *VirtualizationIweHostAllOf) UnsetMemoryAllocation() {
	o.MemoryAllocation.Unset()
}

// GetStorageCapacity returns the StorageCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweHostAllOf) GetStorageCapacity() VirtualizationStorageCapacity {
	if o == nil || o.StorageCapacity.Get() == nil {
		var ret VirtualizationStorageCapacity
		return ret
	}
	return *o.StorageCapacity.Get()
}

// GetStorageCapacityOk returns a tuple with the StorageCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweHostAllOf) GetStorageCapacityOk() (*VirtualizationStorageCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageCapacity.Get(), o.StorageCapacity.IsSet()
}

// HasStorageCapacity returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasStorageCapacity() bool {
	if o != nil && o.StorageCapacity.IsSet() {
		return true
	}

	return false
}

// SetStorageCapacity gets a reference to the given NullableVirtualizationStorageCapacity and assigns it to the StorageCapacity field.
func (o *VirtualizationIweHostAllOf) SetStorageCapacity(v VirtualizationStorageCapacity) {
	o.StorageCapacity.Set(&v)
}

// SetStorageCapacityNil sets the value for StorageCapacity to be an explicit nil
func (o *VirtualizationIweHostAllOf) SetStorageCapacityNil() {
	o.StorageCapacity.Set(nil)
}

// UnsetStorageCapacity ensures that no value is present for StorageCapacity, not even an explicit nil
func (o *VirtualizationIweHostAllOf) UnsetStorageCapacity() {
	o.StorageCapacity.Unset()
}

// GetStorageVmPowerState returns the StorageVmPowerState field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetStorageVmPowerState() string {
	if o == nil || o.StorageVmPowerState == nil {
		var ret string
		return ret
	}
	return *o.StorageVmPowerState
}

// GetStorageVmPowerStateOk returns a tuple with the StorageVmPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetStorageVmPowerStateOk() (*string, bool) {
	if o == nil || o.StorageVmPowerState == nil {
		return nil, false
	}
	return o.StorageVmPowerState, true
}

// HasStorageVmPowerState returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasStorageVmPowerState() bool {
	if o != nil && o.StorageVmPowerState != nil {
		return true
	}

	return false
}

// SetStorageVmPowerState gets a reference to the given string and assigns it to the StorageVmPowerState field.
func (o *VirtualizationIweHostAllOf) SetStorageVmPowerState(v string) {
	o.StorageVmPowerState = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VirtualizationIweHostAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweHostAllOf) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || o.ClusterMember == nil {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMember == nil {
		return nil, false
	}
	return o.ClusterMember, true
}

// HasClusterMember returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasClusterMember() bool {
	if o != nil && o.ClusterMember != nil {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given AssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *VirtualizationIweHostAllOf) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember = &v
}

// GetPhysicalServer returns the PhysicalServer field value if set, zero value otherwise.
func (o *VirtualizationIweHostAllOf) GetPhysicalServer() ComputePhysicalRelationship {
	if o == nil || o.PhysicalServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.PhysicalServer
}

// GetPhysicalServerOk returns a tuple with the PhysicalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweHostAllOf) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.PhysicalServer == nil {
		return nil, false
	}
	return o.PhysicalServer, true
}

// HasPhysicalServer returns a boolean if a field has been set.
func (o *VirtualizationIweHostAllOf) HasPhysicalServer() bool {
	if o != nil && o.PhysicalServer != nil {
		return true
	}

	return false
}

// SetPhysicalServer gets a reference to the given ComputePhysicalRelationship and assigns it to the PhysicalServer field.
func (o *VirtualizationIweHostAllOf) SetPhysicalServer(v ComputePhysicalRelationship) {
	o.PhysicalServer = &v
}

func (o VirtualizationIweHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Age != nil {
		toSerialize["Age"] = o.Age
	}
	if o.ChassisVersion != nil {
		toSerialize["ChassisVersion"] = o.ChassisVersion
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.CpuAllocation.IsSet() {
		toSerialize["CpuAllocation"] = o.CpuAllocation.Get()
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.HwPowerState != nil {
		toSerialize["HwPowerState"] = o.HwPowerState
	}
	if o.InternalIpAddress != nil {
		toSerialize["InternalIpAddress"] = o.InternalIpAddress
	}
	if o.ManagementIpAddress != nil {
		toSerialize["ManagementIpAddress"] = o.ManagementIpAddress
	}
	if o.MasterRole != nil {
		toSerialize["MasterRole"] = o.MasterRole
	}
	if o.MemoryAllocation.IsSet() {
		toSerialize["MemoryAllocation"] = o.MemoryAllocation.Get()
	}
	if o.StorageCapacity.IsSet() {
		toSerialize["StorageCapacity"] = o.StorageCapacity.Get()
	}
	if o.StorageVmPowerState != nil {
		toSerialize["StorageVmPowerState"] = o.StorageVmPowerState
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.ClusterMember != nil {
		toSerialize["ClusterMember"] = o.ClusterMember
	}
	if o.PhysicalServer != nil {
		toSerialize["PhysicalServer"] = o.PhysicalServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweHostAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationIweHostAllOf := _VirtualizationIweHostAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationIweHostAllOf); err == nil {
		*o = VirtualizationIweHostAllOf(varVirtualizationIweHostAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Age")
		delete(additionalProperties, "ChassisVersion")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "CpuAllocation")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "HwPowerState")
		delete(additionalProperties, "InternalIpAddress")
		delete(additionalProperties, "ManagementIpAddress")
		delete(additionalProperties, "MasterRole")
		delete(additionalProperties, "MemoryAllocation")
		delete(additionalProperties, "StorageCapacity")
		delete(additionalProperties, "StorageVmPowerState")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "ClusterMember")
		delete(additionalProperties, "PhysicalServer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationIweHostAllOf struct {
	value *VirtualizationIweHostAllOf
	isSet bool
}

func (v NullableVirtualizationIweHostAllOf) Get() *VirtualizationIweHostAllOf {
	return v.value
}

func (v *NullableVirtualizationIweHostAllOf) Set(val *VirtualizationIweHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweHostAllOf(val *VirtualizationIweHostAllOf) *NullableVirtualizationIweHostAllOf {
	return &NullableVirtualizationIweHostAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationIweHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
