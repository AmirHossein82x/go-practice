package main

import (
	first "server/Servers"
	// two "server/ServerWithMultiplexer"
)

func main() {
	// go Servers.FirstServer()
	// Servers.SecondServer()
	// two.CreateFirstServerWithMux()
	// two.CreateFirstServerWithCustomHandler()
	// two.CreateFirstServerWithMux2()

	first.ThirdServer()

}
