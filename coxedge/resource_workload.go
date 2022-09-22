/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"coxedge/terraform-provider/coxedge/utils"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkloadCreate,
		ReadContext:   resourceWorkloadRead,
		UpdateContext: resourceWorkloadUpdate,
		DeleteContext: resourceWorkloadDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getWorkloadSchema(),
	}
}

func resourceWorkloadCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newWorkload := convertResourceDataToWorkloadCreateAPIObject(d)

	for _, deployment := range newWorkload.Deployments {
		if !deployment.EnableAutoScaling {
			if deployment.InstancesPerPop == -1 {
				return diag.Errorf("instances_per_pop must be set when autoscaling is disabled.")
			}
		}
	}

	organizationId := d.Get("organization_id").(string)

	//Call the API
	createdWorkload, err := coxEdgeClient.CreateWorkload(newWorkload, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, createdWorkload.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceWorkloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <workload_id>:<environment_name>:<organization_id>
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.SetId(keys[0])
		d.Set("environment_name", keys[1])
		d.Set("organization_id", keys[2])
	}
	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	//Get the resource
	workload, err := coxEdgeClient.GetWorkload(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertWorkloadAPIObjectToResourceData(d, workload)

	return diags
}

func resourceWorkloadUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()

	//Convert resource data to API object
	updatedWorkload := convertResourceDataToWorkloadCreateAPIObject(d)
	organizationId := d.Get("organization_id").(string)

	//Call the API
	createdWorkload, err := coxEdgeClient.UpdateWorkload(resourceId, updatedWorkload, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, createdWorkload.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}
	for _, deployment := range updatedWorkload.Deployments {
		for i := 0; i < (deployment.InstancesPerPop)*len(deployment.Pops); i++ {
			time.Sleep(20 * time.Second)
		}
	}
	//Set last_updated
	//d.Set("last_updated", time.Now().Format(time.RFC850))
	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return resourceWorkloadRead(ctx, d, m)
}

func resourceWorkloadDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	//Delete the Workload
	err := coxEdgeClient.DeleteWorkload(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToWorkloadCreateAPIObject(d *schema.ResourceData) apiclient.WorkloadCreateRequest {
	//Create update workload struct
	updatedWorkload := apiclient.WorkloadCreateRequest{
		Name:                d.Get("name").(string),
		EnvironmentName:     d.Get("environment_name").(string),
		Image:               d.Get("image").(string),
		Specs:               d.Get("specs").(string),
		Type:                d.Get("type").(string),
		AddAnyCastIpAddress: d.Get("add_anycast_ip_address").(bool),
		FirstBootSSHKey:     d.Get("first_boot_ssh_key").(string),
		UserData:            d.Get("user_data").(string),
		ContainerEmail:      d.Get("container_email").(string),
		ContainerServer:     d.Get("container_server").(string),
		ContainerUsername:   d.Get("container_username").(string),
		ContainerPassword:   d.Get("container_password").(string),
		Slug:                d.Get("slug").(string),
	}

	//Set commands
	updatedWorkload.Commands = make([]string, len(d.Get("commands").([]interface{})))
	for i, command := range d.Get("commands").([]interface{}) {
		updatedWorkload.Commands[i] = command.(string)
	}

	//Convert ports
	for _, entry := range d.Get("ports").([]interface{}) {
		convertedEntry := entry.(map[string]interface{})
		portSpec := apiclient.WorkloadPort{
			Protocol:       convertedEntry["protocol"].(string),
			PublicPort:     convertedEntry["public_port"].(string),
			PublicPortDesc: convertedEntry["public_port_desc"].(string),
			PublicPortSrc:  convertedEntry["public_port_src"].(string),
		}
		updatedWorkload.Ports = append(updatedWorkload.Ports, portSpec)
	}

	//Convert Persistent Storage
	for _, entry := range d.Get("persistent_storages").([]interface{}) {
		convertedEntry := entry.(map[string]interface{})
		storageSpec := apiclient.WorkloadPersistentStorage{
			Path: convertedEntry["path"].(string),
			Size: convertedEntry["size"].(int),
		}
		updatedWorkload.PersistentStorage = append(updatedWorkload.PersistentStorage, storageSpec)
	}

	//Convert env vars
	for key, value := range d.Get("environment_variables").(map[string]interface{}) {
		newVar := apiclient.WorkloadEnvironmentVariable{
			Key:   key,
			Value: value.(string),
		}
		updatedWorkload.EnvironmentVariables = append(updatedWorkload.EnvironmentVariables, newVar)
	}

	//Convert secret env vars
	for key, value := range d.Get("secret_environment_variables").(map[string]interface{}) {
		newVar := apiclient.WorkloadEnvironmentVariable{
			Key:   key,
			Value: value.(string),
		}
		updatedWorkload.SecretEnvironmentVariables = append(updatedWorkload.SecretEnvironmentVariables, newVar)
	}

	//Convert deployments
	for _, entry := range d.Get("deployment").([]interface{}) {
		convertedEntry := entry.(map[string]interface{})
		deploymentEntry := apiclient.WorkloadAutoscaleDeployment{
			Name:               convertedEntry["name"].(string),
			Pops:               utils.ConvertListInterfaceToStringArray(convertedEntry["pops"]),
			EnableAutoScaling:  convertedEntry["enable_autoscaling"].(bool),
			InstancesPerPop:    convertedEntry["instances_per_pop"].(int),
			MaxInstancesPerPop: convertedEntry["max_instances_per_pop"].(int),
			MinInstancesPerPop: convertedEntry["min_instances_per_pop"].(int),
			CPUUtilization:     convertedEntry["cpu_utilization"].(int),
		}
		updatedWorkload.Deployments = append(updatedWorkload.Deployments, deploymentEntry)
	}

	return updatedWorkload
}

func convertWorkloadAPIObjectToResourceData(d *schema.ResourceData, workload *apiclient.Workload) {
	//Store the data
	d.Set("id", workload.Id)
	d.Set("name", workload.Name)
	d.Set("image", workload.Image)
	d.Set("specs", workload.Specs)
	d.Set("type", workload.Type)
	d.Set("anycast_ip_address", workload.AnycastIpAddress)
	d.Set("commands", workload.Commands)
	d.Set("container_email", workload.ContainerEmail)
	d.Set("container_username", workload.ContainerUsername)
	d.Set("container_server", workload.ContainerServer)
	d.Set("first_boot_ssh_key", workload.FirstBootSshKey)
	d.Set("user_data", workload.UserData)
	//Now the list structures
	deployments := make([]map[string]interface{}, len(workload.Deployments), len(workload.Deployments))
	for i, deployment := range workload.Deployments {
		item := make(map[string]interface{})
		item["name"] = deployment.Name
		item["pops"] = deployment.Pops
		item["enable_autoscaling"] = deployment.EnableAutoScaling
		item["instances_per_pop"] = deployment.InstancesPerPop
		item["max_instances_per_pop"] = deployment.MaxInstancesPerPop
		item["min_instances_per_pop"] = deployment.MinInstancesPerPop
		item["cpu_utilization"] = deployment.CPUUtilization
		deployments[i] = item
	}
	d.Set("deployment", deployments)

	ports := make([]map[string]string, len(workload.Ports), len(workload.Ports))
	for i, portObj := range workload.Ports {
		item := make(map[string]string)
		item["protocol"] = portObj.Protocol
		item["public_port"] = portObj.PublicPort
		item["public_port_desc"] = portObj.PublicPortDesc
		item["public_port_src"] = portObj.PublicPortSrc
		ports[i] = item
	}
	d.Set("ports", ports)

	persistentStorageMap := make([]map[string]interface{}, len(workload.PersistentStorages), len(workload.PersistentStorages))
	for i, persistentStorageObj := range workload.PersistentStorages {
		item := make(map[string]interface{})
		item["path"] = persistentStorageObj.Path
		item["size"] = persistentStorageObj.Size
		persistentStorageMap[i] = item
	}
	d.Set("persistent_storages", persistentStorageMap)

	envVars := make(map[string]string, len(workload.EnvironmentVariables))
	for _, envVarObj := range workload.EnvironmentVariables {
		envVars[envVarObj.Key] = envVarObj.Value
	}
	d.Set("environment_variables", envVars)

	secretEnvVars := make(map[string]string, len(workload.SecretEnvironmentVariables))
	for _, envVarObj := range workload.SecretEnvironmentVariables {
		secretEnvVars[envVarObj.Key] = envVarObj.Value
	}
	d.Set("secret_environment_variables", secretEnvVars)
}
