/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// IaasUcsdInfoAllOf Definition of the list of properties defined in 'iaas.UcsdInfo', excluding properties defined in parent classes.
type IaasUcsdInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Moid of the UCS Director device connector's asset.DeviceRegistration.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Unique ID of UCS Director getting registerd with Intersight.
	Guid *string `json:"Guid,omitempty"`
	// The UCS Director hostname for management.
	HostName *string `json:"HostName,omitempty"`
	// The UCS Director IP address for management.
	Ip *string `json:"Ip,omitempty"`
	// Last successful backup created for this UCS Director appliance if backup is configured.
	LastBackup *time.Time `json:"LastBackup,omitempty"`
	// NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node.
	NodeType *string `json:"NodeType,omitempty"`
	// The UCS Director product name.
	ProductName *string `json:"ProductName,omitempty"`
	// The UCS Director product vendor.
	ProductVendor *string `json:"ProductVendor,omitempty"`
	// The UCS Director product/platform version.
	ProductVersion *string `json:"ProductVersion,omitempty"`
	// The UCS Director status. Possible values are Active, Inactive, Unknown.
	Status *string `json:"Status,omitempty"`
	// An array of relationships to iaasConnectorPack resources.
	ConnectorPack []IaasConnectorPackRelationship `json:"ConnectorPack,omitempty"`
	// An array of relationships to iaasDeviceStatus resources.
	DeviceStatus []IaasDeviceStatusRelationship `json:"DeviceStatus,omitempty"`
	LicenseInfo  *IaasLicenseInfoRelationship   `json:"LicenseInfo,omitempty"`
	// An array of relationships to iaasMostRunTasks resources.
	MostRunTasks         []IaasMostRunTasksRelationship       `json:"MostRunTasks,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	UcsdManagedInfra     *IaasUcsdManagedInfraRelationship    `json:"UcsdManagedInfra,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasUcsdInfoAllOf IaasUcsdInfoAllOf

// NewIaasUcsdInfoAllOf instantiates a new IaasUcsdInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasUcsdInfoAllOf(classId string, objectType string) *IaasUcsdInfoAllOf {
	this := IaasUcsdInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasUcsdInfoAllOfWithDefaults instantiates a new IaasUcsdInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasUcsdInfoAllOfWithDefaults() *IaasUcsdInfoAllOf {
	this := IaasUcsdInfoAllOf{}
	var classId string = "iaas.UcsdInfo"
	this.ClassId = classId
	var objectType string = "iaas.UcsdInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasUcsdInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasUcsdInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasUcsdInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasUcsdInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *IaasUcsdInfoAllOf) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *IaasUcsdInfoAllOf) SetGuid(v string) {
	o.Guid = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *IaasUcsdInfoAllOf) SetHostName(v string) {
	o.HostName = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *IaasUcsdInfoAllOf) SetIp(v string) {
	o.Ip = &v
}

// GetLastBackup returns the LastBackup field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetLastBackup() time.Time {
	if o == nil || o.LastBackup == nil {
		var ret time.Time
		return ret
	}
	return *o.LastBackup
}

// GetLastBackupOk returns a tuple with the LastBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetLastBackupOk() (*time.Time, bool) {
	if o == nil || o.LastBackup == nil {
		return nil, false
	}
	return o.LastBackup, true
}

// HasLastBackup returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasLastBackup() bool {
	if o != nil && o.LastBackup != nil {
		return true
	}

	return false
}

// SetLastBackup gets a reference to the given time.Time and assigns it to the LastBackup field.
func (o *IaasUcsdInfoAllOf) SetLastBackup(v time.Time) {
	o.LastBackup = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasNodeType() bool {
	if o != nil && o.NodeType != nil {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *IaasUcsdInfoAllOf) SetNodeType(v string) {
	o.NodeType = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *IaasUcsdInfoAllOf) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductVendor returns the ProductVendor field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetProductVendor() string {
	if o == nil || o.ProductVendor == nil {
		var ret string
		return ret
	}
	return *o.ProductVendor
}

// GetProductVendorOk returns a tuple with the ProductVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetProductVendorOk() (*string, bool) {
	if o == nil || o.ProductVendor == nil {
		return nil, false
	}
	return o.ProductVendor, true
}

// HasProductVendor returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasProductVendor() bool {
	if o != nil && o.ProductVendor != nil {
		return true
	}

	return false
}

