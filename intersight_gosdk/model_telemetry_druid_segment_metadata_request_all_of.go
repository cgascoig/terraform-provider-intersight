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

// TelemetryDruidSegmentMetadataRequestAllOf struct for TelemetryDruidSegmentMetadataRequestAllOf
type TelemetryDruidSegmentMetadataRequestAllOf struct {
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. If an interval is not specified, the query will use a default interval that spans a configurable period before the end time of the most recent segment.
	Intervals *[]string `json:"intervals,omitempty"`
	// A JSON Object representing what columns should be included in the result. Defaults to \"all\".
	ToInclude *map[string]interface{} `json:"toInclude,omitempty"`
	// Merge all individual segment metadata results into a single result.
	Merge   *bool                       `json:"merge,omitempty"`
	Context *TelemetryDruidQueryContext `json:"context,omitempty"`
	// A list of Strings specifying what column properties (e.g. cardinality, size) should be calculated and returned in the result. Defaults to [\"cardinality\", \"interval\", \"minmax\"], but can be overridden with using the segment metadata query config. * cardinality - in the result will return the estimated floor of cardinality for each column. Only relevant for dimension columns. * minmax - Estimated min/max values for each column. Only relevant for dimension columns. * size - in the result will contain the estimated total segment byte size as if the data were stored in text format. * intervals - in the result will contain the list of intervals associated with the queried segments. * timestampSpec - in the result will contain timestampSpec of data stored in segments. This can be null if timestampSpec of segments was unknown or unmergeable (if merging is enabled). * queryGranularity - in the result will contain query granularity of data stored in segments. This can be null if query granularity of segments was unknown or unmergeable (if merging is enabled). * aggregators - in the result will contain the list of aggregators usable for querying metric columns. This may be null if the aggregators are unknown or unmergeable (if merging is enabled). Merging can be strict or lenient. The form of the result is a map of column name to aggregator. * rollup - in the result is true/false/null. When merging is enabled, if some are rollup, others are not, result is null.
	AnalysisTypes *[]string `json:"analysisTypes,omitempty"`
	// If true, and if the \"aggregators\" analysisType is enabled, aggregators will be merged leniently.
	LenientAggregatorMerge *bool `json:"lenientAggregatorMerge,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _TelemetryDruidSegmentMetadataRequestAllOf TelemetryDruidSegmentMetadataRequestAllOf

// NewTelemetryDruidSegmentMetadataRequestAllOf instantiates a new TelemetryDruidSegmentMetadataRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidSegmentMetadataRequestAllOf(dataSource TelemetryDruidDataSource) *TelemetryDruidSegmentMetadataRequestAllOf {
	this := TelemetryDruidSegmentMetadataRequestAllOf{}
	this.DataSource = dataSource
	return &this
}

// NewTelemetryDruidSegmentMetadataRequestAllOfWithDefaults instantiates a new TelemetryDruidSegmentMetadataRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidSegmentMetadataRequestAllOfWithDefaults() *TelemetryDruidSegmentMetadataRequestAllOf {
	this := TelemetryDruidSegmentMetadataRequestAllOf{}
	return &this
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetIntervals returns the Intervals field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetIntervals() []string {
	if o == nil || o.Intervals == nil {
		var ret []string
		return ret
	}
	return *o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetIntervalsOk() (*[]string, bool) {
	if o == nil || o.Intervals == nil {
		return nil, false
	}
	return o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasIntervals() bool {
	if o != nil && o.Intervals != nil {
		return true
	}

	return false
}

// SetIntervals gets a reference to the given []string and assigns it to the Intervals field.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetIntervals(v []string) {
	o.Intervals = &v
}

// GetToInclude returns the ToInclude field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetToInclude() map[string]interface{} {
	if o == nil || o.ToInclude == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ToInclude
}

// GetToIncludeOk returns a tuple with the ToInclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetToIncludeOk() (*map[string]interface{}, bool) {
	if o == nil || o.ToInclude == nil {
		return nil, false
	}
	return o.ToInclude, true
}

// HasToInclude returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasToInclude() bool {
	if o != nil && o.ToInclude != nil {
		return true
	}

	return false
}

// SetToInclude gets a reference to the given map[string]interface{} and assigns it to the ToInclude field.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetToInclude(v map[string]interface{}) {
	o.ToInclude = &v
}

// GetMerge returns the Merge field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetMerge() bool {
	if o == nil || o.Merge == nil {
		var ret bool
		return ret
	}
	return *o.Merge
}

// GetMergeOk returns a tuple with the Merge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetMergeOk() (*bool, bool) {
	if o == nil || o.Merge == nil {
		return nil, false
	}
	return o.Merge, true
}

// HasMerge returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasMerge() bool {
	if o != nil && o.Merge != nil {
		return true
	}

	return false
}

// SetMerge gets a reference to the given bool and assigns it to the Merge field.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetMerge(v bool) {
	o.Merge = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

// GetAnalysisTypes returns the AnalysisTypes field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetAnalysisTypes() []string {
	if o == nil || o.AnalysisTypes == nil {
		var ret []string
		return ret
	}
	return *o.AnalysisTypes
}

// GetAnalysisTypesOk returns a tuple with the AnalysisTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetAnalysisTypesOk() (*[]string, bool) {
	if o == nil || o.AnalysisTypes == nil {
		return nil, false
	}
	return o.AnalysisTypes, true
}

// HasAnalysisTypes returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasAnalysisTypes() bool {
	if o != nil && o.AnalysisTypes != nil {
		return true
	}

	return false
}

// SetAnalysisTypes gets a reference to the given []string and assigns it to the AnalysisTypes field.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetAnalysisTypes(v []string) {
	o.AnalysisTypes = &v
}

// GetLenientAggregatorMerge returns the LenientAggregatorMerge field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetLenientAggregatorMerge() bool {
	if o == nil || o.LenientAggregatorMerge == nil {
		var ret bool
		return ret
	}
	return *o.LenientAggregatorMerge
}

// GetLenientAggregatorMergeOk returns a tuple with the LenientAggregatorMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetLenientAggregatorMergeOk() (*bool, bool) {
	if o == nil || o.LenientAggregatorMerge == nil {
		return nil, false
	}
	return o.LenientAggregatorMerge, true
}

// HasLenientAggregatorMerge returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasLenientAggregatorMerge() bool {
	if o != nil && o.LenientAggregatorMerge != nil {
		return true
	}

	return false
}

// SetLenientAggregatorMerge gets a reference to the given bool and assigns it to the LenientAggregatorMerge field.
func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetLenientAggregatorMerge(v bool) {
	o.LenientAggregatorMerge = &v
}

func (o TelemetryDruidSegmentMetadataRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if o.Intervals != nil {
		toSerialize["intervals"] = o.Intervals
	}
	if o.ToInclude != nil {
		toSerialize["toInclude"] = o.ToInclude
	}
	if o.Merge != nil {
		toSerialize["merge"] = o.Merge
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.AnalysisTypes != nil {
		toSerialize["analysisTypes"] = o.AnalysisTypes
	}
	if o.LenientAggregatorMerge != nil {
		toSerialize["lenientAggregatorMerge"] = o.LenientAggregatorMerge
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidSegmentMetadataRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidSegmentMetadataRequestAllOf := _TelemetryDruidSegmentMetadataRequestAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidSegmentMetadataRequestAllOf); err == nil {
		*o = TelemetryDruidSegmentMetadataRequestAllOf(varTelemetryDruidSegmentMetadataRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "toInclude")
		delete(additionalProperties, "merge")
		delete(additionalProperties, "context")
		delete(additionalProperties, "analysisTypes")
		delete(additionalProperties, "lenientAggregatorMerge")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidSegmentMetadataRequestAllOf struct {
	value *TelemetryDruidSegmentMetadataRequestAllOf
	isSet bool
}

func (v NullableTelemetryDruidSegmentMetadataRequestAllOf) Get() *TelemetryDruidSegmentMetadataRequestAllOf {
	return v.value
}

func (v *NullableTelemetryDruidSegmentMetadataRequestAllOf) Set(val *TelemetryDruidSegmentMetadataRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidSegmentMetadataRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidSegmentMetadataRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidSegmentMetadataRequestAllOf(val *TelemetryDruidSegmentMetadataRequestAllOf) *NullableTelemetryDruidSegmentMetadataRequestAllOf {
	return &NullableTelemetryDruidSegmentMetadataRequestAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidSegmentMetadataRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidSegmentMetadataRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
