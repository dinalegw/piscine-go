# QuadChecker

## 📌 Project Overview

QuadChecker is a pattern-recognition system that identifies ASCII art rectangles and matches them against five known quad patterns. Given an ASCII rectangle as input, it determines which quad(s) could have created it and reports their dimensions.

### The Problem
- **Input**: ASCII art pattern
- **Output**: Matching quad(s) with dimensions, or "Not a quad function"
- **Challenge**: Handle ambiguous cases where multiple quads match the same pattern

### Real-World Analogy
Like identifying a car manufacturer from visual design:
- **Quad A**: Signature style with `o`, `-`, `|`
- **Quad B**: Diagonal style with `/`, `\`, `*`
- **Quads C/D/E**: Combinations using `A`, `B`, `C`

---

## 🏗️ Architecture

QuadChecker follows a **pipeline architecture**:

```
INPUT → VALIDATE → MATCH → FORMAT → OUTPUT
```

### Components

**1. Input Reader** - Reads ASCII lines from stdin
**2. Validator** - Ensures input is a valid rectangle
**3. Pattern Matcher** - Compares against 5 quad templates
**4. Formatter** - Sorts matches and formats output

---

## 🧱 Core Data Structures

### Quad Specification

```go
type Quad struct {
  Name     string  // "quadA", "quadB", etc.
  TopLeft  rune    // Corner character
  TopRight rune    // Corner character
  BotLeft  rune    // Corner character
  BotRight rune    // Corner character
  Horiz    rune    // Horizontal border
  Vert     rune    // Vertical border
}
```

### Pattern Database

```go
var quads = []Quad{
  {"quadA", 'o', 'o', 'o', 'o', '-', '|'},
  {"quadB", '/', '\\', '\\', '/', '*', '*'},
  {"quadC", 'A', 'A', 'C', 'C', 'B', 'B'},
  {"quadD", 'A', 'C', 'A', 'C', 'B', 'B'},
  {"quadE", 'A', 'C', 'C', 'A', 'B', 'B'},
}
```

### Quad Fingerprints

| Quad | Corners | Feature |
|------|---------|---------|
| quadA | o,o,o,o | All identical |
| quadB | /,\,\,/ | Diagonal |
| quadC | A,A,C,C | Top vs bottom differ |
| quadD | A,C,A,C | Left vs right differ |
| quadE | A,C,C,A | Cross-diagonal |

---

## 📥 Input Processing

### Step 1: Read Input

```go
func readInput() ([]string, error) {
  scanner := bufio.NewScanner(os.Stdin)
  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}
