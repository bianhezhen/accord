package accord

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/mistsys/accord/db"
)

func generateKey() ([]byte, error) {
	key := make([]byte, KeySize)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}

	return key, nil
}

// run tests
// - try encrypting and decrypting
// - check that the nonce is random enough
func TestAESGCMEndToEnd(t *testing.T) {
	psks := make(map[uint32][]byte)
	// test = random key generated by
	// < /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c32
	psks[1] = []byte(`JpUtbRukLuIFyjeKpA4fIpjgs6MTV8eH`)

	pskStore := db.NewLocalPSKStore(psks)
	aesgcm := InitAESGCM(pskStore)
	message := make([]byte, 100)
	_, err := io.ReadFull(rand.Reader, message[:])

	encrypted, err := aesgcm.Encrypt(message, 1)
	if err != nil {
		t.Errorf("Failed to encrypt. %s", err)
	}

	// Now lets try decrypting the same message and see if
	// we get the same message back

	decrypted, err := aesgcm.Decrypt(encrypted)
	if err != nil {
		t.Errorf("Failed to decrypt. %s", err)
	}

	if !bytes.Equal(decrypted, message) {
		t.Errorf("decrypted(encrypted(message)) != message")
		t.Fail()
	}
}

func BenchmarkAESEncrypt100BytesSingleKey(b *testing.B) {

	pskStore := db.NewDummyPSKStore()
	messages := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		msg := make([]byte, 100)
		_, err := io.ReadFull(rand.Reader, msg[:])
		if err != nil {
			b.Fatalf("Failed to allocate enough ")
		}
		messages[n] = msg
	}
	aesgcm := InitAESGCM(pskStore)
	for n := 0; n < b.N; n++ {
		_, err := aesgcm.Encrypt(messages[n], 160394189)
		if err != nil {
			b.Errorf("Failed to encrypt message[%d]", n)
		}
	}
}

func BenchmarkAESEncrypt100BytesNKeys(b *testing.B) {

	keys := make(map[uint32][]byte)
	messages := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		msg := make([]byte, 100)
		key := make([]byte, KeySize)
		_, err := io.ReadFull(rand.Reader, msg[:])
		if err != nil {
			b.Fatalf("Failed to allocate enough ")
		}
		_, err = io.ReadFull(rand.Reader, key[:])
		if err != nil {
			b.Fatalf("Failed to allocate enough ")
		}
		messages[n] = msg
		keys[uint32(n)] = key
	}
	pskStore := db.NewLocalPSKStore(keys)
	aesgcm := InitAESGCM(pskStore)
	for n := 0; n < b.N; n++ {
		_, err := aesgcm.Encrypt(messages[n], uint32(n))
		if err != nil {
			b.Errorf("Failed to encrypt message[%d]", n)
		}
	}
}
