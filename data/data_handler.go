//----------------------------------------------------------------------------------------------
// The application will store the data in a in-memory database (map).
// 
// As the data comes in from the API responses, the map will either be updated with the

// new data, or if the data is new, it will be added to the map. 
// It could also possibly just be added to a splice depending on the endpoint
// and the data being returned as well as what 
// prometheus metrics are being generated for that data.
//----------------------------------------------------------------------------------------------
package data

import (
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"fmt"
	"sync"
	"time"
	"strings"
)

type DataHandler struct {
	DataMutex *sync.Mutex
	Metrics map[string]uint64
	MetricsLastUpdated map[string]time.Time
	Incoming chan config.EndpointResponse
	Request chan string
	Outgoing chan string
}

//-----------------------------------------------------------------------------
//  This will be used in a go rountine to handle the incoming data from 
//  the clients API responses and update the Metrics map with the data.
//-----------------------------------------------------------------------------
func (dh *DataHandler) HandleIncomingData() {

	for incomingData := range dh.Incoming {
		
		fmt.Printf("DataHandler.HandleIncomingData: Received data from endpoint response type: %s\n", incomingData.ResponseDataType)
		
		dh.DataMutex.Lock()

		fmt.Printf("DataHandler.HandleIncomingData: Mutex Locked.")

		// Process the incoming data and update the Metrics map
		switch incomingData.ResponseDataType {
		case "FirewallLogEntry":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			fmt.Printf("DataHandler.HandleIncomingData: Processing FirewallLogEntry data.\n")
			dh.ProcessFirewallLogEntries(incomingData.Data.([]FirewallLogEntry))
		case "ArpTable":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			fmt.Printf("DataHandler.HandleIncomingData: Processing ArpTable data.\n")
			dh.ProcessArpTable(incomingData.Data.(ArpTable))
		case "IfaceStatistics":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			fmt.Printf("DataHandler.HandleIncomingData: Processing IfaceStatistics data.\n")
			dh.ProcessIfaceStatistics(incomingData.Data.(IfaceStatistics))
		case "FirewallSessions":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			fmt.Printf("DataHandler.HandleIncomingData: Processing FirewallSessions data.\n")
			dh.ProcessFirewallSessions(incomingData.Data.(FirewallSessions))
		case "FirewallStates":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			fmt.Printf("DataHandler.HandleIncomingData: Processing FirewallStates data.\n")
			dh.ProcessFirewallStates(incomingData.Data.(FirewallStates))
		case "Interfaces":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			fmt.Printf("DataHandler.HandleIncomingData: Processing Interfaces data.\n")
			dh.ProcessInterfaces(incomingData.Data.(Interfaces))
		default:
			// Handle unknown response data types if necessary
			fmt.Printf("DataHandler.HandleIncomingData: Unknown response data type: %s\n", incomingData.ResponseDataType)
		}
		dh.DataMutex.Unlock()
		fmt.Printf("DataHandler.HandleIncomingData: Mutex Unlocked.")
		// Clear old data from the Metrics and MetricsLastUpdated maps
		dh.ClearOldData()

	}
}

func (dh *DataHandler) ClearOldData() {
	// Implement logic to clear old data from Metrics and MetricsLastUpdated maps
	dh.DataMutex.Lock()
	for metricName, lastUpdated := range dh.MetricsLastUpdated {
		if time.Since(lastUpdated) > time.Hour {
			delete(dh.Metrics, metricName)
			delete(dh.MetricsLastUpdated, metricName)
		}
	}
	dh.DataMutex.Unlock()
}

