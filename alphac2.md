# Workspace Overview — belajar-golang-dasar

This document provides a comprehensive overview of each file in the workspace, explaining their purpose, functions, and common use cases in Go programming.

## Project Structure Overview

### Basic Concepts
- `helloworld.go` - Simple "Hello World" program
- `greaterHello.go` - HTTP server example
- `sample.go` - Basic sample program

### Variables and Types
- `variable.go` - Variable declaration examples
- `constant.go` - Constant declaration and usage
- `type_declaration.go` - Custom type declarations
- `conversion.go` - Type conversion examples

### Control Flow
- `if.go` - Conditional statements
- `switch.go` - Switch statement examples
- `for.go` - Loop variations
- `break.go` - Break statement usage
- `continue.go` - Continue statement usage

### Functions
- `function.go` - Basic function declarations
- `function_parameter.go` - Function parameters
- `function_return_value.go` - Return values
- `return_multiple_values.go` - Multiple return values
- `named_return_values.go` - Named return values
- `variadic_function.go` - Variadic functions
- `recursive_function.go` - Recursion examples
- `anonymous_function.go` - Anonymous functions
- `function_as_parameter.go` - Functions as parameters
- `function_as_value.go` - Functions as values
- `closure.go` - Closure examples

### Data Structures
- `array.go` - Fixed-size arrays
- `slice.go` - Dynamic slices
- `map.go` - Key-value maps
- `string.go` - String operations

### Error Handling
- `panic.go` - Panic mechanism
- `recover.go` - Recover from panics
- `defer.go` - Deferred execution

### Operators and Math
- `comparison.go` - Comparison operators
- `boolean.go` - Boolean operations
- `opr_bool.go` - Boolean expressions
- `number.go` - Numeric types
- `math.go` - Mathematical operations

## Detailed File Descriptions

### Core Language Features

#### Arrays, Slices, and Maps
1. **array.go**
   - Purpose: Demonstrates fixed-size array operations
   - Features: Array declaration, initialization, indexing
   - Use cases: When you need fixed-size collections

2. **slice.go**
   - Purpose: Shows dynamic array (slice) operations
   - Features: Slicing, append, capacity management
   - Use cases: Dynamic collections, most Go programs

3. **map.go**
   - Purpose: Key-value pair operations
   - Features: Map creation, access, deletion
   - Use cases: Lookups, counting, caching

#### Functions and Advanced Features
1. **function.go**
   - Purpose: Basic function syntax
   - Features: Declaration, parameters, return values
   - Use cases: Code organization

2. **function_as_parameter.go**
   - Purpose: Higher-order functions
   - Features: Functions as arguments
   - Use cases: Callbacks, middleware

3. **anonymous_function.go**
   - Purpose: Lambda functions
   - Features: Inline function definitions
   - Use cases: One-off callbacks

4. **closure.go**
   - Purpose: Closure behavior
   - Features: Variable capture
   - Use cases: State encapsulation

#### Error Handling
1. **defer.go**
   - Purpose: Resource cleanup
   - Features: Deferred execution
   - Use cases: File handling, cleanup

2. **panic.go** and **recover.go**
   - Purpose: Error handling
   - Features: Panic recovery
   - Use cases: Critical errors

## Common Operations and Patterns

### Slices
```go
// Creation
slice := []int{1, 2, 3}
slice := make([]int, 3, 5) // length 3, capacity 5

// Append
slice = append(slice, 4)

// Slicing
sub := slice[1:3]

// Copy
dest := make([]int, len(src))
copy(dest, src)
```

### Maps
```go
// Creation
m := make(map[string]int)
m := map[string]int{"one": 1}

// Operations
m["key"] = value
value, exists := m["key"]
delete(m, "key")
```

### Functions
```go
// Basic function
func add(x, y int) int {
    return x + y
}

// Multiple return values
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## Common Use Cases and Examples

### Slice Operations
```go
// Example: Dynamic list management
func addUser(users []string, user string) []string {
    return append(users, user)
}

// Example: Filtering
func filterPositive(numbers []int) []int {
    var result []int
    for _, n := range numbers {
        if n > 0 {
            result = append(result, n)
        }
    }
    return result
}
```

### Map Operations
```go
// Example: Counting occurrences
func countWords(words []string) map[string]int {
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++
    }
    return counts
}

// Example: Grouping
func groupByFirstLetter(words []string) map[string][]string {
    groups := make(map[string][]string)
    for _, word := range words {
        firstLetter := string(word[0])
        groups[firstLetter] = append(groups[firstLetter], word)
    }
    return groups
}
```

### Function Patterns
```go
// Example: Middleware pattern
type HandlerFunc func(string) string

func logMiddleware(next HandlerFunc) HandlerFunc {
    return func(input string) string {
        fmt.Printf("Processing: %s\n", input)
        result := next(input)
        fmt.Printf("Result: %s\n", result)
        return result
    }
}

// Example: Builder pattern with closure
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Best Practices

1. Slices
   - Pre-allocate with make() when size is known
   - Use append() for dynamic growth
   - Consider copy() for slice duplication

2. Maps
   - Always initialize with make()
   - Check existence with two-value assignment
   - Delete unused keys to prevent memory leaks

3. Functions
   - Keep functions focused and small
   - Use named return values for clarity
   - Prefer multiple return values over complex structs

## Common Mistakes to Avoid

1. Slices
   - Forgetting that slices share underlying arrays
   - Not checking capacity before append
   - Incorrect slice bounds

2. Maps
   - Forgetting to initialize with make()
   - Not handling missing keys
   - Concurrent access without synchronization

3. Functions
   - Ignoring returned errors
   - Unnecessary pointer parameters
   - Not documenting side effects

This overview provides a solid foundation for working with Go's core features. Refer to the individual files for more detailed examples and implementations.
