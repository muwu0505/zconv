package zconv

import (
	"testing"
)

func TestToCamel(t *testing.T) {
	cases := map[string]string{
		"HiChina":  "HiChina",
		"Hi.China": "HiChina",
		"Hi_China": "HiChina",
		"Hi-China": "HiChina",
		"Hi China": "HiChina",

		"hi.china": "HiChina",
		"hi_china": "HiChina",
		"hi-china": "HiChina",
		"hi china": "HiChina",

		"":           "",
		" Hi China":  "HiChina",
		"  Hi China": "HiChina",
		" Hi  China": "HiChina",
		" hi china":  "HiChina",
	}

	for ori, dst := range cases {
		result := ToCamel(ori)
		if result != dst {
			t.Errorf("ToCamel(%q) = %q, want %q", ori, result, dst)
		}
	}
}

func TestToDotCamel(t *testing.T) {
	cases := map[string]string{
		"HiChina":  "Hi.China",
		"Hi.China": "Hi.China",
		"Hi_China": "Hi.China",
		"Hi-China": "Hi.China",
		"Hi China": "Hi.China",

		"hi.china": "Hi.China",
		"hi_china": "Hi.China",
		"hi-china": "Hi.China",
		"hi china": "Hi.China",

		"":           "",
		" Hi China":  "Hi.China",
		"  Hi China": "Hi.China",
		" Hi  China": "Hi.China",
		" hi china":  "Hi.China",
	}

	for ori, dst := range cases {
		result := ToDotCamel(ori)
		if result != dst {
			t.Errorf("ToDotCamel(%q) = %q, want %q", ori, result, dst)
		}
	}
}

func TestToSpaceCamel(t *testing.T) {
	cases := map[string]string{
		"HiChina":  "Hi China",
		"Hi.China": "Hi China",
		"Hi_China": "Hi China",
		"Hi-China": "Hi China",
		"Hi China": "Hi China",

		"hi.china": "Hi China",
		"hi_china": "Hi China",
		"hi-china": "Hi China",
		"hi china": "Hi China",

		"":           "",
		" Hi China":  "Hi China",
		"  Hi China": "Hi China",
		" Hi  China": "Hi China",
		" hi china":  "Hi China",
	}

	for ori, dst := range cases {
		result := ToSpaceCamel(ori)
		if result != dst {
			t.Errorf("ToDotCamel(%q) = %q, want %q", ori, result, dst)
		}
	}
}
