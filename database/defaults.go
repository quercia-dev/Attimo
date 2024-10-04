package database

func GetDefaultDatatypes() []Datatype {
	timeType := "time.Time"
	// Populate the database with default datatypes
	return []Datatype{
		// 1
		{Name: "Opened", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: "default"},
		// 2
		{Name: "Closed", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: "default"},
		// 3
		{Name: "Note", VariableType: "string", CompletionValue: "no", CompletionSort: "default", ValueCheck: "default"},
		// 4
		{Name: "Project", VariableType: "string", CompletionValue: "unique", CompletionSort: "last", ValueCheck: "default"},
		// 5
		{Name: "Person", VariableType: "string", CompletionValue: "unique", CompletionSort: "frequency", ValueCheck: "default"},
		// 6
		{Name: "Location", VariableType: "string", CompletionValue: "unique", CompletionSort: "last", ValueCheck: "default"},
		// 7
		{Name: "URL", VariableType: "string", CompletionValue: "no", CompletionSort: "default", ValueCheck: "URL"},
		// 8
		{Name: "Cost (EUR)", VariableType: "integer", CompletionValue: "no", CompletionSort: "default", ValueCheck: "default"},
		// 9
		{Name: "Deadline", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: "default"},
		// 10
		{Name: "Rating", VariableType: "integer", CompletionValue: "{1,2,3,4,5}", CompletionSort: "frequency", ValueCheck: "default"},
		// 11
		{Name: "Email", VariableType: "string", CompletionValue: "unique", CompletionSort: "default", ValueCheck: "mail_ping"},
		// 12
		{Name: "Phone", VariableType: "string", CompletionValue: "no", CompletionSort: "default", ValueCheck: "phone"},
		// 13
		{Name: "File", VariableType: "string", CompletionValue: "file", CompletionSort: "default", ValueCheck: "file_exists"},
	}
}

func GetDefaultCategories() []CategoryTemplate {
	return []CategoryTemplate{
		{Name: "General", ColumnsID: []int{1, 2, 3, 4, 6, 13}},
		{Name: "Contact", ColumnsID: []int{1, 2, 3, 11, 12, 13}},
		{Name: "Financial", ColumnsID: []int{1, 2, 3, 6, 8}},
	}
}
