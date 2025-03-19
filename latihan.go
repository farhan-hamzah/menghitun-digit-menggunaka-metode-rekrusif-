// Menghitung Berapa Banyak Digit
package main
import "fmt"

func jumlahDigit(n int)int{
	if n <= 0{
		return 0
		}
	return 1 + jumlahDigit(n/10)
}
func main(){
	var masukan int
	fmt.Scan(&masukan)
	fmt.Print(jumlahDigit(masukan))
}
