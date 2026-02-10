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
