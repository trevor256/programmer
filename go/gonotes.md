=========Types in Go

Literals are also unicode ??????

11 built-in integer numeric types: int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, and uintptr. However just use int, it will default choose  (32 or 64) based on the build. latter update the intagers for performance optimizations.  

Two built-in floating-point numeric types: float32 and float64. floats arnt perfect so dont use them to represent money, use a exact decimal represintaion. 

Two built-in complex numeric types: complex64 and complex128. These arnt great and should not be used. If you want to use crazy mathmatical computing use a third party package. 

one built-in string type: string.

-------------Intager operators

+ - * / % 
can combine with = , += -= ...etc 

compare intagers with == != > >= < <=
bit-manipulation by bitshifting left << and right >>
bit mask with & and , | or , ^ XOR , &^ and not .  you can aso cobine with = , &= |= ...etc

-------------converting variables 
int(variable name) 

-------------var Versus :=

if you are in a function you can declare a var like this:  x := 10

most verbose:  var x int = 10 
less verbose:  var x = 10

multible var: var x, y int = 10 , 20 
diffrent types: var x, y = 10, "hello"

a lot of multiable var: var (
				x int
				y  =20 
				d, e = 40, "ello"
			     )



=========Build issues 

for some reason when building on fedora for lambda the c libs cause an issue and it must be built with command CGO_ENABLED=0
	// GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
