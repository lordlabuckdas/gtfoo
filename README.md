# gtfo

go program to fetch gtfobins and lolbas

## setup

**to build and compile to an executable:**

```
$ git clone https://github.com/lordlabuckdas/gtfoo.git && cd gtfoo
$ go build -ldflags="-s -w" main.go
$ ./main <args>
```

**to temporarily run it:**

```
$ git clone https://github.com/lordlabuckdas/gtfoo.git && cd gtfoo
$ go run main.go <args>
```

## note

* for now, use `go run` to execute the program

* run `go fmt ./..` to fix indentation and format the `go` files

* populate [gtfobins.go](./gtfobins/gtfobins.go) with helper functions

* later, lolbas

* don't upload the binary to the repo, we'll do it under releases later
