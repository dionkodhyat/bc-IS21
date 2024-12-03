package rules

import (
	"bc-public-service-assessment/models"
)

// CalculateSupplement processes InputData and returns OutputData
func CalculateSupplement(input models.InputData) models.OutputData {
	outputData := models.OutputData{
		ID:               input.ID,
		IsEligible:       input.FamilyUnitInPayForDecember,
		BaseAmount:       0.0,
		ChildrenAmount:   0.0,
		SupplementAmount: 0.0,
	}

	if !input.FamilyUnitInPayForDecember {
		return outputData
	}

	// Calculate base amount
	if input.NumberOfChildren > 0 {
		outputData.BaseAmount = 120.0
	} else {
		switch input.FamilyComposition {
		case "single":
			outputData.BaseAmount = 60.0
		case "couple":
			outputData.BaseAmount = 120.0
		}
	}

	outputData.ChildrenAmount = float64(input.NumberOfChildren) * 20.0
	outputData.SupplementAmount = outputData.BaseAmount + outputData.ChildrenAmount

	return outputData
}
