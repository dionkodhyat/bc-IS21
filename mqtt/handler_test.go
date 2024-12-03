package mqtt

import (
	"bc-public-service-assessment/models"
	"testing"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name        string
		input       models.InputData
		expectError bool
	}{
		{
			name: "Valid input",
			input: models.InputData{
				ID:                         "test-id-1",
				NumberOfChildren:           1,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: true,
			},
			expectError: false,
		}, {
			name: "Invalid ID",
			input: models.InputData{
				ID:                         "",
				NumberOfChildren:           1,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: true,
			},
			expectError: true,
		}, {
			name: "Negative number of children",
			input: models.InputData{
				ID:                         "test-id-2",
				NumberOfChildren:           -1,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: true,
			},
			expectError: true,
		}, {
			name: "Invalid family composition",
			input: models.InputData{
				ID:                         "test-id-2",
				NumberOfChildren:           -1,
				FamilyComposition:          "",
				FamilyUnitInPayForDecember: true,
			},
			expectError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateInput(test.input)
			if (err != nil) != test.expectError {
				t.Errorf("validateInput() expected error: %v, got: %v", test.expectError, err)
			}
		})
	}
}
