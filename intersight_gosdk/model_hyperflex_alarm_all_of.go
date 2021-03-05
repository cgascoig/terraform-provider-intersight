/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-24T06:47:07Z.
 *
 * API version: 1.0.9-3824
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexAlarmAllOf Definition of the list of properties defined in 'hyperflex.Alarm', excluding properties defined in parent classes.
type HyperflexAlarmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The acknowledgement state of the alarm. It is 'true' when the alarm is acknowledged and false otherwise.
	Acknowledged *bool `json:"Acknowledged,omitempty"`
	// The username of the user who acknowledged the alarm.
	AcknowledgedBy *string `json:"AcknowledgedBy,omitempty"`
	// The time when the alarm was acknowledged, represented as a Unix timestamp.
	AcknowledgedTime *int64 `json:"AcknowledgedTime,omitempty"`
	// The time when the alarm was acknowledged in ISO 6801 format.
	AcknowledgedTimeAsUtc *string `json:"AcknowledgedTimeAsUtc,omitempty"`
	// The description of the alarm which includes information about the fault that triggered the alarm.
	Description *string `json:"Description,omitempty"`
	// The data pertaining to this entity.
	EntityData *string `json:"EntityData,omitempty"`
	// The name of entity which triggered the alarm.
	EntityName *string `json:"EntityName,omitempty"`
	// The type of entity which triggered the alarm. For example, this can be the cluster, a node, or VM running on the cluster. * `UNKNOWN` - The type of entity is not known. * `DISK` - The entity is a physical storage device. * `NODE` - The entity is a HyperFlex cluster node. * `CLUSTER` - The entity is the HyperFlex cluster itself. * `DATASTORE` - The entity is a logical datastore configured on the HyperFlex cluster. * `ZONE` - The entity is a logical or physical zone configured on the HyperFlex cluster. * `VIRTUALMACHINE` - The entity is a virtual machine running on the HyperFlex cluster.
	EntityType *string `json:"EntityType,omitempty"`
	// The unique identifier of the entity which triggered the alarm.
	EntityUuId *string `json:"EntityUuId,omitempty"`
	// The localized message displayed to the user which describes the alarm.
	Message *string `json:"Message,omitempty"`
	// The name of the alarm. This name identifies the type of alarm that was triggered.
	Name *string `json:"Name,omitempty"`
	// The severity of the alarm. * `UNKNOWN` - The alarm status is not known. * `CLEARED` - The event that triggered the alarm has been remedied and no longer requires the user's attention. * `INFO` - The alarm represents a message that does not require the user's immediate attention. * `WARNING` - The alarm represents a moderate fault. * `CRITICAL` - The alarm represents a critical fault.
	Status *string `json:"Status,omitempty"`
	// The time when alarm was triggered as a Unix timestamp.
	TriggeredTime *int64 `json:"TriggeredTime,omitempty"`
	// The time when alarm was triggered in ISO 6801 UTC format.
	TriggeredTimeAsUtc *string `json:"TriggeredTimeAsUtc,omitempty"`
	// The unique identifier for this alarm instance.
	Uuid                 *string                       `json:"Uuid,omitempty"`
	Cluster              *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAlarmAllOf HyperflexAlarmAllOf

// NewHyperflexAlarmAllOf instantiates a new HyperflexAlarmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAlarmAllOf(classId string, objectType string) *HyperflexAlarmAllOf {
	this := HyperflexAlarmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var entityType string = "UNKNOWN"
	this.EntityType = &entityType
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// NewHyperflexAlarmAllOfWithDefaults instantiates a new HyperflexAlarmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAlarmAllOfWithDefaults() *HyperflexAlarmAllOf {
	this := HyperflexAlarmAllOf{}
	var classId string = "hyperflex.Alarm"
	this.ClassId = classId
	var objectType string = "hyperflex.Alarm"
	this.ObjectType = objectType
	var entityType string = "UNKNOWN"
	this.EntityType = &entityType
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAlarmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAlarmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAlarmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAlarmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAcknowledged returns the Acknowledged field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetAcknowledged() bool {
	if o == nil || o.Acknowledged == nil {
		var ret bool
		return ret
	}
	return *o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetAcknowledgedOk() (*bool, bool) {
	if o == nil || o.Acknowledged == nil {
		return nil, false
	}
	return o.Acknowledged, true
}

// HasAcknowledged returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasAcknowledged() bool {
	if o != nil && o.Acknowledged != nil {
		return true
	}

	return false
}

