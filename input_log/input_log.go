package input_log

type InputLog [][]string
type Interface interface {
	AddInput([]string) Interface
	MaxEntropy(string) string
	SplitOnKey(string) (Interface, Interface)
	Any() bool
	At(int) []string
	Len() int
}

func (self InputLog) AddInput(input []string) Interface {
	self = append(self, input)
	return Interface(self)
}

func (self InputLog) MaxEntropy(targetKey string) string {
	existentialMap := self.existentialMap()
	entropy := make(map[string]int, len(existentialMap[0])-1)
	for _, input := range existentialMap {
		for key, value := range input {
			if key == targetKey {
				if value {
					entropy[targetKey] += 1
				}
				continue
			}
			if value == input[targetKey] {
				entropy[key] += 1
			} else {
				entropy[key] += 0
			}
		}
	}
	if entropy[targetKey]%len(self) == 0 {
		return targetKey
	}
	maxEntropy := 0.0
	resultKey := targetKey
	for key, value := range entropy {
		ent := float64(value - len(self)/2)
		if ent < 0 {
			ent *= -1
		}
		if ent >= maxEntropy {
			maxEntropy = ent
			resultKey = key
		}
	}
	return resultKey
}

func (self InputLog) SplitOnKey(targetKey string) (left Interface, right Interface) {
	left, right = make(InputLog, 0), make(InputLog, 0)
	for _, v := range self {
		remainder := make([]string, 0)
		containsTarget := false
		for _, k := range v {
			if k != targetKey {
				remainder = append(remainder, k)
			} else {
				containsTarget = true
			}
		}
		if containsTarget {
			right = right.AddInput(remainder)
		} else {
			left = left.AddInput(remainder)
		}
	}
	return
}

func (self InputLog) Any() bool {
	return len(self) > 0
}

func (self InputLog) At(index int) []string {
	return self[index]
}

func (self InputLog) Keys() []string {
	allKeys := make(map[string]int)
	for _, input := range self {
		for _, key := range input {
			if _, ok := allKeys[key]; !ok {
				allKeys[key] = len(allKeys)
			}
		}
	}
	result := make([]string, len(allKeys))
	for key, value := range allKeys {
		result[value] = key
	}
	return result
}

func (self InputLog) Len() int {
	return len(self)
}

func (self InputLog) existentialMap() (result []map[string]bool) {
	allKeys := make(map[string]bool)
	for index, input := range self {
		result = append(result, map[string]bool{})
		for _, key := range input {
			allKeys[key] = true
			result[index][key] = true
		}
	}
	for _, input := range result {
		for key, _ := range allKeys {
			if !input[key] {
				input[key] = false
			}
		}
	}
	return
}
