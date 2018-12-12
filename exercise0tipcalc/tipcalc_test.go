package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTipCalc(t *testing.T) {
	Convey(`Given I have a bill of $30 and a tip rate of 15%`, t, func() {
		billAmount := 30.00
		tipRate := 15.00
		expectedTotal := 35
		expectedTipTotal := 5

		Convey(`When I pass these values through the calculator`, func() {
			total, tipAmount := tipCalcuator(tipRate, billAmount)

			Convey(`Then I should receive a total amount and a tip amount`, func() {
				So(total, ShouldEqual, expectedTotal)
				So(tipAmount, ShouldEqual, expectedTipTotal)
			})
		})
	})

}
