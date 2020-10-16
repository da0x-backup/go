# da0x/rest
Welcome to this simple rest client for golang. This package aims to give
a minimal interface for calling rest services. Some headers are applied by
default.
## Instalation
To install the library simply run:
```
$ go get github.com/da0x/rest
```
To import the library, use:
```
import "github.com/da0x/rest"
```
## Get
You can 
Example:
```
var o struct {
    UserId int `json:"user_id"`
    Name string `json:"name"`
}
code, err := rest.Get("https://example.com/json", &o)
if err != nil {
    log.Fatalln(err)
}
println("response code:", code)
println("response object:", o)
```

## Request Headers
By default, this library will send default headers to indicate that the request content
type is json and to close the connection.
You can reset or clear out the request headers by directly manipulating or replacing rest.RequestHeaders.
Example
```
headers := rest.DefaultHeaders()
headers["new-header"] = "value"
rest.RequestHeaders = headers
...
_, err := rest.Get("...", &o)
```
