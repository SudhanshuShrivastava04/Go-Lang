package main
import(
	"fmt"
)
var print = fmt.Printf
func main(){
	// %o ==> base 8
	// %x ==> base 16
	// %v ==> guess based upon datatype
	// %T ==> type of supplied value
	print("%d %c %s %t %f %o %x\n", 1,'A',"hello world",true,3.14,1,1)
}