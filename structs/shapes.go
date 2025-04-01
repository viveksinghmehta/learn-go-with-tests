package structs

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width)
}

func Area(rect Rectangle) float64 {
	return rect.height * rect.width
}
