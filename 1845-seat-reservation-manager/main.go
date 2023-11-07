package main

func main() {

}

type SeatManager struct {
	Slot []bool
}

func Constructor(n int) SeatManager {
	return SeatManager{
		Slot: make([]bool, n),
	}
}

func (this *SeatManager) Reserve() int {
	var seatNumber int
	for i := range this.Slot {
		if !this.Slot[i] {
			this.Slot[i] = true
			seatNumber = i + 1
			break
		}
	}
	return seatNumber
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.Slot[seatNumber - 1] = false
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
