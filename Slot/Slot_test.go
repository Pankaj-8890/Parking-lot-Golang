package slot

import (
	carPkg "parkinglot/Car"
	"testing"
	"reflect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// func TestNewSlot(t *testing.T){
//     slot := &Slot{}
//     assert.Equal(t,slot,&Slot{})
// }

func TestIsEmpty(t *testing.T) {
	s := NewSlot()
	assert.True(t, s.IsEmpty(), "Newly created slot should be empty")

	car := &carPkg.Car{VehicleNo: "abc1234",Color: "blue"}
	ticket,_ := s.ParkCar(car)
	assert.False(t, s.IsEmpty(), "Slot should not be empty after parking a car")

	s.UnparkCar(ticket)
	assert.True(t, s.IsEmpty(), "Slot should be empty after unparking the car")
}

func TestIsEmptyWhenSlotIsOccupied(t *testing.T){
	slot := NewSlot()
	car := &carPkg.Car{VehicleNo: "DL1099000",Color: "blue"}
    slot.Car = car
    if slot.IsEmpty() {
        t.Errorf("Expected slot to be occupied")
    }
}

func TestParkCar(t *testing.T) {
    slot := NewSlot()
    car := &carPkg.Car{VehicleNo: "MH0011990",Color: "white"}

    ticket, _ := slot.ParkCar(car)
    if !reflect.DeepEqual(slot.Car, car) {
        t.Errorf("Expected %v, got %v", slot.Car, car)
    }
    if ticket == "" {
        t.Errorf("Ticket not generated")
    }
}

func TestParkCarInOccupiedSlot(t *testing.T) {
    slot := NewSlot()
    car := &carPkg.Car{VehicleNo: "MH0011990",Color: "white"}

    ticket, _ := slot.ParkCar(car)
    if !reflect.DeepEqual(slot.Car, car) {
        t.Errorf("Expected %v, got %v", slot.Car, car)
    }
    if ticket == "" {
        t.Errorf("Ticket not generated")
    }

    _, err := slot.ParkCar(car)
    if err!=nil {
		assert.Equal(t,"Slot is already occupied",err.Error())
    }
}

func TestUnparkCar(t *testing.T) {
    slot := NewSlot()
    car := &carPkg.Car{VehicleNo: "MH0011990",Color: "white"}
    ticket,_ := slot.ParkCar(car)
    
    unparkedCar, _ := slot.UnparkCar(ticket)

    if !reflect.DeepEqual(unparkedCar, car) {
        t.Errorf("Expected %v, got %v", slot.Car, car)
    }

}

func TestUnparkCarFromEmptySlot(t *testing.T) {
    slot := NewSlot()
    car := &carPkg.Car{VehicleNo: "MH0011990",Color: "white"}
    ticket,_ := slot.ParkCar(car)
    
    unparkedCar, _ := slot.UnparkCar(ticket)

    if !reflect.DeepEqual(unparkedCar, car) {
        t.Errorf("Expected %v, got %v", slot.Car, car)
    }
    
    _, err := slot.UnparkCar(ticket)
	if err!=nil {
		assert.Equal(t,"Slot is empty",err.Error())
    }
}

func TestUnparkCarWithInvalidTicket(t *testing.T) {
    slot := NewSlot()
    car := &carPkg.Car{VehicleNo: "MH0011990",Color: "white"}
	ticket,_ := slot.ParkCar(car)
    invalid_ticket := uuid.New().String()
    
    _, err := slot.UnparkCar(invalid_ticket)

	if err!=nil || !reflect.DeepEqual(ticket, invalid_ticket){
		assert.Equal(t,"Invalid ticket",err.Error())
    }

}

func TestIsValidTicket(t *testing.T) {
	s := NewSlot()
	car := &carPkg.Car{VehicleNo: "ABC123",Color: "blue"}
	ticket, _ := s.ParkCar(car)

	assert.True(t, s.IsValidTicket(ticket), "Ticket should be valid for the parked car")
	assert.False(t, s.IsValidTicket("invalidTicket"), "Invalid ticket should be invalid")
}
