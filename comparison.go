package constraints

// Comparable is a constraint that matches any type that can be compared with the == and != operators.
type Comparable interface {
	comparable
}

// Orderable is a constraint that matches any type that can be ordered with the <, <=, >, and >= operators.
type Orderable interface {
	Integer | Float | ~string
}

// Addable is a constraint that matches any type that can be added with the + operator.
type Addable interface {
	Number | ~string
}
