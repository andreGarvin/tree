package allPaths

import "testing"

func TestPaths(t *testing.T) {
		TestCases := []struct {
				in string; excepted int;
		}{
				{"./test", 5},
				{"go", 1},
				{"./test", 1},
		}

		Case, _ := All("./test")
		if pathsLen := len( Case ); pathsLen != TestCases[0].excepted {
				t.Errorf("All func: Failed returned len of '%d' instead of '%d'", len( Case ), TestCases[0].excepted)
		}

		Case1, _ := WithExt("./test", TestCases[1].in)
		if len( Case1 ) != TestCases[1].excepted {
				t.Errorf("WithExt func: Failed returned length of '%d' instead of '%d'", len( Case1 ), TestCases[1].excepted)
		}

		Case2, _ := Dirs("./test")
		if len( Case2 ) != TestCases[2].excepted {
				t.Errorf("WithExt func: Failed returned length of '%d' instead of '%d'", len( Case2 ), TestCases[2].excepted)
		}
}
