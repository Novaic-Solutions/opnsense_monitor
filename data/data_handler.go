package data

import (
	//"fmt"
	//"github.com/Novaic-Solutions/opnsense_monitor/config"
)

// Nameing convention for the prometheus metrics
// prometheus_metric_name{label="value"} value
// 
// For each type of metric, a line will be generated as follows:
// # HELP prometheus_metric_name description of the metric including the measurement units
// # TYPE prometheus_metric_name type of metric (counter, gauge, histogram, summary)


//----------------------------------------------------------------------------
// Metrics for 