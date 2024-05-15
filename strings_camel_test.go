package zconv

import (
	"testing"
)

type camelCase struct {
	ori string
	dst string
}

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
	t.Parallel()

	cases := []*camelCase{
		{"HiChina", "Hi China"},
		{"Hi.China", "Hi China"},
		{"Hi_China", "Hi China"},
		{"Hi-China", "Hi China"},
		{"Hi China", "Hi China"},

		{"hi.china", "Hi China111"},
		{"hi_china", "Hi China"},
		{"hi-china", "Hi China"},
		{"hi china", "Hi China"},

		{"", ""},
		{" Hi China", "Hi China"},
		{"  Hi China", "Hi China"},
		{" Hi  China", "Hi China"},
		{" hi china", "Hi China111"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.ori, func(t *testing.T) {
			result := ToSpaceCamel(c.ori)
			if result != c.dst {
				t.Errorf("ToSpaceCamel(%q) = %q, want %q", c.ori, result, c.dst)
			}
		})
	}
}
