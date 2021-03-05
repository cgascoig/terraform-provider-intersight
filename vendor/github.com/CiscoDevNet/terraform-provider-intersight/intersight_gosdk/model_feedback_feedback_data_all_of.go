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

// FeedbackFeedbackDataAllOf Definition of the list of properties defined in 'feedback.FeedbackData', excluding properties defined in parent classes.
type FeedbackFeedbackDataAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Account name of the feedback sender. Copied in order to be persisted in case of account removal.
	AccountName               *string  `json:"AccountName,omitempty"`
	AlternativeFollowUpEmails []string `json:"AlternativeFollowUpEmails,omitempty"`
	// Text of the feedback as provided by the user, if it is a bug or a comment.
	Comment *string `json:"Comment,omitempty"`
	// User's email address details.
	Email *string `json:"Email,omitempty"`
	// Evalation rating as provided by the user to capture user sentiment regarding the issue. * `Excellent` - Option that specifies user's excelent evaluation. * `Poor` - Option that specifies user's poor evaluation. * `Fair` - Option that specifies user's fair evaluation. * `Good` - Option that specifies user's good evaluation.
	Evaluation *string `json:"Evaluation,omitempty"`
	// If a user is open for follow-up or not.
	FollowUp *bool `json:"FollowUp,omitempty"`
	// Bunch of last traceId for reproducing user last activity.
	TraceIds interface{} `json:"TraceIds,omitempty"`
	// Type of the feedback from user. * `Evaluation` - User's feedback classified as an evaluation. * `Bug` - User's feedback classified as a bug.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeedbackFeedbackDataAllOf FeedbackFeedbackDataAllOf

// NewFeedbackFeedbackDataAllOf instantiates a new FeedbackFeedbackDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackFeedbackDataAllOf(classId string, objectType string) *FeedbackFeedbackDataAllOf {
	this := FeedbackFeedbackDataAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var evaluation string = "Excellent"
	this.Evaluation = &evaluation
	var type_ string = "Evaluation"
	this.Type = &type_
	return &this
}

// NewFeedbackFeedbackDataAllOfWithDefaults instantiates a new FeedbackFeedbackDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackFeedbackDataAllOfWithDefaults() *FeedbackFeedbackDataAllOf {
	this := FeedbackFeedbackDataAllOf{}
	var classId string = "feedback.FeedbackData"
	this.ClassId = classId
	var objectType string = "feedback.FeedbackData"
	this.ObjectType = objectType
	var evaluation string = "Excellent"
	this.Evaluation = &evaluation
	var type_ string = "Evaluation"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *FeedbackFeedbackDataAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FeedbackFeedbackDataAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FeedbackFeedbackDataAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FeedbackFeedbackDataAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *FeedbackFeedbackDataAllOf) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *FeedbackFeedbackDataAllOf) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAlternativeFollowUpEmails returns the AlternativeFollowUpEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeedbackFeedbackDataAllOf) GetAlternativeFollowUpEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AlternativeFollowUpEmails
}

// GetAlternativeFollowUpEmailsOk returns a tuple with the AlternativeFollowUpEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeedbackFeedbackDataAllOf) GetAlternativeFollowUpEmailsOk() (*[]string, bool) {
	if o == nil || o.AlternativeFollowUpEmails == nil {
		return nil, false
	}
	return &o.AlternativeFollowUpEmails, true
}

// HasAlternativeFollowUpEmails returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasAlternativeFollowUpEmails() bool {
	if o != nil && o.AlternativeFollowUpEmails != nil {
		return true
	}

	return false
}

// SetAlternativeFollowUpEmails gets a reference to the given []string and assigns it to the AlternativeFollowUpEmails field.
func (o *FeedbackFeedbackDataAllOf) SetAlternativeFollowUpEmails(v []string) {
	o.AlternativeFollowUpEmails = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *FeedbackFeedbackDataAllOf) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *FeedbackFeedbackDataAllOf) SetComment(v string) {
	o.Comment = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *FeedbackFeedbackDataAllOf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *FeedbackFeedbackDataAllOf) SetEmail(v string) {
	o.Email = &v
}

// GetEvaluation returns the Evaluation field value if set, zero value otherwise.
func (o *FeedbackFeedbackDataAllOf) GetEvaluation() string {
	if o == nil || o.Evaluation == nil {
		var ret string
		return ret
	}
	return *o.Evaluation
}

