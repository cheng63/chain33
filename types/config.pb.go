// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Config struct {
	Title      string      `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Log        *Log        `protobuf:"bytes,2,opt,name=log" json:"log,omitempty"`
	Store      *Store      `protobuf:"bytes,3,opt,name=store" json:"store,omitempty"`
	Consensus  *Consensus  `protobuf:"bytes,5,opt,name=consensus" json:"consensus,omitempty"`
	MemPool    *MemPool    `protobuf:"bytes,6,opt,name=memPool" json:"memPool,omitempty"`
	BlockChain *BlockChain `protobuf:"bytes,7,opt,name=blockChain" json:"blockChain,omitempty"`
	Wallet     *Wallet     `protobuf:"bytes,8,opt,name=wallet" json:"wallet,omitempty"`
	P2P        *P2P        `protobuf:"bytes,9,opt,name=p2p" json:"p2p,omitempty"`
	Rpc        *Rpc        `protobuf:"bytes,10,opt,name=rpc" json:"rpc,omitempty"`
	Exec       *Exec       `protobuf:"bytes,11,opt,name=exec" json:"exec,omitempty"`
	TestNet    bool        `protobuf:"varint,12,opt,name=testNet" json:"testNet,omitempty"`
	FixTime    bool        `protobuf:"varint,13,opt,name=fixTime" json:"fixTime,omitempty"`
	Pprof      *Pprof      `protobuf:"bytes,14,opt,name=pprof" json:"pprof,omitempty"`
	Auth       *Authority  `protobuf:"bytes,15,opt,name=auth" json:"auth,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Config) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Config) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Config) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Config) GetConsensus() *Consensus {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *Config) GetMemPool() *MemPool {
	if m != nil {
		return m.MemPool
	}
	return nil
}

func (m *Config) GetBlockChain() *BlockChain {
	if m != nil {
		return m.BlockChain
	}
	return nil
}

