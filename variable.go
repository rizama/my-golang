package main
import "fmt"

func main()  {
	var name string
	name = "Sam"
	fmt.Println(name)

	var lastname = "Pratama"
	fmt.Println(lastname)

	address := "Bandung"
	fmt.Println(address)

	// Multiple
	var(
		country = "Indonesia"
		province = "Jawa Barat"
	)

	fmt.Println(country)
	fmt.Println(province)
}