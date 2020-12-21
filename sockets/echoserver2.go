package main

import ("io"
	"log"
	"net")

func echoserver(conn net.Conn){
	defer conn.Close()
		prompt :=
`
WelCome to my LegalLoli Paradise			
`
		conn.Write([]byte(prompt))
	for {
		message := make([]byte,512)
		_,err := conn.Read(message)
		if err == io.EOF{
			bye := []byte("bie bye senpai UwU\n")
			conn.Write(bye)
			break
		}
		if err!=nil{
			log.Fatalln("Could not read data sent from client ->",err)
		}
		_,err = conn.Write(message)
		if err!=nil{
			log.Fatalln("Error send data to clinet")
		}
	}
}


func main(){
	server,err := net.Listen("tcp","127.0.0.1:8080")
	if err!=nil{
		panic("Unable to bind to port")
	}
	log.Println("Sever Started at port 8080")
	for {
		conn,err := server.Accept()
		if err!=nil{
			log.Fatalln("Client Dropped")
		}
		log.Println("Connection Recived from ",conn.RemoteAddr())
		go echoserver(conn)
	}

}
