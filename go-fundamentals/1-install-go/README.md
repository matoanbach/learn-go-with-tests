## Go Modules
- What is GOPATH?
    - It is an environment variable in Go that specifies the root directory of your workspace, where Go stores source code, compiled packages, and binaries. In modern Go (with modules), it's mainly used for the module cache and installed binaries.
    - People say forget about GOPATH in modern Go.
- go mod init <modulepath>
    - it is to initialize a new module the current directory