//----------------------------------------------------------------------------
// This will be used in a go routine to handle the requests from the main application
//----------------------------------------------------------------------------
func (dh *DataHandler) HandleRequests() {
	for request := range dh.Request {
		fmt.Printf("DataHandler.HandleRequests: Received new request")
		
		dh.DataMutex.Lock()
		fmt.Printf("DataHandler.HandleRequests: Mutex Locked.")
		defer dh.DataMutex.Unlock()

		switch request {
		case "metrics":
			fmt.Printf("DataHandler.HandleRequests: Creating metrics string.\n")
			metricsString := dh.CreateMetricsString()
			fmt.Printf("DataHandler.HandleRequests: Sending metrics string of length %d.\n", len(metricsString))
			dh.Outgoing <- metricsString
		default:
			fmt.Printf("DataHandler.HandleRequests: Unknown request: %s\n", request)
		}
	}
}

func NewDataHandler(dataMutex *sync.Mutex, request chan string, incoming chan config.EndpointResponse, outgoing chan string) *DataHandler {
	return &DataHandler{
		DataMutex: dataMutex,
		Metrics: make(map[string]uint64),
		MetricsLastUpdated: make(map[string]time.Time),
		Request: request,
		Incoming: incoming,
		Outgoing: outgoing,
	}
}

func (dh *DataHandler) CreateMetricsString() string {
	metricsString := ""
	firewallIfaceStatistics := 0
	firewallLogEnginers := 0

	fmt.Printf("DataHandler.CreateMetricsString: Preparing to loop over metrics\n")

	for metricName, metricValue := range dh.Metrics {
		fmt.Printf("DataHandler.CreateMetricsString: Processing metric: %s with value: %d\n", metricName, metricValue)
		if strings.HasPrefix(metricName, "firewall_interface_statistics") {
			if firewallIfaceStatistics == 0 {
				helpAndTypeLines := dh.CreateHelpAndTypeLines(metricName)
				if helpAndTypeLines != "" {
					metricsString += helpAndTypeLines
				}
				firewallIfaceStatistics++
			}
		} else if strings.HasPrefix(metricName, "firewall_log_entries") {
			if firewallLogEnginers == 0 {
				helpAndTypeLines := dh.CreateHelpAndTypeLines(metricName)
				if helpAndTypeLines != "" {
					metricsString += helpAndTypeLines
				}
				firewallLogEnginers++
			}
		} else {
			helpAndTypeLines := dh.CreateHelpAndTypeLines(metricName)
			if helpAndTypeLines != "" {
				metricsString += helpAndTypeLines
			}
		}
		metricsString += fmt.Sprintf("%s %d\n", metricName, metricValue)
	}

	return metricsString
}

