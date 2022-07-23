# HTTP API voting service
# Running the application without Docker
```
go run .
```
## HTTP Methods
```
"GET" /ping — Checking the server connection

    example: 
        "GET" :8000/ping
```
```
"POST" /poll — Create a poll with answer options
    options:
        title — Name of poll
        options — Answer options

    example: 
        "POST" :8000/poll?title=RustVSGolang&options=Golang,Rust
```
```
"GET" /poll — Get a result on a particular poll
    options:
        title — Name of poll
        id — Poll id

    example:
        "GET" :8000/poll?title=RustVSGolang
        "GET" :8000/poll?id=000f574a
```
```
"PATCH" /poll — Vote for a specific option
    options:
        title — Name of poll
        id — Poll id
        option — Selected option

    example: 
        "PATCH" :8000/poll?title=RustVSGolang&option=Golang
        "PATCH" :8000/poll?id=000f574a&option=Golang
```
```
"DELETE" /poll — Delete poll
    options:
        title — Name of poll
        id — Poll id

    example: 
        "DELETE" :8000/poll?title=RustVSGolang
        "DELETE" :8000/poll?id=000f574a
```
### Params for ```.env``` file
```
MONGO=mongodb://localhost:27017
DATABASE=voting_service
COLLECTION=polls
HOST=127.0.0.1
PORT=8000
```