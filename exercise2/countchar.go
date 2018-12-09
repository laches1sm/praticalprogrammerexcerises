package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`What is the input string?`)
	word, err := reader.ReadString('\n')
	if err != nil{
		panic(err)
	}
	word = strings.TrimSpace(word)
	fmt.Println(countChar(word))


}

func countChar(word string)string{
	if word == ``{
	   return `Sorry, you cannot count the characters of an empty string`
   }else{
	   length := len(word)
	   output := fmt.Sprintf(`%s has %d characters`, word, length)
	   return output
   }




}
