package grpcserver

import (
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"github.com/spacemeshos/go-spacemesh/api"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/log"
	"github.com/spacemeshos/go-spacemesh/p2p/peers"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GlobalStateService exposes global state data, output from the STF
type GlobalStateService struct {
	Network     api.NetworkAPI // P2P Swarm
	Mesh        api.TxAPI      // Mesh
	GenTime     api.GenesisTimeAPI
	PeerCounter api.PeerCounter
	Syncer      api.Syncer
}

// RegisterService registers this service with a grpc server instance
func (s GlobalStateService) RegisterService(server *Server) {
	pb.RegisterGlobalStateServiceServer(server.GrpcServer, s)
}

// NewGlobalStateService creates a new grpc service using config data.
func NewGlobalStateService(
	net api.NetworkAPI, tx api.TxAPI, genTime api.GenesisTimeAPI,
	syncer api.Syncer) *GlobalStateService {
	return &GlobalStateService{
		Network:     net,
		Mesh:        tx,
		GenTime:     genTime,
		PeerCounter: peers.NewPeers(net, log.NewDefault("grpc_server.GlobalStateService")),
		Syncer:      syncer,
	}
}

// GlobalStateHash returns the latest layer and its computed global state hash
func (s GlobalStateService) GlobalStateHash(ctx context.Context, in *pb.GlobalStateHashRequest) (*pb.GlobalStateHashResponse, error) {
	log.Info("GRPC GlobalStateService.GlobalStateHash")
	return &pb.GlobalStateHashResponse{Response: &pb.GlobalStateHash{
		RootHash:    s.Mesh.GetStateRoot().Bytes(),
		LayerNumber: s.Mesh.LatestLayerInState().Uint64(),
	}}, nil
}

// Account returns data for one account
func (s GlobalStateService) Account(ctx context.Context, in *pb.AccountRequest) (*pb.AccountResponse, error) {
	log.Info("GRPC GlobalStateService.Account")

	// Load data
	addr := types.BytesToAddress(in.AccountId.Address)
	balance := s.Mesh.GetBalance(addr)
	counter := s.Mesh.GetNonce(addr)

	return &pb.AccountResponse{Account: &pb.Account{
		Address: &pb.AccountId{Address: addr.Bytes()},
		Counter: counter,
		Balance: &pb.Amount{Value: balance},
	}}, nil
}

// AccountDataQuery returns historical account data such as rewards and receipts
func (s GlobalStateService) AccountDataQuery(ctx context.Context, in *pb.AccountDataQueryRequest) (*pb.AccountDataQueryResponse, error) {
	log.Info("GRPC GlobalStateService.AccountDataQuery")

	if in.Filter == nil {
		return nil, status.Errorf(codes.InvalidArgument, "`Filter` must be provided")
	}
	if in.Filter.AccountId == nil {
		return nil, status.Errorf(codes.InvalidArgument, "`Filter.AccountId` must be provided")
	}
	if in.Filter.AccountDataFlags == uint32(pb.AccountDataFlag_ACCOUNT_DATA_FLAG_UNSPECIFIED) {
		return nil, status.Errorf(codes.InvalidArgument, "`Filter.AccountMeshDataFlags` must set at least one bitfield")
	}

	// Read the filter flags
	filterTxReceipt := in.Filter.AccountDataFlags&uint32(pb.AccountDataFlag_ACCOUNT_DATA_FLAG_TRANSACTION_RECEIPT) != 0
	filterReward := in.Filter.AccountDataFlags&uint32(pb.AccountDataFlag_ACCOUNT_DATA_FLAG_REWARD) != 0
	filterAccount := in.Filter.AccountDataFlags&uint32(pb.AccountDataFlag_ACCOUNT_DATA_FLAG_ACCOUNT) != 0

	addr := types.BytesToAddress(in.Filter.AccountId.Address)
	res := &pb.AccountDataQueryResponse{}

	if filterTxReceipt {
		// TODO: Implement this. The node does not implement tx receipts yet.
	}

	if filterReward {
		dbRewards, err := s.Mesh.GetRewards(addr)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error getting rewards data")
		}
		for _, r := range dbRewards {
			res.AccountItem = append(res.AccountItem, &pb.AccountData{Item: &pb.AccountData_Reward{
				Reward: &pb.Reward{
					Layer:       r.Layer.Uint64(),
					Total:       &pb.Amount{Value: r.TotalReward},
					LayerReward: &pb.Amount{Value: r.LayerRewardEstimate},
					// Leave this out for now as this is changing
					//LayerComputed: 0,
					Coinbase: &pb.AccountId{Address: addr.Bytes()},
					Smesher:  nil,
				},
			}})
		}

	}

	if filterAccount {

	}

	// Adjust for max results, offset

	return nil, nil
}

// SmesherDataQuery returns historical info on smesher rewards
func (s GlobalStateService) SmesherDataQuery(ctx context.Context, in *pb.SmesherDataQueryRequest) (*pb.SmesherDataQueryResponse, error) {
	log.Info("GRPC GlobalStateService.SmesherDataQuery")
	return nil, nil
}

// STREAMS

// AccountDataStream exposes a stream of account-related data
func (s GlobalStateService) AccountDataStream(request *pb.AccountDataStreamRequest, stream pb.GlobalStateService_AccountDataStreamServer) error {
	log.Info("GRPC GlobalStateService.AccountDataStream")
	return nil
}

// SmesherRewardStream exposes a stream of smesher rewards
func (s GlobalStateService) SmesherRewardStream(request *pb.SmesherRewardStreamRequest, stream pb.GlobalStateService_SmesherRewardStreamServer) error {
	log.Info("GRPC GlobalStateService.SmesherRewardStream")
	return nil
}

// AppEventStream exposes a stream of emitted app events
func (s GlobalStateService) AppEventStream(request *pb.AppEventStreamRequest, stream pb.GlobalStateService_AppEventStreamServer) error {
	log.Info("GRPC GlobalStateService.AppEventStream")
	return nil
}

// GlobalStateStream exposes a stream of global data data items: rewards, receipts, account info, global state hash
func (s GlobalStateService) GlobalStateStream(request *pb.GlobalStateStreamRequest, stream pb.GlobalStateService_GlobalStateStreamServer) error {
	log.Info("GRPC GlobalStateService.GlobalStateStream")
	return nil
}