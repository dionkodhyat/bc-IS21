package rules

import (
	"bc-public-service-assessment/models"
	"testing"
)

func TestCalculateSupplement(t *testing.T) {
	tests := []struct {
		name                   string
		input                  models.InputData
		expectedBase           float64
		expectedChildrenAmount float64
		expectedSupplement     float64
	}{
		{
			name: "Single with 0 Child",
			input: models.InputData{
				ID:                         "test-id-1",
				NumberOfChildren:           0,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: true,
			},
			expectedBase:           60.0,
			expectedChildrenAmount: 0.0,
			expectedSupplement:     60.0,
		},
		{
			name: "Single with 1 Child",
			input: models.InputData{
				ID:                         "test-id-2",
				NumberOfChildren:           1,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: true,
			},
			expectedBase:           120.0,
			expectedChildrenAmount: 20.0,
			expectedSupplement:     140.0,
		},
		{
			name: "Single with 2 Children",
			input: models.InputData{
				ID:                         "test-id-3",
				NumberOfChildren:           2,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: true,
			},
			expectedBase:           120.0,
			expectedChildrenAmount: 40.0,
			expectedSupplement:     160,
		},
		{
			name: "Single with no children and did not pay in December",
			input: models.InputData{
				ID:                         "test-id-4",
				NumberOfChildren:           0,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: false,
			},
			expectedBase:           0,
			expectedChildrenAmount: 0,
			expectedSupplement:     0,
		}, {
			name: "Single with 2 children and did not pay in December",
			input: models.InputData{
				ID:                         "test-id-5",
				NumberOfChildren:           2,
				FamilyComposition:          "single",
				FamilyUnitInPayForDecember: false,
			},
			expectedBase:           0,
			expectedChildrenAmount: 0,
			expectedSupplement:     0,
		},
		{
			name: "Eligible couples with 0 children",
			input: models.InputData{
				ID:                         "test-id-6",
				NumberOfChildren:           0,
				FamilyComposition:          "couple",
				FamilyUnitInPayForDecember: true,
			},
			expectedBase:           120,
			expectedChildrenAmount: 0,
			expectedSupplement:     120,
		},
		{
			name: "Eligible couples with 1 children",
			input: models.InputData{
				ID:                         "test-id-7",
				NumberOfChildren:           1,
				FamilyComposition:          "couple",
				FamilyUnitInPayForDecember: true,
			},
			expectedBase:           120,
			expectedChildrenAmount: 20,
			expectedSupplement:     140,
		},
		{
			name: "Eligible couples with 2 children",
			input: models.InputData{
				ID:                         "test-id-7",
				NumberOfChildren:           2,
				FamilyComposition:          "couple",
				FamilyUnitInPayForDecember: true,
			},
			expectedBase:           120,
			expectedChildrenAmount: 40,
			expectedSupplement:     160,
		},
		{
			name: "Ineligible couples with no children",
			input: models.InputData{
				ID:                         "test-id-7",
				NumberOfChildren:           0,
				FamilyComposition:          "couple",
				FamilyUnitInPayForDecember: false,
			},
			expectedBase:           0,
			expectedChildrenAmount: 0,
			expectedSupplement:     0,
		},
		{
			name: "Ineligible couples with children",
			input: models.InputData{
				ID:                         "test-id-7",
				NumberOfChildren:           2,
				FamilyComposition:          "couple",
				FamilyUnitInPayForDecember: false,
			},
			expectedBase:           0,
			expectedChildrenAmount: 0,
			expectedSupplement:     0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := CalculateSupplement(test.input)
			if output.BaseAmount != test.expectedBase {
				t.Errorf("Expected BaseAmount %f, got %f", test.expectedBase, output.BaseAmount)
			}

			if output.ChildrenAmount != test.expectedChildrenAmount {
				t.Errorf("Expected ChildrenAmount %f, got %f", test.expectedChildrenAmount, output.ChildrenAmount)
			}

			if output.SupplementAmount != test.expectedSupplement {
				t.Errorf("Expected SupplementAmount %f, got %f", test.expectedSupplement, output.SupplementAmount)
			}
		})
	}
}
