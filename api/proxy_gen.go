// Code generated by github.com/filecoin-project/boost/gen/api. DO NOT EDIT.

package api

import (
	"context"
	"errors"

	smtypes "github.com/filecoin-project/boost/storagemarket/types"
	"github.com/filecoin-project/go-address"
	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-fil-markets/piecestore"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

var ErrNotSupported = errors.New("method not supported")

type BoostStruct struct {
	CommonStruct

	NetStruct

	Internal struct {
		ActorSectorSize func(p0 context.Context, p1 address.Address) (abi.SectorSize, error) `perm:"read"`

		BoostDagstoreGC func(p0 context.Context) ([]DagstoreShardResult, error) `perm:"admin"`

		BoostDagstoreInitializeAll func(p0 context.Context, p1 DagstoreInitializeAllParams) (<-chan DagstoreInitializeAllEvent, error) `perm:"admin"`

		BoostDagstoreInitializeShard func(p0 context.Context, p1 string) error `perm:"admin"`

		BoostDagstoreListShards func(p0 context.Context) ([]DagstoreShardInfo, error) `perm:"read"`

		BoostDagstoreRecoverShard func(p0 context.Context, p1 string) error `perm:"admin"`

		BoostDagstoreRegisterShard func(p0 context.Context, p1 string) error `perm:"admin"`

		BoostDeal func(p0 context.Context, p1 uuid.UUID) (*smtypes.ProviderDealState, error) `perm:"admin"`

		BoostDummyDeal func(p0 context.Context, p1 smtypes.DealParams) (*ProviderDealRejectionInfo, error) `perm:"admin"`

		BoostIndexerAnnounceAllDeals func(p0 context.Context) error `perm:"admin"`

		BoostOfflineDealWithData func(p0 context.Context, p1 uuid.UUID, p2 string) (*ProviderDealRejectionInfo, error) `perm:"admin"`

		DealsConsiderOfflineRetrievalDeals func(p0 context.Context) (bool, error) `perm:"admin"`

		DealsConsiderOfflineStorageDeals func(p0 context.Context) (bool, error) `perm:"admin"`

		DealsConsiderOnlineRetrievalDeals func(p0 context.Context) (bool, error) `perm:"admin"`

		DealsConsiderOnlineStorageDeals func(p0 context.Context) (bool, error) `perm:"admin"`

		DealsConsiderUnverifiedStorageDeals func(p0 context.Context) (bool, error) `perm:"admin"`

		DealsConsiderVerifiedStorageDeals func(p0 context.Context) (bool, error) `perm:"admin"`

		DealsPieceCidBlocklist func(p0 context.Context) ([]cid.Cid, error) `perm:"admin"`

		DealsSetConsiderOfflineRetrievalDeals func(p0 context.Context, p1 bool) error `perm:"admin"`

		DealsSetConsiderOfflineStorageDeals func(p0 context.Context, p1 bool) error `perm:"admin"`

		DealsSetConsiderOnlineRetrievalDeals func(p0 context.Context, p1 bool) error `perm:"admin"`

		DealsSetConsiderOnlineStorageDeals func(p0 context.Context, p1 bool) error `perm:"admin"`

		DealsSetConsiderUnverifiedStorageDeals func(p0 context.Context, p1 bool) error `perm:"admin"`

		DealsSetConsiderVerifiedStorageDeals func(p0 context.Context, p1 bool) error `perm:"admin"`

		DealsSetPieceCidBlocklist func(p0 context.Context, p1 []cid.Cid) error `perm:"admin"`

		MarketCancelDataTransfer func(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error `perm:"write"`

		MarketDataTransferUpdates func(p0 context.Context) (<-chan lapi.DataTransferChannel, error) `perm:"write"`

		MarketGetAsk func(p0 context.Context) (*storagemarket.SignedStorageAsk, error) `perm:"read"`

		MarketGetRetrievalAsk func(p0 context.Context) (*retrievalmarket.Ask, error) `perm:"read"`

		MarketImportDealData func(p0 context.Context, p1 cid.Cid, p2 string) error `perm:"write"`

		MarketListDataTransfers func(p0 context.Context) ([]lapi.DataTransferChannel, error) `perm:"write"`

		MarketListIncompleteDeals func(p0 context.Context) ([]storagemarket.MinerDeal, error) `perm:"read"`

		MarketListRetrievalDeals func(p0 context.Context) ([]retrievalmarket.ProviderDealState, error) `perm:"read"`

		MarketRestartDataTransfer func(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error `perm:"write"`

		MarketSetAsk func(p0 context.Context, p1 types.BigInt, p2 types.BigInt, p3 abi.ChainEpoch, p4 abi.PaddedPieceSize, p5 abi.PaddedPieceSize) error `perm:"admin"`

		MarketSetRetrievalAsk func(p0 context.Context, p1 *retrievalmarket.Ask) error `perm:"admin"`

		PiecesGetCIDInfo func(p0 context.Context, p1 cid.Cid) (*piecestore.CIDInfo, error) `perm:"read"`

		PiecesGetPieceInfo func(p0 context.Context, p1 cid.Cid) (*piecestore.PieceInfo, error) `perm:"read"`

		PiecesListCidInfos func(p0 context.Context) ([]cid.Cid, error) `perm:"read"`

		PiecesListPieces func(p0 context.Context) ([]cid.Cid, error) `perm:"read"`

		RuntimeSubsystems func(p0 context.Context) (lapi.MinerSubsystems, error) `perm:"read"`
	}
}

type BoostStub struct {
	CommonStub

	NetStub
}

type ChainIOStruct struct {
	Internal struct {
		ChainHasObj func(p0 context.Context, p1 cid.Cid) (bool, error) ``

		ChainReadObj func(p0 context.Context, p1 cid.Cid) ([]byte, error) ``
	}
}

type ChainIOStub struct {
}

type CommonStruct struct {
	Internal struct {
		AuthNew func(p0 context.Context, p1 []auth.Permission) ([]byte, error) `perm:"admin"`

		AuthVerify func(p0 context.Context, p1 string) ([]auth.Permission, error) `perm:"read"`

		LogList func(p0 context.Context) ([]string, error) `perm:"write"`

		LogSetLevel func(p0 context.Context, p1 string, p2 string) error `perm:"write"`

		Shutdown func(p0 context.Context) error `perm:"admin"`
	}
}

type CommonStub struct {
}

type CommonNetStruct struct {
	CommonStruct

	NetStruct

	Internal struct {
	}
}

type CommonNetStub struct {
	CommonStub

	NetStub
}

type NetStruct struct {
	Internal struct {
		ID func(p0 context.Context) (peer.ID, error) `perm:"read"`

		NetAddrsListen func(p0 context.Context) (peer.AddrInfo, error) `perm:"read"`

		NetAgentVersion func(p0 context.Context, p1 peer.ID) (string, error) `perm:"read"`

		NetAutoNatStatus func(p0 context.Context) (NatInfo, error) `perm:"read"`

		NetBandwidthStats func(p0 context.Context) (metrics.Stats, error) `perm:"read"`

		NetBandwidthStatsByPeer func(p0 context.Context) (map[string]metrics.Stats, error) `perm:"read"`

		NetBandwidthStatsByProtocol func(p0 context.Context) (map[protocol.ID]metrics.Stats, error) `perm:"read"`

		NetBlockAdd func(p0 context.Context, p1 NetBlockList) error `perm:"admin"`

		NetBlockList func(p0 context.Context) (NetBlockList, error) `perm:"read"`

		NetBlockRemove func(p0 context.Context, p1 NetBlockList) error `perm:"admin"`

		NetConnect func(p0 context.Context, p1 peer.AddrInfo) error `perm:"write"`

		NetConnectedness func(p0 context.Context, p1 peer.ID) (network.Connectedness, error) `perm:"read"`

		NetDisconnect func(p0 context.Context, p1 peer.ID) error `perm:"write"`

		NetFindPeer func(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) `perm:"read"`

		NetPeerInfo func(p0 context.Context, p1 peer.ID) (*ExtendedPeerInfo, error) `perm:"read"`

		NetPeers func(p0 context.Context) ([]peer.AddrInfo, error) `perm:"read"`
	}
}

type NetStub struct {
}

type WalletStruct struct {
	Internal struct {
		WalletDelete func(p0 context.Context, p1 address.Address) error `perm:"admin"`

		WalletExport func(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) `perm:"admin"`

		WalletHas func(p0 context.Context, p1 address.Address) (bool, error) `perm:"admin"`

		WalletImport func(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) `perm:"admin"`

		WalletList func(p0 context.Context) ([]address.Address, error) `perm:"admin"`

		WalletNew func(p0 context.Context, p1 types.KeyType) (address.Address, error) `perm:"admin"`

		WalletSign func(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) `perm:"admin"`
	}
}

type WalletStub struct {
}

func (s *BoostStruct) ActorSectorSize(p0 context.Context, p1 address.Address) (abi.SectorSize, error) {
	if s.Internal.ActorSectorSize == nil {
		return *new(abi.SectorSize), ErrNotSupported
	}
	return s.Internal.ActorSectorSize(p0, p1)
}

func (s *BoostStub) ActorSectorSize(p0 context.Context, p1 address.Address) (abi.SectorSize, error) {
	return *new(abi.SectorSize), ErrNotSupported
}

func (s *BoostStruct) BoostDagstoreGC(p0 context.Context) ([]DagstoreShardResult, error) {
	if s.Internal.BoostDagstoreGC == nil {
		return *new([]DagstoreShardResult), ErrNotSupported
	}
	return s.Internal.BoostDagstoreGC(p0)
}

func (s *BoostStub) BoostDagstoreGC(p0 context.Context) ([]DagstoreShardResult, error) {
	return *new([]DagstoreShardResult), ErrNotSupported
}

func (s *BoostStruct) BoostDagstoreInitializeAll(p0 context.Context, p1 DagstoreInitializeAllParams) (<-chan DagstoreInitializeAllEvent, error) {
	if s.Internal.BoostDagstoreInitializeAll == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostDagstoreInitializeAll(p0, p1)
}

func (s *BoostStub) BoostDagstoreInitializeAll(p0 context.Context, p1 DagstoreInitializeAllParams) (<-chan DagstoreInitializeAllEvent, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) BoostDagstoreInitializeShard(p0 context.Context, p1 string) error {
	if s.Internal.BoostDagstoreInitializeShard == nil {
		return ErrNotSupported
	}
	return s.Internal.BoostDagstoreInitializeShard(p0, p1)
}

func (s *BoostStub) BoostDagstoreInitializeShard(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *BoostStruct) BoostDagstoreListShards(p0 context.Context) ([]DagstoreShardInfo, error) {
	if s.Internal.BoostDagstoreListShards == nil {
		return *new([]DagstoreShardInfo), ErrNotSupported
	}
	return s.Internal.BoostDagstoreListShards(p0)
}

func (s *BoostStub) BoostDagstoreListShards(p0 context.Context) ([]DagstoreShardInfo, error) {
	return *new([]DagstoreShardInfo), ErrNotSupported
}

func (s *BoostStruct) BoostDagstoreRecoverShard(p0 context.Context, p1 string) error {
	if s.Internal.BoostDagstoreRecoverShard == nil {
		return ErrNotSupported
	}
	return s.Internal.BoostDagstoreRecoverShard(p0, p1)
}

func (s *BoostStub) BoostDagstoreRecoverShard(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *BoostStruct) BoostDagstoreRegisterShard(p0 context.Context, p1 string) error {
	if s.Internal.BoostDagstoreRegisterShard == nil {
		return ErrNotSupported
	}
	return s.Internal.BoostDagstoreRegisterShard(p0, p1)
}

func (s *BoostStub) BoostDagstoreRegisterShard(p0 context.Context, p1 string) error {
	return ErrNotSupported
}

func (s *BoostStruct) BoostDeal(p0 context.Context, p1 uuid.UUID) (*smtypes.ProviderDealState, error) {
	if s.Internal.BoostDeal == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostDeal(p0, p1)
}

func (s *BoostStub) BoostDeal(p0 context.Context, p1 uuid.UUID) (*smtypes.ProviderDealState, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) BoostDummyDeal(p0 context.Context, p1 smtypes.DealParams) (*ProviderDealRejectionInfo, error) {
	if s.Internal.BoostDummyDeal == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostDummyDeal(p0, p1)
}

func (s *BoostStub) BoostDummyDeal(p0 context.Context, p1 smtypes.DealParams) (*ProviderDealRejectionInfo, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) BoostIndexerAnnounceAllDeals(p0 context.Context) error {
	if s.Internal.BoostIndexerAnnounceAllDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.BoostIndexerAnnounceAllDeals(p0)
}

func (s *BoostStub) BoostIndexerAnnounceAllDeals(p0 context.Context) error {
	return ErrNotSupported
}

func (s *BoostStruct) BoostOfflineDealWithData(p0 context.Context, p1 uuid.UUID, p2 string) (*ProviderDealRejectionInfo, error) {
	if s.Internal.BoostOfflineDealWithData == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.BoostOfflineDealWithData(p0, p1, p2)
}

func (s *BoostStub) BoostOfflineDealWithData(p0 context.Context, p1 uuid.UUID, p2 string) (*ProviderDealRejectionInfo, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) DealsConsiderOfflineRetrievalDeals(p0 context.Context) (bool, error) {
	if s.Internal.DealsConsiderOfflineRetrievalDeals == nil {
		return false, ErrNotSupported
	}
	return s.Internal.DealsConsiderOfflineRetrievalDeals(p0)
}

func (s *BoostStub) DealsConsiderOfflineRetrievalDeals(p0 context.Context) (bool, error) {
	return false, ErrNotSupported
}

func (s *BoostStruct) DealsConsiderOfflineStorageDeals(p0 context.Context) (bool, error) {
	if s.Internal.DealsConsiderOfflineStorageDeals == nil {
		return false, ErrNotSupported
	}
	return s.Internal.DealsConsiderOfflineStorageDeals(p0)
}

func (s *BoostStub) DealsConsiderOfflineStorageDeals(p0 context.Context) (bool, error) {
	return false, ErrNotSupported
}

func (s *BoostStruct) DealsConsiderOnlineRetrievalDeals(p0 context.Context) (bool, error) {
	if s.Internal.DealsConsiderOnlineRetrievalDeals == nil {
		return false, ErrNotSupported
	}
	return s.Internal.DealsConsiderOnlineRetrievalDeals(p0)
}

func (s *BoostStub) DealsConsiderOnlineRetrievalDeals(p0 context.Context) (bool, error) {
	return false, ErrNotSupported
}

func (s *BoostStruct) DealsConsiderOnlineStorageDeals(p0 context.Context) (bool, error) {
	if s.Internal.DealsConsiderOnlineStorageDeals == nil {
		return false, ErrNotSupported
	}
	return s.Internal.DealsConsiderOnlineStorageDeals(p0)
}

func (s *BoostStub) DealsConsiderOnlineStorageDeals(p0 context.Context) (bool, error) {
	return false, ErrNotSupported
}

func (s *BoostStruct) DealsConsiderUnverifiedStorageDeals(p0 context.Context) (bool, error) {
	if s.Internal.DealsConsiderUnverifiedStorageDeals == nil {
		return false, ErrNotSupported
	}
	return s.Internal.DealsConsiderUnverifiedStorageDeals(p0)
}

func (s *BoostStub) DealsConsiderUnverifiedStorageDeals(p0 context.Context) (bool, error) {
	return false, ErrNotSupported
}

func (s *BoostStruct) DealsConsiderVerifiedStorageDeals(p0 context.Context) (bool, error) {
	if s.Internal.DealsConsiderVerifiedStorageDeals == nil {
		return false, ErrNotSupported
	}
	return s.Internal.DealsConsiderVerifiedStorageDeals(p0)
}

func (s *BoostStub) DealsConsiderVerifiedStorageDeals(p0 context.Context) (bool, error) {
	return false, ErrNotSupported
}

func (s *BoostStruct) DealsPieceCidBlocklist(p0 context.Context) ([]cid.Cid, error) {
	if s.Internal.DealsPieceCidBlocklist == nil {
		return *new([]cid.Cid), ErrNotSupported
	}
	return s.Internal.DealsPieceCidBlocklist(p0)
}

func (s *BoostStub) DealsPieceCidBlocklist(p0 context.Context) ([]cid.Cid, error) {
	return *new([]cid.Cid), ErrNotSupported
}

func (s *BoostStruct) DealsSetConsiderOfflineRetrievalDeals(p0 context.Context, p1 bool) error {
	if s.Internal.DealsSetConsiderOfflineRetrievalDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSetConsiderOfflineRetrievalDeals(p0, p1)
}

func (s *BoostStub) DealsSetConsiderOfflineRetrievalDeals(p0 context.Context, p1 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) DealsSetConsiderOfflineStorageDeals(p0 context.Context, p1 bool) error {
	if s.Internal.DealsSetConsiderOfflineStorageDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSetConsiderOfflineStorageDeals(p0, p1)
}

func (s *BoostStub) DealsSetConsiderOfflineStorageDeals(p0 context.Context, p1 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) DealsSetConsiderOnlineRetrievalDeals(p0 context.Context, p1 bool) error {
	if s.Internal.DealsSetConsiderOnlineRetrievalDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSetConsiderOnlineRetrievalDeals(p0, p1)
}

func (s *BoostStub) DealsSetConsiderOnlineRetrievalDeals(p0 context.Context, p1 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) DealsSetConsiderOnlineStorageDeals(p0 context.Context, p1 bool) error {
	if s.Internal.DealsSetConsiderOnlineStorageDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSetConsiderOnlineStorageDeals(p0, p1)
}

func (s *BoostStub) DealsSetConsiderOnlineStorageDeals(p0 context.Context, p1 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) DealsSetConsiderUnverifiedStorageDeals(p0 context.Context, p1 bool) error {
	if s.Internal.DealsSetConsiderUnverifiedStorageDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSetConsiderUnverifiedStorageDeals(p0, p1)
}

func (s *BoostStub) DealsSetConsiderUnverifiedStorageDeals(p0 context.Context, p1 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) DealsSetConsiderVerifiedStorageDeals(p0 context.Context, p1 bool) error {
	if s.Internal.DealsSetConsiderVerifiedStorageDeals == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSetConsiderVerifiedStorageDeals(p0, p1)
}

func (s *BoostStub) DealsSetConsiderVerifiedStorageDeals(p0 context.Context, p1 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) DealsSetPieceCidBlocklist(p0 context.Context, p1 []cid.Cid) error {
	if s.Internal.DealsSetPieceCidBlocklist == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSetPieceCidBlocklist(p0, p1)
}

func (s *BoostStub) DealsSetPieceCidBlocklist(p0 context.Context, p1 []cid.Cid) error {
	return ErrNotSupported
}

func (s *BoostStruct) MarketCancelDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	if s.Internal.MarketCancelDataTransfer == nil {
		return ErrNotSupported
	}
	return s.Internal.MarketCancelDataTransfer(p0, p1, p2, p3)
}

func (s *BoostStub) MarketCancelDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) MarketDataTransferUpdates(p0 context.Context) (<-chan lapi.DataTransferChannel, error) {
	if s.Internal.MarketDataTransferUpdates == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.MarketDataTransferUpdates(p0)
}

func (s *BoostStub) MarketDataTransferUpdates(p0 context.Context) (<-chan lapi.DataTransferChannel, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) MarketGetAsk(p0 context.Context) (*storagemarket.SignedStorageAsk, error) {
	if s.Internal.MarketGetAsk == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.MarketGetAsk(p0)
}

func (s *BoostStub) MarketGetAsk(p0 context.Context) (*storagemarket.SignedStorageAsk, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) MarketGetRetrievalAsk(p0 context.Context) (*retrievalmarket.Ask, error) {
	if s.Internal.MarketGetRetrievalAsk == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.MarketGetRetrievalAsk(p0)
}

func (s *BoostStub) MarketGetRetrievalAsk(p0 context.Context) (*retrievalmarket.Ask, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) MarketImportDealData(p0 context.Context, p1 cid.Cid, p2 string) error {
	if s.Internal.MarketImportDealData == nil {
		return ErrNotSupported
	}
	return s.Internal.MarketImportDealData(p0, p1, p2)
}

func (s *BoostStub) MarketImportDealData(p0 context.Context, p1 cid.Cid, p2 string) error {
	return ErrNotSupported
}

func (s *BoostStruct) MarketListDataTransfers(p0 context.Context) ([]lapi.DataTransferChannel, error) {
	if s.Internal.MarketListDataTransfers == nil {
		return *new([]lapi.DataTransferChannel), ErrNotSupported
	}
	return s.Internal.MarketListDataTransfers(p0)
}

func (s *BoostStub) MarketListDataTransfers(p0 context.Context) ([]lapi.DataTransferChannel, error) {
	return *new([]lapi.DataTransferChannel), ErrNotSupported
}

func (s *BoostStruct) MarketListIncompleteDeals(p0 context.Context) ([]storagemarket.MinerDeal, error) {
	if s.Internal.MarketListIncompleteDeals == nil {
		return *new([]storagemarket.MinerDeal), ErrNotSupported
	}
	return s.Internal.MarketListIncompleteDeals(p0)
}

func (s *BoostStub) MarketListIncompleteDeals(p0 context.Context) ([]storagemarket.MinerDeal, error) {
	return *new([]storagemarket.MinerDeal), ErrNotSupported
}

func (s *BoostStruct) MarketListRetrievalDeals(p0 context.Context) ([]retrievalmarket.ProviderDealState, error) {
	if s.Internal.MarketListRetrievalDeals == nil {
		return *new([]retrievalmarket.ProviderDealState), ErrNotSupported
	}
	return s.Internal.MarketListRetrievalDeals(p0)
}

func (s *BoostStub) MarketListRetrievalDeals(p0 context.Context) ([]retrievalmarket.ProviderDealState, error) {
	return *new([]retrievalmarket.ProviderDealState), ErrNotSupported
}

func (s *BoostStruct) MarketRestartDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	if s.Internal.MarketRestartDataTransfer == nil {
		return ErrNotSupported
	}
	return s.Internal.MarketRestartDataTransfer(p0, p1, p2, p3)
}

