package main
import(
	"fmt"
	"unicode/utf8"
)
var print = fmt.Println
func main(){
	rStr := "abcdefghijklmnopqrstuvwxyz"
	print("Count :", utf8.RuneCountInString(rStr))
	for i,value := range rStr{
		fmt.Printf("%d : %#U \n", i, value)
	}
}