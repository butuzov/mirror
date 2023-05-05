package mirror

import "github.com/butuzov/mirror/internal/checker"

func newHTTPTestChecker() *checker.Checker {
	c := checker.New("net/http/httptest")
	c.Methods["net/http/httptest.ResponseRecorder"] = HTTPTestMethods

	return c
}

var HTTPTestMethods = map[string]checker.Violation{
	"Write": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*httptest.ResponseRecorder).WriteString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Method: "WriteString",
		},
		Generate: &checker.Generate{
			PreCondition: `h := httptest.ResponseRecorder{}`,
			Pattern:      `Write($0)`,
			Returns:      2,
		},
	},
	"WriteString": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*httptest.ResponseRecorder).Write",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Method: "Write",
		},
		Generate: &checker.Generate{
			PreCondition: `h := httptest.ResponseRecorder{}`,
			Pattern:      `WriteString($0)`,
			Returns:      2,
		},
	},
}
