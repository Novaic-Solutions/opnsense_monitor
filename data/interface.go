package data


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
