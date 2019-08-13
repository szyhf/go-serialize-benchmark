package easyjson

type ValMsg struct {
	Int         int
	Str         string
	Bool        bool
	Struct      *HoldValMsg
	ByteSlice   []byte
	BoolSlice   []bool
	IntSlice    []int
	StringSlice []string
	StructSlice []*HoldValMsg
	IntNeg      int
}

type HoldValMsg struct {
	Int         int
	Str         string
	Bool        bool
	ByteSlice   []byte
	BoolSlice   []bool
	IntSlice    []int
	StringSlice []string
}
