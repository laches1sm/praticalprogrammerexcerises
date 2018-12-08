package main

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestTipCalc(*testing.T){
    Convey(`Given I have a bill of $30 and a tip rate of 15%`, t, func(){
        billAmount := 30
        tipRate := 15
        expectedTotal := 34.5
        expectedTipTotal := 4.5

        Convey(`When I pass these values through the calculator`, func(){
            tipAmount, total := TipCalculator(billAmount, tipRate)

            Convey(`Then I should receive a total amount and a tip amount`, func(){
                So(total, ShouldEqual, expectedTotal)
                So (tipAmount, ShouldEqual, expectedTipTotal)
            })
        })
    })

}



