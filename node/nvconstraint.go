package node

// import "fmt"

// type (
// 	NVConstraintType int
// 	NVConstraintInfo string
// )

// const (
// 	Integer NVConstraintType = iota
// 	Positive
// 	NonNegative
// 	NyquistLimited

// 	IntegerStr        = "Integer"
// 	PositiveStr       = "Positive"
// 	NonNegativeStr    = "NonNegative"
// 	NyquistLimitedStr = "NyquistLimited"
// )

// func (t NVConstraintType) Info() NVConstraintInfo {
// 	switch t {
// 	case Integer:
// 		return IntegerStr
// 	case Positive:
// 		return PositiveStr
// 	case NonNegative:
// 		return NonNegativeStr
// 	case NyquistLimited:
// 		return NyquistLimitedStr
// 	}

// 	panic(fmt.Sprintf("unexpected type %d", t))
// }

// func (info NVConstraintInfo) Type() NVConstraintType {
// 	switch info {
// 	case IntegerStr:
// 		return Integer
// 	case PositiveStr:
// 		return Positive
// 	case NonNegativeStr:
// 		return NonNegative
// 	case NyquistLimitedStr:
// 		return NyquistLimited
// 	}

// 	panic(fmt.Sprintf("unexpected info %s", info))
// }
