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

func resourceComputeVPC() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeVPCCreate,
		ReadContext:   resourceComputeVPCRead,
		UpdateContext: resourceComputeVPCUpdate,
		DeleteContext: resourceComputeVPCDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeVPCSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeVPCCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vpcRequest := convertResourceDataToComputeVPCRequest(data)

	existingList, err := coxEdgeClient.GetComputeVPC(environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	existingIDs := make(map[string]bool)
	for _, item := range existingList {
		existingIDs[item.ID] = true
	}

	//Call the API
	vpcResponse, err := coxEdgeClient.CreateComputeVPC(vpcRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, vpcResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResult.Data.TaskStatus == "SUCCESS" {
		afterList, err := coxEdgeClient.GetComputeVPC(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		var missingItem *apiclient.ComputeVPC
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

func resourceComputeVPCRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <vpcId>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("environment_name", keys[2])
		data.Set("organization_id", keys[3])
	}

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	vpcId := data.Id()

	computeVPC, err := coxEdgeClient.GetComputeVPCById(environmentName, organizationId, vpcId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertVPCToResourceData(data, computeVPC)
	data.SetId(vpcId)
	return diags
}

func resourceComputeVPCUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeVPCDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Delete the Storage
	err := coxEdgeClient.DeleteComputeVPCById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeVPCRequest(data *schema.ResourceData) apiclient.ComputeVPCRequest {
	vpcRequest := apiclient.ComputeVPCRequest{
		LocationID:    data.Get("location_id").(string),
		V4SubnetMask:  data.Get("v4_subnet_mask").(int),
		NetworkPrefix: data.Get("network_prefix").(int),
		IPRange:       data.Get("ip_range").(string),
		RouteID:       data.Get("route_id").(string),
		V4Subnet:      data.Get("v4_subnet").(string),
		Description:   data.Get("description").(string),
	}

	// Parse and add routes
	if routesData, ok := data.GetOk("routes"); ok {
		routes := make([]apiclient.ComputeVPCRouteRequest, 0)
		for _, route := range routesData.([]interface{}) {
			if route != nil {
				routeData, ok := route.(map[string]interface{})
				if !ok {
					// Handle error: route is not a map[string]interface{}
					continue
				}
				routes = append(routes, apiclient.ComputeVPCRouteRequest{
					Destination:   routeData["destination"].(string),
					NetworkPrefix: routeData["network_prefix"].(string),
					TargetAddress: routeData["target_address"].(string),
				})
			}
		}
		vpcRequest.Routes = routes
	}

	return vpcRequest
}

func convertVPCToResourceData(d *schema.ResourceData, vpc2 *apiclient.ComputeVPC) {
	d.Set("id", vpc2.ID)
	d.Set("date_created", vpc2.DateCreated)
	d.Set("region", vpc2.Region)
	d.Set("location", vpc2.Location)
	d.Set("description", vpc2.Description)
	d.Set("v4_subnet", vpc2.V4Subnet)
	d.Set("v4_subnet_mask", vpc2.V4SubnetMask)
	d.Set("subnet", vpc2.Subnet)
}
