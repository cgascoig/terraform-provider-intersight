package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplianceDiagSetting() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceDiagSettingRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"message": {
				Description: "Status message of the password change operation.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Description: "Password of the Intersight Appliance's OS diagnostic user account.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceApplianceDiagSetting().Schema},
				Computed: true,
			}},
	}
}

func dataSourceApplianceDiagSettingRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceDiagSetting{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}
	if v, ok := d.GetOk("message"); ok {
		x := (v.(string))
		o.SetMessage(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.SetPassword(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceDiagSetting object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceDiagSettingList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ApplianceDiagSetting: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ApplianceDiagSettingList.GetCount()
	var i int32
	var applianceDiagSettingResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceDiagSettingList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceDiagSetting: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ApplianceDiagSettingList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ApplianceDiagSetting data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["is_password_set"] = (s.GetIsPasswordSet())
				temp["message"] = (s.GetMessage())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				applianceDiagSettingResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceDiagSettingResults))
	if err := d.Set("results", applianceDiagSettingResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceDiagSettingResults[0]["moid"].(string))
	return de
}
