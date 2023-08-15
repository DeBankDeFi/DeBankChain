package p2p

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/p2p/gating"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"
	ds "github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core"
	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/metrics"
	cmgr "github.com/libp2p/go-libp2p/p2p/net/connmgr"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
)

var DefaultBootnodes = []*enode.Node{
	// Node ID: 5845e1395d7e120912ba2588a8f0451f538caf98d69b9af3602d6f3ab8cc0759
	// URLv4:   enode://e8154cad86b66c3b1cb9b2c875c263f0cdff8f98ddbfe89fc9516855d141e95593fa3cc5708da8d6c979ecc7a0a919fbc2eaefb7fa8f4198d9753d82151156c8@43.207.77.108:9222?discport=51787
	// Record has sequence number 1691653398385 and 6 key/value pairs.
	//   "id"        "v4"
	//   "ip"        43.207.77.108
	//   "opstack"   8496b07b00
	//   "secp256k1" a102e8154cad86b66c3b1cb9b2c875c263f0cdff8f98ddbfe89fc9516855d141e955
	//   "tcp"       9222
	//   "udp"       51787
	//enode.MustParse("enr:-J64QA8c432-TY4uxlxQ4VXQOKWqPlh2g3fGZd_iTAE7PqmWaitTfZOyByHmux87_fqFlZhKtZ_sm-TPYxgIMaaGRBSGAYneZl9xgmlkgnY0gmlwhCvPTWyHb3BzdGFja4SWsHsAiXNlY3AyNTZrMaEC6BVMrYa2bDscubLIdcJj8M3_j5jdv-ifyVFoVdFB6VWDdGNwgiQGg3VkcILKSw"),

	// Node ID: 0decb516278633804bfded1e1e5b5f684bbfa77f6fbd13cacfea6216db43c63a
	// URLv4:   enode://43b5fa8dde7b216034ad97c2067ac21d4be62f3860b4d3956abd5baffecf48cedd36791c08e7697c2efec8c2012ce5ca8e919ca7d33c06652924737a0e8cc988@43.207.221.181:9222?discport=49268
	// Record has sequence number 1691660727153 and 6 key/value pairs.
	//   "id"        "v4"
	//   "ip"        43.207.221.181
	//   "opstack"   8496b07b00
	//   "secp256k1" a10243b5fa8dde7b216034ad97c2067ac21d4be62f3860b4d3956abd5baffecf48ce
	//   "tcp"       9222
	//   "udp"       49268
	//enode.MustParse("enr:-J64QFI65ajctEyDMwkqoNt9tNAsyDFauQA45ugqOz9POiKrU--o2sHNpy3acA6muag_K0fEGRkmGTdkFwwMeMyf6_GGAYne1jNxgmlkgnY0gmlwhCvP3bWHb3BzdGFja4SWsHsAiXNlY3AyNTZrMaECQ7X6jd57IWA0rZfCBnrCHUvmLzhgtNOVar1br_7PSM6DdGNwgiQGg3VkcILAdA"),

	enode.MustParse("enode://43b5fa8dde7b216034ad97c2067ac21d4be62f3860b4d3956abd5baffecf48cedd36791c08e7697c2efec8c2012ce5ca8e919ca7d33c06652924737a0e8cc988@43.207.221.181:9222?discport=49268"),
	enode.MustParse("enode://43b5fa8dde7b216034ad97c2067ac21d4be62f3860b4d3956abd5baffecf48cedd36791c08e7697c2efec8c2012ce5ca8e919ca7d33c06652924737a0e8cc988@43.207.221.181:9222?discport=49268"),
}

type HostMetrics interface {
	gating.UnbanMetrics
	gating.ConnectionGaterMetrics
}

// SetupP2P provides a host and discovery service for usage in the rollup node.
type SetupP2P interface {
	Check() error
	Disabled() bool
	// Host creates a libp2p host service. Returns nil, nil if p2p is disabled.
	Host(log log.Logger, reporter metrics.Reporter, metrics HostMetrics) (host.Host, error)
	// Discovery creates a disc-v5 service. Returns nil, nil, nil if discovery is disabled.
	Discovery(log log.Logger, rollupCfg *rollup.Config, tcpPort uint16) (*enode.LocalNode, *discover.UDPv5, error)
	TargetPeers() uint
	BanPeers() bool
	BanThreshold() float64
	BanDuration() time.Duration
	GossipSetupConfigurables
	ReqRespSyncEnabled() bool
}

// ScoringParams defines the various types of peer scoring parameters.
type ScoringParams struct {
	PeerScoring        pubsub.PeerScoreParams
	ApplicationScoring ApplicationScoreParams
}

