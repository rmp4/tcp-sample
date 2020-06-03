package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	res, err := sendTCP("192.168.0.152:4000", "!")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)

	}
}

// func main() {
// 	conn, err := net.Dial("tcp", "192.168.0.152:4000")
// 	if err != nil {
// 		fmt.Println("dial error:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	fmt.Fprintf(conn, "33")
// 	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
// 	// 讀取response
// 	data, err := ioutil.ReadAll(conn)
// 	if err != nil {
// 		if err != io.EOF {
// 			fmt.Println("read error:", err)
// 		}
// 		panic(err)
// 	}
// 	for {
// 		result, err := ioutil.ReadAll(conn)
// 		fmt.Println("data:", string(result))
// 		if err != nil {
// 			if err != io.EOF {
// 				fmt.Println("read error:", err)
// 			}
// 			panic(err)
// 		}
// 	}

// 	// 顯示結果
// 	fmt.Println("response:", string(data))
// 	fmt.Println("total response size:", len(data))
// }
func sendTCP(addr, msg string) (string, error) {
	// connect to this socket
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// send to socket
	//var test = []byte(msg)
	op := make([]byte, 33)

	for i := 0; i < 16; i++ {
		op[i] = 0
	}

	copy(op[13:17], []byte("admin"))
	copy(op[29:33], []byte("moxa"))
	conn.Write(op)

	// listen for reply
	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	return string(result), err
}
func readFully(conn net.Conn) ([]byte, error) {
	//defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
