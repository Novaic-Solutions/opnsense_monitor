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

//----------------------------------------------------------------------------
//	/api/diagnostics/interface/getInterfaceStatistics
//----------------------------------------------------------------------------
type IfaceStatistics struct {
	Statistics map[string]IfaceStatistic
}

type IfaceStatistic struct {
	Name string
	Flags string
	Mtu int
	Network string
	Address string
	ReceivedPackets string
	ReceivedErrors string
	DroppedPackets string
	ReceivedBytes string
	SentPackets string
	SendErrors string
	SentBytes string
	Collisions string
}

//----------------------------------------------------------------------------
// /api/diagnostics/firewall/pf_statistics/interfaces
//----------------------------------------------------------------------------
type IfaceTraffic struct {
	Cleared string
	References int
	In4_pass_packets uint64
	In4_pass_bytes uint64
	In4_block_packets uint64
	In4_block_bytes uint64
	Out4_pass_packets uint64
	Out4_pass_bytes uint64
	Out4_block_packets uint64
	Out4_block_bytes uint64
	In6_pass_packets uint64
	In6_pass_bytes uint64
	In6_block_packets uint64
	In6_block_bytes uint64
	Out6_pass_packets uint64
	Out6_pass_bytes uint64
	Out6_block_packets uint64
	Out6_block_bytes uint64
}


//----------------------------------------------------------------------------
// /api/diagnostics/firewall/pf_statistics/rules
//----------------------------------------------------------------------------
type FirewallRuleTraffic struct {
	Evaluations uint64
	Packets uint64
	Bytes uint64
	States int
	Inserted string
	State_Creations int
	Time string
}

//------------------------------------------------------------------------------
//	/api/diagnostics/firewall/query_states
//------------------------------------------------------------------------------
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

//------------------------------------------------------------------------------
// 	/api/diagnostics/firewall/query_pf_top
//------------------------------------------------------------------------------
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




//------------------------------------------------------------------------------
// 	/api/diagnostics/interface/search_arp/
//------------------------------------------------------------------------------
type ArpTableEntry struct {
	Mac string
	Ip string
	Interface string
	Expired bool
	Expires int
	Permanent bool
	Type string
	Manufacturer string
	Hostname string
	Intf_description string	
}

type DataEntry struct {
	URI string
	Timestamp string
}

type DataRetriever struct {
	
}

//----------------------------------------------------------------------------
// 	/api/diagnostics/firewall/log/
//------------------------------------------------------------------------------
type FirewallLogEntry struct {
	Rulenr string
	Subrulenr string
	Anchorname string
	Rid string
	Interface string
	Reason string
	Action string
	Dir string
	Ipversion string
	Tos string
	Ecn string
	Ttl string
	Id string
	Offset string
	Ipflags string
	Protonum string
	Protoname string
	Length string
	Src string
	Dst string
	Srcport string
	Dstport string
	Datalen string
	Tcpflags string
	Seq string
	Ack string
	Urp string
	Tcpopts string
	Timestamp string
	Host string
	Digest string
	Label string
}

