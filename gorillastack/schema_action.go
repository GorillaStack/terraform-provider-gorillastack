package gorillastack

import (
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/constants"
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func actionsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// AWS Actions
		"check_tag_compliance": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: checkTagComplianceSchema()},
			Optional: true,
		},
		"copy_db_snapshots": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: copyDbSnapshotsSchema()},
			Optional: true,
		},
		"copy_snapshots": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: copySnapshotsSchema()},
			Optional: true,
		},
		"create_db_snapshots": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: createDbSnapshotsSchema()},
			Optional: true,
		},
		"create_images": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: createImagesActionSchema()},
			Optional: true,
		},
		"create_snapshots": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: createSnapshotsActionSchema()},
			Optional: true,
		},
		"create_vss_snapshots": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: createVssSnapshotsActionSchema()},
			Optional: true,
		},
		"delete_detached_volumes": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deleteDetachedVolumesActionSchema()},
			Optional: true,
		},
		"delete_images": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deleteImagesActionSchema()},
			Optional: true,
		},
		"delete_orphaned_snapshots": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deleteOrphanedSnapshotsSchema()},
			Optional: true,
		},
		"delete_snapshots": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deleteSnapshotsActionSchema()},
			Optional: true,
		},
		"ec2_command_run_powershell_script": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: ec2CommandRunShellScriptActionSchema()},
			Optional: true,
		},
		"ec2_command_run_shell_script": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: ec2CommandRunShellScriptActionSchema()},
			Optional: true,
		},
		"invoke_named_lambda_function": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: invokeNamedLambdaFunctionActionSchema()},
			Optional: true,
		},
		"invoke_tagged_lambda_functions": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: invokeTaggedLambdaFunctionsActionSchema()},
			Optional: true,
		},
		"notify_cost": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: notifyCostActionSchema()},
			Optional: true,
		},
		"notify_event": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: notifyEventActionSchema()},
			Optional: true,
		},
		"notify_instance_count": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: notifyInstanceCountActionSchema()},
			Optional: true,
		},
		"reboot_instances": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: rebootInstancesActionSchema()},
			Optional: true,
		},
		"release_disassociated_ips": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: releaseDisassociatedIpsActionSchema()},
			Optional: true,
		},
		"resume_autoscaling_processes": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: resumeAutoscalingProcessesActionSchema()},
			Optional: true,
		},
		"start_instances": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: startInstancesActionSchema()},
			Optional: true,
		},
		"start_rds_instances": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: startRdsInstancesActionSchema()},
			Optional: true,
		},
		"start_workspaces": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: rebootInstancesActionSchema()},
			Optional: true,
		},
		"stop_instances": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: startInstancesActionSchema()},
			Optional: true,
		},
		"stop_rds_instances": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: startRdsInstancesActionSchema()},
			Optional: true,
		},
		"stop_workspaces": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: rebootInstancesActionSchema()},
			Optional: true,
		},
		"suspend_autoscaling_processes": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: resumeAutoscalingProcessesActionSchema()},
			Optional: true,
		},
		"update_autoscaling_groups": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateAutoscalingGroupsActionSchema()},
			Optional: true,
		},
		"update_dynamodb_table_throughput": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateDynamodbTableThroughputActionSchema()},
			Optional: true,
		},
		"update_ecs_service_scale": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateEcsServiceScaleActionSchema()},
			Optional: true,
		},
		"update_provisioned_iops": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateProvisionedIopsActionSchema()},
			Optional: true,
		},
		"update_security_groups": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateSecurityGroupsActionSchema()},
			Optional: true,
		},
		// Azure Actions
		"deallocate_vms": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"restore_sql_databases": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: restoreSqlDatabasesActionSchema()},
			Optional: true,
		},
		"start_container_groups": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"start_sql_databases": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"start_vms": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"start_wvd_session_hosts": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"stop_container_groups": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"stop_sql_databases": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"stop_wvd_session_hosts": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"update_aks_node_pool_scale": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateAksNodePoolScaleActionSchema()},
			Optional: true,
		},
		"update_application_autoscaling_settings": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateApplicationAutoscalingSettings()},
			Optional: true,
		},
		"update_cosmos_container_throughput": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateCosmosContainerThroughputActionSchema()},
			Optional: true,
		},
		"update_cosmos_table_throughput": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateCosmosContainerThroughputActionSchema()},
			Optional: true,
		},
		"update_scale_sets": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateScaleSetsActionSchema()},
			Optional: true,
		},
		"update_autoscale_settings": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: updateAutoscaleSettingsActionSchema()},
			Optional: true,
		},
		// Pauses
		"delay": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: delayPauseSchema()},
			Optional: true,
		},
		"manual_approval": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: manualApprovalSchema()},
			Optional: true,
		},
	}
}

