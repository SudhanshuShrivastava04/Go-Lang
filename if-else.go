package main
import(
	"fmt"
)
var print = fmt.Println
func main(){
	print("enter your Age")
	var age int
	_,err := fmt.Scan(&age)
	if err != nil {
		print("Error : ", err)
		return
	}
	if age < 18{
		print("you are under 18")
	}else{
		print("you are above 18")
	}
}