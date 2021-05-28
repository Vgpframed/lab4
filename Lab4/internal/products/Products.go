package products

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Product struct {
	Flag  string
	Name  string
	Type  string
	Rate  int
	Price int
}

func (p *Product) GetProducts(filename string) []Product {

	var inc = 0
	cnt := p.GetCount(filename)
	fmt.Println(cnt)
	Products := make([]Product, cnt)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'

	for {
		record, e := reader.Read()
		if e == io.EOF {
			break
		}
		p.Flag = record[0]
		p.Name = record[1]
		p.Type = record[2]
		p.Rate, _ = strconv.Atoi(record[3])
		p.Price, _ = strconv.Atoi(record[4])
		Products[inc] = *p
		inc++
	}
	fmt.Println(Products)
	return Products
}

func (p *Product) GetCount(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	var inc = 1
	for {
		_, e := reader.Read()
		if e == io.EOF {
			return inc - 1
		}

		inc++
	}
}

