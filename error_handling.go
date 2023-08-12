package main
import(
	"fmt"
	"bufio"
	"os"
	"log"
)
var print = fmt.Println
func main(){
	print("What's your name")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if(err == nil){
		print("Hello" ,name)
	}else{
		log.Fatal(err)
	}
} 