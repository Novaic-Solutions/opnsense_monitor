package data

import "time"

//-------------------------------------------------------------------------------------------------------------------
// Metrics for "/api/diagnostics/firewall/query_states"
// Firewall -> Diagnostics -> States
//----------------------------------------------------------------------------
// response_obj_type: "FirewallStates"
//
// Loop over the FirewallStates, which is a slice that contains FirewallState
// 	|--   Descr
//  |     Nat_addr
//  |     Nat_port
//  |     Gateway
//  |     Interface
//  |     Proto
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
// 		firewall_state_packets{descr="", nat_addr="", nat_port="", gateway="", interface="", proto="", direction="", dst_addr="", dst_port="", src_addr="", src_port="", state=""} pkts
//      ...
//	  	# HELP firewall_state_bytes  Total bytes passed through the firewall state
//		# TYPE firewall_state_bytes counter
// 		firewall_state_bytes{descr="", nat_addr="", nat_port="", gateway="", interface="", proto="", direction="", dst_addr="", dst_port="", src_addr="", src_port="", state=""} bytes
//-------------------------------------------------------------------------------------------------------------------
func (dh *DataHandler) ProcessFirewallStates(fwStates FirewallStates) {
	for _, fwState := range fwStates.Rows {
		dh.ProcessFirewallState(fwState)
	}
}

func (dh *DataHandler) ProcessFirewallState(fwState FirewallState) {
	keyTitle := "{descr=\"" + fwState.Descr + "\",nat_addr=\"" + fwState.Nat_addr + "\",nat_port=\"" + fwState.Nat_port + "\",gateway=\"" + fwState.Gateway + "\",interface=\"" + fwState.Interface + "\",proto=\"" + fwState.Proto + "\",direction=\"" + fwState.Direction + "\",dst_addr=\"" + fwState.Dst_addr + "\",dst_port=\"" + fwState.Dst_port + "\",src_addr=\"" + fwState.Src_addr + "\",src_port=\"" + fwState.Src_port + "\",state=\"" + fwState.State + "\"}"
	dh.Metrics["firewall_state_packets" + keyTitle] = uint64(fwState.Pkts[0]) // Assuming Pkts is a slice with at least one element
	dh.MetricsLastUpdated["firewall_state_packets" + keyTitle] = time.Now()
	dh.Metrics["firewall_state_bytes" + keyTitle] = uint64(fwState.Bytes[0]) // Assuming Bytes is a slice with at least one element
	dh.MetricsLastUpdated["firewall_state_bytes" + keyTitle] = time.Now()
}
