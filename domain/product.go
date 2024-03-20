package domain

type product struct {
	title               string
	linkProduct         string
	linkImage           string
	rating              float32
	stock               int8
	price               float64
	priceBeforeDiscount float64
	discountPercentage  int8
}