// SetAcknowledged gets a reference to the given bool and assigns it to the Acknowledged field.
func (o *HyperflexAlarmAllOf) SetAcknowledged(v bool) {
	o.Acknowledged = &v
}

// GetAcknowledgedBy returns the AcknowledgedBy field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetAcknowledgedBy() string {
	if o == nil || o.AcknowledgedBy == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgedBy
}

// GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetAcknowledgedByOk() (*string, bool) {
	if o == nil || o.AcknowledgedBy == nil {
		return nil, false
	}
	return o.AcknowledgedBy, true
}

// HasAcknowledgedBy returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasAcknowledgedBy() bool {
	if o != nil && o.AcknowledgedBy != nil {
		return true
	}

	return false
}

// SetAcknowledgedBy gets a reference to the given string and assigns it to the AcknowledgedBy field.
func (o *HyperflexAlarmAllOf) SetAcknowledgedBy(v string) {
	o.AcknowledgedBy = &v
}

// GetAcknowledgedTime returns the AcknowledgedTime field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetAcknowledgedTime() int64 {
	if o == nil || o.AcknowledgedTime == nil {
		var ret int64
		return ret
	}
	return *o.AcknowledgedTime
}

// GetAcknowledgedTimeOk returns a tuple with the AcknowledgedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetAcknowledgedTimeOk() (*int64, bool) {
	if o == nil || o.AcknowledgedTime == nil {
		return nil, false
	}
	return o.AcknowledgedTime, true
}

// HasAcknowledgedTime returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasAcknowledgedTime() bool {
	if o != nil && o.AcknowledgedTime != nil {
		return true
	}

	return false
}

// SetAcknowledgedTime gets a reference to the given int64 and assigns it to the AcknowledgedTime field.
func (o *HyperflexAlarmAllOf) SetAcknowledgedTime(v int64) {
	o.AcknowledgedTime = &v
}

// GetAcknowledgedTimeAsUtc returns the AcknowledgedTimeAsUtc field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetAcknowledgedTimeAsUtc() string {
	if o == nil || o.AcknowledgedTimeAsUtc == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgedTimeAsUtc
}

// GetAcknowledgedTimeAsUtcOk returns a tuple with the AcknowledgedTimeAsUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetAcknowledgedTimeAsUtcOk() (*string, bool) {
	if o == nil || o.AcknowledgedTimeAsUtc == nil {
		return nil, false
	}
	return o.AcknowledgedTimeAsUtc, true
}

// HasAcknowledgedTimeAsUtc returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasAcknowledgedTimeAsUtc() bool {
	if o != nil && o.AcknowledgedTimeAsUtc != nil {
		return true
	}

	return false
}

