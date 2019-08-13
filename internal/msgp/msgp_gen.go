package msgp

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *HoldValMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Int"
	o = append(o, 0x87, 0xa3, 0x49, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.Int)
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	// string "Bool"
	o = append(o, 0xa4, 0x42, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendBool(o, z.Bool)
	// string "ByteSlice"
	o = append(o, 0xa9, 0x42, 0x79, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendBytes(o, z.ByteSlice)
	// string "BoolSlice"
	o = append(o, 0xa9, 0x42, 0x6f, 0x6f, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.BoolSlice)))
	for zxvk := range z.BoolSlice {
		o = msgp.AppendBool(o, z.BoolSlice[zxvk])
	}
	// string "IntSlice"
	o = append(o, 0xa8, 0x49, 0x6e, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.IntSlice)))
	for zbzg := range z.IntSlice {
		o = msgp.AppendInt(o, z.IntSlice[zbzg])
	}
	// string "StringSlice"
	o = append(o, 0xab, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.StringSlice)))
	for zbai := range z.StringSlice {
		o = msgp.AppendString(o, z.StringSlice[zbai])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HoldValMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Int":
			z.Int, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Str":
			z.Str, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Bool":
			z.Bool, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "ByteSlice":
			z.ByteSlice, bts, err = msgp.ReadBytesBytes(bts, z.ByteSlice)
			if err != nil {
				return
			}
		case "BoolSlice":
			var zajw uint32
			zajw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.BoolSlice) >= int(zajw) {
				z.BoolSlice = (z.BoolSlice)[:zajw]
			} else {
				z.BoolSlice = make([]bool, zajw)
			}
			for zxvk := range z.BoolSlice {
				z.BoolSlice[zxvk], bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					return
				}
			}
		case "IntSlice":
			var zwht uint32
			zwht, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.IntSlice) >= int(zwht) {
				z.IntSlice = (z.IntSlice)[:zwht]
			} else {
				z.IntSlice = make([]int, zwht)
			}
			for zbzg := range z.IntSlice {
				z.IntSlice[zbzg], bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					return
				}
			}
		case "StringSlice":
			var zhct uint32
			zhct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.StringSlice) >= int(zhct) {
				z.StringSlice = (z.StringSlice)[:zhct]
			} else {
				z.StringSlice = make([]string, zhct)
			}
			for zbai := range z.StringSlice {
				z.StringSlice[zbai], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HoldValMsg) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 4 + msgp.StringPrefixSize + len(z.Str) + 5 + msgp.BoolSize + 10 + msgp.BytesPrefixSize + len(z.ByteSlice) + 10 + msgp.ArrayHeaderSize + (len(z.BoolSlice) * (msgp.BoolSize)) + 9 + msgp.ArrayHeaderSize + (len(z.IntSlice) * (msgp.IntSize)) + 12 + msgp.ArrayHeaderSize
	for zbai := range z.StringSlice {
		s += msgp.StringPrefixSize + len(z.StringSlice[zbai])
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ValMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "Int"
	o = append(o, 0x8a, 0xa3, 0x49, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.Int)
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	// string "Bool"
	o = append(o, 0xa4, 0x42, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendBool(o, z.Bool)
	// string "Struct"
	o = append(o, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74)
	if z.Struct == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Struct.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "ByteSlice"
	o = append(o, 0xa9, 0x42, 0x79, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendBytes(o, z.ByteSlice)
	// string "BoolSlice"
	o = append(o, 0xa9, 0x42, 0x6f, 0x6f, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.BoolSlice)))
	for zcua := range z.BoolSlice {
		o = msgp.AppendBool(o, z.BoolSlice[zcua])
	}
	// string "IntSlice"
	o = append(o, 0xa8, 0x49, 0x6e, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.IntSlice)))
	for zxhx := range z.IntSlice {
		o = msgp.AppendInt(o, z.IntSlice[zxhx])
	}
	// string "StringSlice"
	o = append(o, 0xab, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.StringSlice)))
	for zlqf := range z.StringSlice {
		o = msgp.AppendString(o, z.StringSlice[zlqf])
	}
	// string "StructSlice"
	o = append(o, 0xab, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.StructSlice)))
	for zdaf := range z.StructSlice {
		if z.StructSlice[zdaf] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.StructSlice[zdaf].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "IntNeg"
	o = append(o, 0xa6, 0x49, 0x6e, 0x74, 0x4e, 0x65, 0x67)
	o = msgp.AppendInt(o, z.IntNeg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ValMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Int":
			z.Int, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Str":
			z.Str, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Bool":
			z.Bool, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Struct":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Struct = nil
			} else {
				if z.Struct == nil {
					z.Struct = new(HoldValMsg)
				}
				bts, err = z.Struct.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "ByteSlice":
			z.ByteSlice, bts, err = msgp.ReadBytesBytes(bts, z.ByteSlice)
			if err != nil {
				return
			}
		case "BoolSlice":
			var zjfb uint32
			zjfb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.BoolSlice) >= int(zjfb) {
				z.BoolSlice = (z.BoolSlice)[:zjfb]
			} else {
				z.BoolSlice = make([]bool, zjfb)
			}
			for zcua := range z.BoolSlice {
				z.BoolSlice[zcua], bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					return
				}
			}
		case "IntSlice":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.IntSlice) >= int(zcxo) {
				z.IntSlice = (z.IntSlice)[:zcxo]
			} else {
				z.IntSlice = make([]int, zcxo)
			}
			for zxhx := range z.IntSlice {
				z.IntSlice[zxhx], bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					return
				}
			}
		case "StringSlice":
			var zeff uint32
			zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.StringSlice) >= int(zeff) {
				z.StringSlice = (z.StringSlice)[:zeff]
			} else {
				z.StringSlice = make([]string, zeff)
			}
			for zlqf := range z.StringSlice {
				z.StringSlice[zlqf], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "StructSlice":
			var zrsw uint32
			zrsw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.StructSlice) >= int(zrsw) {
				z.StructSlice = (z.StructSlice)[:zrsw]
			} else {
				z.StructSlice = make([]*HoldValMsg, zrsw)
			}
			for zdaf := range z.StructSlice {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.StructSlice[zdaf] = nil
				} else {
					if z.StructSlice[zdaf] == nil {
						z.StructSlice[zdaf] = new(HoldValMsg)
					}
					bts, err = z.StructSlice[zdaf].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "IntNeg":
			z.IntNeg, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ValMsg) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 4 + msgp.StringPrefixSize + len(z.Str) + 5 + msgp.BoolSize + 7
	if z.Struct == nil {
		s += msgp.NilSize
	} else {
		s += z.Struct.Msgsize()
	}
	s += 10 + msgp.BytesPrefixSize + len(z.ByteSlice) + 10 + msgp.ArrayHeaderSize + (len(z.BoolSlice) * (msgp.BoolSize)) + 9 + msgp.ArrayHeaderSize + (len(z.IntSlice) * (msgp.IntSize)) + 12 + msgp.ArrayHeaderSize
	for zlqf := range z.StringSlice {
		s += msgp.StringPrefixSize + len(z.StringSlice[zlqf])
	}
	s += 12 + msgp.ArrayHeaderSize
	for zdaf := range z.StructSlice {
		if z.StructSlice[zdaf] == nil {
			s += msgp.NilSize
		} else {
			s += z.StructSlice[zdaf].Msgsize()
		}
	}
	s += 7 + msgp.IntSize
	return
}
