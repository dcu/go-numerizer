# Go Numerizer

Numerizer is a library to help with parsing numbers in natural language from strings (ex forty two). It's a straight port of the `Numerizer` gem [https://github.com/jduff/numerize](https://github.com/jduff/numerize).


## Installation


	$ go get github.com/dcu/go-numerizer

## Usage

	import("github.com/dcu/go-numerizer")
	
	Numerizer.numerize('forty two') // "42"
	Numerizer.numerize('two and a half') // "2.5"
	Numerizer.numerize('three quarters') // "3/4"
	Numerizer.numerize('two and three eighths') // "2.375"
