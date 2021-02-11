package client

import (
	"log"

	"github.com/BESTSELLER/terraform-provider-harbor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetExecutionBody(d *schema.ResourceData) models.ExecutionBody {

	body := models.ExecutionBody{
		PolicyId:     d.Get("policy_id").(int),
		Id:			  d.Get("policy_id").(int),
	}

	log.Println(body)
	return body
}
