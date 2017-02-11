package color

///////////////////////////////////////////////////////
func Black(cc string) string {
    col_export := "\x1b[30m" + cc + "\x1b[0m"
    return col_export
}
func BBlack(cc string) string {
    col_export := "\x1b[30;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
func Red(cc string) string {
    col_export := "\x1b[31m" + cc + "\x1b[0m"
    return col_export
}
func BRed(cc string) string {
    col_export := "\x1b[31;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
func Green(cc string) string {
    col_export := "\x1b[32m" + cc + "\x1b[0m"
    return col_export
}
func BGreen(cc string) string {
    col_export := "\x1b[32;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
func Yellow(cc string) string {
    col_export := "\x1b[33m" + cc + "\x1b[0m"
    return col_export
}
func BYellow(cc string) string {
    col_export := "\x1b[33;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
func Blue(cc string) string {
    col_export := "\x1b[34m" + cc + "\x1b[0m"
    return col_export
}
func BBlue(cc string) string {
    col_export := "\x1b[34;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
func Magenta(cc string) string {
    col_export := "\x1b[35m" + cc + "\x1b[0m"
    return col_export
}
func BMagenta(cc string) string {
    col_export := "\x1b[35;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
func Cyan(cc string) string {
    col_export := "\x1b[36m" + cc + "\x1b[0m"
    return col_export
}
func BCyan(cc string) string {
    col_export := "\x1b[36;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
func White(cc string) string {
    col_export := "\x1b[37m" + cc + "\x1b[0m"
    return col_export
}
func BWhite(cc string) string {
    col_export := "\x1b[37;1m" + cc + "\x1b[0m"
    return col_export
}
///////////////////////////////////////////////////////
