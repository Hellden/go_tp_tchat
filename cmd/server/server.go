package main

import (
	"fmt"
	"net"
)

func gestionErreur(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	IP   = "127.0.0.1"
	PORT = "3569"
)

func main() {
	fmt.Println("Lancement du serveur...")

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	gestionErreur(err)

	conn, err := ln.Accept()
	gestionErreur(err)

	fmt.Println("Un client est connecté depuis", conn.RemoteAddr())
	// boucle pour toujours écouter les connexions entrantes (ctrl-c pour quitter)
	for {
		// On écoute les messages émis par les clients
		buffer := make([]byte, 4096)       // taille maximum du message qui sera envoyé par le client
		length, err := conn.Read(buffer)   // lire le message envoyé par client
		message := string(buffer[:length]) // supprimer les bits qui servent à rien et convertir les bytes en string

		if err != nil {
			fmt.Println("Le client s'est déconnecté")
		}

		// on affiche le message du client en le convertissant de byte à string
		fmt.Print("Client:", message)

		// On envoie le message au client pour qu'il l'affiche
		conn.Write([]byte("Le serveur te dit : " + message + "\n"))
	}
}
