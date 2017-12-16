package virt

import (
	"fmt"
	"reflect"
	"strings"
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

	// Id returns the ID of the given type. IDs are most likely to be unique
	Id() TypeID

	// Equals determines whether two given types are the same or not.
	// Two types are equal should their TypeString be the same
	Equals(Type) bool

	fmt.Stringer
}

// UnitType represents the absence of type (typeless)
type UnitType struct{}

// UndefinedType represents the type of undefined values/symbols
type UndefinedType struct{}

// TypeAlias is simply used to create type aliases
type TypeAlias struct {
	// The name of the alias
	Name string

	// The aliased type
	Of Type
}

// BaseType represents primitive types such comp integers and floats
type BaseType struct {
	Alias string

	// Which base type (hear primitive) is concerned
	Which TypeID
}

// PointerType represents all types that are pointers to other types
type PointerType struct {
	Alias string

	// The Type pointed to by the PointerType
	To Type
}

// ArrayType represents all types of fixed-length
type ArrayType struct {
	Alias     string
	ElemCount int

	// The type of the elements held by such an array
	Of Type
}

// A StructType represents struct-like data structure
type StructType struct {
	Alias string

	// The Fields listed in the struct type
	Fields []Type
}

// An EnumType represents type enumerations
type EnumType struct {
	Alias string

	// The Entries (named) of the enum type
	Entries map[string]struct{}
}

// A FunctionType holds the signature of a function
type FunctionType struct {
	Alias string

	// The return type of a function type
	Returns Type

	// The arguments of a function type
	Args []Type
}

// Compile-time check
var (
	_ Type = (*UnitType)(nil)
	_ Type = (*UndefinedType)(nil)
	_ Type = (*TypeAlias)(nil)
	_ Type = (*BaseType)(nil)
	_ Type = (*PointerType)(nil)
	_ Type = (*ArrayType)(nil)
	_ Type = (*StructType)(nil)
	_ Type = (*EnumType)(nil)
	_ Type = (*FunctionType)(nil)
)

// TypeKind represents the type class of a type
type TypeKind = int

const (
	UnitKind = iota
	UndefinedKind
	BasicKind
	PointerKind
	ArrayKind
	StructKind
	EnumKind
	FunctionKind
)

// TypeID is used to hold hashes of types to identify them
type TypeID = uint64

// hash hashes a string using the java.lang.String.hashCode method
func hash(str string) TypeID {
	var hash TypeID = 7

	for _, c := range str {
		hash = hash*31 + TypeID(c)
	}

	return hash
}

// These values are used to identify builtin types
const (
	u8typeID TypeID = iota
	u16typeID
	u32typeID
	u64typeID
	i8typeID
	i16typeID
	i32typeID
	i64typeID
	f64typeID
	f32typeID
)

const (
	unitTypeString     = "unit"
	undefTypeString    = "undef"
	funcTypeArrow      = "->"
	nilTypeReplacement = "%!(nil)"
)

/******************************************************************************/
/*** UnitType implementation                                                ***/
/******************************************************************************/

func (t *UnitType) Size() int {
	return 0
}

func (t *UnitType) TypeString() string {
	return unitTypeString
}

func (t *UnitType) Kind() TypeKind {
	return UnitKind
}

func (t *UnitType) Id() TypeID {
	return hash(t.TypeString())
}

func (t *UnitType) Equals(other Type) bool {
	return other != nil && other.Kind() == UnitKind
}

func (t *UnitType) String() string {
	return unitTypeString
}

/******************************************************************************/
/*** TypeAlias implementation                                               ***/
/******************************************************************************/

func (t *TypeAlias) Size() int {
	if t.Of == nil {
		return 0
	}

	return t.Of.Size()
}

func (t *TypeAlias) TypeString() string {
	switch t.Name {
	case "":
		if t.Of == nil {
			return nilTypeReplacement
		}

		return t.Of.String()
	default:
		return t.Name
	}
}

