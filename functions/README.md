# Functions

Example of defining a custom function within an `hcl` configuration.

## Configuration

```hcl
// String will be converted to lowercase. This is using the stdlib.LowerFunc provided by
// https://github.com/zclconf/go-cty/blob/master/cty/function/stdlib
username = lower("FLORESJ")

// The input string will have the `strings.Title` func applied to it
full_name = title("john flores")
```

## Function Definition

In this example, we create a function that executes `strings.Title` against the provided input. This allows you to define functions that transform input in your configuration.

The underlying [go-cty](https://github.com/zclconf/go-cty) library contains a set of helper `functions` that can be leveraged as well. [Check those out here](https://github.com/zclconf/go-cty/tree/master/cty/function/stdlib)

```go
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
```

## Run it!

Run the [main](./main.go) to see it in action.

```bash
go run main.go
```
