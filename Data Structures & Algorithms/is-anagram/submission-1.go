func isAnagram(s string, t string) bool {

	count := make(map[rune]int)

	if len(s)!=len(t){
		return false
	}

	for _, ch := range s{
		count[ch]++
	}

	for _, ch := range t{

		count[ch]--

		if count[ch]<0{
			return false
		}
	}

	return true

}
