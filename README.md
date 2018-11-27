# Chat in Golang Using Gorilla Toolkit Framekework

## Description

This Project is a simple application using [http://www.gorillatoolkit.org](URL "Gorilla Toolkit") to creates a chat with session and websocket communication.

### Gorilla Toolkit mux

Implements a request router and dispatcher.

### Gorilla Toolkit sessions

Provides cookie and filesystem sessions and infrastructure for custom session backends.

##### Environment Variable

`SESSION_KEY` is a value to identificate session id.

Note: Don't store your key in your source code. Pass it via an environmental variable, or flag (or both), and don't accidentally commit it alongside your code.

### Commands

`make build` Build Docker Image with app.

`make push` Upload Docker Image to repository.

`make run` Run app with Docker for port `8000`.
