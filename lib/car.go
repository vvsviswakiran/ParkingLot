package parkinglot

type Car struct {
	RegistrationNumber string
	Color              string
}

func CreateCar(regNo, color string) {
	if color == "" {
		panic("Color cannot be an empty string")
	}
}
