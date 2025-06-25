/*
8. Cree un programa que maneje una lista de contactos de manera
concurrente. La lista de contactos debe permitir agregar, eliminar y
buscar contactos de manera segura desde múltiples goroutines.

a) Defina una estructura Contact que contenga campos como
Nombre, Apellido, CorreoElectronico, y Telefono.

b) Cree una estructura llamada Agenda que contenga un mapa de
Contact con el correo electrónico como clave.

c) Implemente los siguientes métodos para la estructura Agenda:

i. AgregarContacto(contacto Contact): Agrega un
nuevo contacto a la agenda.

ii. EliminarContacto(correo string): Elimina un
contacto de la agenda dado su correo electrónico.

iii. BuscarContacto(correo string) Contact: Busca y
devuelve un contacto dado su correo electrónico.

d) Asegúrese de que las operaciones de agregar, eliminar y buscar
contactos se realicen de manera concurrente y que la estructura
Agenda sea segura para ser accedida desde múltiples
goroutines.

e) Cree una función main() que cree una agenda, inicie varias
goroutines para agregar, eliminar y buscar contactos de manera
simultánea, y luego imprima el contenido de la agenda después
de un tiempo para verificar que las operaciones se hayan
realizado correctamente. */

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type contact struct {
	nom    string
	ape    string
	correo string
	tel    string
}

var (
	agenda   = make(map[string]contact)
	semaforo sync.Mutex
)

func AgregarContacto(contact contact) {
	semaforo.Lock()
	defer semaforo.Unlock()
	agenda[contact.tel] = contact
}

func EliminarContacto(correo string) {
	semaforo.Lock()
	defer semaforo.Unlock()
	for clave, contact := range agenda {
		if contact.correo == correo {
			delete(agenda, clave)
			break
		}
	}
}

func BuscarContacto(correo string) (contact, error) {
	semaforo.Lock()
	defer semaforo.Unlock()
	for _, contact := range agenda {
		if contact.correo == correo {
			return contact, nil
		}
	}
	return contact{}, errors.New("No existe el contacto ")
}

func main() {
	go AgregarContacto(contact{"Pepe", "Argento", "pepeargento@gmail.com", "+5491112345678"})
	go AgregarContacto(contact{"Moni", "Argento", "moniargento@gmail.com", "+5491187654321"})
	go AgregarContacto(contact{"Coqui", "Argento", "coquiargento@gmail.com", "+5491186427531"})
	go AgregarContacto(contact{"Paola", "Argento", "paolaargento@gmail.com", "+5491113572468"})
	go EliminarContacto("moniargento@gmail.com")
	go EliminarContacto("paolaargento@gmail.com")
	go func() {
		contact, err := BuscarContacto("pepeargento@gmail.com")
		if err == nil {
			fmt.Println(contact.nom, contact.ape, "se encuentra en la agenda")
		} else {
			fmt.Println(err)
		}
	}()
	go func() {
		contact, err := BuscarContacto("moniargento@gmail.com")
		if err == nil {
			fmt.Println(contact.nom, contact.ape, "se encuentra en la agenda")
		} else {
			fmt.Println(err)
		}
	}()
	go func() {
		contact, err := BuscarContacto("coquiargento@gmail.com")
		if err == nil {
			fmt.Println(contact.nom, contact.ape, "se encuentra en la agenda")
		} else {
			fmt.Println(err)
		}
	}()
	go func() {
		contact, err := BuscarContacto("paolaargento@gmail.com")
		if err == nil {
			fmt.Println(contact.nom, contact.ape, "se encuentra en la agenda")
		} else {
			fmt.Println(err)
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("----- Contenido de la Agenda -----")
	for _, contact := range agenda {
		fmt.Println(contact.nom, contact.ape)
	}
}
