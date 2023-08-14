package main
import(
	"fmt"
	"strconv"
	"reflect"
)
var print = fmt.Println
func main(){
	// conversion of one datatype to another
	a := 2.2 //float64
	b := int(a) //int
	print(b)
	//string -> int
	C := "1000" //string
	D, err := strconv.Atoi(C)
	print(D, err, reflect.TypeOf(D)) //int
	//int -> string
	E := 10
	F := strconv.Itoa(E)
	print(F, reflect.TypeOf(F))
	G := fmt.Sprintf("%f is string", 3.14) //float to string
	print(G)
}