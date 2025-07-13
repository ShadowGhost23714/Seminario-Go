/* 2. Implemente una blockchain para que sea soporte de una cryptomoneda que
incluya la creación de billeteras para los clientes. Una blockchain, o cadena de
bloques, es un sistema digital distribuído que funciona como un libro de
contabilidad público e inmutable. Almacena información sobre transacciones de
forma segura y descentralizada, sin necesidad de intermediarios. Cada
transacción se agrupa en un bloque, que se enlaza con el bloque anterior,
creando una cadena irrompible.

Utilice structs para representar la transacción (con el monto, el identificador de
quien envía dinero, el identificador de quien recibe ese dinero y la fecha/hora
de la transacción - timestamp -), el bloque (que tienen el hash, el hash previo,
la transacción (data) y la fecha/hora de creación de dicho bloque), la cadena de
bloques y la billetera (con el identificador, nombre y apellido del cliente).

Tip: puede utilizar la librería crypto/sha256 para crear el hash del bloque
anterior.

a) Defina todos los tipos de datos que va a utilizar.

b) Programe funciones para:

i) Crear una billetera

ii) Enviar una transacción

iii) Insertar un bloque en la cadena

iv) Obtener el saldo de un usuario (recorriendo toda la cadena)

v) Realizar una función que valide que la cadena sea consistente
recorriéndola y verificando que el hash almacenado del bloque
anterior es válido

vi) Si utilizó un slice para la estructura de la cadena de bloques
cambie la implementación con una lista enlazada. Puede reutilizar
la implementación del ejercicio 9. ¿Cuál fue el impacto que tuvo en
su programa?

i) ¿Cómo validar que la transacción solo se pueda hacer si la
billetera que va a enviar los fondos tiene saldo suficiente? */

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Transacción
type Transaction struct {
	Amount    float64
	From      string
	To        string
	Timestamp time.Time
}

// Bloque
type Block struct {
	PreviousHash string
	Hash         string
	Timestamp    time.Time
	Data         Transaction
	Next         *Block // para usarlo como lista enlazada
}

// Blockchain (como lista enlazada)
type Blockchain struct {
	Genesis *Block
	Tail    *Block
}

// Billetera
type Wallet struct {
	ID        string
	FirstName string
	LastName  string
}

func CreateWallet(id, firstName, lastName string) Wallet { // Punto i
	return Wallet{ID: id, FirstName: firstName, LastName: lastName}
}

func NewTransaction(amount float64, from, to string) Transaction { // Punto ii
	return Transaction{Amount: amount, From: from, To: to, Timestamp: time.Now()}
}

func (bc *Blockchain) AddBlock(tx Transaction) { // Punto iii
	var prevHash string
	if bc.Tail != nil {
		prevHash = bc.Tail.Hash
	}

	block := Block{
		PreviousHash: prevHash,
		Timestamp:    time.Now(),
		Data:         tx,
	}
	block.Hash = CalculateHash(block)

	if bc.Genesis == nil {
		bc.Genesis = &block
		bc.Tail = &block
	} else {
		bc.Tail.Next = &block
		bc.Tail = &block
	}
}

func CalculateHash(b Block) string {
	record := fmt.Sprintf("%s%f%s%s", b.PreviousHash, b.Data.Amount, b.Data.From, b.Data.To)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (bc *Blockchain) GetBalance(walletID string) float64 { // Punto iv
	balance := 0.0
	current := bc.Genesis
	for current != nil {
		if current.Data.To == walletID {
			balance += current.Data.Amount
		}
		if current.Data.From == walletID {
			balance -= current.Data.Amount
		}
		current = current.Next
	}
	return balance
}

func (bc *Blockchain) IsValid() bool { // // Punto v
	current := bc.Genesis
	for current != nil && current.Next != nil {
		if current.Next.PreviousHash != current.Hash {
			return false
		}
		expectedHash := CalculateHash(*current)
		if current.Hash != expectedHash {
			return false
		}
		current = current.Next
	}
	return true
}

func (bc *Blockchain) CanTransact(fromID string, amount float64) bool { // Punto vi
	return bc.GetBalance(fromID) >= amount
}

func main() {
	bc := &Blockchain{}

	// Crear billeteras
	alice := CreateWallet("A1", "Alice", "Doe")
	bob := CreateWallet("B1", "Bob", "Smith")

	// Dar saldo inicial a Alice
	tx0 := NewTransaction(100.0, "", alice.ID)
	bc.AddBlock(tx0)

	// Alice transfiere a Bob
	if bc.CanTransact(alice.ID, 30.0) {
		tx1 := NewTransaction(30.0, alice.ID, bob.ID)
		bc.AddBlock(tx1)
	}

	fmt.Println("Saldo Alice:", bc.GetBalance(alice.ID))
	fmt.Println("Saldo Bob:", bc.GetBalance(bob.ID))
	fmt.Println("Blockchain válida?", bc.IsValid())
}
