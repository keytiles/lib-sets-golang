package ktsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// We will use .Equals method a lot in other tests so better to well test it
func TestSetEquals(t *testing.T) {

	/* Scenario 1
	   empty sets
	*/

	// ---- GIVEN
	emptySet1 := NewSet[string]()
	emptySet2 := NewSet[string]()

	// ---- WHEN-THEN cases
	// reflexivity
	assert.True(t, emptySet1.Equals(emptySet1))
	// symmetric
	assert.True(t, emptySet1.Equals(emptySet2))
	assert.True(t, emptySet2.Equals(emptySet1))

	/* Scenario 2
	   non empty but equals sets
	   with mixed order
	*/

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSet[string]("c", "a", "b")

	// ---- WHEN-THEN cases
	// reflexivity
	assert.True(t, set1.Equals(set2))
	// symmetric
	assert.True(t, set1.Equals(set2))
	assert.True(t, set2.Equals(set1))

	/* Scenario 3
	   non empty, same sized but not equals sets
	*/

	// ---- GIVEN
	set1 = NewSet[string]("a", "b", "c")
	set2 = NewSet[string]("c", "a", "d")

	// ---- WHEN-THEN cases
	// reflexivity
	assert.False(t, set1.Equals(set2))
	// symmetric
	assert.False(t, set1.Equals(set2))
	assert.False(t, set2.Equals(set1))

}

func TestSetBasicAddAndRemoveAndSize(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")

	// ---- WHEN
	wasRemoved := set.Remove("b")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("a", "c")))
	assert.True(t, wasRemoved)
	assert.Equal(t, 2, set.Size())

	// ---- WHEN
	// removing again - should not matter
	wasRemoved = set.Remove("b")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("c", "a")))
	assert.False(t, wasRemoved)
	assert.Equal(t, 2, set.Size())

	// ---- WHEN
	// adding back
	wasAdded := set.Add("b")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("a", "b", "c")))
	assert.True(t, wasAdded)
	assert.Equal(t, 3, set.Size())

	// ---- WHEN
	// adding back again - should not matter
	wasAdded = set.Add("b")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("a", "b", "c")))
	assert.False(t, wasAdded)
	assert.Equal(t, 3, set.Size())

}

func TestSetMultipleAdd(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")

	// ---- WHEN
	// adding multiple elements
	howManyWasAdded := set.AddAll("b", "d", "a", "e")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("a", "b", "c", "d", "e")))
	assert.Equal(t, 2, howManyWasAdded)

	// ---- WHEN
	// adding existings
	howManyWasAdded = set.AddAll("a", "b")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("a", "b", "c", "d", "e")))
	assert.Equal(t, 0, howManyWasAdded)

}

func TestSetMultipleRemove(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")

	// ---- WHEN
	// adding multiple elements
	howManyWasRemoved := set.RemoveAll("b", "d", "a", "e")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("c")))
	assert.Equal(t, 2, howManyWasRemoved)

	// ---- WHEN
	// adding existings
	howManyWasRemoved = set.RemoveAll("a", "b")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("c")))
	assert.Equal(t, 0, howManyWasRemoved)

}

func TestSetRetainsAll(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c", "d")

	// ---- WHEN
	// retaining multiple elements
	howManyWasRemoved, howManyWasRetained := set.RetainsAll("b", "d", "a", "e")
	// ---- THEN
	assert.True(t, set.Equals(NewSet[string]("a", "b", "d")))
	assert.Equal(t, 3, howManyWasRetained)
	assert.Equal(t, 1, howManyWasRemoved)

}

func TestSetCloning(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")
	// ---- WHEN
	clonedSet := set.Clone()
	// ---- THEN
	assert.True(t, set.Equals(clonedSet))

	// ---- WHEN
	// manipulating the clone should not affect the original
	clonedSet.Remove("b")
	// ---- THEN
	assert.False(t, set.Equals(clonedSet))
	assert.True(t, set.Equals(NewSet("a", "b", "c")))
}

func TestSetUnion(t *testing.T) {

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSet[string]("c", "d", "e")

	// reflexivity
	// ---- WHEN
	set1.Union(set1)
	// ---- THEN
	// set1 not changed
	assert.True(t, set1.Equals(NewSet("a", "b", "c")))

	// ---- WHEN
	set1.Union(set2)
	// ---- THEN
	// set2 not changed
	assert.True(t, set2.Equals(NewSet("c", "d", "e")))
	assert.True(t, set1.Equals(NewSet("a", "b", "c", "d", "e")))
}

func TestSetIntersection(t *testing.T) {

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSet[string]("c", "d", "e")

	// reflexivity
	// ---- WHEN
	set1.Intersect(set1)
	// ---- THEN
	// set1 not changed
	assert.True(t, set1.Equals(NewSet("a", "b", "c")))

	// ---- WHEN
	set1.Intersect(set2)
	// ---- THEN
	// set2 not changed
	assert.True(t, set2.Equals(NewSet("c", "d", "e")))
	assert.True(t, set1.Equals(NewSet("c")))

}

func TestSetSubtract(t *testing.T) {

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSet[string]("c", "d", "e")

	// ---- WHEN
	set1.Subtract(set1)
	// ---- THEN
	// set1 should become empty
	assert.Equal(t, 0, set1.Size())

	// ---- GIVEN
	// set1 restored
	set1 = NewSet[string]("a", "b", "c")

	// ---- WHEN
	set1.Subtract(set2)
	// ---- THEN
	// set2 not changed
	assert.True(t, set2.Equals(NewSet("c", "d", "e")))
	assert.True(t, set1.Equals(NewSet("a", "b")))
}

func TestSetsUnion(t *testing.T) {

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSet[string]("c", "d", "e")

	// ---- WHEN
	union := Union[string](set1, set2)
	// ---- THEN
	// set1 and set2 not changed
	assert.True(t, set1.Equals(NewSet("a", "b", "c")))
	assert.True(t, set2.Equals(NewSet("c", "d", "e")))
	assert.True(t, union.Equals(NewSet("a", "b", "c", "d", "e")))
}
