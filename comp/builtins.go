package virt

var (
	// Unit represents the absence of type (functional-inspired)
	Unit = &UnitType{}

	// Undefined is used to describe symbols which type is not yet known
	Undefined = &UndefinedType{}

	// Unsigned8 represents 1-byte unsigned integers
	Unsigned8 = &BaseType{"u8", u8typeID}

	// Byte represents 1-byte unsigned integers used to hold characters
	Byte = &TypeAlias{"byte", Unsigned8}

	// Unsigned16 represents 2-byte unsigned integers
	Unsigned16 = &BaseType{"u16", u16typeID}

	// Unsigned32 represents 4-byte unsigned integers
	Unsigned32 = &BaseType{"u32", u32typeID}

	// UnsignedInt represents 4-byte unsigned integers (platform dependant)
	UnsignedInt = &TypeAlias{"uint", Unsigned32}

	// Unsigned64 represents 8-byte unsigned integers
	Unsigned64 = &BaseType{"u64", u64typeID}

	// Signed8 represents 1-byte signed integers
	Signed8 = &BaseType{"i8", i8typeID}

	// Signed16 represents 2-byte signed integers
	Signed16 = &BaseType{"i16", i16typeID}

	// Signed32 represents 4-byte signed integers
	Signed32 = &BaseType{"i32", i32typeID}

	// Int represents 4-byte signed integers (platform dependant)
	Int = &TypeAlias{"int", Signed32}

	// Signed64 is an 8-byte signed integer
	Signed64 = &BaseType{"i64", i64typeID}

	// Float64 represents 8-byte floating point numbers
	Float64 = &BaseType{"f64", f64typeID}

	// Float64 represents 4-byte floating point numbers
	Float32 = &BaseType{"f32", f32typeID}

	// String represents strings of characters (C-like strings)
	String = &TypeAlias{"string", &PointerType{"", Byte}}
)