func (dh *DataHandler) CreateHelpAndTypeLines(metricName string) string {
	switch {
		//-------------------------------------------------------------------------------
		// ARP Table Metrics
		//----------------------------------------------------------------------------
		case strings.HasPrefix(metricName, "arp_table_entry"):
			return "# HELP arp_table_entry Total number of entries in the ARP table\n# TYPE arp_table_entry counter\n"
		
		//-----------------------------------------------------------------------
		// Interface Statistics
		//-----------------------------------------------------------------------
		case strings.HasPrefix(metricName, "interface_statistic"):
			measurement := ""
			if strings.HasSuffix(metricName, "packets") {
				measurement = "packets"
			} else if strings.HasSuffix(metricName, "errors") {
				measurement = "errors"
			} else if strings.HasSuffix(metricName, "bytes") {
				measurement = "bytes"
			}

			nameSplit := strings.Split(metricName, "_")
			measurementType := "collisions"
			if len(nameSplit) >= 4 {
				measurementType = nameSplit[3]
			}

			return "# HELP " + metricName + " Number of " + measurementType + " " + measurement + " on the interface\n# TYPE " + metricName + " counter\n"
		
		//-----------------------------------------------------------------------
		// Firewall Sessions
		//-----------------------------------------------------------------------
		case strings.HasPrefix(metricName, "firewall_session"):
			measurement := ""
			if strings.HasSuffix(metricName, "packets") {
				measurement = "packets"
			} else if strings.HasSuffix(metricName, "bytes") {
				measurement = "bytes"
			}

			return "# HELP " + metricName + " Total number of " + measurement + " for the firewall session\n# TYPE " + metricName + " counter\n"

		//-----------------------------------------------------------------------
		// Firewall States
		//-----------------------------------------------------------------------
		case strings.HasPrefix(metricName, "firewall_state"):
			measurement := ""
			if strings.HasSuffix(metricName, "packets") {
				measurement = "packets"
			} else if strings.HasSuffix(metricName, "bytes") {
				measurement = "bytes"
			}
			
			return "# HELP " + metricName + " Total number of " + measurement + " for the firewall state\n# TYPE " + metricName + " counter\n"
		
		//-----------------------------------------------------------------------
		// Firewall Interface Statistics
		//-----------------------------------------------------------------------
		case strings.HasPrefix(metricName, "firewall_interface_statistics"):
			// These may not be needed currently, but I am leaving them here in case its decided at a later date
			// to split the metrics up into packets, bytes, errors, and collisions.  For now, they are all just counters for the interface.
			// if strings.HasSuffix(metricName, "packets") {
			// 	measurement = "packets"
			// } else if strings.HasSuffix(metricName, "bytes") {
			// 	measurement = "bytes"
			// } else if strings.HasSuffix(metricName, "errors") {
			// 	measurement = "errors"
			// } else if strings.HasSuffix(metricName, "collisions") {
			// 	measurement = "collisions"
			// }

			return "# HELP " + metricName + " Total number of all statistics for the firewall interface\n# TYPE " + metricName + " counter\n"
		
		//-----------------------------------------------------------------------
		// Firewall Log Entries
		//-----------------------------------------------------------------------
		case strings.HasPrefix(metricName, "firewall_log_entries"):
			return "# HELP firewall_log_entries Total number of log entries for the unique combination of values\n# TYPE firewall_log_entries counter\n"
		
		default:
			return ""
	}
}

//-------------------------------------------------------------------------------------------------------------------
// Naming convention for the prometheus metrics
// prometheus_metric_name{label="value"} value
// 
// For each type of metric, a line will be generated as follows:
// # HELP prometheus_metric_name description of the metric including the measurement units
// # TYPE prometheus_metric_name type of metric (counter, gauge, histogram, summary)
//----------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------
// Metrics for "/api/diagnostics/interface/search_arp/"
// Interfaces -> Diagnostics -> ARP Table
//----------------------------------------------------------------------------
//  response_obj_type: "ArpTable"

// Gathers the following data:
//     IP address
//     MAC address
//     Interface
//     Interface Name
//     Manufacturer
//     Hostname

// Prometheus metrics will be generated for each entry in the ARP table as follows:
// arp_table_entry{ip="", mac="", interface="", manufacturer="", hostname=""} 1
//-------------------------------------------------------------------------------------------------------------------
func (dh *DataHandler) GetArpEntryTitle(arpEntry ArpTableEntry) string {
	// Create the key string for the prometheus metric
	keyTitle := "arp_table_entry{ip=\"" + arpEntry.Ip + "\", mac=\"" + arpEntry.Mac + "\", interface=\"" + arpEntry.Interface + "\", manufacturer=\"" + arpEntry.Manufacturer + "\", hostname=\"" + arpEntry.Hostname + "\"}"
	return keyTitle
}

func (dh *DataHandler) ProcessArpTable(arpTable ArpTable) {
	for _, arpEntry := range arpTable.Rows {
		dh.ProcessArpTableEntry(arpEntry)
	}
}

func (dh *DataHandler) ProcessArpTableEntry(arpEntry ArpTableEntry) {
	// Create the key string for the prometheus metric
	keyTitle := dh.GetArpEntryTitle(arpEntry)

	// Check if the key already exiztzs in the Metrics map, if not, 
	// add it with a value of 1, if it does exist, increment the value by 1
	if _, exists := dh.Metrics[keyTitle]; !exists {
		dh.Metrics[keyTitle] = 1
	} else {
		dh.Metrics[keyTitle]++
	}
	dh.MetricsLastUpdated[keyTitle] = time.Now()
}

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

