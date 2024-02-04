package parkingattendant

import (
	"testing"
	"reflect"
	"errors"
	carPkg "parkinglot/Car"
	parkingLotPkg "parkinglot/ParkingLot"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


func TestNewParkingAttendant(t *testing.T) {
    pa := NewParkingAttendant()
    assert.Equal(t, NEAREST, pa.statergy)
    assert.Empty(t, pa.parkingLots)
}

func TestAddParkingLot(t *testing.T) {
    pa := NewParkingAttendant()
    parkingLot,_ := parkingLotPkg.NewParkingLot(1)
    pa.Add(parkingLot)
    assert.Equal(t, []*parkingLotPkg.ParkingLot{parkingLot}, pa.parkingLots)
}

func TestChangeStrategy(t *testing.T) {
    pa := NewParkingAttendant()
    pa.ChangeStrategy(FARTHEST)
    assert.Equal(t, FARTHEST, pa.statergy)
}

func TestPark2CarThroughParkingLotsWhenThere_is_2SlotsAvailable(t *testing.T) {

	parkingLot1,err := parkingLotPkg.NewParkingLot(1)
	if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }
	parkingLot2,err := parkingLotPkg.NewParkingLot(1)
	if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }
	pa := NewParkingAttendant()
	pa.Add(parkingLot1)
	pa.Add(parkingLot2)

	carA := carPkg.NewCar("ABC12312", "Red")
	carB := carPkg.NewCar("ABC12311", "Blue")

	ticketA, err1 := pa.ParkCar(carA)
	ticketB, err2 := pa.ParkCar(carB)
	if  err1 != nil || err2 != nil || ticketA == "" || ticketB == "" {
		wantErr := errors.New("Parking lot is full")
        assert.Equal(t,wantErr,err1)
		return
    }

}

func TestPark3CarThroughParkingLotsWhenThere_is_2SlotsAvailable(t *testing.T) {

	pa := NewParkingAttendant()
	parkingLot1,err := parkingLotPkg.NewParkingLot(1)
	if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }
	pa.Add(parkingLot1)
	parkingLot2,err := parkingLotPkg.NewParkingLot(1)
	if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }
	pa.Add(parkingLot2)

	carA := carPkg.NewCar("ABC12312", "Red")
	carB := carPkg.NewCar("ABC12311", "Blue")
	carC := carPkg.NewCar("ABC12311", "Blue")

	pa.ParkCar(carA)
	pa.ParkCar(carB)
	ticketC,err := pa.ParkCar(carC)

	if  err != nil || ticketC != ""{
		wantErr := errors.New("can't parked")
		assert.Equal(t,wantErr,err)
		return
    }

}
func TestUnParkCarThroughParkingLotsWithValidCarTicket(t *testing.T) {
    parkingAttendant := NewParkingAttendant()
    parkingLot, err := parkingLotPkg.NewParkingLot(1)
    if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }
    parkingAttendant.Add(parkingLot)
    car := carPkg.NewCar("xyz01", "blue")
    ticket, err := parkingLot.ParkCar(car)
    if err != nil {
        t.Fatal("Error parking car:", err)
    }

    getCar,err := parkingAttendant.UnPark(ticket)
    if err != nil {
		wantErr := errors.New("Car not found")
		assert.Equal(t,wantErr,err)
		return
    }

    if !reflect.DeepEqual(car, getCar) {
        t.Errorf("Expected %v, got %v", car, getCar)
    }
}

func TestUnParkCarThroughParkingLotsWithInvalidCarTicket(t *testing.T) {
    parkingAttendant := NewParkingAttendant()
    parkingLot, err := parkingLotPkg.NewParkingLot(1)
    if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }
    parkingAttendant.Add(parkingLot)
    car := carPkg.NewCar("xyz01", "blue")
    ticket, err := parkingLot.ParkCar(car)
    if err != nil || ticket == "" {
		wantErr := errors.New("Parking lot is full can't park ")
        assert.Equal(t,wantErr,err)
		return
    }

	invalidTicket := uuid.New().String()
    getCar,err := parkingAttendant.UnPark(invalidTicket)
    if err != nil {
		wantErr := errors.New("Car not found")
		assert.Equal(t,wantErr,err)
		return
    }

    if !reflect.DeepEqual(car, getCar) {
        t.Errorf("Expected %v, got %v", car, getCar)
    }
}

func TestParkCarThroughTwoParkingAttendants(t *testing.T) {
    parkingAttendant := NewParkingAttendant()
    secondParkingAttendant := NewParkingAttendant()
    parkingLot, err := parkingLotPkg.NewParkingLot(2)
    if err != nil {
		wantErr := errors.New("Slots can't be empty or negative")
        assert.Equal(t,wantErr,err)
		return
    }
    parkingAttendant.Add(parkingLot)
    secondParkingAttendant.Add(parkingLot)

    carA := carPkg.NewCar("xyz01", "blue")
    carB := carPkg.NewCar("xyz01", "blue")
    parkingAttendant.ParkCar(carA)
    secondParkingAttendant.ParkCar(carB)

    assert.True(t, parkingLot.IsFull(), "Expected parking lot to be full")
}

func TestParkCarThroughTwoParkingAttendantsWhenParkingLotIsFull(t *testing.T) {
    parkingAttendant := NewParkingAttendant()
    secondParkingAttendant := NewParkingAttendant()
    parkingLot, _ := parkingLotPkg.NewParkingLot(1)
    parkingAttendant.Add(parkingLot)
    secondParkingAttendant.Add(parkingLot)

    car := carPkg.NewCar("xyz01", "blue")
    parkingAttendant.ParkCar(car)
    
    _, err := secondParkingAttendant.ParkCar(car)
    assert.Error(t, err, "Expected an error when parking lot is full")
}

func TestParkCarFromOneParkingAttendantAndUnParkCarFromSecondParkingAttendant(t *testing.T) {
    parkingAttendant := NewParkingAttendant()
    secondParkingAttendant := NewParkingAttendant()
    parkingLot, _ := parkingLotPkg.NewParkingLot(1)
    parkingAttendant.Add(parkingLot)
    secondParkingAttendant.Add(parkingLot)

    carA := carPkg.NewCar("xyz01", "blue")
    ticket, err := parkingAttendant.ParkCar(carA)
    if err != nil {
        t.Fatal("Error parking car:", err)
    }

    getCar, err := secondParkingAttendant.UnPark(ticket)
    if err != nil {
        t.Fatal("Error unparking car:", err)
    }

    assert.Equal(t, carA, getCar)
}

func TestParkCarOnTheFarthestParkingLot(t *testing.T) {
    parkingAttendant := NewParkingAttendant()
    parkingLot,_ := parkingLotPkg.NewParkingLot(1)
    anotherParkingLot,_ := parkingLotPkg.NewParkingLot(3)
    parkingAttendant.ChangeStrategy(FARTHEST)
    parkingAttendant.Add(parkingLot)
    parkingAttendant.Add(anotherParkingLot)

    carA := carPkg.NewCar("xyz01", "blue")
    carB := carPkg.NewCar("xyz02", "red")
    carC := carPkg.NewCar("xyz03", "green")

    parkingAttendant.ParkCar(carA)
    parkingAttendant.ParkCar(carB)
    parkingAttendant.ParkCar(carC)

    assert.False(t, parkingLot.IsFull(), "Expected parking lot not to be full")
    assert.True(t, anotherParkingLot.IsFull(), "Expected another parking lot to be full")
}


