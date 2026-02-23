# Brainfuck Compiler

Brainfuck is an esoteric programming language created in 1993 by Urban Müller. It is designed to be extremely minimalist, using only 8 simple commands. Despite its simplicity, Brainfuck is Turing complete, meaning it can theoretically compute anything that any other programming language can compute.

The language operates on an array of memory cells (called the "tape"), each initially set to zero. A pointer (called the "data pointer") begins at the first memory cell. The programmer can move this pointer left and right along the tape and modify the values in the cells.

## Brainfuck Commands

Brainfuck has only 8 commands:

| Command | Description |
|---------|-------------|
| `>` | Move the data pointer one cell to the right |
| `<` | Move the data pointer one cell to the left |
| `+` | Increment the value at the current cell by 1 |
| `-` | Decrement the value at the current cell by 1 |
| `.` | Output the value at the current cell as an ASCII character |
| `,` | Read one byte of input and store it in the current cell |
| `[` | If the current cell value is zero, jump forward to the matching `]` |
| `]` | If the current cell value is not zero, jump back to the matching `[` |

Any other characters in the source code are ignored as comments.

## How Brainfuck Works

The language uses a tape of 30,000 memory cells (in this implementation), each storing a value from 0 to 255. Here is a simple example:

```
+++       Set cell 0 to 3
>         Move to cell 1
++        Set cell 1 to 2
<         Move back to cell 0
.         Print cell 0 (prints ASCII character 3)
```

Loops are created using square brackets. The loop continues as long as the current cell is not zero:

```
+++++[>++<-]    Multiply: 5 * 2 = 10
                Cell 0 starts at 5
                Loop: add 2 to cell 1, subtract 1 from cell 0
                When cell 0 reaches 0, loop ends
                Result: cell 1 = 10
```

## What This Program Does

This program is a Brainfuck compiler written in Go. It takes Brainfuck source code as input and produces a native executable for x86_64 Linux systems.

The compiler performs four main stages:

### Stage 1: Scanning (Lexical Analysis)
The scanner reads the input file byte by byte and extracts only the valid Brainfuck commands (><+-.,[]). All other characters, including whitespace and comments, are ignored.

### Stage 2: Parsing (Syntax Validation)
The parser validates the structure of the program. It ensures that:
- Every opening bracket `[` has a matching closing bracket `]`
- Brackets are properly nested
- If any bracket is unmatched, the compiler stops and reports an error

### Stage 3: Code Generation
The code generator translates each Brainfuck command into x86_64 assembly language instructions. It generates a `.s` assembly file with the same base name as the input file.

The generated assembly:
- Creates a 30,000-byte memory tape in the `.bss` section
- Uses register `%r12` as the data pointer
- Implements each BF command with appropriate assembly instructions
- Uses Linux system calls for input/output operations
- Generates unique labels for nested loops

### Stage 4: Assembly and Linking
The compiler automatically invokes the system assembler (`as`) and linker (`ld`) to produce a standalone executable. The user does not need to run these tools manually.

## Requirements

- Go programming language (version 1.16 or later)
- Linux x86_64 system
- GNU assembler (`as`)
- GNU linker (`ld`)

## Building the Compiler

```bash
go build -o bf_compiler main.go
```

This creates an executable named `bf_compiler`.

## Usage

```bash
./bf_compiler <filename>.bf
```

The compiler expects a Brainfuck source file with a `.bf` extension. It will generate three files:

- `<filename>.s` - The generated assembly code
- `<filename>.o` - The compiled object file
- `<filename>` - The final executable

### Example

Create a simple Brainfuck program:

```bash
echo '>+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.>>>++++++++[<++++>-]
<.>>>++++++++++[<+++++++++>-]<---.<<<<.+++.------.--------.>>+.>++++++++++.
' > hello.bf
```

Compile it:

```bash
./bf_compiler hello.bf
```

Run the generated executable:

```bash
./hello
```

This will print 'Hello World!'

## Error Handling

The compiler will report errors in the following cases:

- No input file specified
- Input file does not have a `.bf` extension
- Input file cannot be read
- Unmatched brackets in the source code
- Assembly or linking errors

## Examples

### Example 1: Print a Single Character

```brainfuck
+++.
```

This program increments cell 0 three times and prints it (ASCII value 3).

### Example 2: Hello World

```brainfuck
++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.
```

This is a classic Hello World program in Brainfuck.