func (s *BoostStub) MarketRestartDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return ErrNotSupported
}

func (s *BoostStruct) MarketSetAsk(p0 context.Context, p1 types.BigInt, p2 types.BigInt, p3 abi.ChainEpoch, p4 abi.PaddedPieceSize, p5 abi.PaddedPieceSize) error {
	if s.Internal.MarketSetAsk == nil {
		return ErrNotSupported
	}
	return s.Internal.MarketSetAsk(p0, p1, p2, p3, p4, p5)
}

func (s *BoostStub) MarketSetAsk(p0 context.Context, p1 types.BigInt, p2 types.BigInt, p3 abi.ChainEpoch, p4 abi.PaddedPieceSize, p5 abi.PaddedPieceSize) error {
	return ErrNotSupported
}

func (s *BoostStruct) MarketSetRetrievalAsk(p0 context.Context, p1 *retrievalmarket.Ask) error {
	if s.Internal.MarketSetRetrievalAsk == nil {
		return ErrNotSupported
	}
	return s.Internal.MarketSetRetrievalAsk(p0, p1)
}

func (s *BoostStub) MarketSetRetrievalAsk(p0 context.Context, p1 *retrievalmarket.Ask) error {
	return ErrNotSupported
}

func (s *BoostStruct) PiecesGetCIDInfo(p0 context.Context, p1 cid.Cid) (*piecestore.CIDInfo, error) {
	if s.Internal.PiecesGetCIDInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.PiecesGetCIDInfo(p0, p1)
}

