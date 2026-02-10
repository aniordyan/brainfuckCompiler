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
