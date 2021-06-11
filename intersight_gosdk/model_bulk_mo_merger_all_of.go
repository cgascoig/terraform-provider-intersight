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

// BulkMoMergerAllOf Definition of the list of properties defined in 'bulk.MoMerger', excluding properties defined in parent classes.
type BulkMoMergerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of merge action to be applied on the target MOs.  * `Merge` - The null properties/relationships of the source MO will be ignored for the target MO. The non-null properties/relationships of the source will override the target MO properties/relationships. * `Replace` - Merge action as described in RFC 7386. The null properties/relationships of the source MO will be deleted on the target MO.The non-null properties/relationships of the source will override the target MO properties/relationships.When source object type is different from target, only the properties common to both source and target  will be affected.Other properties on the target will be ignored.
	MergeAction          *string                               `json:"MergeAction,omitempty"`
	Responses            []BulkRestResult                      `json:"Responses,omitempty"`
	Sources              []MoBaseMo                            `json:"Sources,omitempty"`
	TargetConfig         *MoBaseMo                             `json:"TargetConfig,omitempty"`
	Targets              []MoBaseMo                            `json:"Targets,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkMoMergerAllOf BulkMoMergerAllOf

// NewBulkMoMergerAllOf instantiates a new BulkMoMergerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkMoMergerAllOf(classId string, objectType string) *BulkMoMergerAllOf {
	this := BulkMoMergerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mergeAction string = "Merge"
	this.MergeAction = &mergeAction
	return &this
}

// NewBulkMoMergerAllOfWithDefaults instantiates a new BulkMoMergerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkMoMergerAllOfWithDefaults() *BulkMoMergerAllOf {
	this := BulkMoMergerAllOf{}
	var classId string = "bulk.MoMerger"
	this.ClassId = classId
	var objectType string = "bulk.MoMerger"
	this.ObjectType = objectType
	var mergeAction string = "Merge"
	this.MergeAction = &mergeAction
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkMoMergerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkMoMergerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkMoMergerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkMoMergerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkMoMergerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkMoMergerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMergeAction returns the MergeAction field value if set, zero value otherwise.
func (o *BulkMoMergerAllOf) GetMergeAction() string {
	if o == nil || o.MergeAction == nil {
		var ret string
		return ret
	}
	return *o.MergeAction
}

// GetMergeActionOk returns a tuple with the MergeAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoMergerAllOf) GetMergeActionOk() (*string, bool) {
	if o == nil || o.MergeAction == nil {
		return nil, false
	}
	return o.MergeAction, true
}

// HasMergeAction returns a boolean if a field has been set.
func (o *BulkMoMergerAllOf) HasMergeAction() bool {
	if o != nil && o.MergeAction != nil {
		return true
	}

	return false
}

// SetMergeAction gets a reference to the given string and assigns it to the MergeAction field.
func (o *BulkMoMergerAllOf) SetMergeAction(v string) {
	o.MergeAction = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoMergerAllOf) GetResponses() []BulkRestResult {
	if o == nil {
		var ret []BulkRestResult
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoMergerAllOf) GetResponsesOk() (*[]BulkRestResult, bool) {
	if o == nil || o.Responses == nil {
		return nil, false
	}
	return &o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *BulkMoMergerAllOf) HasResponses() bool {
	if o != nil && o.Responses != nil {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []BulkRestResult and assigns it to the Responses field.
func (o *BulkMoMergerAllOf) SetResponses(v []BulkRestResult) {
	o.Responses = v
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoMergerAllOf) GetSources() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoMergerAllOf) GetSourcesOk() (*[]MoBaseMo, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *BulkMoMergerAllOf) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []MoBaseMo and assigns it to the Sources field.
func (o *BulkMoMergerAllOf) SetSources(v []MoBaseMo) {
	o.Sources = v
}

// GetTargetConfig returns the TargetConfig field value if set, zero value otherwise.
func (o *BulkMoMergerAllOf) GetTargetConfig() MoBaseMo {
	if o == nil || o.TargetConfig == nil {
		var ret MoBaseMo
		return ret
	}
	return *o.TargetConfig
}

// GetTargetConfigOk returns a tuple with the TargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoMergerAllOf) GetTargetConfigOk() (*MoBaseMo, bool) {
	if o == nil || o.TargetConfig == nil {
		return nil, false
	}
	return o.TargetConfig, true
}

// HasTargetConfig returns a boolean if a field has been set.
func (o *BulkMoMergerAllOf) HasTargetConfig() bool {
	if o != nil && o.TargetConfig != nil {
		return true
	}

	return false
}

// SetTargetConfig gets a reference to the given MoBaseMo and assigns it to the TargetConfig field.
func (o *BulkMoMergerAllOf) SetTargetConfig(v MoBaseMo) {
	o.TargetConfig = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoMergerAllOf) GetTargets() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoMergerAllOf) GetTargetsOk() (*[]MoBaseMo, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *BulkMoMergerAllOf) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MoBaseMo and assigns it to the Targets field.
func (o *BulkMoMergerAllOf) SetTargets(v []MoBaseMo) {
	o.Targets = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkMoMergerAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoMergerAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkMoMergerAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkMoMergerAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o BulkMoMergerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MergeAction != nil {
		toSerialize["MergeAction"] = o.MergeAction
	}
	if o.Responses != nil {
		toSerialize["Responses"] = o.Responses
	}
	if o.Sources != nil {
		toSerialize["Sources"] = o.Sources
	}
	if o.TargetConfig != nil {
		toSerialize["TargetConfig"] = o.TargetConfig
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkMoMergerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkMoMergerAllOf := _BulkMoMergerAllOf{}

	if err = json.Unmarshal(bytes, &varBulkMoMergerAllOf); err == nil {
		*o = BulkMoMergerAllOf(varBulkMoMergerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MergeAction")
		delete(additionalProperties, "Responses")
		delete(additionalProperties, "Sources")
		delete(additionalProperties, "TargetConfig")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkMoMergerAllOf struct {
	value *BulkMoMergerAllOf
	isSet bool
}

func (v NullableBulkMoMergerAllOf) Get() *BulkMoMergerAllOf {
	return v.value
}

func (v *NullableBulkMoMergerAllOf) Set(val *BulkMoMergerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkMoMergerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkMoMergerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkMoMergerAllOf(val *BulkMoMergerAllOf) *NullableBulkMoMergerAllOf {
	return &NullableBulkMoMergerAllOf{value: val, isSet: true}
}

func (v NullableBulkMoMergerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkMoMergerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