func (t *TypeAlias) Kind() TypeKind {
	if t.Of == nil {
		return -1
	}

	return t.Of.Kind()
}

func (t *TypeAlias) Id() TypeID {
	return hash(t.TypeString())
}

func (t *TypeAlias) Equals(other Type) bool {
	if other == nil || t.Of == nil {
		return false
	}

	return t.Of.Equals(other)
}

func (t *TypeAlias) String() string {
	return t.TypeString()
}

/******************************************************************************/
/*** UndefinedType implementation                                           ***/
/******************************************************************************/

func (t *UndefinedType) Size() int {
	return 0
}

func (t *UndefinedType) TypeString() string {
	return undefTypeString
}

func (t *UndefinedType) Kind() TypeKind {
	return UndefinedKind
}

func (t *UndefinedType) Id() TypeID {
	return hash(t.TypeString())
}

func (t *UndefinedType) Equals(other Type) bool {
	return other != nil && other.Kind() == UndefinedKind
}

func (t *UndefinedType) String() string {
	return undefTypeString
}

/******************************************************************************/
/*** BaseType implementation                                                ***/
/******************************************************************************/

func (t *BaseType) Size() int {
	switch t.Which {
	case i8typeID, u8typeID:
		return 1
	case i16typeID, u16typeID:
		return 2
	case i32typeID, u32typeID, f32typeID:
		return 4
	case i64typeID, u64typeID, f64typeID:
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

func (t *BaseType) Id() TypeID {
	return t.Which
}

func (t *BaseType) Equals(other Type) bool {
	if other == nil || other.Kind() != BasicKind {
		return false
	}

	return t.Which == other.(*BaseType).Which
}

func (t *BaseType) String() string {
	if t.Alias != "" {
		return t.Alias
	}

	return t.TypeString()
}

/******************************************************************************/
/*** PointerType implementation                                             ***/
/******************************************************************************/

func (t *PointerType) Size() int {
	return Unsigned64.Size()
}

func (t *PointerType) TypeString() string {
	if t.To == nil {
		return nilTypeReplacement + "&"
	}

	return t.To.String() + "&"
}

func (t *PointerType) Kind() TypeKind {
	return PointerKind
}

func (t *PointerType) Id() TypeID {
	return hash(t.TypeString())
}

func (t *PointerType) Equals(other Type) bool {
	if other == nil || other.Kind() != PointerKind {
		return false
	}

	return t.To.Equals(other.(*PointerType).To)
}

func (t *PointerType) String() string {
	if t.Alias != "" {
		return t.Alias
	}

	return t.TypeString()
}

/******************************************************************************/
/*** ArrayType implementation                                               ***/
/******************************************************************************/

func (t *ArrayType) Size() int {
	if t.Of == nil {
		return 0
	}

	return t.ElemCount * t.Of.Size()
}

func (t *ArrayType) TypeString() string {
	if t.Of == nil {
		return fmt.Sprintf("%s[%d]", nilTypeReplacement, t.ElemCount)
	}

	return fmt.Sprintf("%s[%d]", t.Of.String(), t.ElemCount)
}

func (t *ArrayType) Kind() TypeKind {
	return ArrayKind
}

func (t *ArrayType) Id() TypeID {
	return hash(t.TypeString())
}

func (t *ArrayType) Equals(other Type) bool {
	if other == nil || other.Kind() != ArrayKind {
		return false
	}

	return t.Id() == other.Id()
}

func (t *ArrayType) String() string {
	if t.Alias != "" {
		return t.Alias
	}

	return t.TypeString()
}

/******************************************************************************/
/*** StructType implementation                                              ***/
/******************************************************************************/

func (t *StructType) Size() int {
	if t.Fields == nil {
		return 0
	}

	var size int
	for _, typ := range t.Fields {
		if typ != nil {
			size += typ.Size()
		}
	}

	return size
}

func (t *StructType) TypeString() (repr string) {
	if t.Fields == nil || len(t.Fields) == 0 {
		return "struct{}"
	}

	repr += "struct {\n  "
	for _, typ := range t.Fields {
		if typ == nil {
			repr += nilTypeReplacement
		} else {
			repr += typ.String()
		}
		repr += "\n  "
	}

	// Remove the leading spaces on the last type line
	repr = repr[:len(repr)-2]

	return repr + "}"
}

func (t *StructType) Kind() TypeKind {
	return StructKind
}

func (t *StructType) Id() TypeID {
	return hash(t.TypeString())
}

func (t *StructType) Equals(other Type) bool {
	if other == nil || other.Kind() != StructKind {
		return false
	}

	return t.Id() == other.Id()
}

func (t *StructType) String() string {
	if t.Alias != "" {
		return "struct " + t.Alias
	}

	return t.TypeString()
}

/******************************************************************************/
/*** EnumType implementation                                                ***/
/******************************************************************************/

func (t *EnumType) Size() int {
	return Int.Size()
}

func (t *EnumType) TypeString() (repr string) {
	if t.Entries == nil || len(t.Entries) == 0 {
		return "enum{}"
	}

	repr = "enum {\n  "

	for name := range t.Entries {
		repr += fmt.Sprintf("%s\n  ", name)
	}

	// Remove 2 trailing whitespaces
	repr = repr[:len(repr)-2]

	return repr + "}"
}

func (t *EnumType) Kind() TypeKind {
	return EnumKind
}

func (t *EnumType) Id() TypeID {
	return hash(t.TypeString())
}

func (t *EnumType) Equals(other Type) bool {
	if other == nil || other.Kind() != EnumKind {
		return false
	}

	// Can't use other.Id() == t.Id()
	// for enum type strings may vary over time (maps are unordered)

	return reflect.DeepEqual(t.Entries, other.(*EnumType).Entries)
}

func (t *EnumType) String() string {
	if t.Alias != "" {
		return "enum " + t.Alias
	}

	return fmt.Sprintf("enum <anonymous>[%d]", len(t.Entries))
}

/******************************************************************************/
/*** FunctionType implementation                                            ***/
/******************************************************************************/

// DefaultFunctionReturnValue is used to determine the return type of functions
// that would miss such annotation
var DefaultFunctionReturnValue = Unit

func (t *FunctionType) Size() int {
	return Unsigned64.Size()
}

func (t *FunctionType) TypeString() (repr string) {
	if t.Args == nil {
		repr = "()"

		if t.Returns == nil {
			repr += " " + funcTypeArrow + " " + nilTypeReplacement
			return repr
		} else if len(t.Returns.String()) == 0 {
			repr += " " + funcTypeArrow + " " + DefaultFunctionReturnValue.String()
			return repr
		}

		repr += " " + funcTypeArrow + " " + t.Returns.String()
		return strings.TrimSpace(repr)
	}

	repr = "("
	for _, typ := range t.Args {
		if typ != nil {
			repr += typ.String()
		} else {
			repr += nilTypeReplacement
		}
		repr += ", "
	}

	// Cut the trailing coma+space if there were args
	if len(t.Args) > 0 {
		repr = repr[:len(repr)-2]
	}

	repr += ") "
	if t.Returns != nil {
		// Add arrow only if type actually has a type string
		if len(t.Returns.String()) > 0 {
			repr += funcTypeArrow + " "
		}

		repr += t.Returns.String()
	} else {
		repr += funcTypeArrow + " "

		// TODO nil check
		repr += DefaultFunctionReturnValue.String()
	}

	return strings.TrimSpace(repr)
}

func (t *FunctionType) Kind() TypeKind {
	return FunctionKind
}

func (t *FunctionType) Id() TypeID {
	return hash(t.TypeString())
}

func (t *FunctionType) Equals(other Type) bool {
	if other == nil || other.Kind() != FunctionKind {
		return false
	}

	return t.Id() == other.Id()
}

func (t *FunctionType) String() string {
	if t.Alias != "" {
		return t.Alias
	}

	return t.TypeString()
}
