package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	nakup = 700
	dokup = 6000
)

var (
	norm  = map[int]bool{1: true, 32: true, 60: true, 91: true, 121: true, 152: true, 182: true, 213: true, 244: true, 274: true, 305: true, 335: true}
	prest = map[int]bool{1: true, 32: true, 61: true, 92: true, 122: true, 153: true, 183: true, 214: true, 245: true, 275: true, 306: true, 336: true}
	nacon = map[int]bool{0: true, 3: true}
)

func main() {
	balance := 6000
	year := norm
	for d := 1; d <= 365; d++ {
		if _, ok := nacon[d%7]; ok {
			balance -= nakup
			if balance < 0 {
				log.Fatal("bust below zero")
			}
			log.Infof("%02d: nakupuju %d; bal %d", d, nakup, balance)
		}
		if _, ok := year[d]; ok {
			balance += dokup
			log.Infof("%02d: dokup %d; bal %d", d, dokup, balance)
		}
	}
	calcSig()
}

func calcSig() {
	clientId := "6"
	publicKey := "CpmRVUJL0OGByT2otAfCKeeDdU6yfi6OzvnXcAwaHvE"
	nonce := strconv.FormatInt(time.Now().Unix(), 10)
	secret := "secret"
	data := fmt.Sprintf("%s%s%s", nonce, clientId, publicKey)
	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	signature := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	fmt.Println("Result: " + signature)
}
