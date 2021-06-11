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

// ChassisProfileAllOf Definition of the list of properties defined in 'chassis.Profile', excluding properties defined in parent classes.
type ChassisProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                     `json:"ObjectType"`
	ConfigChanges NullablePolicyConfigChange `json:"ConfigChanges,omitempty"`
	// The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `FIAttached` - Chassis which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform    *string                       `json:"TargetPlatform,omitempty"`
	AssignedChassis   *EquipmentChassisRelationship `json:"AssignedChassis,omitempty"`
	AssociatedChassis *EquipmentChassisRelationship `json:"AssociatedChassis,omitempty"`
	// An array of relationships to chassisConfigChangeDetail resources.
	ConfigChangeDetails []ChassisConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
	ConfigResult        *ChassisConfigResultRelationship        `json:"ConfigResult,omitempty"`
	// An array of relationships to chassisIomProfile resources.
	IomProfiles  []ChassisIomProfileRelationship       `json:"IomProfiles,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisProfileAllOf ChassisProfileAllOf

// NewChassisProfileAllOf instantiates a new ChassisProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisProfileAllOf(classId string, objectType string) *ChassisProfileAllOf {
	this := ChassisProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetPlatform string = "FIAttached"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewChassisProfileAllOfWithDefaults instantiates a new ChassisProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisProfileAllOfWithDefaults() *ChassisProfileAllOf {
	this := ChassisProfileAllOf{}
	var classId string = "chassis.Profile"
	this.ClassId = classId
	var objectType string = "chassis.Profile"
	this.ObjectType = objectType
	var targetPlatform string = "FIAttached"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfileAllOf) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ChassisProfileAllOf) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ChassisProfileAllOf) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ChassisProfileAllOf) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *ChassisProfileAllOf) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfileAllOf) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *ChassisProfileAllOf) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetAssignedChassis returns the AssignedChassis field value if set, zero value otherwise.
func (o *ChassisProfileAllOf) GetAssignedChassis() EquipmentChassisRelationship {
	if o == nil || o.AssignedChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.AssignedChassis
}

// GetAssignedChassisOk returns a tuple with the AssignedChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfileAllOf) GetAssignedChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.AssignedChassis == nil {
		return nil, false
	}
	return o.AssignedChassis, true
}

// HasAssignedChassis returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasAssignedChassis() bool {
	if o != nil && o.AssignedChassis != nil {
		return true
	}

	return false
}

// SetAssignedChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the AssignedChassis field.
func (o *ChassisProfileAllOf) SetAssignedChassis(v EquipmentChassisRelationship) {
	o.AssignedChassis = &v
}

// GetAssociatedChassis returns the AssociatedChassis field value if set, zero value otherwise.
func (o *ChassisProfileAllOf) GetAssociatedChassis() EquipmentChassisRelationship {
	if o == nil || o.AssociatedChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.AssociatedChassis
}

// GetAssociatedChassisOk returns a tuple with the AssociatedChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfileAllOf) GetAssociatedChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.AssociatedChassis == nil {
		return nil, false
	}
	return o.AssociatedChassis, true
}

// HasAssociatedChassis returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasAssociatedChassis() bool {
	if o != nil && o.AssociatedChassis != nil {
		return true
	}

	return false
}

// SetAssociatedChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the AssociatedChassis field.
func (o *ChassisProfileAllOf) SetAssociatedChassis(v EquipmentChassisRelationship) {
	o.AssociatedChassis = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfileAllOf) GetConfigChangeDetails() []ChassisConfigChangeDetailRelationship {
	if o == nil {
		var ret []ChassisConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfileAllOf) GetConfigChangeDetailsOk() (*[]ChassisConfigChangeDetailRelationship, bool) {
	if o == nil || o.ConfigChangeDetails == nil {
		return nil, false
	}
	return &o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasConfigChangeDetails() bool {
	if o != nil && o.ConfigChangeDetails != nil {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []ChassisConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *ChassisProfileAllOf) SetConfigChangeDetails(v []ChassisConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *ChassisProfileAllOf) GetConfigResult() ChassisConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret ChassisConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfileAllOf) GetConfigResultOk() (*ChassisConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given ChassisConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ChassisProfileAllOf) SetConfigResult(v ChassisConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetIomProfiles returns the IomProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfileAllOf) GetIomProfiles() []ChassisIomProfileRelationship {
	if o == nil {
		var ret []ChassisIomProfileRelationship
		return ret
	}
	return o.IomProfiles
}

// GetIomProfilesOk returns a tuple with the IomProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfileAllOf) GetIomProfilesOk() (*[]ChassisIomProfileRelationship, bool) {
	if o == nil || o.IomProfiles == nil {
		return nil, false
	}
	return &o.IomProfiles, true
}

// HasIomProfiles returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasIomProfiles() bool {
	if o != nil && o.IomProfiles != nil {
		return true
	}

	return false
}

// SetIomProfiles gets a reference to the given []ChassisIomProfileRelationship and assigns it to the IomProfiles field.
func (o *ChassisProfileAllOf) SetIomProfiles(v []ChassisIomProfileRelationship) {
	o.IomProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ChassisProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ChassisProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfileAllOf) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return &o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *ChassisProfileAllOf) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *ChassisProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

func (o ChassisProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.AssignedChassis != nil {
		toSerialize["AssignedChassis"] = o.AssignedChassis
	}
	if o.AssociatedChassis != nil {
		toSerialize["AssociatedChassis"] = o.AssociatedChassis
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.IomProfiles != nil {
		toSerialize["IomProfiles"] = o.IomProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varChassisProfileAllOf := _ChassisProfileAllOf{}

	if err = json.Unmarshal(bytes, &varChassisProfileAllOf); err == nil {
		*o = ChassisProfileAllOf(varChassisProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "AssignedChassis")
		delete(additionalProperties, "AssociatedChassis")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "IomProfiles")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RunningWorkflows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChassisProfileAllOf struct {
	value *ChassisProfileAllOf
	isSet bool
}

func (v NullableChassisProfileAllOf) Get() *ChassisProfileAllOf {
	return v.value
}

func (v *NullableChassisProfileAllOf) Set(val *ChassisProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisProfileAllOf(val *ChassisProfileAllOf) *NullableChassisProfileAllOf {
	return &NullableChassisProfileAllOf{value: val, isSet: true}
}

func (v NullableChassisProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
