# da0x/rest

Welcome to this simple rest client for golang. This package aims to give
a very lean interface for calling rest services.

Example:
```
var o struct {
    UserId int `json:"user_id"`
    Name string `json:"name"`
}
_, err := api.Get("https://example.com/json", &o)
if err != nil {
    log.Fatalln(err)
}
println(o)
```

To import the library, use:

```
import "github.com/da0x/rest"
```


