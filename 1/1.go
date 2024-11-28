package main
import "fmt"

func main() {
	var jumlahAnakKelinci int
	var beratAnakKelinci [1000]float64
	var beratTerkecil, beratTerbesar float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahAnakKelinci)
	
	if jumlahAnakKelinci > 1000 {
		fmt.Println("Jumlah anak kelinci tidak boleh lebih dari 1000.")
		return
	}

	fmt.Println("Masukkan berat masing-masing anak kelinci:")
	for i := 0; i < jumlahAnakKelinci; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&beratAnakKelinci[i])
	}

	beratTerkecil = beratAnakKelinci[0]
	beratTerbesar = beratAnakKelinci[0]
	for i := 1; i < jumlahAnakKelinci; i++ {
		if beratAnakKelinci[i] < beratTerkecil {
			beratTerkecil = beratAnakKelinci[i]
		}
		if beratAnakKelinci[i] > beratTerbesar {
			beratTerbesar = beratAnakKelinci[i]
		}
	}
	
	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar)
}