// SetProductVendor gets a reference to the given string and assigns it to the ProductVendor field.
func (o *IaasUcsdInfoAllOf) SetProductVendor(v string) {
	o.ProductVendor = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetProductVersion() string {
	if o == nil || o.ProductVersion == nil {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetProductVersionOk() (*string, bool) {
	if o == nil || o.ProductVersion == nil {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasProductVersion() bool {
	if o != nil && o.ProductVersion != nil {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *IaasUcsdInfoAllOf) SetProductVersion(v string) {
	o.ProductVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IaasUcsdInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetConnectorPack returns the ConnectorPack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfoAllOf) GetConnectorPack() []IaasConnectorPackRelationship {
	if o == nil {
		var ret []IaasConnectorPackRelationship
		return ret
	}
	return o.ConnectorPack
}

// GetConnectorPackOk returns a tuple with the ConnectorPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfoAllOf) GetConnectorPackOk() (*[]IaasConnectorPackRelationship, bool) {
	if o == nil || o.ConnectorPack == nil {
		return nil, false
	}
	return &o.ConnectorPack, true
}

// HasConnectorPack returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasConnectorPack() bool {
	if o != nil && o.ConnectorPack != nil {
		return true
	}

	return false
}

// SetConnectorPack gets a reference to the given []IaasConnectorPackRelationship and assigns it to the ConnectorPack field.
func (o *IaasUcsdInfoAllOf) SetConnectorPack(v []IaasConnectorPackRelationship) {
	o.ConnectorPack = v
}

// GetDeviceStatus returns the DeviceStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfoAllOf) GetDeviceStatus() []IaasDeviceStatusRelationship {
	if o == nil {
		var ret []IaasDeviceStatusRelationship
		return ret
	}
	return o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfoAllOf) GetDeviceStatusOk() (*[]IaasDeviceStatusRelationship, bool) {
	if o == nil || o.DeviceStatus == nil {
		return nil, false
	}
	return &o.DeviceStatus, true
}

// HasDeviceStatus returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasDeviceStatus() bool {
	if o != nil && o.DeviceStatus != nil {
		return true
	}

	return false
}

// SetDeviceStatus gets a reference to the given []IaasDeviceStatusRelationship and assigns it to the DeviceStatus field.
func (o *IaasUcsdInfoAllOf) SetDeviceStatus(v []IaasDeviceStatusRelationship) {
	o.DeviceStatus = v
}

// GetLicenseInfo returns the LicenseInfo field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetLicenseInfo() IaasLicenseInfoRelationship {
	if o == nil || o.LicenseInfo == nil {
		var ret IaasLicenseInfoRelationship
		return ret
	}
	return *o.LicenseInfo
}

// GetLicenseInfoOk returns a tuple with the LicenseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetLicenseInfoOk() (*IaasLicenseInfoRelationship, bool) {
	if o == nil || o.LicenseInfo == nil {
		return nil, false
	}
	return o.LicenseInfo, true
}

// HasLicenseInfo returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasLicenseInfo() bool {
	if o != nil && o.LicenseInfo != nil {
		return true
	}

	return false
}

// SetLicenseInfo gets a reference to the given IaasLicenseInfoRelationship and assigns it to the LicenseInfo field.
func (o *IaasUcsdInfoAllOf) SetLicenseInfo(v IaasLicenseInfoRelationship) {
	o.LicenseInfo = &v
}

// GetMostRunTasks returns the MostRunTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfoAllOf) GetMostRunTasks() []IaasMostRunTasksRelationship {
	if o == nil {
		var ret []IaasMostRunTasksRelationship
		return ret
	}
	return o.MostRunTasks
}

// GetMostRunTasksOk returns a tuple with the MostRunTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfoAllOf) GetMostRunTasksOk() (*[]IaasMostRunTasksRelationship, bool) {
	if o == nil || o.MostRunTasks == nil {
		return nil, false
	}
	return &o.MostRunTasks, true
}

