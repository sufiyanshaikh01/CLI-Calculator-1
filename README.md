# CLI-Calculator-1
# Using Go Lang command line parsing in built library
Go programs can be witten using only the standard library(also called system in-built libraries).These are package that come the Go installation and do not require any external dependencies and run such program directly using [go run] or compile them using [go build]

# System in-built library
*[fmt]- for printing and formatting
*[flag]- for parsing command line argument
These are part of the Go Standard Library

1. The Os loads the binary into memory.
2. The [main()] funcation is called -this is the program's entry point.
3. [flag.parse()] reads argument like [-op=add, -n1=5, -n2=3]
4. The [switch] statement check the opration and performs the calculation.
5. [fmt.Printf()] prints the result to the console.

# How to Run the Program
1. Run directly using go run:
go run main.go -op=add -n1=10 -n2=5
#Output: Result: 15.00

2. Or commpile and run the binary:
go built main.go
./main -op=sub -n1=8 -n2=3
#OutPut: Result: 5.00
