/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// UuidpoolBlockRelationship - A relationship to the 'uuidpool.Block' resource, or the expanded 'uuidpool.Block' resource, or the 'null' value.
type UuidpoolBlockRelationship struct {
	MoMoRef       *MoMoRef
	UuidpoolBlock *UuidpoolBlock
}

// MoMoRefAsUuidpoolBlockRelationship is a convenience function that returns MoMoRef wrapped in UuidpoolBlockRelationship
func MoMoRefAsUuidpoolBlockRelationship(v *MoMoRef) UuidpoolBlockRelationship {
	return UuidpoolBlockRelationship{MoMoRef: v}
}

// UuidpoolBlockAsUuidpoolBlockRelationship is a convenience function that returns UuidpoolBlock wrapped in UuidpoolBlockRelationship
func UuidpoolBlockAsUuidpoolBlockRelationship(v *UuidpoolBlock) UuidpoolBlockRelationship {
	return UuidpoolBlockRelationship{UuidpoolBlock: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UuidpoolBlockRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal UuidpoolBlockRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'uuidpool.Block'
	if jsonDict["ClassId"] == "uuidpool.Block" {
		// try to unmarshal JSON data into UuidpoolBlock
		err = json.Unmarshal(data, &dst.UuidpoolBlock)
		if err == nil {
			return nil // data stored in dst.UuidpoolBlock, return on the first match
		} else {
			dst.UuidpoolBlock = nil
			return fmt.Errorf("Failed to unmarshal UuidpoolBlockRelationship as UuidpoolBlock: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UuidpoolBlockRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.UuidpoolBlock != nil {
		return json.Marshal(&src.UuidpoolBlock)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UuidpoolBlockRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.UuidpoolBlock != nil {
		return obj.UuidpoolBlock
	}

	// all schemas are nil
	return nil
}

type NullableUuidpoolBlockRelationship struct {
	value *UuidpoolBlockRelationship
	isSet bool
}

func (v NullableUuidpoolBlockRelationship) Get() *UuidpoolBlockRelationship {
	return v.value
}

func (v *NullableUuidpoolBlockRelationship) Set(val *UuidpoolBlockRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolBlockRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolBlockRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolBlockRelationship(val *UuidpoolBlockRelationship) *NullableUuidpoolBlockRelationship {
	return &NullableUuidpoolBlockRelationship{value: val, isSet: true}
}

func (v NullableUuidpoolBlockRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolBlockRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
