package sensitive

import (
	"github.com/gogf/gf/v2/encoding/ghash"
)

// 字符串数据结构
type StringPointer struct {
	Value  []rune
	Offset int
	Length int
	Hash   int
}

// 字符串构造，将字符串拆解成字符
func (this *StringPointer) StringPointer(str string) *StringPointer {
	b := make([]rune, len(str))
	for i, s := range str {
		b[i] = s
	}
	obj := &StringPointer{
		Value:  b,
		Offset: 0,
		Hash:   0,
		Length: len(b),
	}
	return obj
}

// 字符串构造，将字符串拆解成字符
func (this *StringPointer) StringPointerByThree(str string, offset int, length int) *StringPointer {
	b := make([]rune, len(str))
	for i, s := range str {
		b[i] = s
	}
	obj := &StringPointer{
		Value:  b,
		Offset: offset,
		Hash:   0,
		Length: len(b),
	}
	return obj
}

// 计算该位置后（包含）2个字符的hash值
func (this *StringPointer) nextTwoCharHash(index int) uint64 {
	first := string(this.Value[this.Offset+index])
	secont := string(this.Value[this.Offset+index+1])
	return 31*ghash.ELF64([]byte(first)) + ghash.ELF64([]byte(secont))
}

// 计算该位置后（包含）2个字符和为1个int型的值
func (this *StringPointer) nextTwoCharMix(index int) int32 {
	first := this.Value[this.Offset+index] << 16
	second := this.Value[this.Offset+index+1]
	return first | second
}
