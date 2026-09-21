package main

import (
	"fmt"
	"wallet/internal/wallet"
)

func main() {
	w := wallet.New()

	acc1, _ := w.CreateAccount("acc-1", "Anna")
	acc2, _ := w.CreateAccount("acc-2", "Boris")

	_ = acc1.Deposit(100000) // 1000.00

	fmt.Println("=== до переводов ===")
	fmt.Println("acc-1:", acc1.Balance)
	fmt.Println("acc-2:", acc2.Balance)

	// 1. перевод самому себе
	err := w.Transfer("acc-1", "acc-1", 5000)
	fmt.Println("\n1. самому себе:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 2. отрицательная сумма
	err = w.Transfer("acc-1", "acc-2", -100)
	fmt.Println("\n2. отрицательная сумма:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 3. несуществующий получатель
	err = w.Transfer("acc-1", "acc-99", 5000)
	fmt.Println("\n3. нет получателя:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 4. не хватает денег
	err = w.Transfer("acc-1", "acc-2", 999999)
	fmt.Println("\n4. мало денег:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 5. успешный перевод
	err = w.Transfer("acc-1", "acc-2", 30000)
	fmt.Println("\n5. успех:", err)
	fmt.Println("   acc-1:", acc1.Balance)
	fmt.Println("   acc-2:", acc2.Balance)
}
