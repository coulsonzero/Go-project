package main

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if m1[s[i]] > 0 && m1[s[i]] != t[i] || m2[t[i]] > 0 && m2[t[i]] != s[i] {
			return false
		}

		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}

	return true
}
