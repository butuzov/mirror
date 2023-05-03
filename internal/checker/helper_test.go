package checker

// func Test_isConverter(t *testing.T) {
// 	isConverterTestCases := []struct {
// 		name     string
// 		code     string
// 		expected bool
// 	}{
// 		{"byte", `[]byte("foobar")`, true},
// 		{"string", `string("foobar")`, true},
// 		{"string to byte to string", `string([]byte("foobar"))`, true},
// 		{"int conversion", `int(1.5)`, false},
// 		{"??", "foo(\"bar\")\ntype foo = string", true},
// 	}

// 	t.Parallel()
// 	for i, test := range isConverterTestCases {
// 		test := test
// 		t.Run(fmt.Sprintf("(%d)_%s", i, test.name), func(t *testing.T) {
// 			astree := parse(fmt.Sprintf("package foo\nvar _ = %s", test.code))
// 			result := isConverterCall(
// 				astree.Decls[0].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0].(*ast.CallExpr))

// 			assert.Equal(t, test.expected, result, "Expected (%t) vs Got (%t) Code (%s)",
// 				test.expected, result, test.code)
// 		})
// 	}
// }

// func Test_exprIsBytes(t *testing.T) {
// 	isConverterTestCases := []struct {
// 		name     string
// 		code     string
// 		expected bool
// 	}{
// 		// {"byte", `[]byte("foobar")`, true},
// 		{"byte alias", "[]bar(\"foobar\")\n type bar = byte", true},
// 		{"byte type", "[]baz(\"foobar\")\n type baz byte", true},
// 		{"byte type", "[]byte(\"foobar\")", true},
// 	}

// 	t.Parallel()
// 	for i, test := range isConverterTestCases {
// 		test := test
// 		t.Run(fmt.Sprintf("(%d)_%s", i, test.name), func(t *testing.T) {
// 			code := fmt.Sprintf("package foo\nvar _ = %s", test.code)
// 			typesInfo, astree := mustParse(code)
// 			el := astree.Decls[0].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0].(*ast.CallExpr)

// 			// el is call expression
// 			fmt.Printf(":%t\n", exprArrayType(typesInfo, el.Fun.(*ast.ArrayType)))
// 			fmt.Println(typesInfo.TypeOf(el).String())
// 			// _ = fset
// 			// spew.Dump(el.Fun.(*ast.ArrayType))
// 			// spew.Dump(el.Args[0])
// 			// result := exprIsBytes(el)

// 			// assert.Equal(t, test.expected, result, "Expected (%t) vs Got (%t) Code (%s)",
// 			// 	test.expected, result, test.code)
// 		})
// 	}
// }
