package internel

type Map struct {
	coordinates [10][10]string
}

func (m *Map) Generate(xLength, yLength int) {

	for y := 0; y < yLength; y++ {
		for x := 0; x < xLength; x++ {
			m.coordinates[y][x] =  "#"
		}
	}
}
