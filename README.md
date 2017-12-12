# preg

package preg

func Dump(v interface{})  
func Match(regex string, subject string) bool  
func MatchByte(regex string, subject []byte) bool  
func MatchList(regex string, subject string) ([]string, bool)  
func MatchListByte(regex string, subject []byte) ([][]byte, bool)  
func ReplaceAll(regex, replace, subject string) string  
func ReplaceAllByte(regex string, replace, subject []byte) []byte  
func ReplaceAllList(regex, replace, subject string) (string, []string)  
func ReplaceAllListByte(regex string, replace, subject []byte) ([]byte, [][]byte)  
func Split(regex, subject string) []string  

