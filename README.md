# Generate and sign security.txt file (RFC-9116)
[![MIT license](https://img.shields.io/github/license/korteke/gensectext.svg)](https://github.com/korteke/gensectext/blob/main/LICENSE.md)
![Build](https://github.com/korteke/gensectext/actions/workflows/go.yaml/badge.svg)

More information about security.txt -file can be found: [Securitytxt.org](https://securitytxt.org/)

This program generates a signed security.txt file based on the values found in the config.json file. The config.json file should be self-explanatory.

The generated `security.txt` file should be placed on the web server so that it is served from the path ".well-known/security.txt", so the final address should be "https://www.example.test/.well-known/security.txt"
 
# Installation 
- [Download](https://github.com/korteke/gensectext/releases/latest) a prebuilt binary, unpack it and run!    
    _or_ 

- Compile with go `go install github.com/korteke/gensectext@latest` (the same command works for updating)

   _or_   
- `git clone https://github.com/korteke/gensectext ; cd gensectext ; go get ; go build .`   

  _or_   
- Pre-baked Docker image is available on [Docker hub](https://hub.docker.com/r/korteke/gensectext).
      ```
      docker pull korteke/gensectext:latest
      docker run --rm -v $(pwd):/app korteke/gensectext:latest -generateTmpl
      docker run --rm -v $(pwd):/app korteke/gensectext:latest -privKey /app/priv.key -passphrase "testtest"
      ```

# Usage
For digital signature, one need private PGP key and passphrase for that key.

`gensectext -h` shows the usage instructions 

```
Usage of ./gensectext:
  -configFile string
        Configuration file for template (default "config.json")
  -date string
    	Custom expires date. Format: YYYY-MM-DD (default now+1year)
  -email string
        Email address for PGP key
  -generateKeys
        Generate private PGP key
  -generateTmpl
    	Generate sample files
  -name string
        Display name for PGP key
  -passphrase string
        Passphrase for private PGP key
  -printPlain
    	Print unsigned file to stdout
  -printSig
    	Print signed file to stdout
  -privKey string
        Private PGP key (default "priv.key")
  -sign
        Sign security.txt with PGP (default true)
```

### Generate sample input files (Mandatory)
Generate sample input files, and follow [Usage](#usage) instructions.
```
➜  gensectext git:(main) ✗ ./gensectext -generateTmpl
2023/03/03 01:17:02 config.json created
2023/03/03 01:17:02 security.tmpl created
➜  gensectext git:(main) ✗
```
With docker you need to create these files manually to bind-mount directory.
### Generate private PGP key (Optional)
Generate a new private pgp key if you do not have one already.
```
➜  gensectext git:(main) ✗ ./gensectext -generateKeys -name "Test" -email "security@example.text" -passphrase testtest
2022/11/27 14:48:17 Generated private PGP key: priv.key
➜  gensectext git:(main) ✗ 
```

### Default usage - Generate security.txt with PGP signature
```
➜  gensectext git:(main) ✗ ./gensectext -privKey priv.key -passphrase RealSecretPassphrase
2022/11/27 02:35:36 Security.txt file(s) generated!
➜  gensectext git:(main) ✗
```

### Just generate file without signature
```
➜  gensectext git:(main) ✗ ./gensectext -sign=false
2022/11/27 02:39:43 Security.txt file(s) generated!
➜  gensectext git:(main) ✗
```

# Configuration
* security.tmpl is a template for the security.txt -file, using [Go templating engine](https://pkg.go.dev/text/template)
* config.json contains values for template. All RFC-fields should be supported. If you don't want something, just remove it from config.json. Contact and Expires fields are required, all others are optional.   
  
The Expires field is calculated to be 11 months from *time.Now()*

# Output
security.txt.asc (without signature)
```
Contact: mailto:security[at]EXAMPLE.com
Contact: https://hackerone.com/EXAMPLECO
Expires: 2023-10-26T23:54:05.428Z
Acknowledgments: https://hackerone.com/EXAMPLECO/thanks?type=team
Preferred-Languages: en, XX
Canonical: https://www.example.com/.well-known/security.txt
Policy: https://hackerone.com/EXAMPLECO/policy
```

security.txt (with signature)
```
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

Contact: mailto:security[at]EXAMPLE.com
Contact: https://hackerone.com/EXAMPLECO
Expires: 2023-10-26T23:54:05.428Z
Acknowledgments: https://hackerone.com/EXAMPLECO/thanks?type=team
Preferred-Languages: en, XX
Canonical: https://www.example.com/.well-known/security.txt
Policy: https://hackerone.com/EXAMPLECO/policy

-----BEGIN PGP SIGNATURE-----


wsBzBAABCgAnBQJjgrINCZD1ftKlLvOwJBYhBOQPlX7A9Tz3/LL/RPV+0qUu87Ak
AABB/AgA4cQ6eJeMv9EkcBABgtoVOmilnzixfQTJ31tYt3Y7Z+XyC4FDzdOAV4Yx
...
-----END PGP SIGNATURE-----
```






