package virt

import (
	"fmt"
	"strings"
	"unsafe"
)

// The Type interface is responsible for providing a uniform base for
// all subsequent types of the type system
type Type interface {
	// Size is used to get the size (in bytes) of the given type
	Size() int

	// TypeString returns the bare type representation
	TypeString() string

	// Kind returns the kind of the type (ie. Void, Pointer, ...)
	Kind() TypeKind

	// Equals determines whether two given types are the same or not.
	// Two types are equal should their TypeString be the same
	Equals(Type) bool

	fmt.Stringer
}

// VoidType represents the absence of type (typeless)
type VoidType struct{}

// UndefinedType represents the type of undefined values/symbols
type UndefinedType struct{}

// BaseType represents primitive types such comp integers and floats
type BaseType struct {
	alias string

	// Which base type (hear primitive) is concerned
	Which int
}

// PointerType represents all types that are pointers to other types
type PointerType struct {
	alias string

	// The Type pointed to by the PointerType
	To Type
}

// ArrayType represents all types of fixed-length
type ArrayType struct {
	alias string
	size  int

	// The type of the elements held by such an array
	Of Type
}

// A Signature is the underlying representation of a function type
type Signature struct {
	// The return type of a function type
	Returns Type

	// The arguments of a function type
	Args []Type
}

// A FunctionType holds the signature of a function
type FunctionType struct {
	alias string

	// The Signature of the function (types)
	Signature *Signature
}

// Compile-time check
var (
	_ Type = (*VoidType)(nil)
	_ Type = (*UndefinedType)(nil)
	_ Type = (*BaseType)(nil)
	_ Type = (*PointerType)(nil)
	_ Type = (*ArrayType)(nil)
	_ Type = (*FunctionType)(nil)
)

// TypeKind represents the type class of a type
type TypeKind = int

const (
	VoidKind TypeKind = iota
	UndefinedKind
	BasicKind
	PointerKind
	ArrayKind
	FunctionKind
)

// These values are used to identify builtin types
const (
	u8typeId = iota
	u16typeId
	u32typeId
	u64typeId
	i8typeId
	i16typeId
	i32typeId
	i64typeId
	f64typeId
	f32typeId
)

/******************************************************************************/
/*** VoidType implementation                                                ***/
/******************************************************************************/

func (t *VoidType) Size() int {
	return 0
}

func (t *VoidType) TypeString() string {
	return t.String()
}

func (t *VoidType) Kind() TypeKind {
	return VoidKind
}

func (t *VoidType) Equals(other Type) bool {
	return other.Kind() == VoidKind
}

func (t *VoidType) String() string {
	return ""
}

/******************************************************************************/
/*** UndefinedType implementation                                           ***/
/******************************************************************************/

func (t *UndefinedType) Size() int {
	return 0
}

func (t *UndefinedType) TypeString() string {
	return t.String()
}

func (t *UndefinedType) Kind() TypeKind {
	return UndefinedKind
}

func (t *UndefinedType) Equals(other Type) bool {
	return other.Kind() == UndefinedKind
}

func (t *UndefinedType) String() string {
	return "undefined"
}

/******************************************************************************/
/*** BaseType implementation                                                ***/
/******************************************************************************/

func (t *BaseType) Size() int {
	switch t.Which {
	case i8typeId, u8typeId:
		return 1
	case i16typeId, u16typeId:
		return 2
	case i32typeId, u32typeId, f32typeId:
		return 4
	case i64typeId, u64typeId, f64typeId:
		return 8
	}

	return 0
}

func (t *BaseType) TypeString() string {
	return fmt.Sprintf("primitive<%d>", t.Which)
}

func (t *BaseType) Kind() TypeKind {
	return BasicKind
}

func (t *BaseType) Equals(other Type) bool {
	if other.Kind() != BasicKind {
		return false
	}

	return t.TypeString() == other.TypeString()
}

func (t *BaseType) String() string {
	if t.alias != "" {
		return t.alias
	}

	return t.TypeString()
}

/******************************************************************************/
/*** PointerType implementation                                             ***/
/******************************************************************************/

func (t *PointerType) Size() int {
	return int(unsafe.Sizeof(uintptr(0)))
}
func (t *PointerType) TypeString() string {
	return t.To.String() + "&"
}

func (t *PointerType) Kind() TypeKind {
	return PointerKind
}

func (t *PointerType) Equals(other Type) bool {
	if other.Kind() != PointerKind {
		return false
	}

	return t.TypeString() == other.TypeString()
}

func (t *PointerType) String() string {
	if t.alias != "" {
		return t.alias
	}

	return t.TypeString()
}

/******************************************************************************/
/*** ArrayType implementation                                               ***/
/******************************************************************************/

func (t *ArrayType) Size() int {
	return t.size * t.Of.Size()
}

func (t *ArrayType) TypeString() string {
	return fmt.Sprintf("[%d]%s", t.size, t.Of.String())
}

func (t *ArrayType) Kind() TypeKind {
	return ArrayKind
}

func (t *ArrayType) Equals(other Type) bool {
	if other.Kind() != ArrayKind {
		return false
	}

	return t.TypeString() == other.TypeString()
}

func (t *ArrayType) String() string {
	if t.alias != "" {
		return t.alias
	}

	return t.TypeString()
}

/******************************************************************************/
/*** FunctionType implementation                                            ***/
/******************************************************************************/

func (t *FunctionType) Size() int {
	return int(unsafe.Sizeof(uintptr(0))) // size of pointer
}

func (t *FunctionType) TypeString() string {
	if t.Signature == nil {
		return "func()"
	}

	repr := "func("
	for _, typ := range t.Signature.Args {
		repr += typ.String()
		repr += ", "
	}

	// Cut the trailing coma+space if there were args
	if len(t.Signature.Args) > 0 {
		repr = repr[:len(repr)-2]
	}

	repr += ") "
	repr += t.Signature.Returns.String()

	return strings.TrimSpace(repr)
}

func (t *FunctionType) Kind() TypeKind {
	return FunctionKind
}

func (t *FunctionType) Equals(other Type) bool {
	if other.Kind() != FunctionKind {
		return false
	}

	return t.TypeString() == other.TypeString()
}

func (t *FunctionType) String() string {
	if t.alias != "" {
		return t.alias
	}

	return t.TypeString()
}
