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

// WorkflowDynamicWorkflowActionTaskListAllOf Definition of the list of properties defined in 'workflow.DynamicWorkflowActionTaskList', excluding properties defined in parent classes.
type WorkflowDynamicWorkflowActionTaskListAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The action of the Dynamic Workflow.
	Action *string `json:"Action,omitempty"`
	// The task list that has precedence which dictates how the workflow should be constructed.
	Tasks                interface{} `json:"Tasks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDynamicWorkflowActionTaskListAllOf WorkflowDynamicWorkflowActionTaskListAllOf

// NewWorkflowDynamicWorkflowActionTaskListAllOf instantiates a new WorkflowDynamicWorkflowActionTaskListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDynamicWorkflowActionTaskListAllOf(classId string, objectType string) *WorkflowDynamicWorkflowActionTaskListAllOf {
	this := WorkflowDynamicWorkflowActionTaskListAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowDynamicWorkflowActionTaskListAllOfWithDefaults instantiates a new WorkflowDynamicWorkflowActionTaskListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDynamicWorkflowActionTaskListAllOfWithDefaults() *WorkflowDynamicWorkflowActionTaskListAllOf {
	this := WorkflowDynamicWorkflowActionTaskListAllOf{}
	var classId string = "workflow.DynamicWorkflowActionTaskList"
	this.ClassId = classId
	var objectType string = "workflow.DynamicWorkflowActionTaskList"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) SetAction(v string) {
	o.Action = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetTasks() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) GetTasksOk() (*interface{}, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given interface{} and assigns it to the Tasks field.
func (o *WorkflowDynamicWorkflowActionTaskListAllOf) SetTasks(v interface{}) {
	o.Tasks = v
}

func (o WorkflowDynamicWorkflowActionTaskListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Tasks != nil {
		toSerialize["Tasks"] = o.Tasks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDynamicWorkflowActionTaskListAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowDynamicWorkflowActionTaskListAllOf := _WorkflowDynamicWorkflowActionTaskListAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowDynamicWorkflowActionTaskListAllOf); err == nil {
		*o = WorkflowDynamicWorkflowActionTaskListAllOf(varWorkflowDynamicWorkflowActionTaskListAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "Tasks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowDynamicWorkflowActionTaskListAllOf struct {
	value *WorkflowDynamicWorkflowActionTaskListAllOf
	isSet bool
}

func (v NullableWorkflowDynamicWorkflowActionTaskListAllOf) Get() *WorkflowDynamicWorkflowActionTaskListAllOf {
	return v.value
}

func (v *NullableWorkflowDynamicWorkflowActionTaskListAllOf) Set(val *WorkflowDynamicWorkflowActionTaskListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDynamicWorkflowActionTaskListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDynamicWorkflowActionTaskListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDynamicWorkflowActionTaskListAllOf(val *WorkflowDynamicWorkflowActionTaskListAllOf) *NullableWorkflowDynamicWorkflowActionTaskListAllOf {
	return &NullableWorkflowDynamicWorkflowActionTaskListAllOf{value: val, isSet: true}
}

func (v NullableWorkflowDynamicWorkflowActionTaskListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDynamicWorkflowActionTaskListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
