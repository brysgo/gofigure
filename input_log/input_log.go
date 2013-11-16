package input_log

type Input []string
type InputLog []Input

func (self InputLog) AddInput(input Input) InputLog {
	self = append(self, input)
	return self
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

func (self InputLog) SplitOnKey(targetKey string) (left InputLog, right InputLog) {
	for _, v := range self {
		remainder := make(Input, 1)
		containsTarget := false
		for _, k := range v {
			if k != targetKey {
				remainder = append(remainder, k)
			} else {
				containsTarget = true
			}
		}
		if containsTarget {
			right = append(right, remainder)
		} else {
			left = append(left, remainder)
		}
	}
	return
}

func (self Input) HasKey(a string) bool {
	for _, b := range self {
		if b == a {
			return true
		}
	}
	return false
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
