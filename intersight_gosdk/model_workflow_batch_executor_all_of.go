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

// WorkflowBatchExecutorAllOf Definition of the list of properties defined in 'workflow.BatchExecutor', excluding properties defined in parent classes.
type WorkflowBatchExecutorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType  string                          `json:"ObjectType"`
	Batch       []WorkflowApi                   `json:"Batch,omitempty"`
	Constraints NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
	// A detailed description about the batch APIs.
	Description *string `json:"Description,omitempty"`
	// Name for the batch API task.
	Name *string `json:"Name,omitempty"`
	// All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
	Outcomes interface{} `json:"Outcomes,omitempty"`
	// Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property.
	Output interface{} `json:"Output,omitempty"`
	// When an execution of a nth API in the Batch fails, Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false.
	RetryFromFailedApi *bool `json:"RetryFromFailedApi,omitempty"`
	// The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
	SkipOnCondition *string `json:"SkipOnCondition,omitempty"`
	// This will hold the data needed for task to be rendered in the user interface.
	UiRenderingData      interface{} `json:"UiRenderingData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBatchExecutorAllOf WorkflowBatchExecutorAllOf

// NewWorkflowBatchExecutorAllOf instantiates a new WorkflowBatchExecutorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBatchExecutorAllOf(classId string, objectType string) *WorkflowBatchExecutorAllOf {
	this := WorkflowBatchExecutorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBatchExecutorAllOfWithDefaults instantiates a new WorkflowBatchExecutorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBatchExecutorAllOfWithDefaults() *WorkflowBatchExecutorAllOf {
	this := WorkflowBatchExecutorAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBatchExecutorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBatchExecutorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBatchExecutorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBatchExecutorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBatch returns the Batch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutorAllOf) GetBatch() []WorkflowApi {
	if o == nil {
		var ret []WorkflowApi
		return ret
	}
	return o.Batch
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutorAllOf) GetBatchOk() (*[]WorkflowApi, bool) {
	if o == nil || o.Batch == nil {
		return nil, false
	}
	return &o.Batch, true
}

// HasBatch returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasBatch() bool {
	if o != nil && o.Batch != nil {
		return true
	}

	return false
}

// SetBatch gets a reference to the given []WorkflowApi and assigns it to the Batch field.
func (o *WorkflowBatchExecutorAllOf) SetBatch(v []WorkflowApi) {
	o.Batch = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutorAllOf) GetConstraints() WorkflowTaskConstraints {
	if o == nil || o.Constraints.Get() == nil {
		var ret WorkflowTaskConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutorAllOf) GetConstraintsOk() (*WorkflowTaskConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableWorkflowTaskConstraints and assigns it to the Constraints field.
func (o *WorkflowBatchExecutorAllOf) SetConstraints(v WorkflowTaskConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *WorkflowBatchExecutorAllOf) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *WorkflowBatchExecutorAllOf) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowBatchExecutorAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutorAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowBatchExecutorAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowBatchExecutorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowBatchExecutorAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutorAllOf) GetOutcomes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutorAllOf) GetOutcomesOk() (*interface{}, bool) {
	if o == nil || o.Outcomes == nil {
		return nil, false
	}
	return &o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given interface{} and assigns it to the Outcomes field.
func (o *WorkflowBatchExecutorAllOf) SetOutcomes(v interface{}) {
	o.Outcomes = v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutorAllOf) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutorAllOf) GetOutputOk() (*interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *WorkflowBatchExecutorAllOf) SetOutput(v interface{}) {
	o.Output = v
}

// GetRetryFromFailedApi returns the RetryFromFailedApi field value if set, zero value otherwise.
func (o *WorkflowBatchExecutorAllOf) GetRetryFromFailedApi() bool {
	if o == nil || o.RetryFromFailedApi == nil {
		var ret bool
		return ret
	}
	return *o.RetryFromFailedApi
}

// GetRetryFromFailedApiOk returns a tuple with the RetryFromFailedApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutorAllOf) GetRetryFromFailedApiOk() (*bool, bool) {
	if o == nil || o.RetryFromFailedApi == nil {
		return nil, false
	}
	return o.RetryFromFailedApi, true
}

// HasRetryFromFailedApi returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasRetryFromFailedApi() bool {
	if o != nil && o.RetryFromFailedApi != nil {
		return true
	}

	return false
}

// SetRetryFromFailedApi gets a reference to the given bool and assigns it to the RetryFromFailedApi field.
func (o *WorkflowBatchExecutorAllOf) SetRetryFromFailedApi(v bool) {
	o.RetryFromFailedApi = &v
}

// GetSkipOnCondition returns the SkipOnCondition field value if set, zero value otherwise.
func (o *WorkflowBatchExecutorAllOf) GetSkipOnCondition() string {
	if o == nil || o.SkipOnCondition == nil {
		var ret string
		return ret
	}
	return *o.SkipOnCondition
}

// GetSkipOnConditionOk returns a tuple with the SkipOnCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchExecutorAllOf) GetSkipOnConditionOk() (*string, bool) {
	if o == nil || o.SkipOnCondition == nil {
		return nil, false
	}
	return o.SkipOnCondition, true
}

// HasSkipOnCondition returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasSkipOnCondition() bool {
	if o != nil && o.SkipOnCondition != nil {
		return true
	}

	return false
}

// SetSkipOnCondition gets a reference to the given string and assigns it to the SkipOnCondition field.
func (o *WorkflowBatchExecutorAllOf) SetSkipOnCondition(v string) {
	o.SkipOnCondition = &v
}

// GetUiRenderingData returns the UiRenderingData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchExecutorAllOf) GetUiRenderingData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiRenderingData
}

// GetUiRenderingDataOk returns a tuple with the UiRenderingData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchExecutorAllOf) GetUiRenderingDataOk() (*interface{}, bool) {
	if o == nil || o.UiRenderingData == nil {
		return nil, false
	}
	return &o.UiRenderingData, true
}

// HasUiRenderingData returns a boolean if a field has been set.
func (o *WorkflowBatchExecutorAllOf) HasUiRenderingData() bool {
	if o != nil && o.UiRenderingData != nil {
		return true
	}

	return false
}

// SetUiRenderingData gets a reference to the given interface{} and assigns it to the UiRenderingData field.
func (o *WorkflowBatchExecutorAllOf) SetUiRenderingData(v interface{}) {
	o.UiRenderingData = v
}

func (o WorkflowBatchExecutorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Batch != nil {
		toSerialize["Batch"] = o.Batch
	}
	if o.Constraints.IsSet() {
		toSerialize["Constraints"] = o.Constraints.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Outcomes != nil {
		toSerialize["Outcomes"] = o.Outcomes
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.RetryFromFailedApi != nil {
		toSerialize["RetryFromFailedApi"] = o.RetryFromFailedApi
	}
	if o.SkipOnCondition != nil {
		toSerialize["SkipOnCondition"] = o.SkipOnCondition
	}
	if o.UiRenderingData != nil {
		toSerialize["UiRenderingData"] = o.UiRenderingData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowBatchExecutorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowBatchExecutorAllOf := _WorkflowBatchExecutorAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowBatchExecutorAllOf); err == nil {
		*o = WorkflowBatchExecutorAllOf(varWorkflowBatchExecutorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Batch")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Outcomes")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "RetryFromFailedApi")
		delete(additionalProperties, "SkipOnCondition")
		delete(additionalProperties, "UiRenderingData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowBatchExecutorAllOf struct {
	value *WorkflowBatchExecutorAllOf
	isSet bool
}

func (v NullableWorkflowBatchExecutorAllOf) Get() *WorkflowBatchExecutorAllOf {
	return v.value
}

func (v *NullableWorkflowBatchExecutorAllOf) Set(val *WorkflowBatchExecutorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBatchExecutorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBatchExecutorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBatchExecutorAllOf(val *WorkflowBatchExecutorAllOf) *NullableWorkflowBatchExecutorAllOf {
	return &NullableWorkflowBatchExecutorAllOf{value: val, isSet: true}
}

func (v NullableWorkflowBatchExecutorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBatchExecutorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
