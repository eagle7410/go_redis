package lib

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	"reflect"
)

type env struct {
	HOST,
	PORT,
	AUTH_PASS string
}

func (i *env) Init () error {

	pwd, err := os.Getwd()

	fmt.Println(pwd)

	if err != nil {
		return err
	}


	envPath := path.Join(pwd, ".env")

	if _, err := os.Stat(envPath); err == nil {
		fmt.Println("Env load from file")
		err := godotenv.Load(envPath)

		if err != nil {
			return err;
		}
	}

	props := map[string]bool {
		"HOST"      : true,
		"PORT"      : true,
		"AUTH_PASS" : false,
	}

	for prop, isRequired := range props {

		v := os.Getenv(prop)

		if (isRequired == true && v == "") {
			return errors.New("Bad " + prop);
		}

		reflect.ValueOf(i).Elem().FieldByName(prop).SetString(v)
	}

	if (!FileExists("logs")) {
		err = os.Mkdir("logs", 0777)

		if err != nil {
			return err
		}
	}

	return nil;
}

var ENV env

