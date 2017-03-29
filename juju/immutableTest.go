package main

type myStruct struct {
	cars int
}

// You are allowed to initialise an immutable struct
var mazdaCnst = myStruct{4}
var immutableToyota = myStruct{4}
var anImmutableAudi = myStruct{34}

func immutableMain() {
	// And you're allowed to access them and their properties
	_ = mazdaCnst
	_ = immutableToyota
	_ = anImmutableAudi.cars

	// But you are not allowed to assign anything to them
	mazdaCnst = myStruct{34}
	immutableToyota = myStruct{34}
	anImmutableAudi = myStruct{34}
}
