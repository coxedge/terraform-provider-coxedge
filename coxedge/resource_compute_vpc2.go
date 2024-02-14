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

func resourceComputeVPC2() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeVPC2Create,
		ReadContext:   resourceComputeVPC2Read,
		UpdateContext: resourceComputeVPC2Update,
		DeleteContext: resourceComputeVPC2Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeVPC2Schema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeVPC2Create(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vpc2Request := convertResourceDataToComputeVPC2CreateAPIObject(data)

	existingList, err := coxEdgeClient.GetComputeVPC2(environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	existingIDs := make(map[string]bool)
	for _, item := range existingList {
		existingIDs[item.ID] = true
	}

	//Call the API
	firewallResponse, err := coxEdgeClient.CreateComputeVPC2(vpc2Request, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, firewallResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResult.Data.TaskStatus == "SUCCESS" {
		afterList, err := coxEdgeClient.GetComputeVPC2(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		var missingItem *apiclient.ComputeVPC2
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

func resourceComputeVPC2Read(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <vpc2Id>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("environment_name", keys[2])
		data.Set("organization_id", keys[3])
	}

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	vpc2Id := data.Id()

	computeVPC2, err := coxEdgeClient.GetComputeVPC2ById(environmentName, organizationId, vpc2Id)
	if err != nil {
		return diag.FromErr(err)
	}
	convertVPC2ToResourceData(data, computeVPC2)
	data.SetId(vpc2Id)
	return diags
}

func resourceComputeVPC2Update(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeVPC2Delete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Delete the Storage
	err := coxEdgeClient.DeleteComputeVPC2ById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeVPC2CreateAPIObject(data *schema.ResourceData) apiclient.ComputeVPC2Request {
	vpc2Request := apiclient.ComputeVPC2Request{
		LocationID:   data.Get("location_id").(string),
		PrefixLength: data.Get("prefix_length").(string),
		IPRange:      data.Get("ip_range").(string),
		IPBlock:      data.Get("ip_block").(string),
		Description:  data.Get("description").(string),
	}
	return vpc2Request
}

func convertVPC2ToResourceData(d *schema.ResourceData, vpc2 *apiclient.ComputeVPC2) {
	d.Set("id", vpc2.ID)
	d.Set("date_created", vpc2.DateCreated)
	d.Set("region", vpc2.Region)
	d.Set("location", vpc2.Location)
	d.Set("description", vpc2.Description)
	d.Set("ip_block", vpc2.IPBlock)
	d.Set("prefix_length", vpc2.PrefixLength)
	d.Set("subnet", vpc2.Subnet)
}
