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

// HyperflexHypervisorVirtualMachineAllOf Definition of the list of properties defined in 'hyperflex.HypervisorVirtualMachine', excluding properties defined in parent classes.
type HyperflexHypervisorVirtualMachineAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The connectivity state of the HyperFlex virtual machine.
	ConnectionState *string `json:"ConnectionState,omitempty"`
	// Guest operating system state of the HyperFlex virtual machine.
	GuestOsState *string `json:"GuestOsState,omitempty"`
	// Host UUID of the HyperFlex virtual machine.
	HostUuid *string                                `json:"HostUuid,omitempty"`
	Ip       NullableNetworkHyperFlexNetworkAddress `json:"Ip,omitempty"`
	// Directory path where virtual machine is stored.
	Path *string `json:"Path,omitempty"`
	// The instance id of platform which a virtual machine is running on.
	PlatformInstanceId *string `json:"PlatformInstanceId,omitempty"`
	// Total provisioned storage to the HyperFlex virtual machine in bytes.
	StorageProvisionedInBytes *int64 `json:"StorageProvisionedInBytes,omitempty"`
	// Storage used by HyperFlex virtual machine in bytes.
	StorageUsedInBytes *int64 `json:"StorageUsedInBytes,omitempty"`
	// Flag indicating whether or not this virtual machine is a template. Apply to the ESXi platform only.
	Template *bool `json:"Template,omitempty"`
	// The instance UUID of a virtual machine.
	VmInstanceUuid       *string                              `json:"VmInstanceUuid,omitempty"`
	Cluster              *HyperflexClusterRelationship        `json:"Cluster,omitempty"`
	Host                 *HyperflexHypervisorHostRelationship `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHypervisorVirtualMachineAllOf HyperflexHypervisorVirtualMachineAllOf

// NewHyperflexHypervisorVirtualMachineAllOf instantiates a new HyperflexHypervisorVirtualMachineAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHypervisorVirtualMachineAllOf(classId string, objectType string) *HyperflexHypervisorVirtualMachineAllOf {
	this := HyperflexHypervisorVirtualMachineAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHypervisorVirtualMachineAllOfWithDefaults instantiates a new HyperflexHypervisorVirtualMachineAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHypervisorVirtualMachineAllOfWithDefaults() *HyperflexHypervisorVirtualMachineAllOf {
	this := HyperflexHypervisorVirtualMachineAllOf{}
	var classId string = "hyperflex.HypervisorVirtualMachine"
	this.ClassId = classId
	var objectType string = "hyperflex.HypervisorVirtualMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHypervisorVirtualMachineAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHypervisorVirtualMachineAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHypervisorVirtualMachineAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHypervisorVirtualMachineAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetConnectionState() string {
	if o == nil || o.ConnectionState == nil {
		var ret string
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetConnectionStateOk() (*string, bool) {
	if o == nil || o.ConnectionState == nil {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasConnectionState() bool {
	if o != nil && o.ConnectionState != nil {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given string and assigns it to the ConnectionState field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetConnectionState(v string) {
	o.ConnectionState = &v
}

// GetGuestOsState returns the GuestOsState field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetGuestOsState() string {
	if o == nil || o.GuestOsState == nil {
		var ret string
		return ret
	}
	return *o.GuestOsState
}

// GetGuestOsStateOk returns a tuple with the GuestOsState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetGuestOsStateOk() (*string, bool) {
	if o == nil || o.GuestOsState == nil {
		return nil, false
	}
	return o.GuestOsState, true
}

// HasGuestOsState returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasGuestOsState() bool {
	if o != nil && o.GuestOsState != nil {
		return true
	}

	return false
}

// SetGuestOsState gets a reference to the given string and assigns it to the GuestOsState field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetGuestOsState(v string) {
	o.GuestOsState = &v
}

// GetHostUuid returns the HostUuid field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetHostUuid() string {
	if o == nil || o.HostUuid == nil {
		var ret string
		return ret
	}
	return *o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetHostUuidOk() (*string, bool) {
	if o == nil || o.HostUuid == nil {
		return nil, false
	}
	return o.HostUuid, true
}

// HasHostUuid returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasHostUuid() bool {
	if o != nil && o.HostUuid != nil {
		return true
	}

	return false
}

// SetHostUuid gets a reference to the given string and assigns it to the HostUuid field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetHostUuid(v string) {
	o.HostUuid = &v
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHypervisorVirtualMachineAllOf) GetIp() NetworkHyperFlexNetworkAddress {
	if o == nil || o.Ip.Get() == nil {
		var ret NetworkHyperFlexNetworkAddress
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHypervisorVirtualMachineAllOf) GetIpOk() (*NetworkHyperFlexNetworkAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableNetworkHyperFlexNetworkAddress and assigns it to the Ip field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetIp(v NetworkHyperFlexNetworkAddress) {
	o.Ip.Set(&v)
}

// SetIpNil sets the value for Ip to be an explicit nil
func (o *HyperflexHypervisorVirtualMachineAllOf) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *HyperflexHypervisorVirtualMachineAllOf) UnsetIp() {
	o.Ip.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetPath(v string) {
	o.Path = &v
}

// GetPlatformInstanceId returns the PlatformInstanceId field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetPlatformInstanceId() string {
	if o == nil || o.PlatformInstanceId == nil {
		var ret string
		return ret
	}
	return *o.PlatformInstanceId
}

// GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetPlatformInstanceIdOk() (*string, bool) {
	if o == nil || o.PlatformInstanceId == nil {
		return nil, false
	}
	return o.PlatformInstanceId, true
}

// HasPlatformInstanceId returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasPlatformInstanceId() bool {
	if o != nil && o.PlatformInstanceId != nil {
		return true
	}

	return false
}

// SetPlatformInstanceId gets a reference to the given string and assigns it to the PlatformInstanceId field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetPlatformInstanceId(v string) {
	o.PlatformInstanceId = &v
}

// GetStorageProvisionedInBytes returns the StorageProvisionedInBytes field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageProvisionedInBytes() int64 {
	if o == nil || o.StorageProvisionedInBytes == nil {
		var ret int64
		return ret
	}
	return *o.StorageProvisionedInBytes
}

// GetStorageProvisionedInBytesOk returns a tuple with the StorageProvisionedInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageProvisionedInBytesOk() (*int64, bool) {
	if o == nil || o.StorageProvisionedInBytes == nil {
		return nil, false
	}
	return o.StorageProvisionedInBytes, true
}

// HasStorageProvisionedInBytes returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasStorageProvisionedInBytes() bool {
	if o != nil && o.StorageProvisionedInBytes != nil {
		return true
	}

	return false
}

// SetStorageProvisionedInBytes gets a reference to the given int64 and assigns it to the StorageProvisionedInBytes field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetStorageProvisionedInBytes(v int64) {
	o.StorageProvisionedInBytes = &v
}

// GetStorageUsedInBytes returns the StorageUsedInBytes field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageUsedInBytes() int64 {
	if o == nil || o.StorageUsedInBytes == nil {
		var ret int64
		return ret
	}
	return *o.StorageUsedInBytes
}

// GetStorageUsedInBytesOk returns a tuple with the StorageUsedInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageUsedInBytesOk() (*int64, bool) {
	if o == nil || o.StorageUsedInBytes == nil {
		return nil, false
	}
	return o.StorageUsedInBytes, true
}

// HasStorageUsedInBytes returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasStorageUsedInBytes() bool {
	if o != nil && o.StorageUsedInBytes != nil {
		return true
	}

	return false
}

// SetStorageUsedInBytes gets a reference to the given int64 and assigns it to the StorageUsedInBytes field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetStorageUsedInBytes(v int64) {
	o.StorageUsedInBytes = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetTemplate() bool {
	if o == nil || o.Template == nil {
		var ret bool
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetTemplateOk() (*bool, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given bool and assigns it to the Template field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetTemplate(v bool) {
	o.Template = &v
}

// GetVmInstanceUuid returns the VmInstanceUuid field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetVmInstanceUuid() string {
	if o == nil || o.VmInstanceUuid == nil {
		var ret string
		return ret
	}
	return *o.VmInstanceUuid
}

// GetVmInstanceUuidOk returns a tuple with the VmInstanceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetVmInstanceUuidOk() (*string, bool) {
	if o == nil || o.VmInstanceUuid == nil {
		return nil, false
	}
	return o.VmInstanceUuid, true
}

// HasVmInstanceUuid returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasVmInstanceUuid() bool {
	if o != nil && o.VmInstanceUuid != nil {
		return true
	}

	return false
}

// SetVmInstanceUuid gets a reference to the given string and assigns it to the VmInstanceUuid field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetVmInstanceUuid(v string) {
	o.VmInstanceUuid = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetHost() HyperflexHypervisorHostRelationship {
	if o == nil || o.Host == nil {
		var ret HyperflexHypervisorHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) GetHostOk() (*HyperflexHypervisorHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachineAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given HyperflexHypervisorHostRelationship and assigns it to the Host field.
func (o *HyperflexHypervisorVirtualMachineAllOf) SetHost(v HyperflexHypervisorHostRelationship) {
	o.Host = &v
}

func (o HyperflexHypervisorVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionState != nil {
		toSerialize["ConnectionState"] = o.ConnectionState
	}
	if o.GuestOsState != nil {
		toSerialize["GuestOsState"] = o.GuestOsState
	}
	if o.HostUuid != nil {
		toSerialize["HostUuid"] = o.HostUuid
	}
	if o.Ip.IsSet() {
		toSerialize["Ip"] = o.Ip.Get()
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.PlatformInstanceId != nil {
		toSerialize["PlatformInstanceId"] = o.PlatformInstanceId
	}
	if o.StorageProvisionedInBytes != nil {
		toSerialize["StorageProvisionedInBytes"] = o.StorageProvisionedInBytes
	}
	if o.StorageUsedInBytes != nil {
		toSerialize["StorageUsedInBytes"] = o.StorageUsedInBytes
	}
	if o.Template != nil {
		toSerialize["Template"] = o.Template
	}
	if o.VmInstanceUuid != nil {
		toSerialize["VmInstanceUuid"] = o.VmInstanceUuid
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHypervisorVirtualMachineAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHypervisorVirtualMachineAllOf := _HyperflexHypervisorVirtualMachineAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHypervisorVirtualMachineAllOf); err == nil {
		*o = HyperflexHypervisorVirtualMachineAllOf(varHyperflexHypervisorVirtualMachineAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionState")
		delete(additionalProperties, "GuestOsState")
		delete(additionalProperties, "HostUuid")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "PlatformInstanceId")
		delete(additionalProperties, "StorageProvisionedInBytes")
		delete(additionalProperties, "StorageUsedInBytes")
		delete(additionalProperties, "Template")
		delete(additionalProperties, "VmInstanceUuid")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHypervisorVirtualMachineAllOf struct {
	value *HyperflexHypervisorVirtualMachineAllOf
	isSet bool
}

func (v NullableHyperflexHypervisorVirtualMachineAllOf) Get() *HyperflexHypervisorVirtualMachineAllOf {
	return v.value
}

func (v *NullableHyperflexHypervisorVirtualMachineAllOf) Set(val *HyperflexHypervisorVirtualMachineAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHypervisorVirtualMachineAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHypervisorVirtualMachineAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHypervisorVirtualMachineAllOf(val *HyperflexHypervisorVirtualMachineAllOf) *NullableHyperflexHypervisorVirtualMachineAllOf {
	return &NullableHyperflexHypervisorVirtualMachineAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHypervisorVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHypervisorVirtualMachineAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
