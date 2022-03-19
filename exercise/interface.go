package main

import "fmt"

type bangunDatar interface {
	luas() float64
	jenisBangunDatar() string
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) jenisBangunDatar() string {
	return "persegi"
}

type perhitunganLuasBangunDatar struct {
	kumpulanBangunDatar []bangunDatar
}

func (p *perhitunganLuasBangunDatar) tambahBangunDatar(b bangunDatar) {
	p.kumpulanBangunDatar = append(p.kumpulanBangunDatar, b)
}
func (p *perhitunganLuasBangunDatar) bangunDatarTerdaftar() []bangunDatar {
	return p.kumpulanBangunDatar
}

func main() {
	persegi := persegi{sisi: 4}
	bangunDatarService := perhitunganLuasBangunDatar{}
	bangunDatarService.tambahBangunDatar(persegi)
	bangunDatar := bangunDatarService.bangunDatarTerdaftar()
	for _, bangun := range bangunDatar {
		cetak(bangun)
	}
}

func cetak(b bangunDatar) {
	fmt.Println(b.jenisBangunDatar())
}
