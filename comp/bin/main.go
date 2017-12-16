package main

import (
	"fmt"

	"github.com/Spriithy/virt/comp"
)

func main() {
	var testArray virt.Type
	testArray = &virt.ArrayType{
		Alias:     "",
		ElemCount: 4,
		Of:        virt.Byte,
	}

	var testStruct virt.Type
	testStruct = &virt.StructType{
		Alias: "test",
		Fields: []virt.Type{
			virt.String,
			nil, // test out the nil type replacement
			testArray,
		},
	}

	fmt.Printf("%s\n", testStruct.String())
	fmt.Printf("%s\n", testStruct.TypeString())

	var testFunction virt.Type
	testFunction = &virt.FunctionType{
		Alias:   "",
		Returns: nil,
		Args: []virt.Type{
			virt.Int, virt.String,
		},
	}

	fmt.Printf("%s\n", testFunction.TypeString())
}
