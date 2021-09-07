# Go Chat

In this lab, you'll build a simple chat client and server with Go. To ensure 
compatibility across all the different student implementations in the class, we
will use Protocol Buffers to create our chat *wire format*.

You are given a .proto file and some starter code; it is your job to finish the
required functionality outlined below.


## Client

The basics are already done for you, so start here to familiarize yourself with
the codebase. You will need to implement registration and direct messages.

### Registration

The implementation you start with auto-generates usernames, so adapt it to send
a registration message to the server. If registration fails (likely due to the
username already being taken), disconnect the client. This should be the **first**
thing a client does.

### Private Messages

If the user inputs a `/` followed by a username, send a private
message to that user. For example:

```
/matthew hey there! Go is pretty cool, right?
```

Will be sent to the user 'matthew' only. The message needs to be routed through
the server since it will be able to map usernames to corresponding sockets.


## Server

The server supports three message types:
* Registrations
* Basic messages (these should be sent to everyone)
* Direct messages (these are sent to a specific user)

## Creating a Module

Go modules are sort of like Java packages. To create a module for this lab, you can use:

```
go mod init chat
```

And then

```
go mod tidy
```

## Running the Program

```
cd chat
go run server/server.go

# (on another machine)
go run client/client.go username hostname
```

# Grading

Since we are all using the same wire format, testing will be fun: you should
be able to handle several connections from other students in the class and
demonstrate client/server functionality is working correctly.

If this lab is too easy and you're looking for a challenge, write a benchmark
script that launches a large number of clients to determine what the upper
bound for client connections is.
