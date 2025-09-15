package restkit

type Query map[string]string

func (qp Query) String() string {
	if len(qp) == 0 {
		return ""
	}
	result := "?"
	first := true
	for k, v := range qp {
		if !first {
			result += "&"
		}
		result += k + "=" + v
		first = false
	}
	return result
}
