package examples

import (
	"errors"
	"fmt"
)

func VariablesAndCollections() {
	fmt.Println("---- Variables and collections ----")
	variablesExample()
	arrayExampleResult := arrayExample()
	slicesExample(arrayExampleResult)
	mapsExample()
	structsExample()
	err := returningErrorsExample()
	fmt.Println(err)
	// The _ allows to to take a result you don't use
	_, err = returningMultipleValues()
}

func variablesExample() {
	// Showing the memory location of an element (reference)
	firstName := "David"
	fmt.Printf("The memory location of firstName %v (type %T) is %v\n\n", firstName, firstName, &firstName)
	// Output: The memory location of firstName David (type string) is 0xc000096220

	// Creating a variable using dereferencing
	var a *string = new(string)
	fmt.Printf("A string created using dereferencing is %v (type %T) at %v\n", a, a, &a)
	// Output: A string created using dereferencing is 0xc000096250 (type *string) at 0xc0000b2020
	// Even when assigning, "a" is just a reference to a memory location
	*a = "Hello"
	fmt.Printf("After assigning a string, it's still %v (type %T) at %v\n", a, a, &a)
	// Output: After assigning a string, it's still 0xc000096250 (type *string) at 0xc0000b2020
	// Need to dereference to get the actual value
	fmt.Printf("After dereferencing, it's now %v (type %T) at %v\n\n", *a, *a, &a)
	// Output: After dereferencing, it's now Hello (type string) at 0xc0000b2020
}

func arrayExample() [3]int {
	/*
		Arrays (lists that aren't resizeable)
	*/

	// The [3] denotes the number of elements in the array
	arr := [3]int{1, 3, 5}
	fmt.Printf("Array is %v (type %T) at %v\n\n", arr, arr, &arr[0]) // Output: Array is [1 3 5] (type [3]int) at 0xc0000bc000
	return arr
}

func slicesExample(arr [3]int) {
	/*
		Slices (an array that's resizeable like a python list)
	*/

	// When copying the array to a slice, the slice just points at the array rather than creating a new one.
	sliceFromArray := arr[:]
	fmt.Printf("sliceFromArray is %v (type %T) at %v\n", sliceFromArray, sliceFromArray, &sliceFromArray[0])
	// Output: sliceFromArray is [1 3 5] (type []int) at 0xc0000bc000

	// Any changes to the slice will create a new value
	sliceFromArray = append(sliceFromArray, 7)
	fmt.Printf("Appended sliceFromArray is %v (type %T) at %v\n", sliceFromArray, sliceFromArray, &sliceFromArray[0])
	// Output:  Appended sliceFromArray is [1 3 5 7] (type []int) at 0xc0000ba030
	fmt.Printf("Array is now %v (type %T) at %v\n\n", arr, arr, &arr[0])
	// Output: Array is now [1 3 5] (type [3]int) at 0xc0000bc00

	// [] denotes a slice over an array
	slice := []int{7, 8, 9}
	fmt.Printf("The initial slice is %v (type %T) at %v\n", slice, slice, &slice[0])
	// Output: The initial slice is [7 8 9] (type []int) at 0xc0000bc078
	slice = append(slice, 10, 11, 12, 56, 88)
	fmt.Printf("The appended slice is %v (type %T) at %v\n\n", slice, slice, &slice[0])
	// Output: The appended slice is [7 8 9 10 11 12 56 88] (type []int) at 0xc0000ac040

	sliceWithLastElementRemoved := slice[:len(slice)-1]
	fmt.Printf("The sliceWithLastElementRemoved is %v (type %T) at %v\n\n", sliceWithLastElementRemoved, sliceWithLastElementRemoved, &sliceWithLastElementRemoved[0])
	// Output: The sliceWithLastElementRemoved is [7 8 9 10 11 12 56] (type []int) at 0xc0000ac040
}

func mapsExample() {
	/*
		maps (like a python dictionary)
	*/

	// [string] is the key type, int is the value type
	m := map[string]int{"foo": 1, "bar": 3}
	fmt.Printf("The initial map is %v (type %T)\n", m, m) // Output: The initial map is map[bar:3 foo:1] (type map[string]int)

	m["moo"] = 7
	fmt.Printf("The appended map is %v (type %T)\n", m, m) // Output: The appended map is map[bar:3 foo:1 moo:7] (type map[string]int)

	delete(m, "foo")
	fmt.Printf("The map with an element deleted is %v (type %T)\n\n", m, m)
	// Output: The map with an element deleted is map[bar:3 moo:7] (type map[string]int)
}

func structsExample() {
	/*
		Structs (way of associating different types of data together)
	*/

	// Creating a struct called "user"
	// First, define the fields the struct is going to contain (could be defined at package level too)
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	// Then, create a variable
	var u user
	// Everything is created with its default value
	fmt.Printf("A struct created without data is %v (type %T)\n", u, u) // A struct created without data is {0  } (type main.user)
	// Adding data to the struct
	u.ID = 1
	u.FirstName = "John"
	u.LastName = "Smith"
	fmt.Printf("A struct with data is %v (type %T)\n", u, u)                        // Output: A struct with data is {1 John Smith} (type main.user)
	fmt.Printf("The firstName element is %v (type %T)\n", u.FirstName, u.FirstName) // Output: The firstName element is John (type string)
	// You can also create the value in a single statement (implicit initialisation)
	u2 := user{ID: 2, FirstName: "Jill", LastName: "Jones"}
	fmt.Printf("A struct with data added implictly is %v (type %T)\n", u2, u2) // Output: A struct with data added implictly is {2 Jill Jones} (type main.user)
	// Can also create on multiple lines (remember to add the comma after the last element or the complier will add a semicolon)
	u3 := user{
		ID:        3,
		FirstName: "Percy",
		LastName:  "Pea",
	}
	fmt.Printf("A struct with data added implictly with nice formatting is %v (type %T)\n", u3, u3)
	// Output: A struct with data added implictly with nice formatting is {3 Percy Pea} (type main.user)
}

func returningErrorsExample() error {
	fmt.Println("returningErrorsExample()")
	return errors.New("Something went wrong.")
}

func returningMultipleValues() (int, error) {
	value := 1
	return value, nil
}
