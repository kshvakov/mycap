package agent

import (
	"bytes"
	"fmt"
	"log"
	"mycap/libs"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Collector struct {
	Device string `json:"device"`

	BPFFilter string `json:"bpf_filter"`
	TXTFilter string `json:"txt_filter"`

	MaxQueryLen int `json:"max_query_len"`

	buffer  map[string]libs.Query
	queries libs.Queries
}

func (self *Collector) Collect() {

	self.buffer = make(map[string]libs.Query)

	handle, err := pcap.OpenLive(self.Device, int32(self.MaxQueryLen)+5, true, time.Second)
	defer handle.Close()

	if err != nil {
		log.Fatal(err)
	}

	handle.SetBPFFilter(self.BPFFilter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	var (
		ipLayer  *layers.IPv4
		tcpLayer *layers.TCP
		ok       bool
	)

	for packet := range packetSource.Packets() {
		if applicationLayer := packet.ApplicationLayer(); applicationLayer != nil {

			if ipLayer, ok = packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4); !ok {
				continue
			}
			if tcpLayer, ok = packet.Layer(layers.LayerTypeTCP).(*layers.TCP); !ok {
				continue
			}

			playload := applicationLayer.Payload()

			if len(playload) < 5 {
				continue
			}

			length := int(playload[0]) | int(playload[1])<<8 | int(playload[2])<<16

			if length > len(playload)-4 {
				continue
			}

			ID := fmt.Sprintf("%s%d:%s%d\n", ipLayer.SrcIP, tcpLayer.SrcPort, ipLayer.DstIP, tcpLayer.DstPort)

			if uint8(playload[4]) == 3 {
				// is request
				if self.isAllowedByFilter(&playload, length) {
					self.buffer[ID] = libs.Query{

						Query: string(playload[5 : length+4]),
						Start: packet.Metadata().Timestamp,

						SrcIP: ipLayer.SrcIP,
						DstIP: ipLayer.DstIP,

						SrcPort: tcpLayer.SrcPort,
						DstPort: tcpLayer.DstPort,

						ID: ID,
					}
				}
			} else {
				// is response
				if query, found := self.buffer[ID]; found {

					query.ID = fmt.Sprintf(
						"%s%d:%s%d:%d\n",
						ipLayer.SrcIP, tcpLayer.SrcPort, ipLayer.DstIP, tcpLayer.DstPort, time.Now().UnixNano(),
					)
					query.Stop = packet.Metadata().Timestamp
					query.Duration = packet.Metadata().Timestamp.Sub(query.Start)
					query.WithResponse = true

					self.queries = append(self.queries, query)
					delete(self.buffer, ID)
				}
			}
		}
	}
}

func (self *Collector) isAllowedByFilter(playload *[]byte, length int) bool {
	return self.TXTFilter == "" || bytes.Contains(bytes.ToLower((*playload)[5:length+4]), bytes.ToLower([]byte(self.TXTFilter)))
}
