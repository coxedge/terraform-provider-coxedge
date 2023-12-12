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

func resourceComputeWorkloadIPv6ReverseDNS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadIPv6ReverseDNSCreate,
		ReadContext:   resourceComputeWorkloadIPv6ReverseDNSRead,
		UpdateContext: resourceComputeWorkloadIPv6ReverseDNSUpdate,
		DeleteContext: resourceComputeWorkloadIPv6ReverseDNSDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeWorkloadIPv6ReverseDNSSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeWorkloadIPv6ReverseDNSCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//convert resource data to API object
	reverseDNSRequest := convertResourceDataToComputeWorkloadIPv6ReverseDNSCreateAPIObject(data)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	//Call the API
	createdReverseDNS, err := coxEdgeClient.CreateComputeWorkloadIPv6ReverseDNSById(reverseDNSRequest, environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdReverseDNS.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(taskResult.Data.TaskId)
	resourceComputeWorkloadIPv6ReverseDNSRead(ctx, data, i)
	return diags
}

func resourceComputeWorkloadIPv6ReverseDNSRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <workload_id>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.Set("workload_id", keys[0])
		data.Set("environment_name", keys[1])
		data.Set("organization_id", keys[2])
	}
	//Get the resource Id
	workloadId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Get the resource
	reverseDNS, err := coxEdgeClient.GetComputeWorkloadIPv6ReverseDNSById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertComputeWorkloadIPv6ReverseDNSAPIObjectToResourceData(data, &reverseDNS[0],workloadId)

	return diags
}

func resourceComputeWorkloadIPv6ReverseDNSUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeWorkloadIPv6ReverseDNSDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	workloadId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)
	ip := data.Get("ip").(string)

	//Delete the Workload
	err := coxEdgeClient.DeleteComputeWorkloadIPv6ReverseDNSById(environmentName, organizationId, workloadId, ip)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeWorkloadIPv6ReverseDNSCreateAPIObject(data *schema.ResourceData) apiclient.ComputeWorkloadIPv6ReverseDNSRequest {

	reverseDNSRequest := apiclient.ComputeWorkloadIPv6ReverseDNSRequest{
		Ip:      data.Get("ip").(string),
		Reverse: data.Get("reverse").(string),
	}

	return reverseDNSRequest
}

func convertComputeWorkloadIPv6ReverseDNSAPIObjectToResourceData(d *schema.ResourceData, reverseDNS *apiclient.IPv6ReverseDNSConfiguration,workloadId string) {
	d.Set("id", reverseDNS.Id)
	d.Set("workload_id", workloadId)
	d.Set("ip", reverseDNS.Ip)
	d.Set("reverse", reverseDNS.Reverse)
}
