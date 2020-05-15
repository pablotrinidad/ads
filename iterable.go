package ads

// Iterable provides an enumeration strategy for collections.
type Iterable interface {
	// Scan returns a boolean indicating if there's a next element or not.
	Scan() bool
	// Next returns the next element in the iterable.
	Next() (interface{}, error)
}
