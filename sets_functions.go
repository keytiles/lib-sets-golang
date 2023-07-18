package ktsets

// Takes arbitrary number of Sets as argument and returns a new Set which is the union of the given ones.
// If no sets are sent as arguments then returns an empty set
func Sets_Union[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	result := sets[0].Clone()
	for i := 1; i < len(sets); i++ {
		result.Union(sets[i])
	}
	return result
}

// Takes arbitrary number of Sets as argument and returns a new Set which is the intersection of the given ones.
// If no sets are sent as arguments then returns an empty set
func Sets_Intersection[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	result := sets[0].Clone()
	for i := 1; i < len(sets); i++ {
		result.Intersect(sets[i])
	}
	return result
}
