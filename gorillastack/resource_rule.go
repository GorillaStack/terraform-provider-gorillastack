package gorillastack

import (
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func constructContext(d *schema.ResourceData) *Context {
	var context Context
	rawContext := d.Get("context").([]interface{})[0].(map[string]interface{})

	if rawAws := rawContext["aws"].([]interface{}); rawAws != nil && len(rawAws) > 0 {
		aws := rawAws[0].(map[string]interface{})
		accGroupIds := util.StringArrayOrNil(aws["account_group_ids"].([]interface{}))
		context = Context{
			Platform: util.StringAddress("aws"),
			AccountIds: &StringArrayOrNull{
				StringArray: util.StringArrayOrNil(aws["account_ids"].([]interface{})),
				ShowEmpty:   len(util.StringArrayOrNil(aws["account_group_ids"].([]interface{}))) > 0,
			},
			Regions:         &StringArrayOrNull{StringArray: util.StringArrayOrNil(aws["regions"].([]interface{}))},
			AccountGroupIds: accGroupIds,
		}
	} else if rawAzure := rawContext["azure"].([]interface{}); rawAzure != nil && len(rawAzure) > 0 {
		azure := rawAzure[0].(map[string]interface{})
		context = Context{
			Platform:        util.StringAddress("azure"),
			SubscriptionIds: &StringArrayOrNull{StringArray: util.StringArrayOrNil(azure["subscription_ids"].([]interface{}))},
		}
	} else {
		// not a supported platform - will error on API call
		return nil
	}

	return &context
}

func constructNotificationFieldMapping(fieldMappingInput map[string]interface{}) *NotificationFieldMapping {
	return &NotificationFieldMapping{
		MappingId:  util.StringAddress(fieldMappingInput["mapping_id"].(string)),
		Label:      util.StringAddress(fieldMappingInput["label"].(string)),
		Expression: util.StringAddress(fieldMappingInput["expression"].(string)),
	}
}

func constructNotificationFieldMappings(rawFieldMappings []interface{}) []*NotificationFieldMapping {
	var fieldMappings []*NotificationFieldMapping
	for _, rawFieldMapping := range rawFieldMappings {
		fieldMappings = append(fieldMappings, constructNotificationFieldMapping(rawFieldMapping.(map[string]interface{})))
	}

	return fieldMappings
}

func constructSGRuleChanges(rawSGRuleChanges []interface{}) []*SGRuleChanges {
	// TODO
	return nil
}

func constructNotification(notificationData map[string]interface{}) *Notification {
	notification := Notification{}
	rawSlackWebhooks := notificationData["slack_webhook"].([]interface{})
	if len(rawSlackWebhooks) > 0 {
		rawSlackWebhook := rawSlackWebhooks[0].(map[string]interface{})

		notification.Slack = &SlackNotificationConfig{
			RoomId: util.StringAddress(rawSlackWebhook["room_id"].(string)),
		}
	}

	rawSlackApps := notificationData["slack_app"].([]interface{})
	if len(rawSlackApps) > 0 {
		rawSlackApp := rawSlackApps[0].(map[string]interface{})

		notification.SlackApp = &SlackAppNotificationConfig{
			InstallationId: util.StringAddress(rawSlackApp["installation_id"].(string)),
		}
	}

	rawEmailConfigs := notificationData["email"].([]interface{})
	if len(rawEmailConfigs) > 0 {
		rawEmailConfig := rawEmailConfigs[0].(map[string]interface{})

		notification.Email = &EmailNotificationConfig{
			SendToTeam:      util.BoolAddress(rawEmailConfig["send_to_team"].(bool)),
			SendToUserGroup: util.BoolAddress(rawEmailConfig["send_to_user_group"].(bool)),
		}
		if addrs := util.ArrayOfStringPointers(rawEmailConfig["email_addresses"].([]interface{})); len(addrs) > 0 {
			notification.Email.EmailAddresses = addrs
		}
	}

	return &notification
}

func constructNotifications(rawNotifications []interface{}) *Notification {
	var notifications []*Notification
	for _, rawNotification := range rawNotifications {
		notifications = append(notifications, constructNotification(rawNotification.(map[string]interface{})))
	}
	return notifications[0]
}

func constructMatchFields(rawMatchFields []interface{}) *MatchFields {
	matchField := rawMatchFields[0].(map[string]interface{})
	return &MatchFields{
		EventName: util.ArrayOfStringPointers(matchField["event_name"].([]interface{})),
	}
}

func constructMatchExpression(rawMatchExpression []interface{}) *MatchExpression {
	if len(rawMatchExpression) > 0 {
		matchExpression := rawMatchExpression[0].(map[string]interface{})
		return &MatchExpression{
			Expression:         util.StringAddress(matchExpression["expression"].(string)),
			ExpressionLanguage: util.StringAddress("jmespath"),
		}
	}
	return nil
}

func constructTrigger(d *schema.ResourceData) *Trigger {
	var trigger Trigger
	rawTrigger := d.Get("trigger").([]interface{})[0].(map[string]interface{})
	for k, rawV := range rawTrigger {
		triggerTypeArr := rawV.([]interface{})
		if len(triggerTypeArr) == 0 {
			continue
		}
		v := triggerTypeArr[0].(map[string]interface{})
		switch k {
		case "schedule":
			trigger = Trigger{
				Trigger:               util.StringAddress("schedule"),
				Cron:                  util.StringAddress(v["cron"].(string)),
				Timezone:              util.StringAddress(v["timezone"].(string)),
				NotificationOffset:    util.IntAddress(v["notification_offset"].(int)),
				DefaultSnoozeDuration: util.IntAddress(v["default_snooze_duration"].(int)),
				Notifications:         constructNotifications(v["notifications"].([]interface{})),
			}
		case "cloudtrail_event":
			trigger = Trigger{
				Trigger:         util.StringAddress("cloudtrail_event"),
				MatchFields:     constructMatchFields(v["match_fields"].([]interface{})),
				MatchExpression: constructMatchExpression(v["match_expression"].([]interface{})),
			}
		case "manual":
			trigger = Trigger{
				Trigger: util.StringAddress("manual_trigger"),
			}
		default:
			break
		}
	}

	return &trigger
}

func getIndex(defn map[string]interface{}) int {
	return defn["index"].(int)
}

func constructAction(actionName string, defn map[string]interface{}) *Action {
	var action Action
	switch actionName {
	/* AWS Actions */
	case "check_tag_compliance":
		action = Action{
			Action:               &actionName,
			TagTargeted:          util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:            util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:     util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			ResourceTypes:        &StringArrayOrNull{StringArray: util.StringArrayOrNil(defn["resource_types"].([]interface{}))},
			ReportType:           util.StringAddress(defn["report_type"].(string)),
			NotificationsTrigger: util.StringAddress(defn["notifications_trigger"].(string)),
			Notifications:        constructNotifications(defn["notifications"].([]interface{})),
		}
	case "copy_db_snapshots":
		action = Action{
			Action:            &actionName,
			Operator:          util.StringAddress(defn["operator"].(string)),
			Value:             util.IntAddress(defn["value"].(int)),
			DestinationRegion: util.StringAddress(defn["destination_region"].(string)),
			Mode:              util.StringAddress(defn["mode"].(string)),
			KeyMapping:        util.MapAddress(defn["key_mapping"].(map[string]string)),
			UseDefaultKmsKey:  util.BoolAddress(defn["use_default_kms_key"].(bool)),
			TagTargeted:       util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:         util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:  util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "copy_snapshots":
		action = Action{
			Action:            &actionName,
			Operator:          util.StringAddress(defn["operator"].(string)),
			Value:             util.IntAddress(defn["value"].(int)),
			DestinationRegion: util.StringAddress(defn["destination_region"].(string)),
			Mode:              util.StringAddress(defn["mode"].(string)),
			KeyMapping:        util.MapAddress(defn["key_mapping"].(map[string]string)),
			UseDefaultKmsKey:  util.BoolAddress(defn["use_default_kms_key"].(bool)),
			TagTargeted:       util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:         util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:  util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "create_db_snapshots":
		action = Action{
			Action:             &actionName,
			TagTargeted:        util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:          util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:   util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			CopyDbInstanceTags: util.BoolAddress(defn["copy_db_instance_tags"].(bool)),
			MultiAzOnly:        util.BoolAddress(defn["multi_az_only"].(bool)),
			AdditionalTags:     util.ArrayOfMapsAddress(defn["additional_tags"].([]interface{})),
		}
	case "create_images":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			NoReboot:         util.BoolAddress(defn["no_reboot"].(bool)),
			CopyVolumeTags:   util.BoolAddress(defn["copy_volume_tags"].(bool)),
			CopyInstanceTags: util.BoolAddress(defn["copy_instance_tags"].(bool)),
			AdditionalTags:   util.ArrayOfMapsAddress(defn["additional_tags"].([]interface{})),
		}
	case "create_snapshots":
		action = Action{
			Action:           &actionName,
			TagTargeted:      util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			CopyVolumeTags:   util.BoolAddress(defn["copy_volume_tags"].(bool)),
			CopyInstanceTags: util.BoolAddress(defn["copy_instance_tags"].(bool)),
			AdditionalTags:   util.ArrayOfMapsAddress(defn["additional_tags"].([]interface{})),
		}
	case "create_vss_snapshots":
		action = Action{
			Action:            &actionName,
			TagTargeted:       util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:         util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:  util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			CopyVolumeTags:    util.BoolAddress(defn["copy_volume_tags"].(bool)),
			CopyInstanceTags:  util.BoolAddress(defn["copy_instance_tags"].(bool)),
			ExcludeBootVolume: util.BoolAddress(defn["exclude_boot_volume"].(bool)),
			UseAdditionalTags: util.BoolAddress(defn["use_additional_tags"].(bool)),
			AdditionalTags:    util.ArrayOfMapsAddress(defn["additional_tags"].([]interface{})),
		}
	case "delete_detached_volumes":
		action = Action{
			Action:           &actionName,
			DryRun:           util.BoolAddress(defn["dry_run"].(bool)),
			DaysDetached:     util.IntAddress(defn["days_detached"].(int)),
			TagTargeted:      util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "delete_images":
		action = Action{
			Action:           &actionName,
			Operator:         util.StringAddress(defn["operator"].(string)),
			Value:            util.IntAddress(defn["value"].(int)),
			KeepLatest:       util.BoolAddress(defn["keep_latest"].(bool)),
			TagTargeted:      util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "delete_orphaned_snapshots":
		action = Action{
			Action:           &actionName,
			DryRun:           util.BoolAddress(defn["dry_run"].(bool)),
			TagTargeted:      util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "delete_snapshots":
		action = Action{
			Action:           &actionName,
			DryRun:           util.BoolAddress(defn["dry_run"].(bool)),
			Operator:         util.StringAddress(defn["operator"].(string)),
			Value:            util.IntAddress(defn["value"].(int)),
			KeepLatest:       util.BoolAddress(defn["keep_latest"].(bool)),
			KeepByVolume:     util.BoolAddress(defn["keep_by_volume"].(bool)),
			TagTargeted:      util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "ec2_command_run_powershell_script":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Commands:         util.ArrayOfStringPointers(defn["commands"].([]interface{})),
			WorkingDirectory: util.StringAddress(defn["working_directory"].(string)),
			ExecutionTimeout: util.IntAddress(defn["execution_timeout"].(int)),
		}
	case "ec2_command_run_shell_script":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Commands:         util.ArrayOfStringPointers(defn["commands"].([]interface{})),
			WorkingDirectory: util.StringAddress(defn["working_directory"].(string)),
			ExecutionTimeout: util.IntAddress(defn["execution_timeout"].(int)),
		}
	case "invoke_named_lambda_function":
		action = Action{
			Action:                 &actionName,
			FunctionName:           util.StringAddress(defn["function_name"].(string)),
			InvocationType:         util.StringAddress(defn["invocation_type"].(string)),
			Payload:                util.StringAddress(defn["payload"].(string)),
			ReplaceConflictingVars: util.BoolAddress(defn["replace_conflicting_vars"].(bool)),
			EnvironmentVariables:   util.ArrayOfMapsAddress(defn["environment_variables"].([]interface{})),
		}
	case "invoke_tagged_lambda_functions":
		action = Action{
			Action:                 &actionName,
			TagGroups:              util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:       util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			InvocationType:         util.StringAddress(defn["invocation_type"].(string)),
			Payload:                util.StringAddress(defn["payload"].(string)),
			ReplaceConflictingVars: util.BoolAddress(defn["replace_conflicting_vars"].(bool)),
			EnvironmentVariables:   util.ArrayOfMapsAddress(defn["environment_variables"].([]interface{})),
		}
	case "notify_cost":
		action = Action{
			Action:        &actionName,
			Service:       util.StringAddress(defn["service"].(string)),
			Notifications: constructNotifications(defn["notifications"].([]interface{})),
		}
	case "notify_event":
		action = Action{
			Action:                    &actionName,
			NotificationFieldMappings: constructNotificationFieldMappings(defn["notification_field_mapping"].([]interface{})),
			Notifications:             constructNotifications(defn["notifications"].([]interface{})),
		}
	case "notify_instance_count":
		action = Action{
			Action:        &actionName,
			Notifications: constructNotifications(defn["notifications"].([]interface{})),
		}
	case "reboot_instances":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "release_disassociated_ips":
		action = Action{
			Action:            &actionName,
			DaysDisassociated: util.IntAddress(defn["days_disassociated"].(int)),
			TagTargeted:       util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:         util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:  util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "resume_autoscaling_processes":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Processes:        util.ArrayOfStringPointers(defn["processes"].([]interface{})),
		}
	case "start_instances":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Wait: &Wait{
				InstanceState:  util.BoolAddress(defn["wait_instance_state"].(bool)),
				InstanceStatus: util.BoolAddress(defn["wait_instance_status"].(bool)),
				SystemStatus:   util.BoolAddress(defn["wait_system_status"].(bool)),
			},
		}
	case "start_rds_instances":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			TargetClusters:   util.BoolAddress(defn["target_clusters"].(bool)),
			Wait: &Wait{
				InstanceState:  util.BoolAddress(defn["wait_instance_state"].(bool)),
				InstanceStatus: util.BoolAddress(defn["wait_instance_status"].(bool)),
				SystemStatus:   util.BoolAddress(defn["wait_system_status"].(bool)),
			},
		}
	case "start_workspaces":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "stop_instances":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Wait: &Wait{
				InstanceState:  util.BoolAddress(defn["wait_instance_state"].(bool)),
				InstanceStatus: util.BoolAddress(defn["wait_instance_status"].(bool)),
				SystemStatus:   util.BoolAddress(defn["wait_system_status"].(bool)),
			},
		}
	case "stop_rds_instances":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			TargetClusters:   util.BoolAddress(defn["target_clusters"].(bool)),
			Wait: &Wait{
				InstanceState:  util.BoolAddress(defn["wait_instance_state"].(bool)),
				InstanceStatus: util.BoolAddress(defn["wait_instance_status"].(bool)),
				SystemStatus:   util.BoolAddress(defn["wait_system_status"].(bool)),
			},
		}
	case "stop_workspaces":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "suspend_autoscaling_processes":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Processes:        util.ArrayOfStringPointers(defn["processes"].([]interface{})),
		}
	case "update_autoscaling_groups":
		action = Action{
			Action:                       &actionName,
			TagGroups:                    util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:             util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			StoreExistingAsgSettings:     util.BoolAddress(defn["store_existing_asg_settings"].(bool)),
			RestoreToPreviousAsgSettings: util.BoolAddress(defn["restore_to_previous_asg_settings"].(bool)),
			IgnoreIfNoCachedAsgSettings:  util.BoolAddress(defn["ignore_if_no_cached_asg_settings"].(bool)),
			Params: &AutoscalingParams{
				Min:     util.IntAddress(defn["min"].(int)),
				Max:     util.IntAddress(defn["max"].(int)),
				Desired: util.IntAddress(defn["desired"].(int)),
			},
		}
	case "update_dynamodb_table_throughput":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			ReadUnits:        util.IntAddress(defn["read_units"].(int)),
			WriteUnits:       util.IntAddress(defn["write_units"].(int)),
		}
	case "update_ecs_service_scale":
		action = Action{
			Action:                        &actionName,
			TagGroups:                     util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:              util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			DesiredCount:                  util.IntAddress(defn["desired_count"].(int)),
			StoredExistingDesiredCount:    util.BoolAddress(defn["store_existing_desired_count"].(bool)),
			RestoreToPreviousDesiredCount: util.BoolAddress(defn["restore_to_previous_desired_count"].(bool)),
			IgnoreIfNoCachedDesiredCount:  util.BoolAddress(defn["ignore_if_no_cached_desired_count"].(bool)),
		}
	case "update_provisioned_iops":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Iops:             util.IntAddress(defn["iops"].(int)),
		}
	case "update_security_groups":
		action = Action{
			Action:           &actionName,
			TagTargeted:      util.BoolAddress(defn["tag_targeted"].(bool)),
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			RuleChanges:      constructSGRuleChanges(defn["rule_changes"].([]interface{})),
		}
	case "update_application_autoscaling_settings":
		action = Action{
			Action:                              &actionName,
			TagGroups:                           util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner:                    util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			ScalableDimension:                   util.StringAddress(defn["scalable_dimension"].(string)),
			ServiceNamespace:                    util.StringAddress(defn["service_namespace"].(string)),
			MinCapacity:                         util.IntAddress(defn["min_capacity"].(int)),
			MaxCapacity:                         util.IntAddress(defn["max_capacity"].(int)),
			StoreExistingAutoscalingSettings:    util.BoolAddress(defn["store_existing_autoscaling_settings"].(bool)),
			RestoreExistingAutoscalingSettings:  util.BoolAddress(defn["restore_existing_autoscaling_settings"].(bool)),
			IgnoreIfNoCachedAutoscalingSettings: util.BoolAddress(defn["ignore_if_no_cached_autoscaling_settings"].(bool)),
		}
		/* Azure Actions */
	case "deallocate_vms":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "restore_sql_databases":
		action = Action{
			Action:         &actionName,
			DatabaseName:   util.StringAddress(defn["database_name"].(string)),
			DatabaseServer: util.StringAddress(defn["database_server"].(string)),
			ResourceGroup:  util.StringAddress(defn["resource_group"].(string)),
		}
	case "start_wvd_session_hosts":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "start_sql_databases":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "start_container_groups":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "start_vms":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "stop_container_groups":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "stop_sql_databases":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "stop_wvd_session_hosts":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}
	case "update_aks_node_pool_scale":
		var minCount = util.IntAddress(defn["min_count"].(int))
		var maxCount = util.IntAddress(defn["max_count"].(int))
		var restoreToPrevScale = util.BoolAddress(defn["restore_to_previous_scale"].(bool))

		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}

		if restoreToPrevScale == nil || *restoreToPrevScale == false {
			action.Params = &AutoscalingParams{
				MinCount: minCount,
				MaxCount: maxCount,
			}
		} else {
			action.RestoreToPreviousScale = restoreToPrevScale
		}
	case "update_cosmos_container_throughput":
		fallthrough
	case "update_cosmos_table_throughput":
		var throughput = util.IntAddress(defn["throughput"].(int))
		var restoreToPreviousThroughput = util.BoolAddress(defn["restore_to_previous_throughput"].(bool))

		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
		}

		if restoreToPreviousThroughput == nil || *restoreToPreviousThroughput == false {
			action.Params = &AutoscalingParams{
				Throughput: throughput,
			}
		} else {
			action.RestoreToPreviousThroughput = restoreToPreviousThroughput
		}
	case "update_scale_sets":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Capacity:         util.IntAddress(defn["capacity"].(int)),
		}
	case "update_autoscale_settings":
		action = Action{
			Action:           &actionName,
			TagGroups:        util.ArrayOfStringPointers(defn["tag_groups"].([]interface{})),
			TagGroupCombiner: util.GetTagGroupCombiner(defn["tag_group_combiner"].(string)),
			Params: &AutoscalingParams{
				Min:     util.IntAddress(defn["min"].(int)),
				Max:     util.IntAddress(defn["max"].(int)),
				Desired: util.IntAddress(defn["desired"].(int)),
			},
		}
		/* Pauses */
	case "delay":
		action = Action{
			Action:       &actionName,
			WaitDuration: util.IntAddress(defn["wait_duration"].(int)),
		}
	case "manual_approval":
		action = Action{
			Action: &actionName,
		}
	}

	return &action
}

