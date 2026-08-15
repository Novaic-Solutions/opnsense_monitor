package data

import (
	//"fmt"
	//"github.com/Novaic-Solutions/opnsense_monitor/config"
)


//----------------------------------------------------------------------------------------------
// The application will store the data in a in-memory database (map).
// 
// As the data comes in from the API responses, the map will either be updated with the
// new data, or if the data is new, it will be added to the map. 
// It could also possibly just be added to a splice depending on the endpoint
// and the data being returned as well as what prometheus metrics are being generated for that data.
//----------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------
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

//----------------------------------------------------------------------------




//----------------------------------------------------------------------------
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

//----------------------------------------------------------------------------


//----------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/query_pf_top"
// Firewall -> Diagnostics -> Sessions
//----------------------------------------------------------------------------
//  response_obj_type: "FirewallSessions"

//----------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/query_states"
// Firewall -> Diagnostics -> States
//----------------------------------------------------------------------------
// response_obj_type: "FirewallStates"


//----------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/pf_statistics/interfaces"
// Firewall -> Diagnostics -> Statistics -> Interfaces Tab
//----------------------------------------------------------------------------
// response_obj_type: "Interfaces"

//----------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/log"
// Firewall -> Log Files -> Live View
//----------------------------------------------------------------------------
// Count the number of log entries for each type of log.
// response_obj_type: "FirewallLogEntry"