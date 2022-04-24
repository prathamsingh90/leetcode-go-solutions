func checkAlmostEquivalent(word1 string, word2 string) bool {
    count := make([]int, 26)
    
    for _, word := range word1{
        count[word-'a']++
    }
    
    for _, word := range word2{
        count[word-'a']--
    }
    
    for _, val := range count {
        if int(math.Abs(float64(val)))>3{
            return false
        }
    }
    
    return true
}