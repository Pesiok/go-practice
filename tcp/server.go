package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		// a listener
		// can handle one connection at time
		// when it closes the connection it can handle another one
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// go handleConnection(conn)
		// go handleRoot13(conn)
		go httpServer(conn)
	}
}

func httpServer(conn net.Conn) {
	defer conn.Close()

	request(conn)

	// respond(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line

			// router
			mux(conn, ln)
		}
		if ln == "" {
			// headers are finished
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("***METHOD", m)
	fmt.Println("***URL", u)

	// router
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		// ...etc
	}
}

func index(conn net.Conn) {
	body := `
	<!doctype html>
	<html lang="en">
		<head>
			<meta charset="utf-8">
			<title>WILLKOMMEN</title>
		</head>
		<body>
			<h1>Hi there</h1>
			<script>alert('and hi there!')</script>
		</body>
	</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func handleConnection(conn net.Conn) {
	// close after 10 seconds
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	// scanner will go on forever...
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You said: %s\n", ln)
	}
	defer conn.Close()
}

func handleRoot13(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		// ascii 97 - 122
		if v <= 109 {
			// we are shifting numbers by 13 places
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
