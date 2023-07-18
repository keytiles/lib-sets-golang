# lib-sets-golang

Implementation of Set data structure and a few useful methods

# How to use

It's very simple. The below is just a taste - check the available operations and functions for yourself! ;-)

```go

    // An empty set - containing strings
    myStringSet := ktsets.NewSet[string]()

    // A Set of integers with initial values
    myIntSet := ktsets.NewSet[int](1, 2, 3)

    // adding values
    wasAdded := myStringSet.Add("a string")
    howManyWasAdded := myStringSet.AddAll("one", "two", "three")

    if myStringSet.Contains("two") {
        // yes we will come in here in this case...
    }

    // And some typical set operations
    unionSet := ktsets.Union(myStringSet, ktsets.NewSet[string]("apple", "orange"))
    // sets have a toString function
    fmt.Printf("union is: %v", unionSet.String())

```

