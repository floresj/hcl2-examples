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
