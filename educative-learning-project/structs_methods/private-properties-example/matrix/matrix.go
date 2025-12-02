package matrix

// This data structure is not export-able bcz it's name is in small letters
type matrix struct {
	row int
	col int
}

// Factory Method/Constructor is to be made Public to access in main and initiliaze the matrix
func NewMatrix(row int, col int) *matrix {
	if (row <= 0) || (col <= 0) {
		return nil
	}
	return &matrix{row, col}
}