func (m *Config) GetWallet() *Wallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func (m *Config) GetP2P() *P2P {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Config) GetRpc() *Rpc {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetExec() *Exec {
	if m != nil {
		return m.Exec
	}
	return nil
}

func (m *Config) GetTestNet() bool {
	if m != nil {
		return m.TestNet
	}
	return false
}

func (m *Config) GetFixTime() bool {
	if m != nil {
		return m.FixTime
	}
	return false
}

func (m *Config) GetPprof() *Pprof {
	if m != nil {
		return m.Pprof
	}
	return nil
}

func (m *Config) GetAuth() *Authority {
	if m != nil {
		return m.Auth
	}
	return nil
}

type Log struct {
	// 日志级别，支持debug(dbug)/info/warn/error(eror)/crit
	Loglevel        string `protobuf:"bytes,1,opt,name=loglevel" json:"loglevel,omitempty"`
	LogConsoleLevel string `protobuf:"bytes,2,opt,name=logConsoleLevel" json:"logConsoleLevel,omitempty"`
	// 日志文件名，可带目录，所有生成的日志文件都放到此目录下
	LogFile string `protobuf:"bytes,3,opt,name=logFile" json:"logFile,omitempty"`
	// 单个日志文件的最大值（单位：兆）
	MaxFileSize uint32 `protobuf:"varint,4,opt,name=maxFileSize" json:"maxFileSize,omitempty"`
	// 最多保存的历史日志文件个数
	MaxBackups uint32 `protobuf:"varint,5,opt,name=maxBackups" json:"maxBackups,omitempty"`
	// 最多保存的历史日志消息（单位：天）
	MaxAge uint32 `protobuf:"varint,6,opt,name=maxAge" json:"maxAge,omitempty"`
	// 日志文件名是否使用本地事件（否则使用UTC时间）
	LocalTime bool `protobuf:"varint,7,opt,name=localTime" json:"localTime,omitempty"`
	// 历史日志文件是否压缩（压缩格式为gz）
	Compress bool `protobuf:"varint,8,opt,name=compress" json:"compress,omitempty"`
	// 是否打印调用源文件和行号
	CallerFile bool `protobuf:"varint,9,opt,name=callerFile" json:"callerFile,omitempty"`
	// 是否打印调用方法
	CallerFunction bool `protobuf:"varint,10,opt,name=callerFunction" json:"callerFunction,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Log) GetLoglevel() string {
	if m != nil {
		return m.Loglevel
	}
	return ""
}

func (m *Log) GetLogConsoleLevel() string {
	if m != nil {
		return m.LogConsoleLevel
	}
	return ""
}

func (m *Log) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *Log) GetMaxFileSize() uint32 {
	if m != nil {
		return m.MaxFileSize
	}
	return 0
}

func (m *Log) GetMaxBackups() uint32 {
	if m != nil {
		return m.MaxBackups
	}
	return 0
}

func (m *Log) GetMaxAge() uint32 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *Log) GetLocalTime() bool {
	if m != nil {
		return m.LocalTime
	}
	return false
}

func (m *Log) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

func (m *Log) GetCallerFile() bool {
	if m != nil {
		return m.CallerFile
	}
	return false
}

func (m *Log) GetCallerFunction() bool {
	if m != nil {
		return m.CallerFunction
	}
	return false
}

type MemPool struct {
	PoolCacheSize      int64 `protobuf:"varint,1,opt,name=poolCacheSize" json:"poolCacheSize,omitempty"`
	MinTxFee           int64 `protobuf:"varint,2,opt,name=minTxFee" json:"minTxFee,omitempty"`
	ForceAccept        bool  `protobuf:"varint,3,opt,name=forceAccept" json:"forceAccept,omitempty"`
	MaxTxNumPerAccount int64 `protobuf:"varint,4,opt,name=maxTxNumPerAccount" json:"maxTxNumPerAccount,omitempty"`
}

func (m *MemPool) Reset()                    { *m = MemPool{} }
func (m *MemPool) String() string            { return proto.CompactTextString(m) }
func (*MemPool) ProtoMessage()               {}
func (*MemPool) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *MemPool) GetPoolCacheSize() int64 {
	if m != nil {
		return m.PoolCacheSize
	}
	return 0
}

func (m *MemPool) GetMinTxFee() int64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

func (m *MemPool) GetForceAccept() bool {
	if m != nil {
		return m.ForceAccept
	}
	return false
}

func (m *MemPool) GetMaxTxNumPerAccount() int64 {
	if m != nil {
		return m.MaxTxNumPerAccount
	}
	return 0
}

type Consensus struct {
	Name                      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Genesis                   string   `protobuf:"bytes,2,opt,name=genesis" json:"genesis,omitempty"`
	Minerstart                bool     `protobuf:"varint,3,opt,name=minerstart" json:"minerstart,omitempty"`
	GenesisBlockTime          int64    `protobuf:"varint,4,opt,name=genesisBlockTime" json:"genesisBlockTime,omitempty"`
	HotkeyAddr                string   `protobuf:"bytes,5,opt,name=hotkeyAddr" json:"hotkeyAddr,omitempty"`
	ForceMining               bool     `protobuf:"varint,6,opt,name=forceMining" json:"forceMining,omitempty"`
	NodeId                    int64    `protobuf:"varint,7,opt,name=NodeId" json:"NodeId,omitempty"`
	PeersURL                  string   `protobuf:"bytes,8,opt,name=PeersURL" json:"PeersURL,omitempty"`
	ClientAddr                string   `protobuf:"bytes,9,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
	RaftApiPort               int64    `protobuf:"varint,15,opt,name=raftApiPort" json:"raftApiPort,omitempty"`
	IsNewJoinNode             bool     `protobuf:"varint,16,opt,name=isNewJoinNode" json:"isNewJoinNode,omitempty"`
	ReadOnlyPeersURL          string   `protobuf:"bytes,17,opt,name=readOnlyPeersURL" json:"readOnlyPeersURL,omitempty"`
	AddPeersURL               string   `protobuf:"bytes,18,opt,name=addPeersURL" json:"addPeersURL,omitempty"`
	DefaultSnapCount          int64    `protobuf:"varint,19,opt,name=defaultSnapCount" json:"defaultSnapCount,omitempty"`
	WriteBlockSeconds         int64    `protobuf:"varint,20,opt,name=writeBlockSeconds" json:"writeBlockSeconds,omitempty"`
	HeartbeatTick             int32    `protobuf:"varint,21,opt,name=heartbeatTick" json:"heartbeatTick,omitempty"`
	ParaRemoteGrpcClient      string   `protobuf:"bytes,22,opt,name=paraRemoteGrpcClient" json:"paraRemoteGrpcClient,omitempty"`
	StartHeight               int64    `protobuf:"varint,23,opt,name=startHeight" json:"startHeight,omitempty"`
	EmptyBlockInterval        int64    `protobuf:"varint,24,opt,name=emptyBlockInterval" json:"emptyBlockInterval,omitempty"`
	AuthAccount               string   `protobuf:"bytes,25,opt,name=authAccount" json:"authAccount,omitempty"`
	WaitBlocks4CommitMsg      int32    `protobuf:"varint,26,opt,name=waitBlocks4CommitMsg" json:"waitBlocks4CommitMsg,omitempty"`
	TimeoutPropose            int32    `protobuf:"varint,30,opt,name=timeoutPropose" json:"timeoutPropose,omitempty"`
	TimeoutProposeDelta       int32    `protobuf:"varint,31,opt,name=timeoutProposeDelta" json:"timeoutProposeDelta,omitempty"`
	TimeoutPrevote            int32    `protobuf:"varint,32,opt,name=timeoutPrevote" json:"timeoutPrevote,omitempty"`
	TimeoutPrevoteDelta       int32    `protobuf:"varint,33,opt,name=timeoutPrevoteDelta" json:"timeoutPrevoteDelta,omitempty"`
	TimeoutPrecommit          int32    `protobuf:"varint,34,opt,name=timeoutPrecommit" json:"timeoutPrecommit,omitempty"`
	TimeoutPrecommitDelta     int32    `protobuf:"varint,35,opt,name=timeoutPrecommitDelta" json:"timeoutPrecommitDelta,omitempty"`
	TimeoutCommit             int32    `protobuf:"varint,36,opt,name=timeoutCommit" json:"timeoutCommit,omitempty"`
	SkipTimeoutCommit         bool     `protobuf:"varint,37,opt,name=skipTimeoutCommit" json:"skipTimeoutCommit,omitempty"`
	CreateEmptyBlocks         bool     `protobuf:"varint,38,opt,name=createEmptyBlocks" json:"createEmptyBlocks,omitempty"`
	CreateEmptyBlocksInterval int32    `protobuf:"varint,39,opt,name=createEmptyBlocksInterval" json:"createEmptyBlocksInterval,omitempty"`
	Seeds                     []string `protobuf:"bytes,40,rep,name=seeds" json:"seeds,omitempty"`
}

func (m *Consensus) Reset()                    { *m = Consensus{} }
func (m *Consensus) String() string            { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()               {}
func (*Consensus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Consensus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consensus) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *Consensus) GetMinerstart() bool {
	if m != nil {
		return m.Minerstart
	}
	return false
}

func (m *Consensus) GetGenesisBlockTime() int64 {
	if m != nil {
		return m.GenesisBlockTime
	}
	return 0
}

func (m *Consensus) GetHotkeyAddr() string {
	if m != nil {
		return m.HotkeyAddr
	}
	return ""
}

func (m *Consensus) GetForceMining() bool {
	if m != nil {
		return m.ForceMining
	}
	return false
}

func (m *Consensus) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *Consensus) GetPeersURL() string {
	if m != nil {
		return m.PeersURL
	}
	return ""
}

func (m *Consensus) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *Consensus) GetRaftApiPort() int64 {
	if m != nil {
		return m.RaftApiPort
	}
	return 0
}

func (m *Consensus) GetIsNewJoinNode() bool {
	if m != nil {
		return m.IsNewJoinNode
	}
	return false
}

func (m *Consensus) GetReadOnlyPeersURL() string {
	if m != nil {
		return m.ReadOnlyPeersURL
	}
	return ""
}

func (m *Consensus) GetAddPeersURL() string {
	if m != nil {
		return m.AddPeersURL
	}
	return ""
}

func (m *Consensus) GetDefaultSnapCount() int64 {
	if m != nil {
		return m.DefaultSnapCount
	}
	return 0
}

func (m *Consensus) GetWriteBlockSeconds() int64 {
	if m != nil {
		return m.WriteBlockSeconds
	}
	return 0
}

func (m *Consensus) GetHeartbeatTick() int32 {
	if m != nil {
		return m.HeartbeatTick
	}
	return 0
}

func (m *Consensus) GetParaRemoteGrpcClient() string {
	if m != nil {
		return m.ParaRemoteGrpcClient
	}
	return ""
}

