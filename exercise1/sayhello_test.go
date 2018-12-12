package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func testSayHelloHappyPath(t *testing.T) {
	Convey(`When I input a name into the program`, t, func() {
		name := `Rebecca`
		realInput := sayHello(name)
		Convey(`I should expect the message to be "Hello Rebecca, nice to meet you!"`, func() {
			So(realInput, ShouldEqual, `Hello Rebecca, nice to meet you!`)
		})
	})
}

func testSayHelloEmptyString(t *testing.T) {
	Convey(`When I input an empty string into the program`, t, func() {
		name := ``
		realInput := sayHello(name)
		Convey(`I should expect the message to be "Sorry, cannot accept an empty string"`, func() {
			So(realInput, ShouldEqual, `Sorry, cannot accept an empty string`)
		})
	})
}
