package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Leer variables
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASS")
	to := []string{os.Getenv("EMAIL_TO")}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Mensaje
	message := []byte("Subject: Prueba desde Go\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n\r\n" +
		"Hola,\n\nEste es un correo enviado desde Golang usando net/smtp.\n\nSaludos.")

	// Autenticación
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Enviar correo
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Correo enviado exitosamente")
}
