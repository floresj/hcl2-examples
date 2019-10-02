package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

type Configuration struct {
	Username string `hcl:"username"`
	FullName string `hcl:"full_name"`
}

func main() {
	// Get a handle to the hclparser
	parser := hclparse.NewParser()

	// Specify file to parse.
	// Alternative is directly pass in a []byte  containing the contents of an `hcl` configuration
	// using parser.ParseHCL([]byte, string)
	file, diags := parser.ParseHCLFile("./config.hcl")

	// diags represents an error value containing diagnostic information. Utilize the
	// HasErrors() func to check if any errors exist. These parse methods return hcl.Diagnostics
	// instead of `error` values. However, the hcl.Diagnostics implements the `error` interface
	// so you can use them for passing up error values.
	if diags.HasErrors() {
		log.Fatal(diags)
	}

	var config Configuration

	// The EvalContext is where you define the variables anjd functions you want used when the
	// hcl configuration is evaluated.
	evalCtx := &hcl.EvalContext{

		Functions: map[string]function.Function{
			// `title` is an example of a custom function that will execute `stings.Title` against
			// the input
			"title": titleFunc,

			// This is one of the built-in standard lib functions provided by the
			// go-cty library. Check out this package for more
			// https://github.com/zclconf/go-cty/blob/master/cty/function/stdlib
			"lower": stdlib.LowerFunc,
		},
	}
	confDiags := gohcl.DecodeBody(file.Body, evalCtx, &config)

	if confDiags.HasErrors() {
		log.Fatal(confDiags)
	}

	fmt.Printf("Username: %s\n", config.Username)
	fmt.Printf("FullName: %s\n", config.FullName)

}

var titleFunc = function.New(&function.Spec{
	// Params defines the positional parameters. These will all be required
	Params: []function.Parameter{
		{
			Name:             "str",
			Type:             cty.String,
			AllowDynamicType: true,
		},
	},
	Type: function.StaticReturnType(cty.String),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		in := args[0].AsString()
		out := strings.Title(in)

		// Since we're dealing with `cty.Value` types, we execute `cty.StringVal()` and
		// return the result of that function
		return cty.StringVal(out), nil
	},
})