// Config sets up a p2p host and discv5 service from configuration.
// This implements SetupP2P.
type Config struct {
	Priv *crypto.Secp256k1PrivateKey

	DisableP2P  bool
	NoDiscovery bool

	// Enable P2P-based alt-syncing method (req-resp protocol, not gossip)
	AltSync bool

	ScoringParams *ScoringParams

	// Whether to ban peers based on their [PeerScoring] score. Should be negative.
	BanningEnabled bool
	// Minimum score before peers are disconnected and banned
	BanningThreshold float64
	BanningDuration  time.Duration

	ListenIP      net.IP
	ListenTCPPort uint16

	// Port to bind discv5 to
	ListenUDPPort uint16

	AdvertiseIP      net.IP
	AdvertiseTCPPort uint16
	AdvertiseUDPPort uint16
	Bootnodes        []*enode.Node
	DiscoveryDB      *enode.DB

	StaticPeers []core.Multiaddr

	HostMux             []libp2p.Option
	HostSecurity        []libp2p.Option
	NoTransportSecurity bool

	PeersLo    uint
	PeersHi    uint
	PeersGrace time.Duration

	MeshD     int // topic stable mesh target count
	MeshDLo   int // topic stable mesh low watermark
	MeshDHi   int // topic stable mesh high watermark
	MeshDLazy int // gossip target

	// FloodPublish publishes messages from ourselves to peers outside of the gossip topic mesh but supporting the same topic.
	FloodPublish bool

	// If true a NAT manager will host a NAT port mapping that is updated with PMP and UPNP by libp2p/go-nat
	NAT bool

	UserAgent string

	TimeoutNegotiation time.Duration
	TimeoutAccept      time.Duration
	TimeoutDial        time.Duration

	// Underlying store that hosts connection-gater and peerstore data.
	Store ds.Batching

	EnableReqRespSync bool
}

func DefaultConnManager(conf *Config) (connmgr.ConnManager, error) {
	return cmgr.NewConnManager(
		int(conf.PeersLo),
		int(conf.PeersHi),
		cmgr.WithGracePeriod(conf.PeersGrace),
		cmgr.WithSilencePeriod(time.Minute),
		cmgr.WithEmergencyTrim(true))
}

func (conf *Config) TargetPeers() uint {
	return conf.PeersLo
}

func (conf *Config) Disabled() bool {
	return conf.DisableP2P
}

func (conf *Config) PeerScoringParams() *ScoringParams {
	if conf.ScoringParams == nil {
		return nil
	}
	return conf.ScoringParams
}

func (conf *Config) BanPeers() bool {
	return conf.BanningEnabled
}

func (conf *Config) BanThreshold() float64 {
	return conf.BanningThreshold
}

func (conf *Config) BanDuration() time.Duration {
	return conf.BanningDuration
}

func (conf *Config) ReqRespSyncEnabled() bool {
	return conf.EnableReqRespSync
}

const maxMeshParam = 1000

func (conf *Config) Check() error {
	if conf.DisableP2P {
		return nil
	}
	if conf.Store == nil {
		return errors.New("p2p requires a persistent or in-memory peerstore, but found none")
	}
	if !conf.NoDiscovery {
		if conf.DiscoveryDB == nil {
			return errors.New("discovery requires a persistent or in-memory discv5 db, but found none")
		}
	}
	if conf.PeersLo == 0 || conf.PeersHi == 0 || conf.PeersLo > conf.PeersHi {
		return fmt.Errorf("peers lo/hi tides are invalid: %d, %d", conf.PeersLo, conf.PeersHi)
	}
	if conf.MeshD <= 0 || conf.MeshD > maxMeshParam {
		return fmt.Errorf("mesh D param must not be 0 or exceed %d, but got %d", maxMeshParam, conf.MeshD)
	}
	if conf.MeshDLo <= 0 || conf.MeshDLo > maxMeshParam {
		return fmt.Errorf("mesh Dlo param must not be 0 or exceed %d, but got %d", maxMeshParam, conf.MeshDLo)
	}
	if conf.MeshDHi <= 0 || conf.MeshDHi > maxMeshParam {
		return fmt.Errorf("mesh Dhi param must not be 0 or exceed %d, but got %d", maxMeshParam, conf.MeshDHi)
	}
	if conf.MeshDLazy <= 0 || conf.MeshDLazy > maxMeshParam {
		return fmt.Errorf("mesh Dlazy param must not be 0 or exceed %d, but got %d", maxMeshParam, conf.MeshDLazy)
	}
	return nil
}
