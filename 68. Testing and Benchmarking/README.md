Go by Example: Testing and Benchmarking

Unit testing is an important part of writing principled Go programs. The testing package provides the tools we need to write unit tests and the go test command runs tests.
	

For the sake of demonstration, this code is in package main, but it could be any package. Testing code typically lives in the same package as the code it tests.




We’ll be testing this simple implementation of an integer minimum. Typically, the code we’re testing would be in a source file named something like intutils.go, and the test file for it would then be named intutils_test.go.
	



A test is created by writing a function with a name beginning with Test.
	


t.Error* will report test failures but continue executing the test. t.Fatal* will report test failures and stop the test immediately.
	


Writing tests can be repetitive, so it’s idiomatic to use a table-driven style, where test inputs and expected outputs are listed in a table and a single loop walks over them and performs the test logic.
	


t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.
	



Benchmark tests typically go in _test.go files and are named beginning with Benchmark. The testing runner executes each benchmark function several times, increasing b.N on each run until it collects a precise measurement.
	



Typically the benchmark runs a function we’re benchmarking in a loop b.N times.

Run all tests in the current project in verbose mode.
	



Run all benchmarks in the current project. All tests are run prior to benchmarks. The bench flag filters benchmark function names with a regexp.