package day5

type MapperRule struct {
	destinationStart int
	sourceStart      int
	length           int
}

type Mapper struct {
	name  string
	rules []MapperRule
}