// We cant just go off the number of entries in the map, as
// any sequences of actions with more than 1 of the same action
// type will go into an array of definitions
func actionCount(rawActions map[string]interface{}) int {
	var count int
	for _, arrayOfDefns := range rawActions {
		count += len(arrayOfDefns.([]interface{}))
	}

	return count
}

func constructActions(d *schema.ResourceData) []*Action {
	rawActions := d.Get("actions").([]interface{})[0].(map[string]interface{})
	actions := make([]*Action, actionCount(rawActions))
	smallestIndex := util.GetSmallestArrayIndex(rawActions)

	for actionName, rawArrayOfMaps := range rawActions {
		arrayOfMaps := util.ConvertToArrayOfMaps(rawArrayOfMaps.([]interface{}))
		for _, defn := range arrayOfMaps {
			action := constructAction(actionName, defn)
			index := getIndex(defn)
			actions[index-smallestIndex] = action
		}
	}

	return actions
}

func constructRuleFromResourceData(d *schema.ResourceData, teamId string) *Rule {
	return &Rule{
		TeamId:    &teamId,
		Name:      util.StringAddress(d.Get("name").(string)),
		Slug:      util.StringAddress(d.Get("slug").(string)),
		Enabled:   util.BoolAddress(d.Get("enabled").(bool)),
		Labels:    util.ArrayOfStringPointers(d.Get("labels").([]interface{})),
		UserGroup: util.StringAddress(d.Get("user_group").(string)),

		Context: constructContext(d),
		Trigger: constructTrigger(d),
		Actions: constructActions(d),
	}
}