func (s *BoostStub) PiecesGetCIDInfo(p0 context.Context, p1 cid.Cid) (*piecestore.CIDInfo, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) PiecesGetPieceInfo(p0 context.Context, p1 cid.Cid) (*piecestore.PieceInfo, error) {
	if s.Internal.PiecesGetPieceInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.PiecesGetPieceInfo(p0, p1)
}

func (s *BoostStub) PiecesGetPieceInfo(p0 context.Context, p1 cid.Cid) (*piecestore.PieceInfo, error) {
	return nil, ErrNotSupported
}

func (s *BoostStruct) PiecesListCidInfos(p0 context.Context) ([]cid.Cid, error) {
	if s.Internal.PiecesListCidInfos == nil {
		return *new([]cid.Cid), ErrNotSupported
	}
	return s.Internal.PiecesListCidInfos(p0)
}

func (s *BoostStub) PiecesListCidInfos(p0 context.Context) ([]cid.Cid, error) {
	return *new([]cid.Cid), ErrNotSupported
}

func (s *BoostStruct) PiecesListPieces(p0 context.Context) ([]cid.Cid, error) {
	if s.Internal.PiecesListPieces == nil {
		return *new([]cid.Cid), ErrNotSupported
	}
	return s.Internal.PiecesListPieces(p0)
}

