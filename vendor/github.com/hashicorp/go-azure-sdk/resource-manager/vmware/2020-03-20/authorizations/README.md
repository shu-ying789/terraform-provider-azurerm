
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2020-03-20/authorizations` Documentation

The `authorizations` SDK allows for interaction with the Azure Resource Manager Service `vmware` (API Version `2020-03-20`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2020-03-20/authorizations"
```


### Client Initialization

```go
client := authorizations.NewAuthorizationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `AuthorizationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := authorizations.NewAuthorizationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "authorizationValue")

payload := authorizations.ExpressRouteAuthorization{
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


### Example Usage: `AuthorizationsClient.Delete`

```go
ctx := context.TODO()
id := authorizations.NewAuthorizationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "authorizationValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `AuthorizationsClient.Get`

```go
ctx := context.TODO()
id := authorizations.NewAuthorizationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "authorizationValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthorizationsClient.List`

```go
ctx := context.TODO()
id := authorizations.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
