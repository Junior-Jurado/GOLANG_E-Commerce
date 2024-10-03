package variables

import (
	"fmt"
)

func MuestroEnteros() {
	var intComun int
	intde32 := int32(10);
	intde64 := int64(234);

	fmt.Println("intcomun =", intComun)
	fmt.Println("int32 =", intde32)
	fmt.Println("int64 =", intde64)
}