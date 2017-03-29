package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)
	i := 0
	t0 := time.Now()
	for {
		if i == 0 {
			t0 = time.Now()
		}
		_, _, err := ServerConn.ReadFromUDP(buf)
		t1 := time.Now()
		fmt.Printf("Program start time:%s,sent:%06d,sent time:%40s,Elapsed time:%13s,Avg:%.3f\n", t0, i, t1, t1.Sub(t0), float64(i)/t1.Sub(t0).Seconds())
		// fmt.Println("Received:", i, "ReceviedTime:", t0)
		i++
		// fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
