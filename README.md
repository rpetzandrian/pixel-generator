# Pixel Generator
This program is written in Go and serves as a pixel generator, designed to create visual elements by manipulating individual pixels. It uses either a recursive approach or loop-based logic to generate and display patterns or images. The flexibility of switching between recursion and iteration makes it adaptable to various use cases and optimization scenarios.

## Group Member
> Rico Andrian Firmansyah (103042310035)<br>
> Daffa Naufal R (103042310019)<br>
> Dimas Adriansyah (103042310038)<br>

## How to run
```
$ go run . <size> <recursize/loop>
```

## Example

### Run as Recursive
```
$ go run . 20 recursive

 ░▓ ▓ ░   ▓░  ▓ ▒░░ 
░ ░░▓ ▓▒▒░▓░░▓   ▒▒▒
░▒▒▓▓░ ▒░▓▒░▓▒ ░▒  ░
░░ ▒ ▓░   ░▒ ░ ░▓▓  
░▓░▓ ▒▒░ ▓▒ ▓▒ ░░░ ▓
░  ▓░▓▒▒▒   ▓░▒ ▓▒▒▒
▒▓▓░░▒▒ ░░▓▒░▓▒░▒░▒▓
▒ ░▒ ░▒ ▒░▒░░▒░▓▒▒▒░
 ▒▓░▓░ ▒▒▓▓▓ ▒▓  ▓▒░
▓  ▓░▒ ▒▓ ▒▒▓░░░░▒▒ 
▓░░▒▓ ░░ ▒▒░▓░░▒▓░▓ 
▓ ▓▓░░▓ ░ ▒▒ ▓▓░▒▓▒ 
 ░  ▒▒▓░ ▒░ ▓▓▒▒   ░
   ▒░▒▓▓▓▒░▒ ░▒░  ▒
 ▓░▒ ░▒ ▒▒▓░▓ ░▓ ▒▓
▒░▓ ▓░ ▓░▒ ▓▓▓▓ ░░▒▒
 ▒▓▓   ░▒░▓░░░▒░░▓▒▒
 ▓▓▓▒▒  ▓ ░▒ ░▓▒▓▓░░
░▓▓░░░ ▓░ ░ ▒  ▓▓▒
▓ ▓▒▓▒▓▒▓▒  ░▒▓▒▓ ▒▓
Elapsed 21.0849ms
```

### Run as Loop
```
$ go run . 20 Loop

▒▒░▒ ▒▒▒░   ░░▓▒▒▓▓▓
▒▒▒▓  ▒░ ░░▓▒░▒░ ▒▓
▓▓▒░░▓▓░▒░░░▓░▒▒  ░░
▓▒░▒▒▓▓ ▓▒░▒ ▒▒░▓▓
▒  ▓░▒▒  ▒▓▓░▒ ░▒░▒▒
▓░▒▓▓▒ ▒░░▓░▒▒▒  ▓▒▓
░░░▓ ░▓▒ ░ ░   ▒░ ░░
░░▓▓▓▒ ░░▒░▒ ▒░░▒▒▓░
░   ░ ░░░▒▓ ▓░▓▒  ▒▒
░▒░ ▓▒ ▓▓░▒▒ ▓▒ ▒░▒▒
▓▒   ▒▒ ░▒░▓▒▒▒ ▓▒ ▒
▒▓▒░▓  ▒▒░░▓ ▒▓▒  ▓ 
░░░▓▒░▒▒▒▒░▓  ▓▓▓▓▒
▓░  ▓▒▒▒ ▒ ▒▓▒▓▓▒▓▓▒
   ▓▒▒▓▓ ░░▒░▒░▓░ ░
 ▒ ▒▓▓░ ▓░  ░▒ ▒░░▓░
 ▓▓▓▒▒▒▒ ▓▒  ▓▓ ▒░░▒
▓▓ ▒  ▓░▒▒▓▓▓▒  ▒▓ ▓
  ▓▒░░▓▓░▒▓▒▓  ░▒▓ ▒
░▓░▒▓░ ▓▓ ▒▒ ▓ ▒░▓▓▓
Elapsed 106.2274ms
```