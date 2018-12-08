package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "math"
    "strings"
)


func main(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Println(`What was the tip percentage?`)
    tipStr, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    tipAmount, err1 := strconv.ParseFloat(strings.TrimSpace(tipStr),64)
    if err1 != nil {
        panic(err)
     }
    fmt.Println(`What was the base price?`)
    basePriceStr, err2 := reader.ReadString('\n')
    if err2 != nil{
        panic(err)
    }
    basePrice, err3 := strconv.ParseFloat(strings.TrimSpace(basePriceStr), 64)
    if err3 != nil{
        panic(err)
    }
    tipCalcuator(tipAmount, basePrice)
}

func tipCalcuator(tipAmount float64, basePrice float64)(total float64, totalTip float64){

    tip := basePrice * (tipAmount/100)
    totalTip = math.Round(tip)
    total = totalTip + basePrice
    fmt.Printf(`Total tip is %f `, totalTip)
    fmt.Printf(`Total price is %f `, total)
    return total, totalTip


}
