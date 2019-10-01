# Basic

A simple example of decoding a `hcl` configuration into native `go` values. This examples showcases the following:

- Defining top-level and object level `attributes`
- Defining `blocks` with `attributes`
- Specifying optional fields using the `optional` tag


## Configuration

```hcl
// version string attribute
version = "1.0.0"

// tags map attribute
tags = {
  "env" = "dev"
}

// Example of two user blocks. These will be handled as a slice of users in the `Configuration`
// struct.
user {
  username        = "floresj"
  first_name      = "John"
  last_name       = "Flores"
  cloud_providers = ["AWS"]
  enabled         = false
}

user {
  username   = "foo"
  first_name = "Foo"
  last_name  = "Bar"
  enabled    = true
}
```

## Run it!

Run the [main](./main.go) to see it in action.

```bash
go run main.go
```
