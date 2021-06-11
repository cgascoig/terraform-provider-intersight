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

// EquipmentPsuAllOf Definition of the list of properties defined in 'equipment.Psu', excluding properties defined in parent classes.
type EquipmentPsuAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field is to provide description for the power supply unit.
	Description *string  `json:"Description,omitempty"`
	OperReason  []string `json:"OperReason,omitempty"`
	// This field identifies the psu operational state.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the Part Number for this Power Supply Unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for the Power Supply.
	Pid *string `json:"Pid,omitempty"`
	// This field identifies the Firmware Version of the Power Supply.
	PsuFwVersion *string `json:"PsuFwVersion,omitempty"`
	// This represents power supply unit identifier in chassis/server.
	PsuId *int64 `json:"PsuId,omitempty"`
	// This field identifies the input source for the Power Supply.
	PsuInputSrc *string `json:"PsuInputSrc,omitempty"`
	// This field identifies the type of the Power Supply.
	PsuType *string `json:"PsuType,omitempty"`
	// This field identifies the Wattage of the Power Supply.
	PsuWattage *string `json:"PsuWattage,omitempty"`
	// This field identifies the Stockkeeping Unit for this Power Supply.
	Sku *string `json:"Sku,omitempty"`
	// This field identifies the Vendor ID for this Power Supply Unit.
	Vid *string `json:"Vid,omitempty"`
	// This field is used to indicate the voltage state for this Power Supply.
	Voltage                *string                              `json:"Voltage,omitempty"`
	ComputeRackUnit        *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	EquipmentChassis       *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	EquipmentFex           *EquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
	EquipmentRackEnclosure *EquipmentRackEnclosureRelationship  `json:"EquipmentRackEnclosure,omitempty"`
	InventoryDeviceInfo    *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement         *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _EquipmentPsuAllOf EquipmentPsuAllOf

// NewEquipmentPsuAllOf instantiates a new EquipmentPsuAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentPsuAllOf(classId string, objectType string) *EquipmentPsuAllOf {
	this := EquipmentPsuAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentPsuAllOfWithDefaults instantiates a new EquipmentPsuAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentPsuAllOfWithDefaults() *EquipmentPsuAllOf {
	this := EquipmentPsuAllOf{}
	var classId string = "equipment.Psu"
	this.ClassId = classId
	var objectType string = "equipment.Psu"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentPsuAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentPsuAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentPsuAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentPsuAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentPsuAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentPsuAllOf) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentPsuAllOf) GetOperReasonOk() (*[]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return &o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentPsuAllOf) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentPsuAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentPsuAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentPsuAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetPsuFwVersion returns the PsuFwVersion field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetPsuFwVersion() string {
	if o == nil || o.PsuFwVersion == nil {
		var ret string
		return ret
	}
	return *o.PsuFwVersion
}

// GetPsuFwVersionOk returns a tuple with the PsuFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetPsuFwVersionOk() (*string, bool) {
	if o == nil || o.PsuFwVersion == nil {
		return nil, false
	}
	return o.PsuFwVersion, true
}

// HasPsuFwVersion returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasPsuFwVersion() bool {
	if o != nil && o.PsuFwVersion != nil {
		return true
	}

	return false
}

// SetPsuFwVersion gets a reference to the given string and assigns it to the PsuFwVersion field.
func (o *EquipmentPsuAllOf) SetPsuFwVersion(v string) {
	o.PsuFwVersion = &v
}

// GetPsuId returns the PsuId field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetPsuId() int64 {
	if o == nil || o.PsuId == nil {
		var ret int64
		return ret
	}
	return *o.PsuId
}

// GetPsuIdOk returns a tuple with the PsuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetPsuIdOk() (*int64, bool) {
	if o == nil || o.PsuId == nil {
		return nil, false
	}
	return o.PsuId, true
}

// HasPsuId returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasPsuId() bool {
	if o != nil && o.PsuId != nil {
		return true
	}

	return false
}

// SetPsuId gets a reference to the given int64 and assigns it to the PsuId field.
func (o *EquipmentPsuAllOf) SetPsuId(v int64) {
	o.PsuId = &v
}

// GetPsuInputSrc returns the PsuInputSrc field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetPsuInputSrc() string {
	if o == nil || o.PsuInputSrc == nil {
		var ret string
		return ret
	}
	return *o.PsuInputSrc
}

// GetPsuInputSrcOk returns a tuple with the PsuInputSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetPsuInputSrcOk() (*string, bool) {
	if o == nil || o.PsuInputSrc == nil {
		return nil, false
	}
	return o.PsuInputSrc, true
}

// HasPsuInputSrc returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasPsuInputSrc() bool {
	if o != nil && o.PsuInputSrc != nil {
		return true
	}

	return false
}

// SetPsuInputSrc gets a reference to the given string and assigns it to the PsuInputSrc field.
func (o *EquipmentPsuAllOf) SetPsuInputSrc(v string) {
	o.PsuInputSrc = &v
}

// GetPsuType returns the PsuType field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetPsuType() string {
	if o == nil || o.PsuType == nil {
		var ret string
		return ret
	}
	return *o.PsuType
}

// GetPsuTypeOk returns a tuple with the PsuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetPsuTypeOk() (*string, bool) {
	if o == nil || o.PsuType == nil {
		return nil, false
	}
	return o.PsuType, true
}

