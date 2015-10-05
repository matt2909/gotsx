package gotsx

import (
	"fmt"
	"testing"
)

func TestTSX(t *testing.T) {
	count, test, abort := 0, 0, 0

	for i := 0; i < 1000000; i++ {
		if result := XBegin(XBEGIN_STARTED); result == XBEGIN_STARTED {
			count++
			if XTest() != 0 {
				test++
			}
			if i%150000 == 0 {
				XAbort2()
			}
			XEnd()
		} else if result&XABORT_EXPLICIT == XABORT_EXPLICIT {
			abort++
		}
	}
	fmt.Printf("count %d, test %d, abort %d\n", count, test, abort)
	if count != test {
		t.Errorf("TSX count(%d) should equal test(%d)", count, test)
	}
	if abort == 0 {
		t.Errorf("TSX error(%d) should be greater than 0", abort)
	}
}
