package mock

type Retriever struct{}

func (Retriever) Get(url string) string {
	return "fake response"
}