// HasPsuType returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasPsuType() bool {
	if o != nil && o.PsuType != nil {
		return true
	}

	return false
}

// SetPsuType gets a reference to the given string and assigns it to the PsuType field.
func (o *EquipmentPsuAllOf) SetPsuType(v string) {
	o.PsuType = &v
}

// GetPsuWattage returns the PsuWattage field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetPsuWattage() string {
	if o == nil || o.PsuWattage == nil {
		var ret string
		return ret
	}
	return *o.PsuWattage
}

// GetPsuWattageOk returns a tuple with the PsuWattage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetPsuWattageOk() (*string, bool) {
	if o == nil || o.PsuWattage == nil {
		return nil, false
	}
	return o.PsuWattage, true
}

// HasPsuWattage returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasPsuWattage() bool {
	if o != nil && o.PsuWattage != nil {
		return true
	}

	return false
}

// SetPsuWattage gets a reference to the given string and assigns it to the PsuWattage field.
func (o *EquipmentPsuAllOf) SetPsuWattage(v string) {
	o.PsuWattage = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentPsuAllOf) SetSku(v string) {
	o.Sku = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentPsuAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetVoltage() string {
	if o == nil || o.Voltage == nil {
		var ret string
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetVoltageOk() (*string, bool) {
	if o == nil || o.Voltage == nil {
		return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasVoltage() bool {
	if o != nil && o.Voltage != nil {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given string and assigns it to the Voltage field.
func (o *EquipmentPsuAllOf) SetVoltage(v string) {
	o.Voltage = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *EquipmentPsuAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentPsuAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentFex returns the EquipmentFex field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetEquipmentFex() EquipmentFexRelationship {
	if o == nil || o.EquipmentFex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.EquipmentFex
}

// GetEquipmentFexOk returns a tuple with the EquipmentFex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetEquipmentFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.EquipmentFex == nil {
		return nil, false
	}
	return o.EquipmentFex, true
}

// HasEquipmentFex returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasEquipmentFex() bool {
	if o != nil && o.EquipmentFex != nil {
		return true
	}

	return false
}

// SetEquipmentFex gets a reference to the given EquipmentFexRelationship and assigns it to the EquipmentFex field.
func (o *EquipmentPsuAllOf) SetEquipmentFex(v EquipmentFexRelationship) {
	o.EquipmentFex = &v
}

// GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship {
	if o == nil || o.EquipmentRackEnclosure == nil {
		var ret EquipmentRackEnclosureRelationship
		return ret
	}
	return *o.EquipmentRackEnclosure
}

// GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool) {
	if o == nil || o.EquipmentRackEnclosure == nil {
		return nil, false
	}
	return o.EquipmentRackEnclosure, true
}

// HasEquipmentRackEnclosure returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasEquipmentRackEnclosure() bool {
	if o != nil && o.EquipmentRackEnclosure != nil {
		return true
	}

	return false
}

// SetEquipmentRackEnclosure gets a reference to the given EquipmentRackEnclosureRelationship and assigns it to the EquipmentRackEnclosure field.
func (o *EquipmentPsuAllOf) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship) {
	o.EquipmentRackEnclosure = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentPsuAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentPsuAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentPsuAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentPsuAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentPsuAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentPsuAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.PsuFwVersion != nil {
		toSerialize["PsuFwVersion"] = o.PsuFwVersion
	}
	if o.PsuId != nil {
		toSerialize["PsuId"] = o.PsuId
	}
	if o.PsuInputSrc != nil {
		toSerialize["PsuInputSrc"] = o.PsuInputSrc
	}
	if o.PsuType != nil {
		toSerialize["PsuType"] = o.PsuType
	}
	if o.PsuWattage != nil {
		toSerialize["PsuWattage"] = o.PsuWattage
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.Voltage != nil {
		toSerialize["Voltage"] = o.Voltage
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentFex != nil {
		toSerialize["EquipmentFex"] = o.EquipmentFex
	}
	if o.EquipmentRackEnclosure != nil {
		toSerialize["EquipmentRackEnclosure"] = o.EquipmentRackEnclosure
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentPsuAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentPsuAllOf := _EquipmentPsuAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentPsuAllOf); err == nil {
		*o = EquipmentPsuAllOf(varEquipmentPsuAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "PsuFwVersion")
		delete(additionalProperties, "PsuId")
		delete(additionalProperties, "PsuInputSrc")
		delete(additionalProperties, "PsuType")
		delete(additionalProperties, "PsuWattage")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "Voltage")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentFex")
		delete(additionalProperties, "EquipmentRackEnclosure")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentPsuAllOf struct {
	value *EquipmentPsuAllOf
	isSet bool
}

func (v NullableEquipmentPsuAllOf) Get() *EquipmentPsuAllOf {
	return v.value
}

func (v *NullableEquipmentPsuAllOf) Set(val *EquipmentPsuAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentPsuAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentPsuAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentPsuAllOf(val *EquipmentPsuAllOf) *NullableEquipmentPsuAllOf {
	return &NullableEquipmentPsuAllOf{value: val, isSet: true}
}

func (v NullableEquipmentPsuAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentPsuAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
