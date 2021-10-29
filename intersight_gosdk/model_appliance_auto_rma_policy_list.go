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

// ApplianceAutoRmaPolicyList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type ApplianceAutoRmaPolicyList struct {
	MoBaseResponse
	// The total number of 'appliance.AutoRmaPolicy' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'appliance.AutoRmaPolicy' resources matching the request.
	Results              []ApplianceAutoRmaPolicy `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceAutoRmaPolicyList ApplianceAutoRmaPolicyList

// NewApplianceAutoRmaPolicyList instantiates a new ApplianceAutoRmaPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAutoRmaPolicyList(objectType string) *ApplianceAutoRmaPolicyList {
	this := ApplianceAutoRmaPolicyList{}
	this.ObjectType = objectType
	return &this
}

// NewApplianceAutoRmaPolicyListWithDefaults instantiates a new ApplianceAutoRmaPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAutoRmaPolicyListWithDefaults() *ApplianceAutoRmaPolicyList {
	this := ApplianceAutoRmaPolicyList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ApplianceAutoRmaPolicyList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAutoRmaPolicyList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ApplianceAutoRmaPolicyList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ApplianceAutoRmaPolicyList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceAutoRmaPolicyList) GetResults() []ApplianceAutoRmaPolicy {
	if o == nil {
		var ret []ApplianceAutoRmaPolicy
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceAutoRmaPolicyList) GetResultsOk() (*[]ApplianceAutoRmaPolicy, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApplianceAutoRmaPolicyList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ApplianceAutoRmaPolicy and assigns it to the Results field.
func (o *ApplianceAutoRmaPolicyList) SetResults(v []ApplianceAutoRmaPolicy) {
	o.Results = v
}

func (o ApplianceAutoRmaPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceAutoRmaPolicyList) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceAutoRmaPolicyListWithoutEmbeddedStruct struct {
		// The total number of 'appliance.AutoRmaPolicy' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'appliance.AutoRmaPolicy' resources matching the request.
		Results []ApplianceAutoRmaPolicy `json:"Results,omitempty"`
	}

	varApplianceAutoRmaPolicyListWithoutEmbeddedStruct := ApplianceAutoRmaPolicyListWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceAutoRmaPolicyListWithoutEmbeddedStruct)
	if err == nil {
		varApplianceAutoRmaPolicyList := _ApplianceAutoRmaPolicyList{}
		varApplianceAutoRmaPolicyList.Count = varApplianceAutoRmaPolicyListWithoutEmbeddedStruct.Count
		varApplianceAutoRmaPolicyList.Results = varApplianceAutoRmaPolicyListWithoutEmbeddedStruct.Results
		*o = ApplianceAutoRmaPolicyList(varApplianceAutoRmaPolicyList)
	} else {
		return err
	}

	varApplianceAutoRmaPolicyList := _ApplianceAutoRmaPolicyList{}

	err = json.Unmarshal(bytes, &varApplianceAutoRmaPolicyList)
	if err == nil {
		o.MoBaseResponse = varApplianceAutoRmaPolicyList.MoBaseResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")

		// remove fields from embedded structs
		reflectMoBaseResponse := reflect.ValueOf(o.MoBaseResponse)
		for i := 0; i < reflectMoBaseResponse.Type().NumField(); i++ {
			t := reflectMoBaseResponse.Type().Field(i)

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

type NullableApplianceAutoRmaPolicyList struct {
	value *ApplianceAutoRmaPolicyList
	isSet bool
}

func (v NullableApplianceAutoRmaPolicyList) Get() *ApplianceAutoRmaPolicyList {
	return v.value
}

func (v *NullableApplianceAutoRmaPolicyList) Set(val *ApplianceAutoRmaPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAutoRmaPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAutoRmaPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAutoRmaPolicyList(val *ApplianceAutoRmaPolicyList) *NullableApplianceAutoRmaPolicyList {
	return &NullableApplianceAutoRmaPolicyList{value: val, isSet: true}
}

func (v NullableApplianceAutoRmaPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAutoRmaPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
