Terraform Provider for GorillaStack
==================

[![Release](https://img.shields.io/github/tag/GorillaStack/terraform-provider-gorillastack.svg?style=for-the-badge)](https://github.com/GorillaStack/terraform-provider-gorillastack/releases/latest)
[![Software License](https://img.shields.io/badge/license-MPL-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Build Status](https://img.shields.io/travis/GorillaStack/terraform-provider-gorillastack/master.svg?style=for-the-badge)](https://travis-ci.org/GorillaStack/terraform-provider-gorillastack)
[![Powered By: GorillaStack](https://img.shields.io/badge/powered%20by-GorillaStack-green.svg?style=for-the-badge)](https://www.gorillastack.com)

- Terraform Website: https://www.terraform.io
- GorillaStack Website: https://www.gorillastack.com
<!-- - Documentation: https://www.terraform.io/docs/providers/gorillastack/index.html -->

Maintainers
-----------

This provider plugin is maintained by:

* The unusually friendly team at [GorillaStack](https://www.gorillastack.com/)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10+
-	[Go](https://golang.org/doc/install) 1.11.0 or higher


Install the Provider
--------------------

```bash
go get -d github.com/gorillastack/terraform-provider-gorillastack
cd $GOPATH/src/github.com/gorillastack/terraform-provider-gorillastack
make
```

Using the Provider
------------------

Configure the provider using a .tfvars file. Here is what your terraform.tfvars file should look like.

```
# Your GorillaStack API Key
api_key = "<api_key>"
# Your GorillaStack Team's Id
team_id = "<gorillastack team_id>"
```

We recommend adding `*.tfvars` as an entry in your `.gitignore` file if you plan to commit your templates to a repository.
We do not recommend hardcoding API keys and other information in your templates.

Then, define Rule and Tag Group resources in your templates.

```terraform
variable "api_key" {}
variable "team_id" {}

provider "gorillastack" {
  api_key = "${var.api_key}"
  team_id = "${var.team_id}"
}

resource "gorillastack_tag_group" "ec2_instance_targets" {
  name    = "EC2 Instance Targets"
  tag_expression = "i \"application\":\"crm\" and not /environment/:/production/"
}

resource "gorillastack_tag_group" "autoscaling_group_targets" {
  name    = "AutoScaling Group Targets"
  tag_expression = "i \"type\":\"api\" and not /environment/:/production/"
}

resource "gorillastack_tag_group" "rds_targets" {
  name    = "RDS Targets"
  tag_expression = "not /environment/:/production/"
}

resource "gorillastack_rule" "test" {
  name      = "test rules creation from terraform"
  labels    = ["terraform", "test", "local"]
  enabled   = false

  context {
    aws {
    }  # should mean all accounts all regions
  }

  trigger {
    schedule {
      cron                = "0 0 9 1 1"
      timezone            = "Australia/Sydney"
      notification_offset = 30
      notifications {
        slack_webhook {
          room_id         = "5cfe861335112656bda2f80a"
        }
      }
    }
  }

  actions {
    delete_detached_volumes {
      index         = 1
      dry_run       = true
      tag_targeted  = false
      days_detached = 0
    }

    stop_instances {
      index         = 2
      tag_groups    = ["${gorillastack_tag_group.ec2_instance_targets.id}"]
    }

    update_autoscaling_groups {
      index         = 3
      tag_groups    = ["${gorillastack_tag_group.autoscaling_group_targets.id}"]
      min           = 0
      max           = 0
      desired       = 0
      store_existing_asg_settings = true
    }

    delay {
      index         = 4
      wait_duration = 10
    }

    stop_rds_instances {
      index         = 5
      tag_groups    = ["${gorillastack_tag_group.rds_targets.id}"]
    }
  }
}
```


Upgrading the provider
----------------------

The GorillaStack provider doesn't upgrade automatically once you've started using it. After a new release you can run 

```bash
go get -d github.com/gorillastack/terraform-provider-gorillastack
cd $GOPATH/src/github.com/gorillastack/terraform-provider-gorillastack
make
```

to upgrade to the latest stable version of the GorillaStack provider.

Building the provider
---------------------

Clone the repository. In the steps above, `go get -d` installs the repository to: `$GOPATH/src/github.com/gorillastack/terraform-provider-gorillastack`

```sh
$ mkdir -p $GOPATH/src/github.com/gorillastack; cd $GOPATH/src/github.com/gorillastack
$ git clone git@github.com:gorillastack/terraform-provider-gorillastack
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/gorillastack/terraform-provider-gorillastack
$ make
```

Contributing/Developing the provider
------------------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11.4+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

For guidance on common development practices such as testing changes or vendoring libraries, see the [contribution guidelines](https://github.com/gorillastack/terraform-provider-gorillastack/blob/master/CONTRIBUTING.md). If you have other development questions we don't cover, please file an issue!
