package main
import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0.0

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var jumlahBalita int
	var beratBalita arrBalita
	var beratMin, beratMax, rataRata float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jumlahBalita)

	if jumlahBalita > 100 {
		fmt.Println("Jumlah balita tidak boleh lebih dari 100.")
		return
	}

	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	hitungMinMax(beratBalita, jumlahBalita, &beratMin, &beratMax)
	rataRata = rerata(beratBalita, jumlahBalita)

	fmt.Printf("Berat minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}