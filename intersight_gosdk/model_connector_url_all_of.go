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
)

// ConnectorUrlAllOf Definition of the list of properties defined in 'connector.Url', excluding properties defined in parent classes.
type ConnectorUrlAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flag to append a query to the url even if rawQuery is empty.
	ForceQuery *bool `json:"ForceQuery,omitempty"`
	// The fragment identifier component of a URI allows indirect identification of a secondary resource by reference to a primary resource and additional identifying information. The identified secondary resource may be some portion or subset of the primary resource, some view on representations of the primary resource, or some other resource defined or described by those representations. A fragment identifier component is indicated by the presence of a number sign (\"#\") character and terminated by the end of the URI.
	Fragment *string `json:"Fragment,omitempty"`
	// The host name identifies the host that holds the resource. The host can be an IP or a hostname that is resolvable by the dns server configured on the platform.
	Host *string `json:"Host,omitempty"`
	// A URI is opaque if, and only if, it is absolute and its scheme-specific part does not begin with a slash character ('/'). An opaque URI has a scheme, a scheme-specific part, and possibly a fragment; all other components are undefined.
	Opaque *string `json:"Opaque,omitempty"`
	// The path identifies the specific resource in the host that the web client wants to access. Value is the decoded form of the path. e.g. '/foo/bar'.
	Path *string `json:"Path,omitempty"`
	// The URI encoded form of the path property. e.g. '%2Fapi%2Fv1%2F'.
	RawPath *string `json:"RawPath,omitempty"`
	// The query component, as defined in RFC 3986, contains non-hierarchical data that, along with data in the path component, serves to identify a resource within the scope of the URI's scheme and naming authority (if any). The query component is indicated by the first question mark character and terminated by a number sign character or by the end of the URI. The rawQuery contains the URIs encoded query component, excluding the ? character.
	RawQuery *string `json:"RawQuery,omitempty"`
	// The scheme identifies the protocol to be used to access the resource on the Internet. It can be HTTP (without SSL) or HTTPS (with SSL).
	Scheme               *string `json:"Scheme,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorUrlAllOf ConnectorUrlAllOf

// NewConnectorUrlAllOf instantiates a new ConnectorUrlAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorUrlAllOf(classId string, objectType string) *ConnectorUrlAllOf {
	this := ConnectorUrlAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorUrlAllOfWithDefaults instantiates a new ConnectorUrlAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorUrlAllOfWithDefaults() *ConnectorUrlAllOf {
	this := ConnectorUrlAllOf{}
	var classId string = "connector.Url"
	this.ClassId = classId
	var objectType string = "connector.Url"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorUrlAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorUrlAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorUrlAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorUrlAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetForceQuery returns the ForceQuery field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetForceQuery() bool {
	if o == nil || o.ForceQuery == nil {
		var ret bool
		return ret
	}
	return *o.ForceQuery
}

// GetForceQueryOk returns a tuple with the ForceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetForceQueryOk() (*bool, bool) {
	if o == nil || o.ForceQuery == nil {
		return nil, false
	}
	return o.ForceQuery, true
}

// HasForceQuery returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasForceQuery() bool {
	if o != nil && o.ForceQuery != nil {
		return true
	}

	return false
}

// SetForceQuery gets a reference to the given bool and assigns it to the ForceQuery field.
func (o *ConnectorUrlAllOf) SetForceQuery(v bool) {
	o.ForceQuery = &v
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetFragment() string {
	if o == nil || o.Fragment == nil {
		var ret string
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetFragmentOk() (*string, bool) {
	if o == nil || o.Fragment == nil {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasFragment() bool {
	if o != nil && o.Fragment != nil {
		return true
	}

	return false
}

// SetFragment gets a reference to the given string and assigns it to the Fragment field.
func (o *ConnectorUrlAllOf) SetFragment(v string) {
	o.Fragment = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ConnectorUrlAllOf) SetHost(v string) {
	o.Host = &v
}

// GetOpaque returns the Opaque field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetOpaque() string {
	if o == nil || o.Opaque == nil {
		var ret string
		return ret
	}
	return *o.Opaque
}

// GetOpaqueOk returns a tuple with the Opaque field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetOpaqueOk() (*string, bool) {
	if o == nil || o.Opaque == nil {
		return nil, false
	}
	return o.Opaque, true
}

// HasOpaque returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasOpaque() bool {
	if o != nil && o.Opaque != nil {
		return true
	}

	return false
}

// SetOpaque gets a reference to the given string and assigns it to the Opaque field.
func (o *ConnectorUrlAllOf) SetOpaque(v string) {
	o.Opaque = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ConnectorUrlAllOf) SetPath(v string) {
	o.Path = &v
}

// GetRawPath returns the RawPath field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetRawPath() string {
	if o == nil || o.RawPath == nil {
		var ret string
		return ret
	}
	return *o.RawPath
}

// GetRawPathOk returns a tuple with the RawPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetRawPathOk() (*string, bool) {
	if o == nil || o.RawPath == nil {
		return nil, false
	}
	return o.RawPath, true
}

// HasRawPath returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasRawPath() bool {
	if o != nil && o.RawPath != nil {
		return true
	}

	return false
}

// SetRawPath gets a reference to the given string and assigns it to the RawPath field.
func (o *ConnectorUrlAllOf) SetRawPath(v string) {
	o.RawPath = &v
}

// GetRawQuery returns the RawQuery field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetRawQuery() string {
	if o == nil || o.RawQuery == nil {
		var ret string
		return ret
	}
	return *o.RawQuery
}

// GetRawQueryOk returns a tuple with the RawQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetRawQueryOk() (*string, bool) {
	if o == nil || o.RawQuery == nil {
		return nil, false
	}
	return o.RawQuery, true
}

// HasRawQuery returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasRawQuery() bool {
	if o != nil && o.RawQuery != nil {
		return true
	}

	return false
}

// SetRawQuery gets a reference to the given string and assigns it to the RawQuery field.
func (o *ConnectorUrlAllOf) SetRawQuery(v string) {
	o.RawQuery = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ConnectorUrlAllOf) GetScheme() string {
	if o == nil || o.Scheme == nil {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorUrlAllOf) GetSchemeOk() (*string, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ConnectorUrlAllOf) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *ConnectorUrlAllOf) SetScheme(v string) {
	o.Scheme = &v
}

func (o ConnectorUrlAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ForceQuery != nil {
		toSerialize["ForceQuery"] = o.ForceQuery
	}
	if o.Fragment != nil {
		toSerialize["Fragment"] = o.Fragment
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.Opaque != nil {
		toSerialize["Opaque"] = o.Opaque
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.RawPath != nil {
		toSerialize["RawPath"] = o.RawPath
	}
	if o.RawQuery != nil {
		toSerialize["RawQuery"] = o.RawQuery
	}
	if o.Scheme != nil {
		toSerialize["Scheme"] = o.Scheme
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorUrlAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorUrlAllOf := _ConnectorUrlAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorUrlAllOf); err == nil {
		*o = ConnectorUrlAllOf(varConnectorUrlAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ForceQuery")
		delete(additionalProperties, "Fragment")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "Opaque")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "RawPath")
		delete(additionalProperties, "RawQuery")
		delete(additionalProperties, "Scheme")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorUrlAllOf struct {
	value *ConnectorUrlAllOf
	isSet bool
}

func (v NullableConnectorUrlAllOf) Get() *ConnectorUrlAllOf {
	return v.value
}

func (v *NullableConnectorUrlAllOf) Set(val *ConnectorUrlAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorUrlAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorUrlAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorUrlAllOf(val *ConnectorUrlAllOf) *NullableConnectorUrlAllOf {
	return &NullableConnectorUrlAllOf{value: val, isSet: true}
}

func (v NullableConnectorUrlAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorUrlAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
