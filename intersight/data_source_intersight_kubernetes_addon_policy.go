package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesAddonPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesAddonPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"system_managed": {
				Description: "To determine if Addon Policy is automatically managed by the system.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesAddonPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesAddonPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAddonPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("system_managed"); ok {
		x := (v.(bool))
		o.SetSystemManaged(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesAddonPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesAddonPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesAddonPolicyList.GetCount()
	var i int32
	var kubernetesAddonPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesAddonPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesAddonPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesAddonPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["addons"] = flattenListKubernetesAddonRelationship(s.GetAddons(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_profiles"] = flattenListKubernetesClusterProfileRelationship(s.GetClusterProfiles(), d)
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["system_managed"] = (s.GetSystemManaged())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				kubernetesAddonPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesAddonPolicyResults))
	if err := d.Set("results", kubernetesAddonPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesAddonPolicyResults[0]["moid"].(string))
	return de
}
