package virt

import (
	"testing"
)

func TestUnitType(t *testing.T) {
	if !Unit.Equals(Unit) {
		t.Error("unit != unit: reflexivity rule broken")
	}

	if Unit.Equals(Unsigned32) {
		t.Error("unit == u32: equality rule broken")
	}
}

func TestTypeAlias(t *testing.T) {
	var alias Type
	alias = &TypeAlias{
		Name: "char",
		Of:   Unsigned8,
	}

	if !alias.Equals(Unsigned8) {
		t.Error("char != u8: aliasing broken")
	}
}

func TestArrayType(t *testing.T) {
	var arr1, arr2 Type
	arr1 = &ArrayType{
		Alias:     "arrType1",
		ElemCount: 32,
		Of:        Unsigned32,
	}

	arr2 = &ArrayType{
		Alias:     "arrType2",
		ElemCount: 32,
		Of:        Unsigned32,
	}

	if !arr1.Equals(arr2) {
		t.Errorf("arr1 != arr2: type mismatch\narr1: %s\narr2: %s", arr1.TypeString(), arr2.TypeString())
	}
}

func TestStructType(t *testing.T) {
	var struct1, struct2 Type
	struct1 = &StructType{
		Alias: "structType1",
		Fields: []Type{
			Unsigned64, Float32,
		},
	}

	struct2 = &StructType{
		Alias: "structType2",
		Fields: []Type{
			Unsigned64, Float32,
		},
	}

	if !struct1.Equals(struct2) {
		t.Errorf("struct1 != struct2: type mismatch\nstruct1: %s\nstruct2: %s", struct1.TypeString(), struct2.TypeString())
	}
}

func TestEnumType(t *testing.T) {
	var enum1, enum2 Type
	enum1 = &EnumType{
		Alias: "enumType1",
		Entries: map[string]struct{}{
			"X": {},
			"Y": {},
			"Z": {},
		},
	}

	enum2 = &EnumType{
		Alias: "enumType2",
		Entries: map[string]struct{}{
			"X": {},
			"Y": {},
			"Z": {},
		},
	}

	if !enum1.Equals(enum2) {
		t.Errorf("enum1 != enum2: type mismatch\nenum1: %s\nenum2: %s", enum1.TypeString(), enum2.TypeString())
	}
}

func TestFunctionType(t *testing.T) {
	var fun1, fun2 Type
	fun1 = &FunctionType{
		Alias:   "funType1",
		Returns: Unit,
		Args: []Type{
			&PointerType{"", Byte},
		},
	}

	fun2 = &FunctionType{
		Alias:   "funType2",
		Returns: Unit,
		Args: []Type{
			&PointerType{"", Byte},
		},
	}

	if !fun1.Equals(fun2) {
		t.Errorf("fun1 != fun2: type mismatch\nfun1: %s\nfun2: %s", fun1.TypeString(), fun2.TypeString())
	}
}
