# golang-dasar-pznm

## Sources

**Slides GO Materials**
<br>
https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit

**Main Sources Code of GO Material**
<br>
https://github.com/ProgrammerZamanNow/belajar-golang-dasar-2023

## Appendix

- **In one package, there MUST NOT be the same function.**

### Type Data Map

Type data Map is a collection of key-value data, where the keys are unique and cannot be duplicated.

It is very flexible because it can contain as much data as possible.

If there are duplicate keys, the keys and values will overwrite with the new one.

##### Function Map

| Operation                     | Desc                            |
| :---------------------------- | :------------------------------ |
| `len(map)`                    | getting total data in map       |
| `map[key]`                    | taking data in map by key       |
| `make[key] = value`           | changing data in map by key     |
| `make(map[TypeKey]TypeValue)` | making new map                  |
| `delete(map,key) `            | deleting all data in map by key |

### Type Data Slice

**Array** = stiff Container with fixed Capacity

**Slice** = flexible and dynamic Container

- **Type Data Slice is fragment from data Array**, Simillar to Array, but the things make it diffrence was size, Slice is changeable or dynamic
- **Slice and Array always connected**, where Slice was accessing data to partial or complete
- Has 3 Main Data, Pointer, length, and Capacity
  - Pointer is first pointer in an array slice
  - Length is length of the slice
  - Capacity of the slice, which length musn't more than Capacity

##### Creating Slice From Array

|  Creating Slice   |                              Description                               |
| :---------------: | :--------------------------------------------------------------------: |
| `array[low:high]` |  Creating Slice from Array starts from index Low to index before High  |
|   `array[low:]`   | Creating Slice from Array starts from index Low to last index in Array |
|  `array[:high]`   |   Creating Slice from Array starts from index 0 to index before High   |
|    `array[:]`     |  Creating Slice from Array starts from index 0 to last index in Array  |

##### Function Slice

|              Operation               |                                                   Description                                                   |
| :----------------------------------: | :-------------------------------------------------------------------------------------------------------------: |
|             `len(slice)`             |                                               For getting Length                                                |
|             `cap(slice)`             |                                              For getting Capacity                                               |
|        `append(slice, data)`         | Making new Slice that adds data to last position Slice, <br> if the capacity is full, make an new array instead |
| `make([]TypeData, length, capacity)` |                                                Making new Slice                                                 |
|     `copy(destination, source)`      |                                    Copying slice from source to destination                                     |

##### Type Data Number

Divided into 2 Type :

- `Integer`
- `Floating Point`

##### Type Data Integer 1

| Type Data |      Min Value       |      Max Value      |
| :-------: | :------------------: | :-----------------: |
|  `int8`   |         -128         |         127         |
|  `int16`  |        -32768        |        32767        |
|  `int32`  |     -2147483648      |     2147483647      |
|  `int64`  | -9223372036854775808 | 9223372036854775807 |

##### Type Data Integer 2

| Type Data | Min Value |      Max Value       |
| :-------: | :-------: | :------------------: |
|  `uint8`  |     0     |         255          |
| `uint16`  |     0     |        65535         |
| `uint32`  |     0     |      4294967295      |
| `uint64`  |     0     | 18446744073709551615 |

##### Type Data Floating Point

|  Type Data   |               Description                |
| :----------: | :--------------------------------------: |
|  `float32`   |    **1.18x 10⁻³⁸** up To **3.4x10³⁸**    |
|  `float64`   |   **2.23x 10⁻³⁰⁸** up To **3.4x10³⁰⁸**   |
| `complex64`  | float32 but has real and imaginary parts |
| `complex128` | float64 but has real and imaginary parts |

##### Alias

| Type Data |   Alias For    |
| :-------: | :------------: |
|  `byte`   |     uint8      |
|  `rune`   |     int32      |
|   `int`   | Min for int32  |
|  `uint`   | Min for uint32 |

##### Type Data Boolean

Divided into 2 Type :

- `true`
- `false`

##### Type Data String