```

### Step 2: Validate Rectangle

```go
func isValidRectangle(lines []string) bool {
  if len(lines) == 0 {
    return false
  }
  width := len(lines[0])
  if width == 0 {
    return false
  }
  for _, line := range lines {
    if len(line) != width {
      return false
    }
  }
  return true
}
```

---

## 🔍 Pattern Matching Algorithm

### Core Logic

```go
func checkQuad(lines []string, quad Quad) bool {
  height := len(lines)
  width := len(lines[0])
  
  // Special case: 1×1
  if width == 1 && height == 1 {
    return lines[0][0] == byte(quad.TopLeft)
  }
  
  // Check 4 corners
  if rune(lines[0][0]) != quad.TopLeft { return false }
  if rune(lines[0][width-1]) != quad.TopRight { return false }
  if rune(lines[height-1][0]) != quad.BotLeft { return false }
  if rune(lines[height-1][width-1]) != quad.BotRight { return false }
  
  // Check horizontal borders
  for col := 1; col < width-1; col++ {
    if rune(lines[0][col]) != quad.Horiz { return false }
    if rune(lines[height-1][col]) != quad.Horiz { return false }
  }
  
  // Check vertical borders
  for row := 1; row < height-1; row++ {
    if rune(lines[row][0]) != quad.Vert { return false }
    if rune(lines[row][width-1]) != quad.Vert { return false }
  }
  
  // Check interior is all spaces
  for row := 1; row < height-1; row++ {
    for col := 1; col < width-1; col++ {
      if lines[row][col] != ' ' { return false }
    }
  }
  
  return true
}
```

### Matching Strategy

1. **1×1 check**: Special case—only check TopLeft corner
2. **Corners**: 4 comparisons to eliminate non-matching quads
3. **Borders**: Check horizontal and vertical patterns
4. **Interior**: Verify all interior cells are spaces
5. **Fail-fast**: Return false at first mismatch

---

## 🧠 Finding All Matches

```go
func findMatchingQuads(lines []string) []string {
  var matches []string
  height := len(lines)
  width := len(lines[0])
  
  for _, quad := range quads {
    if checkQuad(lines, quad) {
      matches = append(matches,
        fmt.Sprintf("[%s] [%d] [%d]", quad.Name, width, height))
    }
  }
  
  return matches
}
```

**Why check all quads?** Some patterns match multiple quads (e.g., 1×1 "A" matches quadC, D, and E).

---

## 📤 Output Processing

```go
func main() {
  lines, err := readInput()
  if err != nil || len(lines) == 0 {
    fmt.Println("Not a quad function")
    return
  }
  
  if !isValidRectangle(lines) {
    fmt.Println("Not a quad function")
    return
  }
  
  matches := findMatchingQuads(lines)
  
  if len(matches) == 0 {
    fmt.Println("Not a quad function")
  } else {
    sort.Strings(matches)
    result := strings.Join(matches, " || ")
    fmt.Println(result)
  }
}
```

### Output Rules

- Single match: `[quadX] [width] [height]\n`
- Multiple matches: Sort alphabetically, join with ` || `
- No match: `Not a quad function\n`

---

## ⚠️ Special Cases

### 1×1 Edge Case

Single character rectangles have all corners at one position:

```
Pattern: "A"
Result:  [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

Only TopLeft corner matters for 1×1 matching.

### Empty or Invalid Input

Rejected at validation stage:
- Empty input
- Lines of different lengths
- Zero-width lines

### Interior Contamination

Interior must contain only spaces:

```
Valid:              Invalid:
o---o               o---o
|   |               |XXX|
o---o               o---o
```

---

## ⏱️ Complexity Analysis

### Time Complexity

- **readInput()**: O(N) – reads all characters
- **isValidRectangle()**: O(L) – checks each line
- **checkQuad()**: O(W × H) – examines every cell
- **findMatchingQuads()**: 5 × O(W × H) = O(N) – 5 quads total

**Total: O(N)** where N = total input characters

### Space Complexity

- **Storage**: O(N) – stores all input lines
- **Processing**: O(1) – loop counters only
- **Output**: O(M) where M ≤ 5 matches

---

## 🧪 Key Test Cases

```
Input:              Expected Output:
o---o               [quadA] [5] [3]
|   |
o---o

A                   [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]

invalid             Not a quad function

o---o               Not a quad function
| |
o---o
```

---

## 🚨 Common Pitfalls

### Off-by-One Errors

❌ **Wrong**: `for col := 0; col < width; col++` (includes corners)
✅ **Right**: `for col := 1; col < width-1; col++` (excludes corners)

### Missing Interior Check

❌ **Wrong**: Only checking borders
✅ **Right**: Verify all interior cells are spaces

### Not Handling 1×1 Separately

❌ **Wrong**: Treating 1×1 like normal rectangles
✅ **Right**: Special case for single character

### Not Sorting Multiple Matches

❌ **Wrong**: `[quadE] || [quadC] || [quadD]`
✅ **Right**: `[quadC] || [quadD] || [quadE]`

### Missing Newline

❌ **Wrong**: `fmt.Print("result")` 
✅ **Right**: `fmt.Println("result")`

---

## 📊 Complete Execution Flow

```
START
  ↓
readInput() ──→ Error/Empty? ──→ Output: "Not a quad function" → EXIT
  ↓ Success
isValidRectangle() ──→ Invalid? ──→ Output: "Not a quad function" → EXIT
  ↓ Valid
findMatchingQuads()
  ├─ for each quad:
  │  └─ checkQuad()
  │     ├─ Check 1×1?
  │     ├─ Check corners
  │     ├─ Check borders
  │     ├─ Check interior
  │     └─ Return true/false
  └─ Collect matches
  ↓
len(matches) == 0? ──→ Yes ──→ Output: "Not a quad function" → EXIT
  ↓ No
sort.Strings(matches)
strings.Join(matches, " || ")
fmt.Println(result)
  ↓
EXIT
```

---


## 🎯 Key Learning Points

### Design Principles
- **Separation of Concerns**: Each function has one responsibility
- **Fail-Fast**: Reject invalid input early
- **Pipeline Architecture**: Each stage transforms data
- **Template Matching**: Compare data against known patterns

### Go Patterns
- **Error Handling**: Explicit error returns
- **Slice Management**: Dynamic arrays with append()
- **Rune vs Byte**: Character type awareness
- **String Indexing**: Returns bytes, not characters

### Algorithm Concepts
- **Pattern Recognition**: Breaking complex patterns into components
- **Validation**: Multi-layer defensive checking
- **Multiple Solutions**: Some problems have multiple correct answers
- **Ambiguity Handling**: Report all valid possibilities

---

# Line-by-Line Analysis of QuadChecker

## 📦 Package Imports Analysis

### 1. `bufio` - Buffered I/O
**Purpose:** Efficient reading of input stream

```go
scanner := bufio.NewScanner(os.Stdin)
// Creates a scanner that reads from standard input in chunks (4096 bytes buffer)
// More efficient than reading byte-by-byte
```

### 2. `fmt` - Formatting and I/O
**Purpose:** Formatted output and string formatting

```go
fmt.Println("Not a quad function")      // Lines 129, 134, 142
fmt.Sprintf("[%s] [%d] [%d]", ...)      // Line 109
// Provides Print, Printf, Println, Sprintf for formatted output
```

### 3. `os` - Operating System Interface
**Purpose:** Access to stdin and other OS features

```go
os.Stdin  // Line 38
// Represents the standard input stream (file descriptor 0)
```

### 4. `sort` - Sorting Algorithms
**Purpose:** Sorting slices and collections

```go
sort.Strings(matches)  // Line 144
// Sorts string slice in-place using quicksort (O(n log n))
```

### 5. `strings` - String Manipulation
**Purpose:** String operations and manipulation

```go
strings.Join(matches, " || ")  // Line 146
// Joins string slice with separator, more efficient than manual concatenation
```

---

## 🧬 Line-by-Line Code Analysis

### Lines 1-5: Package Declaration and Imports

```go
package main  // Line 1: Declares this as an executable program (not a library)

import (  // Line 2: Start import block
  "bufio"   // Line 3: Buffered I/O for efficient input reading
  "fmt"     // Line 4: Formatted I/O operations
  "os"      // Line 5: OS interface for stdin access
  "sort"    // Line 6: Sorting algorithms
  "strings" // Line 7: String manipulation utilities
)  // Line 8: End import block
```

### Lines 10-17: Quad Structure Definition

```go
type Quad struct {  // Line 10: Defines a new struct type named 'Quad'
  Name     string  // Line 11: Name of the quad pattern (e.g., "quadA")
  TopLeft  rune    // Line 12: Character at top-left corner position
  TopRight rune    // Line 13: Character at top-right corner position  
  BotLeft  rune    // Line 14: Character at bottom-left corner position
  BotRight rune    // Line 15: Character at bottom-right corner position
  Horiz    rune    // Line 16: Character for horizontal borders
  Vert     rune    // Line 17: Character for vertical borders
}  // Line 18: End of struct definition
```

**Why `rune` instead of `byte`?**
- `rune` represents a Unicode code point (int32)
- `byte` is unsigned 8-bit integer (uint8)
- Using `rune` is more general, though quad patterns use only ASCII
- Clear semantic meaning: "this is a character"

### Lines 20-26: Quad Pattern Database

```go
var quads = []Quad{  // Line 20: Declares and initializes slice of Quad structs
  {"quadA", 'o', 'o', 'o', 'o', '-', '|'},      // Line 21: Pattern A
  {"quadB", '/', '\\', '\\', '/', '*', '*'},    // Line 22: Pattern B (note escaped backslash)
  {"quadC", 'A', 'A', 'C', 'C', 'B', 'B'},      // Line 23: Pattern C
  {"quadD", 'A', 'C', 'A', 'C', 'B', 'B'},      // Line 24: Pattern D  
  {"quadE", 'A', 'C', 'C', 'A', 'B', 'B'},      // Line 25: Pattern E
}  // Line 26: End of slice initialization
```

**Note on Line 22:** `'\\'` is an escaped backslash. In Go string/char literals:
- `\` is the escape character
- `\\` represents a literal backslash character
- This matches quadB's pattern which uses `\` character

### Lines 28-41: `readInput()` Function

```go
func readInput() ([]string, error) {  // Line 28: Function returns (slice of strings, error)
  scanner := bufio.NewScanner(os.Stdin)  // Line 29: Create scanner reading from stdin
  
  var lines []string  // Line 30: Declare empty string slice (nil, length 0, capacity 0)
  
  for scanner.Scan() {  // Line 31: Loop while scanner can read more input
    lines = append(lines, scanner.Text())  // Line 32: Add line to slice
    // scanner.Text() returns current line WITHOUT newline character
    // append() grows slice as needed (amortized O(1) cost)
  }
  
  if err := scanner.Err(); err != nil {  // Line 34: Check for scanner errors
    return nil, err  // Line 35: Return nil slice and error
  }
  
  return lines, nil  // Line 37: Return slice and nil error (success)
}  // Line 38: End function
```

**Memory Behavior:**
- `lines` starts as nil slice (zero value)
- Each `append()` may cause reallocation with exponential growth
- Typical growth: 0 → 1 → 2 → 4 → 8 → 16 ... capacity

### Lines 43-66: `isValidRectangle()` Function

```go
func isValidRectangle(lines []string) bool {  // Line 43: Function takes string slice, returns bool
  if len(lines) == 0 {  // Line 44: Check if slice is empty
    return false  // Line 45: Empty input is not a rectangle
  }
  
  width := len(lines[0])  // Line 47: Get length of first line
  if width == 0 {  // Line 48: Check if first line is empty
    return false  // Line 49: Zero-width rectangle not valid
  }
  
  for _, line := range lines {  // Line 52: Iterate over each line in slice
    if len(line) != width {  // Line 53: Compare current line length to width
      return false  // Line 54: Early return if lengths differ
    }
    
    for _, ch := range line {  // Line 56: Iterate over each rune in the line
      if ch < 32 || ch > 126 {  // Line 57: Check if outside printable ASCII range
        if ch != ' ' && ch != '\t' {  // Line 58: Allow space and tab
          return false  // Line 59: Reject other non-printable chars
        }
      }
    }
  }
  
  return true  // Line 63: All checks passed, it's a valid rectangle
}  // Line 64: End function
```

**ASCII Range Check Details:**
- 0-31: Control characters (tab=9, newline=10, carriage return=13)
- 32: Space (allowed)
- 33-126: Printable ASCII (allowed)
- 127+: Extended ASCII/Unicode (would be rejected by pattern matching anyway)

### Lines 68-112: `checkQuad()` Function

```go
func checkQuad(lines []string, quad Quad) bool {  // Line 68: Pattern matching function
  height := len(lines)      // Line 69: Number of rows (lines)
  width := len(lines[0])    // Line 70: Length of first line (all same length)
  
  // Special case: 1x1 rectangle
  if width == 1 && height == 1 {  // Line 73: Check for 1×1 case
    return lines[0][0] == byte(quad.TopLeft)  // Line 74: Compare single character
    // For 1×1, all corners are the same position, so only check TopLeft
  }
  
  // Check all four corners
  if rune(lines[0][0]) != quad.TopLeft {  // Line 77: Top-left corner check
    return false  // Line 78: Early return on mismatch
  }
  if rune(lines[0][width-1]) != quad.TopRight {  // Line 79: Top-right corner
    return false  // Line 80
  }
  if rune(lines[height-1][0]) != quad.BotLeft {  // Line 81: Bottom-left corner
    return false  // Line 82
  }
  if rune(lines[height-1][width-1]) != quad.BotRight {  // Line 83: Bottom-right corner
    return false  // Line 84
  }
  
  // Check top and bottom borders (excluding corners)
  for col := 1; col < width-1; col++ {  // Line 87: Loop from second to second-last column
    if rune(lines[0][col]) != quad.Horiz || rune(lines[height-1][col]) != quad.Horiz {  // Line 88
      return false  // Line 89
    }
  }
  
  // Check left and right borders (excluding corners)
  for row := 1; row < height-1; row++ {  // Line 92: Loop from second to second-last row
    if rune(lines[row][0]) != quad.Vert || rune(lines[row][width-1]) != quad.Vert {  // Line 93
      return false  // Line 94
    }
  }
  
  // Check interior is only spaces
  for row := 1; row < height-1; row++ {  // Line 97: Interior rows only
    for col := 1; col < width-1; col++ {  // Line 98: Interior columns only
      if lines[row][col] != ' ' {  // Line 99: Compare to space character
        return false  // Line 100
      }
    }
  }
  
  return true  // Line 103: All checks passed - pattern matches this quad!
}  // Line 104: End function
```

**Loop Boundary Details:**
- Top border: `lines[0][1:width-1]` (exclude corners at `[0][0]` and `[0][width-1]`)
- Bottom border: `lines[height-1][1:width-1]`
- Left border: `lines[1:height-1][0]` (each row, first column)
- Right border: `lines[1:height-1][width-1]` (each row, last column)
- Interior: `lines[1:height-1][1:width-1]` (exclude all borders)

### Lines 106-115: `findMatchingQuads()` Function

```go
func findMatchingQuads(lines []string) []string {  // Line 106: Returns slice of match strings
  var matches []string  // Line 107: Declare empty string slice
  height := len(lines)  // Line 108: Get dimensions once (efficiency)
  width := len(lines[0])  // Line 109
  
  for _, quad := range quads {  // Line 110: Iterate over all 5 quad patterns
    if checkQuad(lines, quad) {  // Line 111: Check if pattern matches this quad
      matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad.Name, width, height))  // Line 112
    }
  }
  
  return matches  // Line 114: Return all matches found (could be empty)
}  // Line 115: End function
```

**Format String Breakdown:**
- `%s`: quad.Name (string)
- `%d`: width (integer)
- `%d`: height (integer)
- Result: `"[quadA] [5] [3]"`

### Lines 117-148: `main()` Function

```go
func main() {  // Line 117: Program entry point
  lines, err := readInput()  // Line 119: Call readInput, get lines and error
  if err != nil || len(lines) == 0 {  // Line 120: Check for errors OR empty input
    fmt.Println("Not a quad function")  // Line 121: Output error message
    return  // Line 122: Exit function (and program) immediately
  }
  
  if !isValidRectangle(lines) {  // Line 125: Call validation function
    fmt.Println("Not a quad function")  // Line 126: Same error message
    return  // Line 127: Early exit
  }
  
  matches := findMatchingQuads(lines)  // Line 130: Get all possible matches
  
  if len(matches) == 0 {  // Line 133: Check if no matches found
    fmt.Println("Not a quad function")  // Line 134: Output error
  } else {  // Line 135: At least one match found
    sort.Strings(matches)  // Line 137: Sort in-place (modifies matches slice)
    result := strings.Join(matches, " || ")  // Line 139: Create single string
    fmt.Println(result)  // Line 140: Output final result
  }
}  // Line 142: End main function (program ends here)
```

**Important Output Details:**
- Single match: `"[quadA] [5] [3]\n"`
- Multiple matches: `"[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]\n"`
- Error case: `"Not a quad function\n"`
- Always ends with newline due to `fmt.Println()`

---

## 🔑 Key Concepts Explained

### 1. Slice vs Array

```go
var lines []string  // Slice (dynamic size, reference type)
var lines [100]string  // Array (fixed size, value type)
```

Slices are used because input size is unknown at compile time. Slices are 3-word structures: pointer, length, capacity. `append()` handles growth automatically.

### 2. Rune vs Byte Conversions

```go
if rune(lines[0][0]) != quad.TopLeft { ... }
if lines[0][0] != byte(quad.TopLeft) { ... }
```

Both work for ASCII characters. Using `rune()` is:
- More general (handles Unicode)
- Clearer intent ("this is a character comparison")
- Type-safe (explicit conversion)

### 3. Early Return Pattern

```go
if condition {
  return false  // Fail fast
}
// Continue only if condition passed
```

Improves efficiency (don't do unnecessary work). Simplifies logic (no deep nesting). Clearer flow (each check is independent).

### 4. Zero-Based Indexing

```go
lines[0][0]  // First line, first character (top-left corner)
lines[0][width-1]  // First line, last character (top-right corner)
lines[height-1][0]  // Last line, first character (bottom-left corner)
```

Go uses 0-based indexing (like C, Java, Python). Last element is at index `length-1`. Critical for avoiding off-by-one errors.

### 5. String Immutability

```go
line := lines[0]  // line is a string
// Strings are immutable in Go
// line[0] = 'X'  // This would be a compile error
```

Strings cannot be modified after creation. Operations return new strings. Slicing a string (`line[1:4]`) creates new string sharing underlying bytes.

### 6. Memory Layout Visualization

```
lines slice header: [ptr→] [len=3] [cap=3]
          ↓
        Array of string headers:
        [0]: [ptr→"o---o"] [len=5]
        [1]: [ptr→"|   |"] [len=5]  
        [2]: [ptr→"o---o"] [len=5]
          ↓
        Actual character data in memory
