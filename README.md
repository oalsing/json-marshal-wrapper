# json-marshal-wrapper

Adds functionality to wrap marshaled json around a key.

When using json.Marshal from the standard json library, it returns the struct with each field as a json key.
```json
{
  "sku": "1337",
  "name": "Golang"
}
```

However, in many cases it might be useful to wrap the marshaled struct with another key as below.
```json
{
  "product": {
    "sku": "1337",
    "name": "Golang"
  }
}
```

This library provides an easy way to wrap the marshaled json around an arbitary key, as in the example below.
```go
mockProduct := product.Product{
  SKU:       "1337",
  Name:      "Golang",
}

product_j, err := marshal.Wrap(mockProduct, "product")
```

## Installation
```sh
dep ensure -add github.com/oalsing/json-marshal-wrapper
```
