package colors

const (
	Reset    = "\033[0m"
	Red      = "\033[31m"
	Green    = "\033[32m"
	Yellow   = "\033[33m"
	Orange   = "\033[93m"
	Blue     = "\033[34m"
	DarkBlue = "\033[94m"
	Purple   = "\033[35m"
	Cyan     = "\033[36m"
	Gray     = "\033[37m"
	DarkGray = "\033[90m"
	White    = "\033[97m"
)

func SetColor(text string, col string) string {
	return col + text + Reset
}
