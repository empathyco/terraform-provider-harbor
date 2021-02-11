package provider

import (
	"fmt"

	"github.com/BESTSELLER/terraform-provider-harbor/client"
	"github.com/BESTSELLER/terraform-provider-harbor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExecution() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
		Create: resourceExecutionCreate,
		Read:   resourceExecutionRead,
		Update: resourceExecutionUpdate,
		Delete: resourceExecutionDelete,
	}
}

func resourceExecutionCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	body := client.GetExecutionBody(d)

	_, headers, err := apiClient.SendRequest("POST", models.PathExecution, body, 201)
	if err != nil {
		return err
	}

	id, err := client.GetID(headers)
	if err != nil {
		return err
	}

	d.SetId(id)
	return resourceExecutionRead(d, m)
}

func resourceExecutionRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	_, _, err := apiClient.SendRequest("GET", d.Id(), nil, 200)
	if err != nil {
		return fmt.Errorf("Resource not found %s", d.Id())
	}

	return nil
}

func resourceExecutionUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceExecutionDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
