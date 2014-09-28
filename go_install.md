---
layout: main
---

## Golang installation Guide


### 1. Download Golang
Download Golang package according to your OS.

[Golang Official Download Site](https://golang.org/dl/)
	
If your network blocks by GFW, use the link below
	
[Golang China mirror Download Site](http://www.golangtc.com/download)
	
**NOTE:** for windows platform, please download the msi installer file.
	
### 2. Install Golang
#### Windows Platform:

#####MSI installer

Open the MSI file and follow the prompts to install the Go tools. By default, the installer puts the Go distribution in **c:\Go**.

The installer should put the **c:\Go\bin** directory in your PATH environment variable. You may need to restart any open command prompts for the change to take effect. 
	
#### Linux Platform:

Extract the package into **/usr/local**, creating a Go tree in **/usr/local/go**. For example:

```bash
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

NOTE: Choose the archive file appropriate for your installation. For instance, if you are installing Go version 1.2.1 for 64-bit x86 on Linux, the archive you want is called go1.2.1.linux-amd64.tar.gz.

(Typically these commands must be run as root or through sudo.)

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

```bash
export PATH=$PATH:/usr/local/go/bin
```


##### Installing to a custom location

The Go binary distributions assume they will be installed in **/usr/local/go** (or c:\Go under Windows), but it is possible to install the Go tools to a different location. In this case you must set the GOROOT environment variable to point to the directory in which it was installed.

For example, if you installed Go to your home directory you should add the following commands to **$HOME/.profile**:

```bash
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin	
```

**Note:** GOROOT must be set only when installing to a custom location. 
	
### 3. Test your installation

Check that Go is installed correctly by building a simple program, as follows.

Create a file named hello.go and put the following program in it:

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

Then run it with the go tool:

```bash
$ go run hello.go
hello, world
```

If you see the "hello, world" message then your Go installation is working. 
	
### 4. The GOPATH environment variable

The GOPATH environment variable specifies the location of your workspace. It is likely the only environment variable you'll need to set when developing Go code.

To get started, create a workspace directory and set GOPATH accordingly. Your workspace can be located wherever you like, but we'll use **$HOME/go_workspace** in this guide. 


```bash
$ mkdir $HOME/go_workspace
$ export GOPATH=$HOME/go_workspace
```

For convenience, add the workspace's bin subdirectory to your PATH:

```bash
$ export PATH=$PATH:$GOPATH/bin
```