func defangIPaddr(address string) string {
    var defangIP strings.Builder
    
    for _, value := range address {
        if value == '.'{
            defangIP.WriteString("[.]")
        } else {
            defangIP.WriteString(string(value))
        }
    }
    
    return defangIP.String()
}