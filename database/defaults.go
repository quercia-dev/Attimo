package database

import (
	"strings"
)

const (
	// Datatype constants define behavior of the database in inserting and editing the data
	// by design choice, there is no default value for the behavior of the datatypes

	StringType string = "string"
	IntType    string = "int"
	FloatType  string = "float64"
	BoolType   string = "bool"
	TimeType   string = "time.Time"
	csvType    string = "[]string"

	LastCompletion   string = "last"
	NoCompletion     string = "no"
	UniqueCompletion string = "unique"
	SetCompletion    string = "in"
	DateCompletion   string = "date"
	FileCompletion   string = "file"

	NoSort         string = "no"
	FrequencySort  string = "frequency"
	LastSort       string = "last"
	alphabeticSort string = "alphabetic"

	nonemptyCheck string = "nonempty"
	RangeCheck    string = "range"
	SetCheck      string = "in"
	NoCheck       string = "no"
	URLCheck      string = "url"
	MailCheck     string = "mail"
	PhoneCheck    string = "phone"
	FileCheck     string = "file_exists"
	DateCheck     string = "date"
)

// ComposeArguments takes a list of strings and returns a string with the arguments formatted as a function call
// Example: ComposeArguments("func", "arg1", "arg2") -> "func(arg1,arg2)"
func ComposeArguments(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.WriteString(args[0])
	builder.WriteString("(")

	for i := 1; i < len(args); i++ {
		builder.WriteString(args[i])
		if i != len(args)-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString(")")
	return builder.String()
}

func getDefaultDatatypes() []Datatype {
	// Populate the database with default datatypes
	return []Datatype{
		// 1
		{Name: "Opened", VariableType: TimeType, CompletionValue: LastCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		// 2
		{Name: "Closed", VariableType: TimeType, CompletionValue: DateCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		// 3
		{Name: "Note", VariableType: StringType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: nonemptyCheck},
		// 4
		{Name: "Project", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		// 5
		{Name: "Person", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		// 6
		{Name: "Location", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		// 7
		{Name: "URL", VariableType: StringType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: URLCheck},
		// 8
		{Name: "Cost_EUR", VariableType: IntType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: nonemptyCheck},
		// 9
		{Name: "Deadline", VariableType: TimeType, CompletionValue: DateCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		// 10
		{Name: "Rating", VariableType: IntType, CompletionValue: SetCompletion + "(1,2,3,4,5)", CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		// 11
		{Name: "Email", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: NoSort, ValueCheck: MailCheck},
		// 12
		{Name: "Phone", VariableType: StringType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: PhoneCheck},
		// 13
		{Name: "File", VariableType: StringType, CompletionValue: FileCompletion, CompletionSort: NoSort, ValueCheck: FileCheck},
		// 14
		{Name: "Priority", VariableType: StringType, CompletionValue: SetCompletion + "(Low,Medium,High,Urgent)", CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		// 15
		{Name: "Status", VariableType: StringType, CompletionValue: SetCompletion + "(Not Started,In Progress,On Hold,Completed,Cancelled)", CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		// 16
		{Name: "Tags", VariableType: csvType, CompletionValue: UniqueCompletion, CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		// 18
		{Name: "Progress", VariableType: IntType, CompletionValue: SetCompletion + "(0,25,50,75,100)", CompletionSort: LastSort, ValueCheck: nonemptyCheck},
	}
}

func getDefaultCategories() []CategoryTemplate {
	return []CategoryTemplate{
		{Name: "General", ColumnsID: []int{1, 2, 3, 4, 6, 13}},
		{Name: "Contact", ColumnsID: []int{1, 2, 3, 11, 12, 13}},
		{Name: "Financial", ColumnsID: []int{1, 2, 3, 6, 8}},
	}
}
