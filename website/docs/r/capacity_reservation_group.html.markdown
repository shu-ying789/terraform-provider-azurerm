---
subcategory: "Compute"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_capacity_reservation_group"
description: |-
  Manages a Compute.
---

# azurerm_capacity_reservation_group

Manages a Compute.

## Example Usage

```hcl
resource "azurerm_capacity_reservation_group" "example" {
  name                = "example"
  resource_group_name = "example"
  location            = "West Europe"
}
```

## Arguments Reference

The following arguments are supported:

* `location` - (Required) The Azure Region where the Compute should exist. Changing this forces a new Compute to be created.

* `name` - (Required) The name which should be used for this Compute. Changing this forces a new Compute to be created.

* `resource_group_name` - (Required) The name of the Resource Group where the Compute should exist. Changing this forces a new Compute to be created.

---

* `tags` - (Optional) A mapping of tags which should be assigned to the Compute.

* `zones` - (Optional) Specifies a list of TODO.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported: 

* `id` - The ID of the Compute.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Compute.
* `read` - (Defaults to 5 minutes) Used when retrieving the Compute.
* `update` - (Defaults to 30 minutes) Used when updating the Compute.
* `delete` - (Defaults to 30 minutes) Used when deleting the Compute.

## Import

Computes can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_capacity_reservation_group.example C:/Program Files/Git/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1

```
