package main

import ("io"
	"net"
	"log")


func echo(conn net.Conn){
	defer conn.Close()
	b := make([]byte,512)
	for {
		size,err := conn.Read(b)
		if err == io.EOF {
			log.Println("Client Disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Recived ",size," bytes ",string(b))
		log.Println("Writing Data")
		if _,err := conn.Write(b); err!=nil{
				log.Fatalln("Unable to write data")
			}
	}

}

func main(){
	listener,err := net.Listen("tcp","127.0.0.1:8080")
	defer listener.Close()
	if err!=nil{
		panic("Unable to bind to port 8080")
	}
	log.Println("Started Listening on 127.0.0.1:8080")
	for {
		conn,err := listener.Accept()
		log.Println("Received Connection from",conn.RemoteAddr())
		if err!=nil{
			log.Fatalln("Failed to connect from client")
			}
		go echo(conn)

	}
}
