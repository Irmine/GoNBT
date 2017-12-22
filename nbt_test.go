package GoNBT

import "fmt"

func init() {
	fmt.Println("\n")

	var compound = NewCompound("test", map[string]INamedTag{
		"testShort": NewShort("testShort", 321),
		"testCompound": NewCompound("testCompound", map[string]INamedTag{
			"testFloat": NewFloat("testFloat", 321.34109),
			"nested": NewCompound("nested", map[string]INamedTag{
				"testString": NewString("testString", "This is a test string"),
			}),
			"list": NewList("list", TAG_String, []INamedTag{
				NewString("test", "Hi :D"),
			}),
		}),
	})

	compound.GetCompound("testCompound").GetCompound("nested").SetCompound("dynamicCompound", map[string]INamedTag{
		"testDynamic": NewList("testDynamic", TAG_Float, []INamedTag{}),
	})

	fmt.Println(compound.ToString())

	fmt.Println("\n")
}
