package slot

import (
   "fmt"
   carPkg "parkinglot/Car"
   "github.com/google/uuid"
)

type Slot struct {
   Car    *carPkg.Car
   Ticket string
}

type Operation interface{
	ParkCar(*carPkg.Car) (string, error)
	UnparkCar(string) (*carPkg.Car, error)
}

func NewSlot() *Slot {
   return &Slot{}
}

func (s *Slot) IsEmpty() bool {
   return s.Car == nil
}

func (s *Slot) ParkCar(car *carPkg.Car) (string, error) {
	if !s.IsEmpty(){
		return "", fmt.Errorf("Slot is already occupied")
	}
	s.Car = car
	s.Ticket = uuid.New().String()
	return s.Ticket, nil
}

func (s *Slot) UnparkCar(ticket string) (*carPkg.Car, error) {
   if s.IsEmpty() {
       return nil, fmt.Errorf("Slot is empty")
   }

   if s.Ticket != ticket {
       return nil, fmt.Errorf("Invalid ticket")
   }

   parkedCar := s.Car
   s.Car = nil
   s.Ticket = ""
   return parkedCar, nil
}

func (s *Slot) IsValidTicket(ticket string) bool {
   return s.Ticket == ticket
}
