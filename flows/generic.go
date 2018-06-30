package flows

import (
	"errors"
	"github.com/golang/protobuf/proto"
	actions_proto "www.velocidex.com/golang/velociraptor/actions/proto"
	config "www.velocidex.com/golang/velociraptor/config"
	crypto_proto "www.velocidex.com/golang/velociraptor/crypto/proto"
	"www.velocidex.com/golang/velociraptor/datastore"
	utils "www.velocidex.com/golang/velociraptor/testing"
)

const (
	_                          = iota
	processVQLResponses uint64 = iota
)

type VQLCollector struct{}

func (self *VQLCollector) Start(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject,
	args proto.Message) (*string, error) {
	vql_collector_args, ok := args.(*actions_proto.VQLCollectorArgs)
	if !ok {
		return nil, errors.New("Expected args of type VQLCollectorArgs")
	}

	db, err := datastore.GetDB(config_obj)
	if err != nil {
		return nil, err
	}

	flow_id := GetNewFlowIdForClient(flow_obj.RunnerArgs.ClientId)
	err = db.QueueMessageForClient(
		config_obj, flow_obj.RunnerArgs.ClientId,
		flow_id,
		"VQLClientAction",
		vql_collector_args, processVQLResponses)
	if err != nil {
		return nil, err
	}

	return &flow_id, nil
}

func (self *VQLCollector) ProcessMessage(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject,
	message *crypto_proto.GrrMessage) error {

	switch message.RequestId {
	case processVQLResponses:
		err := flow_obj.FailIfError(message)
		if err != nil {
			return err
		}

		if flow_obj.IsRequestComplete(message) {
			flow_obj.Complete()
			return nil
		}

		utils.Debug(message)
		err = StoreResultInFlow(config_obj, flow_obj, message)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	impl := VQLCollector{}
	RegisterImplementation("VQLCollector", &impl)
}