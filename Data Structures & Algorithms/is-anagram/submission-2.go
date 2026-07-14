func isAnagram(s string, t string) bool {
    m1 := make(map[byte]int)
    m2 := make(map[byte]int)

    for i:=0; i < len(s); i++ {
        current := s[i]
        freq, found := m1[current]
        if !found {
            m1[current] = 1
        } else {
            m1[current] = freq + 1
        }
    }

    for i:=0; i < len(t); i++ {
        current := t[i]
        freq, found := m2[current]
        if !found {
            m2[current] = 1
        } else {
            m2[current] = freq + 1
        }
    }

    if len(m1) != len(m2) {
        return false 
    }

    for m1key, m1freq := range m1 {
        m2freq, m2found := m2[m1key]
        if !m2found || m2freq != m1freq {
            return false
        }
    }

    return true
}
