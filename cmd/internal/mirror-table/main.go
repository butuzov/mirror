// This generates MIRROR_FUNCS.md
package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/butuzov/mirror"
)

func main() {
	seen, keys := map[string]bool{}, map[string]string{}
	c := chaining(
		mirror.BufioMethods,
		mirror.BytesBufferMethods,
		mirror.BytesFunctions,
		mirror.HTTPTestMethods,
		mirror.MaphashMethods,
		mirror.OsFileMethods,
		mirror.RegexpFunctions,
		mirror.RegexpRegexpMethods,
		mirror.StringFunctions,
		mirror.StringsBuilderMethods,
		mirror.UTF8Functions,
	)

	// Create callers and sort them
	for v := range c {
		a, b := formCaller(v), formAltCaller(v)

		if !seen[a] {
			seen[a] = true
			seen[b] = true
			keys[a] = b
		}

		if !seen[b] {
			seen[b] = true
			seen[a] = true
			keys[b] = a
		}
	}

	sortKeys := []string{}
	for k := range keys {
		sortKeys = append(sortKeys, k)
	}

	sort.Slice(sortKeys, func(i, j int) bool {
		return strings.Compare(cleanSortKey(sortKeys[i]), cleanSortKey(sortKeys[j])) < 0
	})

	var bb bytes.Buffer

	for _, k := range sortKeys {
		fmt.Fprintf(&bb, "<tr>\n<td><code>%s</code></td>\n<td><code>%s</code></td>\n</tr>\n", k, keys[k])
	}

	fmt.Println(bb.String())
}