// HasMostRunTasks returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasMostRunTasks() bool {
	if o != nil && o.MostRunTasks != nil {
		return true
	}

	return false
}

// SetMostRunTasks gets a reference to the given []IaasMostRunTasksRelationship and assigns it to the MostRunTasks field.
func (o *IaasUcsdInfoAllOf) SetMostRunTasks(v []IaasMostRunTasksRelationship) {
	o.MostRunTasks = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasUcsdInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetUcsdManagedInfra returns the UcsdManagedInfra field value if set, zero value otherwise.
func (o *IaasUcsdInfoAllOf) GetUcsdManagedInfra() IaasUcsdManagedInfraRelationship {
	if o == nil || o.UcsdManagedInfra == nil {
		var ret IaasUcsdManagedInfraRelationship
		return ret
	}
	return *o.UcsdManagedInfra
}

// GetUcsdManagedInfraOk returns a tuple with the UcsdManagedInfra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfoAllOf) GetUcsdManagedInfraOk() (*IaasUcsdManagedInfraRelationship, bool) {
	if o == nil || o.UcsdManagedInfra == nil {
		return nil, false
	}
	return o.UcsdManagedInfra, true
}

// HasUcsdManagedInfra returns a boolean if a field has been set.
func (o *IaasUcsdInfoAllOf) HasUcsdManagedInfra() bool {
	if o != nil && o.UcsdManagedInfra != nil {
		return true
	}

	return false
}

// SetUcsdManagedInfra gets a reference to the given IaasUcsdManagedInfraRelationship and assigns it to the UcsdManagedInfra field.
func (o *IaasUcsdInfoAllOf) SetUcsdManagedInfra(v IaasUcsdManagedInfraRelationship) {
	o.UcsdManagedInfra = &v
}

func (o IaasUcsdInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	if o.LastBackup != nil {
		toSerialize["LastBackup"] = o.LastBackup
	}
	if o.NodeType != nil {
		toSerialize["NodeType"] = o.NodeType
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.ProductVendor != nil {
		toSerialize["ProductVendor"] = o.ProductVendor
	}
	if o.ProductVersion != nil {
		toSerialize["ProductVersion"] = o.ProductVersion
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.ConnectorPack != nil {
		toSerialize["ConnectorPack"] = o.ConnectorPack
	}
	if o.DeviceStatus != nil {
		toSerialize["DeviceStatus"] = o.DeviceStatus
	}
	if o.LicenseInfo != nil {
		toSerialize["LicenseInfo"] = o.LicenseInfo
	}
	if o.MostRunTasks != nil {
		toSerialize["MostRunTasks"] = o.MostRunTasks
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.UcsdManagedInfra != nil {
		toSerialize["UcsdManagedInfra"] = o.UcsdManagedInfra
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasUcsdInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIaasUcsdInfoAllOf := _IaasUcsdInfoAllOf{}

	if err = json.Unmarshal(bytes, &varIaasUcsdInfoAllOf); err == nil {
		*o = IaasUcsdInfoAllOf(varIaasUcsdInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Guid")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "LastBackup")
		delete(additionalProperties, "NodeType")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "ProductVendor")
		delete(additionalProperties, "ProductVersion")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "ConnectorPack")
		delete(additionalProperties, "DeviceStatus")
		delete(additionalProperties, "LicenseInfo")
		delete(additionalProperties, "MostRunTasks")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "UcsdManagedInfra")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIaasUcsdInfoAllOf struct {
	value *IaasUcsdInfoAllOf
	isSet bool
}

func (v NullableIaasUcsdInfoAllOf) Get() *IaasUcsdInfoAllOf {
	return v.value
}

func (v *NullableIaasUcsdInfoAllOf) Set(val *IaasUcsdInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasUcsdInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasUcsdInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasUcsdInfoAllOf(val *IaasUcsdInfoAllOf) *NullableIaasUcsdInfoAllOf {
	return &NullableIaasUcsdInfoAllOf{value: val, isSet: true}
}

func (v NullableIaasUcsdInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasUcsdInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
