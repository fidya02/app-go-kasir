package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Barang adalah struktur data untuk menyimpan informasi barang.
type Barang struct {
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

// InvoiceItem adalah struktur data untuk menyimpan item dalam invoice.
type InvoiceItem struct {
	Nama    string `json:"nama"`
	Harga   int    `json:"harga"`
	Jumlah  int    `json:"jumlah"`
	Total   int    `json:"total"`
}

func main() {
	daftarBarang := []Barang{
		{"Barang A", 10000, 50},
		{"Barang B", 20000, 30},
		{"Barang C", 15000, 40},
	}

	var invoiceItems []InvoiceItem
	totalHarga := 0

	scanner := bufio.NewScanner(os.Stdin)

	for {
		tampilkanDaftarBarang(daftarBarang)

		fmt.Print("Pilih barang (nomor): ")
		scanner.Scan()
		pilihan, err := strconv.Atoi(scanner.Text())
		if err != nil || pilihan < 1 || pilihan > len(daftarBarang) {
			fmt.Println("Pilihan tidak valid. Silakan pilih nomor yang benar.")
			continue
		}

		fmt.Print("Jumlah yang dibeli: ")
		scanner.Scan()
		jumlah, err := strconv.Atoi(scanner.Text())
		if err != nil || jumlah <= 0 || jumlah > daftarBarang[pilihan-1].Stok {
			fmt.Println("Jumlah tidak valid atau stok tidak mencukupi.")
			continue
		}

		barang := daftarBarang[pilihan-1]
		totalHarga += barang.Harga * jumlah

		item := InvoiceItem{
			Nama:   barang.Nama,
			Harga:  barang.Harga,
			Jumlah: jumlah,
			Total:  barang.Harga * jumlah,
		}

		invoiceItems = append(invoiceItems, item)

		barang.Stok -= jumlah

		fmt.Print("Beli barang lain? (y/n): ")
		scanner.Scan()
		if scanner.Text() != "y" {
			break
		}
	}

	tampilkanInvoice(invoiceItems, totalHarga)

	err := simpanInvoiceKeFile(invoiceItems, totalHarga, "invoice.txt")
	if err != nil {
		fmt.Println("Gagal menyimpan invoice:", err)
	} else {
		fmt.Println("Invoice telah disimpan dalam file 'invoice.txt'")
	}
}

func tampilkanDaftarBarang(daftar []Barang) {
	fmt.Println("Daftar Barang:")
	for i, barang := range daftar {
		fmt.Printf("%d. %s - Harga: Rp%d - Stok: %d\n", i+1, barang.Nama, barang.Harga, barang.Stok)
	}
}

func tampilkanInvoice(items []InvoiceItem, total int) {
	fmt.Println("Invoice:")
	fmt.Println("===================================")
	for _, item := range items {
		fmt.Printf("%s x%d - Rp%d\n", item.Nama, item.Jumlah, item.Total)
	}
	fmt.Println("===================================")
	fmt.Printf("Total: Rp%d\n", total)
}

func simpanInvoiceKeFile(items []InvoiceItem, total int, filename string) error {
	data, err := json.MarshalIndent(map[string]interface{}{"items": items, "total": total}, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}
