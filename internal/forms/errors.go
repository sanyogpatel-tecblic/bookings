package forms

type errors map[string][]string

func (e errors) Add(fields, message string) {
	e[fields] = append(e[fields], message)
}

func (e errors) Get(fields string) string {
	es := e[fields]

	if len(es) == 0 {
		return ""
	}
	return es[0]
}
