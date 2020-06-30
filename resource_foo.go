package main

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFoo() *schema.Resource {
	return &schema.Resource{
		Create: resourceFooCreate,
		Read:   resourceFooRead,
		Update: resourceFooUpdate,
		Delete: resourceFooDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},

			"value": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Sensitive:    true,
			},

			"end_date": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
			},
		},
	}
}

func resourceFooCreate(d *schema.ResourceData, m interface{}) error {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return err
	}

	d.SetId(id)
	return resourceFooRead(d, m)
}

func resourceFooRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFooUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFooRead(d, m)
}

func resourceFooDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}