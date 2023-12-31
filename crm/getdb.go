package crm

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(self, s)
}

func (self *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (self *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return self.client.Dial(network, address)
}

var CrmDB *sql.DB

func RunDB() (err error) {
	CrmDB, err = GetDB()
	return
}

var SshHost = ""
var SshPort = 22
var SshUser = ""
var DbUser = ""
var DbPass = ""
var DbHost = ""
var DbName = ""

var PrivateKeyPath = ""

func GetDB() (db *sql.DB, err error) {

	// Загрузите файл с ключом
	privateKey, err := loadPrivateKey(PrivateKeyPath)
	if err != nil {
		return
	}
	// Настройка параметров SSH-туннеля
	sshConfig := &ssh.ClientConfig{
		User: SshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH Server
	sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", SshHost, SshPort), sshConfig)
	if err != nil {
		return
	}

	// Now we register the ViaSSHDialer with the ssh connection as a parameter
	sql.Register("postgres+ssh", &ViaSSHDialer{sshcon})

	// And now we can use our new driver with the regular postgres connection string tunneled through the SSH connection
	db, err = sql.Open("postgres+ssh", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DbUser, DbPass, DbHost, DbName))
	if err != nil {
		return
	}
	return
}

func testDB() {
	sshHost := "ssh.example.com" // адрес SSH сервера
	sshPort := 22                // порт SSH
	sshUser := "ssh-user"        // пользователь SSH
	sshPassword := "ssh-pass"    // пароль SSH

	//dbHost := "127.0.0.1:3306" // адрес базы данных относительно SSH сервера
	localPort := "8080" // локальный порт для проброски

	// Настройка SSH клиента
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}

	// Подключение к SSH серверу
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Создание SSH туннеля
	listener, err := conn.Listen("tcp", "localhost:"+localPort)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// Подключение к базе данных через SSH туннель
	db, err := sqlx.Connect("mysql", "dbuser:dbpassword@tcp(localhost:"+localPort+")/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Выполнение SQL запроса
	var result string
	err = db.Get(&result, "SELECT 'Hello, World!'")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
