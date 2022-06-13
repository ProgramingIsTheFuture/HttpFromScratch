package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)
func response(conn net.Conn, n int) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", n)
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
}

func homeHandler(conn net.Conn) {
	body := "<h1>Home page</h1><a href=\"/about\">About page</a>"
	response(conn, len(body))
	fmt.Fprint(conn, body)
}

func aboutHandler(conn net.Conn) {
	body := "<h1>About page</h1><h4>Http server made with TCP from scratch!</h4><a href=\"/\">Go back</a>"
	response(conn, len(body))
  fmt.Fprint(conn, body)
}

func handleRequest(conn net.Conn) {
  defer conn.Close()

  fmt.Println("Reading the connection")

  b, err := bufio.NewReader(conn).ReadBytes('\n')
  // b, err := io.ReadAll(conn)
  if err != nil {
  	fmt.Printf("Failed read all conn: %s", err.Error())
  	return
  }

	header := string(b)
	method := strings.Fields(header)[0]
	url := strings.Fields(header)[1]
  if url == "/"  && method == "GET"{
  	homeHandler(conn)
  } else if url == "/about" && method == "GET" {
		aboutHandler(conn)
  }
}

func main() {

  ls, err := net.Listen("tcp", "127.0.0.1:8000")
  if err != nil {
    panic(err)
  }

  for {
    conn, err := ls.Accept()
    if err != nil {
      fmt.Println("Failed accepting the connection")
      continue
    }

    go handleRequest(conn)
  }

}
