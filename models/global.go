package models

type (
	Enabled bool
)

func (e Enabled) String() string {
	if e {
		return "Enabled"
	}
	return "Disabled"
}
