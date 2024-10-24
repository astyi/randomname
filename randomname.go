package randomname

import (
	"math/rand"
	"unicode/utf8"
)

// GenerateName 生成随机昵称
func GenerateName() string {
	var name string
	selectedType := RandomType(rand.Intn(3))
	switch selectedType {
	case AdjectiveAndPerson:
		name = AdjectiveSlice[rand.Intn(AdjectiveSliceCount)] + PersonSlice[rand.Intn(PersonSliceCount)]
	case PersonActSomething:
		name = PersonSlice[rand.Intn(PersonSliceCount)] + ActSomethingSlice[rand.Intn(ActSomethingSliceCount)]
	case AdjectiveNouns:
		name = AdjectiveSlice[rand.Intn(AdjectiveSliceCount)] + NounsSlice[rand.Intn(NounsSliceCount)]
	}
	return name
}

// GenerateNameNotExceedLen 生成不超过指定长度的随机昵称
func GenerateNameNotExceedLen(len int) string {
	var name string
	selectedType := RandomType(rand.Intn(3))
	switch selectedType {
	case AdjectiveAndPerson:
		name = AdjectiveSlice[rand.Intn(AdjectiveSliceCount)] + PersonSlice[rand.Intn(PersonSliceCount)]
	case PersonActSomething:
		name = PersonSlice[rand.Intn(PersonSliceCount)] + ActSomethingSlice[rand.Intn(ActSomethingSliceCount)]
	case AdjectiveNouns:
		name = AdjectiveSlice[rand.Intn(AdjectiveSliceCount)] + NounsSlice[rand.Intn(NounsSliceCount)]
	}
	if utf8.RuneCountInString(name) > len {
		s := []rune(name)
		name = string(s[:len])
	}
	return name
}
