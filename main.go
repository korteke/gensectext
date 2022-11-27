package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	configFile *string
	privateKey *string
	passphrase *string
	generate   *bool
	sign       *bool
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
	privateKey = flag.String("privKey", "priv.key", "Private GPG key")
	passphrase = flag.String("passphrase", "", "Passphrase for private GPG key")
	generate = flag.Bool("generate", false, "Generate private GPG key")
	sign = flag.Bool("sign", true, "Sign security.txt with GPG")
}

func main() {
	flag.Parse()

	if *generate == true {
		fmt.Println("WIP, this will generate key in future!")
	}

	file, err := os.Open(*configFile)
	checkErr(err)
	decoder := json.NewDecoder(file)
	defer file.Close()
	sectext := SecText{Expires: ExpiresTime()}

	err = decoder.Decode(&sectext)
	checkErr(err)

	f, err := os.Create(securityTextFileUnsigned)
	checkErr(err)
	defer f.Close()

	t := template.Must(template.ParseFiles(secTextTemplate))
	err = t.Execute(f, sectext)
	checkErr(err)

	if *sign == true {
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

// TODO: Need to implement
func generateKey(name string, email string, pass []byte) string {
	rsaKey, err := helper.GenerateKey(name, email, pass, keyType, rsaBits)
	checkErr(err)
	return rsaKey
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
	for i, line := range lines {
		lines[i] = strings.Trim(re.ReplaceAllString(line, ""), "\r\n")
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(securityTextFile, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
