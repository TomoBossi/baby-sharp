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

func (m matrix2) closest(row, col, pad int) float64 {
	mRows := m.rows()
	mCols := m.cols()
	if row < pad {
		if col < pad {
			return m[0][0]
		} else if col >= pad+mCols {
			return m[0][mCols-1]
		} else {
			return m[0][col-pad]
		}
	} else if row >= pad+mRows {
		if col < pad {
			return m[mRows-1][0]
		} else if col >= pad+mCols {
			return m[mRows-1][mCols-1]
		} else {
			return m[mRows-1][col-pad]
		}
	} else if col < pad {
		return m[row-pad][0]
	} else {
		return m[row-pad][mCols-1]
	}
}

func (m matrix2) extend(pad int) matrix2 {
	var extended matrix2
	mRows := m.rows()
	mCols := m.cols()
	for i := range mRows + 2*pad {
		extended = append(extended, []float64{})
		for j := range mCols + 2*pad {
			if i < pad || i >= pad+mRows || j < pad || j >= pad+mCols {
				extended[i] = append(extended[i], m.closest(i, j, pad))
			}
		}
	}
	return extended
}

func (m matrix2) convolveExtended(kernel matrix2) matrix2 {
	pad := kernel.rows() / 2
	extended := m.extend(pad)
	return extended.convolve(kernel)
}