/* Common sub-schemas */

// This is used where the customer defines multiple AWS tags
// to add when creating a snapshot etc.
// The difference is that here we don't want to allow any
// tags that start with the /^aws:/ reserved namespace.
// However, in other cases we'll want no restriction on keys
// and values
func awsTagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringMatch(util.GetAwsNamespaceRegex(), "Cannot use the aws: namespace in tags"),
			Required:     true,
		},
		"value": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringMatch(util.GetAwsNamespaceRegex(), "Cannot use the aws: namespace in tags"),
			Required:     true,
		},
	}
}

func environmentVariableSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
		"secret": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}
}

/* AWS Schema functions */
func checkTagComplianceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"resource_types": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(constants.ResourceTypes, false),
			},
		},
		"report_type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.ReportType, false),
		},
		"notifications_trigger": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.NotificationsTrigger, false),
		},
		"notifications": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &schema.Resource{Schema: notificationSchema()},
		},
	}
}

func copyDbSnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"operator": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.CopyDbSnapshotsOperators, false),
		},
		"value": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"destination_region": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.AwsRegions, false),
		},
		"mode": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.CopySnapshotsModes, false),
		},
		"key_mapping": {
			Type:     schema.TypeMap,
			Optional: true,
		},
		"use_default_kms_key": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func copySnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"operator": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.CopySnapshotsOperators, false),
		},
		"value": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"destination_region": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.AwsRegions, false),
		},
		"mode": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.CopySnapshotsModes, false),
		},
		"key_mapping": {
			Type:     schema.TypeMap,
			Optional: true,
		},
		"use_default_kms_key": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func createDbSnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"copy_db_instance_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"multi_az_only": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"additional_tags": {
			Type:       schema.TypeList,
			MinItems:   1,
			MaxItems:   100,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem:       &schema.Resource{Schema: awsTagSchema()},
		},
	}
}

func createImagesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"no_reboot": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"copy_volume_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"copy_instance_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"additional_tags": {
			Type:       schema.TypeList,
			MinItems:   1,
			MaxItems:   100,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem:       &schema.Resource{Schema: awsTagSchema()},
		},
	}
}

func createSnapshotsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"copy_volume_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"copy_instance_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"additional_tags": {
			Type:       schema.TypeList,
			MinItems:   1,
			MaxItems:   100,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem:       &schema.Resource{Schema: awsTagSchema()},
		},
	}
}

func createVssSnapshotsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"copy_volume_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"copy_instance_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"exclude_boot_volume": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"use_additional_tags": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"additional_tags": {
			Type:       schema.TypeList,
			MinItems:   1,
			MaxItems:   100,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem:       &schema.Resource{Schema: awsTagSchema()},
		},
	}
}

func deleteDetachedVolumesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"dry_run": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"days_detached": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func deleteImagesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"operator": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.DeleteImagesOperators, false),
		},
		"value": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"keep_latest": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func deleteOrphanedSnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"dry_run": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func deleteSnapshotsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"dry_run": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"operator": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(constants.DeleteImagesOperators, false),
		},
		"value": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"keep_latest": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"keep_by_volume": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func ec2CommandRunShellScriptActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"commands": {
			Type:     schema.TypeList,
			MinItems: 1,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"working_directory": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"execution_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func invokeNamedLambdaFunctionActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"function_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"payload": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"replace_conflicting_vars": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"environment_variables": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Elem:     &schema.Resource{Schema: environmentVariableSchema()},
		},
		"invocation_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(constants.InvocationTypes, false),
		},
	}
}

func invokeTaggedLambdaFunctionsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"payload": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"replace_conflicting_vars": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"environment_variables": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Elem:     &schema.Resource{Schema: environmentVariableSchema()},
		},
		"invocation_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(constants.InvocationTypes, false),
		},
	}
}

func notifyCostActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"service": {
			Type:     schema.TypeString,
			Required: true,
		},
		"notifications": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &schema.Resource{Schema: notificationSchema()},
		},
	}
}

func fieldMappingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mapping_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"label": {
			Type:     schema.TypeString,
			Required: true,
		},
		"expression": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func notifyEventActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"notification_field_mapping": {
			Type:       schema.TypeList,
			MinItems:   1,
			MaxItems:   100,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem:       &schema.Resource{Schema: fieldMappingSchema()},
		},
		"notifications": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &schema.Resource{Schema: notificationSchema()},
		},
	}
}

func notifyInstanceCountActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"notifications": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &schema.Resource{Schema: notificationSchema()},
		},
	}
}

func rebootInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func releaseDisassociatedIpsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"days_disassociated": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resumeAutoscalingProcessesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"processes": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(constants.ASGProcesses, false),
			},
		},
	}
}

func startInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"wait_instance_state": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"wait_instance_status": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"wait_system_status": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func startRdsInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"target_clusters": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"wait_instance_state": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"wait_instance_status": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"wait_system_status": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func updateAutoscalingGroupsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"min": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"desired": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"store_existing_asg_settings": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"restore_to_previous_asg_settings": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"ignore_if_no_cached_asg_settings": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func updateDynamodbTableThroughputActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"read_units": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"write_units": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func updateEcsServiceScaleActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"desired_count": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"store_existing_desired_count": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"restore_to_previous_desired_count": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"ignore_if_no_cached_desired_count": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}
}

func updateProvisionedIopsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"iops": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

/* helper functions with schemas for updateSecurityGroupsActionSchema */
func ruleChangeMatchFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol_int": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"protocol_name"},
		},
		"protocol_name": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"protocol_int"},
		},
		"port": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"endpoint": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"endpoint_description": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func ruleChangeMatchSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"fields"}, false),
		},
		"direction": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"ingress", "egress"}, false),
		},
		"fields": {
			Type:     schema.TypeMap,
			Optional: true,
		},
	}
}

func ruleChangeChangeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"operation": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"delete"}, false),
		},
	}
}

func securityGroupRuleChangesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"match": {
			Type:     schema.TypeMap,
			Required: true,
			Elem:     &schema.Resource{Schema: ruleChangeMatchSchema()},
		},
		"change": {
			Type:     schema.TypeMap,
			Required: true,
			Elem:     &schema.Resource{Schema: ruleChangeChangeSchema()},
		},
	}
}

func updateSecurityGroupsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"rule_changes": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			MaxItems: 100,
			Elem:     &schema.Resource{Schema: securityGroupRuleChangesSchema()},
		},
	}
}

/* Azure Schema functions */

func deallocateVmsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func updateAksNodePoolScaleActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"min_count": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max_count": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"restore_to_previous_scale": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func updateApplicationAutoscalingSettings() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"scalable_dimension": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"service_namespace": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"min_capacity": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max_capacity": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"store_existing_autoscaling_settings": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"restore_existing_autoscaling_settings": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"ignore_if_no_cached_autoscaling_settings": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func restoreSqlDatabasesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"database_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"database_server": {
			Type:     schema.TypeString,
			Required: true,
		},
		"resource_group": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func updateScaleSetsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"capacity": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func updateAutoscaleSettingsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"min": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"desired": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func updateCosmosContainerThroughputActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"throughput": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"restore_to_previous_throughput": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

/* Pause Schema functions */
func delayPauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"wait_duration": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func manualApprovalSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"index": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"wait_duration": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"approve_on_expiry": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"notifications": {
			Type:     schema.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &schema.Resource{Schema: notificationSchema()},
		},
	}
}
