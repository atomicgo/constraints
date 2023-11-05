package constraints

// Numeric is a constraint that matches any numeric type, including complex numbers.
type Numeric interface {
	Number | Complex
}

// Complex is a constraint that matches any complex number type.
type Complex interface {
	~complex64 | ~complex128
}

// Number is a constraint that matches any real number type.
// The Number constraint does not include Complex numbers, as those are not ordered.
// If you also need Complex numbers, use the Numeric constraint.
type Number interface {
	Float | Integer
}

// Float is a constraint that matches any floating point number type.
type Float interface {
	~float32 | ~float64
}

// Integer is a constraint that matches any integer type.
type Integer interface {
	Signed | Unsigned
}

// Ordered is a constraint that matches any ordered type.
type Ordered interface {
	Integer | Float | ~string
}

// Signed is a constraint that matches any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that matches any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
