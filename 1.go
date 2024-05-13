package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WiMAXStation struct {
	ID      int
	Channel int
}

func NewWiMAXStation(id, channel int) *WiMAXStation {
	return &WiMAXStation{
		ID:      id,
		Channel: channel,
	}
}

func (ws *WiMAXStation) Transmit(data string) {
	fmt.Printf("WiMAX Station %d on Channel %d transmitting: %s\n", ws.ID, ws.Channel, data)
}

type WiMAXNetwork struct {
	Stations []*WiMAXStation
}

func NewWiMAXNetwork(numStations, numChannels int) *WiMAXNetwork {
	network := &WiMAXNetwork{}
	for i := 0; i < numStations; i++ {
		channel := rand.Intn(numChannels) + 1
		network.Stations = append(network.Stations, NewWiMAXStation(i+1, channel))
	}
	return network
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numStations := 50
	numChannels := 100
	time.Sleep(8 * time.Second) 

	network := NewWiMAXNetwork(numStations, numChannels)

	for i := 0; i < 10; i++ {
		stationID := rand.Intn(numStations)
		data := fmt.Sprintf("Data Packet %d", i+1)
		network.Stations[stationID].Transmit(data)
	}
	time.Sleep(8 * time.Second) 

	fmt.Println("WiMAX Simulation Completed.")
}