func (m *Consensus) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *Consensus) GetEmptyBlockInterval() int64 {
	if m != nil {
		return m.EmptyBlockInterval
	}
	return 0
}

func (m *Consensus) GetAuthAccount() string {
	if m != nil {
		return m.AuthAccount
	}
	return ""
}

func (m *Consensus) GetWaitBlocks4CommitMsg() int32 {
	if m != nil {
		return m.WaitBlocks4CommitMsg
	}
	return 0
}

func (m *Consensus) GetTimeoutPropose() int32 {
	if m != nil {
		return m.TimeoutPropose
	}
	return 0
}

func (m *Consensus) GetTimeoutProposeDelta() int32 {
	if m != nil {
		return m.TimeoutProposeDelta
	}
	return 0
}

func (m *Consensus) GetTimeoutPrevote() int32 {
	if m != nil {
		return m.TimeoutPrevote
	}
	return 0
}

func (m *Consensus) GetTimeoutPrevoteDelta() int32 {
	if m != nil {
		return m.TimeoutPrevoteDelta
	}
	return 0
}

func (m *Consensus) GetTimeoutPrecommit() int32 {
	if m != nil {
		return m.TimeoutPrecommit
	}
	return 0
}

func (m *Consensus) GetTimeoutPrecommitDelta() int32 {
	if m != nil {
		return m.TimeoutPrecommitDelta
	}
	return 0
}

func (m *Consensus) GetTimeoutCommit() int32 {
	if m != nil {
		return m.TimeoutCommit
	}
	return 0
}

func (m *Consensus) GetSkipTimeoutCommit() bool {
	if m != nil {
		return m.SkipTimeoutCommit
	}
	return false
}

func (m *Consensus) GetCreateEmptyBlocks() bool {
	if m != nil {
		return m.CreateEmptyBlocks
	}
	return false
}

func (m *Consensus) GetCreateEmptyBlocksInterval() int32 {
	if m != nil {
		return m.CreateEmptyBlocksInterval
	}
	return 0
}

func (m *Consensus) GetSeeds() []string {
	if m != nil {
		return m.Seeds
	}
	return nil
}

type Wallet struct {
	MinFee         int64    `protobuf:"varint,1,opt,name=minFee" json:"minFee,omitempty"`
	Driver         string   `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath         string   `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache        int32    `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
	SignType       string   `protobuf:"bytes,5,opt,name=signType" json:"signType,omitempty"`
	ForceMining    bool     `protobuf:"varint,6,opt,name=forceMining" json:"forceMining,omitempty"`
	Minerdisable   bool     `protobuf:"varint,7,opt,name=minerdisable" json:"minerdisable,omitempty"`
	Minerwhitelist []string `protobuf:"bytes,8,rep,name=minerwhitelist" json:"minerwhitelist,omitempty"`
}

func (m *Wallet) Reset()                    { *m = Wallet{} }
func (m *Wallet) String() string            { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()               {}
func (*Wallet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Wallet) GetMinFee() int64 {
	if m != nil {
		return m.MinFee
	}
	return 0
}

func (m *Wallet) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Wallet) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *Wallet) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

func (m *Wallet) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

func (m *Wallet) GetForceMining() bool {
	if m != nil {
		return m.ForceMining
	}
	return false
}

func (m *Wallet) GetMinerdisable() bool {
	if m != nil {
		return m.Minerdisable
	}
	return false
}

func (m *Wallet) GetMinerwhitelist() []string {
	if m != nil {
		return m.Minerwhitelist
	}
	return nil
}

type Store struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver  string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath  string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache int32  `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *Store) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Store) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Store) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *Store) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

type BlockChain struct {
	DefCacheSize          int64  `protobuf:"varint,1,opt,name=defCacheSize" json:"defCacheSize,omitempty"`
	MaxFetchBlockNum      int64  `protobuf:"varint,2,opt,name=maxFetchBlockNum" json:"maxFetchBlockNum,omitempty"`
	TimeoutSeconds        int64  `protobuf:"varint,3,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	BatchBlockNum         int64  `protobuf:"varint,4,opt,name=batchBlockNum" json:"batchBlockNum,omitempty"`
	Driver                string `protobuf:"bytes,5,opt,name=driver" json:"driver,omitempty"`
	DbPath                string `protobuf:"bytes,6,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache               int32  `protobuf:"varint,7,opt,name=dbCache" json:"dbCache,omitempty"`
	IsStrongConsistency   bool   `protobuf:"varint,8,opt,name=isStrongConsistency" json:"isStrongConsistency,omitempty"`
	SingleMode            bool   `protobuf:"varint,9,opt,name=singleMode" json:"singleMode,omitempty"`
	Batchsync             bool   `protobuf:"varint,10,opt,name=batchsync" json:"batchsync,omitempty"`
	IsRecordBlockSequence bool   `protobuf:"varint,11,opt,name=isRecordBlockSequence" json:"isRecordBlockSequence,omitempty"`
	IsParaChain           bool   `protobuf:"varint,12,opt,name=isParaChain" json:"isParaChain,omitempty"`
	EnableTxQuickIndex    bool   `protobuf:"varint,13,opt,name=enableTxQuickIndex" json:"enableTxQuickIndex,omitempty"`
}

func (m *BlockChain) Reset()                    { *m = BlockChain{} }
func (m *BlockChain) String() string            { return proto.CompactTextString(m) }
func (*BlockChain) ProtoMessage()               {}
func (*BlockChain) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *BlockChain) GetDefCacheSize() int64 {
	if m != nil {
		return m.DefCacheSize
	}
	return 0
}

func (m *BlockChain) GetMaxFetchBlockNum() int64 {
	if m != nil {
		return m.MaxFetchBlockNum
	}
	return 0
}

func (m *BlockChain) GetTimeoutSeconds() int64 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *BlockChain) GetBatchBlockNum() int64 {
	if m != nil {
		return m.BatchBlockNum
	}
	return 0
}

func (m *BlockChain) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *BlockChain) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *BlockChain) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

func (m *BlockChain) GetIsStrongConsistency() bool {
	if m != nil {
		return m.IsStrongConsistency
	}
	return false
}

func (m *BlockChain) GetSingleMode() bool {
	if m != nil {
		return m.SingleMode
	}
	return false
}

func (m *BlockChain) GetBatchsync() bool {
	if m != nil {
		return m.Batchsync
	}
	return false
}

