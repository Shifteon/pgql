package analyzer

import "fmt"

type result struct {
	err error
}

func ok() result {
	return result{}
}

func fail(msg string, lineNumber int, charPosition int) result {
	return result{err: fmt.Errorf("line %v : char %v - error: %v", lineNumber, charPosition, msg)}
}
