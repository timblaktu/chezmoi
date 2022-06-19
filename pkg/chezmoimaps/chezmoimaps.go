package chezmoimaps

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func SortedKeys[K interface {
	comparable
	constraints.Ordered
}, V any](m map[K]V,
) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

func SortedStringerKeys[K interface {
	comparable
	fmt.Stringer
}, V any](m map[K]V,
) []string {
	keyStrs := make([]string, 0, len(m))
	for k := range m {
		keyStrs = append(keyStrs, k.String())
	}
	sort.Strings(keyStrs)
	return keyStrs
}

func SortedStringers[T fmt.Stringer](stringers []T) []string {
	values := make([]string, 0, len(stringers))
	for _, stringer := range stringers {
		values = append(values, stringer.String())
	}
	sort.Strings(values)
	return values
}
