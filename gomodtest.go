package gomodtest

func EchoMirror(myStr string) string {
    mirror := ""
    for i := len(myStr)-1; i >= 0; i-- {
        mirror += string(myStr[i])
    }
    return mirror
}