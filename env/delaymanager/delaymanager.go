package delaymanager

import (
	"fmt"
	"sync"
)

//Delay Manager keeps track of creation and delivery time of each packet
var DelayManager = delaymanager{
	creationTime: make(map[int]int64),
	reachTime: make(map[int]int64),
	delivered:    make(map[int]bool),
}

type delaymanager struct {
	creationTime map[int]int64
	reachTime map[int]int64
	delivered    map[int]bool
	mux          sync.Mutex
}

func Erase() {
	for key, _ := range DelayManager.creationTime {
		delete(DelayManager.creationTime, key)
	}

	for key, _ := range DelayManager.reachTime {
		delete(DelayManager.reachTime, key)
	}

	for key, _ := range DelayManager.delivered {
		delete(DelayManager.delivered, key)
	}
}

// Register the packet for deliveries
func Register(packetID int) {
	DelayManager.mux.Lock()
	defer DelayManager.mux.Unlock()

	DelayManager.reachTime[packetID] = 0
	DelayManager.creationTime[packetID] = 0
}

// Initialises Creation Time of a Packet
func UpdateCreationTime(packetID int, CT int64) {
	DelayManager.mux.Lock()
	defer DelayManager.mux.Unlock()

	DelayManager.creationTime[packetID] = CT
}

// MarkSuccess if the packet have not been delivered before and Updates Reach Time
func MarkSuccess(packetID int, RT int64) {
	DelayManager.mux.Lock()
	defer DelayManager.mux.Unlock()

	if !DelayManager.delivered[packetID] {
		DelayManager.delivered[packetID] = true
		DelayManager.reachTime[packetID] = RT
	}
}

// String returns the DelayManager data of delivered packets of bicycles in string form
// e.g. fmt.Fprint(os.Create("result.txt"), DelayManager)
func (m delaymanager) String() string {
	result := "packetid,creationtime,deliverytime\n"
	for packet, creationTime := range m.creationTime {
		if (m.delivered[packet] == true) {
			result += fmt.Sprintf("%d,%d,%d\n", packet, creationTime, m.reachTime[packet])
		}
	}
	return result
}
