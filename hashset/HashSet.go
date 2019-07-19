package hashset

type HashsetMap map[string]struct{}

type Hashset struct {
	set HashsetMap
}

func (h *Hashset) Size() int {
	return len(h.set)
}

func (h *Hashset) Add(value string) {
	h.set[value] = struct{}{}
}

func (h *Hashset) ToArray() []string {
	var result []string
	for key := range h.set {
		result = append(result, key)
	}
	return result
}

func New(values ...string) *Hashset {
	hashset := Hashset{make(HashsetMap)}
	for _, value := range values {
		hashset.set[value] = struct{}{}
	}
	return &hashset
}
