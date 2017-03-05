package go_sproto

type ValMsg struct {
	Int         int           `sproto:"integer,0,name=Int"`
	Str         string        `sproto:"string,1,name=Str"`
	Bool        bool          `sproto:"boolean,2,name=Bool"`
	Struct      *HoldValMsg   `sproto:"struct,3,name=Struct"`
	ByteSlice   []byte        `sproto:"string,4,name=ByteSlice"`
	BoolSlice   []bool        `sproto:"boolean,5,array,name=BoolSlice"`
	IntSlice    []int         `sproto:"integer,6,array,name=IntSlice"`
	StringSlice []string      `sproto:"string,7,array,name=StringSlice"`
	StructSlice []*HoldValMsg `sproto:"struct,8,array,name=StructSlice"`
	IntNeg      int           `sproto:"integer,9,name=IntNeg"`
}

type HoldValMsg struct {
	Int         int      `sproto:"integer,0,name=Int"`
	Str         string   `sproto:"string,1,name=Str"`
	Bool        bool     `sproto:"boolean,2,name=Bool"`
	ByteSlice   []byte   `sproto:"string,3,name=ByteSlice"`
	BoolSlice   []bool   `sproto:"boolean,4,array,name=BoolSlice"`
	IntSlice    []int    `sproto:"integer,5,array,name=IntSlice"`
	StringSlice []string `sproto:"string,6,array,name=StringSlice"`
}
