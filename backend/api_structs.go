package main

type IpscanPost struct {
	IfaceName          string `json:"name"`
	IfaceIP            string `json:"ip"`
	Timeout            int64  `json:"timeout"`
	TimeBetweenPackets int64  `json:"timeBetweenPackets"`
}
