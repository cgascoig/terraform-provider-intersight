package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesDaemonSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesDaemonSetRead,
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
			"name": {
				Description: "Name of the referenced kubernetes resource.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "UUID of the referenced kubernetes resource. It is generated by the kubernetes cluster itself.",
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
					"metadata": {
						Description: "Meta data of the kubernetes resource.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
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
								"creation_timestamp": {
									Description: "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Read-only.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"namespace": {
									Description: "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Read-only.",
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
								"resource_version": {
									Description: "Specific resourceVersion to which this reference is made, if any.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"uuid": {
									Description: "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations. Populated by the system. Read-only.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "Name of the referenced kubernetes resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
					"status": {
						Description: "DaemonSetstatus represents the current status of the corresponding DaemonSetStatus in a Kubernetes cluster. For detail, please refer to kubernetes.DaemonSetStatus.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
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
								"current_number_scheduled": {
									Description: "The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"desired_number_scheduled": {
									Description: "The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod).",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"number_available": {
									Description: "The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds).",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"number_misscheduled": {
									Description: "The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"number_ready": {
									Description: "The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"observed_generation": {
									Description: "The most recent generation observed by the daemon set controller.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"updated_number_scheduled": {
									Description: "The total number of nodes that are running updated daemon pod.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
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
					"uuid": {
						Description: "UUID of the referenced kubernetes resource. It is generated by the kubernetes cluster itself.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesDaemonSetRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesDaemonSet{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
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
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesDaemonSet object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesDaemonSetList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesDaemonSet: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesDaemonSetList.GetCount()
	var i int32
	var kubernetesDaemonSetResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesDaemonSetList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesDaemonSet: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesDaemonSetList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesDaemonSet data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["metadata"] = flattenMapKubernetesObjectMeta(s.GetMetadata(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)

				temp["status"] = flattenMapKubernetesDaemonSetStatus(s.GetStatus(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["uuid"] = (s.GetUuid())
				kubernetesDaemonSetResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesDaemonSetResults))
	if err := d.Set("results", kubernetesDaemonSetResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesDaemonSetResults[0]["moid"].(string))
	return de
}
