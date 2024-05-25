package ewt
import (    
	"crypto/ecdsa"    
	"encoding/json"    
	"fmt"    
	"log"  
	"strings"  
	"github.com/ethereum/go-ethereum/common/hexutil"    
	"github.com/ethereum/go-ethereum/crypto"
)
	// SignatureResponse represents the structure of the signature response.
type SignatureResponse struct {    
	Address string `json:"address,omitempty"`    
	Msg     string `json:"msg,omitempty"`    
	Sig     string `json:"sig,omitempty"`    
	Version string `json:"version,omitempty"`
}

// signMessage signs a message using the provided private key.
func signMessage(message, privKey string) (string, string) {    
	// Convert the private key from hex to ECDSA format    	
	ecdsaPrivateKey, err := crypto.HexToECDSA(privKey)    
	if err != nil {        
		log.Fatalln(err)    
	}    

	// Construct the message prefix    
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message)))    
	messageBytes := []byte(message)    

	// Hash the prefix and message using Keccak-256    
	hash := crypto.Keccak256Hash(prefix, messageBytes)    
	
	// Sign the hashed message    
	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)    
	if err != nil {        
		log.Fatalln(err)    
	}    

	// Adjust signature ID to Ethereum's format    
	sig[64] += 27    

	// Derive the public key from the private key    	
	publicKeyBytes := crypto.FromECDSAPub(ecdsaPrivateKey.Public().(*ecdsa.PublicKey))    
	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)    
	if err != nil {        
		log.Fatal(err)    
	}    
	rAddress := crypto.PubkeyToAddress(*pub)    

	// Construct the signature response    
	res := SignatureResponse{        
		Address: rAddress.String(),        
		Msg:     message,        
		Sig:     hexutil.Encode(sig),        
		Version: "2",    }    

	// Marshal the response to JSON with proper formatting    	
	resBytes, err := json.MarshalIndent(res, " ", "\t")    
	if err != nil {        
		log.Fatalln(err)    
	}    

	return res.Sig, string(resBytes)
}
// handleVerifySig verifies the signature against the provided public key and hash.
func verifySig(signature, address, message string) bool {    
	// Decode the signature into bytes    
	sig, err := hexutil.Decode(signature)    
	if err != nil {        
		log.Fatalln(err)    
	}    

	// Adjust signature to standard format (remove Ethereum's recovery ID)    
	sig[64] = sig[64] - 27    

	// Construct the message prefix    
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message)))    
	data := []byte(message)    
	
	// Hash the prefix and data using Keccak-256    
	hash := crypto.Keccak256Hash(prefix, data)    

	// Recover the public key bytes from the signature    	
	sigPublicKeyBytes, err := crypto.Ecrecover(hash.Bytes(), sig)    
	if err != nil {        
		log.Fatalln(err)    
	}    
	ecdsaPublicKey, err := crypto.UnmarshalPubkey(sigPublicKeyBytes)    
	if err != nil {        
		log.Fatalln(err)    
	}    

	// Derive the address from the recovered public key    
	rAddress := crypto.PubkeyToAddress(*ecdsaPublicKey)    
	
	// Check if the recovered address matches the provided address    
	isSigner := strings.EqualFold(rAddress.String(), address)    

	return isSigner
}