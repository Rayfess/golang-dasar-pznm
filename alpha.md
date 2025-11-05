# Go Programming Cheatsheet - Essential Patterns & Common Cases

## 🎯 Quick Navigation

- [Basic Syntax](#basic-syntax)
- [Data Structures](#data-structures)
- [Functions](#functions)
- [Error Handling](#error-handling)
- [Common Patterns](#common-patterns)
- [Best Practices](#best-practices)

## Basic Syntax

### Variables & Types

```go
// Declaration
var name string = "John"
age := 25                          // Type inference
const pi = 3.14

// Type conversion
intValue := 42
floatValue := float64(intValue)

// Multiple assignment
a, b := 1, "hello"
```

### Control Flow

```go
// If with short statement
if score := getScore(); score > 90 {
    fmt.Println("A grade")
}

// Switch (no break needed)
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("Weekend coming")
default:
    fmt.Println("Regular day")
}

// For loops
for i := 0; i < 5; i++ { }        // Traditional
for index, value := range slice { } // Range loop
for { }                            // Infinite loop
```

## Data Structures

### Arrays & Slices (Most Common)

```go
// Arrays (fixed size - less common)
var scores [3]int = [3]int{90, 85, 92}

// Slices (dynamic - USE THIS)
numbers := []int{1, 2, 3, 4, 5}
subset := numbers[1:4]              // [2, 3, 4]

// Slice operations
slice := make([]int, 0, 10)         // Pre-allocate capacity
slice = append(slice, 6)            // Add element
copy(dest, src)                     // Copy slices

// Common slice patterns
// Filtering
func filterEven(numbers []int) []int {
    var result []int
    for _, n := range numbers {
        if n%2 == 0 {
            result = append(result, n)
        }
    }
    return result
}

// Transformation
func doubleNumbers(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, n := range numbers {
        result[i] = n * 2
    }
    return result
}
```

### Maps (Essential Operations)

```go
// Creation
scores := map[string]int{
    "Alice": 95,
    "Bob":   89,
}

// Operations
scores["Charlie"] = 91              // Add/Update
delete(scores, "Bob")               // Delete
score, exists := scores["Alice"]    // Check existence

// Common map patterns
// Word frequency counter
func wordFrequency(text string) map[string]int {
    words := strings.Fields(text)
    freq := make(map[string]int)
    for _, word := range words {
        freq[word]++
    }
    return freq
}

// Grouping data
func groupByCategory(items []Item) map[string][]Item {
    groups := make(map[string][]Item)
    for _, item := range items {
        groups[item.Category] = append(groups[item.Category], item)
    }
    return groups
}
```

## Functions

### Essential Function Patterns

```go
// Multiple return values (VERY COMMON)
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

// Variadic functions
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Named return values
func calculate(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return // naked return
}
```

### Advanced Function Patterns

```go
// Function as parameter (callbacks)
func processNumbers(numbers []int, processor func(int) int) []int {
    result := make([]int, len(numbers))
    for i, n := range numbers {
        result[i] = processor(n)
    }
    return result
}

// Closures (stateful functions)
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Usage
c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

## Error Handling

### Essential Error Patterns

```go
// Basic error handling
result, err := someOperation()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// Defer for cleanup (CRITICAL)
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Always executed

    // Process file
    return nil
}

// Custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

## Common Patterns

### Data Processing Pipeline

```go
func processUserData(users []User) []string {
    // Filter active users
    activeUsers := filter(users, func(u User) bool {
        return u.Active
    })

    // Extract emails
    emails := mapSlice(activeUsers, func(u User) string {
        return u.Email
    })

    // Remove duplicates
    return removeDuplicates(emails)
}
```

### Configuration Pattern

```go
type Config struct {
    Port    int
    Timeout time.Duration
}

func NewConfig() *Config {
    return &Config{
        Port:    8080,
        Timeout: 30 * time.Second,
    }
}

func (c *Config) WithPort(port int) *Config {
    c.Port = port
    return c
}
```

### Middleware Pattern

```go
type Handler func(http.ResponseWriter, *http.Request) error

func logMiddleware(next Handler) Handler {
    return func(w http.ResponseWriter, r *http.Request) error {
        start := time.Now()
        err := next(w, r)
        log.Printf("Request %s %s took %v", r.Method, r.URL.Path, time.Since(start))
        return err
    }
}
```

## Best Practices

### Slices

```go
// ✅ GOOD: Pre-allocate when size known
items := make([]string, 0, expectedSize)

// ✅ GOOD: Copy when you need independence
independent := make([]int, len(original))
copy(independent, original)

// ❌ BAD: Creating slices without capacity
var slowSlice []int // This will cause many reallocations
```

### Maps

```go
// ✅ GOOD: Initialize before use
counts := make(map[string]int)

// ✅ GOOD: Check existence
value, exists := m["key"]
if exists {
    // use value
}

// ❌ BAD: Assuming key exists
value := m["key"] // Returns zero value if key doesn't exist
```

### Functions

```go
// ✅ GOOD: Handle errors properly
result, err := someFunction()
if err != nil {
    return err // Handle immediately
}

// ✅ GOOD: Use defer for cleanup
func process() error {
    resource, err := acquireResource()
    if err != nil {
        return err
    }
    defer resource.Close() // Cleanup guaranteed
}

// ❌ BAD: Ignoring errors
result, _ := someFunction() // Error silently ignored
```

## Common Mistakes & Solutions

### 1. Slice Sharing Problem

```go
// ❌ PROBLEM: Slices share underlying array
original := []int{1, 2, 3}
shared := original[1:]
shared[0] = 999 // Also modifies original!

// ✅ SOLUTION: Use copy for independence
original := []int{1, 2, 3}
independent := make([]int, len(original))
copy(independent, original)
independent[0] = 999 // Safe - doesn't affect original
```

### 2. Map Modification During Iteration

```go
// ❌ PROBLEM: Modifying map during iteration
for k, v := range m {
    delete(m, k) // RUNTIME PANIC!
}

// ✅ SOLUTION: Collect keys first
var keys []string
for k := range m {
    keys = append(keys, k)
}
for _, k := range keys {
    delete(m, k)
}
```

### 3. Goroutine Race Conditions

```go
// ❌ PROBLEM: Data race
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i) // All goroutines see final value of i
    }()
}

// ✅ SOLUTION: Pass parameter
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n) // Each goroutine gets its own copy
    }(i)
}
```

## Quick Reference Tables

### Type Zero Values

| Type      | Zero Value |
| --------- | ---------- |
| `int`     | `0`        |
| `float64` | `0.0`      |
| `string`  | `""`       |
| `bool`    | `false`    |
| `slice`   | `nil`      |
| `map`     | `nil`      |
| `pointer` | `nil`      |

### Common Functions Cheatsheet

```go
// Slice operations
len(slice)                    // Length
cap(slice)                    // Capacity
append(slice, elements...)    // Add elements
copy(dest, src)              // Copy slice

// Map operations
len(map)                      // Number of keys
delete(map, key)              // Remove key

// String operations
len(str)                      // String length
strings.Split(str, delimiter) // Split string
strings.Join(slice, delimiter) // Join slice
strconv.Atoi(string)          // String to int
```

This cheatsheet focuses on the 20% of Go that you'll use 80% of the time, with practical examples and real-world patterns. Keep it handy while coding! 🚀
