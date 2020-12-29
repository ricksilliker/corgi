package corgi

import (
	"github.com/nknorg/nkn-sdk-go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"os"
)

func InitializeClient() bool {
	var seed []byte

	path, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Panic("missing home directory env var")
		return false
	}
	path = path + string(os.PathSeparator) + ".corgi" + "account.corgi"

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		account, err := nkn.NewAccount(nil)
		if err != nil {
			logrus.WithError(err).Panic("something went wrong creating account")
		}
		seed = account.Seed()

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			logrus.WithError(err).Panic("failed to create new account file")
			return false
		}

		if _, err := f.Write(seed); err != nil {
			logrus.WithError(err).Panic("failed to write new account file")
			return false
		}

		if err := f.Close(); err != nil {
			logrus.WithError(err).Panic("failed to close new account file")
			return false
		}

	} else {
		f, err := os.Open(path)
		if err != nil {
			logrus.WithError(err).Panic("trouble opening existing account")
			return false
		}
		seed, err = ioutil.ReadAll(f)
		if err != nil {
			logrus.WithError(err).Panic("trouble reading existing account")
			return false
		}
	}

	account, err := nkn.NewAccount(seed)
	if err != nil {
		logrus.WithError(err).Panic("failed to initialize transaction account")
		return false
	}

	var client *nkn.MultiClient
	client, err = nkn.NewMultiClient(account, "", 3, false, &nkn.ClientConfig{
		ConnectRetries: 1000,
	})
	if err != nil {
		PushError(rt, err.Error(), "do you have an active internet connection?")
	}

	<-client.OnConnect.C

	logrus.Info("client connected")

	client.Listen(nil)
	go Listen(client)

	return true
}

func Listen(client *nkn.MultiClient) {
	for !client.IsClosed() {
		session, err := client.Accept()
		if err != nil {
			logrus.WithError(err).Error("error on client accept")
			continue
		}
		AcceptSession(session)
	}
}

func AcceptSession(session net.Conn) {
	addr := session.RemoteAddr().String()
	logrus.Info("session address: " + addr)
	//reader := bufio.NewReader(session)
}