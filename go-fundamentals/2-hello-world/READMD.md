# Writing Tests
You need to satisfy a few requirements below to write a test in Go:
- It needs to be in a file named like `xxx_test.go`
- A test function must start with `Test`
- A test function must take only one argument `t *testing.T`
- To use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in other file.