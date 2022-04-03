package gsr7

import "fmt"

type validator func() error

func validate(validators ...validator) error {
	for _, v := range validators {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}

func validateCookieName(name string) validator {
	return func() error {
		separators := map[int32]struct{}{
			'(':  {},
			')':  {},
			'<':  {},
			'>':  {},
			'@':  {},
			',':  {},
			';':  {},
			':':  {},
			'\\': {},
			'"':  {},
			'/':  {},
			'[':  {},
			']':  {},
			'?':  {},
			'=':  {},
			'{':  {},
			'}':  {},
			' ':  {},
			'\t': {},
		}
		for i, letter := range name {
			if letter < 32 || letter == 127 {
				return fmt.Errorf("invalid character in cookie name position %d (%d)", i, letter)
			}
			if _, ok := separators[letter]; ok {
				return fmt.Errorf("invalid character in cookie name position %d (%d)", i, letter)
			}
		}
		return nil
	}
}

func validateCookieDomain(domain string) validator {
	return func() error {
		for i, letter := range domain {
			if letter < 32 || letter == 127 || letter == ';' {
				return fmt.Errorf("invalid character in domain name position %d (%d)", i, letter)
			}
		}
		return nil
	}
}

func validateCookiePath(path string) validator {
	return func() error {
		for i, letter := range path {
			if letter < 32 || letter == 127 || letter == ';' {
				return fmt.Errorf("invalid character in path position %d (%d)", i, letter)
			}
		}
		return nil
	}
}

func validateExtensions(extensions []string) validator {
	return func() error {
		for i, extension := range extensions {
			for j, letter := range extension {
				if letter < 32 || letter == 127 || letter == ';' {
					return fmt.Errorf(
						"invalid extension %d in cookie, character %d is invalid (%d)",
						i,
						j,
						letter,
					)
				}
			}
		}
		return nil
	}
}
