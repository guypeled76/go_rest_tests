package api

import "testing"

func GetSomethingTest(t *testing.T) {

	_, err := GetSomething("", "ddd")
	if err == nil {
		t.Error("Should return an error if project is not valid")
	}

}
