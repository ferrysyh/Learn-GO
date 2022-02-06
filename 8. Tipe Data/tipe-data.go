package main

import "fmt"

func main(){
	// var number_positif uint8 = 10 //0-255
	// var number_positif uint16 = 256 //0-65535
	// var number_positif uint32 = 65535 //0-4294967295
	var number_positif uint64 = 65535 //0-18446744073709551615
	fmt.Println(number_positif)
	
	// var number_negatif int8 = -10 //-128-127
	// var number_negatif int16 = -256 //-32768-32767
	// var number_negatif int32 = -65535 //-2147483648-2147483647
	var number_negatif int64 = -65535 //-9223372036854775808-9223372036854775807
	fmt.Println(number_negatif)

	var desimal = 2.55
	fmt.Println(desimal)

	var apakah_ada bool = true
	fmt.Println(apakah_ada)

	var pesan string = "hallo"
	fmt.Println(pesan)
}