/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TamSecurityAdvisoryDetails Details pertaining to a security advisory defined by a given advisory.
type TamSecurityAdvisoryDetails struct {
	TamBaseAdvisoryDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CVSS version 3 base score for the security Advisory.
	BaseScore *float32 `json:"BaseScore,omitempty"`
	CveIds    []string `json:"CveIds,omitempty"`
	// CVSS version 3 environmental score for the security Advisory.
	EnvironmentalScore *float32 `json:"EnvironmentalScore,omitempty"`
	// Cisco assigned status of the published advisory. Depends on whether the investigation is complete or on-going. * `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available. * `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.
	Status *string `json:"Status,omitempty"`
	// CVSS version 3 temporal score for the security Advisory.
	TemporalScore        *float32 `json:"TemporalScore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamSecurityAdvisoryDetails TamSecurityAdvisoryDetails

// NewTamSecurityAdvisoryDetails instantiates a new TamSecurityAdvisoryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamSecurityAdvisoryDetails(classId string, objectType string) *TamSecurityAdvisoryDetails {
	this := TamSecurityAdvisoryDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "interim"
	this.Status = &status
	return &this
}

// NewTamSecurityAdvisoryDetailsWithDefaults instantiates a new TamSecurityAdvisoryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamSecurityAdvisoryDetailsWithDefaults() *TamSecurityAdvisoryDetails {
	this := TamSecurityAdvisoryDetails{}
	var classId string = "tam.SecurityAdvisoryDetails"
	this.ClassId = classId
	var objectType string = "tam.SecurityAdvisoryDetails"
	this.ObjectType = objectType
	var status string = "interim"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamSecurityAdvisoryDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamSecurityAdvisoryDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamSecurityAdvisoryDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamSecurityAdvisoryDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBaseScore returns the BaseScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetBaseScore() float32 {
	if o == nil || o.BaseScore == nil {
		var ret float32
		return ret
	}
	return *o.BaseScore
}

// GetBaseScoreOk returns a tuple with the BaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetBaseScoreOk() (*float32, bool) {
	if o == nil || o.BaseScore == nil {
		return nil, false
	}
	return o.BaseScore, true
}

// HasBaseScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasBaseScore() bool {
	if o != nil && o.BaseScore != nil {
		return true
	}

	return false
}

// SetBaseScore gets a reference to the given float32 and assigns it to the BaseScore field.
func (o *TamSecurityAdvisoryDetails) SetBaseScore(v float32) {
	o.BaseScore = &v
}

// GetCveIds returns the CveIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamSecurityAdvisoryDetails) GetCveIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CveIds
}

// GetCveIdsOk returns a tuple with the CveIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamSecurityAdvisoryDetails) GetCveIdsOk() (*[]string, bool) {
	if o == nil || o.CveIds == nil {
		return nil, false
	}
	return &o.CveIds, true
}

// HasCveIds returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasCveIds() bool {
	if o != nil && o.CveIds != nil {
		return true
	}

	return false
}

// SetCveIds gets a reference to the given []string and assigns it to the CveIds field.
func (o *TamSecurityAdvisoryDetails) SetCveIds(v []string) {
	o.CveIds = v
}

// GetEnvironmentalScore returns the EnvironmentalScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetEnvironmentalScore() float32 {
	if o == nil || o.EnvironmentalScore == nil {
		var ret float32
		return ret
	}
	return *o.EnvironmentalScore
}

// GetEnvironmentalScoreOk returns a tuple with the EnvironmentalScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetEnvironmentalScoreOk() (*float32, bool) {
	if o == nil || o.EnvironmentalScore == nil {
		return nil, false
	}
	return o.EnvironmentalScore, true
}

// HasEnvironmentalScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasEnvironmentalScore() bool {
	if o != nil && o.EnvironmentalScore != nil {
		return true
	}

	return false
}

// SetEnvironmentalScore gets a reference to the given float32 and assigns it to the EnvironmentalScore field.
func (o *TamSecurityAdvisoryDetails) SetEnvironmentalScore(v float32) {
	o.EnvironmentalScore = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TamSecurityAdvisoryDetails) SetStatus(v string) {
	o.Status = &v
}

// GetTemporalScore returns the TemporalScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetTemporalScore() float32 {
	if o == nil || o.TemporalScore == nil {
		var ret float32
		return ret
	}
	return *o.TemporalScore
}

// GetTemporalScoreOk returns a tuple with the TemporalScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetTemporalScoreOk() (*float32, bool) {
	if o == nil || o.TemporalScore == nil {
		return nil, false
	}
	return o.TemporalScore, true
}

// HasTemporalScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasTemporalScore() bool {
	if o != nil && o.TemporalScore != nil {
		return true
	}

	return false
}

// SetTemporalScore gets a reference to the given float32 and assigns it to the TemporalScore field.
func (o *TamSecurityAdvisoryDetails) SetTemporalScore(v float32) {
	o.TemporalScore = &v
}

func (o TamSecurityAdvisoryDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTamBaseAdvisoryDetails, errTamBaseAdvisoryDetails := json.Marshal(o.TamBaseAdvisoryDetails)
	if errTamBaseAdvisoryDetails != nil {
		return []byte{}, errTamBaseAdvisoryDetails
	}
	errTamBaseAdvisoryDetails = json.Unmarshal([]byte(serializedTamBaseAdvisoryDetails), &toSerialize)
	if errTamBaseAdvisoryDetails != nil {
		return []byte{}, errTamBaseAdvisoryDetails
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BaseScore != nil {
		toSerialize["BaseScore"] = o.BaseScore
	}
	if o.CveIds != nil {
		toSerialize["CveIds"] = o.CveIds
	}
	if o.EnvironmentalScore != nil {
		toSerialize["EnvironmentalScore"] = o.EnvironmentalScore
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TemporalScore != nil {
		toSerialize["TemporalScore"] = o.TemporalScore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamSecurityAdvisoryDetails) UnmarshalJSON(bytes []byte) (err error) {
	type TamSecurityAdvisoryDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// CVSS version 3 base score for the security Advisory.
		BaseScore *float32 `json:"BaseScore,omitempty"`
		CveIds    []string `json:"CveIds,omitempty"`
		// CVSS version 3 environmental score for the security Advisory.
		EnvironmentalScore *float32 `json:"EnvironmentalScore,omitempty"`
		// Cisco assigned status of the published advisory. Depends on whether the investigation is complete or on-going. * `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available. * `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.
		Status *string `json:"Status,omitempty"`
		// CVSS version 3 temporal score for the security Advisory.
		TemporalScore *float32 `json:"TemporalScore,omitempty"`
	}

	varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct := TamSecurityAdvisoryDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct)
	if err == nil {
		varTamSecurityAdvisoryDetails := _TamSecurityAdvisoryDetails{}
		varTamSecurityAdvisoryDetails.ClassId = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.ClassId
		varTamSecurityAdvisoryDetails.ObjectType = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.ObjectType
		varTamSecurityAdvisoryDetails.BaseScore = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.BaseScore
		varTamSecurityAdvisoryDetails.CveIds = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.CveIds
		varTamSecurityAdvisoryDetails.EnvironmentalScore = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.EnvironmentalScore
		varTamSecurityAdvisoryDetails.Status = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.Status
		varTamSecurityAdvisoryDetails.TemporalScore = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.TemporalScore
		*o = TamSecurityAdvisoryDetails(varTamSecurityAdvisoryDetails)
	} else {
		return err
	}

	varTamSecurityAdvisoryDetails := _TamSecurityAdvisoryDetails{}

	err = json.Unmarshal(bytes, &varTamSecurityAdvisoryDetails)
	if err == nil {
		o.TamBaseAdvisoryDetails = varTamSecurityAdvisoryDetails.TamBaseAdvisoryDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseScore")
		delete(additionalProperties, "CveIds")
		delete(additionalProperties, "EnvironmentalScore")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TemporalScore")

		// remove fields from embedded structs
		reflectTamBaseAdvisoryDetails := reflect.ValueOf(o.TamBaseAdvisoryDetails)
		for i := 0; i < reflectTamBaseAdvisoryDetails.Type().NumField(); i++ {
			t := reflectTamBaseAdvisoryDetails.Type().Field(i)

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

type NullableTamSecurityAdvisoryDetails struct {
	value *TamSecurityAdvisoryDetails
	isSet bool
}

func (v NullableTamSecurityAdvisoryDetails) Get() *TamSecurityAdvisoryDetails {
	return v.value
}

func (v *NullableTamSecurityAdvisoryDetails) Set(val *TamSecurityAdvisoryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTamSecurityAdvisoryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTamSecurityAdvisoryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamSecurityAdvisoryDetails(val *TamSecurityAdvisoryDetails) *NullableTamSecurityAdvisoryDetails {
	return &NullableTamSecurityAdvisoryDetails{value: val, isSet: true}
}

func (v NullableTamSecurityAdvisoryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamSecurityAdvisoryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
