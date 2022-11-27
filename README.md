# Generate and sign security.txt file
[![MIT license](https://img.shields.io/github/license/korteke/gensectext.svg)](https://github.com/korteke/gensectext/blob/main/LICENSE.md)

More information can be found: [Securitytxt.org](https://securitytxt.org/)

This will generate signed security.txt file based on values from config.json. Config.json should be self-explanatory.

After generation one should need to place created `security.txt` file to web-server's .well-know -directory, so that it will be served on https://www.example.test/.well-known/security.txt 

# Usage
For signature, you need private PGP key and passphrase for that key
```
Usage of ./gensectext:
-configFile string
    Configuration file for template (default "config.json")
-generate
    Generate private GPG key
-passphrase string
    Passphrase for private GPG key
-privKey string
    Private GPG key (default "priv.key")'
-sign
    Sign security.txt with GPG (default true)
```
### Default usage - Generate security.txt with PGP signature
```
➜  gensectext git:(master) ✗ ./gensectext -privKey priv.key -passphrase RealSecretPassphrase
2022/11/27 02:35:36 Security.txt file(s) generated!
➜  gensectext git:(master) ✗
```

### Just generate file without signature
```
➜  gensectext git:(master) ✗ ./gensectext -sign=false
2022/11/27 02:39:43 Security.txt file(s) generated!
➜  gensectext git:(master) ✗
```

# Configuration
* security.tmpl is a template for the file, using [Go templating engine](https://pkg.go.dev/text/template)
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






