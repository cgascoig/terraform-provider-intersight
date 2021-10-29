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
	"reflect"
	"strings"
)

// VirtualizationNetworkInterface Virtual machine network interface configuration data.
type VirtualizationNetworkInterface struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Virtual machine network adaptor type. * `Unknown` - The type of the network adaptor type is unknown. * `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets. * `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * `` - Default network adaptor type supported by the hypervisor.
	AdaptorType *string `json:"AdaptorType,omitempty"`
	// Virtual machine network bridge name.
	Bridge *string `json:"Bridge,omitempty"`
	// Connect the adaptor at virtual machine power on.
	ConnectAtPowerOn *bool `json:"ConnectAtPowerOn,omitempty"`
	// Enable the direct path I/O.
	DirectPathIo *bool `json:"DirectPathIo,omitempty"`
	// Set to true, if IP forwarding is enabled on the NIC.
	IpForwardingEnabled *bool `json:"IpForwardingEnabled,omitempty"`
	// Set to true, if IPv6 address should be allocated for the NIC.
	Ipv6Address *bool `json:"Ipv6Address,omitempty"`
	// Virtual machine network mac address.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of the network interface. This may be different from guest operating system assigned.
	Name *string `json:"Name,omitempty"`
	// Identity of the network to which this network interface belongs.
	NetworkId *string `json:"NetworkId,omitempty"`
	// Identity of the network interface.
	NicId *string `json:"NicId,omitempty"`
	// Order of the NIC attachment to the VM.
	Order *int64 `json:"Order,omitempty"`
	// Allocation mode for NIC addresses e.g. DHCP or static. * `DHCP` - Dynamic IP address allocation using DHCP protocol. * `STATIC_IP` - Assign fixed / static IPs to resources for use. * `IPAM_CALLOUT` - Use callout scripts to query cloud IP allocation tools to assign network parameters. * `PREALLOCATE_IP` - Allows the cloud infrastructure IP allocation to be dynamically provided before the server boots up.
	PrivateIpAllocationMode *string `json:"PrivateIpAllocationMode,omitempty"`
	// Set to true, if public IP should be allocated for the NIC.
	PublicIpAllocate *bool                         `json:"PublicIpAllocate,omitempty"`
	SecurityGroups   []string                      `json:"SecurityGroups,omitempty"`
	StaticIpAddress  []VirtualizationIpAddressInfo `json:"StaticIpAddress,omitempty"`
	// Subnet identifier for the NIC.
	SubnetId             *string `json:"SubnetId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationNetworkInterface VirtualizationNetworkInterface

// NewVirtualizationNetworkInterface instantiates a new VirtualizationNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationNetworkInterface(classId string, objectType string) *VirtualizationNetworkInterface {
	this := VirtualizationNetworkInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adaptorType string = "Unknown"
	this.AdaptorType = &adaptorType
	var privateIpAllocationMode string = "DHCP"
	this.PrivateIpAllocationMode = &privateIpAllocationMode
	return &this
}

// NewVirtualizationNetworkInterfaceWithDefaults instantiates a new VirtualizationNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationNetworkInterfaceWithDefaults() *VirtualizationNetworkInterface {
	this := VirtualizationNetworkInterface{}
	var classId string = "virtualization.NetworkInterface"
	this.ClassId = classId
	var objectType string = "virtualization.NetworkInterface"
	this.ObjectType = objectType
	var adaptorType string = "Unknown"
	this.AdaptorType = &adaptorType
	var privateIpAllocationMode string = "DHCP"
	this.PrivateIpAllocationMode = &privateIpAllocationMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationNetworkInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationNetworkInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationNetworkInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationNetworkInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdaptorType returns the AdaptorType field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetAdaptorType() string {
	if o == nil || o.AdaptorType == nil {
		var ret string
		return ret
	}
	return *o.AdaptorType
}

// GetAdaptorTypeOk returns a tuple with the AdaptorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetAdaptorTypeOk() (*string, bool) {
	if o == nil || o.AdaptorType == nil {
		return nil, false
	}
	return o.AdaptorType, true
}

// HasAdaptorType returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasAdaptorType() bool {
	if o != nil && o.AdaptorType != nil {
		return true
	}

	return false
}

// SetAdaptorType gets a reference to the given string and assigns it to the AdaptorType field.
func (o *VirtualizationNetworkInterface) SetAdaptorType(v string) {
	o.AdaptorType = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetBridge() string {
	if o == nil || o.Bridge == nil {
		var ret string
		return ret
	}
	return *o.Bridge
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetBridgeOk() (*string, bool) {
	if o == nil || o.Bridge == nil {
		return nil, false
	}
	return o.Bridge, true
}

// HasBridge returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasBridge() bool {
	if o != nil && o.Bridge != nil {
		return true
	}

	return false
}

// SetBridge gets a reference to the given string and assigns it to the Bridge field.
func (o *VirtualizationNetworkInterface) SetBridge(v string) {
	o.Bridge = &v
}

// GetConnectAtPowerOn returns the ConnectAtPowerOn field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetConnectAtPowerOn() bool {
	if o == nil || o.ConnectAtPowerOn == nil {
		var ret bool
		return ret
	}
	return *o.ConnectAtPowerOn
}

// GetConnectAtPowerOnOk returns a tuple with the ConnectAtPowerOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetConnectAtPowerOnOk() (*bool, bool) {
	if o == nil || o.ConnectAtPowerOn == nil {
		return nil, false
	}
	return o.ConnectAtPowerOn, true
}

// HasConnectAtPowerOn returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasConnectAtPowerOn() bool {
	if o != nil && o.ConnectAtPowerOn != nil {
		return true
	}

	return false
}

// SetConnectAtPowerOn gets a reference to the given bool and assigns it to the ConnectAtPowerOn field.
func (o *VirtualizationNetworkInterface) SetConnectAtPowerOn(v bool) {
	o.ConnectAtPowerOn = &v
}

// GetDirectPathIo returns the DirectPathIo field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetDirectPathIo() bool {
	if o == nil || o.DirectPathIo == nil {
		var ret bool
		return ret
	}
	return *o.DirectPathIo
}

// GetDirectPathIoOk returns a tuple with the DirectPathIo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetDirectPathIoOk() (*bool, bool) {
	if o == nil || o.DirectPathIo == nil {
		return nil, false
	}
	return o.DirectPathIo, true
}

// HasDirectPathIo returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasDirectPathIo() bool {
	if o != nil && o.DirectPathIo != nil {
		return true
	}

	return false
}

// SetDirectPathIo gets a reference to the given bool and assigns it to the DirectPathIo field.
func (o *VirtualizationNetworkInterface) SetDirectPathIo(v bool) {
	o.DirectPathIo = &v
}

// GetIpForwardingEnabled returns the IpForwardingEnabled field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetIpForwardingEnabled() bool {
	if o == nil || o.IpForwardingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IpForwardingEnabled
}

// GetIpForwardingEnabledOk returns a tuple with the IpForwardingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetIpForwardingEnabledOk() (*bool, bool) {
	if o == nil || o.IpForwardingEnabled == nil {
		return nil, false
	}
	return o.IpForwardingEnabled, true
}

// HasIpForwardingEnabled returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasIpForwardingEnabled() bool {
	if o != nil && o.IpForwardingEnabled != nil {
		return true
	}

	return false
}

// SetIpForwardingEnabled gets a reference to the given bool and assigns it to the IpForwardingEnabled field.
func (o *VirtualizationNetworkInterface) SetIpForwardingEnabled(v bool) {
	o.IpForwardingEnabled = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetIpv6Address() bool {
	if o == nil || o.Ipv6Address == nil {
		var ret bool
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetIpv6AddressOk() (*bool, bool) {
	if o == nil || o.Ipv6Address == nil {
		return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasIpv6Address() bool {
	if o != nil && o.Ipv6Address != nil {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given bool and assigns it to the Ipv6Address field.
func (o *VirtualizationNetworkInterface) SetIpv6Address(v bool) {
	o.Ipv6Address = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationNetworkInterface) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationNetworkInterface) SetName(v string) {
	o.Name = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetNetworkId() string {
	if o == nil || o.NetworkId == nil {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetNetworkIdOk() (*string, bool) {
	if o == nil || o.NetworkId == nil {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasNetworkId() bool {
	if o != nil && o.NetworkId != nil {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *VirtualizationNetworkInterface) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *VirtualizationNetworkInterface) SetNicId(v string) {
	o.NicId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VirtualizationNetworkInterface) SetOrder(v int64) {
	o.Order = &v
}

// GetPrivateIpAllocationMode returns the PrivateIpAllocationMode field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetPrivateIpAllocationMode() string {
	if o == nil || o.PrivateIpAllocationMode == nil {
		var ret string
		return ret
	}
	return *o.PrivateIpAllocationMode
}

// GetPrivateIpAllocationModeOk returns a tuple with the PrivateIpAllocationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetPrivateIpAllocationModeOk() (*string, bool) {
	if o == nil || o.PrivateIpAllocationMode == nil {
		return nil, false
	}
	return o.PrivateIpAllocationMode, true
}

// HasPrivateIpAllocationMode returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasPrivateIpAllocationMode() bool {
	if o != nil && o.PrivateIpAllocationMode != nil {
		return true
	}

	return false
}

// SetPrivateIpAllocationMode gets a reference to the given string and assigns it to the PrivateIpAllocationMode field.
func (o *VirtualizationNetworkInterface) SetPrivateIpAllocationMode(v string) {
	o.PrivateIpAllocationMode = &v
}

// GetPublicIpAllocate returns the PublicIpAllocate field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetPublicIpAllocate() bool {
	if o == nil || o.PublicIpAllocate == nil {
		var ret bool
		return ret
	}
	return *o.PublicIpAllocate
}

// GetPublicIpAllocateOk returns a tuple with the PublicIpAllocate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetPublicIpAllocateOk() (*bool, bool) {
	if o == nil || o.PublicIpAllocate == nil {
		return nil, false
	}
	return o.PublicIpAllocate, true
}

// HasPublicIpAllocate returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasPublicIpAllocate() bool {
	if o != nil && o.PublicIpAllocate != nil {
		return true
	}

	return false
}

// SetPublicIpAllocate gets a reference to the given bool and assigns it to the PublicIpAllocate field.
func (o *VirtualizationNetworkInterface) SetPublicIpAllocate(v bool) {
	o.PublicIpAllocate = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationNetworkInterface) GetSecurityGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationNetworkInterface) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return &o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *VirtualizationNetworkInterface) SetSecurityGroups(v []string) {
	o.SecurityGroups = v
}

// GetStaticIpAddress returns the StaticIpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationNetworkInterface) GetStaticIpAddress() []VirtualizationIpAddressInfo {
	if o == nil {
		var ret []VirtualizationIpAddressInfo
		return ret
	}
	return o.StaticIpAddress
}

// GetStaticIpAddressOk returns a tuple with the StaticIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationNetworkInterface) GetStaticIpAddressOk() (*[]VirtualizationIpAddressInfo, bool) {
	if o == nil || o.StaticIpAddress == nil {
		return nil, false
	}
	return &o.StaticIpAddress, true
}

// HasStaticIpAddress returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasStaticIpAddress() bool {
	if o != nil && o.StaticIpAddress != nil {
		return true
	}

	return false
}

// SetStaticIpAddress gets a reference to the given []VirtualizationIpAddressInfo and assigns it to the StaticIpAddress field.
func (o *VirtualizationNetworkInterface) SetStaticIpAddress(v []VirtualizationIpAddressInfo) {
	o.StaticIpAddress = v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterface) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterface) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterface) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *VirtualizationNetworkInterface) SetSubnetId(v string) {
	o.SubnetId = &v
}

func (o VirtualizationNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdaptorType != nil {
		toSerialize["AdaptorType"] = o.AdaptorType
	}
	if o.Bridge != nil {
		toSerialize["Bridge"] = o.Bridge
	}
	if o.ConnectAtPowerOn != nil {
		toSerialize["ConnectAtPowerOn"] = o.ConnectAtPowerOn
	}
	if o.DirectPathIo != nil {
		toSerialize["DirectPathIo"] = o.DirectPathIo
	}
	if o.IpForwardingEnabled != nil {
		toSerialize["IpForwardingEnabled"] = o.IpForwardingEnabled
	}
	if o.Ipv6Address != nil {
		toSerialize["Ipv6Address"] = o.Ipv6Address
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NetworkId != nil {
		toSerialize["NetworkId"] = o.NetworkId
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.PrivateIpAllocationMode != nil {
		toSerialize["PrivateIpAllocationMode"] = o.PrivateIpAllocationMode
	}
	if o.PublicIpAllocate != nil {
		toSerialize["PublicIpAllocate"] = o.PublicIpAllocate
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.StaticIpAddress != nil {
		toSerialize["StaticIpAddress"] = o.StaticIpAddress
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationNetworkInterface) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationNetworkInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Virtual machine network adaptor type. * `Unknown` - The type of the network adaptor type is unknown. * `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets. * `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * `` - Default network adaptor type supported by the hypervisor.
		AdaptorType *string `json:"AdaptorType,omitempty"`
		// Virtual machine network bridge name.
		Bridge *string `json:"Bridge,omitempty"`
		// Connect the adaptor at virtual machine power on.
		ConnectAtPowerOn *bool `json:"ConnectAtPowerOn,omitempty"`
		// Enable the direct path I/O.
		DirectPathIo *bool `json:"DirectPathIo,omitempty"`
		// Set to true, if IP forwarding is enabled on the NIC.
		IpForwardingEnabled *bool `json:"IpForwardingEnabled,omitempty"`
		// Set to true, if IPv6 address should be allocated for the NIC.
		Ipv6Address *bool `json:"Ipv6Address,omitempty"`
		// Virtual machine network mac address.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Name of the network interface. This may be different from guest operating system assigned.
		Name *string `json:"Name,omitempty"`
		// Identity of the network to which this network interface belongs.
		NetworkId *string `json:"NetworkId,omitempty"`
		// Identity of the network interface.
		NicId *string `json:"NicId,omitempty"`
		// Order of the NIC attachment to the VM.
		Order *int64 `json:"Order,omitempty"`
		// Allocation mode for NIC addresses e.g. DHCP or static. * `DHCP` - Dynamic IP address allocation using DHCP protocol. * `STATIC_IP` - Assign fixed / static IPs to resources for use. * `IPAM_CALLOUT` - Use callout scripts to query cloud IP allocation tools to assign network parameters. * `PREALLOCATE_IP` - Allows the cloud infrastructure IP allocation to be dynamically provided before the server boots up.
		PrivateIpAllocationMode *string `json:"PrivateIpAllocationMode,omitempty"`
		// Set to true, if public IP should be allocated for the NIC.
		PublicIpAllocate *bool                         `json:"PublicIpAllocate,omitempty"`
		SecurityGroups   []string                      `json:"SecurityGroups,omitempty"`
		StaticIpAddress  []VirtualizationIpAddressInfo `json:"StaticIpAddress,omitempty"`
		// Subnet identifier for the NIC.
		SubnetId *string `json:"SubnetId,omitempty"`
	}

	varVirtualizationNetworkInterfaceWithoutEmbeddedStruct := VirtualizationNetworkInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationNetworkInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationNetworkInterface := _VirtualizationNetworkInterface{}
		varVirtualizationNetworkInterface.ClassId = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.ClassId
		varVirtualizationNetworkInterface.ObjectType = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.ObjectType
		varVirtualizationNetworkInterface.AdaptorType = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.AdaptorType
		varVirtualizationNetworkInterface.Bridge = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.Bridge
		varVirtualizationNetworkInterface.ConnectAtPowerOn = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.ConnectAtPowerOn
		varVirtualizationNetworkInterface.DirectPathIo = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.DirectPathIo
		varVirtualizationNetworkInterface.IpForwardingEnabled = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.IpForwardingEnabled
		varVirtualizationNetworkInterface.Ipv6Address = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.Ipv6Address
		varVirtualizationNetworkInterface.MacAddress = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.MacAddress
		varVirtualizationNetworkInterface.Name = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.Name
		varVirtualizationNetworkInterface.NetworkId = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.NetworkId
		varVirtualizationNetworkInterface.NicId = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.NicId
		varVirtualizationNetworkInterface.Order = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.Order
		varVirtualizationNetworkInterface.PrivateIpAllocationMode = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.PrivateIpAllocationMode
		varVirtualizationNetworkInterface.PublicIpAllocate = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.PublicIpAllocate
		varVirtualizationNetworkInterface.SecurityGroups = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.SecurityGroups
		varVirtualizationNetworkInterface.StaticIpAddress = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.StaticIpAddress
		varVirtualizationNetworkInterface.SubnetId = varVirtualizationNetworkInterfaceWithoutEmbeddedStruct.SubnetId
		*o = VirtualizationNetworkInterface(varVirtualizationNetworkInterface)
	} else {
		return err
	}

	varVirtualizationNetworkInterface := _VirtualizationNetworkInterface{}

	err = json.Unmarshal(bytes, &varVirtualizationNetworkInterface)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationNetworkInterface.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdaptorType")
		delete(additionalProperties, "Bridge")
		delete(additionalProperties, "ConnectAtPowerOn")
		delete(additionalProperties, "DirectPathIo")
		delete(additionalProperties, "IpForwardingEnabled")
		delete(additionalProperties, "Ipv6Address")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetworkId")
		delete(additionalProperties, "NicId")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "PrivateIpAllocationMode")
		delete(additionalProperties, "PublicIpAllocate")
		delete(additionalProperties, "SecurityGroups")
		delete(additionalProperties, "StaticIpAddress")
		delete(additionalProperties, "SubnetId")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVirtualizationNetworkInterface struct {
	value *VirtualizationNetworkInterface
	isSet bool
}

func (v NullableVirtualizationNetworkInterface) Get() *VirtualizationNetworkInterface {
	return v.value
}

func (v *NullableVirtualizationNetworkInterface) Set(val *VirtualizationNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationNetworkInterface(val *VirtualizationNetworkInterface) *NullableVirtualizationNetworkInterface {
	return &NullableVirtualizationNetworkInterface{value: val, isSet: true}
}

func (v NullableVirtualizationNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
