package main
import(
	"fmt"
	"strings"
)
var print = fmt.Println
func main(){
	var sV string
	print("enter any string")
	_, err01 := fmt.Scan(&sV)
	if err01 != nil {
		print("Error :", err01)
	}
	print("enter any alphabet :")
	var x string // for char
	_, err02 := fmt.Scan(&x)
	if err02 != nil {
		print("Error :", err02)
	}
	print("index of 1st occurence of", x)
	print(strings.Index(sV,x))
}