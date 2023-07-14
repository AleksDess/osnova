package crm

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func PingDbase(db *sql.DB) (err error) {
	return db.Ping()
}

func GetResult(db *sql.DB, z string) {
	rows, err := db.Query(z)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println(rows)
	}
}

func loadPrivateKey(path string) (ssh.Signer, error) {

	// Откройте файл с ключом
	privateKeyFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open private key file: %v", err)
	}
	defer privateKeyFile.Close()

	// Прочитайте содержимое файла с ключом
	privateKeyBytes, err := io.ReadAll(privateKeyFile)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	// Получите объект типа ssh.Signer из прочитанного ключа
	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	return privateKey, nil
}

func CRM() {

	privateKeyPath := "C:/Users/38099/.ssh/id_rsa"

	// Загрузите файл с ключом
	privateKey, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	// Настройка параметров SSH-туннеля
	sshConfig := &ssh.ClientConfig{
		User: "<developer>",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Установка соединения с SSH-сервером
	sshConn, err := ssh.Dial("tcp", "<18.185.123.177>:<22>", sshConfig)
	if err != nil {
		log.Fatalf("Failed to connect to SSH server: %v", err)
	}
	defer sshConn.Close()

	// Установка проброса порта к базе данных PostgreSQL
	localAddr := "localhost:5432"
	remoteAddr := "<10.11.0.33>:5432"
	sshListener, err := sshConn.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalf("Failed to listen on local address: %v", err)
	}
	defer sshListener.Close()

	log.Printf("Forwarding PostgreSQL from %s to %s...\n", localAddr, remoteAddr)

	// Установка соединения с базой данных PostgreSQL через SSH-туннель
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "<taxicrm>", "<Ziw51k0mztKoKBa019H5>", "<taxicrm>"))
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}
	defer db.Close()

	// Пример выполнения запроса к базе данных
	var result string
	err = db.QueryRow("select * from auto_parks order by updated_at desc limit 10").Scan(&result)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	fmt.Println(result)

}