```

---

## 🎯 Algorithm Complexity Summary

### Time Complexity:

- **readInput()**: O(N) where N = total characters
- **isValidRectangle()**: O(N) (checks each character)
- **checkQuad()**: O(N) per quad (5× total)
- **findMatchingQuads()**: O(5N) = O(N)
- **sort.Strings()**: O(m log m) where m ≤ 5 (negligible)
- **Total**: O(N) - optimal (must examine every character)

### Space Complexity:

- **Input storage**: O(N) (must store to examine)
- **Quad patterns**: O(1) (5× small structs)
- **Matches**: O(1) (≤5 strings)
- **Total**: O(N) - optimal (must store input)

### Memory Allocation Points:

- `append(lines, scanner.Text())` - Each line allocation
- `fmt.Sprintf()` in matches - Each match string
- `strings.Join()` - Final output string
- Scanner buffer - Internal 4096-byte buffer

---

## 🚨 Critical Implementation Details

### 1. The 1×1 Special Case

```go
if width == 1 && height == 1 {
  return lines[0][0] == byte(quad.TopLeft)
}
```

**Why special?** All four corners occupy same position. **Logic:** Only check TopLeft character. **Result:** "A" matches quadC, D, E (all have 'A' as TopLeft).

### 2. Border Loop Boundaries

```go
for col := 1; col < width-1; col++ {
  // Checks columns 1 through width-2
  // Excludes 0 (left corner) and width-1 (right corner)
}
```

**Critical:** Must exclude corners from border checks. For width=2: Loop doesn't run (1 < 1 is false) - correct (2×2 has no borders). For width=1: Would cause issues, but 1×1 handled separately.

### 3. Interior Space Requirement

```go
if lines[row][col] != ' ' {
  return false
}
```

**Why?** Quad patterns define only borders, interior must be empty. Without this check: Would accept patterns with filled interiors. **Performance:** Most expensive check (O((w-2)×(h-2))).

### 4. Exact Output Formatting

```go
fmt.Sprintf("[%s] [%d] [%d]", quad.Name, width, height)
// NOT: "[%s][%d][%d]" (missing spaces)
// NOT: "quad%s: %dx%d" (wrong format)
```

Spaces matter: Automated tests check exact string. Brackets required: `[quadA]` not `quadA`. Newline required: `fmt.Println()` ensures this.

### 5. Multiple Match Handling

```go
sort.Strings(matches)  // Alphabetical order required
strings.Join(matches, " || ")  // Exact separator format
```

**Specification:** "display them all alphabetically ordered". **Separator:** Space, double pipe, space (`" || "`). Single match: No separator added by `strings.Join()`.

---

## 🔍 Debugging Insights

### Common Issues and Fixes

#### 1. Off-by-One in Border Loops

```go
// WRONG
for col := 0; col < width; col++ { ... }  // Includes corners!

