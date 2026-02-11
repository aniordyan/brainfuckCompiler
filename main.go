package main

import
(
 "fmt"
 "os"
 "os/exec"
 "strings"
)

func main ()
{

}


func scanner (input string) []int
{
 validInput := "><+-.,[]"
 tokens := []int{}

 for i:= 0; i < len(input); i++
 {
   char := input[i]
   if strings.ContainRune(validInput, rune(char))
   {
	   tokens = append (tokens, char)
   }
 }

 return tokens
 
}

func parser (tokens []int) error 
{
 //tried to implement stack logic
 depth := 0

 for i, tok := range tokens
 {
  if tok == '['
  {
    depth++
  } 

  else if tok == ']'
  {
    if depth == 0 
    {
	return fmt.Errorf("unmatched ']' at position %d", i)
    }

    depth--
  }
 }

if depth != 0 
{
	return fmt.Errorf("%d '[' — missing ']'", depth)
}

return nil

}


func generateAssembly(tokens []int, name string) error

{
//create .s file

  file, err := os.Create(name)
//headers

  for_, tok := range tokens 
  {
	switch tok
	{
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
                        fmt.Fprintf(file, "incq %%r12\n")
		case ']':
                        fmt.Fprintf(file, "incq %%r12\n")	
	}
  }
}
