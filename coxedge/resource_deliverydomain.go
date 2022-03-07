package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeliveryDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeliveryDomainCreate,
		ReadContext:   resourceDeliveryDomainRead,
		//TODO: Implement
		//UpdateContext: resourceDeliveryDomainUpdate,
		DeleteContext: resourceDeliveryDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getDeliveryDomainSchema(),
	}
}

func resourceDeliveryDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newDeliveryDomain := convertResourceDataToDeliveryDomainCreateAPIObject(d)

	//Call the API
	createdDeliveryDomain, err := coxEdgeClient.CreateDeliveryDomain(newDeliveryDomain)
	if err != nil {
		return diag.FromErr(err)
	}

	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, createdDeliveryDomain.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the ID
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceDeliveryDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Id()

	//Get the resource
	deliveryDomain, err := coxEdgeClient.GetDeliveryDomain(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertDeliveryDomainAPIObjectToResourceData(d, deliveryDomain)

	//Update state
	resourceDeliveryDomainRead(ctx, d, m)

	return diags
}

/*func resourceDeliveryDomainUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedDeliveryDomain := convertResourceDataToDeliveryDomainCreateAPIObject(d)

	//Call the API
	_, err := coxEdgeClient.UpdateDeliveryDomain(resourceId, updatedDeliveryDomain)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceDeliveryDomainRead(ctx, d, m)
}*/

func resourceDeliveryDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Delete the DeliveryDomain
	err := coxEdgeClient.DeleteDeliveryDomain(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToDeliveryDomainCreateAPIObject(d *schema.ResourceData) apiclient.DeliveryDomainCreateRequest {
	//Create update deliveryDomain struct
	updatedDeliveryDomain := apiclient.DeliveryDomainCreateRequest{
		EnvironmentName: d.Get("environment_name").(string),
		Domain:          d.Get("domain").(string),
	}

	return updatedDeliveryDomain
}

func convertDeliveryDomainAPIObjectToResourceData(d *schema.ResourceData, deliveryDomain *apiclient.DeliveryDomain) {
	//Store the data
	d.Set("id", deliveryDomain.Id)
	d.Set("stack_id", deliveryDomain.StackId)
	d.Set("domain", deliveryDomain.Domain)
	d.Set("updated_at", deliveryDomain.UpdatedAt)
}
