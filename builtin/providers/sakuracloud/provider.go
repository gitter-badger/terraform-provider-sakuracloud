package sakuracloud

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var DefaultZone = "is1b"

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SAKURACLOUD_ACCESS_TOKEN", nil),
				Description: "Your SakuraCloud APIKey(token)",
			},
			"secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SAKURACLOUD_ACCESS_TOKEN_SECRET", nil),
				Description: "Your SakuraCloud APIKey(secret)",
			},
			"zone": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.MultiEnvDefaultFunc([]string{"SAKURACLOUD_ZONE"}, nil),
				Description:  "Target SakuraCloud Zone(is1a | is1b | tk1a | tk1v)",
				InputDefault: DefaultZone,
				ValidateFunc: validateStringInWord([]string{"is1a", "is1b", "tk1a", "tk1v"}),
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SAKURACLOUD_TIMEOUT"}, 20),
			},
			"trace": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SAKURACLOUD_TRACE_MODE", false),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sakuracloud_archive":        dataSourceSakuraCloudArchive(),
			"sakuracloud_bridge":         dataSourceSakuraCloudBridge(),
			"sakuracloud_cdrom":          dataSourceSakuraCloudCDROM(),
			"sakuracloud_database":       dataSourceSakuraCloudDatabase(),
			"sakuracloud_disk":           dataSourceSakuraCloudDisk(),
			"sakuracloud_dns":            dataSourceSakuraCloudDNS(),
			"sakuracloud_gslb":           dataSourceSakuraCloudGSLB(),
			"sakuracloud_internet":       dataSourceSakuraCloudInternet(),
			"sakuracloud_load_balancer":  dataSourceSakuraCloudLoadBalancer(),
			"sakuracloud_note":           dataSourceSakuraCloudNote(),
			"sakuracloud_packet_filter":  dataSourceSakuraCloudPacketFilter(),
			"sakuracloud_simple_monitor": dataSourceSakuraCloudSimpleMonitor(),
			"sakuracloud_server":         dataSourceSakuraCloudServer(),
			"sakuracloud_ssh_key":        dataSourceSakuraCloudSSHKey(),
			"sakuracloud_switch":         dataSourceSakuraCloudSwitch(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"sakuracloud_auto_backup":                    resourceSakuraCloudAutoBackup(),
			"sakuracloud_bridge":                         resourceSakuraCloudBridge(),
			"sakuracloud_database":                       resourceSakuraCloudDatabase(),
			"sakuracloud_disk":                           resourceSakuraCloudDisk(),
			"sakuracloud_dns":                            resourceSakuraCloudDNS(),
			"sakuracloud_dns_record":                     resourceSakuraCloudDNSRecord(),
			"sakuracloud_gslb":                           resourceSakuraCloudGSLB(),
			"sakuracloud_gslb_server":                    resourceSakuraCloudGSLBServer(),
			"sakuracloud_internet":                       resourceSakuraCloudInternet(),
			"sakuracloud_load_balancer":                  resourceSakuraCloudLoadBalancer(),
			"sakuracloud_load_balancer_vip":              resourceSakuraCloudLoadBalancerVIP(),
			"sakuracloud_load_balancer_server":           resourceSakuraCloudLoadBalancerServer(),
			"sakuracloud_note":                           resourceSakuraCloudNote(),
			"sakuracloud_packet_filter":                  resourceSakuraCloudPacketFilter(),
			"sakuracloud_simple_monitor":                 resourceSakuraCloudSimpleMonitor(),
			"sakuracloud_server":                         resourceSakuraCloudServer(),
			"sakuracloud_ssh_key":                        resourceSakuraCloudSSHKey(),
			"sakuracloud_switch":                         resourceSakuraCloudSwitch(),
			"sakuracloud_vpc_router":                     resourceSakuraCloudVPCRouter(),
			"sakuracloud_vpc_router_interface":           resourceSakuraCloudVPCRouterInterface(),
			"sakuracloud_vpc_router_firewall":            resourceSakuraCloudVPCRouterFirewall(),
			"sakuracloud_vpc_router_dhcp_server":         resourceSakuraCloudVPCRouterDHCPServer(),
			"sakuracloud_vpc_router_dhcp_static_mapping": resourceSakuraCloudVPCRouterDHCPStaticMapping(),
			"sakuracloud_vpc_router_port_forwarding":     resourceSakuraCloudVPCRouterPortForwarding(),
			"sakuracloud_vpc_router_pptp":                resourceSakuraCloudVPCRouterPPTP(),
			"sakuracloud_vpc_router_l2tp":                resourceSakuraCloudVPCRouterL2TP(),
			"sakuracloud_vpc_router_static_nat":          resourceSakuraCloudVPCRouterStaticNAT(),
			"sakuracloud_vpc_router_user":                resourceSakuraCloudVPCRouterRemoteAccessUser(),
			"sakuracloud_vpc_router_site_to_site_vpn":    resourceSakuraCloudVPCRouterSiteToSiteIPsecVPN(),
			"sakuracloud_vpc_router_static_route":        resourceSakuraCloudVPCRouterStaticRoute(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	if _, ok := d.GetOk("zone"); !ok {
		d.Set("zone", DefaultZone)
	}

	config := Config{
		AccessToken:       d.Get("token").(string),
		AccessTokenSecret: d.Get("secret").(string),
		Zone:              d.Get("zone").(string),
		TimeoutMinute:     d.Get("timeout").(int),
		TraceMode:         d.Get("trace").(bool),
	}

	return config.NewClient(), nil
}

var sakuraMutexKV = mutexkv.NewMutexKV()
