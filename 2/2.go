package main
import "fmt"

func main() {
	var jumlahIkan, jumlahWadah int
	var ikanPerWadah [1000]int
	var totalBeratPerWadah [1000]float64
	var totalBeratSemuaWadah, rataRataBerat float64

	fmt.Print("Masukkan jumlah wadah (x): ")
	fmt.Scan(&jumlahWadah)
	fmt.Print("Masukkan jumlah ikan (y): ")
	fmt.Scan(&jumlahIkan)

	if jumlahWadah > 1000 || jumlahIkan > 1000 {
		fmt.Println("Jumlah wadah atau ikan tidak boleh lebih dari 1000.")
		return
	}

	fmt.Println("Masukkan banyak ikan untuk setiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Banyak ikan di wadah ke-%d: ", i+1)
		fmt.Scan(&ikanPerWadah[i])
	}

	fmt.Println("Masukkan berat ikan yang dijual:")
	for i := 0; i < jumlahIkan; i++ {
		var beratIkan float64
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan)

		wadah := i % jumlahWadah
		totalBeratPerWadah[wadah] += beratIkan
	}

	for i := 0; i < jumlahWadah; i++ {
		totalBeratSemuaWadah += totalBeratPerWadah[i]
	}

	rataRataBerat = totalBeratSemuaWadah / float64(jumlahWadah)

	fmt.Println("Total berat ikan di setiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah ke-%d: %.2f kg\n", i+1, totalBeratPerWadah[i])
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f kg\n", rataRataBerat)
}