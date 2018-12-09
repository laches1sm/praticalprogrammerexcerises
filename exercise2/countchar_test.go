package main

import(
	"testing"
	. "github.com/smartystreets/goconvey"
)

func testCountChar(t *testing.T){
	Convey(`When I input a string into the program`, t, func(){
		strInput := `parrots`
		realOutput := countChar(strInput)
		Convey(`I should expect to see the number of characters returned`, func(){
			So(realOutput, ShouldEqual, `parrots has 7 characters`)
		})
	})
}

func testCountCharEmptyOutput(t *testing.T){
	Convey(`When I input an empty string into the program`, t, func(){
		strInput = ``
		realOutput := countChar(strInput)
		Convey(`I should expect to see an error message returned`, func(){
			So(realOutput, ShouldEqual, `Sorry, you cannot count the characters of an empty string`)
		})
	})
}
