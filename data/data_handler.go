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

//----------------------------------------------------------------------------


//----------------------------------------------------------------------------
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
//----------------------------------------------------------------------------


//----------------------------------------------------------------------------
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
//----------------------------------------------------------------------------


//----------------------------------------------------------------------------
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
//----------------------------------------------------------------------------
