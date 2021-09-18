package analyzer

var regexpFunctions = map[string][]int{
	"Match":       {1},
	"MatchString": {1},
}

var regexpMethods = map[string][]int{
	// think on do we need expand at all?
	"Expand":                     {1},
	"ExpandString":               {1},
	"Match":                      {0},
	"MatchString":                {0},
	"Find":                       {0},
	"FindString":                 {0},
	"FindAll":                    {0},
	"FindAllString":              {0},
	"FindAllIndex":               {0},
	"FindAllStringIndex":         {0},
	"FindAllSubmatch":            {0},
	"FindAllStringSubmatch":      {0},
	"FindAllSubmatchIndex":       {0},
	"FindAllStringSubmatchIndex": {0},
	"FindIndex":                  {0},
	"FindStringIndex":            {0},
	"FindSubmatch":               {0},
	"FindStringSubmatch":         {0},
	"FindSubmatchIndex":          {0},
	"FindStringSubmatchIndex":    {0},
	"ReplaceAll":                 {0, 1},
	"ReplaceAllString":           {0, 1},
	"ReplaceAllFunc":             {0},
	"ReplaceAllStringFunc":       {0},
	"ReplaceAllLiteral":          {0, 1},
	"ReplaceAllLiteralString":    {0, 1},
}
