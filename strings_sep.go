package zconv

const (
	StringsSepSnack = "_"
	StringsSepKebab = "-"
	StringsSepDot   = "."
	StringsSepSpace = " "
)

var StringsSepMap = map[string]string{
	"snake": StringsSepSnack,
	"kebab": StringsSepKebab,
	"dot":   StringsSepDot,
	"space": StringsSepSpace,
}

func StringsSep(seps ...string) map[string]byte {
	if len(seps) <= 0 {
		seps = []string{StringsSepSnack, StringsSepKebab, StringsSepDot, StringsSepSpace}
	}

	sepMap := make(map[string]byte, len(seps))
	for _, sep := range seps {
		sepMap[sep] = 0
	}
	return sepMap
}
