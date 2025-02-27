---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_listener"
sidebar_current: "docs-oci-resource-network_load_balancer-listener"
description: |-
  Provides the Listener resource in Oracle Cloud Infrastructure Network Load Balancer service
---

# oci_network_load_balancer_listener
This resource provides the Listener resource in Oracle Cloud Infrastructure Network Load Balancer service.

Adds a listener to a network load balancer.

## Example Usage

```hcl
resource "oci_network_load_balancer_listener" "test_listener" {
	#Required
	default_backend_set_name = oci_network_load_balancer_backend_set.test_backend_set.name
	name = var.listener_name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
	port = var.listener_port
	protocol = var.listener_protocol

	#Optional
	ip_version = var.listener_ip_version
}
```

## Argument Reference

The following arguments are supported:

* `default_backend_set_name` - (Required) (Updatable) The name of the associated backend set.  Example: `example_backend_set` 
* `ip_version` - (Optional) (Updatable) IP version associated with the listener.
* `name` - (Required) A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
* `port` - (Required) (Updatable) The communication port for the listener.  Example: `80` 
* `protocol` - (Required) (Updatable) The protocol on which the listener accepts connection requests. For public network load balancers, ANY protocol refers to TCP/UDP. For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true). To get a list of valid protocols, use the [ListNetworkLoadBalancersProtocols](https://docs.cloud.oracle.com/iaas/api/#/en/NetworkLoadBalancer/20200501/networkLoadBalancerProtocol/ListNetworkLoadBalancersProtocols) operation.  Example: `TCP` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `default_backend_set_name` - The name of the associated backend set.  Example: `example_backend_set` 
* `ip_version` - IP version associated with the listener.
* `name` - A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener` 
* `port` - The communication port for the listener.  Example: `80` 
* `protocol` - The protocol on which the listener accepts connection requests. For public network load balancers, ANY protocol refers to TCP/UDP. For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true). To get a list of valid protocols, use the [ListNetworkLoadBalancersProtocols](https://docs.cloud.oracle.com/iaas/api/#/en/NetworkLoadBalancer/20200501/networkLoadBalancerProtocol/ListNetworkLoadBalancersProtocols) operation.  Example: `TCP` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Listener
	* `update` - (Defaults to 20 minutes), when updating the Listener
	* `delete` - (Defaults to 20 minutes), when destroying the Listener


## Import

Listeners can be imported using the `id`, e.g.

```
$ terraform import oci_network_load_balancer_listener.test_listener "networkLoadBalancers/{networkLoadBalancerId}/listeners/{listenerName}" 
```

