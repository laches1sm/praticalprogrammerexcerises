package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`What is your name?`)
	name, err := reader.ReadString('\n')
	if err != nil{
		panic(err)
	}
	name = strings.TrimSpace(name)
	output := sayHello(name)
	fmt.Println(output)

}

func sayHello(name string) string{
	if name == ``{
		return `Sorry, cannnot accept an empty string`
	}else{
		output := fmt.Sprintf(`Hello %s, nice to meet you!`, name)
		return output
}
}
