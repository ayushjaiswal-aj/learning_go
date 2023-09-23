package main

func main() {

	iAge := 0

	if (iAge > 0) && (iAge <= 18) {
		println("You are not a adult")
	} else if (iAge > 18) && (iAge <= 60) {
		println("You are a adult")
	} else if (iAge > 60) && (iAge <= 120) {
		println("You are a senior citizen")
	} else {
		println("Invalid age")
	}

	if !true {
		println("This is not executed")
	} else {
		println("This is executed, not true, else block")
	}
}
