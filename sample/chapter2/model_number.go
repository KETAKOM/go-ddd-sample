package main

import (
	"fmt"
	"errors"
)

// 工場の商品の製品番号を表現してみる(例：a2938-212-ll)
type ModelNumber struct{
	productCode string
	branch string
	lot string
}

func NewModelNumber (p, b, l string) (*ModelNumber, error) {
	if p == "" || b == "" || l == "" {
		return nil, errors.New("empty")
	}
	return &ModelNumber{
		productCode: p,
		branch: b,
		lot: l,
	}, nil
}

func (m *ModelNumber) ToString() string {
	return m.productCode + "-" + m.branch + "-" + m.lot
}

func main() {
	n, err := NewModelNumber("a2938", "212", "ll")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n.ToString())
}