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
	"time"
)

// HyperflexHealthCheckExecutionSnapshotAllOf Definition of the list of properties defined in 'hyperflex.HealthCheckExecutionSnapshot', excluding properties defined in parent classes.
type HyperflexHealthCheckExecutionSnapshotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Category that the HyperFlex health check Definition belongs to.
	Category *string `json:"Category,omitempty"`
	// Information detailing the possible cause of the healthcheck failure, if the check fails.
	Cause *string `json:"Cause,omitempty"`
	// Health check execution completion time.
	CompletionTime *time.Time `json:"CompletionTime,omitempty"`
	// Details of the health check execution result.
	HealthCheckDetails *string `json:"HealthCheckDetails,omitempty"`
	// Error details of a script execution failure.
	HealthCheckExecutionErrorDetails *string `json:"HealthCheckExecutionErrorDetails,omitempty"`
	// Error summary of a script execution failure.
	HealthCheckExecutionErrorSummary *string `json:"HealthCheckExecutionErrorSummary,omitempty"`
	// Status of the health check execution. * `UNKNOWN` - Indicates that the health heck execution results are unknown. * `SUCCEEDED` - Indicates that the health check execution succeeded. * `FAILED` - Indicates that the health check execution failed. * `TIMED_OUT` - Indicates that the health check execution timed out before completion.
	HealthCheckExecutionStatus *string `json:"HealthCheckExecutionStatus,omitempty"`
	// Health check execution result. Valid only if HealthCheckExecutionStatus is SUCCEEDED. * `UNKNOWN` - Indicates that the health check results could not be determined. * `PASS` - Indicates that the health check passed. * `FAIL` - Indicates that the health check failed. * `WARN` - Indicates that the health check completed with a warning. * `NOT_APPLICABLE` - Indicates that the health check is either unsupported, or not applicable on the Cluster.
	HealthCheckResult *string `json:"HealthCheckResult,omitempty"`
	// A brief summary of health check results.
	HealthCheckSummary *string `json:"HealthCheckSummary,omitempty"`
	// HyperFlex Device Name where the healthcheck is executed.
	HxDeviceName *string `json:"HxDeviceName,omitempty"`
	// Information detailing a suggegsted resolution for the healthcheck failure, if the check fails.
	SuggestedResolution   *string                                     `json:"SuggestedResolution,omitempty"`
	HealthCheckDefinition *HyperflexHealthCheckDefinitionRelationship `json:"HealthCheckDefinition,omitempty"`
	HxCluster             *HyperflexClusterRelationship               `json:"HxCluster,omitempty"`
	RegisteredDevice      *AssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
	Workflow              *WorkflowWorkflowInfoRelationship           `json:"Workflow,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HyperflexHealthCheckExecutionSnapshotAllOf HyperflexHealthCheckExecutionSnapshotAllOf

// NewHyperflexHealthCheckExecutionSnapshotAllOf instantiates a new HyperflexHealthCheckExecutionSnapshotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealthCheckExecutionSnapshotAllOf(classId string, objectType string) *HyperflexHealthCheckExecutionSnapshotAllOf {
	this := HyperflexHealthCheckExecutionSnapshotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var healthCheckExecutionStatus string = "UNKNOWN"
	this.HealthCheckExecutionStatus = &healthCheckExecutionStatus
	var healthCheckResult string = "UNKNOWN"
	this.HealthCheckResult = &healthCheckResult
	return &this
}

// NewHyperflexHealthCheckExecutionSnapshotAllOfWithDefaults instantiates a new HyperflexHealthCheckExecutionSnapshotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthCheckExecutionSnapshotAllOfWithDefaults() *HyperflexHealthCheckExecutionSnapshotAllOf {
	this := HyperflexHealthCheckExecutionSnapshotAllOf{}
	var classId string = "hyperflex.HealthCheckExecutionSnapshot"
	this.ClassId = classId
	var objectType string = "hyperflex.HealthCheckExecutionSnapshot"
	this.ObjectType = objectType
	var healthCheckExecutionStatus string = "UNKNOWN"
	this.HealthCheckExecutionStatus = &healthCheckExecutionStatus
	var healthCheckResult string = "UNKNOWN"
	this.HealthCheckResult = &healthCheckResult
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCause() string {
	if o == nil || o.Cause == nil {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCauseOk() (*string, bool) {
	if o == nil || o.Cause == nil {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasCause() bool {
	if o != nil && o.Cause != nil {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetCause(v string) {
	o.Cause = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetHealthCheckDetails returns the HealthCheckDetails field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDetails() string {
	if o == nil || o.HealthCheckDetails == nil {
		var ret string
		return ret
	}
	return *o.HealthCheckDetails
}

// GetHealthCheckDetailsOk returns a tuple with the HealthCheckDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDetailsOk() (*string, bool) {
	if o == nil || o.HealthCheckDetails == nil {
		return nil, false
	}
	return o.HealthCheckDetails, true
}

// HasHealthCheckDetails returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckDetails() bool {
	if o != nil && o.HealthCheckDetails != nil {
		return true
	}

	return false
}

// SetHealthCheckDetails gets a reference to the given string and assigns it to the HealthCheckDetails field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckDetails(v string) {
	o.HealthCheckDetails = &v
}

// GetHealthCheckExecutionErrorDetails returns the HealthCheckExecutionErrorDetails field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorDetails() string {
	if o == nil || o.HealthCheckExecutionErrorDetails == nil {
		var ret string
		return ret
	}
	return *o.HealthCheckExecutionErrorDetails
}

// GetHealthCheckExecutionErrorDetailsOk returns a tuple with the HealthCheckExecutionErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorDetailsOk() (*string, bool) {
	if o == nil || o.HealthCheckExecutionErrorDetails == nil {
		return nil, false
	}
	return o.HealthCheckExecutionErrorDetails, true
}

// HasHealthCheckExecutionErrorDetails returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckExecutionErrorDetails() bool {
	if o != nil && o.HealthCheckExecutionErrorDetails != nil {
		return true
	}

	return false
}

// SetHealthCheckExecutionErrorDetails gets a reference to the given string and assigns it to the HealthCheckExecutionErrorDetails field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckExecutionErrorDetails(v string) {
	o.HealthCheckExecutionErrorDetails = &v
}

// GetHealthCheckExecutionErrorSummary returns the HealthCheckExecutionErrorSummary field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorSummary() string {
	if o == nil || o.HealthCheckExecutionErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.HealthCheckExecutionErrorSummary
}

// GetHealthCheckExecutionErrorSummaryOk returns a tuple with the HealthCheckExecutionErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionErrorSummaryOk() (*string, bool) {
	if o == nil || o.HealthCheckExecutionErrorSummary == nil {
		return nil, false
	}
	return o.HealthCheckExecutionErrorSummary, true
}

// HasHealthCheckExecutionErrorSummary returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckExecutionErrorSummary() bool {
	if o != nil && o.HealthCheckExecutionErrorSummary != nil {
		return true
	}

	return false
}

// SetHealthCheckExecutionErrorSummary gets a reference to the given string and assigns it to the HealthCheckExecutionErrorSummary field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckExecutionErrorSummary(v string) {
	o.HealthCheckExecutionErrorSummary = &v
}

// GetHealthCheckExecutionStatus returns the HealthCheckExecutionStatus field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionStatus() string {
	if o == nil || o.HealthCheckExecutionStatus == nil {
		var ret string
		return ret
	}
	return *o.HealthCheckExecutionStatus
}

// GetHealthCheckExecutionStatusOk returns a tuple with the HealthCheckExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckExecutionStatusOk() (*string, bool) {
	if o == nil || o.HealthCheckExecutionStatus == nil {
		return nil, false
	}
	return o.HealthCheckExecutionStatus, true
}

// HasHealthCheckExecutionStatus returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckExecutionStatus() bool {
	if o != nil && o.HealthCheckExecutionStatus != nil {
		return true
	}

	return false
}

// SetHealthCheckExecutionStatus gets a reference to the given string and assigns it to the HealthCheckExecutionStatus field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckExecutionStatus(v string) {
	o.HealthCheckExecutionStatus = &v
}

// GetHealthCheckResult returns the HealthCheckResult field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckResult() string {
	if o == nil || o.HealthCheckResult == nil {
		var ret string
		return ret
	}
	return *o.HealthCheckResult
}

// GetHealthCheckResultOk returns a tuple with the HealthCheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckResultOk() (*string, bool) {
	if o == nil || o.HealthCheckResult == nil {
		return nil, false
	}
	return o.HealthCheckResult, true
}

// HasHealthCheckResult returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckResult() bool {
	if o != nil && o.HealthCheckResult != nil {
		return true
	}

	return false
}

// SetHealthCheckResult gets a reference to the given string and assigns it to the HealthCheckResult field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckResult(v string) {
	o.HealthCheckResult = &v
}

// GetHealthCheckSummary returns the HealthCheckSummary field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckSummary() string {
	if o == nil || o.HealthCheckSummary == nil {
		var ret string
		return ret
	}
	return *o.HealthCheckSummary
}

// GetHealthCheckSummaryOk returns a tuple with the HealthCheckSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckSummaryOk() (*string, bool) {
	if o == nil || o.HealthCheckSummary == nil {
		return nil, false
	}
	return o.HealthCheckSummary, true
}

// HasHealthCheckSummary returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckSummary() bool {
	if o != nil && o.HealthCheckSummary != nil {
		return true
	}

	return false
}

// SetHealthCheckSummary gets a reference to the given string and assigns it to the HealthCheckSummary field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckSummary(v string) {
	o.HealthCheckSummary = &v
}

// GetHxDeviceName returns the HxDeviceName field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxDeviceName() string {
	if o == nil || o.HxDeviceName == nil {
		var ret string
		return ret
	}
	return *o.HxDeviceName
}

// GetHxDeviceNameOk returns a tuple with the HxDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxDeviceNameOk() (*string, bool) {
	if o == nil || o.HxDeviceName == nil {
		return nil, false
	}
	return o.HxDeviceName, true
}

// HasHxDeviceName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHxDeviceName() bool {
	if o != nil && o.HxDeviceName != nil {
		return true
	}

	return false
}

// SetHxDeviceName gets a reference to the given string and assigns it to the HxDeviceName field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHxDeviceName(v string) {
	o.HxDeviceName = &v
}

// GetSuggestedResolution returns the SuggestedResolution field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetSuggestedResolution() string {
	if o == nil || o.SuggestedResolution == nil {
		var ret string
		return ret
	}
	return *o.SuggestedResolution
}

// GetSuggestedResolutionOk returns a tuple with the SuggestedResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetSuggestedResolutionOk() (*string, bool) {
	if o == nil || o.SuggestedResolution == nil {
		return nil, false
	}
	return o.SuggestedResolution, true
}

// HasSuggestedResolution returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasSuggestedResolution() bool {
	if o != nil && o.SuggestedResolution != nil {
		return true
	}

	return false
}

// SetSuggestedResolution gets a reference to the given string and assigns it to the SuggestedResolution field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetSuggestedResolution(v string) {
	o.SuggestedResolution = &v
}

// GetHealthCheckDefinition returns the HealthCheckDefinition field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDefinition() HyperflexHealthCheckDefinitionRelationship {
	if o == nil || o.HealthCheckDefinition == nil {
		var ret HyperflexHealthCheckDefinitionRelationship
		return ret
	}
	return *o.HealthCheckDefinition
}

// GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHealthCheckDefinitionOk() (*HyperflexHealthCheckDefinitionRelationship, bool) {
	if o == nil || o.HealthCheckDefinition == nil {
		return nil, false
	}
	return o.HealthCheckDefinition, true
}

// HasHealthCheckDefinition returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHealthCheckDefinition() bool {
	if o != nil && o.HealthCheckDefinition != nil {
		return true
	}

	return false
}

// SetHealthCheckDefinition gets a reference to the given HyperflexHealthCheckDefinitionRelationship and assigns it to the HealthCheckDefinition field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHealthCheckDefinition(v HyperflexHealthCheckDefinitionRelationship) {
	o.HealthCheckDefinition = &v
}

// GetHxCluster returns the HxCluster field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxCluster() HyperflexClusterRelationship {
	if o == nil || o.HxCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.HxCluster
}

// GetHxClusterOk returns a tuple with the HxCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetHxClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.HxCluster == nil {
		return nil, false
	}
	return o.HxCluster, true
}

// HasHxCluster returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasHxCluster() bool {
	if o != nil && o.HxCluster != nil {
		return true
	}

	return false
}

// SetHxCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the HxCluster field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetHxCluster(v HyperflexClusterRelationship) {
	o.HxCluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *HyperflexHealthCheckExecutionSnapshotAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o HyperflexHealthCheckExecutionSnapshotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.Cause != nil {
		toSerialize["Cause"] = o.Cause
	}
	if o.CompletionTime != nil {
		toSerialize["CompletionTime"] = o.CompletionTime
	}
	if o.HealthCheckDetails != nil {
		toSerialize["HealthCheckDetails"] = o.HealthCheckDetails
	}
	if o.HealthCheckExecutionErrorDetails != nil {
		toSerialize["HealthCheckExecutionErrorDetails"] = o.HealthCheckExecutionErrorDetails
	}
	if o.HealthCheckExecutionErrorSummary != nil {
		toSerialize["HealthCheckExecutionErrorSummary"] = o.HealthCheckExecutionErrorSummary
	}
	if o.HealthCheckExecutionStatus != nil {
		toSerialize["HealthCheckExecutionStatus"] = o.HealthCheckExecutionStatus
	}
	if o.HealthCheckResult != nil {
		toSerialize["HealthCheckResult"] = o.HealthCheckResult
	}
	if o.HealthCheckSummary != nil {
		toSerialize["HealthCheckSummary"] = o.HealthCheckSummary
	}
	if o.HxDeviceName != nil {
		toSerialize["HxDeviceName"] = o.HxDeviceName
	}
	if o.SuggestedResolution != nil {
		toSerialize["SuggestedResolution"] = o.SuggestedResolution
	}
	if o.HealthCheckDefinition != nil {
		toSerialize["HealthCheckDefinition"] = o.HealthCheckDefinition
	}
	if o.HxCluster != nil {
		toSerialize["HxCluster"] = o.HxCluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHealthCheckExecutionSnapshotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHealthCheckExecutionSnapshotAllOf := _HyperflexHealthCheckExecutionSnapshotAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHealthCheckExecutionSnapshotAllOf); err == nil {
		*o = HyperflexHealthCheckExecutionSnapshotAllOf(varHyperflexHealthCheckExecutionSnapshotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "Cause")
		delete(additionalProperties, "CompletionTime")
		delete(additionalProperties, "HealthCheckDetails")
		delete(additionalProperties, "HealthCheckExecutionErrorDetails")
		delete(additionalProperties, "HealthCheckExecutionErrorSummary")
		delete(additionalProperties, "HealthCheckExecutionStatus")
		delete(additionalProperties, "HealthCheckResult")
		delete(additionalProperties, "HealthCheckSummary")
		delete(additionalProperties, "HxDeviceName")
		delete(additionalProperties, "SuggestedResolution")
		delete(additionalProperties, "HealthCheckDefinition")
		delete(additionalProperties, "HxCluster")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHealthCheckExecutionSnapshotAllOf struct {
	value *HyperflexHealthCheckExecutionSnapshotAllOf
	isSet bool
}

func (v NullableHyperflexHealthCheckExecutionSnapshotAllOf) Get() *HyperflexHealthCheckExecutionSnapshotAllOf {
	return v.value
}

func (v *NullableHyperflexHealthCheckExecutionSnapshotAllOf) Set(val *HyperflexHealthCheckExecutionSnapshotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthCheckExecutionSnapshotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthCheckExecutionSnapshotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthCheckExecutionSnapshotAllOf(val *HyperflexHealthCheckExecutionSnapshotAllOf) *NullableHyperflexHealthCheckExecutionSnapshotAllOf {
	return &NullableHyperflexHealthCheckExecutionSnapshotAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHealthCheckExecutionSnapshotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthCheckExecutionSnapshotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
