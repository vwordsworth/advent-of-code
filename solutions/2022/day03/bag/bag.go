package bag

type Bag struct {
	AllItems   string
	FirstComp  string
	SecondComp string
}

type Group struct {
	First  Bag
	Second Bag
	Third  Bag
}

func NewBag(all string) Bag {
	first, second := splitItems(all)
	return Bag{AllItems: all, FirstComp: first, SecondComp: second}
}

func NewGroup(first, second, third Bag) Group {
	return Group{First: first, Second: second, Third: third}
}

func splitItems(bag string) (string, string) {
	size := len(bag)
	splitIdx := size / 2
	return bag[0:splitIdx], bag[splitIdx:size]
}

func (b Bag) GetSingleBagRepeatedRunePriority() int {
	var repeat rune
	first := map[rune]int{}
	for _, r := range b.FirstComp {
		first[r] = 1
	}

	for _, r := range b.SecondComp {
		if _, ok := first[r]; ok {
			repeat = r
			break
		}
	}
	return convertRuneValueToPriority(repeat)
}

func (g Group) GetGroupRepeatedRunePriority() int {
	var repeat rune
	first := map[rune]int{}
	for _, r := range g.First.AllItems {
		first[r] = 1
	}

	second := map[rune]int{}
	for _, r := range g.Second.AllItems {
		if _, ok := first[r]; ok {
			second[r] = 1
		}
	}

	for _, r := range g.Third.AllItems {
		if _, ok := second[r]; ok {
			repeat = r
			break
		}
	}
	return convertRuneValueToPriority(repeat)
}

func convertRuneValueToPriority(r rune) int {
	val := int(r)
	if isRuneUpperCase(r) {
		val -= 38 // Uppercase runes: 65-90, want 1-26 priority
	} else {
		val -= 96 // Lowercase runes: 97-148, want 27-52 priority
	}
	return val
}

func isRuneUpperCase(r rune) bool {
	return r < 'a' // All uppercase runes come before lowercase runes
}