//-------------------------------------------------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/query_pf_top"
// Firewall -> Diagnostics -> Sessions
//----------------------------------------------------------------------------
//  response_obj_type: "FirewallSessions"

// Gathers the following data:  in a map[string]FirewallSessions
// Loop over the FirewallSessions, which is a slice that contains FirewallSessions
// 	|--   Src_addr
//  |     Src_port
//  |     Dst_addr
//  |     Dst_port
//  |     Proto
//  |     State
//  |     Dir
//  |     Age
//  |     Expire
//  |--   Descr
//        Pkts
//        Bytes

// 
// Create metrics like the following for each entry in the FirewallSessions slice:
//		# HELP firewall_session_packets  Total packets passed through the firewall session
//		# TYPE firewall_session_packets counter
// 		firewall_session_packets{src_ip="", src_port="", dst_ip="", dst_port="", proto="", state="", dir="", age="", expirs="", descr=""} pkts
//      ...

//      # HELP firewall_session_bytes  Total bytes passed through the firewall session
//		# TYPE firewall_session_bytes counter
// 		firewall_session_bytes{src_ip="", src_port="", dst_ip="", dst_port="", proto="", state="", dir="", age="", expirs="", descr=""} bytes

//-------------------------------------------------------------------------------------------------------------------
func (dh *DataHandler) ProcessFirewallSessions(fwSessions FirewallSessions) {
	for _, fwSession := range fwSessions.Rows {
		dh.ProcessFirewallSession(fwSession)
	}
}

func (dh *DataHandler) ProcessFirewallSession(fwSession FirewallSession) {
	keyTitle := "{src_ip=\"" + fwSession.Src_addr + "\", src_port=\"" + fwSession.Src_port + "\", dst_ip=\"" + fwSession.Dst_addr + "\", dst_port=\"" + fwSession.Dst_port + "\", proto=\"" + fwSession.Proto + "\", state=\"" + fwSession.State + "\", dir=\"" + fwSession.Dir + "\", age=\"" + fmt.Sprintf("%d", fwSession.Age) + "\", expires=\"" + fmt.Sprintf("%d", fwSession.Expire) + "\", descr=\"" + fwSession.Descr + "\"}"
	dh.Metrics["firewall_session_packets" + keyTitle] = uint64(fwSession.Pkts)
	dh.Metrics["firewall_session_bytes" + keyTitle] = uint64(fwSession.Bytes)
}

//-------------------------------------------------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/query_states"
// Firewall -> Diagnostics -> States
//----------------------------------------------------------------------------
// response_obj_type: "FirewallStates"
//
// Loop over the FirewallStates, which is a slice that contains FirewallState
// 	|--   Label
//  |     Descr
//  |     Nat_addr
//  |     Nat_port
//  |     Gateway
//  |     Interface
//  |     Proto
//  |     Ipproto
//  |     Direction
//  |     Dst_addr
//  |     Dst_port
//  |     Src_addr
//  |     Src_port
//  |--	  State
//        Pkts
//        Bytes

// Create metrics like the following for each entry in the FirewallStates slice:
//		# HELP firewall_state_packets  Total packets passed through the firewall state
//		# TYPE firewall_state_packets counter
// 		firewall_state_packets{label="", descr="", nat_addr="", nat_port="", gateway="", interface="", proto="", ipproto="", direction="", dst_addr="", dst_port="", src_addr="", src_port="", state=""} pkts
//      ...
//	  	# HELP firewall_state_bytes  Total bytes passed through the firewall state
//		# TYPE firewall_state_bytes counter
// 		firewall_state_bytes{label="", descr="", nat_addr="", nat_port="", gateway="", interface="", proto="", ipproto="", direction="", dst_addr="", dst_port="", src_addr="", src_port="", state=""} bytes
//-------------------------------------------------------------------------------------------------------------------
func (dh *DataHandler) ProcessFirewallStates(fwStates FirewallStates) {
	for _, fwState := range fwStates.Rows {
		dh.ProcessFirewallState(fwState)
	}
}

