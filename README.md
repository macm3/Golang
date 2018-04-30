# My-First-Go-Project
This is my first Go project. Duh.

Beside the basic "Hello World" program, I'm trying to catch up with this course: https://www.alura.com.br/curso-online-golang. This course is all about do a web site monitoring system to see wich one web site is on/off.
Also I'm trying to make my first Go Server. Let's see what I can do.


**0. Running**

In order to run this project you must install Go language here: https://golang.org/doc/install

**1. Clone the repository**
```
$ git clone https://github.com/macm3/Golang.git
```
**2. Travell to repository's folder from terminal command line**

**3. Make sure you have Go installed**

```
$ go version
```

Go could be very trick when it comes to packages and working directory.
Make sure you have your GOPATH and GOROOT set it up by following the steps that follows:

```
$ export GOPATH=$HOME/go-workspace #don't forget to change your path correctly!
```
```
$ export GOROOT=/usr/local/opt/go/libexec
```
```
$ export PATH=$PATH:$GOPATH/bin
```
```
$ export PATH=$PATH:$GOROOT/bin
```

**4. Building and running your project**

You don't have to build your file. Golang it's beautiful and it can build and run your project at the same time.
```
$ go run program.go
```
But you can do it like the old times too.
```
$ go build program.go
```
```
$ go run 
```
**5. Let's learn and have fun!!!!!!**
