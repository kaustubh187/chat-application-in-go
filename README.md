# Chat Room Server in Go
This Go application creates a TCP-based chat server that allows multiple clients to join and communicate within different chat rooms. Clients can choose from one of three available chat rooms (Sports, Games, or Geopolitics) and send messages that are broadcasted to all other participants in the same room.

## Overview
The server application operates as follows:

Create Rooms: It initializes three chat rooms using maps to store each room’s participants and their connections.
Client Connections: It listens for incoming client connections and assigns each client a unique ID.
Room Selection: Once connected, clients are prompted to select a room to join.
Message Handling: The server listens to each client's messages and broadcasts them to all other clients in the same room.


## Features
Multi-room Support: Three separate chat rooms (Sports, Games, Geopolitics) allow clients to select and chat with users interested in specific topics.
Concurrent Connections: Uses Go routines to handle multiple client connections simultaneously.
Synchronized Map Access: Uses a mutex to synchronize access to each room's map, ensuring thread safety for concurrent reads and writes.

# Getting Started
## Prerequisites
-> Go (v1.15+)
