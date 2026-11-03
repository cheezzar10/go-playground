package lang

import (
	"fmt"
	"strings"
)

func makeNilSlice() []string {
	// return []string{"foo"}
	return nil
}

func IterateNilSlice() {
	names := makeNilSlice()
	if names == nil {
		fmt.Println("it's a 'nil' slice")
	}
	fmt.Printf("slice length: %d\n", len(names))
	fmt.Println("iterating nil slice")
	for index, name := range names {
		fmt.Printf("slice element [%d] = %s\n", index, name)
	}
}

func namedReturns(packageName string) (groupId string, artifactId string) {
	f := func(pn string) bool {
		if len(pn) > len(groupId) {
			return true
		}
		return false
	}

	groupId = "org.springframework.security.supersecurity"

	if f(packageName) {
		packageNameParts := strings.Split(packageName, ":")
		if len(packageNameParts) < 2 {
			return
		}
		return packageNameParts[0], packageNameParts[1]
	}

	fmt.Println("here")
	return
}

func NamedReturnsDemo() {
	var defStr string
	fmt.Printf("default string: '%s'\n", defStr)
	groupId, artifactId := namedReturns("org.springframework:spring-core")
	fmt.Printf("groupId length: %d\n", len(groupId))
	fmt.Printf("group: '%s', artifact: '%s'\n", groupId, artifactId)
}
