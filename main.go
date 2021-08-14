package main

import "fmt"

func main() {
	/*
		Array
	*/
	var names = [4]string{"Rahmatulah", "Sidik", "Siapa", "Ya"}
	fmt.Println("Panjang elemen array names :", len(names))
	fmt.Println(names)

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Mangga"
	fruits[2] = "Sirsak"
	fruits[3] = "Jeruk"
	fmt.Println(fruits)
	fmt.Println(fruits[2])

	students := [2]string{
		"Riski",
		"Anwar",
	}

	fmt.Println(students)

	/*
		Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	*/
	numbers := [...]int{
		1,
		2,
		3,
		4,
	}

	fmt.Println("Panjang element array numbers :", len(numbers))

	var data = [...]string{"ayam"}

	fmt.Println(data)

	/*
		Array Multidimensi
	*/
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{10, 11, 12}}
	fmt.Println(numbers1)

	var numbers2 = [2][2]int{
		{4, 5},
		{7, 8},
	}
	fmt.Println(numbers2)

	for i := 0; i < len(names); i++ {
		fmt.Printf("[%d] : %s\n", i, names[i])
	}

	for j, fruit := range fruits {
		fmt.Printf("Elemen ke %d = %s\n", j, fruit)
	}

	/*
		Penggunaan Variabel Underscore _ Dalam for - range
		salah satu kegunaan variabel pengangguran, atau underscore ( _ ).
		Tampung saja nilai yang tidak ingin digunakan ke underscore.
	*/
	fmt.Println()

	for _, name := range names {
		fmt.Println(name)
	}

	// hanya menampilkan index loop, sedangkan value array ditampung sama underscore (_)
	for i, _ := range fruits {
		fmt.Printf("%d \t", i)
	}

	/*
		Alokasi Elemen Array Menggunakan Keyword make
	*/

	fmt.Println()
	var fruits1 = make([]string, 2)
	fruits1[0] = "apple"
	fruits1[1] = "manggo"
	fmt.Println(fruits1) // [apple manggo]

}
