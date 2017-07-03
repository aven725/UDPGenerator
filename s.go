package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		//		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	var tS int64
	var tR int64
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)
	i := 0
	//	t0 := time.Now()
	for {
		//		if i == 0 {
		//			t0 = time.Now()
		//		}
		_, _, err := ServerConn.ReadFromUDP(buf)
		//		t1 := time.Now()
		//		fmt.Printf("Program start time:%s,recevie:%06d,recevie time:%40s,Elapsed time:%13s,Avg:%.3f\n", t0, i, t1, t1.Sub(t0), float64(i)/t1.Sub(t0).Seconds())
		// fmt.Println("Received:", i, "ReceviedTime:", t0)
		i++
		tS = BytesToInt64(buf[0:8])
		tR = time.Now().UnixNano()

//		fmt.Printf("Sent Time:%d,Received Time:%d,Sub Time(ms):,%f\n", tS, tR, float64(tR-tS)/1000000)
		fmt.Printf("%d,%.3f\n",i,float64(tR-tS)/1000000)

		// fmt.Println(tS, tR)

		if err != nil {
			//			fmt.Println("Error: ", err)
		}
	}
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
