
## `github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/namespaces` Documentation

The `namespaces` SDK allows for interaction with the Azure Resource Manager Service `relay` (API Version `2017-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/relay/2017-04-01/namespaces"
```


### Client Initialization

```go
client := namespaces.NewNamespacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `NamespacesClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := namespaces.NewSubscriptionID()

payload := namespaces.CheckNameAvailability{
	// ...
}

read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := namespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

payload := namespaces.RelayNamespace{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `NamespacesClient.CreateOrUpdateAuthorizationRule`

```go
ctx := context.TODO()
id := namespaces.NewAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "authorizationRuleValue")

payload := namespaces.AuthorizationRule{
	// ...
}

read, err := client.CreateOrUpdateAuthorizationRule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesClient.Delete`

```go
ctx := context.TODO()
id := namespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `NamespacesClient.DeleteAuthorizationRule`

```go
ctx := context.TODO()
id := namespaces.NewAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "authorizationRuleValue")
read, err := client.DeleteAuthorizationRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesClient.Get`

```go
ctx := context.TODO()
id := namespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesClient.GetAuthorizationRule`

```go
ctx := context.TODO()
id := namespaces.NewAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "authorizationRuleValue")
read, err := client.GetAuthorizationRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesClient.List`

```go
ctx := context.TODO()
id := namespaces.NewSubscriptionID()
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NamespacesClient.ListAuthorizationRules`

```go
ctx := context.TODO()
id := namespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")
// alternatively `client.ListAuthorizationRules(ctx, id)` can be used to do batched pagination
items, err := client.ListAuthorizationRulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NamespacesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := namespaces.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NamespacesClient.ListKeys`

```go
ctx := context.TODO()
id := namespaces.NewAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "authorizationRuleValue")
read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesClient.RegenerateKeys`

```go
ctx := context.TODO()
id := namespaces.NewAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "authorizationRuleValue")

payload := namespaces.RegenerateAccessKeyParameters{
	// ...
}

read, err := client.RegenerateKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesClient.Update`

```go
ctx := context.TODO()
id := namespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

payload := namespaces.RelayUpdateParameters{
	// ...
}

read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
