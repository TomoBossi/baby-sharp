package main

type matrix2 [][]float64

func (m matrix2) rows() int {
	return len(m)
}

func (m matrix2) cols() int {
	return len(m[0])
}

func (m matrix2) apply(kernel matrix2) float64 {
	sum := 0.0
	for i := range kernel.rows() {
		for j := range kernel.cols() {
			sum += m[i][j] * kernel[i][j]
		}
	}
	return sum
}

func (m matrix2) slice(rowStart, rowEnd, colStart, colEnd int) matrix2 {
	var sliced matrix2
	for i := range rowEnd - rowStart {
		sliced = append(sliced, m[rowStart+i][colStart:colEnd])
	}
	return sliced
}

func (m matrix2) convolve(kernel matrix2) matrix2 {
	var convolved matrix2
	kRows := kernel.rows()
	kCols := kernel.cols()
	for i := range m.rows() - kRows + 1 {
		convolved = append(convolved, []float64{})
		for j := range m.cols() - kernel.cols() + 1 {
			convolved[i] = append(convolved[i], m.slice(i, i+kRows, j, j+kCols).apply(kernel))
		}
	}
	return convolved
}

func (m matrix2) closest(row, col, padRows, padCols int) float64 {
	mRows := m.rows()
	mCols := m.cols()
	if row < padRows {
		if col < padCols {
			return m[0][0]
		} else if col >= padCols+mCols {
			return m[0][mCols-1]
		} else {
			return m[0][col-padCols]
		}
	} else if row >= padRows+mRows {
		if col < padCols {
			return m[mRows-1][0]
		} else if col >= padCols+mCols {
			return m[mRows-1][mCols-1]
		} else {
			return m[mRows-1][col-padCols]
		}
	} else if col < padCols {
		return m[row-padRows][0]
	} else {
		return m[row-padRows][mCols-1]
	}
}

func (m matrix2) extend(padRows, padCols int) matrix2 {
	var extended matrix2
	mRows := m.rows()
	mCols := m.cols()
	for i := range mRows + 2*padRows {
		extended = append(extended, []float64{})
		for j := range mCols + 2*padCols {
			if i < padRows || i >= padRows+mRows || j < padCols || j >= padCols+mCols {
				extended[i] = append(extended[i], m.closest(i, j, padRows, padCols))
			} else {
				extended[i] = append(extended[i], m[i-padRows][j-padCols])
			}
		}
	}
	return extended
}

func (m matrix2) convolveExtended(kernel matrix2) matrix2 {
	padRows := kernel.rows() / 2
	padCols := kernel.cols() / 2
	extended := m.extend(padRows, padCols)
	return extended.convolve(kernel)
}
