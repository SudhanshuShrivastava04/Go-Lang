package main
import(
	"fmt"
	"time"
)
var print = fmt.Println
func main(){
	present := time.Now()
	print(present.Year(), present.Month(), present.Day())
	print(present.Hour(), present.Minute(), present.Second())
}