// SetAcknowledgedTimeAsUtc gets a reference to the given string and assigns it to the AcknowledgedTimeAsUtc field.
func (o *HyperflexAlarmAllOf) SetAcknowledgedTimeAsUtc(v string) {
	o.AcknowledgedTimeAsUtc = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexAlarmAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEntityData returns the EntityData field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetEntityData() string {
	if o == nil || o.EntityData == nil {
		var ret string
		return ret
	}
	return *o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetEntityDataOk() (*string, bool) {
	if o == nil || o.EntityData == nil {
		return nil, false
	}
	return o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasEntityData() bool {
	if o != nil && o.EntityData != nil {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given string and assigns it to the EntityData field.
func (o *HyperflexAlarmAllOf) SetEntityData(v string) {
	o.EntityData = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetEntityName() string {
	if o == nil || o.EntityName == nil {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetEntityNameOk() (*string, bool) {
	if o == nil || o.EntityName == nil {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasEntityName() bool {
	if o != nil && o.EntityName != nil {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *HyperflexAlarmAllOf) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *HyperflexAlarmAllOf) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityUuId returns the EntityUuId field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetEntityUuId() string {
	if o == nil || o.EntityUuId == nil {
		var ret string
		return ret
	}
	return *o.EntityUuId
}

// GetEntityUuIdOk returns a tuple with the EntityUuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetEntityUuIdOk() (*string, bool) {
	if o == nil || o.EntityUuId == nil {
		return nil, false
	}
	return o.EntityUuId, true
}

// HasEntityUuId returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasEntityUuId() bool {
	if o != nil && o.EntityUuId != nil {
		return true
	}

	return false
}

// SetEntityUuId gets a reference to the given string and assigns it to the EntityUuId field.
func (o *HyperflexAlarmAllOf) SetEntityUuId(v string) {
	o.EntityUuId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HyperflexAlarmAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexAlarmAllOf) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexAlarmAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTriggeredTime returns the TriggeredTime field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetTriggeredTime() int64 {
	if o == nil || o.TriggeredTime == nil {
		var ret int64
		return ret
	}
	return *o.TriggeredTime
}

// GetTriggeredTimeOk returns a tuple with the TriggeredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetTriggeredTimeOk() (*int64, bool) {
	if o == nil || o.TriggeredTime == nil {
		return nil, false
	}
	return o.TriggeredTime, true
}

// HasTriggeredTime returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasTriggeredTime() bool {
	if o != nil && o.TriggeredTime != nil {
		return true
	}

	return false
}

// SetTriggeredTime gets a reference to the given int64 and assigns it to the TriggeredTime field.
func (o *HyperflexAlarmAllOf) SetTriggeredTime(v int64) {
	o.TriggeredTime = &v
}

// GetTriggeredTimeAsUtc returns the TriggeredTimeAsUtc field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetTriggeredTimeAsUtc() string {
	if o == nil || o.TriggeredTimeAsUtc == nil {
		var ret string
		return ret
	}
	return *o.TriggeredTimeAsUtc
}

// GetTriggeredTimeAsUtcOk returns a tuple with the TriggeredTimeAsUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetTriggeredTimeAsUtcOk() (*string, bool) {
	if o == nil || o.TriggeredTimeAsUtc == nil {
		return nil, false
	}
	return o.TriggeredTimeAsUtc, true
}

// HasTriggeredTimeAsUtc returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasTriggeredTimeAsUtc() bool {
	if o != nil && o.TriggeredTimeAsUtc != nil {
		return true
	}

	return false
}

// SetTriggeredTimeAsUtc gets a reference to the given string and assigns it to the TriggeredTimeAsUtc field.
func (o *HyperflexAlarmAllOf) SetTriggeredTimeAsUtc(v string) {
	o.TriggeredTimeAsUtc = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexAlarmAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexAlarmAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexAlarmAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexAlarmAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

func (o HyperflexAlarmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Acknowledged != nil {
		toSerialize["Acknowledged"] = o.Acknowledged
	}
	if o.AcknowledgedBy != nil {
		toSerialize["AcknowledgedBy"] = o.AcknowledgedBy
	}
	if o.AcknowledgedTime != nil {
		toSerialize["AcknowledgedTime"] = o.AcknowledgedTime
	}
	if o.AcknowledgedTimeAsUtc != nil {
		toSerialize["AcknowledgedTimeAsUtc"] = o.AcknowledgedTimeAsUtc
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EntityData != nil {
		toSerialize["EntityData"] = o.EntityData
	}
	if o.EntityName != nil {
		toSerialize["EntityName"] = o.EntityName
	}
	if o.EntityType != nil {
		toSerialize["EntityType"] = o.EntityType
	}
	if o.EntityUuId != nil {
		toSerialize["EntityUuId"] = o.EntityUuId
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TriggeredTime != nil {
		toSerialize["TriggeredTime"] = o.TriggeredTime
	}
	if o.TriggeredTimeAsUtc != nil {
		toSerialize["TriggeredTimeAsUtc"] = o.TriggeredTimeAsUtc
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAlarmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexAlarmAllOf := _HyperflexAlarmAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexAlarmAllOf); err == nil {
		*o = HyperflexAlarmAllOf(varHyperflexAlarmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Acknowledged")
		delete(additionalProperties, "AcknowledgedBy")
		delete(additionalProperties, "AcknowledgedTime")
		delete(additionalProperties, "AcknowledgedTimeAsUtc")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EntityData")
		delete(additionalProperties, "EntityName")
		delete(additionalProperties, "EntityType")
		delete(additionalProperties, "EntityUuId")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TriggeredTime")
		delete(additionalProperties, "TriggeredTimeAsUtc")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Cluster")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAlarmAllOf struct {
	value *HyperflexAlarmAllOf
	isSet bool
}

func (v NullableHyperflexAlarmAllOf) Get() *HyperflexAlarmAllOf {
	return v.value
}

func (v *NullableHyperflexAlarmAllOf) Set(val *HyperflexAlarmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAlarmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAlarmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAlarmAllOf(val *HyperflexAlarmAllOf) *NullableHyperflexAlarmAllOf {
	return &NullableHyperflexAlarmAllOf{value: val, isSet: true}
}

func (v NullableHyperflexAlarmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAlarmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
