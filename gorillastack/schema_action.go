package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func actionsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// AWS Actions
		"copy_db_snapshots": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: copyDbSnapshotsSchema()},
			Optional: true,
		},
		"copy_snapshots": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: copySnapshotsSchema()},
			Optional: true,
		},
		"create_db_snapshots": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: createDbSnapshotsSchema()},
			Optional: true,
		},
		"create_images_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: createImagesActionSchema()},
			Optional: true,
		},
		"create_snapshots_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: createSnapshotsActionSchema()},
			Optional: true,
		},
		"create_vss_snapshots_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: createVssSnapshotsActionSchema()},
			Optional: true,
		},
		"delete_detached_volumes_action": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deleteDetachedVolumesActionSchema()},
			Optional: true,
		},
		"delete_images_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: deleteImagesActionSchema()},
			Optional: true,
		},
		"delete_orphaned_snapshots": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: deleteOrphanedSnapshotsSchema()},
			Optional: true,
		},
		"delete_snapshots_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: deleteSnapshotsActionSchema()},
			Optional: true,
		},
		"ec2_command_run_powershell_script_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: ec2CommandRunPowershellScriptActionSchema()},
			Optional: true,
		},
		"ec2_command_run_shell_script_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: ec2CommandRunShellScriptActionSchema()},
			Optional: true,
		},
		"invoke_named_lambda_function_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: invokeNamedLambdaFunctionActionSchema()},
			Optional: true,
		},
		"invoke_tagged_lambda_functions_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: invokeTaggedLambdaFunctionsActionSchema()},
			Optional: true,
		},
		"notify_cost_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: notifyCostActionSchema()},
			Optional: true,
		},
		"notify_instance_count_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: notifyInstanceCountActionSchema()},
			Optional: true,
		},
		"reboot_instances_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: rebootInstancesActionSchema()},
			Optional: true,
		},
		"release_disassociated_ips_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: releaseDisassociatedIpsActionSchema()},
			Optional: true,
		},
		"resume_autoscaling_processes_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: resumeAutoscalingProcessesActionSchema()},
			Optional: true,
		},
		"start_instances_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: startInstancesActionSchema()},
			Optional: true,
		},
		"start_rds_instances_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: startRdsInstancesActionSchema()},
			Optional: true,
		},
		"stop_instances_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: stopInstancesActionSchema()},
			Optional: true,
		},
		"stop_rds_instances_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: stopRdsInstancesActionSchema()},
			Optional: true,
		},
		"suspend_autoscaling_processes_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: suspendAutoscalingProcessesActionSchema()},
			Optional: true,
		},
		"update_autoscaling_groups_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: updateAutoscalingGroupsActionSchema()},
			Optional: true,
		},
		"update_dynamodb_table_throughput_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: updateDynamodbTableThroughputActionSchema()},
			Optional: true,
		},
		"update_ecs_service_scale_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: updateEcsServiceScaleActionSchema()},
			Optional: true,
		},
		"update_provisioned_iops_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: updateProvisionedIopsActionSchema()},
			Optional: true,
		},
		"update_security_groups_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: updateSecurityGroupsActionSchema()},
			Optional: true,
		},
		// Azure Actions
		"deallocate_vms_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: deallocateVmsActionSchema()},
			Optional: true,
		},
		"start_vms_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: startVmsActionSchema()},
			Optional: true,
		},
		"update_scale_sets_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: updateScaleSetsActionSchema()},
			Optional: true,
		},
		"update_autoscale_settings_action": {
			Type:			schema.TypeList,
			Elem:			&schema.Resource{Schema: updateAutoscaleSettingsActionSchema()},
			Optional: true,
		},
		// Pauses
		"delay_pause": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: delayPauseSchema()},
			Optional: true,
		},
		"manual_approval_pause": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: manualApprovalSchema()},
			Optional: true,
		},
	}
}

/* AWS Schema functions */
func copyDbSnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func copySnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func createDbSnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func createImagesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func createSnapshotsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func createVssSnapshotsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
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
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func deleteImagesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func deleteOrphanedSnapshotsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func deleteSnapshotsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func ec2CommandRunPowershellScriptActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func ec2CommandRunShellScriptActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func invokeNamedLambdaFunctionActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func invokeTaggedLambdaFunctionsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func notifyCostActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func notifyInstanceCountActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func rebootInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func releaseDisassociatedIpsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func resumeAutoscalingProcessesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func startInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func startRdsInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func stopInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func stopRdsInstancesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func suspendAutoscalingProcessesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func updateAutoscalingGroupsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func updateDynamodbTableThroughputActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func updateEcsServiceScaleActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func updateProvisionedIopsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func updateSecurityGroupsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

/* Azure Schema functions */

func deallocateVmsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func startVmsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func updateScaleSetsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
		},
	}
}

func updateAutoscaleSettingsActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"action_id": {
			Type:			schema.TypeString,
			Computed:	true,
		},
		"index": {
			Type:			schema.TypeInt,
			Required: true,
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
			Type:			schema.TypeBool,
			Required:	true,
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