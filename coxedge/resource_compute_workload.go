package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strings"
	"time"
)

func resourceComputeWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadCreate,
		ReadContext:   resourceComputeWorkloadRead,
		UpdateContext: resourceComputeWorkloadUpdate,
		DeleteContext: resourceComputeWorkloadDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeWorkloadSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeWorkloadRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <workload_id>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("environment_name", keys[1])
		data.Set("organization_id", keys[2])
	}
	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Get the resource
	workload, err := coxEdgeClient.GetComputeWorkloadById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertComputeWorkloadAPIObjectToResourceData(data, workload)

	return diags
}

func resourceComputeWorkloadCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//convert resource data to API object
	workloadRequest := convertResourceDataToComputeWorkloadCreateAPIObject(data)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	//Call the API
	createdWorkload, err := coxEdgeClient.CreateComputeWorkload(workloadRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdWorkload.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(taskResult.Data.TaskId)
	return diags
}

func resourceComputeWorkloadUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeWorkloadDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Delete the Workload
	err := coxEdgeClient.DeleteComputeWorkloadById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeWorkloadCreateAPIObject(data *schema.ResourceData) apiclient.ComputeWorkloadRequest {

	workloadRequest := apiclient.ComputeWorkloadRequest{
		IsIPv6:                 data.Get("is_ipv6").(bool),
		NoPublicIPv4:           data.Get("no_public_ipv4").(bool),
		IsVirtualPrivateClouds: data.Get("is_virtual_private_clouds").(bool),
		IsVPC2:                 data.Get("is_vpc2").(bool),
		OperatingSystemId:      data.Get("operating_system_id").(string),
		LocationId:             data.Get("location_id").(string),
		PlanId:                 data.Get("plan_id").(string),
		Hostname:               data.Get("hostname").(string),
		Label:                  data.Get("label").(string),
		Name:                   data.Get("label").(string),
		FirstBootSshKey:        data.Get("first_boot_ssh_key").(string),
		SshKeyName:             data.Get("ssh_key_name").(string),
		FirewallId:             data.Get("firewall_id").(string),
		UserData:               data.Get("user_data").(string),
	}

	return workloadRequest
}

func convertComputeWorkloadAPIObjectToResourceData(d *schema.ResourceData, workload *apiclient.ComputeWorkload) {
	d.Set("id", workload.Id)
	d.Set("hostname", workload.Hostname)
	d.Set("label", workload.Label)
	d.Set("status", workload.Status)
	d.Set("os", workload.OS)
	d.Set("ram", workload.RAM)
	d.Set("date_created", workload.DateCreated)
	d.Set("region", workload.Region)
	d.Set("disk", workload.Disk)
	d.Set("main_ip", workload.MainIP)
	d.Set("vcpu_count", workload.VCPUCount)
	d.Set("plan", workload.Plan)
	d.Set("allowed_bandwidth", workload.AllowedBandwidth)
	d.Set("netmask_v4", workload.NetmaskV4)
	d.Set("gateway_v4", workload.GatewayV4)
	d.Set("power_status", workload.PowerStatus)
	d.Set("server_status", workload.ServerStatus)
	d.Set("v6_network", workload.V6Network)
	d.Set("v6_main_ip", workload.V6MainIP)
	d.Set("v6_network_size", workload.V6NetworkSize)
	d.Set("internal_ip", workload.InternalIP)
	d.Set("kvm", workload.KVM)
	d.Set("os_id", workload.OSID)
	d.Set("app_id", workload.AppID)
	d.Set("image_id", workload.ImageID)
	d.Set("firewall_group_id", workload.FirewallGroupID)
	d.Set("features", workload.Features)
	d.Set("tags", workload.Tags)
}