func (m *BlockChain) GetIsRecordBlockSequence() bool {
	if m != nil {
		return m.IsRecordBlockSequence
	}
	return false
}

func (m *BlockChain) GetIsParaChain() bool {
	if m != nil {
		return m.IsParaChain
	}
	return false
}

func (m *BlockChain) GetEnableTxQuickIndex() bool {
	if m != nil {
		return m.EnableTxQuickIndex
	}
	return false
}

type P2P struct {
	SeedPort        int32    `protobuf:"varint,1,opt,name=seedPort" json:"seedPort,omitempty"`
	Driver          string   `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath          string   `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache         int32    `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
	GrpcLogFile     string   `protobuf:"bytes,5,opt,name=grpcLogFile" json:"grpcLogFile,omitempty"`
	IsSeed          bool     `protobuf:"varint,6,opt,name=isSeed" json:"isSeed,omitempty"`
	ServerStart     bool     `protobuf:"varint,7,opt,name=serverStart" json:"serverStart,omitempty"`
	Seeds           []string `protobuf:"bytes,8,rep,name=seeds" json:"seeds,omitempty"`
	Enable          bool     `protobuf:"varint,9,opt,name=enable" json:"enable,omitempty"`
	MsgCacheSize    int32    `protobuf:"varint,10,opt,name=msgCacheSize" json:"msgCacheSize,omitempty"`
	Version         int32    `protobuf:"varint,11,opt,name=version" json:"version,omitempty"`
	VerMix          int32    `protobuf:"varint,12,opt,name=verMix" json:"verMix,omitempty"`
	VerMax          int32    `protobuf:"varint,13,opt,name=verMax" json:"verMax,omitempty"`
	InnerSeedEnable bool     `protobuf:"varint,14,opt,name=innerSeedEnable" json:"innerSeedEnable,omitempty"`
	InnerBounds     int32    `protobuf:"varint,15,opt,name=innerBounds" json:"innerBounds,omitempty"`
	UseGithub       bool     `protobuf:"varint,16,opt,name=useGithub" json:"useGithub,omitempty"`
}

func (m *P2P) Reset()                    { *m = P2P{} }
func (m *P2P) String() string            { return proto.CompactTextString(m) }
func (*P2P) ProtoMessage()               {}
func (*P2P) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *P2P) GetSeedPort() int32 {
	if m != nil {
		return m.SeedPort
	}
	return 0
}

func (m *P2P) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *P2P) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *P2P) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

func (m *P2P) GetGrpcLogFile() string {
	if m != nil {
		return m.GrpcLogFile
	}
	return ""
}

func (m *P2P) GetIsSeed() bool {
	if m != nil {
		return m.IsSeed
	}
	return false
}

func (m *P2P) GetServerStart() bool {
	if m != nil {
		return m.ServerStart
	}
	return false
}

func (m *P2P) GetSeeds() []string {
	if m != nil {
		return m.Seeds
	}
	return nil
}

func (m *P2P) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *P2P) GetMsgCacheSize() int32 {
	if m != nil {
		return m.MsgCacheSize
	}
	return 0
}

func (m *P2P) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *P2P) GetVerMix() int32 {
	if m != nil {
		return m.VerMix
	}
	return 0
}

func (m *P2P) GetVerMax() int32 {
	if m != nil {
		return m.VerMax
	}
	return 0
}

func (m *P2P) GetInnerSeedEnable() bool {
	if m != nil {
		return m.InnerSeedEnable
	}
	return false
}

func (m *P2P) GetInnerBounds() int32 {
	if m != nil {
		return m.InnerBounds
	}
	return 0
}

func (m *P2P) GetUseGithub() bool {
	if m != nil {
		return m.UseGithub
	}
	return false
}

type Rpc struct {
	JrpcBindAddr      string   `protobuf:"bytes,1,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	GrpcBindAddr      string   `protobuf:"bytes,2,opt,name=grpcBindAddr" json:"grpcBindAddr,omitempty"`
	Whitlist          []string `protobuf:"bytes,3,rep,name=whitlist" json:"whitlist,omitempty"`
	Whitelist         []string `protobuf:"bytes,4,rep,name=whitelist" json:"whitelist,omitempty"`
	JrpcFuncWhitelist []string `protobuf:"bytes,5,rep,name=jrpcFuncWhitelist" json:"jrpcFuncWhitelist,omitempty"`
	GrpcFuncWhitelist []string `protobuf:"bytes,6,rep,name=grpcFuncWhitelist" json:"grpcFuncWhitelist,omitempty"`
	JrpcFuncBlacklist []string `protobuf:"bytes,7,rep,name=jrpcFuncBlacklist" json:"jrpcFuncBlacklist,omitempty"`
	GrpcFuncBlacklist []string `protobuf:"bytes,8,rep,name=grpcFuncBlacklist" json:"grpcFuncBlacklist,omitempty"`
	MainnetJrpcAddr   string   `protobuf:"bytes,9,opt,name=mainnetJrpcAddr" json:"mainnetJrpcAddr,omitempty"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *Rpc) GetJrpcBindAddr() string {
	if m != nil {
		return m.JrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetGrpcBindAddr() string {
	if m != nil {
		return m.GrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetWhitlist() []string {
	if m != nil {
		return m.Whitlist
	}
	return nil
}

func (m *Rpc) GetWhitelist() []string {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func (m *Rpc) GetJrpcFuncWhitelist() []string {
	if m != nil {
		return m.JrpcFuncWhitelist
	}
	return nil
}

func (m *Rpc) GetGrpcFuncWhitelist() []string {
	if m != nil {
		return m.GrpcFuncWhitelist
	}
	return nil
}

func (m *Rpc) GetJrpcFuncBlacklist() []string {
	if m != nil {
		return m.JrpcFuncBlacklist
	}
	return nil
}

func (m *Rpc) GetGrpcFuncBlacklist() []string {
	if m != nil {
		return m.GrpcFuncBlacklist
	}
	return nil
}

func (m *Rpc) GetMainnetJrpcAddr() string {
	if m != nil {
		return m.MainnetJrpcAddr
	}
	return ""
}

type Exec struct {
	MinExecFee int64 `protobuf:"varint,1,opt,name=minExecFee" json:"minExecFee,omitempty"`
	IsFree     bool  `protobuf:"varint,2,opt,name=isFree" json:"isFree,omitempty"`
	EnableStat bool  `protobuf:"varint,3,opt,name=enableStat" json:"enableStat,omitempty"`
	EnableMVCC bool  `protobuf:"varint,4,opt,name=enableMVCC" json:"enableMVCC,omitempty"`
}

func (m *Exec) Reset()                    { *m = Exec{} }
func (m *Exec) String() string            { return proto.CompactTextString(m) }
func (*Exec) ProtoMessage()               {}
func (*Exec) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *Exec) GetMinExecFee() int64 {
	if m != nil {
		return m.MinExecFee
	}
	return 0
}

func (m *Exec) GetIsFree() bool {
	if m != nil {
		return m.IsFree
	}
	return false
}

func (m *Exec) GetEnableStat() bool {
	if m != nil {
		return m.EnableStat
	}
	return false
}

func (m *Exec) GetEnableMVCC() bool {
	if m != nil {
		return m.EnableMVCC
	}
	return false
}

type Authority struct {
	Enable     bool   `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
	CryptoPath string `protobuf:"bytes,2,opt,name=cryptoPath" json:"cryptoPath,omitempty"`
	SignType   string `protobuf:"bytes,3,opt,name=signType" json:"signType,omitempty"`
}

