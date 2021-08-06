/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ConfigExportedItemAllOf Definition of the list of properties defined in 'config.ExportedItem', excluding properties defined in parent classes.
type ConfigExportedItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the file corresponding to item MO.
	FileName *string             `json:"FileName,omitempty"`
	Item     NullableConfigMoRef `json:"Item,omitempty"`
	// MO item identity (the moref corresponding to item) expressed as a string.
	Name *string `json:"Name,omitempty"`
	// Version of the service that owns the item MO.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Status of the item's export operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `RollBackInitiated` - The rollback has been inititiated for import failure. * `RollBackFailed` - The rollback has failed for import failure. * `RollbackCompleted` - The rollback has completed for import failure. * `RollbackAborted` - The rollback has been aborted for import failure. * `OperationTimedOut` - The operation has timed out.
	Status *string `json:"Status,omitempty"`
	// Progress or error message for the MO's export operation.
	StatusMessage *string                         `json:"StatusMessage,omitempty"`
	Exporter      *ConfigExporterRelationship     `json:"Exporter,omitempty"`
	ParentItem    *ConfigExportedItemRelationship `json:"ParentItem,omitempty"`
	// An array of relationships to configExportedItem resources.
	RelatedItems         []ConfigExportedItemRelationship `json:"RelatedItems,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigExportedItemAllOf ConfigExportedItemAllOf

// NewConfigExportedItemAllOf instantiates a new ConfigExportedItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigExportedItemAllOf(classId string, objectType string) *ConfigExportedItemAllOf {
	this := ConfigExportedItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// NewConfigExportedItemAllOfWithDefaults instantiates a new ConfigExportedItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigExportedItemAllOfWithDefaults() *ConfigExportedItemAllOf {
	this := ConfigExportedItemAllOf{}
	var classId string = "config.ExportedItem"
	this.ClassId = classId
	var objectType string = "config.ExportedItem"
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConfigExportedItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConfigExportedItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConfigExportedItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConfigExportedItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ConfigExportedItemAllOf) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ConfigExportedItemAllOf) SetFileName(v string) {
	o.FileName = &v
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigExportedItemAllOf) GetItem() ConfigMoRef {
	if o == nil || o.Item.Get() == nil {
		var ret ConfigMoRef
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigExportedItemAllOf) GetItemOk() (*ConfigMoRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableConfigMoRef and assigns it to the Item field.
func (o *ConfigExportedItemAllOf) SetItem(v ConfigMoRef) {
	o.Item.Set(&v)
}

// SetItemNil sets the value for Item to be an explicit nil
func (o *ConfigExportedItemAllOf) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *ConfigExportedItemAllOf) UnsetItem() {
	o.Item.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigExportedItemAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigExportedItemAllOf) SetName(v string) {
	o.Name = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *ConfigExportedItemAllOf) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *ConfigExportedItemAllOf) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConfigExportedItemAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConfigExportedItemAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ConfigExportedItemAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ConfigExportedItemAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetExporter returns the Exporter field value if set, zero value otherwise.
func (o *ConfigExportedItemAllOf) GetExporter() ConfigExporterRelationship {
	if o == nil || o.Exporter == nil {
		var ret ConfigExporterRelationship
		return ret
	}
	return *o.Exporter
}

// GetExporterOk returns a tuple with the Exporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetExporterOk() (*ConfigExporterRelationship, bool) {
	if o == nil || o.Exporter == nil {
		return nil, false
	}
	return o.Exporter, true
}

// HasExporter returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasExporter() bool {
	if o != nil && o.Exporter != nil {
		return true
	}

	return false
}

// SetExporter gets a reference to the given ConfigExporterRelationship and assigns it to the Exporter field.
func (o *ConfigExportedItemAllOf) SetExporter(v ConfigExporterRelationship) {
	o.Exporter = &v
}

// GetParentItem returns the ParentItem field value if set, zero value otherwise.
func (o *ConfigExportedItemAllOf) GetParentItem() ConfigExportedItemRelationship {
	if o == nil || o.ParentItem == nil {
		var ret ConfigExportedItemRelationship
		return ret
	}
	return *o.ParentItem
}

// GetParentItemOk returns a tuple with the ParentItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExportedItemAllOf) GetParentItemOk() (*ConfigExportedItemRelationship, bool) {
	if o == nil || o.ParentItem == nil {
		return nil, false
	}
	return o.ParentItem, true
}

// HasParentItem returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasParentItem() bool {
	if o != nil && o.ParentItem != nil {
		return true
	}

	return false
}

// SetParentItem gets a reference to the given ConfigExportedItemRelationship and assigns it to the ParentItem field.
func (o *ConfigExportedItemAllOf) SetParentItem(v ConfigExportedItemRelationship) {
	o.ParentItem = &v
}

// GetRelatedItems returns the RelatedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigExportedItemAllOf) GetRelatedItems() []ConfigExportedItemRelationship {
	if o == nil {
		var ret []ConfigExportedItemRelationship
		return ret
	}
	return o.RelatedItems
}

// GetRelatedItemsOk returns a tuple with the RelatedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigExportedItemAllOf) GetRelatedItemsOk() (*[]ConfigExportedItemRelationship, bool) {
	if o == nil || o.RelatedItems == nil {
		return nil, false
	}
	return &o.RelatedItems, true
}

// HasRelatedItems returns a boolean if a field has been set.
func (o *ConfigExportedItemAllOf) HasRelatedItems() bool {
	if o != nil && o.RelatedItems != nil {
		return true
	}

	return false
}

// SetRelatedItems gets a reference to the given []ConfigExportedItemRelationship and assigns it to the RelatedItems field.
func (o *ConfigExportedItemAllOf) SetRelatedItems(v []ConfigExportedItemRelationship) {
	o.RelatedItems = v
}

func (o ConfigExportedItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.Item.IsSet() {
		toSerialize["Item"] = o.Item.Get()
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ServiceVersion != nil {
		toSerialize["ServiceVersion"] = o.ServiceVersion
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.Exporter != nil {
		toSerialize["Exporter"] = o.Exporter
	}
	if o.ParentItem != nil {
		toSerialize["ParentItem"] = o.ParentItem
	}
	if o.RelatedItems != nil {
		toSerialize["RelatedItems"] = o.RelatedItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConfigExportedItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConfigExportedItemAllOf := _ConfigExportedItemAllOf{}

	if err = json.Unmarshal(bytes, &varConfigExportedItemAllOf); err == nil {
		*o = ConfigExportedItemAllOf(varConfigExportedItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "Item")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ServiceVersion")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Exporter")
		delete(additionalProperties, "ParentItem")
		delete(additionalProperties, "RelatedItems")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigExportedItemAllOf struct {
	value *ConfigExportedItemAllOf
	isSet bool
}

func (v NullableConfigExportedItemAllOf) Get() *ConfigExportedItemAllOf {
	return v.value
}

func (v *NullableConfigExportedItemAllOf) Set(val *ConfigExportedItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigExportedItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigExportedItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigExportedItemAllOf(val *ConfigExportedItemAllOf) *NullableConfigExportedItemAllOf {
	return &NullableConfigExportedItemAllOf{value: val, isSet: true}
}

func (v NullableConfigExportedItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigExportedItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
