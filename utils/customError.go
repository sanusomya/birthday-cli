package utils

type FlagMissing struct{}

func (FlagMissing) Error() string {
	return "all flags are necessary"
}
