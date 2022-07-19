package main

import "fmt"

func main() {

	keranjangBelanja := newKeranjang("Keranjang Belanja CPM")

	item1 := item{
		id:        5,
		nama_item: "Kemeja Flanel",
		qty:       2,
		harga:     120000,
	}

	keranjangBelanja.tambahItem(item1)

	items := []item{
		{
			id:        1,
			nama_item: "Kemeja Polos",
			qty:       2,
			harga:     100000,
		},
		{
			id:        2,
			nama_item: "Sepatu",
			qty:       1,
			harga:     300000,
		},
		{
			id:        3,
			nama_item: "Tas",
			qty:       1,
			harga:     400000,
		},
		{
			id:        4,
			nama_item: "Rok",
			qty:       1,
			harga:     100000,
		},
	}

	keranjangBelanja.tambahItemDua(items...)

	updateItem := item{
		id:        2,
		nama_item: "Sepatu Edited",
		qty:       4,
		harga:     100000,
	}
	keranjangBelanja.updateItem(updateItem)
	//fmt.Println(keranjangBelanja.formatOutput())
	//keranjangBelanja.deleteItem(1)
	fmt.Println(keranjangBelanja.formatOutput())

}
