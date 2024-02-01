package car

type Car struct {
    VehicleNo string
    Color     string
}

func NewCar(vehicleNo, color string) *Car {
    return &Car{VehicleNo: vehicleNo, Color: color}
}
