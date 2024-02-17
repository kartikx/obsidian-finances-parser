package parser

import (
	"strings"

	"github.com/kartikx/obsidian-finances-parser/pkg/models"
)

func GetCategoriesForExpense(expense *models.Expense) []models.ExpenseCategory {
	categorySet := map[models.ExpenseCategory]bool{}

	// Check Name to Category Mapping.
	for name, categories := range models.ExpenseNameToCategoryMap {
		if strings.Contains(expense.Name, name) {
			for _, category := range categories {
				categorySet[category] = true
			}
		}
	}

	// Larger expenses are expected to be Split.
	if (expense.Amount) > 500 {
		categorySet[models.SPLIT] = true
	}

	categories := make([]models.ExpenseCategory, 0, len(categorySet))
	for k := range categorySet {
		categories = append(categories, k)
	}

	return categories
}
