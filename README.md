# Piscine Go — Learn2Earn Bootcamp Solutions

![Go](https://img.shields.io/badge/Go-1.22.2-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Exercises](https://img.shields.io/badge/Exercises-140%2B-brightgreen.svg)

A comprehensive collection of **140+ Go programming exercises and projects** completed during the **Learn2Earn Piscine** bootcamp.

## Table of Contents
- [Core Library Exercises](#core-library-exercises)
- [Command-Line Programs](#command-line-programs)
- [Shell Scripts](#shell-scripts)
- [Project Structure](#project-structure)
- [Learning Path](#learning-path)
- [Key Concepts](#key-concepts-covered)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Core Library Exercises

All root-level `.go` files are library functions (package `piscine`).

### String Manipulation
| File | Description |
|------|-------------|
| `strlen.go` | String length |
| `strrev.go` | Reverse string |
| `split.go` | Split by delimiter |
| `splitwhitespaces.go` | Split by whitespace |
| `join.go` | Join strings with separator |
| `concat.go` | Concatenate strings |
| `concatparams.go` | Concatenate CLI args |
| `capitalize.go` | Capitalize words |
| `tolower.go` / `toupper.go` | Case conversion |
| `rot14.go` | ROT14 cipher |
| `alphacount.go` | Count alphabetic chars |
| `compare.go` | Lexicographic comparison |
| `index.go` | Substring index |
| `nrune.go` / `firstrune.go` / `lastrune.go` | Rune access |
| `printstr.go` | Print without fmt |
| `printwordstables.go` | Word table formatter |

### Numeric Conversion
| File | Description |
|------|-------------|
| `atoi.go` | String to integer (full) |
| `basicatoi.go` / `basicatoi2.go` | Simple atoi variants |
| `atoibase.go` | Base conversion to int |
| `convertbase.go` | Between arbitrary bases |
| `trimatoi.go` | Trim + convert |
| `printnbr.go` / `printnbrbase.go` | Print integers |
| `divmod.go` / `ultimatedivmod.go` | Division & modulus |
| `pointone.go` / `ultimatepointone.go` | Float handling |
| `any.go` | Any element satisfies predicate |

### Type Checks
| File | Description |
|------|-------------|
| `isalpha.go` | Letter? | `isnumeric.go` | Digit? |
| `isupper.go` | Uppercase? | `islower.go` | Lowercase? |
| `isprintable.go` | Printable? | `isprime.go` | Prime? |
| `isnegative.go` | Negative? | `issorted.go` | Sorted? |
| `unmatch.go` | Find unique element |

### Sorting
| File | Description |
|------|-------------|
| `sortwordarr.go` / `advancedsortwordarr.go` | String sort |
| `sortintegertable.go` | Integer sort |
| `sortlistinsert.go` / `sortedlistmerge.go` | Sorted list ops |
| `listsort.go` | List bubble sort |
| `revparams.go` / `sortparams.go` | Arg reordering |

### Arrays & Ranges
| File | Description |
|------|-------------|
| `compact.go` | Remove adjacent duplicates |
| `appendrange.go` / `descendappendrange.go` | Range append |
| `makerange.go` | Integer range |
| `countif.go` / `foreach.go` / `any.go` | Functional ops |
| `map.go` | Map values |
| `stringtointslice.go` | String to []int |
| `max.go` | Find max |

### Combinatorics
| File | Description |
|------|-------------|
| `printcomb.go` / `printcomb2.go` / `printcombn.go` | n-digit combos |
| `descendcomb.go` | Descending combinations |
| `printnbrinorder.go` | BST in-order print |

### Mathematics
| File | Description |
|------|-------------|
| `fibonacci.go` | Fibonacci (recursion) |
| `iterativefactorial.go` / `recursivefactorial.go` | Factorial |
| `iterativepower.go` / `recursivepower.go` | Exponentiation |
| `isprime.go` / `findnextprime.go` | Prime operations |
| `sqrt.go` | Integer sqrt |
| `collatzcountdown.go` | Collatz steps |

### Cards & Games
| File | Description |
|------|-------------|
| `dealapackofcards.go` | Deal cards to 4 players |
| `shoppinglistsort.go` / `shoppingsummarycounter.go` | Shopping ops |
| `fooddeliverytime.go` | Food time lookup |
| `rockandroll.go` | Rock-paper-scissors |
| `podiumposition.go` | Medal placement |
| `loafofbread.go` | Bread simulation |
| `jumpover.go` | Array puzzle |
| `activebits.go` | Count set bits |

### Binary Search Trees (13 files)
`btreeinsertdata.go`, `btreedeletenode.go`, `btreetransplant.go`, `btreesearchitem.go`,
`btreemax.go`, `btreemin.go`, `btreelevelcount.go`,
`btreeapplyinorder.go`, `btreeapplypreorder.go`, `btreeapplypostorder.go`, `btreeapplybylevel.go`,
`btreeisbinary.go`, `printnbrinorder.go`

### Linked Lists (16 files)
`listpushback.go`, `listpushfront.go`, `listsize.go`, `listlast.go`, `listat.go`,
`listclear.go`, `listfind.go`, `listreverse.go`, `listmerge.go`, `listsort.go`,
`listforeach.go`, `listforeachif.go`, `listremoveif.go`,
`sortedlistmerge.go`, `sortlistinsert.go`

### Pointers & Memory
| File | Description |
|------|-------------|
| `enigma.go` | Triple pointer manipulation |
| `swap.go` | Pointer swap |
| `divmod.go` / `ultimatedivmod.go` | Divmod with/without pointers |
| `pointone.go` / `ultimatepointone.go` | Float precision |

## Command-Line Programs

| Directory | Description |
|-----------|-------------|
| `cat/` | Unix cat clone |
| `flags/` | CLI flag parser |
| `ztail/` | tail -c clone |
| `boolean/` | Even/odd arg check |
| `comcheck/` | Argument trigger check |
| `displayfile/` | File displayer |
| `revparams/` | Reverse args |
| `sortparams/` | Sort args |
| `printalphabet/` | a-z printer |
| `printdigits/` | 0-9 printer |
| `printreversealphabet/` | z-a printer |
| `printparams/` | Arg printer |
| `printprogramname/` | Name printer |
| `nbrconvertalpha/` | Number to alpha |
| `pilot/` | Struct demo |
| `point/` | Pointer demo |
| `rotatevowels/` | Vowel rotator |
| `fixthemain/` | Main transformer |

## Shell Scripts

`hello.sh`, `countfiles.sh`, `explain.sh`, `lookagain.sh`, `look`, `myfamily.sh`,
`mystery.sh`, `skip.sh`, `teacher.sh`, `to-git-or-not-to-git.sh`, `who-are-you.sh`,
`my_answer.sh`

## Project Structure

```
piscine-go/
├── README.md               ← Documentation
├── LICENSE                 ← MIT License
├── .gitignore              
├── go.mod / go.sum         ← Go 1.22.2 + z01
├── done.tar                ← Submission archive
├── *.go                    ← 100+ library exercises
├── boolean/                ← CLI: even/odd check
├── cat/                    ← CLI: cat clone
├── flags/                  ← CLI: flag parser
├── ztail/                  ← CLI: tail -c clone
├── comcheck/               ← CLI: arg checker
├── displayfile/            ← CLI: file reader
├── printalphabet/          ← CLI: a-z
├── printdigits/            ← CLI: 0-9
├── printreversealphabet/   ← CLI: z-a
├── printparams/            ← CLI: arg printer
├── revparams/              ← CLI: reverse args
├── sortparams/             ← CLI: sort args
├── nbrconvertalpha/        ← CLI: num to alpha
├── pilot/                  ← Struct demo
├── point/                  ← Pointer demo
├── rotatevowels/           ← CLI: vowel rotator
├── fixthemain/             ← Code transformer
├── quadchecker/            ← 🏆 Capstone project
└── 1/, 2/, ... 9/          ← Placeholder dirs
```

## Learning Path

### Level 1 — Foundations (Week 1)
String/char functions → `atoi`/`printnbr` → type checks → basic printing programs

### Level 2 — Control Flow (Week 1-2)
Combinations → prime/coll → factorial/power → sorting basic types

### Level 3 — Pointers & Structs (Week 2-3)
`swap` → `divmod` → `enigma` → pointer-based structs

### Level 4 — Linked Lists (Week 3)
Push/pop → find/remove → traverse → reverse → sort → merge

### Level 5 — Binary Trees (Week 3-4)
Insert → search → min/max → all traversals → delete → validate

### Level 6 — Algorithm Projects
Fibonacci → 8-Queens → card dealing → CLI tools → quadchecker capstone

## Key Concepts Covered

- **String manipulation**: strlen, strrev, split, join, capitalize, rot14
- **Type conversion**: atoi, atoi-base, base conversion, trimatoi
- **Type checking**: isalpha, isnumeric, isprintable, isprime
- **Sorting**: bubble sort, insertion sort, list merge sort
- **Data structures**: singly linked lists, binary search trees, hash maps
- **Recursion**: factorial, power, fibonacci, 8-queens
- **Pointers**: swap, enigma, divmod, point operations
- **CLI tools**: flag parsing, argument manipulation, file I/O
- **Shell scripting**: 12 automation scripts
- **Algorithm Design**: combinations, prime finding, pattern recognition

## Quick Start

```bash
# Clone
git clone https://github.com/dinalegw/piscine-go.git
cd piscine-go

# Run CLI program
cd cat && go run main.go README.md

# Use as library
import "piscine"
piscine.Atoi("42")        # 42
piscine.Fibonacci(10)     # 55
```

## License

MIT License — see [LICENSE](LICENSE) file for details.

```
MIT License
Copyright (c) 2026 dinalegw

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND.
```

---
<div align="center">Built with 💚 during the Learn2Earn Piscine Go Bootcamp</div>
