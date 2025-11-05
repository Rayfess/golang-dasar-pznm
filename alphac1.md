# Go Programming Cheatsheet

This cheatsheet focuses on commonly used features in Go with practical examples and their outputs.

## Arrays

```go
// Fixed-size array declaration
var scores [3]int = [3]int{90, 85, 92}

// Shorthand declaration
grades := [...]int{95, 88, 78}

// Access and modification
scores[1] = 89
```

Example output:
```
scores: [90 89 92]
grades: [95 88 78]
```

## Slices

```go
// Creating slices
numbers := []int{1, 2, 3, 4, 5}
subset := numbers[1:4]  // [2 3 4]

// Using make
slice := make([]int, 3, 5)  // length 3, capacity 5

// Append
slice = append(slice, 6)

// Copy
dest := make([]int, len(numbers))
copy(dest, numbers)
```

Example output:
```
numbers: [1 2 3 4 5]
subset: [2 3 4]
slice before append: [0 0 0]
slice after append: [0 0 0 6]
dest after copy: [1 2 3 4 5]
```

## Maps

```go
// Creating maps
scores := map[string]int{
    "Alice": 95,
    "Bob":   89,
}

// Using make
grades := make(map[string]int)

// Operations
scores["Charlie"] = 91        // Add
delete(scores, "Bob")         // Delete
score, exists := scores["Alice"] // Check existence

// Iteration
for name, score := range scores {
    fmt.Printf("%s: %d\n", name, score)
}
```

Example output:
```
Initial scores: map[Alice:95 Bob:89]
After adding Charlie: map[Alice:95 Bob:89 Charlie:91]
After deleting Bob: map[Alice:95 Charlie:91]
Alice exists: true, score: 95
Iteration:
Alice: 95
Charlie: 91
```

## Functions

### Basic Functions
```go
func add(x, y int) int {
    return x + y
}

result := add(5, 3)
```

Output:
```
8
```

### Multiple Return Values
```go
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

result, err := divide(10, 2)
```

Output:
```
result: 5, err: nil
```

### Variadic Functions
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

total := sum(1, 2, 3, 4)
```

Output:
```
10
```

### Function as Parameter
```go
func applyOperation(x int, operation func(int) int) int {
    return operation(x)
}

double := func(x int) int {
    return x * 2
}

result := applyOperation(5, double)
```

Output:
```
10
```

### Closures
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

Output:
```
1
2
```

## Common Use Cases

### Slice Processing
```go
// Filtering slice
func filterPositive(numbers []int) []int {
    var result []int
    for _, n := range numbers {
        if n > 0 {
            result = append(result, n)
        }
    }
    return result
}

nums := []int{-1, 2, -3, 4, 5}
positives := filterPositive(nums)
```

Output:
```
Original: [-1 2 -3 4 5]
Filtered: [2 4 5]
```

### Map Usage
```go
// Word frequency counter
func wordCount(text string) map[string]int {
    words := strings.Fields(text)
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++
    }
    return counts
}

text := "apple banana apple cherry banana"
freq := wordCount(text)
```

Output:
```
Word frequencies: map[apple:2 banana:2 cherry:1]
```

### Function Composition
```go
func pipeline(x int) int {
    double := func(n int) int { return n * 2 }
    addOne := func(n int) int { return n + 1 }
    
    result := addOne(double(x))
    return result
}

value := pipeline(5)
```

Output:
```
11  // (5 * 2) + 1
```

## Best Practices

1. Slices
   - Pre-allocate when size is known: `make([]int, 0, expectedSize)`
   - Use `copy()` instead of re-slicing when you need a new backing array
   - Check `len()` before accessing elements

2. Maps
   - Always initialize with `make()` before use
   - Use the two-value form of map lookup to check existence
   - Clear maps by reassigning: `m = make(map[string]int)`

3. Functions
   - Keep functions small and focused
   - Return early for error conditions
   - Use named return values for clarity in larger functions

## Common Mistakes to Avoid

### Slices
❌ Don't forget slices share backing arrays
```go
original := []int{1, 2, 3}
shared := original[1:] // Changes to shared affect original
```

✅ Use copy when you need independence
```go
original := []int{1, 2, 3}
independent := make([]int, len(original))
copy(independent, original)
```

### Maps
❌ Don't forget to check existence
```go
value := m["key"] // Might panic if key doesn't exist
```

✅ Always check for existence
```go
value, exists := m["key"]
if exists {
    // use value
}
```

### Functions
❌ Don't ignore error returns
```go
result, _ := someFunction() // Bad: ignoring error
```

✅ Always handle errors
```go
result, err := someFunction()
if err != nil {
    // handle error
}
```

This cheatsheet provides quick reference for common Go patterns with practical examples and their outputs. Use it alongside the main documentation for more detailed explanations.
