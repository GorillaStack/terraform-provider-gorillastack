package gorillastack

import (
	// "fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed values
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"date_created": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, meta interface{}) error {
	// name := d.Get("name").(string)

	// client := meta.(*Client)
	// user, err := client.GetUserByName(name)
	// if err != nil {
	// 	return err
	// }
	// if user == nil {
	// 	d.SetId("")
	// 	return nil
	// }

	// id := fmt.Sprintf("%d", user._id)
	// d.SetId(id)
	// d.Set("email", user.emails[0].address)
	// d.Set("created_at", user.createdAt)

	return nil
}
