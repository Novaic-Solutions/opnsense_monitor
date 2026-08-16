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
	Statistics map[string]IfaceStatistic `json:"statistics"`
}

type IfaceStatistic struct {
	Name string `json:"name"`
	Flags string `json:"flags"`
	Mtu int `json:"mtu"`
	Network string `json:"network"`
	Address string `json:"address"`
	ReceivedPackets uint64 `json:"received-packets"`
	ReceivedErrors uint64 `json:"received-errors"`
	DroppedPackets uint64 `json:"dropped-packets"`
	ReceivedBytes uint64 `json:"received-bytes"`
	SentPackets uint64 `json:"sent-packets"`
	SendErrors uint64 `json:"send-errors"`
	SentBytes uint64 `json:"sent-bytes"`
	Collisions uint64 `json:"collisions"`
}

//----------------------------------------------------------------------------
// /api/diagnostics/firewall/pf_statistics/interfaces
//----------------------------------------------------------------------------
type Interfaces struct {
	Interfaces map[string]IfaceTraffic `json:"interfaces"`
}

type IfaceTraffic struct {
	Cleared string `json:"cleared"`
	References int `json:"references"`
	In4_pass_packets uint64 `json:"in4_pass_packets"`
	In4_pass_bytes uint64 `json:"in4_pass_bytes"`
	In4_block_packets uint64 `json:"in4_block_packets"`
	In4_block_bytes uint64 `json:"in4_block_bytes"`
	Out4_pass_packets uint64 `json:"out4_pass_packets"`
	Out4_pass_bytes uint64 `json:"out4_pass_bytes"`
	Out4_block_packets uint64 `json:"out4_block_packets"`
	Out4_block_bytes uint64 `json:"out4_block_bytes"`
	In6_pass_packets uint64 `json:"in6_pass_packets"`
	In6_pass_bytes uint64 `json:"in6_pass_bytes"`
	In6_block_packets uint64 `json:"in6_block_packets"`
	In6_block_bytes uint64 `json:"in6_block_bytes"`
	Out6_pass_packets uint64 `json:"out6_pass_packets"`
	Out6_pass_bytes uint64 `json:"out6_pass_bytes"`
	Out6_block_packets uint64 `json:"out6_block_packets"`
	Out6_block_bytes uint64 `json:"out6_block_bytes"`
}

//----------------------------------------------------------------------------
// /api/diagnostics/firewall/pf_statistics/rules
//----------------------------------------------------------------------------
type FirewallRulesStatistics struct {
	FilterRules map[string]FirewallRuleTraffic `json:"filter_rules"`
	NatRules map[string]FirewallRuleTraffic `json:"nat_rules"`
}

type FirewallRuleTraffic struct {
	Evaluations uint64 `json:"evaluations"`
	Packets uint64 `json:"packets"`
	Bytes uint64 `json:"bytes"`
	States int `json:"states"`
	Inserted string `json:"inserted"`
	State_Creations int `json:"state_creations"`
	Time string `json:"time"`
}

//------------------------------------------------------------------------------
//	/api/diagnostics/firewall/query_states
//------------------------------------------------------------------------------
type FirewallStates struct {
	Rows []FirewallState `json:"rows"`
	Total int `json:"total"`
	Current int `json:"current"`
	RowCount int `json:"row_count"`
}

type FirewallState struct {
	Label string `json:"label"`
	Descr string `json:"descr"`
	Nat_addr string `json:"nat_addr"`
	Nat_port string `json:"nat_port"`
	Gateway string `json:"gateway"`
	Proto string `json:"proto"`
	Ipproto string `json:"ipproto"`
	Flags []string `json:"flags"`
	Direction string `json:"direction"`
	Dst_addr string `json:"dst_addr"`
	Dst_port string `json:"dst_port"`
	Src_addr string `json:"src_addr"`
	Src_port string `json:"src_port"`
	State string `json:"state"`
	Age string `json:"age"`
	Expires string `json:"expires"`
	Pkts []int `json:"pkts"`
	Bytes []int `json:"bytes"`
	Rule string `json:"rule"`
	Id string `json:"id"`
	Interface string `json:"interface"`
	Route_to string `json:"route-to"`
}

//------------------------------------------------------------------------------
// 	/api/diagnostics/firewall/query_pf_top
//------------------------------------------------------------------------------
type FirewallSessions struct {
	Total int `json:"total"`
	Current int `json:"current"`
	RowCount int `json:"row_count"`
	Rows []FirewallSession `json:"rows"`
}

type FirewallSession struct {
	Proto string `json:"proto"`
	Dir string `json:"dir"`
	Src_addr string `json:"src_addr"`
	Src_port string `json:"src_port"`
	Dst_addr string `json:"dst_addr"`
	Dst_port string `json:"dst_port"`
	Gw_addr string `json:"gw_addr"`
	Gw_port string `json:"gw_port"`
	State string `json:"state"`
	Age int `json:"age"`
	Expire int `json:"expire"`
	Pkts int `json:"pkts"`
	Bytes int `json:"bytes"`
	Avg int `json:"avg"`
	Rule string `json:"rule"`
	Label string `json:"label"`
	Descr string `json:"descr"`
}

//------------------------------------------------------------------------------
// 	/api/diagnostics/interface/search_arp/
//------------------------------------------------------------------------------
type ArpTable struct {
	Total int `json:"total"`
	Current int `json:"current"`
	RowCount int `json:"row_count"`
	Rows []ArpTableEntry `json:"rows"`
}

type ArpTableEntry struct {
	Mac string `json:"mac"`
	Ip string `json:"ip"`
	Interface string `json:"intf"`
	Expired bool `json:"expired"`
	Expires int `json:"expires"`
	Permanent bool `json:"permanent"`
	Type string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Hostname string `json:"hostname"`
	Intf_description string `json:"intf_description"`
}

//----------------------------------------------------------------------------
// 	/api/diagnostics/firewall/log/
//------------------------------------------------------------------------------

type FirewallLogEntry struct {
	Rulenr string `json:"rulenr"`
	Subrulenr string `json:"subrulenr"`
	Anchorname string `json:"anchorname"`
	Rid string `json:"rid"`
	Interface string `json:"interface"`
	Reason string `json:"reason"`
	Action string `json:"action"`
	Dir string `json:"dir"`
	Ipversion string `json:"ipversion"`
	Tos string `json:"tos"`
	Ecn string `json:"ecn"`
	Ttl string `json:"ttl"`
	Id string `json:"id"`
	Offset string `json:"offset"`
	Ipflags string `json:"ipflags"`
	Protonum string `json:"protonum"`
	Protoname string `json:"protoname"`
	Length string `json:"length"`
	Src string `json:"src"`
	Dst string `json:"dst"`
	Srcport string `json:"srcport"`
	Dstport string `json:"dstport"`
	Datalen string `json:"datalen"`
	Tcpflags string `json:"tcpflags"`
	Seq string `json:"seq"`
	Ack string `json:"ack"`
	Urp string `json:"urp"`
	Tcpopts string `json:"tcpopts"`
	Timestamp string `json:"__timestamp__"`
	Host string `json:"__host__"`
	Digest string `json:"__digest__"`
	Label string `json:"label"`
}



//------------------------------------------------------------------------------
//  Struct for holding all of the data from API Responses and for 
//  updating them.  Also will be used to create the data that is sent to the web 
// 	page for display.  This will be updated by the server with data from the 
//  clients that is gathered from the channel.
//------------------------------------------------------------------------------
type OpnSenseData struct {
	
}