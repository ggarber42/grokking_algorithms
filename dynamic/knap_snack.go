package dynamic

type Item struct {
	name   string
	weight int
	value  int
}

func KnapSnackSolve(items []Item, maxWeight int) []Item {
	maxrows := len(items)
	maxcolumns := maxWeight

	grid := make([][][]Item, maxrows)
	fillZeros(&grid, maxcolumns)

	for row, column := range grid {
		rowItem := items[row]

		for index := range column {
			var newItens []Item

			availableWeight := index + 1
			items := getOldItens(grid, index, row)
			items = append(items, rowItem)
			sortedItems := sortByValue(items)

			for _, item := range sortedItems {
				if item.weight <= availableWeight {
					newItens = append(newItens, item)
					availableWeight -= item.weight
				}
			}
			grid[row][index] = append(grid[row][index], newItens...)
		}
	}

	return grid[maxrows-1][maxcolumns-1]
}

func getOldItens(grid [][][]Item, index, row int) []Item {
	var items []Item
	covered := make(map[string]bool)

	for row > 0 {
		for _, gridItems := range grid[row-1] {
			for _, item := range gridItems {
				if !covered[item.name] {
					covered[item.name] = true
					items = append(items, item)
				}
			}
		}

		row -= 1
	}

	return items
}

func fillZeros(grid *[][][]Item, columns int) {
	for i := range *grid {
		(*grid)[i] = make([][]Item, columns)
	}
}

func sortByValue(items []Item) []Item {
	var newItens []Item

	for len(items) > 0 {
		var newItem Item
		var index int

		for i, item := range items {
			if i == 0 {
				newItem = item
				index = i
			} else if newItem.value < item.value {
				newItem = item
				index = i
			}

		}
		newItens = append(newItens, newItem)
		items = append(items[:index], items[index+1:]...)

	}

	return newItens
}