func (s *BoostStub) PiecesListPieces(p0 context.Context) ([]cid.Cid, error) {
	return *new([]cid.Cid), ErrNotSupported
}

func (s *BoostStruct) RuntimeSubsystems(p0 context.Context) (lapi.MinerSubsystems, error) {
	if s.Internal.RuntimeSubsystems == nil {
		return *new(lapi.MinerSubsystems), ErrNotSupported
	}
	return s.Internal.RuntimeSubsystems(p0)
}

func (s *BoostStub) RuntimeSubsystems(p0 context.Context) (lapi.MinerSubsystems, error) {
	return *new(lapi.MinerSubsystems), ErrNotSupported
}

func (s *ChainIOStruct) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	if s.Internal.ChainHasObj == nil {
		return false, ErrNotSupported
	}
	return s.Internal.ChainHasObj(p0, p1)
}

func (s *ChainIOStub) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	return false, ErrNotSupported
}

func (s *ChainIOStruct) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	if s.Internal.ChainReadObj == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.ChainReadObj(p0, p1)
}

func (s *ChainIOStub) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *CommonStruct) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	if s.Internal.AuthNew == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.AuthNew(p0, p1)
}

func (s *CommonStub) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *CommonStruct) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	if s.Internal.AuthVerify == nil {
		return *new([]auth.Permission), ErrNotSupported
	}
	return s.Internal.AuthVerify(p0, p1)
}

