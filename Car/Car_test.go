package car

import (
	"testing"
)

func TestNewParkingLot(t *testing.T) {
    
	car := NewCar("KA10011","black")

	if(car.VehicleNo != "KA10011" || car.Color != "black"){
		t.Errorf("Vehicle no. or color are not match")
	}
}