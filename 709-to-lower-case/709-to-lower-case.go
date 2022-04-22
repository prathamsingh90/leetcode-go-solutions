func toLowerCase(s string) string {
    
    var result strings.Builder
    
    for _, charac := range s{
        if unicode.IsUpper(charac){
            result.WriteRune(unicode.ToLower(charac))
        } else {
        result.WriteRune(charac)
        }
    }
    
    return result.String()
}