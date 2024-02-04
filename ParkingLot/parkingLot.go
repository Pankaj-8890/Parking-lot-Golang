package parkinglot


import(
	slotPkg "parkinglot/Slot"
	"fmt"
	carPkg "parkinglot/Car"
    "reflect"
)

type ParkingLot struct {
   slots []*slotPkg.Slot
}

type Operation interface{
	ParkCar(*carPkg.Car) (string, error)
	UnparkCar(string) (*carPkg.Car, error)
}


func NewParkingLot(numberOfSlots int) (*ParkingLot, error) {
   if numberOfSlots <= 0 {
       return nil, fmt.Errorf("Slots can't be empty or negative")
   }

   slots := make([]*slotPkg.Slot, numberOfSlots)
   for i := range slots {
       slots[i] = slotPkg.NewSlot()
   }

   return &ParkingLot{slots: slots}, nil
}


func (p *ParkingLot) ParkCar(car *carPkg.Car) (string, error) {
	emptySlotIndex := p.findEmptySlot()
    if(p.isCarExist(car)){
        return "", fmt.Errorf("Can't par already exist with the same vechile number")
    }

	if emptySlotIndex == -1 {
		return "", fmt.Errorf("Parking lot is full")
	}
	ticket, err := p.slots[emptySlotIndex].ParkCar(car)
	if err != nil {
		return "", err
	}

	return ticket, nil
}

func (p *ParkingLot) UnparkCar(ticket string) (*carPkg.Car, error) {
   slot := p.getParkedCarSlot(ticket)
   if slot == nil {
       return nil, fmt.Errorf("Car not found")
   }

   car, err := slot.UnparkCar(ticket)
   if err != nil {
       return nil, err
   }

   return car, nil
}

func (p *ParkingLot) IsFull() bool {
   for _, slot := range p.slots {
       if slot.IsEmpty() {
           return false
       }
   }
   return true
}

func (p *ParkingLot) findEmptySlot() int {
   for i, slot := range p.slots {
       if slot.IsEmpty() {
           return i
       }
   }
   return -1
}

func (p *ParkingLot)isCarExist(car *carPkg.Car)bool{
    for _, slot := range p.slots {
        if slot.IsEmpty(){
            continue
        } 
        if  reflect.DeepEqual(car.VehicleNo, slot.Car.VehicleNo) {
            return true
        }
    }
    return false
}
func (p *ParkingLot) getParkedCarSlot(ticket string) *slotPkg.Slot {
   for _, slot := range p.slots {
       if slot.IsValidTicket(ticket) {
           return slot
       }
   }
   return nil
}