func (dh *DataHandler) ProcessFirewallState(fwState FirewallState) {
	keyTitle := "{label=\"" + fwState.Label + "\", descr=\"" + fwState.Descr + "\", nat_addr=\"" + fwState.Nat_addr + "\", nat_port=\"" + fwState.Nat_port + "\", gateway=\"" + fwState.Gateway + "\", interface=\"" + fwState.Interface + "\", proto=\"" + fwState.Proto + "\", ipproto=\"" + fwState.Ipproto + "\", direction=\"" + fwState.Direction + "\", dst_addr=\"" + fwState.Dst_addr + "\", dst_port=\"" + fwState.Dst_port + "\", src_addr=\"" + fwState.Src_addr + "\", src_port=\"" + fwState.Src_port + "\", state=\"" + fwState.State + "\"}"
	dh.Metrics["firewall_state_packets" + keyTitle] = uint64(fwState.Pkts[0]) // Assuming Pkts is a slice with at least one element
	dh.MetricsLastUpdated["firewall_state_packets" + keyTitle] = time.Now()
	dh.Metrics["firewall_state_bytes" + keyTitle] = uint64(fwState.Bytes[0]) // Assuming Bytes is a slice with at least one element
	dh.MetricsLastUpdated["firewall_state_bytes" + keyTitle] = time.Now()
}

//-------------------------------------------------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/pf_statistics/interfaces"
// Firewall -> Diagnostics -> Statistics -> Interfaces Tab
//----------------------------------------------------------------------------
// response_obj_type: "Interfaces"
//
// Loop over the Interfaces, which is a slice that contains IfaceTraffic structs
//   |--  Interface
//        In4_pass_packets
//        In4_pass_bytes
//        In4_block_packets
//        In4_block_bytes
//        In6_pass_packets
//        In6_pass_bytes
//        In6_block_packets
//        In6_block_bytes
//        Out4_pass_packets
//		  Out4_pass_bytes
//		  Out4_block_packets
//		  Out4_block_bytes
//		  Out6_pass_packets
//		  Out6_pass_bytes
//		  Out6_block_packets
//		  Out6_block_bytes

// Create metrics like the following for each entry in the Interfaces map:
//		# HELP firewall_interface_in4_pass_packets  Total number of IPv4 packets passed through the interface
//		# TYPE firewall_interface_in4_pass_packets counter
// 		firewall_interface_statistics_in4_pass_packets{interface=""} in4_pass_packets
//	  ...
//		# HELP firewall_interface_out6_block_bytes  Total number of IPv6 bytes blocked by the interface
//		# TYPE firewall_interface_out6_block_bytes counter
// 		firewall_interface_statistics_out6_block_bytes{interface=""} out6_block_bytes
//-------------------------------------------------------------------------------------------------------------------
func (dh *DataHandler) ProcessInterfaces(interfaces Interfaces) {
	for ifaceName, ifaceTraffic := range interfaces.Interfaces {
		dh.ProcessInterfaceTraffic(ifaceName, ifaceTraffic)
	}
}

