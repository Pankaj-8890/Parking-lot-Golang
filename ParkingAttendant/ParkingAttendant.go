package parkingattendant

import (
    "errors"
    "fmt"
	parkingLotPkg "parkinglot/ParkingLot"
	carPkg "parkinglot/Car"
)


type ParkingAttendant struct {
    parkingLots []*parkingLotPkg.ParkingLot
    statergy    Statergy
}

func NewParkingAttendant() *ParkingAttendant {
    return &ParkingAttendant{
        parkingLots: []*parkingLotPkg.ParkingLot{},
        statergy:    NEAREST,
    }
}

func (p *ParkingAttendant) Add(parkingLot *parkingLotPkg.ParkingLot) {
    p.parkingLots = append(p.parkingLots, parkingLot)
}

func (pa *ParkingAttendant) ChangeStrategy(strategy Statergy) {
    pa.statergy = strategy
}

func (pa *ParkingAttendant) ParkCar(car *carPkg.Car) (string, error) {
    
    parkingLotsTo := make([]*parkingLotPkg.ParkingLot, len(pa.parkingLots))
    copy(parkingLotsTo, pa.parkingLots)

	
    if pa.statergy == FARTHEST {
        for i, j := 0, len(parkingLotsTo)-1; i < j; i, j = i+1, j-1 {
            parkingLotsTo[i], parkingLotsTo[j] = parkingLotsTo[j], parkingLotsTo[i]
        }
    }

    for _, parkingLot := range parkingLotsTo {
        index, err := parkingLot.ParkCar(car)
		fmt.Println(index)
        if index != "" || err == nil {
            return index, nil
        }else {
			continue
		}
        
    }

    return "", errors.New("can't parked")
}

func (pa *ParkingAttendant) UnPark(id string) (*carPkg.Car, error) {
    for _, parkingLot := range pa.parkingLots {
        car, err := parkingLot.UnparkCar(id)
        if err == nil {
            return car, nil
        }else { 
			continue
		}

    }

    return &carPkg.Car{}, errors.New("Car not found")
}
