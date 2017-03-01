package sproto

type PtrMSG struct {
	Int         *int          `sproto:"integer,0,name=Int"`
	String      *string       `sproto:"string,1,name=String"`
	Bool        *bool         `sproto:"boolean,2,name=Bool"`
	Struct      *HoldPtrMSG   `sproto:"struct,3,name=Struct"`
	ByteSlice   []byte        `sproto:"string,4,name=ByteSlice"`
	BoolSlice   []bool        `sproto:"boolean,5,array,name=BoolSlice"`
	IntSlice    []int         `sproto:"integer,6,array,name=IntSlice"`
	StringSlice []string      `sproto:"string,7,array,name=StringSlice"`
	StructSlice []*HoldPtrMSG `sproto:"struct,8,array,name=StructSlice"`
	IntNeg      *int          `sproto:"integer,9,name=IntNeg"`
}

type HoldPtrMSG struct {
	Int         *int     `sproto:"integer,0,name=Int"`
	String      *string  `sproto:"string,1,name=String"`
	Bool        *bool    `sproto:"boolean,2,name=Bool"`
	ByteSlice   []byte   `sproto:"string,3,name=ByteSlice"`
	BoolSlice   []bool   `sproto:"boolean,4,array,name=BoolSlice"`
	IntSlice    []int    `sproto:"integer,5,array,name=IntSlice"`
	StringSlice []string `sproto:"string,6,array,name=StringSlice"`
}

type ValMSG struct {
	Int         int           `sproto:"integer,0,name=Int"`
	Str         string        `sproto:"string,1,name=Str"`
	Bool        bool          `sproto:"boolean,2,name=Bool"`
	Struct      *HoldValMSG   `sproto:"struct,3,name=Struct"`
	ByteSlice   []byte        `sproto:"string,4,name=ByteSlice"`
	BoolSlice   []bool        `sproto:"boolean,5,array,name=BoolSlice"`
	IntSlice    []int         `sproto:"integer,6,array,name=IntSlice"`
	StringSlice []string      `sproto:"string,7,array,name=StringSlice"`
	StructSlice []*HoldValMSG `sproto:"struct,8,array,name=StructSlice"`
	IntNeg      int           `sproto:"integer,9,name=IntNeg"`
}

type HoldValMSG struct {
	Int         int      `sproto:"integer,0,name=Int"`
	Str         string   `sproto:"string,1,name=Str"`
	Bool        bool     `sproto:"boolean,2,name=Bool"`
	ByteSlice   []byte   `sproto:"string,3,name=ByteSlice"`
	BoolSlice   []bool   `sproto:"boolean,4,array,name=BoolSlice"`
	IntSlice    []int    `sproto:"integer,5,array,name=IntSlice"`
	StringSlice []string `sproto:"string,6,array,name=StringSlice"`
}
