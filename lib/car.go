package parkinglot

type Car struct {
	RegistrationNumber string
	Color              string
}

func CreateCar(regNo, color string) Car {
	if color == "" {
		panic("Color cannot be an empty string")
	}
	if len(regNo) != 10 {
		panic("Length of registration number must be 10")
	}
	return Car{regNo, color}
}
