package verifier

type Args struct {
	Token string
	Scope string
}

type Reply struct {
	Valid bool
}
