package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	input := "Mustafa Katipoğlu"

	md5hash := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()

	io.WriteString(md5hash, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
	sha512_256 := sha512.Sum512_256([]byte(input))

	// HMAC (Hash Message Authentication Code)
	hmac512 := hmac.New(sha512.New, []byte("Secret String"))
	hmac512.Write([]byte(input))

	// 6959d38b7f258e352afc9612a907c083
	fmt.Printf("md5: \t\t %x \n", md5hash.Sum(nil))

	// 033da2b3275b139b21632f0ad90ad5a1406961d5e20cacae4413c137e691f289
	fmt.Printf("sha256: \t\t %x \n", sha_256.Sum(nil))

	// 7e649faf2b3fc2c88a18d9425036ff8c07d2b77787d207edc80b648931a657bca1b1978a9a10b81888387f5315bb885a3e3cf0fe96bb7bfd735fe0ca78983327
	fmt.Printf("sha512: \t\t %x \n", sha_512.Sum(nil))

	// 82ed7db125b5dbe9d543705a38e8e6cd7b2dd10e0f3d58f117df538bc529c47d
	fmt.Printf("sha512_256: \t\t %x \n", sha512_256)

	// 7CRmSVX/i69f/hv6xwDJJotDMVneIi46X7ZrG4Hh31s3kEtIzTVqwivGDHacX+WD975BcQvrMxD5IzSOrtthww==
	fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))

}
