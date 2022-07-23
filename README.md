# HTTP API voting service
# Running the application without Docker
```
go run .
```
## HTTP Methods
```
/ping — Checking the server connection

    example: 
        "GET" :8000/ping
```
```
/createpoll — Create a poll with answer options
    options:
        title — Name of poll
        options — Answer options

    example: 
        "POST" :8000/createpoll?title=RustVSGolang&options=Golang,Rust
```
```
/poll — Vote for a specific option
    options:
        title — Name of poll
        id — Poll id
        option — Selected option

    example: 
        "PATCH" :8000/poll?title=RustVSGolang&option=Golang
        "PATCH" :8000/poll?id=000f574a&option=Golang
```
```
/getpoll — Get a result on a particular poll
    options:
        title — Name of poll
        id — Poll id

    example:
        "GET" :8000/getpoll?title=RustVSGolang
        "GET" :8000/getpoll?id=000f574a
```
```
/deletepoll — Delete poll
    options:
        title — Name of poll
        id — Poll id

    example: 
        "DELETE" :8000/deletepoll?title=RustVSGolang
        "DELETE" :8000/deletepoll?id=000f574a
```
### Params for ```.env``` file
```
MONGO=mongodb://localhost:27017
DATABASE=voting_service
COLLECTION=polls
HOST=127.0.0.1
PORT=8000
```