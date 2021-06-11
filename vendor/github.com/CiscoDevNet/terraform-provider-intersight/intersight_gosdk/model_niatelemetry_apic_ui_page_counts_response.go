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
	"fmt"
)

// NiatelemetryApicUiPageCountsResponse - The response body of a HTTP GET request for the 'niatelemetry.ApicUiPageCounts' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.ApicUiPageCounts' resources.
type NiatelemetryApicUiPageCountsResponse struct {
	MoAggregateTransform             *MoAggregateTransform
	MoDocumentCount                  *MoDocumentCount
	MoTagSummary                     *MoTagSummary
	NiatelemetryApicUiPageCountsList *NiatelemetryApicUiPageCountsList
}

// MoAggregateTransformAsNiatelemetryApicUiPageCountsResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryApicUiPageCountsResponse
func MoAggregateTransformAsNiatelemetryApicUiPageCountsResponse(v *MoAggregateTransform) NiatelemetryApicUiPageCountsResponse {
	return NiatelemetryApicUiPageCountsResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsNiatelemetryApicUiPageCountsResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryApicUiPageCountsResponse
func MoDocumentCountAsNiatelemetryApicUiPageCountsResponse(v *MoDocumentCount) NiatelemetryApicUiPageCountsResponse {
	return NiatelemetryApicUiPageCountsResponse{MoDocumentCount: v}
}

// MoTagSummaryAsNiatelemetryApicUiPageCountsResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryApicUiPageCountsResponse
func MoTagSummaryAsNiatelemetryApicUiPageCountsResponse(v *MoTagSummary) NiatelemetryApicUiPageCountsResponse {
	return NiatelemetryApicUiPageCountsResponse{MoTagSummary: v}
}

// NiatelemetryApicUiPageCountsListAsNiatelemetryApicUiPageCountsResponse is a convenience function that returns NiatelemetryApicUiPageCountsList wrapped in NiatelemetryApicUiPageCountsResponse
func NiatelemetryApicUiPageCountsListAsNiatelemetryApicUiPageCountsResponse(v *NiatelemetryApicUiPageCountsList) NiatelemetryApicUiPageCountsResponse {
	return NiatelemetryApicUiPageCountsResponse{NiatelemetryApicUiPageCountsList: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryApicUiPageCountsResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicUiPageCountsResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicUiPageCountsResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicUiPageCountsResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.ApicUiPageCounts.List'
	if jsonDict["ObjectType"] == "niatelemetry.ApicUiPageCounts.List" {
		// try to unmarshal JSON data into NiatelemetryApicUiPageCountsList
		err = json.Unmarshal(data, &dst.NiatelemetryApicUiPageCountsList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryApicUiPageCountsList, return on the first match
		} else {
			dst.NiatelemetryApicUiPageCountsList = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicUiPageCountsResponse as NiatelemetryApicUiPageCountsList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryApicUiPageCountsResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryApicUiPageCountsList != nil {
		return json.Marshal(&src.NiatelemetryApicUiPageCountsList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryApicUiPageCountsResponse) GetActualInstance() interface{} {
	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	if obj.NiatelemetryApicUiPageCountsList != nil {
		return obj.NiatelemetryApicUiPageCountsList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryApicUiPageCountsResponse struct {
	value *NiatelemetryApicUiPageCountsResponse
	isSet bool
}

func (v NullableNiatelemetryApicUiPageCountsResponse) Get() *NiatelemetryApicUiPageCountsResponse {
	return v.value
}

func (v *NullableNiatelemetryApicUiPageCountsResponse) Set(val *NiatelemetryApicUiPageCountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicUiPageCountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicUiPageCountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicUiPageCountsResponse(val *NiatelemetryApicUiPageCountsResponse) *NullableNiatelemetryApicUiPageCountsResponse {
	return &NullableNiatelemetryApicUiPageCountsResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryApicUiPageCountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicUiPageCountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
