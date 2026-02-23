package main

import
(
 "fmt"
 "os"
 "os/exec"
 "strings"
)

func main (){

  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "do: bf_compiler <file>.bf\n")
    os.Exit(1)
  }

 inputFile := os.Args[1]

 if !strings.HasSuffix(inputFile, ".bf") {
 		fmt.Fprintf(os.Stderr, "error: file must end in .bf\n")
 		os.Exit(1)
 	}

 name := strings.TrimSuffix(inputFile, ".bf")
 assembly := name + ".s"
 obj := name + ".o"
 exe := name

 // Read input file
    data, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: cannot read file: %v\n", err)
        os.Exit(1)
    }


    //scanner
    tokens := scanner(string(data))

    //parser
    if err := parser(tokens); err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

    //generate
    if err := generateAssembly(tokens, assembly); err != nil {
		fmt.Fprintf(os.Stderr, "codegen error: %v\n", err)
		os.Exit(1)
	}

    //linking
    if err := executing(assembly, obj, exe); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}


func scanner (input string) []byte{
 validInput := "><+-.,[]"
 tokens := []byte{}

 for i:= 0; i < len(input); i++ {
   char := input[i]
   if strings.ContainsRune(validInput, rune(char)) {
	   tokens = append (tokens, char)
   }
 }

 return tokens

}

func parser (tokens []byte) error{
 //tried to implement stack logic
 depth := 0

 for i, tok := range tokens {
  if tok == '[' {
    depth++
  } else if tok == ']' {
    if depth == 0 {
	return fmt.Errorf("unmatched ']' at position %d", i)
    }

    depth--
  }
 }

if depth != 0 {
	return fmt.Errorf("%d '[' — missing ']'", depth)
}

return nil

}


func generateAssembly(tokens []byte, name string) error {
//create .s file

  file, err := os.Create(name)
  if err != nil {
    return err
  }

  defer file.Close()

  depth := 0
  stack := []int{}
//headers
  fmt.Fprintf(file, ".section .bss\n")
  fmt.Fprintf(file, "tape: .skip 30000\n") //bf's memory size
  fmt.Fprintf(file, ".section .text\n")
  fmt.Fprintf(file, ".global _start\n")
  fmt.Fprintf(file, "_start:\n")
  fmt.Fprintf(file, "leaq tape(%%rip), %%r12\n")




  for _, tok := range tokens{
	switch tok {
		case '>':
			fmt.Fprintf(file, "incq %%r12\n")
		case '<':
      fmt.Fprintf(file, "decq %%r12\n")
		case '+':
      fmt.Fprintf(file, "incb (%%r12)\n")
		case '.':
                        fmt.Fprintf(file, "movq $1, %%rax\n")
			fmt.Fprintf(file, "movq $1, %%rdi\n")
			fmt.Fprintf(file, "movq %%r12, %%rsi\n")
			fmt.Fprintf(file, "movq $1, %%rdx\n")
			fmt.Fprintf(file, "syscall\n")
		case '-':
                        fmt.Fprintf(file, "decb (%%r12)\n")
		case ',':
                        fmt.Fprintf(file, "movq $0, %%rax\n")
                        fmt.Fprintf(file, "movq $0, %%rdi\n")
                        fmt.Fprintf(file, "movq %%r12, %%rsi\n")
                        fmt.Fprintf(file, "movq $1, %%rdx\n")
                        fmt.Fprintf(file, "syscall\n")
		case '[':
			current := depth
			depth++
			stack = append(stack, current)

			fmt.Fprintf(file, "loop_open_%d:\n", current) //to label each nestedloop
			fmt.Fprintf(file, "cmpb $0, (%%r12)\n")
			fmt.Fprintf(file, "je loop_close_%d\n", current)

		case ']':
			//pop from stack
			current := stack[len(stack)-1]
			stack = stack [:len(stack) - 1]

      fmt.Fprintf(file, "cmpb $0, (%%r12)\n")
			fmt.Fprintf(file, "jne loop_open_%d\n", current)
			fmt.Fprintf(file, "loop_close_%d:\n", current)



	}
  }

  //footer
  fmt.Fprintf(file, "movq $60, %%rax\n")
  fmt.Fprintf(file, "xorq %%rdi, %%rdi\n")
  fmt.Fprintf(file, "syscall\n")

  return nil
}


func executing (assembly, obj, exe string) error {
	cmd := exec.Command ("as", assembly, "-o", obj)

	cmd.Stderr = os.Stderr
        if err := cmd.Run(); err != nil {
        return fmt.Errorf("assembler failed: %v", err)
        }

	cmd = exec.Command ("ld", obj, "-o", exe)
	cmd.Stderr = os.Stderr
        if err := cmd.Run(); err != nil {
        return fmt.Errorf("linker failed: %v", err)
        }

    return nil

}