func (m *Authority) Reset()                    { *m = Authority{} }
func (m *Authority) String() string            { return proto.CompactTextString(m) }
func (*Authority) ProtoMessage()               {}
func (*Authority) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *Authority) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Authority) GetCryptoPath() string {
	if m != nil {
		return m.CryptoPath
	}
	return ""
}

func (m *Authority) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

type Pprof struct {
	ListenAddr string `protobuf:"bytes,1,opt,name=listenAddr" json:"listenAddr,omitempty"`
}

func (m *Pprof) Reset()                    { *m = Pprof{} }
func (m *Pprof) String() string            { return proto.CompactTextString(m) }
func (*Pprof) ProtoMessage()               {}
func (*Pprof) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *Pprof) GetListenAddr() string {
	if m != nil {
		return m.ListenAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "types.Config")
	proto.RegisterType((*Log)(nil), "types.Log")
	proto.RegisterType((*MemPool)(nil), "types.MemPool")
	proto.RegisterType((*Consensus)(nil), "types.Consensus")
	proto.RegisterType((*Wallet)(nil), "types.Wallet")
	proto.RegisterType((*Store)(nil), "types.Store")
	proto.RegisterType((*BlockChain)(nil), "types.BlockChain")
	proto.RegisterType((*P2P)(nil), "types.P2P")
	proto.RegisterType((*Rpc)(nil), "types.Rpc")
	proto.RegisterType((*Exec)(nil), "types.Exec")
	proto.RegisterType((*Authority)(nil), "types.Authority")
	proto.RegisterType((*Pprof)(nil), "types.Pprof")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x58, 0xcd, 0x6e, 0x23, 0xc7,
	0x11, 0x06, 0x77, 0x44, 0x49, 0x6c, 0xad, 0xf6, 0x67, 0x76, 0xed, 0x8c, 0x0d, 0xc3, 0x66, 0x98,
	0xb5, 0x4d, 0x04, 0x81, 0xe2, 0x48, 0x3e, 0xe6, 0x22, 0x31, 0xbb, 0x8e, 0x0d, 0x49, 0x61, 0x5a,
	0x4a, 0x0c, 0xe4, 0x12, 0x34, 0x7b, 0x4a, 0xc3, 0x8e, 0x66, 0xa6, 0x27, 0xdd, 0x4d, 0x89, 0xcc,
	0x21, 0x6f, 0x92, 0x87, 0xc9, 0x03, 0xe4, 0x41, 0x02, 0x04, 0xc8, 0x25, 0x97, 0xdc, 0x82, 0xaa,
	0x6e, 0xce, 0x0f, 0xc9, 0x05, 0x72, 0xf0, 0x49, 0xac, 0xef, 0xfb, 0xa6, 0xa7, 0xeb, 0xa7, 0xab,
	0x6b, 0xc4, 0x9e, 0x4a, 0x5d, 0xde, 0xa9, 0xec, 0xa4, 0x32, 0xda, 0xe9, 0xb8, 0xef, 0x56, 0x15,
	0xd8, 0xd1, 0xbf, 0x23, 0xb6, 0x3f, 0x21, 0x3c, 0x7e, 0xcd, 0xfa, 0x4e, 0xb9, 0x1c, 0x92, 0xde,
	0xb0, 0x37, 0x1e, 0x70, 0x6f, 0xc4, 0x9f, 0xb0, 0x28, 0xd7, 0x59, 0xf2, 0x64, 0xd8, 0x1b, 0x1f,
	0x9d, 0xb2, 0x13, 0x7a, 0xea, 0xe4, 0x52, 0x67, 0x1c, 0xe1, 0x78, 0xc4, 0xfa, 0xd6, 0x69, 0x03,
	0x49, 0x44, 0xfc, 0xd3, 0xc0, 0xdf, 0x20, 0xc6, 0x3d, 0x15, 0x9f, 0xb0, 0x81, 0xd4, 0xa5, 0x85,
	0xd2, 0x2e, 0x6c, 0xd2, 0x27, 0xdd, 0x8b, 0xa0, 0x9b, 0xac, 0x71, 0xde, 0x48, 0xe2, 0x31, 0x3b,
	0x28, 0xa0, 0x98, 0x6a, 0x9d, 0x27, 0xfb, 0xa4, 0x7e, 0x16, 0xd4, 0x57, 0x1e, 0xe5, 0x6b, 0x3a,
	0xfe, 0x05, 0x63, 0xb3, 0x5c, 0xcb, 0xfb, 0xc9, 0x5c, 0xa8, 0x32, 0x39, 0x20, 0xf1, 0xcb, 0x20,
	0xbe, 0xa8, 0x09, 0xde, 0x12, 0xc5, 0x9f, 0xb3, 0xfd, 0x47, 0x91, 0xe7, 0xe0, 0x92, 0x43, 0x92,
	0x1f, 0x07, 0xf9, 0xf7, 0x04, 0xf2, 0x40, 0xa2, 0xd7, 0xd5, 0x69, 0x95, 0x0c, 0x3a, 0x5e, 0x4f,
	0x4f, 0xa7, 0x1c, 0x61, 0x64, 0x4d, 0x25, 0x13, 0xd6, 0x61, 0x79, 0x25, 0x39, 0xc2, 0xf1, 0x67,
	0x6c, 0x0f, 0x96, 0x20, 0x93, 0x23, 0xa2, 0x8f, 0x02, 0xfd, 0x76, 0x09, 0x92, 0x13, 0x11, 0x27,
	0xec, 0xc0, 0x81, 0x75, 0xd7, 0xe0, 0x92, 0xa7, 0xc3, 0xde, 0xf8, 0x90, 0xaf, 0x4d, 0x64, 0xee,
	0xd4, 0xf2, 0x56, 0x15, 0x90, 0x1c, 0x7b, 0x26, 0x98, 0x18, 0xe8, 0xaa, 0x32, 0xfa, 0x2e, 0x79,
	0xd6, 0x09, 0xf4, 0x14, 0x31, 0xee, 0xa9, 0xf8, 0x0d, 0xdb, 0x13, 0x0b, 0x37, 0x4f, 0x9e, 0x77,
	0x62, 0x7c, 0xbe, 0x70, 0x73, 0x6d, 0x94, 0x5b, 0x71, 0x62, 0x47, 0x7f, 0x7f, 0xc2, 0xa2, 0x4b,
	0x9d, 0xc5, 0x1f, 0xb3, 0xc3, 0x5c, 0x67, 0x39, 0x3c, 0x40, 0x1e, 0x32, 0x5e, 0xdb, 0xf1, 0x98,
	0x3d, 0xcf, 0x75, 0x86, 0xd9, 0xd1, 0x39, 0x5c, 0x92, 0xe4, 0x09, 0x49, 0x36, 0x61, 0xdc, 0x71,
	0xae, 0xb3, 0x77, 0x2a, 0xf7, 0x25, 0x30, 0xe0, 0x6b, 0x33, 0x1e, 0xb2, 0xa3, 0x42, 0x2c, 0xf1,
	0xe7, 0x8d, 0xfa, 0x0b, 0x24, 0x7b, 0xc3, 0xde, 0xf8, 0x98, 0xb7, 0xa1, 0xf8, 0x53, 0xc6, 0x0a,
	0xb1, 0xbc, 0x10, 0xf2, 0x7e, 0x51, 0xf9, 0xca, 0x38, 0xe6, 0x2d, 0x24, 0xfe, 0x90, 0xed, 0x17,
	0x62, 0x79, 0x9e, 0x01, 0xd5, 0xc1, 0x31, 0x0f, 0x56, 0xfc, 0x09, 0x1b, 0xe4, 0x5a, 0x8a, 0x9c,
	0xe2, 0x74, 0x40, 0x71, 0x6a, 0x00, 0xf4, 0x4b, 0xea, 0xa2, 0x32, 0x60, 0x2d, 0xe5, 0xf8, 0x90,
	0xd7, 0x36, 0xbe, 0x51, 0x62, 0x82, 0x0d, 0x6d, 0x78, 0x40, 0x6c, 0x0b, 0x89, 0xbf, 0x60, 0xcf,
	0x82, 0xb5, 0x28, 0xa5, 0x53, 0xba, 0xa4, 0x1c, 0x1f, 0xf2, 0x0d, 0x74, 0xf4, 0xb7, 0x1e, 0x3b,
	0x08, 0xd5, 0x18, 0xbf, 0x61, 0xc7, 0x95, 0xd6, 0xf9, 0x44, 0xc8, 0xb9, 0xf7, 0x14, 0x83, 0x19,
	0xf1, 0x2e, 0x88, 0xbb, 0x2a, 0x54, 0x79, 0xbb, 0x7c, 0x07, 0x40, 0xa1, 0x8c, 0x78, 0x6d, 0x63,
	0xa4, 0xee, 0xb4, 0x91, 0x70, 0x2e, 0x25, 0x54, 0x8e, 0xe2, 0x78, 0xc8, 0xdb, 0x50, 0x7c, 0xc2,
	0xe2, 0x42, 0x2c, 0x6f, 0x97, 0xd7, 0x8b, 0x62, 0x0a, 0xe6, 0x5c, 0x4a, 0xbd, 0x28, 0x1d, 0x85,
	0x34, 0xe2, 0x3b, 0x98, 0xd1, 0x7f, 0x07, 0x6c, 0x50, 0x9f, 0xad, 0x38, 0x66, 0x7b, 0xa5, 0x28,
	0xd6, 0xe7, 0x9a, 0x7e, 0x63, 0xde, 0x32, 0x28, 0xc1, 0x2a, 0x1b, 0x32, 0xbb, 0x36, 0x29, 0x2b,
	0xaa, 0x04, 0x63, 0x9d, 0x30, 0xeb, 0xcd, 0xb4, 0x90, 0xf8, 0xa7, 0xec, 0x45, 0x90, 0xd2, 0x11,
	0xa3, 0x24, 0xf8, 0x9d, 0x6c, 0xe1, 0xb8, 0xd6, 0x5c, 0xbb, 0x7b, 0x58, 0x9d, 0xa7, 0xa9, 0xa1,
	0x0c, 0x0f, 0x78, 0x0b, 0xa9, 0x3d, 0xbf, 0x52, 0xa5, 0x2a, 0x33, 0x4a, 0xf3, 0xda, 0x73, 0x0f,
	0x61, 0x0d, 0x5c, 0xeb, 0x14, 0xbe, 0x4d, 0x29, 0xd1, 0x11, 0x0f, 0x16, 0xc6, 0x73, 0x0a, 0x60,
	0xec, 0xef, 0xf8, 0x25, 0x65, 0x79, 0xc0, 0x6b, 0x1b, 0xdf, 0x3a, 0xc9, 0x15, 0x94, 0x8e, 0xde,
	0x3a, 0xf0, 0x6f, 0x6d, 0x10, 0x7c, 0xab, 0x11, 0x77, 0xee, 0xbc, 0x52, 0x53, 0x6d, 0x1c, 0x1d,
	0x97, 0x88, 0xb7, 0x21, 0xcc, 0xa9, 0xb2, 0xd7, 0xf0, 0xf8, 0x9d, 0x56, 0x25, 0xbe, 0x30, 0x79,
	0x41, 0x3b, 0xeb, 0x82, 0x18, 0x09, 0x03, 0x22, 0xfd, 0x4d, 0x99, 0xaf, 0xea, 0xbd, 0xbc, 0xa4,
	0xb7, 0x6d, 0xe1, 0xf8, 0x4e, 0x91, 0xa6, 0xb5, 0x2c, 0x26, 0x59, 0x1b, 0xc2, 0xd5, 0x52, 0xb8,
	0x13, 0x8b, 0xdc, 0xdd, 0x94, 0xa2, 0x9a, 0x50, 0x86, 0x5f, 0xf9, 0xb8, 0x6e, 0xe2, 0xf1, 0xcf,
	0xd8, 0xcb, 0x47, 0xa3, 0x1c, 0x50, 0xa4, 0x6f, 0x40, 0xea, 0x32, 0xb5, 0xc9, 0x6b, 0x12, 0x6f,
	0x13, 0xe8, 0xcd, 0x1c, 0x84, 0x71, 0x33, 0x10, 0xee, 0x56, 0xc9, 0xfb, 0xe4, 0x83, 0x61, 0x6f,
	0xdc, 0xe7, 0x5d, 0x30, 0x3e, 0x65, 0xaf, 0x2b, 0x61, 0x04, 0x87, 0x42, 0x3b, 0xf8, 0xc6, 0x54,
	0xd2, 0x47, 0x2c, 0xf9, 0x90, 0xb6, 0xba, 0x93, 0x43, 0xaf, 0xa8, 0x28, 0x7e, 0x0d, 0x2a, 0x9b,
	0xbb, 0xe4, 0x47, 0x3e, 0x92, 0x2d, 0x08, 0x2b, 0x17, 0x8a, 0xca, 0xad, 0x68, 0x43, 0xdf, 0x96,
	0x0e, 0xcc, 0x83, 0xc8, 0x93, 0xc4, 0x57, 0xee, 0x36, 0x43, 0x71, 0x5a, 0xb8, 0xf9, 0xba, 0xc4,
	0x3f, 0x0a, 0x71, 0x6a, 0x20, 0xdc, 0xe7, 0xa3, 0x50, 0x8e, 0x1e, 0xb3, 0x5f, 0x4f, 0x74, 0x51,
	0x28, 0x77, 0x65, 0xb3, 0xe4, 0x63, 0x72, 0x6a, 0x27, 0x87, 0xe7, 0xda, 0xa9, 0x02, 0xf4, 0xc2,
	0x4d, 0x8d, 0xae, 0xb4, 0x85, 0xe4, 0x53, 0x52, 0x6f, 0xa0, 0xf1, 0x57, 0xec, 0x55, 0x17, 0xf9,
	0x15, 0xe4, 0x4e, 0x24, 0x9f, 0x91, 0x78, 0x17, 0xd5, 0x59, 0x19, 0x1e, 0xb4, 0x83, 0x64, 0xb8,
	0xb1, 0x32, 0xa1, 0x9d, 0x95, 0x09, 0xf1, 0x2b, 0xff, 0x78, 0x63, 0xe5, 0x86, 0xc2, 0x7a, 0x68,
	0x60, 0x49, 0xae, 0x24, 0x23, 0x92, 0x6f, 0xe1, 0xf1, 0xd7, 0xec, 0x83, 0x4d, 0xcc, 0xaf, 0xff,
	0x13, 0x7a, 0x60, 0x37, 0x89, 0x75, 0x11, 0x08, 0x1f, 0xa9, 0xe4, 0x8d, 0xaf, 0x8b, 0x0e, 0x88,
	0xb5, 0x66, 0xef, 0x55, 0x75, 0xdb, 0x51, 0x7e, 0x4e, 0xe7, 0x61, 0x9b, 0x40, 0xb5, 0x34, 0x20,
	0x1c, 0xbc, 0xad, 0x73, 0x6b, 0x93, 0x2f, 0xbc, 0x7a, 0x8b, 0x88, 0x7f, 0xc9, 0x3e, 0xda, 0x02,
	0xeb, 0x22, 0xf9, 0x92, 0x76, 0xf3, 0x7e, 0x01, 0x0e, 0x2c, 0x16, 0x20, 0xb5, 0xc9, 0x78, 0x18,
	0xe1, 0xc0, 0x42, 0xc6, 0xe8, 0x3f, 0x3d, 0xb6, 0xef, 0x6f, 0x73, 0xba, 0x40, 0x54, 0x89, 0x2d,
	0xd7, 0xf7, 0xe4, 0x60, 0x21, 0x9e, 0x1a, 0xf5, 0x00, 0x26, 0xf4, 0xbe, 0x60, 0x11, 0x3e, 0x9b,
	0x0a, 0x37, 0x0f, 0x77, 0x59, 0xb0, 0xb0, 0x59, 0xa6, 0x33, 0xea, 0xe5, 0xd4, 0xe9, 0xfa, 0x7c,
	0x6d, 0x62, 0x1b, 0xb2, 0x2a, 0x2b, 0x6f, 0x57, 0x15, 0x84, 0xf6, 0x56, 0xdb, 0xff, 0x47, 0x73,
	0x1b, 0xb1, 0xa7, 0xd4, 0x58, 0x53, 0x65, 0xc5, 0x2c, 0x5f, 0xdf, 0x65, 0x1d, 0x0c, 0x0b, 0x8c,
	0xec, 0xc7, 0xb9, 0x72, 0x90, 0x2b, 0x8b, 0x83, 0x0b, 0x7a, 0xbb, 0x81, 0x8e, 0x80, 0xf5, 0x69,
	0xea, 0xda, 0xd9, 0xed, 0x7f, 0x30, 0x87, 0x47, 0xff, 0x8a, 0x18, 0x6b, 0x46, 0x2b, 0xf4, 0x20,
	0x85, 0xbb, 0xcd, 0xbb, 0xaf, 0x83, 0x61, 0x21, 0xe3, 0xad, 0x0f, 0x4e, 0xce, 0xe9, 0xc9, 0xeb,
	0x45, 0x11, 0xae, 0xc0, 0x2d, 0xbc, 0x75, 0x9c, 0xd6, 0x5d, 0x2d, 0x22, 0xe5, 0x06, 0x8a, 0xa5,
	0x3b, 0x13, 0xed, 0x05, 0xfd, 0x0d, 0xd4, 0x05, 0x5b, 0x6e, 0xf7, 0xdf, 0xe3, 0xf6, 0xfe, 0xfb,
	0xdc, 0x3e, 0xe8, 0xe6, 0xf9, 0x2b, 0xf6, 0x4a, 0xd9, 0x1b, 0x67, 0x74, 0x49, 0xe3, 0x8f, 0xb2,
	0x0e, 0x4a, 0xb9, 0x0a, 0xf3, 0xc5, 0x2e, 0x0a, 0x2f, 0x21, 0xab, 0xca, 0x2c, 0x87, 0x2b, 0xbc,
	0x3f, 0xc2, 0xa8, 0xd1, 0x20, 0x38, 0xc4, 0xd0, 0x66, 0xed, 0xaa, 0x94, 0x61, 0xca, 0x68, 0x00,
	0x3c, 0xd0, 0xca, 0x72, 0x90, 0xda, 0xa4, 0xa1, 0x95, 0xff, 0x79, 0x01, 0xa5, 0x04, 0x1a, 0x2a,
	0x0f, 0xf9, 0x6e, 0x12, 0x2b, 0x4e, 0xd9, 0xa9, 0x30, 0xc2, 0x0f, 0xc4, 0x7e, 0xb8, 0x6c, 0x43,
	0xd4, 0x8e, 0x4b, 0xac, 0xab, 0xdb, 0xe5, 0x6f, 0x17, 0x0a, 0xfb, 0x6e, 0x0a, 0xcb, 0x30, 0x6b,
	0xee, 0x60, 0x46, 0xff, 0x88, 0x58, 0x34, 0x3d, 0x9d, 0x52, 0x9d, 0x03, 0xa4, 0x74, 0x5f, 0xf6,
	0x28, 0x34, 0xb5, 0xfd, 0x03, 0x9e, 0xa6, 0x21, 0x3b, 0xca, 0x4c, 0x25, 0x2f, 0xc3, 0x40, 0xe9,
	0x93, 0xd6, 0x86, 0x70, 0x4d, 0x65, 0x6f, 0x00, 0xd2, 0x70, 0x9c, 0x82, 0x45, 0x17, 0x11, 0x98,
	0x07, 0x30, 0x37, 0x34, 0xb5, 0xf8, 0x83, 0xd4, 0x86, 0x9a, 0x66, 0x71, 0xd8, 0x6a, 0x16, 0xb8,
	0x9e, 0xf7, 0x3a, 0x64, 0x28, 0x58, 0x74, 0x32, 0x6d, 0xd6, 0xd4, 0x35, 0xa3, 0x8d, 0x76, 0x30,
	0xf4, 0xe3, 0x01, 0x8c, 0xc5, 0x29, 0xf1, 0xc8, 0xfb, 0x11, 0x4c, 0x5c, 0xf5, 0x01, 0xcc, 0x95,
	0x5a, 0x52, 0x0a, 0xfa, 0x3c, 0x58, 0x6b, 0x5c, 0xf8, 0x88, 0x07, 0x5c, 0x2c, 0x71, 0xdc, 0x56,
	0x65, 0x09, 0x06, 0x5d, 0x79, 0xeb, 0xb7, 0xf3, 0x8c, 0xb6, 0xb3, 0x09, 0x53, 0x86, 0x11, 0xba,
	0xd0, 0x0b, 0x3c, 0x1c, 0xcf, 0x69, 0x99, 0x36, 0x84, 0x75, 0xb5, 0xb0, 0xf0, 0x8d, 0x72, 0xf3,
	0xc5, 0x2c, 0x8c, 0x2d, 0x0d, 0x30, 0xfa, 0xe7, 0x13, 0x16, 0xf1, 0x4a, 0xa2, 0x7f, 0x7f, 0x32,
	0x95, 0xbc, 0x50, 0x65, 0x4a, 0x43, 0x92, 0x6f, 0x16, 0x1d, 0x0c, 0x35, 0x59, 0x5b, 0xe3, 0xb3,
	0xdb, 0xc1, 0xb0, 0x2e, 0xb0, 0x05, 0x51, 0x5f, 0x8a, 0x28, 0xb0, 0xb5, 0x8d, 0x3b, 0x69, 0x9a,
	0xd6, 0x1e, 0x91, 0x0d, 0x80, 0x17, 0x05, 0xbe, 0x0d, 0x47, 0xea, 0xef, 0x6b, 0x55, 0x9f, 0x54,
	0xdb, 0x04, 0xaa, 0xb3, 0x2d, 0xf5, 0xbe, 0x57, 0x67, 0xbb, 0xd4, 0xeb, 0x25, 0x2e, 0x72, 0x21,
	0xef, 0x49, 0x7d, 0xd0, 0x5d, 0xbb, 0x26, 0xda, 0x6b, 0x37, 0xea, 0xc3, 0xee, 0xda, 0x8d, 0x7a,
	0xcc, 0x9e, 0x17, 0x02, 0x03, 0xee, 0xbe, 0x33, 0x95, 0x6c, 0x4d, 0x98, 0x9b, 0xf0, 0xe8, 0xaf,
	0x6c, 0x0f, 0x3f, 0xfa, 0xc2, 0x40, 0x8d, 0x3f, 0x9b, 0x9b, 0xa8, 0x85, 0xf8, 0x9a, 0x7e, 0x67,
	0xc2, 0x87, 0x01, 0xd5, 0x34, 0x5a, 0xf8, 0x9c, 0xaf, 0xc6, 0x1b, 0x27, 0xea, 0x41, 0xbc, 0x41,
	0x1a, 0xfe, 0xea, 0xf7, 0x93, 0x09, 0x1d, 0xa5, 0x9a, 0x47, 0x64, 0xf4, 0x47, 0x36, 0xa8, 0xbf,
	0xfd, 0x5a, 0x85, 0xde, 0xeb, 0x14, 0x3a, 0x7e, 0x11, 0x99, 0x55, 0xe5, 0x34, 0x1d, 0x54, 0x9f,
	0xe2, 0x16, 0xd2, 0xb9, 0xe0, 0xa2, 0xee, 0x05, 0x37, 0xfa, 0x92, 0xf5, 0xe9, 0xfb, 0x13, 0x17,
	0xc9, 0xa9, 0xef, 0xb5, 0x6a, 0xa9, 0x85, 0x5c, 0xbc, 0xf9, 0xc3, 0x28, 0x53, 0x2e, 0x17, 0xb3,
	0x93, 0xb3, 0xb3, 0x13, 0x59, 0xfe, 0x5c, 0x62, 0x2f, 0x3a, 0x3b, 0xab, 0xff, 0xd2, 0x87, 0xea,
	0x6c, 0x9f, 0xfe, 0x31, 0x71, 0xf6, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x6c, 0x88, 0x61,
	0xa8, 0x10, 0x00, 0x00,
}
