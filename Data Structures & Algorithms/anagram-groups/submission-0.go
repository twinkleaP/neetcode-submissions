func groupAnagrams(strs []string) [][]string {

	result := [][]string{}
	m := make(map[[26]int][]string)

	for _, s := range strs{
		var count[26]int


		for i:= 0; i < len(s); i++{
			count[s[i]-'a']++
		}

		m[count]=append(m[count],s)
	}

	for _, group := range m{
		result = append(result, group)
	}
  return result
}
