package internel

type Map struct {
	coordinates [10][10]string
}

func (m *Map) Generate(xLength, yLength int) {

	for y := 0; y < yLength; y++ {
		var xString = ""
		for x := 0; x < xLength; x++ {
			xString += "#"
		}
		m.coordinates[y][0] = xString
	}
}
