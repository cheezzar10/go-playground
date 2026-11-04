package lang

import (
	"fmt"
	"strings"
	"unsafe"

	m "github.com/cheezzar10/go-playground/model"
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

func MapsDemo() {
	carId := "BMW:320d"
	cars := map[string]*m.Car{
		carId: {Maker: "BMW", Model: "320d"},
	}

	car := cars[carId]
	fmt.Printf("sizeof a pointer to a car: %d\n", unsafe.Sizeof(car))
	fmt.Printf("car with id %s = %+v\n", carId, car)

	cars["nil"] = nil
	nilCar := cars["nil"]
	fmt.Printf("nil car: %+v\n", nilCar)

	missingCar, ok := cars["missing"]
	fmt.Printf("missing car: %+v\n", missingCar)
	fmt.Printf("missing car found: %t\n", ok)
}

func ValuesMapDemo() {
	cars := map[string]m.Car{}
	cars["BMW:330d"] = m.Car{
		Maker: "BMW",
		Model: "330d",
	}

	// empty struct{}{} value has zero size

	car := cars["nil"]
	// it's not possible to compare car with nil
	fmt.Printf("car value: %+v\n", car)
	fmt.Printf("sizeof Car value: %d\n", unsafe.Sizeof(car))
}

func SliceOfLongsDemo() {
	nums := make([]int64, 4)
	fmt.Printf("longs type: %T\n", nums)

	fmt.Printf("longs length: %d\n", len(nums))
	fmt.Printf("longs capacity: %d\n", cap(nums))

	// experiment with slicing and growing in length without reallocation
}

func SliceCopyDemo() {
	src := []int32{1, 2, 3, 4, 5, 6}
	fmt.Println("source:", src)
	dst := make([]int32, 4, 16)
	fmt.Println("(before copy) destination:", dst)
	fmt.Printf("(before copy) destination length: %d, capacity: %d\n", len(dst), cap(dst))
	copied := copy(dst, src)
	fmt.Printf("%d elements copied\n", copied)
	fmt.Println("(after copy) destination:", dst)
	fmt.Printf("(after copy) destination length: %d, capacity: %d\n", len(dst), cap(dst))

}
