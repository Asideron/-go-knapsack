package main

type PackInfo struct {
	Value int
	Items *[]*Item
}

func newPackInfo() *PackInfo {
	return &PackInfo{
		Value: 0,
		Items: &[]*Item{},
	}
}

func (p *PackInfo) getItems() *[]*Item {
	items := make([]*Item, 0)
	items = append(items, *p.Items...)
	return &items
}
