---
subcategory: "Compute"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_capacity_reservation"
description: |-
  Manages a Capacity Reservation.
---

# azurerm_capacity_reservation

Manages a Capacity Reservation.

## Example Usage

```hcl
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "example"
  location = "West Europe"
}

resource "azurerm_capacity_reservation" "example" {
  name                = "example"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  zones               = ["1"]
  sku {
    capacity = 1
    name     = "Standard_D2s_v3"
  }
}
```

## Arguments Reference

The following arguments are supported:

* `location` - (Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.

* `name` - (Required) The name which should be used for this Capacity Reservation. Changing this forces a new Capacity Reservation to be created.

* `resource_group_name` - (Required) The name of the Resource Group where the Capacity Reservation exists.

---

* `tags` - (Optional) A mapping of tags which should be assigned to the Capacity Reservation.

* `zones` - (Optional) Specifies a list of Availability Zones in which this Capacity Reservation should be located. Changing this forces a new Capacity Reservation to be created.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported: 

* `id` - The ID of the Capacity Reservation.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Capacity Reservation.
* `read` - (Defaults to 5 minutes) Used when retrieving the Capacity Reservation.
* `update` - (Defaults to 30 minutes) Used when updating the Capacity Reservation.
* `delete` - (Defaults to 30 minutes) Used when deleting the Capacity Reservation.

## Import

Capacity Reservation can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_capacity_reservation.example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1

```
