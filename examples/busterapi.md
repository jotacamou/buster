## Examples on how to use the `busterapi` library.

Creating a Buster API client:
```
import "github.com/jotacamou/buster/busterapi"

func main() {
    bc := &busterapi.NewClient(BUSTER_API_URL)
}

```

Creating a transaction on the Buster API:
```
// Generate API key with the gen-api-key.sh script
key := "d61b88e58feb9f4f66c85dce878178c0"

// resp is the response from the Buster API in a json byte array
resp, err := bc.CreateTransaction(key, "RID399281", "156")

if err != nil {
    // Handle error
}

println(string(resp))
```
Output:
```
{"id":"4e6a856d-c2da-4266-bbfe-6c3a11161ba9","created":"2020-08-14T03:05:16.082Z","status":"CREATED","referenceId":"RID399281","amount":"156"}
```

Retrieving a transaction by database ID:
```
resp, err := bc.GetTransaction(key, "4e6a856d-c2da-4266-bbfe-6c3a11161ba9")

if err != nil {
    // Handle error
}

println(string(resp))
```
Output:
```
{"id":"4e6a856d-c2da-4266-bbfe-6c3a11161ba9","created":"2020-08-14T03:05:16.000Z","status":"CANCELED","referenceId":"RID399281","amount":156}
```
