Error or Panic
-

Here you can find 2 approaches how to deal with errors,
first one - return error, another one - panic.
Which better - it's up to you!

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

# 1
go run src/app/main.go --strategy=error
go run src/app/main.go --strategy=error --username=bond
go run src/app/main.go --strategy=error --username=james

# 2
go run src/app/main.go --strategy=panic
go run src/app/main.go --strategy=panic --username=bond
go run src/app/main.go --strategy=panic --username=james

# 3
go run src/app/main.go --strategy=errorroutine
go run src/app/main.go --strategy=errorroutine --username=bond
go run src/app/main.go --strategy=errorroutine --username=james

# 4
go run src/app/main.go --strategy=panicroutine
go run src/app/main.go --strategy=panicroutine --username=bond
go run src/app/main.go --strategy=panicroutine --username=james

# 5 😎
go run src/app/main.go --strategy=panicroutinepro
go run src/app/main.go --strategy=panicroutinepro --username=bond
go run src/app/main.go --strategy=panicroutinepro --username=james
````
