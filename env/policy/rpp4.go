package policy

import (
	"iob-dtn/env/sensor/buffer"
	"iob-dtn/env/sensor/buffer/packet"
)

// RPP-4 Received Packet Priority
// when a packet is received, it replaces the newest generated packet.
// But if there are only received packets, then it replaces the newest one.
// If a packet is generated while the buffer is full, it is discarded.
//rp-old> rp-new > gp-old > gp-new

type RPP4 struct {
	basePolicy
}

func (g RPP4) CreateSlot(b buffer.Buffer, p packet.Packet, sensor_id int) (int, error) {
	index, err := g.getFreeSlot(b)
	if err == nil {
		return index, nil
	}

	// if the packet is generated not received then, do not allocate slot
	if sensor_id == p.Parent_id {
		return 0, CAN_NOT_CREATE_SLOT_ERROR
	}

    // Check if there are generated packets in the buffer
    flag := 0
    index = 0

    for i, pac := range b.Packets {
        if pac.Parent_id == sensor_id {
            flag = 1
            index = i
            break
        }
    }

    min_time := b.Packets[index].GetTimestamp()

	// In case generated packets are there in buffer
    if flag == 1 {
        for i, pac := range b.Packets {
            if pac.Parent_id == sensor_id {
                if !min_time.After(pac.GetTimestamp()) {
                    min_time, index = pac.GetTimestamp(), i
                }
            }
        }
    } else { // If only recieved packets are there in buffer
        for i, pac := range b.Packets {
		    if !min_time.After(pac.GetTimestamp()) {
			    min_time, index = pac.GetTimestamp(), i
            }
        }
    }

	return index, err
}
