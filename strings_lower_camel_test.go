package zconv

import (
	"testing"
)

func TestToLowerCamel(t *testing.T) {
	cases := map[string]string{
		"HiChina":  "hiChina",
		"Hi.China": "hiChina",
		"Hi_China": "hiChina",
		"Hi-China": "hiChina",
		"Hi China": "hiChina",

		"hi.china": "hiChina",
		"hi_china": "hiChina",
		"hi-china": "hiChina",
		"hi china": "hiChina",

		"":           "",
		" Hi China":  "hiChina",
		"  Hi China": "hiChina",
		" Hi  China": "hiChina",
		" hi china":  "hiChina",
	}

	for ori, dst := range cases {
		result := ToLowerCamel(ori)
		if result != dst {
			t.Errorf("ToLowerCamel(%q) = %q, want %q", ori, result, dst)
		}
	}
}

func TestToLowerDotCamel(t *testing.T) {
	cases := map[string]string{
		"HiChina":  "hi.China",
		"Hi.China": "hi.China",
		"Hi_China": "hi.China",
		"Hi-China": "hi.China",
		"Hi China": "hi.China",

		"hi.china": "hi.China",
		"hi_china": "hi.China",
		"hi-china": "hi.China",
		"hi china": "hi.China",

		"":           "",
		" Hi China":  "hi.China",
		"  Hi China": "hi.China",
		" Hi  China": "hi.China",
		" hi china":  "hi.China",
	}

	for ori, dst := range cases {
		result := ToLowerDotCamel(ori)
		if result != dst {
			t.Errorf("ToLowerDotCamel(%q) = %q, want %q", ori, result, dst)
		}
	}
}

func TestToLowerSpaceCamel(t *testing.T) {
	cases := map[string]string{
		"HiChina":  "hi China",
		"Hi.China": "hi China",
		"Hi_China": "hi China",
		"Hi-China": "hi China",
		"Hi China": "hi China",

		"hi.china": "hi China",
		"hi_china": "hi China",
		"hi-china": "hi China",
		"hi china": "hi China",

		"":           "",
		" Hi China":  "hi China",
		"  Hi China": "hi China",
		" Hi  China": "hi China",
		" hi china":  "hi China",
	}

	for ori, dst := range cases {
		result := ToLowerSpaceCamel(ori)
		if result != dst {
			t.Errorf("ToLowerSpaceCamel(%q) = %q, want %q", ori, result, dst)
		}
	}
}
