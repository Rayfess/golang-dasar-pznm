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

#### Type Data Number

**Divided into 2 Type :**

- `Integer`
- `Floating Point`

##### Type Data Integer 1

| Type Data | Min Value            | Max Value           |
| :-------- | :------------------- | :------------------ |
| `int8`    | -128                 | 127                 |
| `int16`   | -32768               | 32767               |
| `int32`   | -2147483648          | 2147483647          |
| `int64`   | -9223372036854775808 | 9223372036854775807 |

##### Type Data Integer 2

| Type Data | Min Value | Max Value            |
| :-------- | :-------- | :------------------- |
| `uint8`   | 0         | 255                  |
| `uint16`  | 0         | 65535                |
| `uint32`  | 0         | 4294967295           |
| `uint64`  | 0         | 18446744073709551615 |

##### Type Data Floating Point

| Type Data    | Description                              |
| :----------- | :--------------------------------------- |
| `float32`    | **1.18x 10⁻³⁸** up To **3.4x10³⁸**       |
| `float64`    | **2.23x 10⁻³⁰⁸** up To **3.4x10³⁰⁸**     |
| `complex64`  | float32 but has real and imaginary parts |
| `complex128` | float64 but has real and imaginary parts |

##### Alias

| Type Data | Alias For      |
| :-------- | :------------- |
| `byte`    | uint8          |
| `rune`    | int32          |
| `int`     | Min for int32  |
| `uint`    | Min for uint32 |

#### Type Data Boolean

**Divided into 2 Type :**

- `true`
- `false`

#### Type Data String

**Characters that are included in Quotation Marks. Example:**
<br>

- `"inserthere"`
- `"67"` _still in a string, because the characters are included on quotation marks_

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

### Creating Variable

```go
package main

import "fmt"
func main() {
	var name = "Rhaef"
	fmt.Println(name)
}
```

<br>

**shorthand for declare variable** `:=`

```go
package main

import "fmt"
func main() {
	name := "Rhaef"
	fmt.Println(name)

// dont use := again, because its only needed on first declare
	name = "Jacky"
	fmt.Println(name)
}
```

### Run the App Without Compiling

###### The App MUST BE COMPILED, except for Development stages

```go
go run helloworld.go
```
