package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirmwareSasExpanderDescriptor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirmwareSasExpanderDescriptorRead,
		Schema: map[string]*schema.Schema{
			"brand_string": {
				Description: "The brand string of the endpoint for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Detailed information about the endpoint.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"label": {
				Description: "The label type for the component.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "The model of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"revision": {
				Description: "The revision for the component.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vendor": {
				Description: "The vendor of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "The firmware or software version of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFirmwareSasExpanderDescriptor().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFirmwareSasExpanderDescriptorRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FirmwareSasExpanderDescriptor{}
	if v, ok := d.GetOk("brand_string"); ok {
		x := (v.(string))
		o.SetBrandString(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FirmwareSasExpanderDescriptor object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareSasExpanderDescriptorList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FirmwareSasExpanderDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FirmwareSasExpanderDescriptorList.GetCount()
	var i int32
	var firmwareSasExpanderDescriptorResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareSasExpanderDescriptorList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FirmwareSasExpanderDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FirmwareSasExpanderDescriptorList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FirmwareSasExpanderDescriptor data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["brand_string"] = (s.GetBrandString())

				temp["capabilities"] = flattenListCapabilityCapabilityRelationship(s.GetCapabilities(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["label"] = (s.GetLabel())
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["revision"] = (s.GetRevision())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vendor"] = (s.GetVendor())
				temp["nr_version"] = (s.GetVersion())
				firmwareSasExpanderDescriptorResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(firmwareSasExpanderDescriptorResults))
	if err := d.Set("results", firmwareSasExpanderDescriptorResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(firmwareSasExpanderDescriptorResults[0]["moid"].(string))
	return de
}
