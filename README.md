## Benchmarking the Overhead of Go Generics

This repo includes a basic benchmark analysis of the performance overhead associated with using Go generics. It compares a generic-based stack implementation with a non-generic version. 

The comparison covers 
 - int values and 
 - a simple struct.

## How to Run the Benchmark

```bash
go run main.go
```

## Implementation

The generic stack implementation is based on the following type definition:

```go
type Stack[T any] struct {
  data []T
}
```
This is located in the package github.com/kegliz/gengo/genstack.

The non-generic int stack implementation is based on the following type definition:

```go
type Stack struct {
  data []int
}
```
This is located in the package github.com/kegliz/gengo/intstack.

The non-generic struct stack implementation is based on the following type definition:

```go
type Item struct {
		Name string
		Age  int
	}

type Stack struct {
  data []Item
}
```

This is located in the package github.com/kegliz/gengo/structstack.

The benchmarking code is located in the main package.




