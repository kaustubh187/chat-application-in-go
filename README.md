# Chat Room Server in Go
This Go application creates a TCP-based chat server that allows multiple clients to join and communicate within different chat rooms. Clients can choose from one of three available chat rooms (Sports, Games, or Geopolitics) and send messages that are broadcasted to all other participants in the same room.

## Overview
The server application operates as follows:

- Create Rooms: It initializes three chat rooms using maps to store each room’s participants and their connections.
- Client Connections: It listens for incoming client connections and assigns each client a unique ID.
- Room Selection: Once connected, clients are prompted to select a room to join.
- Message Handling: The server listens to each client's messages and broadcasts them to all other clients in the same room.


## Features
- Multi-room Support: Three separate chat rooms (Sports, Games, Geopolitics) allow clients to select and chat with users interested in specific topics.
Concurrent Connections: Uses Go routines to handle multiple client connections simultaneously.
- Synchronized Map Access: Uses a mutex to synchronize access to each room's map, ensuring thread safety for concurrent reads and writes.

# Getting Started
## Prerequisites
-> Go (v1.15+)

![image](https://github.com/user-attachments/assets/30ee8174-53bc-4ee9-8caf-73d09b3bf9ae)
![image](https://github.com/user-attachments/assets/44f6b828-77bc-4fe9-a37b-247d0a94d4d0)
![image](https://github.com/user-attachments/assets/9ddff246-8c30-4bc8-9117-1c95ea0afd51)
