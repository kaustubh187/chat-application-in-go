package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

/*
	1. Dial at port 8000 to connect to server
	2. Call send and receive data
	3. Check if client is already connected or not
*/
var (
	connected     bool
	connectedSync sync.Mutex
)


func main() {
	fmt.Println("Client Started ....")
			
	for{
		connectedSync.Lock()
		alreadyConnected := connected
		connectedSync.Unlock()
		
		
		// If this client is not connected
		if !alreadyConnected{
			conn,err := net.Dial("tcp","127.0.0.1:8000")

			if err != nil {
				fmt.Println("Failed to connect to server" + err.Error())
				time.Sleep(time.Duration(5) * time.Second)
				continue
			}

			fmt.Println(conn.RemoteAddr().String() + ": connected")
			
			connectedSync.Lock()
			connected = true
			connectedSync.Unlock()


			go receiveMessage(conn)
			go sendMessage(conn)
		}
		
		
	}

	

}


func sendMessage(conn net.Conn){

	for{
		
		reader := bufio.NewReader(os.Stdin)
		message,_ := reader.ReadString('\n')

		
		_,err := fmt.Fprint(conn, message)

		if err != nil {
			fmt.Println("Failed to send message: " + err.Error())
		}
	}
}


func receiveMessage(conn net.Conn){

	for {
		
		message,err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			return
		}
	
		fmt.Println("\n\033[33;1m" + message)
		fmt.Print("\033[32mYou:  ")
		
	}


}