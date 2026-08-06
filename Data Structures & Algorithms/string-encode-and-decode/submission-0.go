type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    encoded := ""
    for _, val := range strs {
        length := len(val)
        encoded = encoded + strconv.Itoa(length)
        encoded = encoded + "#"
        encoded = encoded + val
    }

    return encoded
}

func (s *Solution) Decode(encoded string) []string {
    decoded := []string{}
    i := 0
    lenOfStr := ""
    for i < len(encoded) {
        charAt := string(encoded[i])
        if charAt == "#" {
            // found a delimeter
            l, _ := strconv.Atoi(lenOfStr)
            str := encoded[i+1:i+1+l]
            decoded = append(decoded, str)
            i = i+1+l
            lenOfStr = ""
        } else {
            lenOfStr = lenOfStr + string(encoded[i])
            i++
        }
    }

    return decoded

}