// RIGHT  
for col := 1; col < width-1; col++ { ... }  // Excludes corners
```

#### 2. Forgetting 1×1 Special Case

```go
// WRONG: Tries to access lines[0][width-1] when width=1
// This would be out of bounds or compare wrong positions

// RIGHT: Handle 1×1 separately
if width == 1 && height == 1 { ... }
```

#### 3. Missing Interior Check

```go
// WRONG: Would accept:
// o---o
// |XXX|  ← X's should be rejected
// o---o

// RIGHT: Check interior spaces
for row := 1; row < height-1; row++ {
  for col := 1; col < width-1; col++ {
    if lines[row][col] != ' ' { return false }
  }
}
```

#### 4. Output Format Errors

```go
// WRONG: Missing spaces/brackets/newline
fmt.Print("quadA", width, height)

// RIGHT: Exact format
fmt.Println(fmt.Sprintf("[%s] [%d] [%d]", ...))
```

---

## 🎓 Learning Points from Each Section

**From Package Imports:**
- Choose libraries based on needs (I/O, strings, sorting)
- Go's standard library is comprehensive
- Import only what you need

**From Data Structures:**
- Structs organize related data
- Slices provide flexible collections
- Constants/global variables for configuration

**From Algorithm Design:**
- Break complex problems into functions
- Early validation saves computation
- Fail-fast improves efficiency
- Special cases need special handling

**From Implementation Details:**
- Boundary conditions are critical
- Type conversions matter
- Memory efficiency considerations
- Code clarity aids maintenance

**From Testing Considerations:**
- Edge cases reveal design flaws
- Exact output requirements
- Performance constraints
- Error handling completeness

This complete line-by-line analysis shows how QuadChecker embodies good software engineering practices: clear structure, efficient algorithms, robust error handling, and attention to detail. Every line serves a specific purpose in creating a reliable, correct pattern recognition system.


## 📚 Next Steps

To extend your learning:

1. **Modify the system**: Add more quad patterns
2. **Optimize further**: Profile and optimize hot paths
3. **Add features**: Configurable patterns from files
4. **Related projects**: File type detection, network protocols
5. **Contribute**: Practice with open source projects


## 🎓 Conclusion & Learning Resources

### 🎯 Key Takeaways
- **Complex problems decompose into manageable parts**
- **Edge cases matter** — 1×1 ambiguity teaches critical lessons
- **Specifications are contracts** — exact output formatting is crucial
- **Testing is essential** — unit to system level validation
- **Code clarity enables collaboration and maintenance**

### 📚 Essential Learning Resources

**Go Programming Language**
- [A Tour of Go](https://tour.golang.org) — Interactive tutorial
- [Go by Example](https://gobyexample.com) — Practical code examples
- [Effective Go](https://golang.org/doc/effective_go) — Style guide
- "The Go Programming Language" by Donovan & Kernighan

**Algorithms & Data Structures**
- [VisuAlgo](https://visualgo.net) — Visual learning
- [Big-O Cheat Sheet](https://www.bigocheatsheet.com) — Complexity reference
- "Grokking Algorithms" by Aditya Bhargava
- "Introduction to Algorithms" by Cormen et al.

**Testing & Quality**
- [Go Testing Package](https://golang.org/pkg/testing)
- Table-driven tests and fuzzing documentation
- [Testify Library](https://github.com/stretchr/testify)

**Software Architecture**
- [SOLID Principles in Go](https://dave.cheney.net/2016/08/20/solid-go-design)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Go Design Patterns](https://refactoring.guru/design-patterns/go)

**System Design & Problem Solving**
- "How to Solve It" by George Pólya
- [System Design Primer](https://github.com/donnemartin/system-design-primer)
- "Designing Data-Intensive Applications" by Martin Kleppmann

**Documentation & Communication**
- [Go Doc Comments](https://tip.golang.org/doc/comment)
- [Readme.so](https://readme.so) — Interactive README builder
- [Mermaid.js](https://mermaid.js.org) — Diagram generation

**Communities**
- [Gophers Slack](https://gophers.slack.com) — Community chat
- [r/golang](https://reddit.com/r/golang) — Reddit community
- [Go Blog](https://go.dev/blog) — Official updates

### 💡 Words of Wisdom
- **Learn by doing** — Build projects, not just tutorials
- **Read others' code** — Open source is your classroom
- **Embrace debugging** — Every bug teaches something
- **Write tests first** — TDD leads to better design
- **Document as you go** — Future you will be grateful
- **Seek feedback** — Code reviews accelerate learning
- **Stay curious** — The field evolves constantly

> "The only way to learn a new programming language is by writing programs in it." — Dennis Ritchie

Your QuadChecker skills transfer directly to pattern recognition, validation systems, CLI tools, and protocol parsing. Keep building! 🚀
