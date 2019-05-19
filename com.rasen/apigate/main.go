package main

import "com.rasen/apigate/gate"

func main() {
	gate.NewGateWay([]string{"127.0.0.1:2379"})
}
