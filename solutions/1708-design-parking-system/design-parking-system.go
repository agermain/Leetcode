// Design a parking system for a parking lot. The parking lot has three kinds of parking spaces: big, medium, and small, with a fixed number of slots for each size.
//
// Implement the ParkingSystem class:
//
//
// 	ParkingSystem(int big, int medium, int small) Initializes object of the ParkingSystem class. The number of slots for each parking space are given as part of the constructor.
// 	bool addCar(int carType) Checks whether there is a parking space of carType for the car that wants to get into the parking lot. carType can be of three kinds: big, medium, or small, which are represented by 1, 2, and 3 respectively. A car can only park in a parking space of its carType. If there is no space available, return false, else park the car in that size space and return true.
//
//
//  
// Example 1:
//
//
// Input
// ["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
// [[1, 1, 0], [1], [2], [3], [1]]
// Output
// [null, true, true, false, false]
//
// Explanation
// ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
// parkingSystem.addCar(1); // return true because there is 1 available slot for a big car
// parkingSystem.addCar(2); // return true because there is 1 available slot for a medium car
// parkingSystem.addCar(3); // return false because there is no available slot for a small car
// parkingSystem.addCar(1); // return false because there is no available slot for a big car. It is already occupied.
//
//
//  
// Constraints:
//
//
// 	0 <= big, medium, small <= 1000
// 	carType is 1, 2, or 3
// 	At most 1000 calls will be made to addCar
//
//


type ParkingSystem struct {
    bigCar int
    mediumCar int
    smallCar int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{big, medium, small}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
    case 1:
        if this.bigCar > 0 {
            this.bigCar--
            return true
        }
    case 2:
        if this.mediumCar > 0 {
            this.mediumCar--
            return true
        }
    case 3:
        if this.smallCar > 0 {
            this.smallCar--
            return true
        }
    }
    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
