package main

func greet(name string) {
	println("Hello, " + name + "!")
	greet2(name)
	println("Getting ready to say bye ...")
	bye()
}

func greet2(name string) {
	println("How are you, " + name + "?")
}

func bye() {
	println("OK, bye!")
}

func main() {
	greet("Alan")
}
