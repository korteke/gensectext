package main

import (
	"encoding/json"
	"flag"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

// Constants
const (
	secTextTemplate          = "security.tmpl"
	securityTextFileUnsigned = "security.txt.asc"
	securityTextFile         = "security.txt"
	rsaBits                  = 4096
	keyType                  = "rsa"
)

var (
	configFile   *string
	privateKey   *string
	passphrase   *string
	generateKeys *bool
	name         *string
	email        *string
	sign         *bool
)

type SecText struct {
	Contacts        []string
	Expires         string
	Acknowledgments string
	Languages       string
	Canonical       string
	Policy          string
	Encryption      string
	Hiring          string
}

func init() {
	configFile = flag.String("configFile", "config.json", "Configuration file for template")
	privateKey = flag.String("privKey", "priv.key", "Private PGP key")
	passphrase = flag.String("passphrase", "", "Passphrase for private PGP key")
	generateKeys = flag.Bool("generateKeys", false, "Generate private PGP key")
	name = flag.String("name", "", "Display name for PGP key")
	email = flag.String("email", "", "Email address for PGP key")
	sign = flag.Bool("sign", true, "Sign security.txt with PGP")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if *generateKeys {
		if err := generateKey(*name, *email, []byte(*passphrase)); err != nil {
			log.Fatalln("error creating new PGP private key")
		}
		log.Println("Generated private PGP key:", *privateKey)
		os.Exit(0)
	}

	file, err := os.Open(*configFile)
	checkErr(err)
	decoder := json.NewDecoder(file)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("Error closing file", file.Name(), err)
		}
	}(file)
	sectext := SecText{Expires: ExpiresTime()}

	err = decoder.Decode(&sectext)
	checkErr(err)

	f, err := os.Create(securityTextFileUnsigned)
	checkErr(err)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln("Error closing file", file.Name(), err)
		}
	}(f)

	t := template.Must(template.ParseFiles(secTextTemplate))
	err = t.Execute(f, sectext)
	checkErr(err)

	if *sign {
		unsignedFile, err := os.ReadFile(securityTextFileUnsigned)
		checkErr(err)
		privKey, err := readPrivatekey()
		checkErr(err)

		a, err := helper.SignCleartextMessageArmored(string(privKey), []byte(*passphrase), string(unsignedFile))
		checkErr(err)
		f, err = os.Create(securityTextFile)
		checkErr(err)
		if err = os.WriteFile(securityTextFile, []byte(a), 0644); err != nil {
			log.Fatalln(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(f)

		if err = removeHeaders(); err != nil {
			log.Fatalln(err)
		}
	}

	log.Print("Security.txt file(s) generated!")

}

func ExpiresTime() string {
	return time.Now().AddDate(0, 11, 0).UTC().Format("2006-01-02T15:04:05.000Z07:00")

}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func generateKey(name string, email string, pass []byte) error {
	rsaKey, err := helper.GenerateKey(name, email, pass, keyType, rsaBits)
	checkErr(err)
	if err := os.WriteFile(*privateKey, []byte(rsaKey), 0400); err != nil {
		return err
	}
	return nil
}

func readPrivatekey() ([]byte, error) {
	privKey, err := os.ReadFile(*privateKey)
	return privKey, err
}

func removeHeaders() error {
	input, err := os.ReadFile(securityTextFile)
	checkErr(err)
	re := regexp.MustCompile(`((?m)Version: .*|Comment: .*)`)

	lines := strings.Split(string(input), "\n")
	var newLines []string

	for _, line := range lines {
		if !re.MatchString(line) {
			newLines = append(newLines, line)
		}

	}
	output := strings.Join(newLines, "\n")
	err = os.WriteFile(securityTextFile, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
