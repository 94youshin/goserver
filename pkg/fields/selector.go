package fields

type Selector interface {
	Matches(fields Fields) bool

	Empty() bool

	RequiresExactMatch(field string) (value string, found bool)

	Transform()

	String() string

	DeepCopySelector() Selector
}

func SelectorFromSet(ls Set) Selector {
	return nil
}
