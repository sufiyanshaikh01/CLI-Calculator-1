# CLI-Calculator-1
# Using Go Lang command line parsing in built library
Go programs can be witten using only the standard library(also called system in-built libraries).These are package that come the Go installation and do not require any external dependencies and run such program directly using [go run] or compile them using [go build]

# System in-built library
*[fmt]- for printing and formatting<br>
*[flag]- for parsing command line argument<br>
These are part of the Go Standard Library<br>

1. The Os loads the binary into memory.<br>
2. The [main()] funcation is called -this is the program's entry point.<br>
3. [flag.parse()] reads argument like [-op=add, -n1=5, -n2=3]<br>
4. The [switch] statement check the opration and performs the calculation.<br>
5. [fmt.Printf()] prints the result to the console.<br>

# How to Run the Program
1. Run directly using go run:<br>
go run main.go -op=add -n1=10 -n2=5<br>
#Output: Result: 15.00<br>

2. Or commpile and run the binary:<br>
go built main.go<br>
./main -op=sub -n1=8 -n2=3<br>
#OutPut: Result: 5.00<br>

# What is a Unit Test?
A unit test checks the correctness of a small piece of code — usually a single function. It helps catch bugs early and ensures code reliability during changes.
