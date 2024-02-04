package parkinglot

import (
	"errors"
	carPkg "parkinglot/Car"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"

)

func TestNewParkingLot(t *testing.T) {
    lot, err := NewParkingLot(5)
    if err != nil {
        t.Errorf("Failed to create parking lot: %v", err)
    }
    if len(lot.slots) != 5 {
        t.Errorf("Expected 5 slots, got %d", len(lot.slots))
    }

    _, err = NewParkingLot(0)
    if err == nil {
        t.Errorf("Expected error for invalid slot number")
    }
}


func TestParkCar(t *testing.T) {
    lot, err := NewParkingLot(0)
	
    if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }

    car := carPkg.NewCar("ABC123", "Red")
    lot.ParkCar(car)
    if err != nil {
        t.Errorf("Failed to park car: %v", err)
    }

}
func TestParkCarWhenLotisFull(t *testing.T) {
    lot, err := NewParkingLot(1)
    if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }

    carA := carPkg.NewCar("ABC123", "Red")
	carB := carPkg.NewCar("ABC123", "Red")
    lot.ParkCar(carA)
	ticket,err := lot.ParkCar(carB)
    if err != nil || ticket != "" {
		wantErr := errors.New("Can't par already exist with the same vechile number")
        assert.Equal(t,wantErr,err)
		return
    }


}

func TestParkCarTheSameCarAgain(t *testing.T) {
    lot, err := NewParkingLot(2)
    if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }

    carA := carPkg.NewCar("ABC123", "Red")
	carB := carPkg.NewCar("ABC123", "Red")
    lot.ParkCar(carA)
	ticket,err := lot.ParkCar(carB)
    if err != nil || ticket == "" {
		wantErr := errors.New("Can't par already exist with the same vechile number")
        assert.Equal(t,wantErr,err)
		return
    }


}

func TestUnParkCar(t *testing.T) {
    lot, err := NewParkingLot(1)
    if err != nil {
        t.Fatal(err)
    }

    car := carPkg.NewCar("ABC123", "Red")

    ticket, err := lot.ParkCar(car)
    if err != nil {
		wantErr := errors.New("Parking lot is full can't park "+ ticket)
        assert.Equal(t,wantErr,err)
		return
    }

    getCar, err := lot.UnparkCar(ticket)
    if err != nil {
		wantErr := errors.New("Car not found")
		assert.Equal(t,wantErr,err)
		return
    }

    if !reflect.DeepEqual(car, getCar) {
        t.Errorf("Expected %v, got %v", car, getCar)
    }
}

func TestUnParkCarWithInvalidTicket(t *testing.T) {
    lot, err := NewParkingLot(1)
    if err != nil {
        t.Fatal(err)
    }

    car := carPkg.NewCar("ABC123", "Red")

    ticket, err := lot.ParkCar(car)
    if err != nil {
		wantErr := errors.New("Parking lot is full can't park "+ ticket)
        assert.Equal(t,wantErr,err)
		return
    }

	invalidTicket := uuid.New().String()
    getCar, err := lot.UnparkCar(invalidTicket)
    if err != nil {
		wantErr := errors.New("Car not found")
		assert.Equal(t,wantErr,err)
		return
    }

    if !reflect.DeepEqual(car, getCar) {
        t.Errorf("Expected %v, got %v", car, getCar)
    }
}



