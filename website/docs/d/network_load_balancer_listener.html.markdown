---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_listener"
sidebar_current: "docs-oci-datasource-network_load_balancer-listener"
description: |-
  Provides details about a specific Listener in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_listener
This data source provides details about a specific Listener resource in Oracle Cloud Infrastructure Network Load Balancer service.

Retrieves listener properties associated with a given network load balancer and listener name.

## Example Usage

```hcl
data "oci_network_load_balancer_listener" "test_listener" {
	#Required
	listener_name = oci_network_load_balancer_listener.test_listener.name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `listener_name` - (Required) The name of the listener to get.  Example: `example_listener` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `default_backend_set_name` - The name of the associated backend set.  Example: `example_backend_set` 
* `ip_version` - IP version associated with the listener.
* `name` - A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener` 
* `port` - The communication port for the listener.  Example: `80` 
* `protocol` - The protocol on which the listener accepts connection requests. For public network load balancers, ANY protocol refers to TCP/UDP. For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true). To get a list of valid protocols, use the [ListNetworkLoadBalancersProtocols](https://docs.cloud.oracle.com/iaas/api/#/en/NetworkLoadBalancer/20200501/networkLoadBalancerProtocol/ListNetworkLoadBalancersProtocols) operation.  Example: `TCP` 

