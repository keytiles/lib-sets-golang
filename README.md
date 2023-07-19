# lib-sets-golang

Implementation of Set data structure and a few useful methods

# How to use

It's very simple. The below is just a taste - check the available operations and functions for yourself! ;-)

> **IMPORTANT!** - keep an eye on thread-safety! The currently available `ktsets.Set` implementation is NOT thread safe (intentionally)!  

It is possible in the future we add something like `ktsets.ConcurrentSet` or similar which provides thread-safe implementation but for now we don't need it.

```go

// ---- Creating Sets
//      You need to tell type of elements you keep there
//      You get compiler level type safety in return

// An empty set - containing strings
myStringSet := ktsets.NewSet[string]()
// A Set of integers with initial values
myIntSet := ktsets.NewSet[int](1, 2, 3)

// ---- Some basic operations

// adding values
wasAdded := myStringSet.Add("a string")                      // wasAdded = true
wasAdded = myStringSet.Add("a string")                       // wasAdded = false
howManyWasAdded := myStringSet.AddAll("one", "two", "three") // howManyWasAdded = 3
howManyWasAdded = myStringSet.AddAll("one", "two", "four")   // howManyWasAdded = 1
// removing values
wasRemoved := myStringSet.Remove("a string")               // wasRemoved = true
wasRemoved = myStringSet.Remove("a string")                // wasRemoved = false
howManyWasRemoved := myStringSet.RemoveAll("one", "three") // howManyWasRemoved = 2
howManyWasRemoved = myStringSet.RemoveAll("one", "four")   // howManyWasRemoved = 1


// ---- What is in the Set?

var elements []string
elements = myStringSet.GetAll() // elements: ["two"]

// or Print it? we have .String()
fmt.Printf("myStringSet: %v", myStringSet.String())

// size related operations
stringSetSize := myStringSet.Size() // stringSetSize = 1
if myStringSet.IsEmpty() {
    // we will NOT come here as not empty
}


// ---- Containment checks

// containment check
if myStringSet.Contains("two") {
    // yes we will come in here in this case...
}
containsAll := myStringSet.ContainsAll("one", "two", "three") // containsAll = false
containsAny := myStringSet.ContainsAny("one", "two", "three") // containsAny = true


// ---- Set-Set operations

// Equality test
isEqual := myStringSet.Equals(ktsets.NewSet[string]("one", "two", "three")) // isEqual = false
isEqual = myStringSet.Equals(ktsets.NewSet[string]("two"))                  // isEqual = true

// in-place operations (modifying left hand value)
myOtherStringSet := ktsets.NewSet[string]("apple", "orange")
myStringSet.Union(myOtherStringSet)    // myStringSet changed to: ["two", "apple", "orange"]
myStringSet.Subtract(myOtherStringSet) // myStringSet now goes bacck to: ["two"]

// don't want to screw up "myStringSet"?
// Operate on clone!
myStringSet.Clone().Intersect(myOtherStringSet)


// ---- provided Set-Set operation functions

// Or use Union function - which returns new instance always leaving original Sets untouched
unionSet := ktsets.Union(myStringSet, ktsets.NewSet[string]("apple", "orange"), ktsets.NewSet[string]("strawberry"))
// sets have a toString function
fmt.Printf("union is: %v", unionSet.String())

```

