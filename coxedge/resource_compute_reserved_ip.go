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

func resourceComputeReservedIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeReservedIPCreate,
		ReadContext:   resourceComputeReservedIPRead,
		UpdateContext: resourceComputeReservedIPUpdate,
		DeleteContext: resourceComputeReservedIPDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeReservedIPSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeReservedIPCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	reservedIPRequest := convertResourceDataToComputeReservedIPRequest(data)

	existingList, err := coxEdgeClient.GetComputeReservedIPs(environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	existingIDs := make(map[string]bool)
	for _, item := range existingList {
		existingIDs[item.ID] = true
	}

	//Call the API
	reservedIPResponse, err := coxEdgeClient.CreateComputeReservedIP(reservedIPRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, reservedIPResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResult.Data.TaskStatus == "SUCCESS" {
		afterList, err := coxEdgeClient.GetComputeReservedIPs(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		var missingItem *apiclient.ComputeReservedIP
		for _, item := range afterList {
			if !existingIDs[item.ID] {
				missingItem = &item
				data.SetId(missingItem.ID)
				return diags
			}
		}
	}
	//Save the Id
	data.SetId(taskResult.Data.TaskId)
	return diags
}

func resourceComputeReservedIPRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <reserved_ip>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("environment_name", keys[2])
		data.Set("organization_id", keys[3])
	}

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	reservedIPId := data.Id()

	computeReservedIP, err := coxEdgeClient.GetComputeReservedIPById(environmentName, organizationId, reservedIPId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertReservedIPToResourceData(data, computeReservedIP)
	data.SetId(reservedIPId)
	return diags
}

func resourceComputeReservedIPUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	label := data.Get("label").(string)
	reservedIPId := data.Id()

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	reservedIPRequest := apiclient.UpdateComputeReservedIPRequest{
		Label: label,
	}
	//Call the API
	updateResponse, err := coxEdgeClient.UpdateComputeReservedIP(reservedIPRequest, environmentName, organizationId, reservedIPId)
	if err != nil {
		return diag.FromErr(err)
	}
	timeout := data.Timeout(schema.TimeoutUpdate)

	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, updateResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(reservedIPId)
	return diags
}

func resourceComputeReservedIPDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Delete the Storage
	err := coxEdgeClient.DeleteComputeReservedIPById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeReservedIPRequest(data *schema.ResourceData) apiclient.ComputeReservedIPRequest {
	reservedIPRequest := apiclient.ComputeReservedIPRequest{
		Region: data.Get("region").(string),
		IpType: data.Get("ip_type").(string),
		Label:  data.Get("label").(string),
	}
	return reservedIPRequest
}

func convertReservedIPToResourceData(d *schema.ResourceData, reservedIP *apiclient.ComputeReservedIP) {
	d.Set("id", reservedIP.ID)
	d.Set("region", reservedIP.Region)
	d.Set("location", reservedIP.Location)
	d.Set("ip_type", reservedIP.IPType)
	d.Set("subnet", reservedIP.Subnet)
	d.Set("subnet_size", reservedIP.SubnetSize)
	d.Set("label", reservedIP.Label)
	d.Set("instance_id", reservedIP.InstanceID)
	d.Set("reserved_ip", reservedIP.ReservedIP)
	if reservedIP.InstanceID != "" {
		d.Set("is_workload_attached", true)
	} else {
		d.Set("is_workload_attached", false)
	}

}
