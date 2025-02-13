/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ConvergedinfraPod A pod is unit of deployment of converged infrastructure. Contains inventory information related to the health, HCL, storage, nodes, etc. of the pod.
type ConvergedinfraPod struct {
	ConvergedinfraBasePod
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The deployment type for this solution pod. * `FlexPodInfra` - The deployment type for a pod is of Infrastructure. * `FlexPodNG` - The deployment type for a pod is of Nextgen type.
	DeploymentType       *string                               `json:"DeploymentType,omitempty"`
	Summary              *ConvergedinfraPodSummary             `json:"Summary,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	PodResourceGroup     *ResourceGroupRelationship            `json:"PodResourceGroup,omitempty"`
	SolutionInstance     *WorkflowSolutionInstanceRelationship `json:"SolutionInstance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraPod ConvergedinfraPod

// NewConvergedinfraPod instantiates a new ConvergedinfraPod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraPod(classId string, objectType string) *ConvergedinfraPod {
	this := ConvergedinfraPod{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "FlexPod"
	this.Type = &type_
	return &this
}

// NewConvergedinfraPodWithDefaults instantiates a new ConvergedinfraPod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraPodWithDefaults() *ConvergedinfraPod {
	this := ConvergedinfraPod{}
	var classId string = "convergedinfra.Pod"
	this.ClassId = classId
	var objectType string = "convergedinfra.Pod"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraPod) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPod) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraPod) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraPod) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPod) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraPod) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *ConvergedinfraPod) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPod) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *ConvergedinfraPod) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *ConvergedinfraPod) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ConvergedinfraPod) GetSummary() ConvergedinfraPodSummary {
	if o == nil || o.Summary == nil {
		var ret ConvergedinfraPodSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPod) GetSummaryOk() (*ConvergedinfraPodSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ConvergedinfraPod) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given ConvergedinfraPodSummary and assigns it to the Summary field.
func (o *ConvergedinfraPod) SetSummary(v ConvergedinfraPodSummary) {
	o.Summary = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ConvergedinfraPod) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPod) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ConvergedinfraPod) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ConvergedinfraPod) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetPodResourceGroup returns the PodResourceGroup field value if set, zero value otherwise.
func (o *ConvergedinfraPod) GetPodResourceGroup() ResourceGroupRelationship {
	if o == nil || o.PodResourceGroup == nil {
		var ret ResourceGroupRelationship
		return ret
	}
	return *o.PodResourceGroup
}

// GetPodResourceGroupOk returns a tuple with the PodResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPod) GetPodResourceGroupOk() (*ResourceGroupRelationship, bool) {
	if o == nil || o.PodResourceGroup == nil {
		return nil, false
	}
	return o.PodResourceGroup, true
}

// HasPodResourceGroup returns a boolean if a field has been set.
func (o *ConvergedinfraPod) HasPodResourceGroup() bool {
	if o != nil && o.PodResourceGroup != nil {
		return true
	}

	return false
}

// SetPodResourceGroup gets a reference to the given ResourceGroupRelationship and assigns it to the PodResourceGroup field.
func (o *ConvergedinfraPod) SetPodResourceGroup(v ResourceGroupRelationship) {
	o.PodResourceGroup = &v
}

// GetSolutionInstance returns the SolutionInstance field value if set, zero value otherwise.
func (o *ConvergedinfraPod) GetSolutionInstance() WorkflowSolutionInstanceRelationship {
	if o == nil || o.SolutionInstance == nil {
		var ret WorkflowSolutionInstanceRelationship
		return ret
	}
	return *o.SolutionInstance
}

// GetSolutionInstanceOk returns a tuple with the SolutionInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPod) GetSolutionInstanceOk() (*WorkflowSolutionInstanceRelationship, bool) {
	if o == nil || o.SolutionInstance == nil {
		return nil, false
	}
	return o.SolutionInstance, true
}

// HasSolutionInstance returns a boolean if a field has been set.
func (o *ConvergedinfraPod) HasSolutionInstance() bool {
	if o != nil && o.SolutionInstance != nil {
		return true
	}

	return false
}

