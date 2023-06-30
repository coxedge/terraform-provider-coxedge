package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSubnetCreate,
		ReadContext:   resourceSubnetRead,
		UpdateContext: resourceSubnetUpdate,
		DeleteContext: resourceSubnetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getSubnets(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceSubnetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	subnetRequest := convertResourceDataToSubnetCreateAPIObject(d)

	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)
	vpcId := d.Get("vpc_id").(string)

	//Call the API
	createdSubnet, err := coxEdgeClient.CreateSubnet(subnetRequest, environmentName, organizationId)
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
		subnetList, err := coxEdgeClient.GetAllSubnets(vpcId, environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, subnt := range subnetList {
			if subnt.Name == d.Get("name").(string) {
				resourceId = subnt.Id
			}
		}

	}
	if resourceId != "" {
		d.SetId(resourceId)
	} else {
		diagn := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Something went wrong. Subnet name not found",
			Detail:   "Something went wrong. Subnet name not found",
		}
		diags = append(diags, diagn)
	}
	return diags
}

func resourceSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSubnetUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSubnetDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func convertResourceDataToSubnetCreateAPIObject(d *schema.ResourceData) apiclient.SubnetRequest {
	subnetObject := apiclient.SubnetRequest{
		Id:    d.Get("id").(string),
		VpcId: d.Get("vpc_id").(string),
		Name:  d.Get("name").(string),
		Slug:  d.Get("slug").(string),
		Cidr:  d.Get("cidr").(string),
	}
	return subnetObject
}
