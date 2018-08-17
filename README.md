Error or Panic
-

[![Maintainability](https://api.codeclimate.com/v1/badges/caf54b60fe8ae4d9a423/maintainability)](https://codeclimate.com/github/cn007b/eop/maintainability)

Here you can find 2 approaches how to deal with errors:
first one - return error, another one - panic.
<br>Which better - it's up to you!

## Prerequisites

Common stuff between all approaches - workflow:
here I have one controller which calls service function
which works like a facade and performs sign up in:

* facebook
* twitter
* pinterest

you can see it on sequence diagram:

![sequence diagram](/sequenceDiagram.png)

## Usage

````bash
git clone https://github.com/cn007b/eop.git && cd eop
export GOPATH=$PWD

# 1: return error
go run -race src/app/main.go --strategy=error
go run -race src/app/main.go --strategy=error --username=bond
go run -race src/app/main.go --strategy=error --username=james

# 2: panic
go run -race src/app/main.go --strategy=panic
go run -race src/app/main.go --strategy=panic --username=bond
go run -race src/app/main.go --strategy=panic --username=james

# 3: return error from goroutine
go run -race src/app/main.go --strategy=errorroutine
go run -race src/app/main.go --strategy=errorroutine --username=bond
go run -race src/app/main.go --strategy=errorroutine --username=james

# 4: panic in goroutine
go run -race src/app/main.go --strategy=panicroutine
go run -race src/app/main.go --strategy=panicroutine --username=bond
go run -race src/app/main.go --strategy=panicroutine --username=james

# 5: panic in goroutine like PRO 😎
go run -race src/app/main.go --strategy=panicroutinepro
go run -race src/app/main.go --strategy=panicroutinepro --username=bond
go run -race src/app/main.go --strategy=panicroutinepro --username=james
````