// GetEvaluationOk returns a tuple with the Evaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetEvaluationOk() (*string, bool) {
	if o == nil || o.Evaluation == nil {
		return nil, false
	}
	return o.Evaluation, true
}

// HasEvaluation returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasEvaluation() bool {
	if o != nil && o.Evaluation != nil {
		return true
	}

	return false
}

// SetEvaluation gets a reference to the given string and assigns it to the Evaluation field.
func (o *FeedbackFeedbackDataAllOf) SetEvaluation(v string) {
	o.Evaluation = &v
}

// GetFollowUp returns the FollowUp field value if set, zero value otherwise.
func (o *FeedbackFeedbackDataAllOf) GetFollowUp() bool {
	if o == nil || o.FollowUp == nil {
		var ret bool
		return ret
	}
	return *o.FollowUp
}

// GetFollowUpOk returns a tuple with the FollowUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetFollowUpOk() (*bool, bool) {
	if o == nil || o.FollowUp == nil {
		return nil, false
	}
	return o.FollowUp, true
}

// HasFollowUp returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasFollowUp() bool {
	if o != nil && o.FollowUp != nil {
		return true
	}

	return false
}

// SetFollowUp gets a reference to the given bool and assigns it to the FollowUp field.
func (o *FeedbackFeedbackDataAllOf) SetFollowUp(v bool) {
	o.FollowUp = &v
}

// GetTraceIds returns the TraceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeedbackFeedbackDataAllOf) GetTraceIds() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TraceIds
}

// GetTraceIdsOk returns a tuple with the TraceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeedbackFeedbackDataAllOf) GetTraceIdsOk() (*interface{}, bool) {
	if o == nil || o.TraceIds == nil {
		return nil, false
	}
	return &o.TraceIds, true
}

// HasTraceIds returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasTraceIds() bool {
	if o != nil && o.TraceIds != nil {
		return true
	}

	return false
}

// SetTraceIds gets a reference to the given interface{} and assigns it to the TraceIds field.
func (o *FeedbackFeedbackDataAllOf) SetTraceIds(v interface{}) {
	o.TraceIds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeedbackFeedbackDataAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackDataAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeedbackFeedbackDataAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FeedbackFeedbackDataAllOf) SetType(v string) {
	o.Type = &v
}

func (o FeedbackFeedbackDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountName != nil {
		toSerialize["AccountName"] = o.AccountName
	}
	if o.AlternativeFollowUpEmails != nil {
		toSerialize["AlternativeFollowUpEmails"] = o.AlternativeFollowUpEmails
	}
	if o.Comment != nil {
		toSerialize["Comment"] = o.Comment
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.Evaluation != nil {
		toSerialize["Evaluation"] = o.Evaluation
	}
	if o.FollowUp != nil {
		toSerialize["FollowUp"] = o.FollowUp
	}
	if o.TraceIds != nil {
		toSerialize["TraceIds"] = o.TraceIds
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FeedbackFeedbackDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFeedbackFeedbackDataAllOf := _FeedbackFeedbackDataAllOf{}

	if err = json.Unmarshal(bytes, &varFeedbackFeedbackDataAllOf); err == nil {
		*o = FeedbackFeedbackDataAllOf(varFeedbackFeedbackDataAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountName")
		delete(additionalProperties, "AlternativeFollowUpEmails")
		delete(additionalProperties, "Comment")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "Evaluation")
		delete(additionalProperties, "FollowUp")
		delete(additionalProperties, "TraceIds")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeedbackFeedbackDataAllOf struct {
	value *FeedbackFeedbackDataAllOf
	isSet bool
}

func (v NullableFeedbackFeedbackDataAllOf) Get() *FeedbackFeedbackDataAllOf {
	return v.value
}

func (v *NullableFeedbackFeedbackDataAllOf) Set(val *FeedbackFeedbackDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackFeedbackDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackFeedbackDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackFeedbackDataAllOf(val *FeedbackFeedbackDataAllOf) *NullableFeedbackFeedbackDataAllOf {
	return &NullableFeedbackFeedbackDataAllOf{value: val, isSet: true}
}

func (v NullableFeedbackFeedbackDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackFeedbackDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
