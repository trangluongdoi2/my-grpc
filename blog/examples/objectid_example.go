package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Tạo một ObjectID mới
	oid := primitive.NewObjectID()

	fmt.Println("=== So sánh oid.Hex() vs oid.String() ===\n")

	// 1. Sử dụng Hex() - trả về hex string thuần
	hexString := oid.Hex()
	fmt.Printf("oid.Hex():    %s\n", hexString)
	fmt.Printf("Type:         %T\n", hexString)
	fmt.Printf("Length:       %d ký tự\n\n", len(hexString))

	// 2. Sử dụng String() - trả về có wrapper
	stringRep := oid.String()
	fmt.Printf("oid.String(): %s\n", stringRep)
	fmt.Printf("Type:         %T\n", stringRep)
	fmt.Printf("Length:       %d ký tự\n\n", len(stringRep))

	// 3. So sánh độ dài
	fmt.Printf("Chênh lệch:   %d ký tự (do có thêm 'ObjectID(\"\")')\n\n", len(stringRep)-len(hexString))

	// 4. Convert từ Hex string về ObjectID
	fmt.Println("=== Convert ngược lại ===\n")

	// Từ Hex() -> ObjectID: ✓ DỄ DÀNG
	newOid, err := primitive.ObjectIDFromHex(hexString)
	if err == nil {
		fmt.Printf("✓ Từ Hex():    %v -> %v\n", hexString, newOid)
	}

	// Từ String() -> ObjectID: ✗ PHỨC TẠP (phải strip wrapper)
	_, err = primitive.ObjectIDFromHex(stringRep)
	if err != nil {
		fmt.Printf("✗ Từ String(): Cannot convert directly!\n")
		fmt.Printf("  Error: %v\n\n", err)
	}

	// 5. Use cases thực tế
	fmt.Println("=== Use Cases ===\n")
	fmt.Printf("Hex() dùng cho:\n")
	fmt.Printf("  - Gửi qua API/gRPC response\n")
	fmt.Printf("  - Lưu trong JSON\n")
	fmt.Printf("  - So sánh ID với string\n")
	fmt.Printf("  - Query MongoDB với ID string\n\n")

	fmt.Printf("String() dùng cho:\n")
	fmt.Printf("  - Logging/debugging\n")
	fmt.Printf("  - Hiển thị cho developer\n")
	fmt.Printf("  - Debug output rõ ràng hơn\n")
}
