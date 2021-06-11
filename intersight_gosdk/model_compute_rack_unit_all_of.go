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

// ComputeRackUnitAllOf Definition of the list of properties defined in 'compute.RackUnit', excluding properties defined in parent classes.
type ComputeRackUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Connectivity Status of RackUnit to Switch - A or B or AB.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// RackUnit ID that uniquely identifies the server.
	ServerId *int64 `json:"ServerId,omitempty"`
	// To maintain the Topology workflow run status.
	TopologyScanStatus *string `json:"TopologyScanStatus,omitempty"`
	// An array of relationships to adapterUnit resources.
	Adapters     []AdapterUnitRelationship `json:"Adapters,omitempty"`
	BiosBootmode *BiosBootModeRelationship `json:"BiosBootmode,omitempty"`
	// An array of relationships to biosUnit resources.
	Biosunits          []BiosUnitRelationship            `json:"Biosunits,omitempty"`
	Bmc                *ManagementControllerRelationship `json:"Bmc,omitempty"`
	Board              *ComputeBoardRelationship         `json:"Board,omitempty"`
	BootDeviceBootmode *BootDeviceBootModeRelationship   `json:"BootDeviceBootmode,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	Fanmodules []EquipmentFanModuleRelationship `json:"Fanmodules,omitempty"`
	// An array of relationships to inventoryGenericInventoryHolder resources.
	GenericInventoryHolders []InventoryGenericInventoryHolderRelationship `json:"GenericInventoryHolders,omitempty"`
	// An array of relationships to graphicsCard resources.
	GraphicsCards       []GraphicsCardRelationship       `json:"GraphicsCards,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	LocatorLed          *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
	// An array of relationships to memoryArray resources.
	MemoryArrays []MemoryArrayRelationship `json:"MemoryArrays,omitempty"`
	// An array of relationships to pciDevice resources.
	PciDevices []PciDeviceRelationship `json:"PciDevices,omitempty"`
	// An array of relationships to processorUnit resources.
	Processors []ProcessorUnitRelationship `json:"Processors,omitempty"`
	// An array of relationships to equipmentPsu resources.
	Psus              []EquipmentPsuRelationship              `json:"Psus,omitempty"`
	RackEnclosureSlot *EquipmentRackEnclosureSlotRelationship `json:"RackEnclosureSlot,omitempty"`
	RegisteredDevice  *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageSasExpander resources.
	SasExpanders []StorageSasExpanderRelationship `json:"SasExpanders,omitempty"`
	// An array of relationships to storageController resources.
	StorageControllers []StorageControllerRelationship `json:"StorageControllers,omitempty"`
	// An array of relationships to storageEnclosure resources.
	StorageEnclosures    []StorageEnclosureRelationship `json:"StorageEnclosures,omitempty"`
	TopSystem            *TopSystemRelationship         `json:"TopSystem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeRackUnitAllOf ComputeRackUnitAllOf

// NewComputeRackUnitAllOf instantiates a new ComputeRackUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeRackUnitAllOf(classId string, objectType string) *ComputeRackUnitAllOf {
	this := ComputeRackUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeRackUnitAllOfWithDefaults instantiates a new ComputeRackUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeRackUnitAllOfWithDefaults() *ComputeRackUnitAllOf {
	this := ComputeRackUnitAllOf{}
	var classId string = "compute.RackUnit"
	this.ClassId = classId
	var objectType string = "compute.RackUnit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeRackUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeRackUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeRackUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeRackUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *ComputeRackUnitAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetServerId() int64 {
	if o == nil || o.ServerId == nil {
		var ret int64
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetServerIdOk() (*int64, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int64 and assigns it to the ServerId field.
func (o *ComputeRackUnitAllOf) SetServerId(v int64) {
	o.ServerId = &v
}

// GetTopologyScanStatus returns the TopologyScanStatus field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetTopologyScanStatus() string {
	if o == nil || o.TopologyScanStatus == nil {
		var ret string
		return ret
	}
	return *o.TopologyScanStatus
}

// GetTopologyScanStatusOk returns a tuple with the TopologyScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetTopologyScanStatusOk() (*string, bool) {
	if o == nil || o.TopologyScanStatus == nil {
		return nil, false
	}
	return o.TopologyScanStatus, true
}

// HasTopologyScanStatus returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasTopologyScanStatus() bool {
	if o != nil && o.TopologyScanStatus != nil {
		return true
	}

	return false
}

// SetTopologyScanStatus gets a reference to the given string and assigns it to the TopologyScanStatus field.
func (o *ComputeRackUnitAllOf) SetTopologyScanStatus(v string) {
	o.TopologyScanStatus = &v
}

// GetAdapters returns the Adapters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetAdapters() []AdapterUnitRelationship {
	if o == nil {
		var ret []AdapterUnitRelationship
		return ret
	}
	return o.Adapters
}

// GetAdaptersOk returns a tuple with the Adapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetAdaptersOk() (*[]AdapterUnitRelationship, bool) {
	if o == nil || o.Adapters == nil {
		return nil, false
	}
	return &o.Adapters, true
}

// HasAdapters returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasAdapters() bool {
	if o != nil && o.Adapters != nil {
		return true
	}

	return false
}

// SetAdapters gets a reference to the given []AdapterUnitRelationship and assigns it to the Adapters field.
func (o *ComputeRackUnitAllOf) SetAdapters(v []AdapterUnitRelationship) {
	o.Adapters = v
}

// GetBiosBootmode returns the BiosBootmode field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetBiosBootmode() BiosBootModeRelationship {
	if o == nil || o.BiosBootmode == nil {
		var ret BiosBootModeRelationship
		return ret
	}
	return *o.BiosBootmode
}

// GetBiosBootmodeOk returns a tuple with the BiosBootmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetBiosBootmodeOk() (*BiosBootModeRelationship, bool) {
	if o == nil || o.BiosBootmode == nil {
		return nil, false
	}
	return o.BiosBootmode, true
}

// HasBiosBootmode returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasBiosBootmode() bool {
	if o != nil && o.BiosBootmode != nil {
		return true
	}

	return false
}

// SetBiosBootmode gets a reference to the given BiosBootModeRelationship and assigns it to the BiosBootmode field.
func (o *ComputeRackUnitAllOf) SetBiosBootmode(v BiosBootModeRelationship) {
	o.BiosBootmode = &v
}

// GetBiosunits returns the Biosunits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetBiosunits() []BiosUnitRelationship {
	if o == nil {
		var ret []BiosUnitRelationship
		return ret
	}
	return o.Biosunits
}

// GetBiosunitsOk returns a tuple with the Biosunits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetBiosunitsOk() (*[]BiosUnitRelationship, bool) {
	if o == nil || o.Biosunits == nil {
		return nil, false
	}
	return &o.Biosunits, true
}

// HasBiosunits returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasBiosunits() bool {
	if o != nil && o.Biosunits != nil {
		return true
	}

	return false
}

// SetBiosunits gets a reference to the given []BiosUnitRelationship and assigns it to the Biosunits field.
func (o *ComputeRackUnitAllOf) SetBiosunits(v []BiosUnitRelationship) {
	o.Biosunits = v
}

// GetBmc returns the Bmc field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetBmc() ManagementControllerRelationship {
	if o == nil || o.Bmc == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Bmc
}

// GetBmcOk returns a tuple with the Bmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetBmcOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Bmc == nil {
		return nil, false
	}
	return o.Bmc, true
}

// HasBmc returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasBmc() bool {
	if o != nil && o.Bmc != nil {
		return true
	}

	return false
}

// SetBmc gets a reference to the given ManagementControllerRelationship and assigns it to the Bmc field.
func (o *ComputeRackUnitAllOf) SetBmc(v ManagementControllerRelationship) {
	o.Bmc = &v
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetBoard() ComputeBoardRelationship {
	if o == nil || o.Board == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.Board == nil {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasBoard() bool {
	if o != nil && o.Board != nil {
		return true
	}

	return false
}

// SetBoard gets a reference to the given ComputeBoardRelationship and assigns it to the Board field.
func (o *ComputeRackUnitAllOf) SetBoard(v ComputeBoardRelationship) {
	o.Board = &v
}

// GetBootDeviceBootmode returns the BootDeviceBootmode field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetBootDeviceBootmode() BootDeviceBootModeRelationship {
	if o == nil || o.BootDeviceBootmode == nil {
		var ret BootDeviceBootModeRelationship
		return ret
	}
	return *o.BootDeviceBootmode
}

// GetBootDeviceBootmodeOk returns a tuple with the BootDeviceBootmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetBootDeviceBootmodeOk() (*BootDeviceBootModeRelationship, bool) {
	if o == nil || o.BootDeviceBootmode == nil {
		return nil, false
	}
	return o.BootDeviceBootmode, true
}

// HasBootDeviceBootmode returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasBootDeviceBootmode() bool {
	if o != nil && o.BootDeviceBootmode != nil {
		return true
	}

	return false
}

// SetBootDeviceBootmode gets a reference to the given BootDeviceBootModeRelationship and assigns it to the BootDeviceBootmode field.
func (o *ComputeRackUnitAllOf) SetBootDeviceBootmode(v BootDeviceBootModeRelationship) {
	o.BootDeviceBootmode = &v
}

// GetFanmodules returns the Fanmodules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetFanmodules() []EquipmentFanModuleRelationship {
	if o == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return o.Fanmodules
}

// GetFanmodulesOk returns a tuple with the Fanmodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool) {
	if o == nil || o.Fanmodules == nil {
		return nil, false
	}
	return &o.Fanmodules, true
}

// HasFanmodules returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasFanmodules() bool {
	if o != nil && o.Fanmodules != nil {
		return true
	}

	return false
}

// SetFanmodules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the Fanmodules field.
func (o *ComputeRackUnitAllOf) SetFanmodules(v []EquipmentFanModuleRelationship) {
	o.Fanmodules = v
}

// GetGenericInventoryHolders returns the GenericInventoryHolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship {
	if o == nil {
		var ret []InventoryGenericInventoryHolderRelationship
		return ret
	}
	return o.GenericInventoryHolders
}

// GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool) {
	if o == nil || o.GenericInventoryHolders == nil {
		return nil, false
	}
	return &o.GenericInventoryHolders, true
}

// HasGenericInventoryHolders returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasGenericInventoryHolders() bool {
	if o != nil && o.GenericInventoryHolders != nil {
		return true
	}

	return false
}

// SetGenericInventoryHolders gets a reference to the given []InventoryGenericInventoryHolderRelationship and assigns it to the GenericInventoryHolders field.
func (o *ComputeRackUnitAllOf) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship) {
	o.GenericInventoryHolders = v
}

// GetGraphicsCards returns the GraphicsCards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetGraphicsCards() []GraphicsCardRelationship {
	if o == nil {
		var ret []GraphicsCardRelationship
		return ret
	}
	return o.GraphicsCards
}

// GetGraphicsCardsOk returns a tuple with the GraphicsCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool) {
	if o == nil || o.GraphicsCards == nil {
		return nil, false
	}
	return &o.GraphicsCards, true
}

// HasGraphicsCards returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasGraphicsCards() bool {
	if o != nil && o.GraphicsCards != nil {
		return true
	}

	return false
}

// SetGraphicsCards gets a reference to the given []GraphicsCardRelationship and assigns it to the GraphicsCards field.
func (o *ComputeRackUnitAllOf) SetGraphicsCards(v []GraphicsCardRelationship) {
	o.GraphicsCards = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeRackUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *ComputeRackUnitAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetMemoryArrays returns the MemoryArrays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetMemoryArrays() []MemoryArrayRelationship {
	if o == nil {
		var ret []MemoryArrayRelationship
		return ret
	}
	return o.MemoryArrays
}

// GetMemoryArraysOk returns a tuple with the MemoryArrays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool) {
	if o == nil || o.MemoryArrays == nil {
		return nil, false
	}
	return &o.MemoryArrays, true
}

// HasMemoryArrays returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasMemoryArrays() bool {
	if o != nil && o.MemoryArrays != nil {
		return true
	}

	return false
}

// SetMemoryArrays gets a reference to the given []MemoryArrayRelationship and assigns it to the MemoryArrays field.
func (o *ComputeRackUnitAllOf) SetMemoryArrays(v []MemoryArrayRelationship) {
	o.MemoryArrays = v
}

// GetPciDevices returns the PciDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetPciDevices() []PciDeviceRelationship {
	if o == nil {
		var ret []PciDeviceRelationship
		return ret
	}
	return o.PciDevices
}

// GetPciDevicesOk returns a tuple with the PciDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetPciDevicesOk() (*[]PciDeviceRelationship, bool) {
	if o == nil || o.PciDevices == nil {
		return nil, false
	}
	return &o.PciDevices, true
}

// HasPciDevices returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasPciDevices() bool {
	if o != nil && o.PciDevices != nil {
		return true
	}

	return false
}

// SetPciDevices gets a reference to the given []PciDeviceRelationship and assigns it to the PciDevices field.
func (o *ComputeRackUnitAllOf) SetPciDevices(v []PciDeviceRelationship) {
	o.PciDevices = v
}

// GetProcessors returns the Processors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetProcessors() []ProcessorUnitRelationship {
	if o == nil {
		var ret []ProcessorUnitRelationship
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return &o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasProcessors() bool {
	if o != nil && o.Processors != nil {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []ProcessorUnitRelationship and assigns it to the Processors field.
func (o *ComputeRackUnitAllOf) SetProcessors(v []ProcessorUnitRelationship) {
	o.Processors = v
}

// GetPsus returns the Psus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetPsus() []EquipmentPsuRelationship {
	if o == nil {
		var ret []EquipmentPsuRelationship
		return ret
	}
	return o.Psus
}

// GetPsusOk returns a tuple with the Psus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool) {
	if o == nil || o.Psus == nil {
		return nil, false
	}
	return &o.Psus, true
}

// HasPsus returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasPsus() bool {
	if o != nil && o.Psus != nil {
		return true
	}

	return false
}

// SetPsus gets a reference to the given []EquipmentPsuRelationship and assigns it to the Psus field.
func (o *ComputeRackUnitAllOf) SetPsus(v []EquipmentPsuRelationship) {
	o.Psus = v
}

// GetRackEnclosureSlot returns the RackEnclosureSlot field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetRackEnclosureSlot() EquipmentRackEnclosureSlotRelationship {
	if o == nil || o.RackEnclosureSlot == nil {
		var ret EquipmentRackEnclosureSlotRelationship
		return ret
	}
	return *o.RackEnclosureSlot
}

// GetRackEnclosureSlotOk returns a tuple with the RackEnclosureSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetRackEnclosureSlotOk() (*EquipmentRackEnclosureSlotRelationship, bool) {
	if o == nil || o.RackEnclosureSlot == nil {
		return nil, false
	}
	return o.RackEnclosureSlot, true
}

// HasRackEnclosureSlot returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasRackEnclosureSlot() bool {
	if o != nil && o.RackEnclosureSlot != nil {
		return true
	}

	return false
}

// SetRackEnclosureSlot gets a reference to the given EquipmentRackEnclosureSlotRelationship and assigns it to the RackEnclosureSlot field.
func (o *ComputeRackUnitAllOf) SetRackEnclosureSlot(v EquipmentRackEnclosureSlotRelationship) {
	o.RackEnclosureSlot = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeRackUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSasExpanders returns the SasExpanders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetSasExpanders() []StorageSasExpanderRelationship {
	if o == nil {
		var ret []StorageSasExpanderRelationship
		return ret
	}
	return o.SasExpanders
}

// GetSasExpandersOk returns a tuple with the SasExpanders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetSasExpandersOk() (*[]StorageSasExpanderRelationship, bool) {
	if o == nil || o.SasExpanders == nil {
		return nil, false
	}
	return &o.SasExpanders, true
}

// HasSasExpanders returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasSasExpanders() bool {
	if o != nil && o.SasExpanders != nil {
		return true
	}

	return false
}

// SetSasExpanders gets a reference to the given []StorageSasExpanderRelationship and assigns it to the SasExpanders field.
func (o *ComputeRackUnitAllOf) SetSasExpanders(v []StorageSasExpanderRelationship) {
	o.SasExpanders = v
}

// GetStorageControllers returns the StorageControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetStorageControllers() []StorageControllerRelationship {
	if o == nil {
		var ret []StorageControllerRelationship
		return ret
	}
	return o.StorageControllers
}

// GetStorageControllersOk returns a tuple with the StorageControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetStorageControllersOk() (*[]StorageControllerRelationship, bool) {
	if o == nil || o.StorageControllers == nil {
		return nil, false
	}
	return &o.StorageControllers, true
}

// HasStorageControllers returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasStorageControllers() bool {
	if o != nil && o.StorageControllers != nil {
		return true
	}

	return false
}

// SetStorageControllers gets a reference to the given []StorageControllerRelationship and assigns it to the StorageControllers field.
func (o *ComputeRackUnitAllOf) SetStorageControllers(v []StorageControllerRelationship) {
	o.StorageControllers = v
}

// GetStorageEnclosures returns the StorageEnclosures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeRackUnitAllOf) GetStorageEnclosures() []StorageEnclosureRelationship {
	if o == nil {
		var ret []StorageEnclosureRelationship
		return ret
	}
	return o.StorageEnclosures
}

// GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeRackUnitAllOf) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosures == nil {
		return nil, false
	}
	return &o.StorageEnclosures, true
}

// HasStorageEnclosures returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasStorageEnclosures() bool {
	if o != nil && o.StorageEnclosures != nil {
		return true
	}

	return false
}

// SetStorageEnclosures gets a reference to the given []StorageEnclosureRelationship and assigns it to the StorageEnclosures field.
func (o *ComputeRackUnitAllOf) SetStorageEnclosures(v []StorageEnclosureRelationship) {
	o.StorageEnclosures = v
}

// GetTopSystem returns the TopSystem field value if set, zero value otherwise.
func (o *ComputeRackUnitAllOf) GetTopSystem() TopSystemRelationship {
	if o == nil || o.TopSystem == nil {
		var ret TopSystemRelationship
		return ret
	}
	return *o.TopSystem
}

// GetTopSystemOk returns a tuple with the TopSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitAllOf) GetTopSystemOk() (*TopSystemRelationship, bool) {
	if o == nil || o.TopSystem == nil {
		return nil, false
	}
	return o.TopSystem, true
}

// HasTopSystem returns a boolean if a field has been set.
func (o *ComputeRackUnitAllOf) HasTopSystem() bool {
	if o != nil && o.TopSystem != nil {
		return true
	}

	return false
}

// SetTopSystem gets a reference to the given TopSystemRelationship and assigns it to the TopSystem field.
func (o *ComputeRackUnitAllOf) SetTopSystem(v TopSystemRelationship) {
	o.TopSystem = &v
}

func (o ComputeRackUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.ServerId != nil {
		toSerialize["ServerId"] = o.ServerId
	}
	if o.TopologyScanStatus != nil {
		toSerialize["TopologyScanStatus"] = o.TopologyScanStatus
	}
	if o.Adapters != nil {
		toSerialize["Adapters"] = o.Adapters
	}
	if o.BiosBootmode != nil {
		toSerialize["BiosBootmode"] = o.BiosBootmode
	}
	if o.Biosunits != nil {
		toSerialize["Biosunits"] = o.Biosunits
	}
	if o.Bmc != nil {
		toSerialize["Bmc"] = o.Bmc
	}
	if o.Board != nil {
		toSerialize["Board"] = o.Board
	}
	if o.BootDeviceBootmode != nil {
		toSerialize["BootDeviceBootmode"] = o.BootDeviceBootmode
	}
	if o.Fanmodules != nil {
		toSerialize["Fanmodules"] = o.Fanmodules
	}
	if o.GenericInventoryHolders != nil {
		toSerialize["GenericInventoryHolders"] = o.GenericInventoryHolders
	}
	if o.GraphicsCards != nil {
		toSerialize["GraphicsCards"] = o.GraphicsCards
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.MemoryArrays != nil {
		toSerialize["MemoryArrays"] = o.MemoryArrays
	}
	if o.PciDevices != nil {
		toSerialize["PciDevices"] = o.PciDevices
	}
	if o.Processors != nil {
		toSerialize["Processors"] = o.Processors
	}
	if o.Psus != nil {
		toSerialize["Psus"] = o.Psus
	}
	if o.RackEnclosureSlot != nil {
		toSerialize["RackEnclosureSlot"] = o.RackEnclosureSlot
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SasExpanders != nil {
		toSerialize["SasExpanders"] = o.SasExpanders
	}
	if o.StorageControllers != nil {
		toSerialize["StorageControllers"] = o.StorageControllers
	}
	if o.StorageEnclosures != nil {
		toSerialize["StorageEnclosures"] = o.StorageEnclosures
	}
	if o.TopSystem != nil {
		toSerialize["TopSystem"] = o.TopSystem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeRackUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeRackUnitAllOf := _ComputeRackUnitAllOf{}

	if err = json.Unmarshal(bytes, &varComputeRackUnitAllOf); err == nil {
		*o = ComputeRackUnitAllOf(varComputeRackUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "ServerId")
		delete(additionalProperties, "TopologyScanStatus")
		delete(additionalProperties, "Adapters")
		delete(additionalProperties, "BiosBootmode")
		delete(additionalProperties, "Biosunits")
		delete(additionalProperties, "Bmc")
		delete(additionalProperties, "Board")
		delete(additionalProperties, "BootDeviceBootmode")
		delete(additionalProperties, "Fanmodules")
		delete(additionalProperties, "GenericInventoryHolders")
		delete(additionalProperties, "GraphicsCards")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "MemoryArrays")
		delete(additionalProperties, "PciDevices")
		delete(additionalProperties, "Processors")
		delete(additionalProperties, "Psus")
		delete(additionalProperties, "RackEnclosureSlot")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SasExpanders")
		delete(additionalProperties, "StorageControllers")
		delete(additionalProperties, "StorageEnclosures")
		delete(additionalProperties, "TopSystem")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeRackUnitAllOf struct {
	value *ComputeRackUnitAllOf
	isSet bool
}

func (v NullableComputeRackUnitAllOf) Get() *ComputeRackUnitAllOf {
	return v.value
}

func (v *NullableComputeRackUnitAllOf) Set(val *ComputeRackUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeRackUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeRackUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeRackUnitAllOf(val *ComputeRackUnitAllOf) *NullableComputeRackUnitAllOf {
	return &NullableComputeRackUnitAllOf{value: val, isSet: true}
}

func (v NullableComputeRackUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeRackUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
