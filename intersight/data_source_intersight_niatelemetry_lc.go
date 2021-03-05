package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryLc() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryLcRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the line cards present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dn": {
				Description: "Dn value for the line cards present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hardware_version": {
				Description: "Hardware version of the line cards present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "Model of the line cards present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"node_id": {
				Description: "Node Id of the line card present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operational_state": {
				Description: "Opretaional state of the line cards present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"power_state": {
				Description: "Power state of the line cards present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"record_type": {
				Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"record_version": {
				Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"redundancy_state": {
				Description: "Redundancy state of the line cards present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial_number": {
				Description: "Serial number of the line card present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_name": {
				Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"description": {
						Description: "Description of the line cards present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"dn": {
						Description: "Dn value for the line cards present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hardware_version": {
						Description: "Hardware version of the line cards present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"model": {
						Description: "Model of the line cards present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"node_id": {
						Description: "Node Id of the line card present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"operational_state": {
						Description: "Opretaional state of the line cards present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"power_state": {
						Description: "Power state of the line cards present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"record_type": {
						Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"record_version": {
						Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"redundancy_state": {
						Description: "Redundancy state of the line cards present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"serial_number": {
						Description: "Serial number of the line card present.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"site_name": {
						Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiatelemetryLcRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryLc{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("hardware_version"); ok {
		x := (v.(string))
		o.SetHardwareVersion(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("node_id"); ok {
		x := (v.(string))
		o.SetNodeId(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("operational_state"); ok {
		x := (v.(string))
		o.SetOperationalState(x)
	}
	if v, ok := d.GetOk("power_state"); ok {
		x := (v.(string))
		o.SetPowerState(x)
	}
	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}
	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}
	if v, ok := d.GetOk("redundancy_state"); ok {
		x := (v.(string))
		o.SetRedundancyState(x)
	}
	if v, ok := d.GetOk("serial_number"); ok {
		x := (v.(string))
		o.SetSerialNumber(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryLc object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryLcList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryLc: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryLcList.GetCount()
	var i int32
	var niatelemetryLcResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryLcList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryLc: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryLcList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryLc data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["dn"] = (s.GetDn())
				temp["hardware_version"] = (s.GetHardwareVersion())
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["node_id"] = (s.GetNodeId())
				temp["object_type"] = (s.GetObjectType())
				temp["operational_state"] = (s.GetOperationalState())
				temp["power_state"] = (s.GetPowerState())
				temp["record_type"] = (s.GetRecordType())
				temp["record_version"] = (s.GetRecordVersion())
				temp["redundancy_state"] = (s.GetRedundancyState())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["serial_number"] = (s.GetSerialNumber())
				temp["site_name"] = (s.GetSiteName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				niatelemetryLcResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryLcResults))
	if err := d.Set("results", niatelemetryLcResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryLcResults[0]["moid"].(string))
	return de
}
