package test_methods

import (
	"container/list"
	"log"
	"net"
	"time"
)

func tcpPing(host string, port string) (int64, int, *list.List) {
	var alt int64
	var suc int
	var fac int
	delays := list.New()

	for {
		if fac >= 3 || (suc != 0 && fac + suc >= 10) {
			break
		}

		now := time.Now()

		conn, err := net.DialTimeout("tcp", host + ":" + port, 2 * time.Second)
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}
		conn.Close()

		delay := time.Now().UnixNano() / 1e6 - now.UnixNano() / 1e6
		alt += delay
		suc += 1
		delays.PushBack(delay)

		log.Println(delay)
	}

	if suc == 0 {
		return 0, 0, delays
	}

	return alt / int64(suc), suc / (suc + fac), delays
}

func googlePing(host string, port string) (int64, int, *list.List) {
	var alt int64
	var suc int
	var fac int
	delays := list.New()

	for {
		if fac >= 3 || (suc != 0 && fac + suc >= 10) {
			break
		}

		now := time.Now()

		conn, err := net.DialTimeout("tcp", host + ":" + port, 2 * time.Second)
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}
		_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		_ = conn.SetWriteDeadline(time.Now().Add(2 * time.Second))

		_, err = conn.Write([]byte("\x05\x01\x00"))
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}

		buf := make([]byte, 2)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}
		log.Println(string(buf[0:cnt]))

		_, err = conn.Write([]byte("\x05\x01\x00\x03\x0agoogle.com\x00\x50"))
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}

		buf2 := make([]byte, 10)
		cnt2, err := conn.Read(buf2)
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}
		log.Println(string(buf2[0:cnt2]))

		_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\nUser-Agent: curl/11.45.14\r\n\r\n"))
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}

		buf3 := make([]byte, 1)
		_, err = conn.Read(buf3)
		if err != nil {
			log.Println(err)
			fac += 1
			delays.PushBack(0)
			continue
		}
		log.Println(buf3)

		conn.Close()

		delay := time.Now().UnixNano() / 1e6 - now.UnixNano() / 1e6
		alt += delay
		suc += 1
		delays.PushBack(delay)
		log.Println(delay)
	}

	if suc == 0 {
		return 0, 0, delays
	}

	return alt / int64(suc), suc / (suc + fac), delays
}