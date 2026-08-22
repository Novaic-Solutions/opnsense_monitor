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
	"context"
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
func (dh *DataHandler) HandleIncomingData(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("DataHandler.HandleIncomingData: Mutex Address --> %p\n", dh.DataMutex)

	for {
		
	}

	for incomingData := range dh.Incoming {
		select {
		case <-ctx.Done():
			fmt.Printf("DataHandler.HandleIncomingData: Context cancelled, exiting.\n")
			return
		default:
		}
		
		fmt.Printf("DataHandler.HandleIncomingData: Received data from endpoint response type: %s\n", incomingData.ResponseDataType)
		
		dh.DataMutex.Lock()

		fmt.Printf("DataHandler.HandleIncomingData: Mutex Locked.\n")

		// Process the incoming data and update the Metrics map
		switch incomingData.ResponseDataType {
		case "FirewallLogEntry":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			dh.ProcessFirewallLogEntries(incomingData.Data.([]FirewallLogEntry))
		case "ArpTable":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			dh.ProcessArpTable(incomingData.Data.(ArpTable))
		case "IfaceStatistics":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			dh.ProcessIfaceStatistics(incomingData.Data.(IfaceStatistics))
		case "FirewallSessions":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			dh.ProcessFirewallSessions(incomingData.Data.(FirewallSessions))
		case "FirewallStates":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			dh.ProcessFirewallStates(incomingData.Data.(FirewallStates))
		case "Interfaces":
			// Pass the incomingData object over to a function that will process the data and update the Metrics map
			dh.ProcessInterfaces(incomingData.Data.(Interfaces))
		default:
			// Handle unknown response data types if necessary
			fmt.Printf("DataHandler.HandleIncomingData: Unknown response data type: %s\n", incomingData.ResponseDataType)
		}
		dh.DataMutex.Unlock()
		fmt.Printf("DataHandler.HandleIncomingData: Mutex Unlocked.\n")
		// Clear old data from the Metrics and MetricsLastUpdated maps
		fmt.Printf("DataHandler.HandleIncomingData: Total metrics stored: %d\n", len(dh.Metrics))
		//dh.ClearOldData()
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
func (dh *DataHandler) HandleRequests(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("DataHandler.HandleRequests: Mutex Address --> %p\n", dh.DataMutex)

	for request := range dh.Request {

		fmt.Printf("DataHandler.HandleRequests: Received new request\n")

		dh.DataMutex.Lock()
		fmt.Printf("DataHandler.HandleRequests: Mutex Locked.\n")

		switch request {
		case "metrics":
			fmt.Printf("DataHandler.HandleRequests: Creating metrics string.\n")
			metricsString := dh.CreateMetricsString()
			fmt.Printf("DataHandler.HandleRequests: Sending metrics string of length %d.\n", len(metricsString))
			dh.Outgoing <- metricsString
		default:
			fmt.Printf("DataHandler.HandleRequests: Unknown request: %s\n", request)
		}
		dh.DataMutex.Unlock()
		fmt.Printf("DataHandler.HandleRequests: Mutex Unlocked.\n")
	}
}

//-----------------------------------------------------------------------------
// Create a new instance of the DataHandler struct and return a pointer to it
//-----------------------------------------------------------------------------
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

//----------------------------------------------------------------------------
// Create the metrics string to be sent to the web page.
// This will loop over the Metrics map and create a string in the prometheus format.
//----------------------------------------------------------------------------
func (dh *DataHandler) CreateMetricsString() string {
	var metricsString strings.Builder
	firewallIfaceStatistics := 0
	firewallLogEntries := 0

	fmt.Printf("DataHandler.CreateMetricsString: Preparing to loop over metrics\n")

	for metricName, metricValue := range dh.Metrics {
		//fmt.Printf("DataHandler.CreateMetricsString: Processing metric: %s\n", metricName)
		if strings.Contains(metricName, "firewall_interface_statistics") {
			if firewallIfaceStatistics == 0 {
				helpAndTypeLines := dh.CreateHelpAndTypeLines(metricName)
				if helpAndTypeLines != "" {
					metricsString.WriteString(helpAndTypeLines)
				}
				firewallIfaceStatistics++
			}
		} else if strings.Contains(metricName, "firewall_log_entries") {
			if firewallLogEntries == 0 {
				helpAndTypeLines := dh.CreateHelpAndTypeLines(metricName)
				if helpAndTypeLines != "" {
					metricsString.WriteString(helpAndTypeLines)
				}
				firewallLogEntries++
			}
		} else {
			helpAndTypeLines := dh.CreateHelpAndTypeLines(metricName)
			if helpAndTypeLines != "" {
				metricsString.WriteString(helpAndTypeLines)
			}
		}
		metricsString.WriteString(fmt.Sprintf("%s %d\n", metricName, metricValue))
	}

	return metricsString.String()
}

//----------------------------------------------------------------------------
// Create the # HELP and # TYPE lines for the prometheus metrics based on the metric name.
//----------------------------------------------------------------------------
func (dh *DataHandler) CreateHelpAndTypeLines(metricName string) string {
	switch {
		//-------------------------------------------------------------------------------
		// ARP Table Metrics
		//----------------------------------------------------------------------------
		case strings.Contains(metricName, "arp_table_entry"):
			return "# HELP arp_table_entry Total number of entries in the ARP table\n# TYPE arp_table_entry counter\n"
		
		//-----------------------------------------------------------------------
		// Interface Statistics
		//-----------------------------------------------------------------------
		case strings.Contains(metricName, "interface_statistic"):
			measurement := ""
			if strings.Contains(metricName, "packets") {
				measurement = "packets"
			} else if strings.Contains(metricName, "errors") {
				measurement = "errors"
			} else if strings.Contains(metricName, "bytes") {
				measurement = "bytes"
			}

			nameSplit := strings.Split(metricName, "_")
			measurementType := "collisions"
			if len(nameSplit) >= 4 {
				measurementType = nameSplit[3]
			}

			return "# HELP interface_statistic Number of " + measurementType + " " + measurement + " on the interface\n# TYPE interface_statistic counter\n"
		
		//-----------------------------------------------------------------------
		// Firewall Sessions
		//-----------------------------------------------------------------------
		case strings.Contains(metricName, "firewall_session"):
			measurement := ""
			if strings.Contains(metricName, "packets") {
				measurement = "packets"
			} else if strings.Contains(metricName, "bytes") {
				measurement = "bytes"
			}

			return "# HELP firewall_session Total number of " + measurement + " for the firewall session\n# TYPE firewall_session counter\n"

		//-----------------------------------------------------------------------
		// Firewall States
		//-----------------------------------------------------------------------
		case strings.Contains(metricName, "firewall_state"):
			measurement := ""
			if strings.Contains(metricName, "packets") {
				measurement = "packets"
			} else if strings.Contains(metricName, "bytes") {
				measurement = "bytes"
			}
			
			return "# HELP firewall_state Total number of " + measurement + " for the firewall state\n# TYPE firewall_state counter\n"
		
		//-----------------------------------------------------------------------
		// Firewall Interface Statistics
		//-----------------------------------------------------------------------
		case strings.Contains(metricName, "firewall_interface_statistics"):
			return "# HELP firewall_interface_statistics Total number of all statistics for the firewall interface\n# TYPE firewall_interface_statistics counter\n"
		
		//-----------------------------------------------------------------------
		// Firewall Log Entries
		//-----------------------------------------------------------------------
		case strings.Contains(metricName, "firewall_log_entries"):
			return "# HELP firewall_log_entries Total number of log entries for the unique combination of values\n# TYPE firewall_log_entries counter\n"
		
		default:
			return ""
	}
}

