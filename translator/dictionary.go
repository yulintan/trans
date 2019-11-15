package translator

type Dictionary interface {
	Translate(s string) (string, error)
}
