func groupAnagrams(strs []string) [][]string {
    ret := [][]string{}

    shrinkingList := strs
    for len(shrinkingList) > 0 { // go until we've pulled everything out
        target := shrinkingList[0]
        anagramsForTarget := []string{}
        anagramsForTarget = append(anagramsForTarget, target)
        shrinkingList = shrinkingList[1:]
        
        i := 0
        for i < len(shrinkingList) {
            check := shrinkingList[i]
            if isAnagram(target, check) {
                anagramsForTarget = append(anagramsForTarget, check)
                shrinkingList = append(shrinkingList[:i], shrinkingList[i+1:]...)
            } else {
                i++
            }
        }

        ret = append(ret, anagramsForTarget)
    }

    return ret
}

func isAnagram(str1 string, str2 string) bool {
    if len(str1) != len(str2) {
        return false 
    }
    m1 := make(map[byte]int)
    m2 := make(map[byte]int)

    for i := 0; i < len(str1); i++ {
        freq, found := m1[str1[i]]
        if !found {
            m1[str1[i]] = 1
        } else {
            m1[str1[i]] = freq + 1
        }
    }

    for i := 0; i < len(str2); i++ {
        freq, found := m2[str2[i]]
        if !found {
            m2[str2[i]] = 1
        } else {
            m2[str2[i]] = freq + 1
        }
    }

    if len(m1) != len(m2) {
        return false
    }

    for key1, value1 := range m1 {
        value2, found := m2[key1]
        if !found || value1 != value2 {
            return false
        }
    }
    return true
}