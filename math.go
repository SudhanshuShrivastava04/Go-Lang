package main
import(
	"fmt"
	"math"
)
var print = fmt.Println
func main(){
	print("Abs :", math.Abs(-42))
	print("Power 4^2 : ", math.Pow(4,2))
	print("Sqrt: ", math.Sqrt(42))
	print("Cbrt : ", math.Cbrt(42))
	print("Ceil(4.2) : ", math.Ceil(4.2))
	print("Floor(4.2) : ", math.Floor(4.2))
	print("Round(4.2) : ", math.Round(4.2))
	//log maths
	print("log10(100) :", math.Log10(100))
	print("log2(8) :", math.Log2(8))
	print("Max(4,2) :", math.Max(4,2))
	print("Min(4,2) :", math.Min(4,2))
	// radian and degree
	r90 := 90 * math.Pi/180
	print(r90)
	d90 := r90 * (180/math.Pi)
	print(d90)
	print("Sin(90) :", math.Sin(r90))
}