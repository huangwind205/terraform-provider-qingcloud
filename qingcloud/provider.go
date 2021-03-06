/**
 * Copyright (c) 2016 Magicshui
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
/**
 * Copyright (c) 2017 yunify
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package qingcloud

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["access_key"],
			},
			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["secret_key"],
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["zone"],
			},
			"endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"qingcloud_vpn_cert": dataSourceQingcloudVpnCert(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"qingcloud_eip":                   resourceQingcloudEip(),
			"qingcloud_keypair":               resourceQingcloudKeypair(),
			"qingcloud_security_group":        resourceQingcloudSecurityGroup(),
			"qingcloud_security_group_rule":   resourceQingcloudSecurityGroupRule(),
			"qingcloud_vxnet":                 resourceQingcloudVxnet(),
			"qingcloud_vpc":                   resourceQingcloudVpc(),
			"qingcloud_instance":              resourceQingcloudInstance(),
			"qingcloud_volume":                resourceQingcloudVolume(),
			"qingcloud_tag":                   resourceQingcloudTag(),
			"qingcloud_vpc_static":            resourceQingcloudVpcStatic(),
			"qingcloud_loadbalancer":          resourceQingcloudLoadBalancer(),
			"qingcloud_loadbalancer_listener": resourceQingcloudLoadBalancerListener(),
			"qingcloud_loadbalancer_backend":  resourceQingcloudLoadBalancerBackend(),
			"qingcloud_server_certificate":    resourceQingcloudServerCertificate(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	accesskey, ok := d.GetOk("access_key")
	if !ok {
		accesskey = os.Getenv("QINGCLOUD_ACCESS_KEY")
	}
	secretkey, ok := d.GetOk("secret_key")
	if !ok {
		secretkey = os.Getenv("QINGCLOUD_SECRET_KEY")
	}
	zone, ok := d.GetOk("zone")
	if !ok {
		zone = os.Getenv("QINGCLOUD_ZONE")
		if zone == "" {
			zone = DEFAULT_ZONE
		}
	}
	endpoint, ok := d.GetOk("endpoint")
	if !ok {
		endpoint = os.Getenv("QINGCLOUD_ENDPOINT")
		if endpoint == "" {
			endpoint = DEFAULT_ENDPOINT
		}
	}

	config := Config{
		ID:       accesskey.(string),
		Secret:   secretkey.(string),
		Zone:     zone.(string),
		EndPoint: endpoint.(string),
	}
	return config.Client()
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key": "qingcloud access key ID ",
		"secret_key": "qingcloud access key secret",
		"zone":       "qingcloud reigon zone",
		"endpoint":   "qingcloud api endpoint",
	}
}
