package main

import (
	"strconv"
	"strings"
)

func byteToIpv4String(byteSlice []byte) string {
	return strconv.Itoa(int(byteSlice[0])) + "." +
		strconv.Itoa(int(byteSlice[1])) + "." +
		strconv.Itoa(int(byteSlice[2])) + "." +
		strconv.Itoa(int(byteSlice[3]))
}

func byteToMACString(byteSlice []byte) string {
	return strconv.FormatInt(int64(byteSlice[0]), 16) + "-" +
		strconv.FormatInt(int64(byteSlice[1]), 16) + "-" +
		strconv.FormatInt(int64(byteSlice[2]), 16) + "-" +
		strconv.FormatInt(int64(byteSlice[3]), 16) + "-" +
		strconv.FormatInt(int64(byteSlice[4]), 16) + "-" +
		strconv.FormatInt(int64(byteSlice[5]), 16)
}

func stringToByteIP(ipString string) []byte {
	slicesIp := strings.Split(ipString, ".")
	returnbyte := make([]byte, 4)
	for i, slice := range slicesIp {
		test, _ := strconv.Atoi(slice)
		returnbyte[i] = byte(test)
	}
	return returnbyte
}
