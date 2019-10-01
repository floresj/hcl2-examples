package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

// Configuration is the top-level struct that contains attributes and blocks defined within
// an `hcl` file.
type Configuration struct {
	Version string            `hcl:"version"`
	Tags    map[string]string `hcl:"tags"`

	// Users refer to the `user` blocks defined in the `hcl` file
	Users []User `hcl:"user,block"`
}

// User is a hcl block containing information about a user. Unless specified, all fields
// are required. In order to indicate that a field is optional, use the `optional` tag or
// define the field as a pointer
type User struct {
	Username       string   `hcl:"username"`
	FirstName      string   `hcl:"first_name"`
	LastName       string   `hcl:"last_name"`
	CloudProviders []string `hcl:"cloud_providers,optional"` // Example of optional field
	Enabled        bool     `hcl:"enabled"`
}

// String prints out a pretty version of the user struct
func (u User) String() string {
	return fmt.Sprintf(
		"Username: %s\nFirstname: %s\nLastName: %s\nCloudProviders: %v\nEnabled: %v\n",
		u.Username, u.FirstName, u.LastName, u.CloudProviders, u.Enabled,
	)
}

func main() {
	// Get a handle to the hclparser
	parser := hclparse.NewParser()

	// Specify file to parse.
	// Alternative is directly pass in a []byte  containing the contents of an `hcl` configuration
	// using parser.ParseHCL([]byte, string)
	file, diags := parser.ParseHCLFile("./basic.hcl")

	// diags is represents an error value containing diagnostic information. Utilize the
	// HasErrors() func to check if any errors exist. These parse methods return hcl.Diagnostics
	// instead of `error` values. However, the hcl.Diagnostics implements the `error` interface
	// so you can use them for passing up error values.
	if diags.HasErrors() {
		log.Fatal(diags)
	}

	// The `gohcl` package contains high-level functions for decoding `hcl` into native
	// go values. The second parameter contains the Evaluation Context which is currently nil.
	// That will be demonstrated in another example.
	var config Configuration
	confDiags := gohcl.DecodeBody(file.Body, nil, &config)

	if confDiags.HasErrors() {
		log.Fatal(confDiags)
	}

	fmt.Println(config.Version)
	fmt.Println(config.Tags)

	for _, user := range config.Users {
		fmt.Println(user)
	}
}
