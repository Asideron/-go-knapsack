package main

type Backpack struct {
	Items  *[]*Item
	Cap    int
	matrix *matrix
}

func NewBackpack(cap int) *Backpack {
	items := make([]*Item, 0)
	matrix := make(matrix, 0)
	return &Backpack{
		Items:  &items,
		Cap:    cap,
		matrix: &matrix,
	}
}

func (b *Backpack) Add(item Item) {
	*b.Items = append(*b.Items, &item)
}

func (b *Backpack) canFitItem(i, j int) bool {
	return j-(*b.Items)[i].Weight >= 0
}

func (b *Backpack) packToCombineWith(i, j int) *PackInfo {
	return (*b.matrix)[i][j-(*b.Items)[i].Weight]
}

func (b *Backpack) itemsCombinationValue(i, j int) int {
	return b.packToCombineWith(i, j).Value + (*b.Items)[i].Value
}

func (b *Backpack) itemsCombinationLessValuable(i, j int) bool {
	return (*b.matrix)[i][j].Value >= b.itemsCombinationValue(i, j)
}

func (b *Backpack) Pack() (*PackInfo, error) {
	if b.Items == nil || len(*b.Items) == 0 {
		return nil, errNoItemsProvided
	}

	*b.matrix = append(*b.matrix, newRow(b.Cap+1))

	for i := range *b.Items {
		r := newRow(1)
		for j := 1; j < b.Cap+1; j++ {
			var element *PackInfo
			if !b.canFitItem(i, j) || b.itemsCombinationLessValuable(i, j) {
				element = (*b.matrix)[i][j]
			} else {
				element = newPackInfo()

				items := b.packToCombineWith(i, j).getItems()
				*items = append(*items, (*b.Items)[i])

				element.Items = items
				element.Value = b.itemsCombinationValue(i, j)
			}
			r = append(r, element)
		}
		*b.matrix = append(*b.matrix, r)
	}

	return (*b.matrix)[len(*b.Items)][b.Cap], nil
}
