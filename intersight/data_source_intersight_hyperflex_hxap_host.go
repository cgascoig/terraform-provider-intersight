package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexHxapHost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHxapHostRead,
		Schema: map[string]*schema.Schema{
			"age": {
				Description: "Denotes age or life time of the Host in nano seconds.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_uuid": {
				Description: "The UUID of the cluster to which this Host belongs to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"failure_reason": {
				Description: "Reason of the failure when host is in failed state.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hw_power_state": {
				Description: "Is the host Powered-up or Powered-down.\n* `Unknown` - The entity's power state is unknown.\n* `PoweredOn` - The entity is powered on.\n* `PoweredOff` - The entity is powered down.\n* `StandBy` - The entity is in standby mode.\n* `Paused` - The entity is in pause state.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_type": {
				Description: "Identifies the broad type of the underlying hypervisor.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identity": {
				Description: "The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"internal_ip_address": {
				Description: "Internal IP Address of the Host.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"maintenance_mode": {
				Description: "Is this host in maintenance mode. Set to true or false.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"management_ip_address": {
				Description: "Management IP Address of the Host.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"master_role": {
				Description: "Is the role of this host is master in the cluster? true or false.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"model": {
				Description: "Commercial model information about this hardware.",
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
				Description: "Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "Serial number of this host (internally generated).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "Host health status, as reported by the hypervisor platform.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"up_time": {
				Description: "The uptime of the host, stored as Duration (from w3c).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"uuid": {
				Description: "Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vendor": {
				Description: "Commercial vendor details of this hardware.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "Product version of the Host.",
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
					"age": {
						Description: "Denotes age or life time of the Host in nano seconds.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster": {
						Description: "A reference to a hyperflexHxapCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"cluster_member": {
						Description: "A reference to a assetClusterMember resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"cluster_uuid": {
						Description: "The UUID of the cluster to which this Host belongs to.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cpu_info": {
						Description: "Details about the CPUs installed on this host are represented here.",
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
								"cores": {
									Description: "Number of cores per CPU, as reported by the manufacturer.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"description": {
									Description: "The vendor provided description of the CPU. For example, Intel Xeon E5-2640 v3 @ 2.60GHz.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"sockets": {
									Description: "Number of CPU sockets available.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"speed": {
									Description: "Speed of the CPUs in Hertz. For example, 2593749663.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"vendor": {
									Description: "Manufacturer of the CPU . For example, Intel.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"failure_reason": {
						Description: "Reason of the failure when host is in failed state.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hardware_info": {
						Description: "The hardware details of this host. It includes capacity, manufacturer, and model information.",
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
								"cpu_cores": {
									Description: "The number of cpu cores on this hardware platform.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"cpu_speed": {
									Description: "Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"memory_size": {
									Description: "The amount of memory allocated (bytes) to this hardware platform.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"hw_power_state": {
						Description: "Is the host Powered-up or Powered-down.\n* `Unknown` - The entity's power state is unknown.\n* `PoweredOn` - The entity is powered on.\n* `PoweredOff` - The entity is powered down.\n* `StandBy` - The entity is in standby mode.\n* `Paused` - The entity is in pause state.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hypervisor_type": {
						Description: "Identifies the broad type of the underlying hypervisor.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"identity": {
						Description: "The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"internal_ip_address": {
						Description: "Internal IP Address of the Host.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"maintenance_mode": {
						Description: "Is this host in maintenance mode. Set to true or false.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"management_ip_address": {
						Description: "Management IP Address of the Host.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"master_role": {
						Description: "Is the role of this host is master in the cluster? true or false.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"memory_capacity": {
						Description: "The memory capacity and usage information on this host.",
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
								"capacity": {
									Description: "The total memory capacity of the entity in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"free": {
									Description: "Free memory (bytes) that is unused and available for allocation, as a point-in-time snapshot. The available memory capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used memory capacity is also reported.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"used": {
									Description: "Memory (bytes) that has been already used up, as a point-in-time snapshot.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"model": {
						Description: "Commercial model information about this hardware.",
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
						Description: "Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"physical_server": {
						Description: "A reference to a computePhysical resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"processor_capacity": {
						Description: "The capacity and usage information for CPU power on this host.",
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
								"capacity": {
									Description: "Total capacity of the entity in MHz.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"free": {
									Description: "Free CPU capacity in MHz, as a point-in-time snapshot. The available CPU capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used CPU capacity is also reported.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"used": {
									Description: "Used CPU capacity of the entity in MHz, as a point-in-time snapshot.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"product_info": {
						Description: "Details of this product, such as vendor, model, etc. are represented here.",
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
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"product_name": {
									Description: "Commercial product name. For example, VMware ESXi.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"product_type": {
									Description: "Product name provided by the vendor. For example, embeddedEsx.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"product_vendor": {
									Description: "Commercial vendor name. For example, VMware Inc.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"nr_version": {
									Description: "Hypervisor version running on the system.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
						Computed: true,
					},
					"serial": {
						Description: "Serial number of this host (internally generated).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Host health status, as reported by the hypervisor platform.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
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
					"up_time": {
						Description: "The uptime of the host, stored as Duration (from w3c).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"uuid": {
						Description: "Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vendor": {
						Description: "Commercial vendor details of this hardware.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Product version of the Host.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexHxapHostRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHxapHost{}
	if v, ok := d.GetOk("age"); ok {
		x := (v.(string))
		o.SetAge(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cluster_uuid"); ok {
		x := (v.(string))
		o.SetClusterUuid(x)
	}
	if v, ok := d.GetOk("failure_reason"); ok {
		x := (v.(string))
		o.SetFailureReason(x)
	}
	if v, ok := d.GetOk("hw_power_state"); ok {
		x := (v.(string))
		o.SetHwPowerState(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("internal_ip_address"); ok {
		x := (v.(string))
		o.SetInternalIpAddress(x)
	}
	if v, ok := d.GetOk("maintenance_mode"); ok {
		x := (v.(bool))
		o.SetMaintenanceMode(x)
	}
	if v, ok := d.GetOk("management_ip_address"); ok {
		x := (v.(string))
		o.SetManagementIpAddress(x)
	}
	if v, ok := d.GetOk("master_role"); ok {
		x := (v.(bool))
		o.SetMasterRole(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
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
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("up_time"); ok {
		x := (v.(string))
		o.SetUpTime(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
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
		return diag.Errorf("json marshal of HyperflexHxapHost object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHxapHostList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexHxapHost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexHxapHostList.GetCount()
	var i int32
	var hyperflexHxapHostResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHxapHostList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexHxapHost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexHxapHostList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexHxapHost data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["age"] = (s.GetAge())
				temp["class_id"] = (s.GetClassId())

				temp["cluster"] = flattenMapHyperflexHxapClusterRelationship(s.GetCluster(), d)

				temp["cluster_member"] = flattenMapAssetClusterMemberRelationship(s.GetClusterMember(), d)
				temp["cluster_uuid"] = (s.GetClusterUuid())

				temp["cpu_info"] = flattenMapVirtualizationCpuInfo(s.GetCpuInfo(), d)
				temp["failure_reason"] = (s.GetFailureReason())

				temp["hardware_info"] = flattenMapInfraHardwareInfo(s.GetHardwareInfo(), d)
				temp["hw_power_state"] = (s.GetHwPowerState())
				temp["hypervisor_type"] = (s.GetHypervisorType())
				temp["identity"] = (s.GetIdentity())
				temp["internal_ip_address"] = (s.GetInternalIpAddress())
				temp["maintenance_mode"] = (s.GetMaintenanceMode())
				temp["management_ip_address"] = (s.GetManagementIpAddress())
				temp["master_role"] = (s.GetMasterRole())

				temp["memory_capacity"] = flattenMapVirtualizationMemoryCapacity(s.GetMemoryCapacity(), d)
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["physical_server"] = flattenMapComputePhysicalRelationship(s.GetPhysicalServer(), d)

				temp["processor_capacity"] = flattenMapVirtualizationComputeCapacity(s.GetProcessorCapacity(), d)

				temp["product_info"] = flattenMapVirtualizationProductInfo(s.GetProductInfo(), d)

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["serial"] = (s.GetSerial())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["up_time"] = (s.GetUpTime())
				temp["uuid"] = (s.GetUuid())
				temp["vendor"] = (s.GetVendor())
				temp["nr_version"] = (s.GetVersion())
				hyperflexHxapHostResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexHxapHostResults))
	if err := d.Set("results", hyperflexHxapHostResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexHxapHostResults[0]["moid"].(string))
	return de
}
