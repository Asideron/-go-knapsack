package main

type matrix []row

type row []*PackInfo

func newRow(len int) row {
	r := make(row, len)
	for i := range r {
		r[i] = newPackInfo()
	}
	return r
}
