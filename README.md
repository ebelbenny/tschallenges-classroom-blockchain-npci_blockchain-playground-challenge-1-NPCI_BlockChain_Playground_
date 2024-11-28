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
Just three commands need to work:```
/users  - See who's online
/quit   - Leave the chat
hello   - Send "hello" to everyone
```

## Getting Started
1. Start with the provided code below
2. Run the server: `go run server.go`
3. Connect using netcat: `nc localhost 8080`
4. Type messages to chat!

## What You Need to Know
- Basic Go syntax
- How to run Go programs
- What a goroutine is

## Grading
- Does it work? (50%)
- Is the code clean? (25%)
- Does it handle errors? (25%)
