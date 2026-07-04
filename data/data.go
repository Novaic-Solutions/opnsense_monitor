package data

//------------------------------------------------------------------------------
//	Connects to the database and will retrieve data from a channel for 
//  data from API responses and enter it into the database.
//  A separate channel will be used to send data back to the main application for
//  display on the web page.
//------------------------------------------------------------------------------

type ResponseData struct {
	rowCount int
	total int
	current int
	rows []any
}

type FirewallState struct {
	Label string
	Descr string
	Nat_addr string
	Nat_port string
	Gateway string
	Proto string
	Flags []string
	Direction string
	Dst_addr string
	Dst_port string
	Src_addr string
	Src_port string
	State string
	Age string
	Expires string
	Pkts []int
	Bytes []int
	Rule string
	Id string
	Interface string
}

type FirewallSession struct {
	Proto string
	Dir string
	Src_addr string
	Src_port string
	Dst_addr string
	Dst_port string
	Gw_addr string
	Gw_port string
	Age int
	Expire int
	Pkts int
	Bytes int
	Avg int
	Rule string
	Label string
	Descr string
}

type DataEntry struct {
	URI string
	Timestamp string
}

type DataRetriever struct {
	
}