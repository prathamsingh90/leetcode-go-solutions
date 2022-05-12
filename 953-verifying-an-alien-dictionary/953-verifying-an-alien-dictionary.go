var orderMap = make(map[byte]int)
func sortedWords(word1, word2 string) bool {
    for i:=0;i<len(word1)&&i<len(word2); i++{
        if word1[i]!=word2[i]{
            if orderMap[word1[i]-'a']>orderMap[word2[i]-'a']{
                return false
            } else {
                return true
            }
        }
    }
    
    return len(word1)<=len(word2)
}

func isAlienSorted(words []string, order string) bool {
    
    for i := range order {
        orderMap[order[i]-'a']=i
    }
    
    for i:=0; i<len(words)-1; i++ {
        if !sortedWords(words[i], words[i+1]){
            return false
        }
    }
    
    return true
}