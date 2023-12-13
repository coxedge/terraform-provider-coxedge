package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"time"
)

func resourceRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouteCreate,
		ReadContext:   resourceRouteRead,
		UpdateContext: resourceRouteUpdate,
		DeleteContext: resourceRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getRoutes(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceRouteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	routeRequest := convertResourceDataToRouteAPIObject(d)

	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)
	vpcId := d.Get("vpc_id").(string)

	//Call the API
	createdSubnet, err := coxEdgeClient.CreateRoute(routeRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	timeout := d.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdSubnet.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}
	//vpc Id is not present in task result
	//calling getAllVPCs api and filtering it out using name  to get vpcId
	//if vpc is not present in getAllVPCs api, throwing error
	var resourceId = ""
	if taskResult.Data.TaskStatus == "SUCCESS" {
		routeList, err := coxEdgeClient.GetAllRoutes(vpcId, environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, route := range routeList {
			if route.Name == d.Get("name").(string) {
				resourceId = route.Id
			}
		}

	}
	if resourceId != "" {
		d.SetId(resourceId)
	} else {
		diagn := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Something went wrong. Route name not found",
			Detail:   "Something went wrong. Route name not found",
		}
		diags = append(diags, diagn)
	}
	return diags
}

func resourceRouteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.SetId(keys[0])
		d.Set("vpc_id", keys[1])
		d.Set("environment_name", keys[2])
		d.Set("organization_id", keys[3])
	}
	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)
	vpcId := d.Get("vpc_id").(string)
	//Get the resource
	route, err := coxEdgeClient.GetRoute(resourceId, vpcId, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertRouteAPIObjectToResourceData(d, route)

	return diags
}

func resourceRouteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diag := diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Updating of Route in VPC is not allowed",
		Detail:   "Updating of Route in VPC is not allowed",
	}
	diags = append(diags, diag)
	return diags
}

func resourceRouteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)

	routeRequest := convertResourceDataToRouteAPIObject(d)

	//Delete the Site
	deleteVPC, err := coxEdgeClient.DeleteRoute(routeRequest, environmentName, organizationId)
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

func convertResourceDataToRouteAPIObject(d *schema.ResourceData) apiclient.RouteRequest {
	routeObject := apiclient.RouteRequest{
		Id:      d.Get("id").(string),
		VpcId:   d.Get("vpc_id").(string),
		Name:    d.Get("name").(string),
		StackId: d.Get("stack_id").(string),
		Status:  d.Get("status").(string),
	}
	routeObject.DestinationCidrs = []string{}
	for _, val := range d.Get("destination_cidrs").([]interface{}) {
		routeObject.DestinationCidrs = append(routeObject.DestinationCidrs, val.(string))
	}
	routeObject.NextHops = []string{}
	for _, val := range d.Get("next_hops").([]interface{}) {
		routeObject.NextHops = append(routeObject.NextHops, val.(string))
	}
	return routeObject
}

func convertRouteAPIObjectToResourceData(d *schema.ResourceData, route *apiclient.Route) {
	//Store the data
	d.Set("id", route.Id)
	d.Set("status", route.Status)
	d.Set("stack_id", route.StackId)
	d.Set("vpc_d", route.VpcId)
	d.Set("vpc_name", route.VpcName)
	d.Set("name", route.Name)
	d.Set("slug", route.Slug)
	d.Set("destination_cidrs", route.DestinationCidrs)
	d.Set("next_hops", route.NextHops)

}
