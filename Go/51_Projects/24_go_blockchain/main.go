package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var blockchain *Blockchain

type Block struct {
	Position     int
	Data         BookCheckout
	TimeStamp    string
	Hash         string
	PreviousHash string
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"release_date"`
	ISBN        string `json:"isbn"`
}

type Blockchain struct {
	Blocks []*Block
}

func (blockchain *Blockchain) AddBlock(data BookCheckout) {
	previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	block := createBlock(previousBlock, data)

	if validBlock(block, previousBlock) {
		blockchain.Blocks = append(blockchain.Blocks, block)
	}
}

func (block *Block) generateHash() {
	bytes, _ := json.Marshal(block.Data)
	data := fmt.Sprintf("%d%s%s%s", block.Position, block.TimeStamp, string(bytes), block.PreviousHash)

	hash := sha256.New()
	hash.Write([]byte(data))
	block.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (block *Block) validateHash(hash string) bool {
	block.generateHash()
	return block.Hash == hash
}

func createBlock(previousBlock *Block, checkoutItem BookCheckout) *Block {
	block := &Block{}
	block.Position = previousBlock.Position + 1
	block.TimeStamp = time.Now().String()
	block.PreviousHash = previousBlock.Hash
	block.Data = checkoutItem
	block.generateHash()
	return block
}

func validBlock(block, previousBlock *Block) bool {
	if previousBlock.Hash != block.PreviousHash {
		return false
	}

	if !block.validateHash(block.Hash) {
		return false
	}

	if previousBlock.Position+1 != block.Position {
		return false
	}

	return true
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(blockchain.Blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	io.WriteString(w, string(jbytes))
}

func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutItem BookCheckout

	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not write block: %v", err)
		w.Write([]byte("could not write block"))
		return
	}

	blockchain.AddBlock(checkoutItem)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("block written successfully"))
}

func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create: %v", err)
		w.Write([]byte("could not create new book"))
		return
	}

	hash := md5.New()
	io.WriteString(hash, book.ISBN+time.Now().String())
	book.ID = fmt.Sprintf("%x", hash.Sum(nil))

	resp, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not marshal payload: %v", err)
		w.Write([]byte("could not save book data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func genesisBlock() *Block {
	return createBlock(&Block{}, BookCheckout{IsGenesis: true})
}

func newBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{genesisBlock()}}
}

func main() {
	blockchain = newBlockchain()

	router := mux.NewRouter()
	router.HandleFunc("/blockchain", getBlockchain).Methods("GET")
	router.HandleFunc("/blockchain", writeBlock).Methods("POST")
	router.HandleFunc("/newbook", newBook).Methods("POST")

	go func() {
		for _, block := range blockchain.Blocks {
			fmt.Printf("Previous hash: %x\n", block.PreviousHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}()

	port := ":3000"

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
