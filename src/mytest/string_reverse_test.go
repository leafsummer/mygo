package strev

import "testing"
import "./string_reverse"

type ReverseTest struct {
    in, out string
}

var ReverseTests = []ReverseTest {
    ReverseTest{"ABCD", "DCBA"},
    ReverseTest{"CVO-AZ", "ZA-OVC"},
    ReverseTest{"Hello 世界", "界世 olleH"},
}

func TestReverse(t *testing.T) {
    for _, r := range ReverseTests{
        exp := strev.Reverse(r.in)
        if r.out != exp {
            t.Error("Reverse of %s expects %s, but got %s", r.in, exp, r.out)
        }
    }
}

func BenchmarkReverse(b *testing.B){
    s := "ABCD"
    for i := 0; i < b.N; i++ {
        strev.Reverse(s)
    }
}