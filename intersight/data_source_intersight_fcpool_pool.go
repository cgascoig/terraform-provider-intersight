package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFcpoolPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFcpoolPoolRead,
		Schema: map[string]*schema.Schema{
			"assigned": {
				Description: "Number of IDs that are currently assigned.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"assignment_order": {
				Description: "Assignment order decides the order in which the next identifier is allocated.\n* `sequential` - Identifiers are assigned in a sequential order.\n* `default` - Assignment order is decided by the system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
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
			"pool_purpose": {
				Description: "Purpose of this WWN pool.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"size": {
				Description: "Total number of identifiers in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFcpoolPool().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFcpoolPoolRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FcpoolPool{}
	if v, ok := d.GetOk("assigned"); ok {
		x := int64(v.(int))
		o.SetAssigned(x)
	}
	if v, ok := d.GetOk("assignment_order"); ok {
		x := (v.(string))
		o.SetAssignmentOrder(x)
	}
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
	if v, ok := d.GetOk("pool_purpose"); ok {
		x := (v.(string))
		o.SetPoolPurpose(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FcpoolPool object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FcpoolApi.GetFcpoolPoolList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FcpoolPool: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FcpoolPoolList.GetCount()
	var i int32
	var fcpoolPoolResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FcpoolApi.GetFcpoolPoolList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FcpoolPool: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FcpoolPoolList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FcpoolPool data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["assigned"] = (s.GetAssigned())
				temp["assignment_order"] = (s.GetAssignmentOrder())

				temp["block_heads"] = flattenListFcpoolFcBlockRelationship(s.GetBlockHeads(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["id_blocks"] = flattenListFcpoolBlock(s.GetIdBlocks(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["pool_purpose"] = (s.GetPoolPurpose())
				temp["size"] = (s.GetSize())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				fcpoolPoolResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fcpoolPoolResults))
	if err := d.Set("results", fcpoolPoolResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fcpoolPoolResults[0]["moid"].(string))
	return de
}
