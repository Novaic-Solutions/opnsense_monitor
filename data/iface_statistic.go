package data

import "fmt"

//-------------------------------------------------------------------------------------------------------------------
// Metrics for "/api/diagnostics/interface/getInterfaceStatistics"
// Interfaces -> Diagnostics -> Netstat
//----------------------------------------------------------------------------
//  response_obj_type: "InterfaceStatistics"

// Gathers the following data:  in a map[string]InterfaceStatistics
//     for the key -> "[INTERFACE_NAME](interface) / mac_address" or "[INTERFACE_NAME](interface) / ip_address"
//	   for the value -> InterfaceStatistics struct
//      |--     Interface Name
//		|       Flags
//		|       MTU
//		|       Network
//		|--     MAC address
//		
//              recieved-packets
//              recieved-errors
//              dropped-packets
//				received-bytes
//				sent-packets
//				send-errors
//				sent-bytes
//				collisions

// Prometheus metrics will be generated for each entry in the InterfaceStatistics map as follows:
// interface_statistics{name="", flags="", mtu="", address=""} value
//
// For each InterfaceStatistics entry, for each value that isnt in the brackets above, provide a
// # HELP and # TYPE line for them

// example:
// # HELP interface_statistics_received_packets Total number of packets received on the interface
// # TYPE interface_statistics_received_packets counter
// interface_statistics_received_packets{name="", flags="", mtu="", address=""} value
// # HELP interface_statistics_received_errors Total number of errors received on the interface
// # TYPE interface_statistics_received_errors counter
// interface_statistics_received_errors{name="", flags="", mtu="", address=""} value

//-------------------------------------------------------------------------------------------------------------------
func (dh *DataHandler) ProcessIfaceStatistics(ifaceStats IfaceStatistics) {
	for _, value := range ifaceStats.Statistics {
		dh.ProcessIfaceStatistic(value)
	}
}

func (dh *DataHandler) ProcessIfaceStatistic(ifaceStat IfaceStatistic) {
	keyTitle := "{name=\"" + ifaceStat.Name + "\", flags=\"" + ifaceStat.Flags + "\", mtu=\"" + fmt.Sprintf("%d", ifaceStat.Mtu) + "\", address=\"" + ifaceStat.Address + "\"}"
	dh.Metrics["interface_statistic_received_packets" + keyTitle] = ifaceStat.ReceivedPackets
	dh.Metrics["interface_statistic_received_errors" + keyTitle] = ifaceStat.ReceivedErrors
	dh.Metrics["interface_statistic_dropped_packets" + keyTitle] = ifaceStat.DroppedPackets
	dh.Metrics["interface_statistic_received_bytes" + keyTitle] = ifaceStat.ReceivedBytes
	dh.Metrics["interface_statistic_sent_packets" + keyTitle] = ifaceStat.SentPackets
	dh.Metrics["interface_statistic_send_errors" + keyTitle] = ifaceStat.SendErrors
	dh.Metrics["interface_statistic_sent_bytes" + keyTitle] = ifaceStat.SentBytes
	dh.Metrics["interface_statistic_collisions" + keyTitle] = ifaceStat.Collisions
}