package goset

type Set map[string]bool

func NewSet(elements []string) (s Set) {
	s = make(Set)
	for _, str := range elements {
		s[str] = true
	}
	return s
}

func (s Set) Contains(str string) (contains bool) {
	return s[str]
}

func (s Set) Add(str string) {
	s[str] = true
}

func (s Set) Remove(str string) {
	delete(s, str)
}

func (s Set) Union(s2 Set) (set Set) {
	set = make(Set)
	for str := range s {
		set.Add(str)
	}
	for str := range s2 {
		set.Add(str)
	}
	return set
}

func (s Set) Intersection(s2 Set) (set Set) {
	set = make(Set)
	var smallSet, bigSet Set
	if len(s) > len(s2) {
		smallSet = s2
		bigSet = s
	} else {
		smallSet = s
		bigSet = s2
	}

	for str := range smallSet {
		if bigSet.Contains(str) {
			set.Add(str)
		}
	}
	return set
}

func (s Set) Difference(s2 Set) (set Set) {
	set = make(Set)
	for str := range s {
		if !s2.Contains(str) {
			set.Add(str)
		}
	}
	return set
}

func (s Set) Copy() (set Set) {
	set = make(Set)
	for str := range s {
		set.Add(str)
	}
	return set
}

func (s Set) String() (str string) {
	str = "Set{"
	first := true
	for st := range s {
		if first {
			str += `"` + st + `"`
			first = false
		} else {
			str += "," + `"` + st + `"`
		}

	}
	return str + "}"
}
