# Simple Chat Server Challenge

## Goal
Build a basic chat server in Go where multiple users can connect and send messages to each other.

## What You'll Build
- A server that lets multiple people chat together
- A way for users to send messages to everyone
- Simple commands like listing users and quitting

## Team Work
- The Playground Challenge 1 is released on Thursday, Nov 28, 2024.
- Prepare for the Challenge individually beforehand.
- Teams will be announced at the start of the session on Saturday, Nov 30, 2024.
- Collaborate with your team in Zoom breakout rooms to implement and finalize your solution.
- Reach out to the instructor for guidance if you have questions or face challenges.
- Present your solution to the instructor for evaluation.

## Requirements
Just three commands need to work:
- /users  - See who's online
- /quit   - Leave the chat
- hello   - Send "hello" to everyone



## Getting Started

1. Start with the provided code in `server.go` and fill in your TODO's.
2. Open the terminal and run the server:`go run server.go`
3. Open a new terminal and install netcat by running the following commandsone by one:

   ` sudo apt update`

   ` sudo apt install -y netcat`
5. Once netcat is installed, connect to the server using: `nc localhost 8080`
6. You will be prompted to enter a username. Type in your desired username and hit Enter. To add more users, repeat `step 4` in additional terminals. You can have as many users as needed.
7. After connecting, you can type messages to chat with other users. You can also use the following commands: `/users` , `/quit` and `hello` 

## What You Need to Know
- Basic Go syntax
- How to run Go programs
- What a goroutine is

## Grading
- Does it work? (50%)
- Is the code clean? (25%)
- Does it handle errors? (25%)