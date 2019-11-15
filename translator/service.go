package translator

type Translator interface {
	Translate(word string) (string, error)
}

type translator struct {
	dic Dictionary
}

func NewTranslator(dic Dictionary) Translator {
	return &translator{
		dic: dic,
	}
}

// Translate translates word into targeted language (English or Chinese auto detected)
func (s *translator) Translate(word string) (string, error) {
	result, err := s.dic.Translate(word)
	if err != nil {
		return "", err
	}

	return result, nil
}
