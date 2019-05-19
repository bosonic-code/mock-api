package mocker

type Request struct {
	Method  string
	Path    string
	Query   map[string]string
	Headers map[string]string
	Body    string
}

type Response struct {
	Status  int32
	Body    string
	Headers map[string]string
}
