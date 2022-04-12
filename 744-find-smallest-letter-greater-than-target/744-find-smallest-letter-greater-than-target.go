func nextGreatestLetter(letters []byte, target byte) byte {
 
    var result byte
    
    result = letters[0]
    
    for i:=0; i<len(letters); i++ {
        if target<letters[i]{
            result = letters[i]
            break
        }
    }
    
    return result
}