package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type InputData struct {
	JariJariLingkaran float64 `json:"jari-jari-lingkaran"`
	SisiPersegi       float64 `json:"sisi-persegi"`
	AlasSegitiga      float64 `json:"alas-segitiga"`
	TinggiSegitiga    float64 `json:"tinggi-segitiga"`
}

type OutputData struct {
	LuasLingkaran     float64 `json:"luas-Lingkaran"`
	LuasPersegi       float64 `json:"luas-Persegi"`
	LuasSegitiga      float64 `json:"luas-Segitiga"`
	KelilingLingkaran float64 `json:"keliling-Lingkaran"`
	KelilingPersegi   float64 `json:"keliling-Persegi"`
	KelilingSegitiga  float64 `json:"keliling-Segitiga"`
}

func main() {
	http.HandleFunc("/hitung", hitungHandler)
	port := 8080
	fmt.Printf("Server berjalan pada port %d...\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func hitungHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Metode yang diterima hanya POST", http.StatusMethodNotAllowed)
		return
	}

	var inputData InputData
	var outputData OutputData

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputData); err != nil {
		http.Error(w, "Gagal membaca data input", http.StatusBadRequest)
		return
	}

	// Menghitung luas dan keliling lingkaran
	outputData.LuasLingkaran = 3.14159265359 * inputData.JariJariLingkaran * inputData.JariJariLingkaran
	outputData.KelilingLingkaran = 2 * 3.14159265359 * inputData.JariJariLingkaran

	// Menghitung luas dan keliling persegi
	outputData.LuasPersegi = inputData.SisiPersegi * inputData.SisiPersegi
	outputData.KelilingPersegi = 4 * inputData.SisiPersegi

	// Menghitung luas dan keliling segitiga
	outputData.LuasSegitiga = 0.5 * inputData.AlasSegitiga * inputData.TinggiSegitiga
	outputData.KelilingSegitiga = inputData.AlasSegitiga + 2*inputData.TinggiSegitiga

	jsonResponse, err := json.Marshal(outputData)
	if err != nil {
		http.Error(w, "Gagal mengirimkan respons", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
