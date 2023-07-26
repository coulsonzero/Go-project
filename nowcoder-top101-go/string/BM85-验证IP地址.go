package main

import "strings"
import "strconv"

/**
 * 验证IP地址
 * @param IP string字符串 一个IP地址字符串
 * @return string字符串
 */
func solve(IP string) string {
	// write code here
	switch {
	case strings.Contains(IP, "."):
		return isIPv4(IP)
	case strings.Contains(IP, ":"):
		return isIPv6(IP)
	default:
		return "Neither"
	}
}

// IPv4只有十进制数和分割点，其中数字在0-255之间，共4组，且不能有零开头的非零数，不能缺省
func isIPv4(ip string) string {
	arr := strings.Split(ip, ".")
	for _, v := range arr {
		if v > "255" || v < "0" || strings.HasPrefix(v, "0") || len(v) == 0 || len(v) > 3 || len(arr) != 4 {
			return "Neither"
		}
	}
	return "IPv4"
}

// IPv6由8组16进制数组成，会出现a-fA-F，通过冒号分割，不可缺省，每组最多4位
func isIPv6(ip string) string {
	arr := strings.Split(ip, ":")
	for _, v := range arr {
		if len(arr) != 8 || len(v) == 0 || len(v) > 4 || !isValid16(v) {
			return "Neither"
		}
	}
	return "IPv6"
}

func isValid16(s string) bool {
	if _, err := strconv.ParseUint(s, 16, 64); err != nil {
		return false
	}
	return true

	// for _, c := range s {
	//     if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
	//         return false
	//     }
	// }
	// return true
}
