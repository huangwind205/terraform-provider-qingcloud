package qingcloud

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["token"],
			},
			"secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["secret"],
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["zone"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			// "qingcloud_eip":     resourceQingcloudEip(),
			"qingcloud_keypair":       resourceQingcloudKeypair(),
			"qingcloud_securitygroup": resourceQingcloudSecuritygroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var qingcloudMutexKV = mutexkv.NewMutexKV()

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ID:     d.Get("id").(string),
		Secret: d.Get("secret").(string),
		Zone:   d.Get("zone").(string),
	}
	return config.Client()
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"id":     "青云的 ID ",
		"secret": "青云的密钥",
		"zone":   "青云的 zone ",
	}
}
