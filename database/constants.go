// database/constants.go
package database

import (
	"strings"
)

// Date format
const (
	dateFormat         = "02-01-2006"
	dbSetupErrorString = "Failed to set up database: %v"
)

// Data types
const (
	StringType = "string"
	IntType    = "int"
	FloatType  = "float64"
	BoolType   = "bool"
	TimeType   = "time.Time"
	csvType    = "[]string"
)

// Completion types
const (
	LastCompletion   = "last"
	NoCompletion     = "no"
	UniqueCompletion = "unique"
	SetCompletion    = "in"
	DateCompletion   = "date"
	FileCompletion   = "file"
)

// Sort types
const (
	NoSort         = "no"
	FrequencySort  = "frequency"
	LastSort       = "last"
	alphabeticSort = "alphabetic"
)

// Validation types
const (
	nonemptyCheck = "nonempty"
	RangeCheck    = "range"
	SetCheck      = "in"
	NoCheck       = "no"
	URLCheck      = "url"
	MailCheck     = "mail"
	PhoneCheck    = "phone"
	FileCheck     = "file_exists"
	DateCheck     = "date"
)

// Current version
const currentVersion = "1.0.0"

// ComposeArguments takes a list of strings and returns a string with the arguments formatted as a function call
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

// getDefaultDatatypes returns the default datatype configurations
func getDefaultDatatypes() []Datatype {
	return []Datatype{
		{Name: "Opened", VariableType: TimeType, CompletionValue: LastCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		{Name: "Closed", VariableType: TimeType, CompletionValue: DateCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		{Name: "Note", VariableType: StringType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: nonemptyCheck},
		{Name: "Project", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		{Name: "Person", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		{Name: "Location", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		{Name: "URL", VariableType: StringType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: URLCheck},
		{Name: "Cost_EUR", VariableType: IntType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: nonemptyCheck},
		{Name: "Deadline", VariableType: TimeType, CompletionValue: DateCompletion, CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		{Name: "Rating", VariableType: IntType, CompletionValue: SetCompletion + "(1,2,3,4,5)", CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		{Name: "Email", VariableType: StringType, CompletionValue: UniqueCompletion, CompletionSort: NoSort, ValueCheck: MailCheck},
		{Name: "Phone", VariableType: StringType, CompletionValue: NoCompletion, CompletionSort: NoSort, ValueCheck: PhoneCheck},
		{Name: "File", VariableType: StringType, CompletionValue: FileCompletion, CompletionSort: NoSort, ValueCheck: FileCheck},
		{Name: "Priority", VariableType: StringType, CompletionValue: SetCompletion + "(Low,Medium,High,Urgent)", CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		{Name: "Status", VariableType: StringType, CompletionValue: SetCompletion + "(Not Started,In Progress,On Hold,Completed,Cancelled)", CompletionSort: LastSort, ValueCheck: nonemptyCheck},
		{Name: "Tags", VariableType: csvType, CompletionValue: UniqueCompletion, CompletionSort: FrequencySort, ValueCheck: nonemptyCheck},
		{Name: "Progress", VariableType: IntType, CompletionValue: SetCompletion + "(0,25,50,75,100)", CompletionSort: LastSort, ValueCheck: nonemptyCheck},
	}
}

// getDefaultCategories returns the default category configurations
func getDefaultCategories() []CategoryTemplate {
	return []CategoryTemplate{
		{Name: "General", ColumnsID: []int{1, 2, 3, 4, 6, 13}},   // Opened, Closed, Note, Project, Location, File
		{Name: "Contact", ColumnsID: []int{1, 2, 3, 11, 12, 13}}, // Opened, Closed, Note, Email, Phone, File
		{Name: "Financial", ColumnsID: []int{1, 2, 3, 6, 8}},     // Opened, Closed, Note, Location, Cost_EUR
	}
}
