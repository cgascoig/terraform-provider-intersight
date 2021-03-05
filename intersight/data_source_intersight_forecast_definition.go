package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceForecastDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceForecastDefinitionRead,
		Schema: map[string]*schema.Schema{
			"alert_threshold_in_percentage": {
				Description: "Threshold above which user needs to be indicated through alarm/alert.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"data_source": {
				Description: "Data source from where we get the data for the metrics to compute regression model. For example Druid.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"metric_name": {
				Description: "Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"min_num_of_days_of_data": {
				Description: "Minimum number of days of data required for computing forecast model.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_of_days_of_historical_data": {
				Description: "Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"platform_type": {
				Description: "The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"alert_threshold_in_percentage": {
						Description: "Threshold above which user needs to be indicated through alarm/alert.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"catalog": {
						Description: "A reference to a forecastCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"data_source": {
						Description: "Data source from where we get the data for the metrics to compute regression model. For example Druid.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"metric_name": {
						Description: "Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"min_num_of_days_of_data": {
						Description: "Minimum number of days of data required for computing forecast model.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"num_of_days_of_historical_data": {
						Description: "Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"platform_type": {
						Description: "The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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

func dataSourceForecastDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ForecastDefinition{}
	if v, ok := d.GetOk("alert_threshold_in_percentage"); ok {
		x := int64(v.(int))
		o.SetAlertThresholdInPercentage(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("data_source"); ok {
		x := (v.(string))
		o.SetDataSource(x)
	}
	if v, ok := d.GetOk("metric_name"); ok {
		x := (v.(string))
		o.SetMetricName(x)
	}
	if v, ok := d.GetOk("min_num_of_days_of_data"); ok {
		x := int64(v.(int))
		o.SetMinNumOfDaysOfData(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("num_of_days_of_historical_data"); ok {
		x := int64(v.(int))
		o.SetNumOfDaysOfHistoricalData(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ForecastDefinition object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ForecastApi.GetForecastDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ForecastDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ForecastDefinitionList.GetCount()
	var i int32
	var forecastDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ForecastApi.GetForecastDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ForecastDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ForecastDefinitionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ForecastDefinition data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["alert_threshold_in_percentage"] = (s.GetAlertThresholdInPercentage())

				temp["catalog"] = flattenMapForecastCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())
				temp["data_source"] = (s.GetDataSource())
				temp["metric_name"] = (s.GetMetricName())
				temp["min_num_of_days_of_data"] = (s.GetMinNumOfDaysOfData())
				temp["moid"] = (s.GetMoid())
				temp["num_of_days_of_historical_data"] = (s.GetNumOfDaysOfHistoricalData())
				temp["object_type"] = (s.GetObjectType())
				temp["platform_type"] = (s.GetPlatformType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				forecastDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(forecastDefinitionResults))
	if err := d.Set("results", forecastDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(forecastDefinitionResults[0]["moid"].(string))
	return de
}
