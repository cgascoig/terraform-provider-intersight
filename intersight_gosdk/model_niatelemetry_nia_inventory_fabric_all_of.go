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

// NiatelemetryNiaInventoryFabricAllOf Definition of the list of properties defined in 'niatelemetry.NiaInventoryFabric', excluding properties defined in parent classes.
type NiatelemetryNiaInventoryFabricAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the aycast gateway mac.
	AnycastGwMac *string `json:"AnycastGwMac,omitempty"`
	// Counts the number of BGP interfaces that are in established state.
	BgpEstablishedInterfaceCount *int64 `json:"BgpEstablishedInterfaceCount,omitempty"`
	// Count number of active interfaces on border gateways.
	BgwInterfaceUpCount *int64 `json:"BgwInterfaceUpCount,omitempty"`
	// Count number of border gateway spines in the fabric inventory.
	BorderGatewaySpineCount *int64 `json:"BorderGatewaySpineCount,omitempty"`
	// Count number of border leafs in the fabric inventory.
	BorderLeafCount *int64 `json:"BorderLeafCount,omitempty"`
	// Returns the dci subnet range.
	DciSubnetRange *string `json:"DciSubnetRange,omitempty"`
	// Returns the dci subnet target mask.
	DciSubnetTargetMask *string `json:"DciSubnetTargetMask,omitempty"`
	// Returns the value of the dcnmtrackerEnabled field.
	DcnmtrackerEnabled *bool `json:"DcnmtrackerEnabled,omitempty"`
	// Count number of ebgp evpn active interfaces.
	EbgpEvpnLinkUpCount *int64 `json:"EbgpEvpnLinkUpCount,omitempty"`
	// Uniquely identifies a fabric.
	FabricId *string `json:"FabricId,omitempty"`
	// Returns the value of the Name of a fabric.
	FabricName *string `json:"FabricName,omitempty"`
	// Checks if border gateway is present in the fabric inventory.
	IsBgwPresent *bool `json:"IsBgwPresent,omitempty"`
	// Returns if ngoam is enabled.
	IsNgoamEnabled *bool `json:"IsNgoamEnabled,omitempty"`
	// Returns if the scheduled backup is enabled.
	IsScheduledBackUpEnabled *bool `json:"IsScheduledBackUpEnabled,omitempty"`
	// Returns total number of leafs in the fabric.
	LeafCount    *int64                    `json:"LeafCount,omitempty"`
	LogicalLinks []NiatelemetryLogicalLink `json:"LogicalLinks,omitempty"`
	// Returns the count of vnis between sites.
	NxosVniBwSitesCount *int64 `json:"NxosVniBwSitesCount,omitempty"`
	// Returns the count of vrfs between sites.
	NxosVrfBwSitesCount *int64 `json:"NxosVrfBwSitesCount,omitempty"`
	// Returns the value of the nxosVrfCount field.
	NxosVrfCount *int64 `json:"NxosVrfCount,omitempty"`
	// Serial number of device being inventoried. The serial number is unique per device.
	Serial *string `json:"Serial,omitempty"`
	// Name of fabric domain of the controller.
	SiteName *string `json:"SiteName,omitempty"`
	// Returns total number of spines in the fabric.
	SpineCount *int64 `json:"SpineCount,omitempty"`
	// VLAN to VNI mappings configured in the DCNM.
	VlanVniMappings *string `json:"VlanVniMappings,omitempty"`
	// Count number of IP addresses configured in the DCNM networks.
	VniIpCount           *int64                               `json:"VniIpCount,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNiaInventoryFabricAllOf NiatelemetryNiaInventoryFabricAllOf

// NewNiatelemetryNiaInventoryFabricAllOf instantiates a new NiatelemetryNiaInventoryFabricAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaInventoryFabricAllOf(classId string, objectType string) *NiatelemetryNiaInventoryFabricAllOf {
	this := NiatelemetryNiaInventoryFabricAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNiaInventoryFabricAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryFabricAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaInventoryFabricAllOfWithDefaults() *NiatelemetryNiaInventoryFabricAllOf {
	this := NiatelemetryNiaInventoryFabricAllOf{}
	var classId string = "niatelemetry.NiaInventoryFabric"
	this.ClassId = classId
	var objectType string = "niatelemetry.NiaInventoryFabric"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNiaInventoryFabricAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNiaInventoryFabricAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnycastGwMac returns the AnycastGwMac field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMac() string {
	if o == nil || o.AnycastGwMac == nil {
		var ret string
		return ret
	}
	return *o.AnycastGwMac
}

// GetAnycastGwMacOk returns a tuple with the AnycastGwMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMacOk() (*string, bool) {
	if o == nil || o.AnycastGwMac == nil {
		return nil, false
	}
	return o.AnycastGwMac, true
}

// HasAnycastGwMac returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasAnycastGwMac() bool {
	if o != nil && o.AnycastGwMac != nil {
		return true
	}

	return false
}

// SetAnycastGwMac gets a reference to the given string and assigns it to the AnycastGwMac field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetAnycastGwMac(v string) {
	o.AnycastGwMac = &v
}

// GetBgpEstablishedInterfaceCount returns the BgpEstablishedInterfaceCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgpEstablishedInterfaceCount() int64 {
	if o == nil || o.BgpEstablishedInterfaceCount == nil {
		var ret int64
		return ret
	}
	return *o.BgpEstablishedInterfaceCount
}

// GetBgpEstablishedInterfaceCountOk returns a tuple with the BgpEstablishedInterfaceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgpEstablishedInterfaceCountOk() (*int64, bool) {
	if o == nil || o.BgpEstablishedInterfaceCount == nil {
		return nil, false
	}
	return o.BgpEstablishedInterfaceCount, true
}

// HasBgpEstablishedInterfaceCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasBgpEstablishedInterfaceCount() bool {
	if o != nil && o.BgpEstablishedInterfaceCount != nil {
		return true
	}

	return false
}

// SetBgpEstablishedInterfaceCount gets a reference to the given int64 and assigns it to the BgpEstablishedInterfaceCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetBgpEstablishedInterfaceCount(v int64) {
	o.BgpEstablishedInterfaceCount = &v
}

// GetBgwInterfaceUpCount returns the BgwInterfaceUpCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgwInterfaceUpCount() int64 {
	if o == nil || o.BgwInterfaceUpCount == nil {
		var ret int64
		return ret
	}
	return *o.BgwInterfaceUpCount
}

// GetBgwInterfaceUpCountOk returns a tuple with the BgwInterfaceUpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgwInterfaceUpCountOk() (*int64, bool) {
	if o == nil || o.BgwInterfaceUpCount == nil {
		return nil, false
	}
	return o.BgwInterfaceUpCount, true
}

// HasBgwInterfaceUpCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasBgwInterfaceUpCount() bool {
	if o != nil && o.BgwInterfaceUpCount != nil {
		return true
	}

	return false
}

// SetBgwInterfaceUpCount gets a reference to the given int64 and assigns it to the BgwInterfaceUpCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetBgwInterfaceUpCount(v int64) {
	o.BgwInterfaceUpCount = &v
}

// GetBorderGatewaySpineCount returns the BorderGatewaySpineCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderGatewaySpineCount() int64 {
	if o == nil || o.BorderGatewaySpineCount == nil {
		var ret int64
		return ret
	}
	return *o.BorderGatewaySpineCount
}

// GetBorderGatewaySpineCountOk returns a tuple with the BorderGatewaySpineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderGatewaySpineCountOk() (*int64, bool) {
	if o == nil || o.BorderGatewaySpineCount == nil {
		return nil, false
	}
	return o.BorderGatewaySpineCount, true
}

// HasBorderGatewaySpineCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasBorderGatewaySpineCount() bool {
	if o != nil && o.BorderGatewaySpineCount != nil {
		return true
	}

	return false
}

// SetBorderGatewaySpineCount gets a reference to the given int64 and assigns it to the BorderGatewaySpineCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetBorderGatewaySpineCount(v int64) {
	o.BorderGatewaySpineCount = &v
}

// GetBorderLeafCount returns the BorderLeafCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderLeafCount() int64 {
	if o == nil || o.BorderLeafCount == nil {
		var ret int64
		return ret
	}
	return *o.BorderLeafCount
}

// GetBorderLeafCountOk returns a tuple with the BorderLeafCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderLeafCountOk() (*int64, bool) {
	if o == nil || o.BorderLeafCount == nil {
		return nil, false
	}
	return o.BorderLeafCount, true
}

// HasBorderLeafCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasBorderLeafCount() bool {
	if o != nil && o.BorderLeafCount != nil {
		return true
	}

	return false
}

// SetBorderLeafCount gets a reference to the given int64 and assigns it to the BorderLeafCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetBorderLeafCount(v int64) {
	o.BorderLeafCount = &v
}

// GetDciSubnetRange returns the DciSubnetRange field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetRange() string {
	if o == nil || o.DciSubnetRange == nil {
		var ret string
		return ret
	}
	return *o.DciSubnetRange
}

// GetDciSubnetRangeOk returns a tuple with the DciSubnetRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetRangeOk() (*string, bool) {
	if o == nil || o.DciSubnetRange == nil {
		return nil, false
	}
	return o.DciSubnetRange, true
}

// HasDciSubnetRange returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasDciSubnetRange() bool {
	if o != nil && o.DciSubnetRange != nil {
		return true
	}

	return false
}

// SetDciSubnetRange gets a reference to the given string and assigns it to the DciSubnetRange field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetDciSubnetRange(v string) {
	o.DciSubnetRange = &v
}

// GetDciSubnetTargetMask returns the DciSubnetTargetMask field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetTargetMask() string {
	if o == nil || o.DciSubnetTargetMask == nil {
		var ret string
		return ret
	}
	return *o.DciSubnetTargetMask
}

// GetDciSubnetTargetMaskOk returns a tuple with the DciSubnetTargetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetTargetMaskOk() (*string, bool) {
	if o == nil || o.DciSubnetTargetMask == nil {
		return nil, false
	}
	return o.DciSubnetTargetMask, true
}

// HasDciSubnetTargetMask returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasDciSubnetTargetMask() bool {
	if o != nil && o.DciSubnetTargetMask != nil {
		return true
	}

	return false
}

// SetDciSubnetTargetMask gets a reference to the given string and assigns it to the DciSubnetTargetMask field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetDciSubnetTargetMask(v string) {
	o.DciSubnetTargetMask = &v
}

// GetDcnmtrackerEnabled returns the DcnmtrackerEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabled() bool {
	if o == nil || o.DcnmtrackerEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DcnmtrackerEnabled
}

// GetDcnmtrackerEnabledOk returns a tuple with the DcnmtrackerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabledOk() (*bool, bool) {
	if o == nil || o.DcnmtrackerEnabled == nil {
		return nil, false
	}
	return o.DcnmtrackerEnabled, true
}

// HasDcnmtrackerEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasDcnmtrackerEnabled() bool {
	if o != nil && o.DcnmtrackerEnabled != nil {
		return true
	}

	return false
}

// SetDcnmtrackerEnabled gets a reference to the given bool and assigns it to the DcnmtrackerEnabled field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetDcnmtrackerEnabled(v bool) {
	o.DcnmtrackerEnabled = &v
}

// GetEbgpEvpnLinkUpCount returns the EbgpEvpnLinkUpCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetEbgpEvpnLinkUpCount() int64 {
	if o == nil || o.EbgpEvpnLinkUpCount == nil {
		var ret int64
		return ret
	}
	return *o.EbgpEvpnLinkUpCount
}

// GetEbgpEvpnLinkUpCountOk returns a tuple with the EbgpEvpnLinkUpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetEbgpEvpnLinkUpCountOk() (*int64, bool) {
	if o == nil || o.EbgpEvpnLinkUpCount == nil {
		return nil, false
	}
	return o.EbgpEvpnLinkUpCount, true
}

// HasEbgpEvpnLinkUpCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasEbgpEvpnLinkUpCount() bool {
	if o != nil && o.EbgpEvpnLinkUpCount != nil {
		return true
	}

	return false
}

// SetEbgpEvpnLinkUpCount gets a reference to the given int64 and assigns it to the EbgpEvpnLinkUpCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetEbgpEvpnLinkUpCount(v int64) {
	o.EbgpEvpnLinkUpCount = &v
}

// GetFabricId returns the FabricId field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricId() string {
	if o == nil || o.FabricId == nil {
		var ret string
		return ret
	}
	return *o.FabricId
}

// GetFabricIdOk returns a tuple with the FabricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricIdOk() (*string, bool) {
	if o == nil || o.FabricId == nil {
		return nil, false
	}
	return o.FabricId, true
}

// HasFabricId returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricId() bool {
	if o != nil && o.FabricId != nil {
		return true
	}

	return false
}

// SetFabricId gets a reference to the given string and assigns it to the FabricId field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricId(v string) {
	o.FabricId = &v
}

// GetFabricName returns the FabricName field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricName() string {
	if o == nil || o.FabricName == nil {
		var ret string
		return ret
	}
	return *o.FabricName
}

// GetFabricNameOk returns a tuple with the FabricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricNameOk() (*string, bool) {
	if o == nil || o.FabricName == nil {
		return nil, false
	}
	return o.FabricName, true
}

// HasFabricName returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricName() bool {
	if o != nil && o.FabricName != nil {
		return true
	}

	return false
}

// SetFabricName gets a reference to the given string and assigns it to the FabricName field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricName(v string) {
	o.FabricName = &v
}

// GetIsBgwPresent returns the IsBgwPresent field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsBgwPresent() bool {
	if o == nil || o.IsBgwPresent == nil {
		var ret bool
		return ret
	}
	return *o.IsBgwPresent
}

// GetIsBgwPresentOk returns a tuple with the IsBgwPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsBgwPresentOk() (*bool, bool) {
	if o == nil || o.IsBgwPresent == nil {
		return nil, false
	}
	return o.IsBgwPresent, true
}

// HasIsBgwPresent returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsBgwPresent() bool {
	if o != nil && o.IsBgwPresent != nil {
		return true
	}

	return false
}

// SetIsBgwPresent gets a reference to the given bool and assigns it to the IsBgwPresent field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsBgwPresent(v bool) {
	o.IsBgwPresent = &v
}

// GetIsNgoamEnabled returns the IsNgoamEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabled() bool {
	if o == nil || o.IsNgoamEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsNgoamEnabled
}

// GetIsNgoamEnabledOk returns a tuple with the IsNgoamEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabledOk() (*bool, bool) {
	if o == nil || o.IsNgoamEnabled == nil {
		return nil, false
	}
	return o.IsNgoamEnabled, true
}

// HasIsNgoamEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsNgoamEnabled() bool {
	if o != nil && o.IsNgoamEnabled != nil {
		return true
	}

	return false
}

// SetIsNgoamEnabled gets a reference to the given bool and assigns it to the IsNgoamEnabled field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsNgoamEnabled(v bool) {
	o.IsNgoamEnabled = &v
}

// GetIsScheduledBackUpEnabled returns the IsScheduledBackUpEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabled() bool {
	if o == nil || o.IsScheduledBackUpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsScheduledBackUpEnabled
}

// GetIsScheduledBackUpEnabledOk returns a tuple with the IsScheduledBackUpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabledOk() (*bool, bool) {
	if o == nil || o.IsScheduledBackUpEnabled == nil {
		return nil, false
	}
	return o.IsScheduledBackUpEnabled, true
}

// HasIsScheduledBackUpEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsScheduledBackUpEnabled() bool {
	if o != nil && o.IsScheduledBackUpEnabled != nil {
		return true
	}

	return false
}

// SetIsScheduledBackUpEnabled gets a reference to the given bool and assigns it to the IsScheduledBackUpEnabled field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsScheduledBackUpEnabled(v bool) {
	o.IsScheduledBackUpEnabled = &v
}

// GetLeafCount returns the LeafCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCount() int64 {
	if o == nil || o.LeafCount == nil {
		var ret int64
		return ret
	}
	return *o.LeafCount
}

// GetLeafCountOk returns a tuple with the LeafCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCountOk() (*int64, bool) {
	if o == nil || o.LeafCount == nil {
		return nil, false
	}
	return o.LeafCount, true
}

// HasLeafCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasLeafCount() bool {
	if o != nil && o.LeafCount != nil {
		return true
	}

	return false
}

// SetLeafCount gets a reference to the given int64 and assigns it to the LeafCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetLeafCount(v int64) {
	o.LeafCount = &v
}

// GetLogicalLinks returns the LogicalLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinks() []NiatelemetryLogicalLink {
	if o == nil {
		var ret []NiatelemetryLogicalLink
		return ret
	}
	return o.LogicalLinks
}

// GetLogicalLinksOk returns a tuple with the LogicalLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinksOk() (*[]NiatelemetryLogicalLink, bool) {
	if o == nil || o.LogicalLinks == nil {
		return nil, false
	}
	return &o.LogicalLinks, true
}

// HasLogicalLinks returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasLogicalLinks() bool {
	if o != nil && o.LogicalLinks != nil {
		return true
	}

	return false
}

// SetLogicalLinks gets a reference to the given []NiatelemetryLogicalLink and assigns it to the LogicalLinks field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetLogicalLinks(v []NiatelemetryLogicalLink) {
	o.LogicalLinks = v
}

// GetNxosVniBwSitesCount returns the NxosVniBwSitesCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVniBwSitesCount() int64 {
	if o == nil || o.NxosVniBwSitesCount == nil {
		var ret int64
		return ret
	}
	return *o.NxosVniBwSitesCount
}

// GetNxosVniBwSitesCountOk returns a tuple with the NxosVniBwSitesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVniBwSitesCountOk() (*int64, bool) {
	if o == nil || o.NxosVniBwSitesCount == nil {
		return nil, false
	}
	return o.NxosVniBwSitesCount, true
}

// HasNxosVniBwSitesCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVniBwSitesCount() bool {
	if o != nil && o.NxosVniBwSitesCount != nil {
		return true
	}

	return false
}

// SetNxosVniBwSitesCount gets a reference to the given int64 and assigns it to the NxosVniBwSitesCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVniBwSitesCount(v int64) {
	o.NxosVniBwSitesCount = &v
}

// GetNxosVrfBwSitesCount returns the NxosVrfBwSitesCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfBwSitesCount() int64 {
	if o == nil || o.NxosVrfBwSitesCount == nil {
		var ret int64
		return ret
	}
	return *o.NxosVrfBwSitesCount
}

// GetNxosVrfBwSitesCountOk returns a tuple with the NxosVrfBwSitesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfBwSitesCountOk() (*int64, bool) {
	if o == nil || o.NxosVrfBwSitesCount == nil {
		return nil, false
	}
	return o.NxosVrfBwSitesCount, true
}

// HasNxosVrfBwSitesCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVrfBwSitesCount() bool {
	if o != nil && o.NxosVrfBwSitesCount != nil {
		return true
	}

	return false
}

// SetNxosVrfBwSitesCount gets a reference to the given int64 and assigns it to the NxosVrfBwSitesCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVrfBwSitesCount(v int64) {
	o.NxosVrfBwSitesCount = &v
}

// GetNxosVrfCount returns the NxosVrfCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCount() int64 {
	if o == nil || o.NxosVrfCount == nil {
		var ret int64
		return ret
	}
	return *o.NxosVrfCount
}

// GetNxosVrfCountOk returns a tuple with the NxosVrfCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCountOk() (*int64, bool) {
	if o == nil || o.NxosVrfCount == nil {
		return nil, false
	}
	return o.NxosVrfCount, true
}

// HasNxosVrfCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVrfCount() bool {
	if o != nil && o.NxosVrfCount != nil {
		return true
	}

	return false
}

// SetNxosVrfCount gets a reference to the given int64 and assigns it to the NxosVrfCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVrfCount(v int64) {
	o.NxosVrfCount = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSpineCount returns the SpineCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCount() int64 {
	if o == nil || o.SpineCount == nil {
		var ret int64
		return ret
	}
	return *o.SpineCount
}

// GetSpineCountOk returns a tuple with the SpineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCountOk() (*int64, bool) {
	if o == nil || o.SpineCount == nil {
		return nil, false
	}
	return o.SpineCount, true
}

// HasSpineCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasSpineCount() bool {
	if o != nil && o.SpineCount != nil {
		return true
	}

	return false
}

// SetSpineCount gets a reference to the given int64 and assigns it to the SpineCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetSpineCount(v int64) {
	o.SpineCount = &v
}

// GetVlanVniMappings returns the VlanVniMappings field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetVlanVniMappings() string {
	if o == nil || o.VlanVniMappings == nil {
		var ret string
		return ret
	}
	return *o.VlanVniMappings
}

// GetVlanVniMappingsOk returns a tuple with the VlanVniMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetVlanVniMappingsOk() (*string, bool) {
	if o == nil || o.VlanVniMappings == nil {
		return nil, false
	}
	return o.VlanVniMappings, true
}

// HasVlanVniMappings returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasVlanVniMappings() bool {
	if o != nil && o.VlanVniMappings != nil {
		return true
	}

	return false
}

// SetVlanVniMappings gets a reference to the given string and assigns it to the VlanVniMappings field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetVlanVniMappings(v string) {
	o.VlanVniMappings = &v
}

// GetVniIpCount returns the VniIpCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetVniIpCount() int64 {
	if o == nil || o.VniIpCount == nil {
		var ret int64
		return ret
	}
	return *o.VniIpCount
}

// GetVniIpCountOk returns a tuple with the VniIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetVniIpCountOk() (*int64, bool) {
	if o == nil || o.VniIpCount == nil {
		return nil, false
	}
	return o.VniIpCount, true
}

// HasVniIpCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasVniIpCount() bool {
	if o != nil && o.VniIpCount != nil {
		return true
	}

	return false
}

// SetVniIpCount gets a reference to the given int64 and assigns it to the VniIpCount field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetVniIpCount(v int64) {
	o.VniIpCount = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryFabricAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNiaInventoryFabricAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNiaInventoryFabricAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AnycastGwMac != nil {
		toSerialize["AnycastGwMac"] = o.AnycastGwMac
	}
	if o.BgpEstablishedInterfaceCount != nil {
		toSerialize["BgpEstablishedInterfaceCount"] = o.BgpEstablishedInterfaceCount
	}
	if o.BgwInterfaceUpCount != nil {
		toSerialize["BgwInterfaceUpCount"] = o.BgwInterfaceUpCount
	}
	if o.BorderGatewaySpineCount != nil {
		toSerialize["BorderGatewaySpineCount"] = o.BorderGatewaySpineCount
	}
	if o.BorderLeafCount != nil {
		toSerialize["BorderLeafCount"] = o.BorderLeafCount
	}
	if o.DciSubnetRange != nil {
		toSerialize["DciSubnetRange"] = o.DciSubnetRange
	}
	if o.DciSubnetTargetMask != nil {
		toSerialize["DciSubnetTargetMask"] = o.DciSubnetTargetMask
	}
	if o.DcnmtrackerEnabled != nil {
		toSerialize["DcnmtrackerEnabled"] = o.DcnmtrackerEnabled
	}
	if o.EbgpEvpnLinkUpCount != nil {
		toSerialize["EbgpEvpnLinkUpCount"] = o.EbgpEvpnLinkUpCount
	}
	if o.FabricId != nil {
		toSerialize["FabricId"] = o.FabricId
	}
	if o.FabricName != nil {
		toSerialize["FabricName"] = o.FabricName
	}
	if o.IsBgwPresent != nil {
		toSerialize["IsBgwPresent"] = o.IsBgwPresent
	}
	if o.IsNgoamEnabled != nil {
		toSerialize["IsNgoamEnabled"] = o.IsNgoamEnabled
	}
	if o.IsScheduledBackUpEnabled != nil {
		toSerialize["IsScheduledBackUpEnabled"] = o.IsScheduledBackUpEnabled
	}
	if o.LeafCount != nil {
		toSerialize["LeafCount"] = o.LeafCount
	}
	if o.LogicalLinks != nil {
		toSerialize["LogicalLinks"] = o.LogicalLinks
	}
	if o.NxosVniBwSitesCount != nil {
		toSerialize["NxosVniBwSitesCount"] = o.NxosVniBwSitesCount
	}
	if o.NxosVrfBwSitesCount != nil {
		toSerialize["NxosVrfBwSitesCount"] = o.NxosVrfBwSitesCount
	}
	if o.NxosVrfCount != nil {
		toSerialize["NxosVrfCount"] = o.NxosVrfCount
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.SpineCount != nil {
		toSerialize["SpineCount"] = o.SpineCount
	}
	if o.VlanVniMappings != nil {
		toSerialize["VlanVniMappings"] = o.VlanVniMappings
	}
	if o.VniIpCount != nil {
		toSerialize["VniIpCount"] = o.VniIpCount
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNiaInventoryFabricAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNiaInventoryFabricAllOf := _NiatelemetryNiaInventoryFabricAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNiaInventoryFabricAllOf); err == nil {
		*o = NiatelemetryNiaInventoryFabricAllOf(varNiatelemetryNiaInventoryFabricAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AnycastGwMac")
		delete(additionalProperties, "BgpEstablishedInterfaceCount")
		delete(additionalProperties, "BgwInterfaceUpCount")
		delete(additionalProperties, "BorderGatewaySpineCount")
		delete(additionalProperties, "BorderLeafCount")
		delete(additionalProperties, "DciSubnetRange")
		delete(additionalProperties, "DciSubnetTargetMask")
		delete(additionalProperties, "DcnmtrackerEnabled")
		delete(additionalProperties, "EbgpEvpnLinkUpCount")
		delete(additionalProperties, "FabricId")
		delete(additionalProperties, "FabricName")
		delete(additionalProperties, "IsBgwPresent")
		delete(additionalProperties, "IsNgoamEnabled")
		delete(additionalProperties, "IsScheduledBackUpEnabled")
		delete(additionalProperties, "LeafCount")
		delete(additionalProperties, "LogicalLinks")
		delete(additionalProperties, "NxosVniBwSitesCount")
		delete(additionalProperties, "NxosVrfBwSitesCount")
		delete(additionalProperties, "NxosVrfCount")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SpineCount")
		delete(additionalProperties, "VlanVniMappings")
		delete(additionalProperties, "VniIpCount")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNiaInventoryFabricAllOf struct {
	value *NiatelemetryNiaInventoryFabricAllOf
	isSet bool
}

func (v NullableNiatelemetryNiaInventoryFabricAllOf) Get() *NiatelemetryNiaInventoryFabricAllOf {
	return v.value
}

func (v *NullableNiatelemetryNiaInventoryFabricAllOf) Set(val *NiatelemetryNiaInventoryFabricAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaInventoryFabricAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaInventoryFabricAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaInventoryFabricAllOf(val *NiatelemetryNiaInventoryFabricAllOf) *NullableNiatelemetryNiaInventoryFabricAllOf {
	return &NullableNiatelemetryNiaInventoryFabricAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaInventoryFabricAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaInventoryFabricAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
