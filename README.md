# Elevator
* ## 描述
所需物件
* Elevator
```go
type Elevator struct {
	elevatorID         int
	currentFloorNumber int          // start floor is 0
	direction          int          // -1 == down, 1 == up
	goalFloorNumber    map[int]bool // stores the floors where the elevetor needs to stop
}
// Constructor
func NewElevator(ID int) *Elevator{

}

// Which elevator
func (e *Elevator)GetElevatorID() int{

}

// Where is elevator,currently
func (e *Elevator)GetCurrentFloorNumber() int{

}

// The elevator is down or up
func (e *Elevator)GetDirection() int{

}

// Where is the elevator to go
func (e *Elevator) GetNumGoalFloors() int {

}

func (e *Elevator) GetGoalFloorNumbers() []int {

}

// Add floor to go
func (e *Elevator) addGoalFloor(floorNumber int){
    e.goalFloorNumber[floorNumber] = true
}
// Cancel floor to go
func (e *Elevator) removeFloor(floorNumber int){
    delete(e.goalFloor, floorNumber)
}


```

