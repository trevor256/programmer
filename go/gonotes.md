Built-in Basic Types in Go
one boolean built-in boolean type: bool.
11 built-in integer numeric types: int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, and uintptr.

two built-in floating-point numeric types: float32 and float64.

two built-in complex numeric types: complex64 and complex128.

one built-in string type: string.



for some reason when building on fedora for lambda the c libs cause an issue and it must be built with command CGO_ENABLED=0
	// GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
