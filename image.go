package main

type image struct {
	r, g, b matrix2
}

func (i image) convolveExtended(kernel matrix2) image {
	return image{
		r: i.r.convolveExtended(kernel),
		g: i.g.convolveExtended(kernel),
		b: i.b.convolveExtended(kernel),
	}
}
