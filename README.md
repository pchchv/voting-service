# HTTP API voting service
# Running the application without Docker
```
go run .
```
## HTTP Methods
```
/ping — Checking the server connection
```
```
/createPoll — Create a poll with answer options
    options:
        title — Name of poll
        options — Answer options

    example: 
        http://localhost:8000/createPoll?title=RustVSGolang&options=Golang,Rust
```
```
/poll — Vote for a specific option
    options:
        title — Name of poll
        id — Poll id
        option — Selected option

    example: 
        http://localhost:8000/poll?title=RustVSGolang&option=Golang
        http://localhost:8000/poll?id=000f574a&option=Golang
```
```
/getResult — Get a result on a particular poll
    options:
        title — Name of poll
        id — Poll id

    example:
        http://localhost:8000/getResult?title=RustVSGolang
        http://localhost:8000/getResult?id=000f574a
```