# Ultimate Go Debugging Course

This repo holds the example code and course information for the Ultimate Go Debugging course.

This course will teach you various techniques that can be used to debug a Go program.

## Getting Started

First, you should have the latest version of Go. You can get the latest version from https://go.dev/dl/.
Once you download the tarball follow the instructions on https://go.dev/doc/install to install Go on your system.
To verify a correct Go install run the following command:

```
go version
```

You should see output similar to:

```
go version go1.17.6 linux/amd64
```

Second, once you have Go installed you must install Delve. You can install the latest version of Delve via:

```
go install github.com/go-delve/delve/cmd/dlv@latest
```

You can verify Delve has been correctly installed by running the following command:

```
dlv help
```

The full install documentation is [here](https://github.com/go-delve/delve/tree/master/Documentation/installation) if needed.