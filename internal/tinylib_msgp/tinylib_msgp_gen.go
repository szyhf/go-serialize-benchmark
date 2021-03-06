package msgp

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

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
	for za0001 := range z.BoolSlice {
		o = msgp.AppendBool(o, z.BoolSlice[za0001])
	}
	// string "IntSlice"
	o = append(o, 0xa8, 0x49, 0x6e, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.IntSlice)))
	for za0002 := range z.IntSlice {
		o = msgp.AppendInt(o, z.IntSlice[za0002])
	}
	// string "StringSlice"
	o = append(o, 0xab, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.StringSlice)))
	for za0003 := range z.StringSlice {
		o = msgp.AppendString(o, z.StringSlice[za0003])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HoldValMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Int":
			z.Int, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Int")
				return
			}
		case "Str":
			z.Str, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Str")
				return
			}
		case "Bool":
			z.Bool, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bool")
				return
			}
		case "ByteSlice":
			z.ByteSlice, bts, err = msgp.ReadBytesBytes(bts, z.ByteSlice)
			if err != nil {
				err = msgp.WrapError(err, "ByteSlice")
				return
			}
		case "BoolSlice":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BoolSlice")
				return
			}
			if cap(z.BoolSlice) >= int(zb0002) {
				z.BoolSlice = (z.BoolSlice)[:zb0002]
			} else {
				z.BoolSlice = make([]bool, zb0002)
			}
			for za0001 := range z.BoolSlice {
				z.BoolSlice[za0001], bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "BoolSlice", za0001)
					return
				}
			}
		case "IntSlice":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IntSlice")
				return
			}
			if cap(z.IntSlice) >= int(zb0003) {
				z.IntSlice = (z.IntSlice)[:zb0003]
			} else {
				z.IntSlice = make([]int, zb0003)
			}
			for za0002 := range z.IntSlice {
				z.IntSlice[za0002], bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "IntSlice", za0002)
					return
				}
			}
		case "StringSlice":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StringSlice")
				return
			}
			if cap(z.StringSlice) >= int(zb0004) {
				z.StringSlice = (z.StringSlice)[:zb0004]
			} else {
				z.StringSlice = make([]string, zb0004)
			}
			for za0003 := range z.StringSlice {
				z.StringSlice[za0003], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "StringSlice", za0003)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
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
	for za0003 := range z.StringSlice {
		s += msgp.StringPrefixSize + len(z.StringSlice[za0003])
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
			err = msgp.WrapError(err, "Struct")
			return
		}
	}
	// string "ByteSlice"
	o = append(o, 0xa9, 0x42, 0x79, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendBytes(o, z.ByteSlice)
	// string "BoolSlice"
	o = append(o, 0xa9, 0x42, 0x6f, 0x6f, 0x6c, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.BoolSlice)))
	for za0001 := range z.BoolSlice {
		o = msgp.AppendBool(o, z.BoolSlice[za0001])
	}
	// string "IntSlice"
	o = append(o, 0xa8, 0x49, 0x6e, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.IntSlice)))
	for za0002 := range z.IntSlice {
		o = msgp.AppendInt(o, z.IntSlice[za0002])
	}
	// string "StringSlice"
	o = append(o, 0xab, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.StringSlice)))
	for za0003 := range z.StringSlice {
		o = msgp.AppendString(o, z.StringSlice[za0003])
	}
	// string "StructSlice"
	o = append(o, 0xab, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.StructSlice)))
	for za0004 := range z.StructSlice {
		if z.StructSlice[za0004] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.StructSlice[za0004].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "StructSlice", za0004)
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
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Int":
			z.Int, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Int")
				return
			}
		case "Str":
			z.Str, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Str")
				return
			}
		case "Bool":
			z.Bool, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bool")
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
					err = msgp.WrapError(err, "Struct")
					return
				}
			}
		case "ByteSlice":
			z.ByteSlice, bts, err = msgp.ReadBytesBytes(bts, z.ByteSlice)
			if err != nil {
				err = msgp.WrapError(err, "ByteSlice")
				return
			}
		case "BoolSlice":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BoolSlice")
				return
			}
			if cap(z.BoolSlice) >= int(zb0002) {
				z.BoolSlice = (z.BoolSlice)[:zb0002]
			} else {
				z.BoolSlice = make([]bool, zb0002)
			}
			for za0001 := range z.BoolSlice {
				z.BoolSlice[za0001], bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "BoolSlice", za0001)
					return
				}
			}
		case "IntSlice":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IntSlice")
				return
			}
			if cap(z.IntSlice) >= int(zb0003) {
				z.IntSlice = (z.IntSlice)[:zb0003]
			} else {
				z.IntSlice = make([]int, zb0003)
			}
			for za0002 := range z.IntSlice {
				z.IntSlice[za0002], bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "IntSlice", za0002)
					return
				}
			}
		case "StringSlice":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StringSlice")
				return
			}
			if cap(z.StringSlice) >= int(zb0004) {
				z.StringSlice = (z.StringSlice)[:zb0004]
			} else {
				z.StringSlice = make([]string, zb0004)
			}
			for za0003 := range z.StringSlice {
				z.StringSlice[za0003], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "StringSlice", za0003)
					return
				}
			}
		case "StructSlice":
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StructSlice")
				return
			}
			if cap(z.StructSlice) >= int(zb0005) {
				z.StructSlice = (z.StructSlice)[:zb0005]
			} else {
				z.StructSlice = make([]*HoldValMsg, zb0005)
			}
			for za0004 := range z.StructSlice {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.StructSlice[za0004] = nil
				} else {
					if z.StructSlice[za0004] == nil {
						z.StructSlice[za0004] = new(HoldValMsg)
					}
					bts, err = z.StructSlice[za0004].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "StructSlice", za0004)
						return
					}
				}
			}
		case "IntNeg":
			z.IntNeg, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IntNeg")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
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
	for za0003 := range z.StringSlice {
		s += msgp.StringPrefixSize + len(z.StringSlice[za0003])
	}
	s += 12 + msgp.ArrayHeaderSize
	for za0004 := range z.StructSlice {
		if z.StructSlice[za0004] == nil {
			s += msgp.NilSize
		} else {
			s += z.StructSlice[za0004].Msgsize()
		}
	}
	s += 7 + msgp.IntSize
	return
}
