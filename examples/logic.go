package examples

import "fmt"

func LoopExamples() {
	fmt.Println("---- Loops ----")
	loopThatTerminatesBasedOnACondition()
	loopWithAPostClause()
	infiniteLoop()
	loopingOverCollections()
	//panicking()
}

func LogicExamples() {
	fmt.Println("---- Logic ----")
	ifStatements()
	switchStatements()
}

func loopThatTerminatesBasedOnACondition() {
	// A loop that terminates based on a condition
	var i int
	for i < 10 {
		println(i)
		i++

		// To short-circuit, use a break statement
		if i == 6 {
			break
		}

		// To skip the rest of a loop depending on a condition, use continue
		if i == 4 {
			continue
		}
		println("Continuing...")
	}
	println()
}

func loopWithAPostClause() {
	// Like an old-fashioned for statement
	for i := 0; i < 5; i++ {
		println(i)
	}
	println()
}

func infiniteLoop() {
	// The ugly way
	var i int
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
	println()

	// The proper way
	var j int
	for {
		if j == 5 {
			break
		}
		println(j)
		j++
	}
	println()
}

func loopingOverCollections() {
	// Over a slice
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	// Over a map
	mappyMcMapface := map[string]int{"hello": 1, "world": 2}
	for j, w := range mappyMcMapface {
		println(j, w)
	}
}

func panicking() {
	println("Starting web server")
	// Stops the programme and prints out a bit of a stacktrace
	panic("Something bad happened")
	println("Stopping web server")
}

func ifStatements() {
	a := 3
	b := 2

	if a < b {
		println("a is less than b")
	} else if a > b {
		println("a is greater than b")
	} else {
		println("a and b are the same")
	}
}

func switchStatements() {
	a := 1

	switch a {
	case 1:
		println("It's 1")
		// No break required - it's implicit (use "fallthrough" if you want to do this)
	case 2:
		println("It's 2")
	default:
		print("I'm a default case")
	}
}
