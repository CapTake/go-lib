package data

import (
	"encoding/json"
	"time"
)

type ProtocolTicket struct {
	ID             int64                 `json:"id"`
	Ticketer       Address               `json:"ticketer"`
	RawType        json.RawMessage       `json:"rawType"`
	RawContent     json.RawMessage       `json:"rawContent"`
	Content        ProtocolTicketContent `json:"content"`
	TypeHash       int64                 `json:"typeHash"`
	ContentHash    int64                 `json:"contentHash"`
	FirstMinter    Address               `json:"firstMinter"`
	FirstLevel     uint64                `json:"firstLevel"`
	FirstTime      time.Time             `json:"firstTime"`
	LastLevel      uint64                `json:"lastLevel"`
	LastTime       time.Time             `json:"lastTime"`
	TransfersCount int64                 `json:"transfersCount"`
	BalancesCount  int64                 `json:"balancesCount"`
	HoldersCount   int64                 `json:"holdersCount"`
	TotalMinted    string                `json:"totalMinted"`
	TotalBurned    string                `json:"totalBurned"`
	TotalSupply    string                `json:"totalSupply"`
}

type ProtocolTicketContent struct {
	Protocol string `json:"string"`
	Payload  string `json:"bytes"`
}

type ProtocolTicketBalance struct {
	ID             int64          `json:"id"`
	Ticket         ProtocolTicket `json:"ticket"`
	Account        Address        `json:"account"`
	Balance        string         `json:"balance"`
	TransfersCount int64          `json:"transfersCount"`
	FirstLevel     uint64         `json:"firstLevel"`
	FirstTime      time.Time      `json:"firstTime"`
	LastLevel      uint64         `json:"lastLevel"`
	LastTime       time.Time      `json:"lastTime"`
}

type ProtocolTicketTransfer struct {
	ID            int64          `json:"id"`
	Level         uint64         `json:"level"`
	Timestamp     time.Time      `json:"timestamp"`
	Ticket        ProtocolTicket `json:"ticket"`
	To            Address        `json:"to"`
	Amount        string         `json:"amount"`
	TransactionID int64          `json:"transactionId"`
}
