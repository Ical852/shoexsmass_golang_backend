package category

type CategoryFormatter struct {
	ID 		int		`json:"id"`
	Name 	string	`json:"name"`
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:   category.ID,
		Name: category.Name,
	}

	return formatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	var categoriesFormated []CategoryFormatter
	for _, category := range categories {
		categoriesFormated = append(categoriesFormated, FormatCategory(category))
	}

	return categoriesFormated
}