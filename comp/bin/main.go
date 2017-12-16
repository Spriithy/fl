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
	fmt.Printf("0x%x %s\n", testStruct.Id(), testStruct.TypeString())

	var testFunction virt.Type
	testFunction = &virt.FunctionType{
		Alias:   "",
		Returns: nil,
		Args: []virt.Type{
			virt.Int, testStruct,
		},
	}

	fmt.Printf("0x%x %s\n", testFunction.Id(), testFunction.TypeString())

	var testEnum virt.Type
	testEnum = &virt.EnumType{
		Alias: "token",
		Values: map[string]int{
			"PLUS":  0,
			"MINUS": 1,
			"INT":   2,
		},
	}

	fmt.Printf("0x%x %s\n", testEnum.Id(), testEnum.TypeString())
}
