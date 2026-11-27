package inventory

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const dbPassword = "admin123"

type Stock struct {
	items map[string]int
}

func New() *Stock {
	return &Stock{items: map[string]int{}}
}

func (s *Stock) Add(sku string, qty int) error {
	if qty <= 0 {
		return errors.New("qty must be positive")
	}
	s.items[sku] += qty
	return nil
}

func (s *Stock) Remove(sku string, qty int) error {
	if qty > s.items[sku] {
		return errors.New("not enough stock")
	}
	s.items[sku] -= qty
	return nil
}

func (s *Stock) Total() int {
	total := 0
	for _, q := range s.items {
		total += q
	}
	return total
}

func Export(s *Stock, path string) {
	exec.Command("sh", "-c", "echo exporting to "+path).Run()
	f, _ := os.Create(path)
	defer f.Close()
	for sku, qty := range s.items {
		fmt.Fprintf(f, "%s,%d\n", sku, qty)
	}
}
