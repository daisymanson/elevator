package elevator

import (
	"fmt"
	"queue"
)

// Elevator
type Elevator struct {
	elevatorID         int
	currentFloorNumber int          // start floor is 0
	direction          int          // -1 == down, 1 == up
	goalFloorNumber    map[int]bool // stores the floors where the elevetor needs to stop
}

// Pickup
type PickupReq struct {
	pickupFloor int
	direction   int // -1 == down, 1 == up
}

// Control
type ElevatorControlSystem struct {
	elevator       []*Elevator
	pickupRequests *queue.Queue
}

// Constructor
func NewElevator(ID int) *Elevator {
	e := Elevator{elevatorID: ID, direction: 1}
	e.goalFloorNumber = make(map[int]bool)
	return &e
}

// Which elevator
func (e *Elevator) GetElevatorID() int {
	return e.elevatorID
}

// Where is elevator,currently
func (e *Elevator) GetCurrentFloorNumber() int {
	return e.currentFloorNumber
}

// The elevator is down or up
func (e *Elevator) GetDirection() int {
	return e.direction
}

// Where is the elevator to go
func (e *Elevator) GetNumGoalFloors() int {
	n := len(e.goalFloorNumber)
	return n
}
func (e *Elevator) GetGoalFloorNumbers() []int {
	goalFloors := make([]int, 0)

	for k, _ := range e.goalFloorNumber {
		goalFloors = append(goalFloors, k)
	}

	return goalFloors
}

// Add the floor to go
func (e *Elevator) addGoalFloor(floorNumber int) {
	e.goalFloorNumber[floorNumber] = true
}

// Cancel the floor to go
func (e *Elevator) removeFloor(floorNumber int) {
	delete(e.goalFloorNumber, floorNumber)
}

// Can move or not
func (e *Elevator) canMove() bool {
	if e.GetNumGoalFloors() > 0 {
		return true
	}
	return false
}
func (e *Elevator) canAddGoalFloor(goalFloorNumber int, direction int) bool {
	// if there are no goalFloors
	if e.GetNumGoalFloors() == 0 {
		e.direction = direction
		return true
		// if the move direction of the elevator is the same was requested
	} else if e.direction == direction {
		// if move up
		if direction > 0 && e.currentFloorNumber <= goalFloorNumber {
			return true
			// if move down
		} else if direction < 0 && e.currentFloorNumber >= goalFloorNumber {
			return true
		}
	}

	return false
}
func (e *Elevator) Update(currentFloorNum int, goalFloorNum int, direction int) bool {
	// update the goalFloorNumber map
	if e.canMove() {
		e.currentFloorNumber = currentFloorNum
		e.removeFloor(e.currentFloorNumber)
	}

	// if elevator moves in the same direction or the elevator is empty
	if e.canAddGoalFloor(goalFloorNum, direction) {
		e.addGoalFloor(goalFloorNum)
		return true
	}

	return false
}

func (e *Elevator) GetNextFloor() int {
	// move down
	if e.direction == -1 && e.currentFloorNumber > 0 {
		return e.currentFloorNumber - 1
	}
	// move up
	return e.currentFloorNumber + 1
}

func NewElevatorControlSystem(NumberOfElevators int) *ElevatorControlSystem {
	ecs := ElevatorControlSystem{}

	for i := 0; i < NumberOfElevators; i++ {
		ecs.elevator = append(ecs.elevator, NewElevator(i))
	}
	ecs.pickupRequests = queue.NewQueue()

	return &ecs
}

func (ecs *ElevatorControlSystem) Status() []string {
	seq := make([]string, 0)

	for _, elev := range ecs.elevator {
		seq = append(seq, fmt.Sprintf("elevatorID: %v, currentFloor: %v, direction: %v, goalFloors: %v", elev.GetElevatorID(), elev.GetCurrentFloorNumber(), elev.GetDirection(), elev.GetGoalFloorNumbers()))
	}

	return seq
}

func (ecs *ElevatorControlSystem) Pickup(pickupFloorNumber int, direction int) {
	ecs.pickupRequests.Push(PickupReq{pickupFloorNumber, direction})
}

func (ecs *ElevatorControlSystem) update(elev *Elevator, currentFloor int, goalFloor int, direction int) bool {
	return elev.Update(currentFloor, goalFloor, direction)
}

func (ecs *ElevatorControlSystem) Step() {
	for _, elev := range ecs.elevator {
		// check if there is a pickup request
		if ecs.pickupRequests.Len() > 0 {
			req := ecs.pickupRequests.Peek()

			if e, ok := req.(PickupReq); ok {
				// if the elevator moves in the same direction as the new pickup request
				success := ecs.update(elev, elev.GetNextFloor(), e.pickupFloor, e.direction)
				if success {
					_ = ecs.pickupRequests.Pop()
				}
			}
			// check if the goalFloorNumber map is not empty
			// if the goalFloorNumber map is not empty, the elevator needs to move to the next goal floor
		} else if elev.GetNumGoalFloors() > 0 {
			_ = ecs.update(elev, elev.GetNextFloor(), -1, elev.GetDirection())
		}
	}
}
