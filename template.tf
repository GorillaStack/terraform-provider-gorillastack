variable "api_key" {}
variable "team_id" {}

provider "gorillastack" {
  api_key = "${var.api_key}"
  team_id = "${var.team_id}"
}

resource "gorillastack_tag_group" "aashwin-tf-test" {
  name           = "aash-dev-test2"
  tag_expression = "i \"something-that-will\":\"never match\""
}

resource "gorillastack_tag_group" "aashwin-start-workspaces" {
  name           = "aash-dev-test"
  tag_expression = "i \"something-that-will\":\"never match\""
}

resource "gorillastack_tag_group" "aashwin-supertest" {
  name           = "aash-dev-supertest"
  tag_expression = "i \"application\":\"crm\" and not i \"environment\":\"production\""
}

resource "gorillastack_rule" "aashwin-stop-workspaces-test" {
  name    = "Aashwin test - stop workspaces"
  labels  = ["terraform", "test", "local"]
  enabled = false

  context {
    aws {} # should mean all accounts all regions
  }

  trigger {
    manual {}
  }

  actions {
    stop_workspaces {
      index      = 1
      tag_groups = ["${gorillastack_tag_group.aashwin-tf-test.id}"]
    }

    start_workspaces {
      index      = 2
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"]
    }

    update_application_autoscaling_settings {
      index      = 3
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"] 
      scalable_dimension  = "ecs:service:DesiredCount"
      service_namespace = "ecs"
      min_capacity = 1
      max_capacity = 3
      restore_existing_autoscaling_settings = true
      ignore_if_no_cached_autoscaling_settings = true
    }

    # restore_sql_database { 
    #   #   # index      = 3 
    #   database_name = 3
    #   database_server = "testing"
    #   resource_group = "testing"
    # }
  }
}

resource "gorillastack_rule" "aashwin-azure-rule-test" {
  name    = "Aashwin test - sql actions"
  labels  = ["terraform", "test", "local"]
  enabled = false

  context {
    azure {
    }

  }

  trigger {
    manual {}
  }

  actions {
    # stop_workspaces {
    #   index      = 1
    #   tag_groups = ["${gorillastack_tag_group.aashwin-tf-test.id}"]
    # }

    start_wvd_session_hosts {
      index      = 1
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}", "${gorillastack_tag_group.aashwin-tf-test.id}"]
      tag_group_combiner = "or"
    }

    start_container_groups {
      index      = 2
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"]
    }

    start_sql_databases {
      index      = 3
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"]

    }
    start_workspaces {
      index      = 4
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"]

    }
    stop_workspaces {
      index      = 5
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"]

    }

    stop_container_groups {
      index      = 6
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"]

    }

    stop_sql_databases {
      index      = 7
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"] 
     } 
     
    # start_vms { #   index      = 1
    #   tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"]
    # }


    restore_sql_databases { 
      index      = 8 
      database_name = "steven-test-sql-database"
      database_server = "steven-test-sql-server"
      resource_group = "STEVEN-RESOURCE-GROUP"
    }

    update_aks_node_pool_scale {
      index      = 9
      tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"] 
      # min_count = 1 
      # max_count = 8
      restore_to_previous_scale = true
    }

    # update_cosmos_container_throughput {
    #   index      = 10
    #   tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"] 
    #   minimum = 400
    #   maximum = 1000000
    #   multiple_of = 100
    # }

    # update_application_autoscale_settings {
    #   index      = 9
    #   tag_groups = ["${gorillastack_tag_group.aashwin-start-workspaces.id}"] 
      
    # }
  }
}

