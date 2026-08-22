package data

import "fmt"

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