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
	"reflect"
	"strings"
)

// WorkflowSolutionInstance Solution instance is one instance of a solution based on a solution definition.
type WorkflowSolutionInstance struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description for this solution instance.
	Description *string `json:"Description,omitempty"`
	// Last status of the solution instance which will be reverted when an ongoing solution action instance is aborted. * `NotCreated` - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state. * `Okay` - The last action on the solution completed and the solution is in Okay state. * `Decommissioned` - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned.
	LastStatus *string `json:"LastStatus,omitempty"`
	// A name of the solution instance. Name of the solution instance must be unique within a type of Solution definition.
	Name *string `json:"Name,omitempty"`
	// Status of the solution instance which controls the actions that can be performed on this instance. * `NotCreated` - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state. * `Okay` - The last action on the solution completed and the solution is in Okay state. * `Decommissioned` - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned.
	Status *string `json:"Status,omitempty"`
	// The user identifier which indicates the user that started this workflow.
	UserId               *string                                 `json:"UserId,omitempty"`
	Organization         *OrganizationOrganizationRelationship   `json:"Organization,omitempty"`
	SolutionDefinition   *WorkflowSolutionDefinitionRelationship `json:"SolutionDefinition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSolutionInstance WorkflowSolutionInstance

// NewWorkflowSolutionInstance instantiates a new WorkflowSolutionInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSolutionInstance(classId string, objectType string) *WorkflowSolutionInstance {
	this := WorkflowSolutionInstance{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowSolutionInstanceWithDefaults instantiates a new WorkflowSolutionInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSolutionInstanceWithDefaults() *WorkflowSolutionInstance {
	this := WorkflowSolutionInstance{}
	var classId string = "workflow.SolutionInstance"
	this.ClassId = classId
	var objectType string = "workflow.SolutionInstance"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSolutionInstance) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSolutionInstance) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSolutionInstance) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSolutionInstance) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowSolutionInstance) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowSolutionInstance) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowSolutionInstance) SetDescription(v string) {
	o.Description = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *WorkflowSolutionInstance) GetLastStatus() string {
	if o == nil || o.LastStatus == nil {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetLastStatusOk() (*string, bool) {
	if o == nil || o.LastStatus == nil {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *WorkflowSolutionInstance) HasLastStatus() bool {
	if o != nil && o.LastStatus != nil {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *WorkflowSolutionInstance) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowSolutionInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowSolutionInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowSolutionInstance) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowSolutionInstance) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowSolutionInstance) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowSolutionInstance) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *WorkflowSolutionInstance) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *WorkflowSolutionInstance) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *WorkflowSolutionInstance) SetUserId(v string) {
	o.UserId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *WorkflowSolutionInstance) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WorkflowSolutionInstance) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *WorkflowSolutionInstance) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSolutionDefinition returns the SolutionDefinition field value if set, zero value otherwise.
func (o *WorkflowSolutionInstance) GetSolutionDefinition() WorkflowSolutionDefinitionRelationship {
	if o == nil || o.SolutionDefinition == nil {
		var ret WorkflowSolutionDefinitionRelationship
		return ret
	}
	return *o.SolutionDefinition
}

// GetSolutionDefinitionOk returns a tuple with the SolutionDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionInstance) GetSolutionDefinitionOk() (*WorkflowSolutionDefinitionRelationship, bool) {
	if o == nil || o.SolutionDefinition == nil {
		return nil, false
	}
	return o.SolutionDefinition, true
}

// HasSolutionDefinition returns a boolean if a field has been set.
func (o *WorkflowSolutionInstance) HasSolutionDefinition() bool {
	if o != nil && o.SolutionDefinition != nil {
		return true
	}

	return false
}

// SetSolutionDefinition gets a reference to the given WorkflowSolutionDefinitionRelationship and assigns it to the SolutionDefinition field.
func (o *WorkflowSolutionInstance) SetSolutionDefinition(v WorkflowSolutionDefinitionRelationship) {
	o.SolutionDefinition = &v
}

func (o WorkflowSolutionInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.LastStatus != nil {
		toSerialize["LastStatus"] = o.LastStatus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.SolutionDefinition != nil {
		toSerialize["SolutionDefinition"] = o.SolutionDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSolutionInstance) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowSolutionInstanceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The description for this solution instance.
		Description *string `json:"Description,omitempty"`
		// Last status of the solution instance which will be reverted when an ongoing solution action instance is aborted. * `NotCreated` - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state. * `Okay` - The last action on the solution completed and the solution is in Okay state. * `Decommissioned` - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned.
		LastStatus *string `json:"LastStatus,omitempty"`
		// A name of the solution instance. Name of the solution instance must be unique within a type of Solution definition.
		Name *string `json:"Name,omitempty"`
		// Status of the solution instance which controls the actions that can be performed on this instance. * `NotCreated` - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state. * `Okay` - The last action on the solution completed and the solution is in Okay state. * `Decommissioned` - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned.
		Status *string `json:"Status,omitempty"`
		// The user identifier which indicates the user that started this workflow.
		UserId             *string                                 `json:"UserId,omitempty"`
		Organization       *OrganizationOrganizationRelationship   `json:"Organization,omitempty"`
		SolutionDefinition *WorkflowSolutionDefinitionRelationship `json:"SolutionDefinition,omitempty"`
	}

	varWorkflowSolutionInstanceWithoutEmbeddedStruct := WorkflowSolutionInstanceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowSolutionInstanceWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowSolutionInstance := _WorkflowSolutionInstance{}
		varWorkflowSolutionInstance.ClassId = varWorkflowSolutionInstanceWithoutEmbeddedStruct.ClassId
		varWorkflowSolutionInstance.ObjectType = varWorkflowSolutionInstanceWithoutEmbeddedStruct.ObjectType
		varWorkflowSolutionInstance.Description = varWorkflowSolutionInstanceWithoutEmbeddedStruct.Description
		varWorkflowSolutionInstance.LastStatus = varWorkflowSolutionInstanceWithoutEmbeddedStruct.LastStatus
		varWorkflowSolutionInstance.Name = varWorkflowSolutionInstanceWithoutEmbeddedStruct.Name
		varWorkflowSolutionInstance.Status = varWorkflowSolutionInstanceWithoutEmbeddedStruct.Status
		varWorkflowSolutionInstance.UserId = varWorkflowSolutionInstanceWithoutEmbeddedStruct.UserId
		varWorkflowSolutionInstance.Organization = varWorkflowSolutionInstanceWithoutEmbeddedStruct.Organization
		varWorkflowSolutionInstance.SolutionDefinition = varWorkflowSolutionInstanceWithoutEmbeddedStruct.SolutionDefinition
		*o = WorkflowSolutionInstance(varWorkflowSolutionInstance)
	} else {
		return err
	}

	varWorkflowSolutionInstance := _WorkflowSolutionInstance{}

	err = json.Unmarshal(bytes, &varWorkflowSolutionInstance)
	if err == nil {
		o.MoBaseMo = varWorkflowSolutionInstance.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "LastStatus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "SolutionDefinition")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableWorkflowSolutionInstance struct {
	value *WorkflowSolutionInstance
	isSet bool
}

func (v NullableWorkflowSolutionInstance) Get() *WorkflowSolutionInstance {
	return v.value
}

func (v *NullableWorkflowSolutionInstance) Set(val *WorkflowSolutionInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSolutionInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSolutionInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSolutionInstance(val *WorkflowSolutionInstance) *NullableWorkflowSolutionInstance {
	return &NullableWorkflowSolutionInstance{value: val, isSet: true}
}

func (v NullableWorkflowSolutionInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSolutionInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
