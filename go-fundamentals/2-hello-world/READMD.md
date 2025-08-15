# Writing Tests
You need to satisfy a few requirements below to write a test in Go:
- It needs to be in a file named like `xxx_test.go`
- A test function must start with `Test`
- A test function must take only one argument `t *testing.T`
- To use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in other file.
- Do `t.Fail()` when you want to fail

# Go's documentation
There are a few options to see documentation:
1. go doc [something]
2. pkgsite -open .
    - prerequisite:  `go install golang.org/x/pkgsite/cmd/pkgsite@latest` 

# Go Subtest
Use `t.Run(...)` to run subtests in Go