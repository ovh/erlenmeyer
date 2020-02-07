# Build erlernmeyer

As we moved from [dep]() to [go modules](https://github.com/golang/go/wiki/Modules#example), you have several step to follow to build this project.

## Requirements

You need to have go1.11 at least. To make it work with go 1.11: 
- You can keep this project in your `GOPATH`. 
- Activate the `GO111MODULE` running:

```sh
export GO111MODULE=on  
```

You can also set a proper env with https://direnv.net/. 

You can easily build this project using go1.12: 
- You can clone this project outside your `GOPATH`
- You will still need to activate the `GO111MODULE`: 

```sh
export GO111MODULE=on # Can be on or auto if you build erlenmeyer outside your GOPATH
```

You also need to have `gofmt` (coming with you go setup) and [golangci](https://github.com/golangci/golangci-lint).

## Init, dependancies and dev mode

First init your project locally running:

```sh
make init
```

Then download the project dependencies

```sh
make dep
```

And finally build the dev version

```sh
make dev
```

If you get any issue during compilation reset you go.mod file from Github:

```sh
git checkout -- go.mod
```

## Release

To compile erlenmeyer release simply run:

```sh
make release
```

## Run

Run the dev compiled version:

```sh
./build/erlenmeyer --config /Path/to/erlenmeyer.yaml
```
