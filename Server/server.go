package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
)

/*
	1. Create a 3 chat rooms -> 3 maps for each room storing userId and net.Conn
	2. Wait for clients to connect -> handleConnection()
	3. Propose clients like which room to join -> joinRoom(conn)
	4. Take users input and assign room to each user -> receiveMessage(conn,map)
	5. Start broadcasting messages -> sendMessage(conn,map)
*/

var room1 = make(map[string]net.Conn)
var room2 = make(map[string]net.Conn)
var room3 = make(map[string]net.Conn)
var connectedSync sync.Mutex


func handleConnection(conn net.Conn){
	// 1. Generate unique Id for any new connection
	// 2. Call joinRoom to make the client join a specfic room

	clientUId := uuid.New().String()
	fmt.Println("New Id generated: " + clientUId)

	go joinRoom(conn, clientUId)
}

func joinRoom(conn net.Conn, clientId string){
	for{
		_,err := fmt.Fprint(conn,"Enter the room of your choice 1. Sports 2. Games  3. Geopolitics\n")
		
		if err != nil {
			fmt.Println("Could not receive clients choice of the room ;( " + err.Error())
			conn.Close()
			fmt.Println("Connection : " + conn.RemoteAddr().String() + " disconnected!")
		}
		
		choice,er2 := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("We recieved: " + choice)
	
		if er2 != nil {
			fmt.Println("Failed to read client input: " + err.Error())
			return
		}
		choice = strings.TrimSpace(choice)
		i, er3 := strconv.Atoi(choice)

		fmt.Printf("After conversion: %d \n", i)
		if er3 != nil || (i < 0 || i > 3) {
			continue
		}

		if i == 1 {
			connectedSync.Lock()
			room1[clientId] = conn
			connectedSync.Unlock()

			go receiveMessage(conn,room1,clientId)
			
		
		} else if i == 2 {
			connectedSync.Lock()
			room2[clientId] = conn 
			connectedSync.Unlock()

			go receiveMessage(conn,room2,clientId)
		} else{
			connectedSync.Lock()
			room3[clientId] = conn
			connectedSync.Unlock()

			go receiveMessage(conn,room3,clientId)
		}
		return 

	}
	
}

func receiveMessage(conn net.Conn, mp map[string]net.Conn, clientUUID string){

	for{
		message,err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error receiving data from: " + conn.RemoteAddr().String() + ": " + err.Error())
			conn.Close()
			break
		}

		for id,cn := range(mp) {
			
			_,er2 := fmt.Fprint(cn,clientUUID + ": " + message)

			if er2 != nil {
				fmt.Println("Could not send message to: " + id)
			}

			
		}
	}
	

}
func main(){

	ln, err := net.Listen("tcp",":8000")

	if err != nil {
		fmt.Println("Failed to start server: " + err.Error())
	}

	fmt.Println("Server started listening ...")

	for{
		conn,err := ln.Accept()

		if err != nil {
			fmt.Println("Error listening to client " + err.Error())
		}

		go handleConnection(conn)
	}
}