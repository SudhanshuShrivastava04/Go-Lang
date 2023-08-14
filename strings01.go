package main
import(
	"fmt"
	"strings"
)
var print = fmt.Println
func main(){
	print("enter name")
	var name string
	_,err := fmt.Scan(&name)
	if err != nil {
		print("Error :", err)
	}
	a := "hello world!"
	print("before :", a,"--> length :", len(a), "--> contains name?", strings.Contains(a,name))
	replacer := strings.NewReplacer("world" , name)
	b := replacer.Replace(a)
	print("after : ", b,"--> length :", len(b), "conatins name?", strings.Contains(b,name))
}