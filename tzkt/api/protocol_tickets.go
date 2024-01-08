package api

import (
	"context"

	"github.com/dipdup-net/go-lib/tzkt/data"
)

// GetProtocolTicketsCount - Returns a total number of tickets.
func (tzkt *API) GetProtocolTicketsCount(ctx context.Context, filters map[string]string) (uint64, error) {
	return tzkt.count(ctx, "/v1/tickets/count", filters)
}

// GetProtocolTickets - Returns a list of tickets.
func (tzkt *API) GetProtocolTickets(ctx context.Context, filters map[string]string) (data []data.ProtocolTicket, err error) {
	err = tzkt.json(ctx, "/v1/tickets", filters, false, &data)
	return
}

// GetProtocolTicketBalancesCount - Returns a total number of ticket balances.
func (tzkt *API) GetProtocolTicketBalancesCount(ctx context.Context, filters map[string]string) (uint64, error) {
	return tzkt.count(ctx, "/v1/tickets/balances/count", filters)
}

// GetProtocolTicketBalances - Returns a list of ticket balances.
func (tzkt *API) GetProtocolTicketBalances(ctx context.Context, filters map[string]string) (data []data.TicketBalance, err error) {
	err = tzkt.json(ctx, "/v1/tickets/balances", filters, false, &data)
	return
}

// GetProtocolTicketTransfersCount - Returns the total number of ticket transfers.
func (tzkt *API) GetProtocolTicketTransfersCount(ctx context.Context, filters map[string]string) (uint64, error) {
	return tzkt.count(ctx, "/v1/tickets/transfers/count", filters)
}

// GetProtocolTicketTransfers - Returns a list of ticket transfers.
func (tzkt *API) GetProtocolTicketTransfers(ctx context.Context, filters map[string]string) (data []data.ProtocolTicketTransfer, err error) {
	err = tzkt.json(ctx, "/v1/tickets/transfers", filters, false, &data)
	return
}
