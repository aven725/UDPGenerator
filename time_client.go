package main

import (
	// "encoding/hex"
	"bytes"
	"encoding/binary"
	// "fmt"
	"net"
	// "strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		//		fmt.Println("Error: ", err)
	}
}

func main() {
	var t0 time.Time
	var t1 time.Time
	breakTime := 10
	ServerAddr, err := net.ResolveUDPAddr("udp", "192.168.100.1:10001")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "192.168.100.2:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	str := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	// packetin, _ := hex.DecodeString(str)
	i := 0
	//	t0 := time.Now()
	//	for i=0;i<5000000;i++{
	for {
		buf := ByteCombine(Int64ToBytes(time.Now().UnixNano()), []byte(str))
		//		t1 := time.Now()
		//		fmt.Printf("Program start time:%s,sent:%06d,sent time:%40s,Elapsed time:%13s,Avg:%.3f\n", t0, i, t1, t1.Sub(t0), float64(i)/t1.Sub(t0).Seconds())
		// fmt.Println("Program Start Time:", t0, "sent:", i, "sent time:", t1, "Elapsed time:", t1.Sub(t0), "avg:", float64(i)/t1.Sub(t0).Seconds())
		// msg := strconv.Itoa(str)
		if i == 0 {
			t0 = time.Now()
		}
		_, err := Conn.Write(buf)
		i++
		t1 = time.Now()

		if t1.Sub(t0).Seconds() > float64(breakTime) {
			break
		}

		if err != nil {
			//			fmt.Println(str, err)
		}
		time.Sleep(time.Microsecond * 1) // 0.03ms
		// time.Sleep(time.Second * 5)
	}
}

func ByteCombine(a []byte, b []byte) []byte {
	var buffer bytes.Buffer

	for i := 0; i < len(a); i++ {
		buffer.WriteByte(a[i])
	}

	for i := 0; i < len(b); i++ {
		buffer.WriteByte(b[i])
	}

	return buffer.Bytes()
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
