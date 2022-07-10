package enums

type Education int8

const (
	Primary = iota
	LowerSecondary
	UpperSecondary
	PostSecondary
	ShortCycleTetriary
	Bachelor
	Master
	Doctorate

)

func (status Education) String() string {
	switch status {
	case Primary:
		return "Primary education"
	case LowerSecondary:
		return "Lower secondary education"
	case UpperSecondary:
		return "Upper secondary education"
	case PostSecondary:
		return "Post Secondary education"
	case ShortCycleTetriary:
		return "Short-cycle Tetriary education"
	case Bachelor:
		return "Bachelor"
	case Master:
		return "Master"
	case Doctorate:
		return "Doctorate"
	}
	return "Unknown"
}
