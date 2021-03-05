package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesAciCniTenantClusterAllocation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesAciCniTenantClusterAllocationRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"node_svc_ip_subnet": {
				Description: "CIDR allocated for ACI node service IPs in this tenant cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pod_ip_subnet": {
				Description: "CIDR allocated for pod IPs in this tenant cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_end": {
				Description: "End of VLAN range allocated to this tenant cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_start": {
				Description: "Start of VLAN range allocated to this tenant cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesAciCniTenantClusterAllocation().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesAciCniTenantClusterAllocationRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAciCniTenantClusterAllocation{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("node_svc_ip_subnet"); ok {
		x := (v.(string))
		o.SetNodeSvcIpSubnet(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pod_ip_subnet"); ok {
		x := (v.(string))
		o.SetPodIpSubnet(x)
	}
	if v, ok := d.GetOk("vlan_end"); ok {
		x := (v.(string))
		o.SetVlanEnd(x)
	}
	if v, ok := d.GetOk("vlan_start"); ok {
		x := (v.(string))
		o.SetVlanStart(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesAciCniTenantClusterAllocation object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAciCniTenantClusterAllocationList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesAciCniTenantClusterAllocation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesAciCniTenantClusterAllocationList.GetCount()
	var i int32
	var kubernetesAciCniTenantClusterAllocationResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAciCniTenantClusterAllocationList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesAciCniTenantClusterAllocation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesAciCniTenantClusterAllocationList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesAciCniTenantClusterAllocation data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["node_svc_ip_subnet"] = (s.GetNodeSvcIpSubnet())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["pod_ip_subnet"] = (s.GetPodIpSubnet())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vlan_end"] = (s.GetVlanEnd())
				temp["vlan_start"] = (s.GetVlanStart())
				kubernetesAciCniTenantClusterAllocationResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesAciCniTenantClusterAllocationResults))
	if err := d.Set("results", kubernetesAciCniTenantClusterAllocationResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesAciCniTenantClusterAllocationResults[0]["moid"].(string))
	return de
}