func (s *CommonStub) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	return *new([]auth.Permission), ErrNotSupported
}

func (s *CommonStruct) LogList(p0 context.Context) ([]string, error) {
	if s.Internal.LogList == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.LogList(p0)
}

func (s *CommonStub) LogList(p0 context.Context) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *CommonStruct) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.LogSetLevel == nil {
		return ErrNotSupported
	}
	return s.Internal.LogSetLevel(p0, p1, p2)
}

func (s *CommonStub) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *CommonStruct) Shutdown(p0 context.Context) error {
	if s.Internal.Shutdown == nil {
		return ErrNotSupported
	}
	return s.Internal.Shutdown(p0)
}

func (s *CommonStub) Shutdown(p0 context.Context) error {
	return ErrNotSupported
}

func (s *NetStruct) ID(p0 context.Context) (peer.ID, error) {
	if s.Internal.ID == nil {
		return *new(peer.ID), ErrNotSupported
	}
	return s.Internal.ID(p0)
}

func (s *NetStub) ID(p0 context.Context) (peer.ID, error) {
	return *new(peer.ID), ErrNotSupported
}

func (s *NetStruct) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	if s.Internal.NetAddrsListen == nil {
		return *new(peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetAddrsListen(p0)
}

func (s *NetStub) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	return *new(peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetAgentVersion(p0 context.Context, p1 peer.ID) (string, error) {
	if s.Internal.NetAgentVersion == nil {
		return "", ErrNotSupported
	}
	return s.Internal.NetAgentVersion(p0, p1)
}

func (s *NetStub) NetAgentVersion(p0 context.Context, p1 peer.ID) (string, error) {
	return "", ErrNotSupported
}

func (s *NetStruct) NetAutoNatStatus(p0 context.Context) (NatInfo, error) {
	if s.Internal.NetAutoNatStatus == nil {
		return *new(NatInfo), ErrNotSupported
	}
	return s.Internal.NetAutoNatStatus(p0)
}

func (s *NetStub) NetAutoNatStatus(p0 context.Context) (NatInfo, error) {
	return *new(NatInfo), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStats(p0 context.Context) (metrics.Stats, error) {
	if s.Internal.NetBandwidthStats == nil {
		return *new(metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStats(p0)
}

func (s *NetStub) NetBandwidthStats(p0 context.Context) (metrics.Stats, error) {
	return *new(metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStatsByPeer(p0 context.Context) (map[string]metrics.Stats, error) {
	if s.Internal.NetBandwidthStatsByPeer == nil {
		return *new(map[string]metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStatsByPeer(p0)
}

func (s *NetStub) NetBandwidthStatsByPeer(p0 context.Context) (map[string]metrics.Stats, error) {
	return *new(map[string]metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStatsByProtocol(p0 context.Context) (map[protocol.ID]metrics.Stats, error) {
	if s.Internal.NetBandwidthStatsByProtocol == nil {
		return *new(map[protocol.ID]metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStatsByProtocol(p0)
}

func (s *NetStub) NetBandwidthStatsByProtocol(p0 context.Context) (map[protocol.ID]metrics.Stats, error) {
	return *new(map[protocol.ID]metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBlockAdd(p0 context.Context, p1 NetBlockList) error {
	if s.Internal.NetBlockAdd == nil {
		return ErrNotSupported
	}
	return s.Internal.NetBlockAdd(p0, p1)
}

func (s *NetStub) NetBlockAdd(p0 context.Context, p1 NetBlockList) error {
	return ErrNotSupported
}

func (s *NetStruct) NetBlockList(p0 context.Context) (NetBlockList, error) {
	if s.Internal.NetBlockList == nil {
		return *new(NetBlockList), ErrNotSupported
	}
	return s.Internal.NetBlockList(p0)
}

func (s *NetStub) NetBlockList(p0 context.Context) (NetBlockList, error) {
	return *new(NetBlockList), ErrNotSupported
}

func (s *NetStruct) NetBlockRemove(p0 context.Context, p1 NetBlockList) error {
	if s.Internal.NetBlockRemove == nil {
		return ErrNotSupported
	}
	return s.Internal.NetBlockRemove(p0, p1)
}

func (s *NetStub) NetBlockRemove(p0 context.Context, p1 NetBlockList) error {
	return ErrNotSupported
}

func (s *NetStruct) NetConnect(p0 context.Context, p1 peer.AddrInfo) error {
	if s.Internal.NetConnect == nil {
		return ErrNotSupported
	}
	return s.Internal.NetConnect(p0, p1)
}

func (s *NetStub) NetConnect(p0 context.Context, p1 peer.AddrInfo) error {
	return ErrNotSupported
}

func (s *NetStruct) NetConnectedness(p0 context.Context, p1 peer.ID) (network.Connectedness, error) {
	if s.Internal.NetConnectedness == nil {
		return *new(network.Connectedness), ErrNotSupported
	}
	return s.Internal.NetConnectedness(p0, p1)
}

func (s *NetStub) NetConnectedness(p0 context.Context, p1 peer.ID) (network.Connectedness, error) {
	return *new(network.Connectedness), ErrNotSupported
}

func (s *NetStruct) NetDisconnect(p0 context.Context, p1 peer.ID) error {
	if s.Internal.NetDisconnect == nil {
		return ErrNotSupported
	}
	return s.Internal.NetDisconnect(p0, p1)
}

func (s *NetStub) NetDisconnect(p0 context.Context, p1 peer.ID) error {
	return ErrNotSupported
}

func (s *NetStruct) NetFindPeer(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) {
	if s.Internal.NetFindPeer == nil {
		return *new(peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetFindPeer(p0, p1)
}

func (s *NetStub) NetFindPeer(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) {
	return *new(peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetPeerInfo(p0 context.Context, p1 peer.ID) (*ExtendedPeerInfo, error) {
	if s.Internal.NetPeerInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.NetPeerInfo(p0, p1)
}

func (s *NetStub) NetPeerInfo(p0 context.Context, p1 peer.ID) (*ExtendedPeerInfo, error) {
	return nil, ErrNotSupported
}

func (s *NetStruct) NetPeers(p0 context.Context) ([]peer.AddrInfo, error) {
	if s.Internal.NetPeers == nil {
		return *new([]peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetPeers(p0)
}

func (s *NetStub) NetPeers(p0 context.Context) ([]peer.AddrInfo, error) {
	return *new([]peer.AddrInfo), ErrNotSupported
}

func (s *WalletStruct) WalletDelete(p0 context.Context, p1 address.Address) error {
	if s.Internal.WalletDelete == nil {
		return ErrNotSupported
	}
	return s.Internal.WalletDelete(p0, p1)
}

func (s *WalletStub) WalletDelete(p0 context.Context, p1 address.Address) error {
	return ErrNotSupported
}

func (s *WalletStruct) WalletExport(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) {
	if s.Internal.WalletExport == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.WalletExport(p0, p1)
}

func (s *WalletStub) WalletExport(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) {
	return nil, ErrNotSupported
}

func (s *WalletStruct) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	if s.Internal.WalletHas == nil {
		return false, ErrNotSupported
	}
	return s.Internal.WalletHas(p0, p1)
}

func (s *WalletStub) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	return false, ErrNotSupported
}

func (s *WalletStruct) WalletImport(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) {
	if s.Internal.WalletImport == nil {
		return *new(address.Address), ErrNotSupported
	}
	return s.Internal.WalletImport(p0, p1)
}

func (s *WalletStub) WalletImport(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) {
	return *new(address.Address), ErrNotSupported
}

func (s *WalletStruct) WalletList(p0 context.Context) ([]address.Address, error) {
	if s.Internal.WalletList == nil {
		return *new([]address.Address), ErrNotSupported
	}
	return s.Internal.WalletList(p0)
}

func (s *WalletStub) WalletList(p0 context.Context) ([]address.Address, error) {
	return *new([]address.Address), ErrNotSupported
}

func (s *WalletStruct) WalletNew(p0 context.Context, p1 types.KeyType) (address.Address, error) {
	if s.Internal.WalletNew == nil {
		return *new(address.Address), ErrNotSupported
	}
	return s.Internal.WalletNew(p0, p1)
}

func (s *WalletStub) WalletNew(p0 context.Context, p1 types.KeyType) (address.Address, error) {
	return *new(address.Address), ErrNotSupported
}

func (s *WalletStruct) WalletSign(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) {
	if s.Internal.WalletSign == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.WalletSign(p0, p1, p2)
}

func (s *WalletStub) WalletSign(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) {
	return nil, ErrNotSupported
}

var _ Boost = new(BoostStruct)
var _ ChainIO = new(ChainIOStruct)
var _ Common = new(CommonStruct)
var _ CommonNet = new(CommonNetStruct)
var _ Net = new(NetStruct)
var _ Wallet = new(WalletStruct)
