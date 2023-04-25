package main

import "time"

func main() {

}

func mySleep(vrem time.Duration) {
	<-time.After(vrem)
}
