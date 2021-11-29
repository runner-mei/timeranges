package timerangesub

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTimeRangesSubtract(t *testing.T) {
	subtractor := [][]int{{900, 1000}}
	minuend := [][]int{{900, 1000}}
	expectStr := ""
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 930}}
	minuend = [][]int{{830, 900}}
	expectStr = "0900-0930"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 930}}
	minuend = [][]int{{1030, 1100}}
	expectStr = "0900-0930"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 930}}
	minuend = [][]int{{830, 915}}
	expectStr = "0915-0930"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 930}}
	minuend = [][]int{{915, 1000}}
	expectStr = "0900-0915"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 930}}
	minuend = [][]int{{815, 1000}}
	expectStr = ""
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 930}}
	minuend = [][]int{{930, 1500}}
	expectStr = "0900-0930"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 930}, {1000, 1030}}
	minuend = [][]int{{915, 1015}}
	expectStr = "0900-0915,1015-1030"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 1100}, {1300, 1500}}
	minuend = [][]int{{900, 915}, {1000, 1015}, {1030, 1045}, {1230, 1600}}
	expectStr = "0915-1000,1015-1030,1045-1100"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}

	subtractor = [][]int{{900, 1100}, {1300, 1500}}
	minuend = [][]int{{1030, 1045}, {900, 915}, {1000, 1015}, {1230, 1600}}
	expectStr = "0915-1000,1015-1030,1045-1100"
	if res := toStr(TimeRangesSubtract(subtractor, minuend)); res != expectStr {
		t.Errorf("expect %s, but got %s", expectStr, res)
	}
}

func toStr(ranges [][]int) string {
	buf := bytes.Buffer{}
	length := len(ranges)
	for idx, value := range ranges {
		buf.WriteString(fmt.Sprintf("%04d-%04d", value[0], value[1]))
		if idx < length-1 {
			buf.WriteString(",")
		}
	}

	return buf.String()
}
