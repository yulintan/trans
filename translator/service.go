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
func (t *translator) Translate(word string) (string, error) {
	raw, err := t.dic.Translate(word)
	if err != nil {
		return "", err
	}

	result, err := t.dic.PrettyPrint(raw)
	if err != nil {
		return "", err
	}

	return result, nil
}
