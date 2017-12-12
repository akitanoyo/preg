
package preg

import "testing"
import "io/ioutil"

func TestMatch(t *testing.T) {
    s := "http://github.com/akitanoyo/preg"
    r := "//\\w+\\.\\w+/"
    if !Match(r, s) {
        t.Errorf("got preg.Match %v %v",
            s, r)
    }
}

func TestMatchList(t *testing.T) {
    s := "http://github.com/akitanoyo/preg"
    r := "//(\\w+\\.\\w+)/"
    ms, ok := MatchList(r, s)
    if !ok {
        t.Errorf("got preg.MatchList %v %v",
            s, r)
    }

    if len(ms) == 0 {
       t.Errorf("got no preg.MatchList %v no match", r) 
    } else if len(ms) != 1 || ms[0] != "github.com" {
        t.Errorf("got preg.MatchList %v != %v",
            ms[0], "github.com")
    }
}

func TestReplaceAll(t *testing.T) {
    s := "http://github.com/akitanoyo/preg"
    r := "/\\w+"
    s = ReplaceAll(r, "X", s)
    if s != "http:/X.comXX"  {
        t.Errorf("got preg.ReplaceAll %v != %v",
            s, "http:/X.comXX")
    }
}

func TestReplaceAllList(t *testing.T) {
    s := "http://github.com/akitanoyo/preg"
    r := "/(\\w+)"
    s, res := ReplaceAllList(r, "X",s)
    if s != "http:/X.comXX"  {
        t.Errorf("got preg.ReplaceAllList %v != %v",
            s, "http:/X.comXX")
    }
    if len(res) != 3 {
        t.Errorf("got preg.ReplaceAllList no 3match %v != %v",
            len(res), "3")
    } else if res[2] != "preg" {
        t.Errorf("got preg.ReplaceAllList third match %v != %v",
            res[2], "preg")
    }
}

func TestSplit(t *testing.T) {
    s := "http://github.com/akitanoyo/preg"
    r := "/."
    ms := Split(r, s)
    if len(ms) != 4 {
        t.Errorf("got preg.Split() splits into %d != %d",
            len(ms), 4)
    }
    if ms[3] != "reg" {
        t.Errorf("got preg.Split() splits of fourth %v != %v",
            ms[3], "reg")
    }
}

func TestMatchByte(t *testing.T) {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        t.Errorf("got test file not found! test.txt")
    }

    r := "//\\w+\\.\\w+/"
    if !MatchByte(r, data) {
        t.Errorf("got test preg.MatchByte no mach %s", r)
    }
}

func TestMatchListByte(t *testing.T) {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        t.Errorf("got test file not found! test.txt")
    }

    // r := "//\\w+\\.\\w+/" // check error case
    r := "//(\\w+\\.\\w+)/"
    ms, ok := MatchListByte(r, data)
    if !ok {
        t.Errorf("got preg.MatchList %v", r)
    }

    if len(ms) == 0 {
        t.Errorf("got no preg.MatchListByte %v no match", r)
    } else if len(ms) != 1 || string(ms[0]) != "github.com" {
        t.Errorf("got preg.MatchList %v != %v (len %v)",
            ms[0], "github.com", len(ms))
    }
}

func TestReplaceAllByte(t *testing.T) {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        t.Errorf("got test file not found! test.txt")
    }
    r := "/\\w+"
    s := ReplaceAllByte(r, []byte("X"), data)

    if string(s) != "http:/X.comXX\n"  {
        t.Errorf("got preg.ReplaceAll %v != %v",
            s, []byte("http:/X.comXX\n"))
    }
}

func TestReplaceAllListByte(t *testing.T) {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        t.Errorf("got test file not found! test.txt")
    }

    // r := "//(\\w+)" // check error case
    r := "/(\\w+)"
    s, res := ReplaceAllListByte(r, []byte("X"), data)
    if string(s) != "http:/X.comXX\n"  {
        t.Errorf("got preg.ReplaceAllList %v != %v",
            s, "http:/X.comXX")
    }
    if len(res) != 3 {
        t.Errorf("got preg.ReplaceAllList no 3match %v != %v",
            len(res), "3")
    } else if string(res[2]) != "preg" {
        t.Errorf("got preg.ReplaceAllList third match %v != %v",
            res[2], "preg")
    }
}

