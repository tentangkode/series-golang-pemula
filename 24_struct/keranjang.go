package main

type item struct {
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
