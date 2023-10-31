package ktsets

import (
	"fmt"
	"reflect"
	"strings"
)

// Set represents a set of elements. It behaves like set in Math :-) You can put in and remove elements but it is guarateed that
// an element is member of the Set only once - does not matter how many times you put in the same element.
//
// Please note that a Set is always a Set of a specific type of elements. You can not mix types you put into a Set. Therefore
// when you create a Set you need to tell at that moment of init what type of elements the Set will contain
//
// **IMPORTANT!** This Set implementation is not thread safe! It is your responsibility to take care of synchronization...
type Set[T comparable] struct {
	// we use a map under the hood
	_map map[T]bool
}

// Creates a new Set - optionally with the given initial content
func NewSet[T comparable](elements ...T) Set[T] {
	s := Set[T]{_map: map[T]bool{}}
	if len(elements) > 0 {
		s.AddAll(elements...)
	}
	return s
}

// Creates a new Set with initial capacity - optionally with the given initial content
func NewSetWithCapacity[T comparable](capacity int, elements ...T) Set[T] {
	s := Set[T]{_map: make(map[T]bool, capacity)}
	if len(elements) > 0 {
		s.AddAll(elements...)
	}
	return s
}

// Returns a string representation of this Set
// If you placed complex objects in the map then result might look funny.. :-) but works
func (s Set[T]) String() string {
	strs := make([]string, s.Size())
	i := 0
	for element := range s._map {
		strs[i] = fmt.Sprintf("%+v", element)
		i++
	}
	return fmt.Sprintf("ktsets.Set {%v}", strings.Join(strs, ", "))
}

// Returns a clone of this Set.
// Optionally, you can pass in a desired capacity of the clone (obviously, if its smaller than will be corrected). If you pass in more
// arguments than one only the first one is used so no point to do that
// PLEASE NOTE! This method is not doing deep-cloning!
func (s Set[T]) Clone(capacity ...int) Set[T] {
	cloneLen := len(s._map)
	if len(capacity) > 0 && capacity[0] > cloneLen {
		cloneLen = capacity[0]
	}
	contentClone := make(map[T]bool, cloneLen)
	for k, v := range s._map {
		contentClone[k] = v
	}
	return Set[T]{_map: contentClone}
}

func (s Set[T]) Size() int {
	return len(s._map)
}

func (s Set[T]) IsEmpty() bool {
	return len(s._map) == 0
}

// Returns all elements as a Slice from this set - order is not guaranteed!
func (s Set[T]) GetAll() []T {
	elements := make([]T, 0, len(s._map))
	for key := range s._map {
		elements = append(elements, key)
	}
	return elements
}

// Adds an element to this set - return true if it was really added or false if it was already in the set so not added
func (s Set[T]) Add(element T) (wasAdded bool) {
	if s._map[element] {
		return
	}
	s._map[element] = true
	wasAdded = true
	return
}

// Adds all given elements to this Set
func (s Set[T]) AddAll(elements ...T) (howManyWasAdded int) {
	for _, element := range elements {
		if s.Add(element) {
			howManyWasAdded++
		}
	}
	return
}

// Tells if the given element is in the set or not
func (s Set[T]) Contains(element T) (contains bool) {
	return s._map[element]
}

// Tells if all listed elements are in the Set
func (s Set[T]) ContainsAll(elements ...T) (containsAll bool) {
	containsAll = false
	for _, element := range elements {
		if !s.Contains(element) {
			return
		}

	}
	containsAll = true
	return
}

// Tells if any of the listed elements is in the Set
func (s Set[T]) ContainsAny(elements ...T) (containsAny bool) {
	containsAny = true
	for _, element := range elements {
		if s.Contains(element) {
			return
		}
	}
	containsAny = false
	return
}

// Makes the Set empty
func (s Set[T]) Clear() {
	for k := range s._map {
		delete(s._map, k)
	}
}

// Removes the given element from the Set - returns TRUE if element was in the Set so really removed
func (s Set[T]) Remove(element T) (wasRemoved bool) {
	wasRemoved = s._map[element]
	delete(s._map, element)
	return
}

// Removes all specified elements from the Set
func (s Set[T]) RemoveAll(elements ...T) (howManyWasRemoved int) {
	for _, element := range elements {
		if s.Remove(element) {
			howManyWasRemoved++
		}
	}
	return
}

// Keeps only those elements in the Set which are present in the givel slice - removes everything else
func (s Set[T]) RetainsAll(elements ...T) (howManyWasRemoved int, howManyWasRetained int) {
	elementsSet := NewSet[T](elements...)
	for _, element := range s.GetAll() {
		if !elementsSet.Contains(element) {
			s.Remove(element)
			howManyWasRemoved++
		} else {
			howManyWasRetained++
		}
	}
	return
}

// Tells if this Set contains the same elements as the Other -> they are the same (by Go equality testing at least)
func (s Set[T]) Equals(other Set[T]) bool {
	return reflect.DeepEqual(s._map, other._map)
}

// Subtracts the given Set from this Set
func (s Set[T]) Subtract(other Set[T]) {
	s.RemoveAll(other.GetAll()...)
}

// This Set will become the Union with the given Set
func (s Set[T]) Union(other Set[T]) {
	s.AddAll(other.GetAll()...)
}

// This Set will become the intersection with the given Set
func (s Set[T]) Intersect(other Set[T]) {
	s.RetainsAll(other.GetAll()...)
}
