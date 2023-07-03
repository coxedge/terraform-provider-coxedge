package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
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

	subnetRequest := convertResourceDataToSubnetAPIObject(d)

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
	subnet, err := coxEdgeClient.GetSubnet(resourceId, vpcId, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertSubnetAPIObjectToResourceData(d, subnet)

	return diags
}

func resourceSubnetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diag := diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Updating of Subnet in VPC is not allowed",
		Detail:   "Updating of Subnet in VPC is not allowed",
	}
	diags = append(diags, diag)
	return diags
}

func resourceSubnetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)

	subnetRequest := convertResourceDataToSubnetAPIObject(d)

	//Delete the Site
	deleteVPC, err := coxEdgeClient.DeleteSubnet(subnetRequest, environmentName, organizationId)
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

func convertResourceDataToSubnetAPIObject(d *schema.ResourceData) apiclient.SubnetRequest {
	subnetObject := apiclient.SubnetRequest{
		Id:      d.Get("id").(string),
		VpcId:   d.Get("vpc_id").(string),
		Name:    d.Get("name").(string),
		Slug:    d.Get("slug").(string),
		Cidr:    d.Get("cidr").(string),
		StackId: d.Get("stack_id").(string),
	}
	return subnetObject
}

func convertSubnetAPIObjectToResourceData(d *schema.ResourceData, subnet *apiclient.Subnets) {
	//Store the data
	d.Set("id", subnet.Id)
	d.Set("status", subnet.Status)
	d.Set("stack_id", subnet.StackId)
	d.Set("vpc_d", subnet.VpcId)
	d.Set("vpc_slug", subnet.VpcSlug)
	d.Set("vpc_name", subnet.VpcName)
	d.Set("name", subnet.Name)
	d.Set("slug", subnet.Slug)
	d.Set("cidr", subnet.Cidr)

}
