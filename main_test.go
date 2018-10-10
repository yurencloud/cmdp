package main

import (
	"github.com/yurencloud/cmdp/src"
	"strings"
	"testing"
)

/*
测试前，把下面的参数改改
20181010全测试通过
*/

func TestVersion(t *testing.T) {
	result, _ := src.Exec("cmdp version")
	if !strings.Contains(result, "version") {
		t.Error()
	}
}

func TestSearch(t *testing.T) {
	result, _ := src.Exec("cmdp search ls")
	if !strings.Contains(result, "id:") {
		t.Error()
	}

	result2, _ := src.Exec("cmdp s ls")
	if !strings.Contains(result2, "id:") {
		t.Error()
	}

	result3, _ := src.Exec("cmdp search docker -p 2 -s 5")
	if !strings.Contains(result3, "id:") || !strings.Contains(result3, "size:5, page:2") {
		t.Error()
	}

	result4, _ := src.Exec("cmdp search -a")
	if !strings.Contains(result4, "id:") {
		t.Error()
	}

	result5, _ := src.Exec("cmdp search chenchen/redis")
	if !strings.Contains(result5, "id:") {
		t.Error()
	}

	result6, _ := src.Exec("cmdp search chenchen/")
	if !strings.Contains(result6, "id:") {
		t.Error()
	}

	result7, _ := src.Exec("cmdp d 68")
	if !strings.Contains(result7, "success") {
		t.Error()
	}
}

func TestExec(t *testing.T) {
	result, _ := src.Exec("cmdp exec lsA")
	if !strings.Contains(result, "main.go") {
		t.Error()
	}

	result2, _ := src.Exec("cmdp e lsA")
	if !strings.Contains(result2, "main.go") {
		t.Error()
	}

	result3, _ := src.Exec("cmdp e tom2/lsA")
	if strings.Contains(result3, "success") {
		t.Error()
	}

	result4, _ := src.Exec("cmdp e tom2/lsA -F")
	if !strings.Contains(result4, "success") {
		t.Error()
	}
}


var username = "test3"

func TestRegister(t *testing.T) {
	result, _ := src.Exec("cmdp register "+username+" 123456")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestLogin(t *testing.T) {
	result, _ := src.Exec("cmdp login "+username+" 123456")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestReset(t *testing.T) {
	result, _ := src.Exec("cmdp reset 1234567")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestLogout(t *testing.T) {
	result, _ := src.Exec("cmdp logout")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestLoginMackWang(t *testing.T) {
	result, _ := src.Exec("cmdp login mackwang 123456")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestInfo(t *testing.T) {
	result, _ := src.Exec("cmdp info \"I am mackwang\"")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestPush(t *testing.T) {
	result, _ := src.Exec("cmdp push test.sh test.sh2 test")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestPull(t *testing.T) {
	result, _ := src.Exec("cmdp pull test.sh2")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestFind(t *testing.T) {
	result, _ := src.Exec("cmdp find test")
	if !strings.Contains(result, "id:") {
		t.Error(result)
	}
	result2, _ := src.Exec("cmdp find chenchen/hello")
	if !strings.Contains(result2, "id:") {
		t.Error(result2)
	}

	result3, _ := src.Exec("cmdp find")
	if !strings.Contains(result, "id:") {
		t.Error(result3)
	}

	result4, _ := src.Exec("cmdp find chenchen/")
	if !strings.Contains(result, "id:") {
		t.Error(result4)
	}
}

func TestRemove(t *testing.T) {
	result, _ := src.Exec("cmdp remove 9")
	if !strings.Contains(result, "success") {
		t.Error(result)
	}
}

func TestStar(t *testing.T) {
	result, _ := src.Exec("cmdp star tom2")
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


