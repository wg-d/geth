package etherpc

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/eth-go/ethpub"
	_ "log"
)

type MainPackage struct {
	ethp *ethpub.PEthereum
}

type JsonArgs interface {
	requirements() error
}

type BlockResponse struct {
	JsonResponse
}
type GetBlockArgs struct {
	BlockNumber int
	Hash        string
}

type ErrorResponse struct {
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type JsonResponse interface {
}

type SuccessRes struct {
	Error  bool         `json:"error"`
	Result JsonResponse `json:"result"`
}

func NewSuccessRes(object JsonResponse) string {
	e := SuccessRes{Error: false, Result: object}
	res, err := json.Marshal(e)
	if err != nil {
		// This should never happen
		panic("Creating json error response failed, help")
	}
	success := string(res)
	return success
}

func NewErrorResponse(msg string) error {
	e := ErrorResponse{Error: true, ErrorText: msg}
	res, err := json.Marshal(e)
	if err != nil {
		// This should never happen
		panic("Creating json error response failed, help")
	}
	newErr := errors.New(string(res))
	return newErr
}

func (b *GetBlockArgs) requirements() error {
	if b.BlockNumber == 0 && b.Hash == "" {
		return NewErrorResponse("GetBlock requires either a block 'number' or a block 'hash' as argument")
	}
	return nil
}

func (p *MainPackage) GetBlock(args *GetBlockArgs, reply *string) error {
	err := args.requirements()
	if err != nil {
		return err
	}
	// Do something
	block := p.ethp.GetBlock(args.Hash)
	*reply = NewSuccessRes(block)
	return nil
}

type NewTxArgs struct {
	Sec       string
	Recipient string
	Value     string
	Gas       string
	GasPrice  string
	Init      string
	Body      string
}
type TxResponse struct {
	Hash string
}

func (a *NewTxArgs) requirements() error {
	if a.Recipient == "" {
		return NewErrorResponse("Transact requires a 'recipient' address as argument")
	}
	if a.Value == "" {
		return NewErrorResponse("Transact requires a 'value' as argument")
	}
	if a.Gas == "" {
		return NewErrorResponse("Transact requires a 'gas' value as argument")
	}
	if a.GasPrice == "" {
		return NewErrorResponse("Transact requires a 'gasprice' value as argument")
	}
	return nil
}

func (a *NewTxArgs) requirementsContract() error {
	if a.Value == "" {
		return NewErrorResponse("Create requires a 'value' as argument")
	}
	if a.Gas == "" {
		return NewErrorResponse("Create requires a 'gas' value as argument")
	}
	if a.GasPrice == "" {
		return NewErrorResponse("Create requires a 'gasprice' value as argument")
	}
	if a.Init == "" {
		return NewErrorResponse("Create requires a 'init' value as argument")
	}
	if a.Body == "" {
		return NewErrorResponse("Create requires a 'body' value as argument")
	}
	return nil
}

func (p *MainPackage) Transact(args *NewTxArgs, reply *string) error {
	err := args.requirements()
	if err != nil {
		return err
	}
	result, _ := p.ethp.Transact(p.ethp.GetKey().PrivateKey, args.Recipient, args.Value, args.Gas, args.GasPrice, args.Body)
	*reply = NewSuccessRes(result)
	return nil
}

func (p *MainPackage) Create(args *NewTxArgs, reply *string) error {
	err := args.requirementsContract()
	if err != nil {
		return err
	}
	result, _ := p.ethp.Create(p.ethp.GetKey().PrivateKey, args.Value, args.Gas, args.GasPrice, args.Init, args.Body)
	*reply = NewSuccessRes(result)
	return nil
}

func (p *MainPackage) getKey(args interface{}, reply *string) error {
	return nil
}

type GetStorageArgs struct {
	Address string
	Key     string
}

func (a *GetStorageArgs) requirements() error {
	if a.Address == "" {
		return NewErrorResponse("GetStorageAt requires an 'address' value as argument")
	}
	if a.Key == "" {
		return NewErrorResponse("GetStorageAt requires an 'key' value as argument")
	}
	return nil
}

func (p *MainPackage) getStorageAt(args *GetStorageArgs, reply *string) error {
	err := args.requirements()
	if err != nil {
		return err
	}
	return nil
}

type GetBalanceArgs struct {
	Address string
}

func (a *GetBalanceArgs) requirements() error {
	if a.Address == "" {
		return NewErrorResponse("GetBalanceAt requires an 'address' value as argument")
	}
	return nil
}

func (p *MainPackage) GetBalanceAt(args *GetBalanceArgs, reply *string) error {
	err := args.requirements()
	if err != nil {
		return err
	}
	return nil
}

type TestRes struct {
	JsonResponse `json:"-"`
	Answer       int `json:"answer"`
}

func (p *MainPackage) Test(args *GetBlockArgs, reply *string) error {
	*reply = NewSuccessRes(TestRes{Answer: 15})
	return nil
}
