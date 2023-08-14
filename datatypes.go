package main
import(
	"fmt"
	"reflect"
)
var print = fmt.Println
func main(){
	// int, float64, bool, string, rune
	print(reflect.TypeOf(1)) //int
	print(reflect.TypeOf(3.3)) //float64
	print(reflect.TypeOf(true)) //bool
	print(reflect.TypeOf("hello")) //string
	print(reflect.TypeOf('😉')) //int32
}