// SetSolutionInstance gets a reference to the given WorkflowSolutionInstanceRelationship and assigns it to the SolutionInstance field.
func (o *ConvergedinfraPod) SetSolutionInstance(v WorkflowSolutionInstanceRelationship) {
	o.SolutionInstance = &v
}

func (o ConvergedinfraPod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConvergedinfraBasePod, errConvergedinfraBasePod := json.Marshal(o.ConvergedinfraBasePod)
	if errConvergedinfraBasePod != nil {
		return []byte{}, errConvergedinfraBasePod
	}
	errConvergedinfraBasePod = json.Unmarshal([]byte(serializedConvergedinfraBasePod), &toSerialize)
	if errConvergedinfraBasePod != nil {
		return []byte{}, errConvergedinfraBasePod
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.Summary != nil {
		toSerialize["Summary"] = o.Summary
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.PodResourceGroup != nil {
		toSerialize["PodResourceGroup"] = o.PodResourceGroup
	}
	if o.SolutionInstance != nil {
		toSerialize["SolutionInstance"] = o.SolutionInstance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraPod) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraPodWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The deployment type for this solution pod. * `FlexPodInfra` - The deployment type for a pod is of Infrastructure. * `FlexPodNG` - The deployment type for a pod is of Nextgen type.
		DeploymentType   *string                               `json:"DeploymentType,omitempty"`
		Summary          *ConvergedinfraPodSummary             `json:"Summary,omitempty"`
		Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		PodResourceGroup *ResourceGroupRelationship            `json:"PodResourceGroup,omitempty"`
		SolutionInstance *WorkflowSolutionInstanceRelationship `json:"SolutionInstance,omitempty"`
	}

	varConvergedinfraPodWithoutEmbeddedStruct := ConvergedinfraPodWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraPodWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraPod := _ConvergedinfraPod{}
		varConvergedinfraPod.ClassId = varConvergedinfraPodWithoutEmbeddedStruct.ClassId
		varConvergedinfraPod.ObjectType = varConvergedinfraPodWithoutEmbeddedStruct.ObjectType
		varConvergedinfraPod.DeploymentType = varConvergedinfraPodWithoutEmbeddedStruct.DeploymentType
		varConvergedinfraPod.Summary = varConvergedinfraPodWithoutEmbeddedStruct.Summary
		varConvergedinfraPod.Organization = varConvergedinfraPodWithoutEmbeddedStruct.Organization
		varConvergedinfraPod.PodResourceGroup = varConvergedinfraPodWithoutEmbeddedStruct.PodResourceGroup
		varConvergedinfraPod.SolutionInstance = varConvergedinfraPodWithoutEmbeddedStruct.SolutionInstance
		*o = ConvergedinfraPod(varConvergedinfraPod)
	} else {
		return err
	}

	varConvergedinfraPod := _ConvergedinfraPod{}

	err = json.Unmarshal(bytes, &varConvergedinfraPod)
	if err == nil {
		o.ConvergedinfraBasePod = varConvergedinfraPod.ConvergedinfraBasePod
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "Summary")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "PodResourceGroup")
		delete(additionalProperties, "SolutionInstance")

		// remove fields from embedded structs
		reflectConvergedinfraBasePod := reflect.ValueOf(o.ConvergedinfraBasePod)
		for i := 0; i < reflectConvergedinfraBasePod.Type().NumField(); i++ {
			t := reflectConvergedinfraBasePod.Type().Field(i)

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

type NullableConvergedinfraPod struct {
	value *ConvergedinfraPod
	isSet bool
}

func (v NullableConvergedinfraPod) Get() *ConvergedinfraPod {
	return v.value
}

func (v *NullableConvergedinfraPod) Set(val *ConvergedinfraPod) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraPod) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraPod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraPod(val *ConvergedinfraPod) *NullableConvergedinfraPod {
	return &NullableConvergedinfraPod{value: val, isSet: true}
}

func (v NullableConvergedinfraPod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraPod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