func (dh *DataHandler) ProcessInterfaceTraffic(ifaceName string, ifaceTraffic IfaceTraffic) {
	keyTitle := "{interface=\"" + ifaceName + "\"}"
	dh.Metrics["firewall_interface_statistics_in4_pass_packets" + keyTitle] = ifaceTraffic.In4_pass_packets
	dh.Metrics["firewall_interface_statistics_in4_pass_bytes" + keyTitle] = ifaceTraffic.In4_pass_bytes
	dh.Metrics["firewall_interface_statistics_in4_block_packets" + keyTitle] = ifaceTraffic.In4_block_packets
	dh.Metrics["firewall_interface_statistics_in4_block_bytes" + keyTitle] = ifaceTraffic.In4_block_bytes
	dh.Metrics["firewall_interface_statistics_out4_pass_packets" + keyTitle] = ifaceTraffic.Out4_pass_packets
	dh.Metrics["firewall_interface_statistics_out4_pass_bytes" + keyTitle] = ifaceTraffic.Out4_pass_bytes
	dh.Metrics["firewall_interface_statistics_out4_block_packets" + keyTitle] = ifaceTraffic.Out4_block_packets
	dh.Metrics["firewall_interface_statistics_out4_block_bytes" + keyTitle] = ifaceTraffic.Out4_block_bytes
	dh.Metrics["firewall_interface_statistics_in6_pass_packets" + keyTitle] = ifaceTraffic.In6_pass_packets
	dh.Metrics["firewall_interface_statistics_in6_pass_bytes" + keyTitle] = ifaceTraffic.In6_pass_bytes
	dh.Metrics["firewall_interface_statistics_in6_block_packets" + keyTitle] = ifaceTraffic.In6_block_packets
	dh.Metrics["firewall_interface_statistics_in6_block_bytes" + keyTitle] = ifaceTraffic.In6_block_bytes
	dh.Metrics["firewall_interface_statistics_out6_pass_packets" + keyTitle] = ifaceTraffic.Out6_pass_packets
	dh.Metrics["firewall_interface_statistics_out6_pass_bytes" + keyTitle] = ifaceTraffic.Out6_pass_bytes
	dh.Metrics["firewall_interface_statistics_out6_block_packets" + keyTitle] = ifaceTraffic.Out6_block_packets
	dh.Metrics["firewall_interface_statistics_out6_block_bytes" + keyTitle] = ifaceTraffic.Out6_block_bytes
}

//-------------------------------------------------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/log"
// Firewall -> Log Files -> Live View
//     Used as live traffic counters for now
//----------------------------------------------------------------------------
// Count the number of log entries for each type of log.
// response_obj_type: "FirewallLogEntry"
// 
// |--  Interface
// |    Action
// |    Src
// |    Dst
// |    Protoname
// |    Src_port
// |    Dst_port
// |    Rulenr
// |--  Ipversion


// Form a string for each log entry using the values from above and use it as a key in a map[string]uint64
// to count the number of log entries for each unique combination of values.
// 
// Create metrics like the following for each unique combination of values in the log entries:
//		# HELP firewall_log_entries  Total number of log entries for the unique combination of values
//		# TYPE firewall_log_entries counter
// 		firewall_log_entries{interface="", action="", src="", dst="", protoname="", src_port="", dst_port="", rulenr="", ipversion=""} count
//      ...
//-------------------------------------------------------------------------------------------------------------------
func (dh *DataHandler) ProcessFirewallLogEntries(fwLogEntries []FirewallLogEntry) {
	for _, fwLogEntry := range fwLogEntries {
		dh.ProcessFirewallLogEntry(fwLogEntry)
	}
}

func (dh *DataHandler) ProcessFirewallLogEntry(fwLogEntry FirewallLogEntry) {
	keyTitle := "{interface=\"" + fwLogEntry.Interface + "\", action=\"" + fwLogEntry.Action + "\", src=\"" + fwLogEntry.Src + "\", dst=\"" + fwLogEntry.Dst + "\", protoname=\"" + fwLogEntry.Protoname + "\", src_port=\"" + fwLogEntry.Srcport + "\", dst_port=\"" + fwLogEntry.Dstport + "\", rulenr=\"" + fwLogEntry.Rulenr + "\", ipversion=\"" + fwLogEntry.Ipversion + "\"}"
	if _, exists := dh.Metrics["firewall_log_entries"+keyTitle]; !exists {
		dh.Metrics["firewall_log_entries"+keyTitle] = 1
	} else {
		dh.Metrics["firewall_log_entries"+keyTitle]++
	}
	dh.MetricsLastUpdated["firewall_log_entries"+keyTitle] = time.Now()
}