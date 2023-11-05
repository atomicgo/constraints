package constraints

import (
	"reflect"
	"testing"
)

func TestConstraints(t *testing.T) {
	t.Run("Numbers", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			// Complex numbers
			acceptNumeric(t, complex64(0))
			acceptNumeric(t, complex128(0))

			// Signed integers
			acceptNumeric(t, int(0))
			acceptNumber(t, int8(0))
			acceptNumber(t, int16(0))
			acceptNumber(t, int32(0))
			acceptNumeric(t, int64(0))

			// Unsigned integers
			acceptNumeric(t, uint(0))
			acceptNumeric(t, uint8(0))
			acceptNumeric(t, uint16(0))
			acceptNumeric(t, uint32(0))
			acceptNumeric(t, uint64(0))
			acceptNumeric(t, uintptr(0))

			// Floating point numbers
			acceptNumeric(t, float32(0))
			acceptNumeric(t, float64(0))
		})

		t.Run("Number", func(t *testing.T) {
			// Signed integers
			acceptNumber(t, int(0))
			acceptNumber(t, int8(0))
			acceptNumber(t, int16(0))
			acceptNumber(t, int32(0))
			acceptNumber(t, int64(0))

			// Unsigned integers
			acceptNumber(t, uint(0))
			acceptNumber(t, uint8(0))
			acceptNumber(t, uint16(0))
			acceptNumber(t, uint32(0))
			acceptNumber(t, uint64(0))
			acceptNumber(t, uintptr(0))

			// Floating point numbers
			acceptNumber(t, float32(0))
			acceptNumber(t, float64(0))
		})

		t.Run("Complex", func(t *testing.T) {
			// Complex numbers
			acceptComplex(t, complex64(0))
			acceptComplex(t, complex128(0))
		})

		t.Run("Integer", func(t *testing.T) {
			// Signed integers
			acceptInteger(t, int(0))
			acceptInteger(t, int8(0))
			acceptInteger(t, int16(0))
			acceptInteger(t, int32(0))
			acceptInteger(t, int64(0))

			// Unsigned integers
			acceptInteger(t, uint(0))
			acceptInteger(t, uint8(0))
			acceptInteger(t, uint16(0))
			acceptInteger(t, uint32(0))
			acceptInteger(t, uint64(0))
			acceptInteger(t, uintptr(0))
		})

		t.Run("Float", func(t *testing.T) {
			// Floating point numbers
			acceptFloat(t, float32(0))
			acceptFloat(t, float64(0))
		})

		t.Run("Ordered", func(t *testing.T) {
			// Signed integers
			acceptOrdered(t, int(0))
			acceptOrdered(t, int8(0))
			acceptOrdered(t, int16(0))
			acceptOrdered(t, int32(0))
			acceptOrdered(t, int64(0))

			// Unsigned integers
			acceptOrdered(t, uint(0))
			acceptOrdered(t, uint8(0))
			acceptOrdered(t, uint16(0))
			acceptOrdered(t, uint32(0))
			acceptOrdered(t, uint64(0))
			acceptOrdered(t, uintptr(0))

			// Floating point numbers
			acceptOrdered(t, float32(0))
			acceptOrdered(t, float64(0))

			// Strings
			acceptOrdered(t, "")
		})

		t.Run("Signed", func(t *testing.T) {
			// Signed integers
			acceptSigned(t, int(0))
			acceptSigned(t, int8(0))
			acceptSigned(t, int16(0))
			acceptSigned(t, int32(0))
			acceptSigned(t, int64(0))
		})

		t.Run("Unsigned", func(t *testing.T) {
			// Unsigned integers
			acceptUnsigned(t, uint(0))
			acceptUnsigned(t, uint8(0))
			acceptUnsigned(t, uint16(0))
			acceptUnsigned(t, uint32(0))
			acceptUnsigned(t, uint64(0))
			acceptUnsigned(t, uintptr(0))
		})
	})

	t.Run("Comparisons", func(t *testing.T) {
		t.Run("Comparable", func(t *testing.T) {
			// Complex numbers
			acceptComparable(t, complex64(0))
			acceptComparable(t, complex128(0))

			// Signed integers
			acceptComparable(t, int(0))
			acceptComparable(t, int8(0))
			acceptComparable(t, int16(0))
			acceptComparable(t, int32(0))
			acceptComparable(t, int64(0))

			// Unsigned integers
			acceptComparable(t, uint(0))
			acceptComparable(t, uint8(0))
			acceptComparable(t, uint16(0))
			acceptComparable(t, uint32(0))
			acceptComparable(t, uint64(0))
			acceptComparable(t, uintptr(0))

			// Floating point numbers
			acceptComparable(t, float32(0))
			acceptComparable(t, float64(0))

			// Strings
			acceptComparable(t, "")
		})

		t.Run("Orderable", func(t *testing.T) {
			// Signed integers
			acceptOrderable(t, int(0))
			acceptOrderable(t, int8(0))
			acceptOrderable(t, int16(0))
			acceptOrderable(t, int32(0))
			acceptOrderable(t, int64(0))

			// Unsigned integers
			acceptOrderable(t, uint(0))
			acceptOrderable(t, uint8(0))
			acceptOrderable(t, uint16(0))
			acceptOrderable(t, uint32(0))
			acceptOrderable(t, uint64(0))
			acceptOrderable(t, uintptr(0))

			// Floating point numbers
			acceptOrderable(t, float32(0))
			acceptOrderable(t, float64(0))

			// Strings
			acceptOrderable(t, "")
		})

		t.Run("Addable", func(t *testing.T) {
			// Signed integers
			acceptAddable(t, int(0))
			acceptAddable(t, int8(0))
			acceptAddable(t, int16(0))
			acceptAddable(t, int32(0))
			acceptAddable(t, int64(0))

			// Unsigned integers
			acceptAddable(t, uint(0))
			acceptAddable(t, uint8(0))
			acceptAddable(t, uint16(0))
			acceptAddable(t, uint32(0))
			acceptAddable(t, uint64(0))
			acceptAddable(t, uintptr(0))

			// Floating point numbers
			acceptAddable(t, float32(0))
			acceptAddable(t, float64(0))

			// Strings
			acceptAddable(t, "")
		})
	})
}

// This function is a helper that tries to accept a value of a constraint type.
// If the constraint is satisfied, this will compile.
func accept[T any](t *testing.T, value T) {
	t.Helper()

	// Get the type of 'value'.
	typeOfValue := reflect.TypeOf(value)

	// Check if the type is not nil (it shouldn't be for a value).
	if typeOfValue == nil {
		t.Error("Cannot determine the type of the value")
		return
	}

	// Run named test
	t.Run(typeOfValue.Name(), func(t *testing.T) {
		// If the code compiles, the test passes.
	})
}

func acceptNumeric[T Numeric](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptNumber[T Number](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptComplex[T Complex](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptFloat[T Float](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptInteger[T Integer](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptOrdered[T Ordered](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptSigned[T Signed](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptUnsigned[T Unsigned](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptComparable[T Comparable](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptOrderable[T Orderable](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}

func acceptAddable[T Addable](t *testing.T, value T) {
	t.Helper()

	accept(t, value)
}
