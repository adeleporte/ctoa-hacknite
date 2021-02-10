package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePeople() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePeopleCreate,
		ReadContext:   resourcePeopleRead,
		UpdateContext: resourcePeopleUpdate,
		DeleteContext: resourcePeopleDelete,
		Schema: map[string]*schema.Schema{
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePeopleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*Client)

	ctoa := Ctoa{
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
	}

	id, err := CreateCtoa(client, ctoa)
	if err != nil {
		return diags
	}

	d.SetId(id)
	return diags

}

func resourcePeopleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)
	id := d.Id()

	ctoa, err := ReadCtoa(client, id)
	if err != nil {
		return diags
	}

	d.Set("first_name", ctoa.FirstName)
	d.Set("last_name", ctoa.LastName)

	return diags
}

func resourcePeopleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*Client)

	ctoa := Ctoa{
		ID:        d.Id(),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
	}

	updatedCtoa, err := UpdateCtoa(client, ctoa)
	if err != nil {
		return diags
	}

	d.SetId(updatedCtoa.ID)
	d.Set("first_name", updatedCtoa.FirstName)
	d.Set("last_name", updatedCtoa.LastName)

	return resourcePeopleRead(ctx, d, m)
}

func resourcePeopleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*Client)
	id := d.Id()

	id, err := DeleteCtoa(client, id)
	if err != nil {
		d.SetId(id)
		return diags
	}

	return diags
}
