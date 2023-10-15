// Copyright © 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

package main

import (
	"github.com/reiver/go-telnet"
)

func main() {
	telnet.DialToAndCall("127.0.0.1:11211", telnet.StandardCaller)
}
