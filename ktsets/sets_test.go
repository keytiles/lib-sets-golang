package ktsets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// We will use .Equals method a lot in other tests so better to well test it
func TestSetEquals(t *testing.T) {

	/* Scenario 1
	   empty sets
	*/

	// ---- GIVEN
	emptySet1 := NewSet[string]()
	emptySet2 := NewSetWithCapacity[string](100)

	// ---- WHEN-THEN cases
	// reflexivity
	require.True(t, emptySet1.Equals(emptySet1))
	// symmetric
	require.True(t, emptySet1.Equals(emptySet2))
	require.True(t, emptySet2.Equals(emptySet1))

	/* Scenario 2
	   non empty but equals sets
	   with mixed order
	*/

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSetWithCapacity[string](100, "c", "a", "b")

	// ---- WHEN-THEN cases
	// reflexivity
	require.True(t, set1.Equals(set2))
	// symmetric
	require.True(t, set1.Equals(set2))
	require.True(t, set2.Equals(set1))

	/* Scenario 3
	   non empty, same sized but not equals sets
	*/

	// ---- GIVEN
	set1 = NewSet[string]("a", "b", "c")
	set2 = NewSetWithCapacity[string](100, "c", "a", "d")

	// ---- WHEN-THEN cases
	// reflexivity
	require.False(t, set1.Equals(set2))
	// symmetric
	require.False(t, set1.Equals(set2))
	require.False(t, set2.Equals(set1))

}

func TestSetBasicAddAndRemoveAndSize(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")

	// ---- WHEN
	wasRemoved := set.Remove("b")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("a", "c")))
	require.True(t, wasRemoved)
	require.Equal(t, 2, set.Size())

	// ---- WHEN
	// removing again - should not matter
	wasRemoved = set.Remove("b")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("c", "a")))
	require.False(t, wasRemoved)
	require.Equal(t, 2, set.Size())

	// ---- WHEN
	// adding back
	wasAdded := set.Add("b")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("a", "b", "c")))
	require.True(t, wasAdded)
	require.Equal(t, 3, set.Size())

	// ---- WHEN
	// adding back again - should not matter
	wasAdded = set.Add("b")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("a", "b", "c")))
	require.False(t, wasAdded)
	require.Equal(t, 3, set.Size())

}

func TestSetMultipleAdd(t *testing.T) {

	// ---- GIVEN
	set := NewSetWithCapacity[string](100, "a", "b", "c")

	// ---- WHEN
	// adding multiple elements
	howManyWasAdded := set.AddAll("b", "d", "a", "e")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("a", "b", "c", "d", "e")))
	require.Equal(t, 2, howManyWasAdded)

	// ---- WHEN
	// adding existings
	howManyWasAdded = set.AddAll("a", "b")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("a", "b", "c", "d", "e")))
	require.Equal(t, 0, howManyWasAdded)

}

func TestSetMultipleRemove(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")

	// ---- WHEN
	// adding multiple elements
	howManyWasRemoved := set.RemoveAll("b", "d", "a", "e")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("c")))
	require.Equal(t, 2, howManyWasRemoved)

	// ---- WHEN
	// adding existings
	howManyWasRemoved = set.RemoveAll("a", "b")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("c")))
	require.Equal(t, 0, howManyWasRemoved)

}

func TestSetClear(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")

	// ---- WHEN
	// adding multiple elements
	set.Clear()
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]()))
	require.Equal(t, 0, set.Size())

}

func TestSetRetainsAll(t *testing.T) {

	// ---- GIVEN
	set := NewSetWithCapacity[string](100, "a", "b", "c", "d")

	// ---- WHEN
	// retaining multiple elements
	howManyWasRemoved, howManyWasRetained := set.RetainsAll("b", "d", "a", "e")
	// ---- THEN
	require.True(t, set.Equals(NewSet[string]("a", "b", "d")))
	require.Equal(t, 3, howManyWasRetained)
	require.Equal(t, 1, howManyWasRemoved)

}

func TestSetCloning(t *testing.T) {

	// ---- GIVEN
	set := NewSet[string]("a", "b", "c")
	// ---- WHEN
	clonedSet := set.Clone()
	// ---- THEN
	require.True(t, set.Equals(clonedSet))

	// ---- WHEN
	// manipulating the clone should not affect the original
	clonedSet.Remove("b")
	// ---- THEN
	require.False(t, set.Equals(clonedSet))
	require.True(t, set.Equals(NewSet("a", "b", "c")))

	// ---- WHEN
	// testing the clone with capacity
	clonedSet = set.Clone(100)
	// ---- THEN
	require.True(t, set.Equals(clonedSet))

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
	require.True(t, set1.Equals(NewSet("a", "b", "c")))

	// ---- WHEN
	set1.Union(set2)
	// ---- THEN
	// set2 not changed
	require.True(t, set2.Equals(NewSet("c", "d", "e")))
	require.True(t, set1.Equals(NewSet("a", "b", "c", "d", "e")))
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
	require.True(t, set1.Equals(NewSet("a", "b", "c")))

	// ---- WHEN
	set1.Intersect(set2)
	// ---- THEN
	// set2 not changed
	require.True(t, set2.Equals(NewSet("c", "d", "e")))
	require.True(t, set1.Equals(NewSet("c")))

}

func TestSetSubtract(t *testing.T) {

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSet[string]("c", "d", "e")

	// ---- WHEN
	set1.Subtract(set1)
	// ---- THEN
	// set1 should become empty
	require.Equal(t, 0, set1.Size())

	// ---- GIVEN
	// set1 restored
	set1 = NewSet[string]("a", "b", "c")

	// ---- WHEN
	set1.Subtract(set2)
	// ---- THEN
	// set2 not changed
	require.True(t, set2.Equals(NewSet("c", "d", "e")))
	require.True(t, set1.Equals(NewSet("a", "b")))
}

func TestSetsUnion(t *testing.T) {

	// ---- GIVEN
	set1 := NewSet[string]("a", "b", "c")
	set2 := NewSet[string]("c", "d", "e")

	// ---- WHEN
	union := Union[string](set1, set2)
	// ---- THEN
	// set1 and set2 not changed
	require.True(t, set1.Equals(NewSet("a", "b", "c")))
	require.True(t, set2.Equals(NewSet("c", "d", "e")))
	require.True(t, union.Equals(NewSet("a", "b", "c", "d", "e")))
}
