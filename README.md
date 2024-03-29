# gtfo

go program to fetch gtfobins and lolbas

## setup

**to build and compile to an executable:**

``` 
$ git clone https://github.com/lordlabuckdas/gtfoo.git && cd gtfoo
$ go build -ldflags="-s -w" gtfoo.go
$ ./gtfoo <args>
```

**to temporarily run it:**

``` 
$ git clone https://github.com/lordlabuckdas/gtfoo.git && cd gtfoo
$ go run gtfoo.go <args>
```

## note

* for now, use `go run` to execute the program
* run `go fmt ./..` to fix indentation and format the `go` files
* populate [gtfobins.go](./gtfobins/gtfobins.go) with helper functions
* upload the binary to this repo, you should not
* put up in releases, it will be

## color scheme

The project uses [ANSI escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code) for colored output

| output                     | _colorVariable_ |
|----------------------------|-----------------|
| function name (ex: `Sudo` )| _colorCyan_     |
| function description       | _colorGreen_    |
| code                       | _colorYellow_   |
| code description           | _colorWhite_    |
| errors                     | _colorRed_      |

## contributing

Please see our [contributing guidelines](./CONTRIBUTING.md)
