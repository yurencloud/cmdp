package main

import (
	"github.com/yurencloud/cmdp/src"
	"strings"
	"testing"
)

/*
测试前，把下面的参数改改
20181010全测试通过
20190807全测试通过
*/

// 每次测试需要数字加1

var TomName = "tom3"

func TestVersion(t *testing.T) {
	result, _ := src.Exec("cmdp version")
	if !strings.Contains(result, "version") {
		t.Error()
	}
}

func TestRegisterAndLogin(t *testing.T) {
	result8, _ := src.Exec("cmdp register " + TomName + " 123456")
	if !strings.Contains(result8, "success") {
		t.Error(result8)
	}

	result9, _ := src.Exec("cmdp login " + TomName + " 123456")
	if !strings.Contains(result9, "success") {
		t.Error(result9)
	}
}

func TestCmd(t *testing.T) {
	result8, _ := src.Exec("cmdp c \"ls -a\" ls-a -p")
	if !strings.Contains(result8, "success") {
		t.Error()
	}

	result9, _ := src.Exec("cmdp fc mackwang/go-build-mac")
	if !strings.Contains(result9, "success") {
		t.Error()
	}

	result, _ := src.Exec("cmdp search ls")
	if !strings.Contains(result, "id:") {
		t.Error()
	}

	result2, _ := src.Exec("cmdp s ls")
	if !strings.Contains(result2, "id:") {
		t.Error()
	}

	result3, _ := src.Exec("cmdp search ls -p 1 -s 5")
	if !strings.Contains(result3, "id:") || !strings.Contains(result3, "size:5, page:1") {
		t.Error()
	}

	result4, _ := src.Exec("cmdp search -a")
	if !strings.Contains(result4, "id:") {
		t.Error()
	}

	result5, _ := src.Exec("cmdp search mackwang/mysql")
	if !strings.Contains(result5, "id:") {
		t.Error()
	}

	result6, _ := src.Exec("cmdp search mackwang/")
	if !strings.Contains(result6, "id:") {
		t.Error()
	}

	result7, _ := src.Exec("cmdp d 68")
	if !strings.Contains(result7, "success") {
		t.Error()
	}
}

func TestExec(t *testing.T) {
	result, _ := src.Exec("cmdp exec ls-a")
	if !strings.Contains(result, "main.go") {
		t.Error()
	}

	result2, _ := src.Exec("cmdp e ls-a")
	if !strings.Contains(result2, "main.go") {
		t.Error()
	}

	result3, _ := src.Exec("cmdp e mackwang/ls-a")
	if strings.Contains(result3, "success") {
		t.Error()
	}

	result4, _ := src.Exec("cmdp e mackwang/ls-a -F")
	if !strings.Contains(result4, "success") {
		t.Error()
	}
}

func TestUser(t *testing.T) {
	result, _ := src.Exec("cmdp reset 1234567")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}

	result2, _ := src.Exec("cmdp logout")
	if !strings.Contains(result2, "success") {
		t.Error(result2)
	}

	result3, _ := src.Exec("cmdp login  " + TomName + " 1234567")
	if !strings.Contains(result3, "success") {
		t.Error(result3)
	}

	result4, _ := src.Exec("cmdp info \"I am mackwang\"")
	if !strings.Contains(result4, "success") {
		t.Error(result4)
	}
}

func TestFile(t *testing.T) {
	result, _ := src.Exec("cmdp push test.sh test.sh2 test -p")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}

	result2, _ := src.Exec("cmdp pull test.sh2")
	if !strings.Contains(result2, "success") {
		t.Error(result2)
	}

	result3, _ := src.Exec("cmdp find test")
	if !strings.Contains(result3, "id:") {
		t.Error(result3)
	}
	result4, _ := src.Exec("cmdp find mackwang/test")
	if !strings.Contains(result4, "id:") {
		t.Error(result4)
	}

	result5, _ := src.Exec("cmdp find")
	if !strings.Contains(result5, "id:") {
		t.Error(result5)
	}

	result6, _ := src.Exec("cmdp find mackwang/")
	if !strings.Contains(result6, "id:") {
		t.Error(result6)
	}

	result7, _ := src.Exec("cmdp ff mackwang/webpack")
	if !strings.Contains(result7, "success") {
		t.Error(result7)
	}

	result8, _ := src.Exec("cmdp remove 9")
	if !strings.Contains(result8, "success") {
		t.Error(result8)
	}
}

func TestStar(t *testing.T) {
	result, _ := src.Exec("cmdp star mackwang")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
	result2, _ := src.Exec("cmdp star")
	if !strings.Contains(result2, "id:") {
		t.Error(result2)
	}
	result3, _ := src.Exec("cmdp star -d 6")
	if !strings.Contains(result3, "success") {
		t.Error(result3)
	}
}
