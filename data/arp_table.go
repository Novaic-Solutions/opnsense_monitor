package data

import "time"

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