Characters that are included in Quotation Marks. Example:

- `"inserthere"`
- `"67"` _still in a string, because the characters are included on quotation marks_

#### Augmented Assignments

| Math Operations | Augmented Assignment |
| :-------------: | :------------------: |
|  `a = a + 10`   |      `a += 10`       |
|  `a = a - 10`   |      `a -= 10`       |
|  `a = a * 10`   |      `a *= 10`       |
|  `a = a / 10`   |      `a /= 10`       |
|  `a = a % 10`   |      `a %= 10`       |

#### Unary Operator

| Operator |   Description   |
| :------: | :-------------: |
|   `++`   |    Increment    |
|   `--`   |    Decrement    |
|   `-`    |    Negative     |
|   `+`    |    Positive     |
|   `!`    | Reverse Boolean |

#### Comparison Operator

| Operator | Description  |
| :------: | :----------: |
|   `>`    |  More Than   |
|   `<`    |  Less Than   |
|   `>=`   | More or Same |
|   `<=`   | Less or Same |
|   `==`   |   Same as    |
|   `!=`   | Not Same as  |

#### Boolean Operator

| Operator | Description |
| :------: | :---------: |
|   `&&`   |     And     |
|   `⏐⏐`   |     Or      |
|   `!`    |  Opposite   |

##### `&&` Operator

| Value 1 | Operator | Value2  | Result  |
| :-----: | :------: | :-----: | :-----: |
| `true`  |   `&&`   | `true`  | `true`  |
| `true`  |   `&&`   | `false` | `false` |
| `false` |   `&&`   | `false` | `true`  |
| `false` |   `&&`   | `true`  | `false` |

##### `⏐⏐` Operator

| Value 1 | Operator | Value2  | Result  |
| :-----: | :------: | :-----: | :-----: |
| `true`  |   `⏐⏐`   | `true`  | `true`  |
| `true`  |   `⏐⏐`   | `false` | `true`  |
| `false` |   `⏐⏐`   | `true`  | `true`  |
| `false` |   `⏐⏐`   | `false` | `false` |

##### `!` Operator

| Operator |  Value  | Result  |
| :------: | :-----: | :-----: |
|   `!`    | `true`  | `false` |
|   `!`    | `false` | `true`  |

##### Index in Array

| Index | Sample Data |
| :---: | :---------: |
|  `0`  |   `Ahmad`   |
|  `1`  |   `Bryan`   |
|  `2`  |  `Cassie`   |

#### Function Array

|       Operation        |       Description       |
| :--------------------: | :---------------------: |
|      `len(array)`      | Getting Length of Array |
|    `arrray[index]`     |      Getting Data       |
| `array[index] = value` |      Changing data      |

## Usage

### Creating Module

```go
go mod init project-name
```

### Compile the program to app

##### The Output of Compiling Build is Correspond to the Module Name

```go
go build
```

### Run the App

```go
// if the program is run on windows
belajar-golang-dasar.exe

// if the program is run on mac/linux, Unix System
./belajar-golang-dasar

// sample output
Hello World!
```

### Declaring Variable

Divided into 2 Type :

- `const` for declaring only one time
- `var` usefull for many uses of creating multiple variables

#### Var Use Case

```go
	var name = "Rhaef"
	fmt.Println(name)
```

<br>

**shorthand for declare variable** `:=`

```go
	name := "Rhaef"
	fmt.Println(name)

  // dont use := again, because its only needed on first declare
	name = "Jacky"
	fmt.Println(name)
```

#### Const Declaring

```go
	// Declare using Const, cannot redeclare value
	const age = 17
    age = 20 // will cause error because rules that CONST, cannot be redeclared
	fmt.Println(age)
```

#### Diffrence Between Const and Var

`const`

**Cannot CHANGE VALUE, it will cause error**

`var`

**Can CHANGE VALUE, so more dynamic use**

### Run the App Without Compiling

###### The App MUST BE COMPILED, except for Development stages

```go
go run helloworld.go
```
