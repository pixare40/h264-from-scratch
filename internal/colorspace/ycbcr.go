package colorspace

func YCbCrToRGB(y, cb, cr float64) (r, g, b float64) {
	r = y + 1.402*(cr-0.5)
	g = y - 0.34414*(cb-0.5) - 0.71414*(cr-0.5)
	b = y + 1.772*(cb-0.5)
	return
}

func RGBToYCbCr(r, g, b float64) (y, cb, cr float64) {
	y = 0.299*r + 0.587*g + 0.114*b
	cb = -0.1687*r - 0.3313*g + 0.5*b + 0.5
	cr = 0.5*r - 0.4187*g - 0.0813*b + 0.5
	return
}
