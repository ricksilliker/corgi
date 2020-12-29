package corgi

import (
	"github.com/nknorg/nkn-sdk-go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func InitializeClient() bool {
	var seed []byte

	path, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Panic("missing home directory env var")
		return false
	}
	path = path + string(os.PathSeparator) + ".corgi"

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		account, err = nkn.NewAccount(nil)
		if err != nil {
			logrus.WithError(err).Panic("something went wrong creating account")
		}
		seed = account.Seed()
	} else {
		f, err = os.Open(path)
		seed, err = ioutil.ReadAll(f)

	}

	account, err := nkn.NewAccount(seed)
	if err != nil {
		logrus.WithError(err).Panic("failed to initialize transaction account")
		return false
	}

}