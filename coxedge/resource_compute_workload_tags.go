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

func resourceComputeWorkloadTags() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadTagsCreate,
		ReadContext:   resourceComputeWorkloadTagsRead,
		UpdateContext: resourceComputeWorkloadTagsUpdate,
		DeleteContext: resourceComputeWorkloadTagsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeWorkloadTagsSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeWorkloadTagsCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	var diags diag.Diagnostics
	resourceId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)
	tagRequest := apiclient.ComputeWorkloadTagRequest{
		Tag: data.Get("tag").(string),
	}
	//Call the API
	tagResponse, err := coxEdgeClient.AddComputeWorkloadTag(tagRequest, environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, tagResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(data.Get("tag").(string))
	return diags
}

func resourceComputeWorkloadTagsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <workload_id>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("workload_id", keys[0])
		data.Set("environment_name", keys[1])
		data.Set("organization_id", keys[2])
		data.Set("tag", keys[3])
	}
	//Get the resource Id
	resourceId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)
	tag := data.Get("tag").(string)

	//Get the resource
	tagResponse, err := coxEdgeClient.GetComputeWorkloadTagByTagId(environmentName, organizationId, resourceId, tag)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertComputeWorkloadTagAPIObjectToResourceData(data, tagResponse)

	return diags
}

func resourceComputeWorkloadTagsUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeWorkloadTagsDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource Id
	resourceId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	tagRequest := apiclient.ComputeWorkloadTagRequest{
		Id:  data.Get("tag").(string),
		Tag: data.Get("tag").(string),
	}
	//Get the resource
	err := coxEdgeClient.DeleteComputeWorkloadTag(tagRequest, environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertComputeWorkloadTagAPIObjectToResourceData(d *schema.ResourceData, userData *apiclient.ComputeWorkloadTag) {
	d.Set("id", userData.ID)
	d.Set("tag", userData.Tag)
}
