package zanzi

import (
    "fmt"
    "os/user"
    "strings"

	rcdb "github.com/sourcenetwork/raccoondb"
        _ "google.golang.org/grpc"
        "go.uber.org/zap"

	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/types"
	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/relation_graph"
	"github.com/sourcenetwork/zanzi/internal/store"
	"github.com/sourcenetwork/zanzi/internal/store/kv_store"
)


const (
        dataFile string = "zanzi.db"
        dataDir string = "zanzi"
)

type option func (*Zanzi) error

// New builds a new instance of Zanzi with the given options.
//
// If opt is nil, initializes Zanzi with an INFO level logger
// and stores its data under DefaultDataDir.
func New(opt ...option) (Zanzi, error)  {
    zanzi := Zanzi{}
    for _, o := range opt {
        err := o(&zanzi)
        if err != nil{
            return Zanzi{}, fmt.Errorf("could not build Zanzi: %w", err)
        }
    }
    zanzi.setDefaults()
    zanzi.init()
    return zanzi, nil
}

// setDefaults initializes required attributes if not supplied by the user.
func (z *Zanzi) setDefaults() {
    if z.logger == nil {
       WithDefaultLogger()(z)
    }
    if z.store == nil {
       WithMemKVStore()(z)
    }
}

// WithDevelopmentConfig configures zanzi to use a local in memory kv store and to output DEBUG level logging.
func WithDevelopmentConfig() option {
    return func(z *Zanzi) error {
        err := WithMemKVStore()(z)
        if err != nil {
            return err
        }
        return WithDevelopmentLogger()(z)
    }
}

// WithDefaultKVStore configures Zanzi to use its default KV Store (LevelDB). Data will be stored in dataDir
func WithDefaultKVStore(path string) option { 
    return func(z *Zanzi) error {
        if strings.HasPrefix(path, "~") {
            usr, err := user.Current()
            if err != nil {
                return fmt.Errorf("error identifying user: %w", err)
            }
            home := usr.HomeDir

            path = strings.Replace(path, "~", home, 1)
        }

        kv, err := rcdb.NewLevelDB(path, dataFile)
        if err != nil {
            return fmt.Errorf("error initializing kv store: %v", err)
        }
        return WithKVStore(kv)(z)
    }
}


// WithLogger sets the logger which will be used by Zanzi.
// The Logger type is a zanzi defined logging facade.
// This approach gives clients the freedom to choose the 
// logging level and logging implementation for Zanzi.
func WithLogger(logger types.Logger) option {
    return func(z *Zanzi) error {
        z.logger = logger
        return nil
    }
}

// WithDefaultLogger configures zanzi to use its default logger implementation (Zap).
// Logging level is set to INFO.
func WithDefaultLogger() option {
    return func(z *Zanzi) error {
        z.logger = zap.S()
        return nil
    }
}

// WithDevelopmentLogger configures zanzi to use Zap as its logger (writes to stderr).
// Logging level is set to DEBUG.
func WithDevelopmentLogger() option {
    return func(z *Zanzi) error {
        zapLogger, err := zap.NewDevelopment()
        if err != nil {
            return err
        }
        z.logger = zapLogger.Sugar()
        return nil
    }
}

// WithoutLogger configures zanzi to not produce any logs.
func WithoutLogger(logger types.Logger) option {
    return func(z *Zanzi) error {
        z.logger = &types.NoopLogger{}
        return nil
    }
}

// WithKVStore configures zanzi to use the given kv store as its
// data store
func WithKVStore(kv rcdb.KVStore) option {
    return func(z *Zanzi) error {
        kvStore, err := kv_store.NewKVStore(kv)
        if err != nil {
            return err
        }
        z.store = &kvStore
        return nil
    }
}

// WithMemKVStore configures zanzi to use an in memory kv store
func WithMemKVStore() option {
    kv := rcdb.NewMemKV()
    return WithKVStore(kv)
}

// Zanzi is a container type responsible for initializing the zanzi's internal types and any additional requirements.
// Clients should use the new function to use zanzi.
type Zanzi struct {
    polService api.PolicyServiceServer
    relGraphService api.RelationGraphServer
    logger types.Logger
    store store.Store
}

func (z *Zanzi) init() {
    z.polService = policy.NewService(z.store.GetPolicyRepository())
    z.relGraphService = relation_graph.NewService(
        z.store.GetRelationNodeRepository(),
        z.store.GetPolicyRepository(),
        z.logger,
    )
}

// GetRelationshipGraphService returns an implementation of the
// GRPC defined RelationGraph service type.
func (z *Zanzi) GetRelationGraphService() api.RelationGraphServer {
    return z.relGraphService
}

// GetPolicyService returns an implementation of the
// GRPC defined RelationGraph service type.
func (z *Zanzi) GetPolicyService() api.PolicyServiceServer {
    return z.polService
}

func (z *Zanzi) GetLogger() types.Logger {
    return z.logger
}
