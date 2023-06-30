package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"time"
)

func resourceVPC() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVPCCreate,
		ReadContext:   resourceVPCRead,
		UpdateContext: resourceVPCUpdate,
		DeleteContext: resourceVPCDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getVPCSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceVPCCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vpcRequest := convertResourceDataToVPCCreateAPIObject(d)

	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)

	//Call the API
	createdVPC, err := coxEdgeClient.CreateVPCNetwork(vpcRequest, environmentName, organizationId)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("------------------err %v", err))
		return diag.FromErr(err)
	}
	tflog.Info(ctx, fmt.Sprintf("------------------createdVPC %v", createdVPC))

	timeout := d.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdVPC.TaskId, timeout)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("------------------ task %v", err))
		return diag.FromErr(err)
	}
	tflog.Error(ctx, fmt.Sprintf("------------------taskResult %v", taskResult))

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceVPCRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.SetId(keys[0])
		d.Set("environment_name", keys[1])
		d.Set("organization_id", keys[2])
	}
	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)
	//Get the resource
	vpc, err := coxEdgeClient.GetVPCNetwork(resourceId, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertVPCAPIObjectToResourceData(d, vpc)

	return diags
}

func resourceVPCUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diag := diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Updating of VPC network is not allowed",
		Detail:   "Updating of VPC network is not allowed",
	}
	diags = append(diags, diag)
	return diags
}

func resourceVPCDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)

	vpcRequest := convertResourceDataToVPCCreateAPIObject(d)

	//Delete the Site
	deleteVPC, err := coxEdgeClient.DeleteVPCNetwork(vpcRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	timeout := d.Timeout(schema.TimeoutDelete)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, deleteVPC.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}
	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToVPCCreateAPIObject(d *schema.ResourceData) apiclient.VPCRequest {
	vpcObject := apiclient.VPCRequest{
		Id:         d.Get("id").(string),
		Name:       d.Get("name").(string),
		Slug:       d.Get("slug").(string),
		StackId:    d.Get("stack_id").(string),
		Cidr:       d.Get("cidr").(string),
		DefaultVpc: d.Get("default_vpc").(bool),
		Status:     d.Get("status").(string),
		Created:    d.Get("created").(string),
	}
	return vpcObject
}

func convertVPCAPIObjectToResourceData(d *schema.ResourceData, vpc *apiclient.VPCs) {
	//Store the data
	d.Set("id", vpc.Id)
	d.Set("name", vpc.Name)
	d.Set("stack_id", vpc.StackId)
	d.Set("slug", vpc.Slug)
	d.Set("cidr", vpc.Cidr)
	d.Set("default_vpc", vpc.DefaultVpc)
	d.Set("created", vpc.Created)
	d.Set("routes", vpc.Routes)
	d.Set("status", vpc.Status)

	subnetList := make([]map[string]string, len(vpc.Subnets), len(vpc.Subnets))
	for i, subnet := range vpc.Subnets {
		item := make(map[string]string)
		item["id"] = subnet.Id
		item["name"] = subnet.Name
		item["stack_id"] = subnet.StackId
		item["vpc_id"] = subnet.VpcId
		item["slug"] = subnet.Slug
		item["cidr"] = subnet.Cidr
		item["status"] = subnet.Status
		subnetList[i] = item
	}
	d.Set("subnets", subnetList)

	routeList := make([]map[string]interface{}, len(vpc.Routes), len(vpc.Routes))
	for i, route := range vpc.Routes {
		item := make(map[string]interface{})
		item["id"] = route.Id
		item["name"] = route.Name
		item["stack_id"] = route.StackId
		item["vpc_id"] = route.VpcId
		item["slug"] = route.Slug
		item["destination_cidrs"] = route.DestinationCidrs
		item["next_hops"] = route.NextHops
		item["status"] = route.Status
		routeList[i] = item
	}
	d.Set("routes", routeList)
}
