package utils

import "testing"

/*
func Test_IsPrime_9(t *testing.T) {
	//Arrange
	sut := 9
	expectedResult := false

	//Act
	actualResult := IsPrime(sut)

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v, but got %v\n", expectedResult, actualResult)
	}
}

func Test_IsPrime_7(t *testing.T) {
	//Arrange
	sut := 7
	expectedResult := true

	//Act
	actualResult := IsPrime(sut)

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v, but got %v\n", expectedResult, actualResult)
	}
}
*/

func Test_IsPrime(t *testing.T) {
	testData := []struct {
		sut            int
		expectedResult bool
		actualResult   bool
		name           string
	}{
		{sut: 7, expectedResult: true, name: "Test_IsPrime_7"},
		{sut: 8, expectedResult: false, name: "Test_IsPrime_8"},
		{sut: 9, expectedResult: true, name: "Test_IsPrime_9"},
		{sut: 11, expectedResult: true, name: "Test_IsPrime_11"},
		{sut: 12, expectedResult: false, name: "Test_IsPrime_12"},
	}
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			//Arrange
			sut := data.sut

			//Act
			data.actualResult = IsPrime(sut)

			//Assert
			if data.expectedResult != data.actualResult {
				t.Errorf("Expected %v, but got %v\n", data.expectedResult, data.actualResult)
			}
		})
	}
}
