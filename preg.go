package  preg

import (
    "regexp"
    "fmt"
)

//////////////////////////////////////////////////////////////

func Match(regex string, subject string) bool {
    r := regexp.MustCompile(regex)
	return r.MatchString(subject)
}	

func MatchList(regex string, subject string) ([]string, bool) {
    res := [][]string{}
    r := regexp.MustCompile(regex)
    ok := r.MatchString(subject)
    if ok {
        res = r.FindAllStringSubmatch(subject, -1)
    }
	return submatch(res), ok
}	

func MatchByte(regex string, subject []byte) bool {
    r := regexp.MustCompile(regex)
	return r.Match(subject)
}	

func MatchListByte(regex string, subject []byte) ([][]byte, bool) {
    res := [][][]byte{}
    r := regexp.MustCompile(regex)
    ok := r.Match(subject)
    if ok {
        res = r.FindAllSubmatch(subject, -1)
    }
	return submatchbyte(res), ok
}	


func ReplaceAll(regex, replace, subject string) string {
    r := regexp.MustCompile(regex)
	subject = r.ReplaceAllString(subject, replace)
	return subject
}

func ReplaceAllList(regex, replace, subject string) (string, []string) {
    res := [][]string{}
    r := regexp.MustCompile(regex)
	s := r.ReplaceAllString(subject, replace)
    res = r.FindAllStringSubmatch(subject, -1)
    return s, submatch(res)
}

func ReplaceAllByte(regex string, replace, subject []byte) []byte {
    r := regexp.MustCompile(regex)
	subject = r.ReplaceAll(subject, replace)
	return subject
}

func ReplaceAllListByte(regex string, replace, subject []byte) ([]byte, [][]byte) {
    res := [][][]byte{}
    r := regexp.MustCompile(regex)
	s := r.ReplaceAll(subject, replace)
    res = r.FindAllSubmatch(subject, -1)
    return s, submatchbyte(res)
}


func Split(regex, subject string) []string {
    r := regexp.MustCompile(regex)
	result := r.Split(subject, -1)
	return result
}


func submatch(mlist [][]string) []string {
    res := []string{}
    // res[0] = regex
    for _,m := range mlist {
        for i := 1; i < len(m); i++ {
            res = append(res, m[i])
        }
    }
    return res
}

func submatchbyte(mlist [][][]byte) [][]byte {
    res := [][]byte{}

    for _,m := range mlist {
        for i := 1; i < len(m); i++ {
            res = append(res, m[i])
        }
    }
    return res
}

func Dump(v interface{}) {
    f := fmt.Sprintf("%#v\n", v)
    f = ReplaceAll("[{}]", " \n", f)
    f = ReplaceAll("[,] ", "\n", f)
    f = ReplaceAll("\n", "\n   ", f)
    fmt.Println(f)
}