func resourceRuleCreate(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	rule, err := client.CreateRule(teamId, constructRuleFromResourceData(d, teamId))

	if err != nil {
		return err
	}

	d.SetId(*rule.Id)
	return resourceRuleRead(d, m)
}

func resourceRuleRead(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	ruleId := d.Id()
	rule, err := client.GetRule(teamId, ruleId)

	if err != nil {
		return err
	}

	d.Set("_id", ruleId)
	d.Set("name", rule.Name)
	d.Set("slug", rule.Slug)
	d.Set("created_at", rule.CreatedAt)
	d.Set("updated_at", rule.UpdatedAt)
	d.Set("created_by", rule.CreatedBy)
	d.Set("team_id", rule.TeamId)
	d.Set("enabled", rule.Enabled)
	d.Set("user_group", rule.UserGroup)
	d.Set("labels", rule.Labels)
	return nil
}

func resourceRuleUpdate(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	ruleId := d.Id()
	_, err := client.UpdateRule(teamId, ruleId, constructRuleFromResourceData(d, teamId))

	if err != nil {
		return err
	}
	return resourceRuleRead(d, m)
}

func resourceRuleDelete(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	ruleId := d.Id()

	err := client.DeleteRule(teamId, ruleId)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRuleCreate,
		Read:   resourceRuleRead,
		Update: resourceRuleUpdate,
		Delete: resourceRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: ruleSchema(),
	}
}
