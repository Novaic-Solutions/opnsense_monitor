package data

import "time"

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
// |--  Rulenr


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
	keyTitle := "{interface=\"" + fwLogEntry.Interface + "\",action=\"" + fwLogEntry.Action + "\",src=\"" + fwLogEntry.Src + "\",dst=\"" + fwLogEntry.Dst + "\",protoname=\"" + fwLogEntry.Protoname + "\",src_port=\"" + fwLogEntry.Srcport + "\",dst_port=\"" + fwLogEntry.Dstport + "\",rulenr=\"" + fwLogEntry.Rulenr + "}"
	if _, exists := dh.Metrics["firewall_log_entries"+keyTitle]; !exists {
		dh.Metrics["firewall_log_entries"+keyTitle] = 1
	} else {
		dh.Metrics["firewall_log_entries"+keyTitle]++
	}
	dh.MetricsLastUpdated["firewall_log_entries"+keyTitle] = time.Now()
}