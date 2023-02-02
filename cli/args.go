package cli

type Args interface {
	Get(n int) string
	First() string
	FetchArgs() []string
}

type args []string

func (a *args) Get(n int) string {
	if len(*a) > n {
		return (*a)[n]
	}
	return ""
}

func (a *args) First() string {
	return a.Get(0)
}

func (a *args) FetchArgs() []string {
	if len(*a) >= 2 {
		tail := []string((*a)[1:])
		return tail
	}
	return []string{}
}
