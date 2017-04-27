package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"mycap/gui"
	"time"
)

var (
	device        = flag.String("device", "lo", "")
	guiPath       = flag.String("gui", "./gui", "")
	BPFFilter     = flag.String("bpf_filter", "tcp and port 3306", "")
	queryFilter   = flag.String("query_filter", "", "not case-sensitive")
	slowQueryTime = flag.Int64("slow_query_time", 0, "in milliseconds")
	maxQueryLen   = flag.Int("max_query_len", 9192, "")
	queries       = make(map[string]query)
)

type query struct {
	query string
	start time.Time
}

func main() {

	flag.Parse()

	gui.InitGui(*guiPath)

	handle, err := pcap.OpenLive(*device, int32(*maxQueryLen)+5, true, time.Second)

	defer handle.Close()

	if err != nil {

		log.Fatal(err)
	}

	handle.SetBPFFilter(*BPFFilter)

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

			// is request
			if uint8(playload[4]) == 3 {

				isAllowedByFilter := *queryFilter == "" || bytes.Contains(bytes.ToLower(playload[5:length+4]), bytes.ToLower([]byte(*queryFilter)))

				if isAllowedByFilter {
					qkey := fmt.Sprintf("%s%d:%s%d\n", ipLayer.SrcIP, tcpLayer.SrcPort, ipLayer.DstIP, tcpLayer.DstPort)
					queries[qkey] = query{
						query: string(playload[5 : length+4]),
						start: packet.Metadata().Timestamp,
					}
				}
			} else { // is response

				qkey := fmt.Sprintf("%s%d:%s%d\n", ipLayer.DstIP, tcpLayer.DstPort, ipLayer.SrcIP, tcpLayer.SrcPort)

				if query, found := queries[qkey]; found {

					queryTime := packet.Metadata().Timestamp.Sub(query.start)

					if *slowQueryTime == 0 || queryTime.Nanoseconds() > *slowQueryTime*1000000 {
						gui.AllQueries.Add(
							// from IP
							fmt.Sprintf("%s", ipLayer.DstIP),
							// fmt.Sprintf("%s:%d", ipLayer.DstIP, tcpLayer.DstPort),
							// to IP:port
							fmt.Sprintf("%s:%d", ipLayer.SrcIP, tcpLayer.SrcPort),
							query.query,
							query.start,
							queryTime,
						)
						// fmt.Printf("-[ QUERY %f s]-:\n%s\n\n\n", queryTime.Seconds(), query.query)
					}

					delete(queries, qkey)
				}
			}

		}
	}
}
