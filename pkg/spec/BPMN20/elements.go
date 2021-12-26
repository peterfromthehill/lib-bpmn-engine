package BPMN20

const (
	StartEventType             string = "StartEvent"
	EndEventType               string = "EndEvent"
	ServiceTaskType            string = "ServiceTask"
	ParallelGatewayType        string = "ParallelGateway"
	ExclusiveGatewayType       string = "ExclusiveGateway"
	IntermediateCatchEventType string = "IntermediateCatchEvent"
)

type BaseElement interface {
	GetId() string
	GetName() string
	GetIncoming() []string
	GetOutgoing() []string

	GetTypeName() string
}

func (startEvent TStartEvent) GetId() string {
	return startEvent.Id
}

func (startEvent TStartEvent) GetName() string {
	return startEvent.Name
}

func (startEvent TStartEvent) GetIncoming() []string {
	return startEvent.IncomingAssociation
}

func (startEvent TStartEvent) GetOutgoing() []string {
	return startEvent.OutgoingAssociation
}

func (startEvent TStartEvent) GetTypeName() string {
	return StartEventType
}

func (endEvent TEndEvent) GetId() string {
	return endEvent.Id
}

func (endEvent TEndEvent) GetName() string {
	return endEvent.Name
}

func (endEvent TEndEvent) GetIncoming() []string {
	return endEvent.IncomingAssociation
}

func (endEvent TEndEvent) GetOutgoing() []string {
	return endEvent.OutgoingAssociation
}

func (endEvent TEndEvent) GetTypeName() string {
	return EndEventType
}

func (serviceTask TServiceTask) GetId() string {
	return serviceTask.Id
}

func (serviceTask TServiceTask) GetName() string {
	return serviceTask.Name
}

func (serviceTask TServiceTask) GetIncoming() []string {
	return serviceTask.IncomingAssociation
}

func (serviceTask TServiceTask) GetOutgoing() []string {
	return serviceTask.OutgoingAssociation
}

func (serviceTask TServiceTask) GetTypeName() string {
	return ServiceTaskType
}

func (parallelGateway TParallelGateway) GetId() string {
	return parallelGateway.Id
}

func (parallelGateway TParallelGateway) GetName() string {
	return parallelGateway.Name
}

func (parallelGateway TParallelGateway) GetIncoming() []string {
	return parallelGateway.IncomingAssociation
}

func (parallelGateway TParallelGateway) GetOutgoing() []string {
	return parallelGateway.OutgoingAssociation
}

func (parallelGateway TParallelGateway) GetTypeName() string {
	return ParallelGatewayType
}

func (exclusiveGateway TExclusiveGateway) GetId() string {
	return exclusiveGateway.Id
}

func (exclusiveGateway TExclusiveGateway) GetName() string {
	return exclusiveGateway.Name
}

func (exclusiveGateway TExclusiveGateway) GetIncoming() []string {
	return exclusiveGateway.IncomingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetOutgoing() []string {
	return exclusiveGateway.OutgoingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetTypeName() string {
	return ExclusiveGatewayType
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetId() string {
	return intermediateCatchEvent.Id
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetName() string {
	return intermediateCatchEvent.Name
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetIncoming() []string {
	return intermediateCatchEvent.IncomingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetOutgoing() []string {
	return intermediateCatchEvent.OutgoingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetTypeName() string {
	return IntermediateCatchEventType
}

func FindTargetRefs(sequenceFlows []TSequenceFlow, withId func(string) bool) (ret []string) {
	for _, flow := range sequenceFlows {
		if withId(flow.Id) {
			ret = append(ret, flow.TargetRef)
		}
	}
	return
}