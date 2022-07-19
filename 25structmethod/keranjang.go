package main

import (
	"fmt"
	"strings"
)

type item struct {
	id        int
	nama_item string
	qty       int32
	harga     float64
}

type keranjang struct {
	nama  string
	items []item
	total float64
}

func newKeranjang(nama string) keranjang {
	k := keranjang{
		nama:  nama,
		items: []item{},
		total: 0,
	}

	return k
}

func (k *keranjang) tambahItem(belanjaan item) {
	k.items = append(k.items, belanjaan)
	k.hitungTotal()
}

func (k *keranjang) tambahItemDua(belanjaan ...item) {
	k.items = append(k.items, belanjaan...)
	k.hitungTotal()
}

func (k *keranjang) updateItem(itemUpdate item) {
	for index, val := range k.items {
		if val.id == itemUpdate.id {
			k.items[index] = itemUpdate
		}
	}
	k.hitungTotal()
}

func (k *keranjang) deleteItem(id int) {
	for index, val := range k.items {
		if val.id == id {
			k.items = append(k.items[:index], k.items[index+1:]...)
		}
	}
	k.hitungTotal()
}

func (k *keranjang) hitungTotal() {
	var total float64
	for _, val := range k.items {
		total += val.harga * float64(val.qty)
	}

	k.total = total
}

func (k *keranjang) formatOutput() string {

	print := strings.Repeat("#", 65) + "\n"
	print += fmt.Sprintf("%40s\n", k.nama)
	print += strings.Repeat("#", 65) + "\n"

	print += fmt.Sprintf("%-30s %-5s %-15s %-15s \n", "NAMA BARANG", "QTY", "HARGA", "SUBTOTAL")
	print += strings.Repeat("-", 65) + "\n"

	for _, item := range k.items {
		print += fmt.Sprintf("%-30s %-5d %-15.2f %-15.2f \n", item.nama_item, item.qty, item.harga, item.harga*float64(item.qty))
	}

	print += strings.Repeat("-", 65) + "\n"
	print += fmt.Sprintf("%-30s %-5s %-15s %-15.2f \n", "", "", "TOTAL BELANJA", k.total)
	return print
}
