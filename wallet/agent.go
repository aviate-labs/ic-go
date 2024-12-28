// Package wallet provides a client for the "wallet" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package wallet

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

type AddressEntry struct {
	Id   principal.Principal `ic:"id" json:"id"`
	Name *string             `ic:"name,omitempty" json:"name,omitempty"`
	Kind Kind                `ic:"kind" json:"kind"`
	Role Role                `ic:"role" json:"role"`
}

// Agent is a client for the "wallet" canister.
type Agent struct {
	*agent.Agent
	CanisterId principal.Principal
}

// NewAgent creates a new agent for the "wallet" canister.
func NewAgent(canisterId principal.Principal, config agent.Config) (*Agent, error) {
	a, err := agent.New(config)
	if err != nil {
		return nil, err
	}
	return &Agent{
		Agent:      a,
		CanisterId: canisterId,
	}, nil
}

// AddAddress calls the "add_address" method on the "wallet" canister.
func (a Agent) AddAddress(address AddressEntry) error {
	if err := a.Agent.Call(
		a.CanisterId,
		"add_address",
		[]any{address},
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// AddController calls the "add_controller" method on the "wallet" canister.
func (a Agent) AddController(arg0 principal.Principal) error {
	if err := a.Agent.Call(
		a.CanisterId,
		"add_controller",
		[]any{arg0},
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// Authorize calls the "authorize" method on the "wallet" canister.
func (a Agent) Authorize(arg0 principal.Principal) error {
	if err := a.Agent.Call(
		a.CanisterId,
		"authorize",
		[]any{arg0},
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// Deauthorize calls the "deauthorize" method on the "wallet" canister.
func (a Agent) Deauthorize(arg0 principal.Principal) (*WalletResult, error) {
	var r0 WalletResult
	if err := a.Agent.Call(
		a.CanisterId,
		"deauthorize",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetChart calls the "get_chart" method on the "wallet" canister.
func (a Agent) GetChart(arg0 *struct {
	Count     *uint32 `ic:"count,omitempty" json:"count,omitempty"`
	Precision *uint64 `ic:"precision,omitempty" json:"precision,omitempty"`
}) (*[]struct {
	Field0 uint64 `ic:"0" json:"0"`
	Field1 uint64 `ic:"1" json:"1"`
}, error) {
	var r0 []struct {
		Field0 uint64 `ic:"0" json:"0"`
		Field1 uint64 `ic:"1" json:"1"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"get_chart",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetControllers calls the "get_controllers" method on the "wallet" canister.
func (a Agent) GetControllers() (*[]principal.Principal, error) {
	var r0 []principal.Principal
	if err := a.Agent.Query(
		a.CanisterId,
		"get_controllers",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetCustodians calls the "get_custodians" method on the "wallet" canister.
func (a Agent) GetCustodians() (*[]principal.Principal, error) {
	var r0 []principal.Principal
	if err := a.Agent.Query(
		a.CanisterId,
		"get_custodians",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetEvents calls the "get_events" method on the "wallet" canister.
func (a Agent) GetEvents(arg0 *struct {
	From *uint32 `ic:"from,omitempty" json:"from,omitempty"`
	To   *uint32 `ic:"to,omitempty" json:"to,omitempty"`
}) (*[]Event, error) {
	var r0 []Event
	if err := a.Agent.Query(
		a.CanisterId,
		"get_events",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetEvents128 calls the "get_events128" method on the "wallet" canister.
func (a Agent) GetEvents128(arg0 *struct {
	From *uint32 `ic:"from,omitempty" json:"from,omitempty"`
	To   *uint32 `ic:"to,omitempty" json:"to,omitempty"`
}) (*[]Event128, error) {
	var r0 []Event128
	if err := a.Agent.Query(
		a.CanisterId,
		"get_events128",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetManagedCanisterEvents calls the "get_managed_canister_events" method on the "wallet" canister.
func (a Agent) GetManagedCanisterEvents(arg0 struct {
	Canister principal.Principal `ic:"canister" json:"canister"`
	From     *uint32             `ic:"from,omitempty" json:"from,omitempty"`
	To       *uint32             `ic:"to,omitempty" json:"to,omitempty"`
}) (**[]ManagedCanisterEvent, error) {
	var r0 *[]ManagedCanisterEvent
	if err := a.Agent.Query(
		a.CanisterId,
		"get_managed_canister_events",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetManagedCanisterEvents128 calls the "get_managed_canister_events128" method on the "wallet" canister.
func (a Agent) GetManagedCanisterEvents128(arg0 struct {
	Canister principal.Principal `ic:"canister" json:"canister"`
	From     *uint32             `ic:"from,omitempty" json:"from,omitempty"`
	To       *uint32             `ic:"to,omitempty" json:"to,omitempty"`
}) (**[]ManagedCanisterEvent128, error) {
	var r0 *[]ManagedCanisterEvent128
	if err := a.Agent.Query(
		a.CanisterId,
		"get_managed_canister_events128",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// HttpRequest calls the "http_request" method on the "wallet" canister.
func (a Agent) HttpRequest(request HttpRequest) (*HttpResponse, error) {
	var r0 HttpResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"http_request",
		[]any{request},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListAddresses calls the "list_addresses" method on the "wallet" canister.
func (a Agent) ListAddresses() (*[]AddressEntry, error) {
	var r0 []AddressEntry
	if err := a.Agent.Query(
		a.CanisterId,
		"list_addresses",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListManagedCanisters calls the "list_managed_canisters" method on the "wallet" canister.
func (a Agent) ListManagedCanisters(arg0 struct {
	From *uint32 `ic:"from,omitempty" json:"from,omitempty"`
	To   *uint32 `ic:"to,omitempty" json:"to,omitempty"`
}) (*[]ManagedCanisterInfo, *uint32, error) {
	var r0 []ManagedCanisterInfo
	var r1 uint32
	if err := a.Agent.Query(
		a.CanisterId,
		"list_managed_canisters",
		[]any{arg0},
		[]any{&r0, &r1},
	); err != nil {
		return nil, nil, err
	}
	return &r0, &r1, nil
}

// Name calls the "name" method on the "wallet" canister.
func (a Agent) Name() (**string, error) {
	var r0 *string
	if err := a.Agent.Query(
		a.CanisterId,
		"name",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// RemoveAddress calls the "remove_address" method on the "wallet" canister.
func (a Agent) RemoveAddress(address principal.Principal) (*WalletResult, error) {
	var r0 WalletResult
	if err := a.Agent.Call(
		a.CanisterId,
		"remove_address",
		[]any{address},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// RemoveController calls the "remove_controller" method on the "wallet" canister.
func (a Agent) RemoveController(arg0 principal.Principal) (*WalletResult, error) {
	var r0 WalletResult
	if err := a.Agent.Call(
		a.CanisterId,
		"remove_controller",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// SetName calls the "set_name" method on the "wallet" canister.
func (a Agent) SetName(arg0 string) error {
	if err := a.Agent.Call(
		a.CanisterId,
		"set_name",
		[]any{arg0},
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// SetShortName calls the "set_short_name" method on the "wallet" canister.
func (a Agent) SetShortName(arg0 principal.Principal, arg1 *string) (**ManagedCanisterInfo, error) {
	var r0 *ManagedCanisterInfo
	if err := a.Agent.Call(
		a.CanisterId,
		"set_short_name",
		[]any{arg0, arg1},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletApiVersion calls the "wallet_api_version" method on the "wallet" canister.
func (a Agent) WalletApiVersion() (*string, error) {
	var r0 string
	if err := a.Agent.Query(
		a.CanisterId,
		"wallet_api_version",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletBalance calls the "wallet_balance" method on the "wallet" canister.
func (a Agent) WalletBalance() (*struct {
	Amount uint64 `ic:"amount" json:"amount"`
}, error) {
	var r0 struct {
		Amount uint64 `ic:"amount" json:"amount"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"wallet_balance",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletBalance128 calls the "wallet_balance128" method on the "wallet" canister.
func (a Agent) WalletBalance128() (*struct {
	Amount idl.Nat `ic:"amount" json:"amount"`
}, error) {
	var r0 struct {
		Amount idl.Nat `ic:"amount" json:"amount"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"wallet_balance128",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletCall calls the "wallet_call" method on the "wallet" canister.
func (a Agent) WalletCall(arg0 struct {
	Canister   principal.Principal `ic:"canister" json:"canister"`
	MethodName string              `ic:"method_name" json:"method_name"`
	Args       []byte              `ic:"args" json:"args"`
	Cycles     uint64              `ic:"cycles" json:"cycles"`
}) (*WalletResultCall, error) {
	var r0 WalletResultCall
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_call",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletCall128 calls the "wallet_call128" method on the "wallet" canister.
func (a Agent) WalletCall128(arg0 struct {
	Canister   principal.Principal `ic:"canister" json:"canister"`
	MethodName string              `ic:"method_name" json:"method_name"`
	Args       []byte              `ic:"args" json:"args"`
	Cycles     idl.Nat             `ic:"cycles" json:"cycles"`
}) (*WalletResultCall, error) {
	var r0 WalletResultCall
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_call128",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletCallWithMaxCycles calls the "wallet_call_with_max_cycles" method on the "wallet" canister.
func (a Agent) WalletCallWithMaxCycles(arg0 struct {
	Canister   principal.Principal `ic:"canister" json:"canister"`
	MethodName string              `ic:"method_name" json:"method_name"`
	Args       []byte              `ic:"args" json:"args"`
}) (*WalletResultCallWithMaxCycles, error) {
	var r0 WalletResultCallWithMaxCycles
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_call_with_max_cycles",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletCreateCanister calls the "wallet_create_canister" method on the "wallet" canister.
func (a Agent) WalletCreateCanister(arg0 CreateCanisterArgs) (*WalletResultCreate, error) {
	var r0 WalletResultCreate
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_create_canister",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletCreateCanister128 calls the "wallet_create_canister128" method on the "wallet" canister.
func (a Agent) WalletCreateCanister128(arg0 CreateCanisterArgs128) (*WalletResultCreate, error) {
	var r0 WalletResultCreate
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_create_canister128",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletCreateWallet calls the "wallet_create_wallet" method on the "wallet" canister.
func (a Agent) WalletCreateWallet(arg0 CreateCanisterArgs) (*WalletResultCreate, error) {
	var r0 WalletResultCreate
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_create_wallet",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletCreateWallet128 calls the "wallet_create_wallet128" method on the "wallet" canister.
func (a Agent) WalletCreateWallet128(arg0 CreateCanisterArgs128) (*WalletResultCreate, error) {
	var r0 WalletResultCreate
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_create_wallet128",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletReceive calls the "wallet_receive" method on the "wallet" canister.
func (a Agent) WalletReceive(arg0 *ReceiveOptions) error {
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_receive",
		[]any{arg0},
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// WalletSend calls the "wallet_send" method on the "wallet" canister.
func (a Agent) WalletSend(arg0 struct {
	Canister principal.Principal `ic:"canister" json:"canister"`
	Amount   uint64              `ic:"amount" json:"amount"`
}) (*WalletResult, error) {
	var r0 WalletResult
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_send",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletSend128 calls the "wallet_send128" method on the "wallet" canister.
func (a Agent) WalletSend128(arg0 struct {
	Canister principal.Principal `ic:"canister" json:"canister"`
	Amount   idl.Nat             `ic:"amount" json:"amount"`
}) (*WalletResult, error) {
	var r0 WalletResult
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_send128",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// WalletStoreWalletWasm calls the "wallet_store_wallet_wasm" method on the "wallet" canister.
func (a Agent) WalletStoreWalletWasm(arg0 struct {
	WasmModule []byte `ic:"wasm_module" json:"wasm_module"`
}) error {
	if err := a.Agent.Call(
		a.CanisterId,
		"wallet_store_wallet_wasm",
		[]any{arg0},
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

type CanisterSettings struct {
	Controller        *principal.Principal   `ic:"controller,omitempty" json:"controller,omitempty"`
	Controllers       *[]principal.Principal `ic:"controllers,omitempty" json:"controllers,omitempty"`
	ComputeAllocation *idl.Nat               `ic:"compute_allocation,omitempty" json:"compute_allocation,omitempty"`
	MemoryAllocation  *idl.Nat               `ic:"memory_allocation,omitempty" json:"memory_allocation,omitempty"`
	FreezingThreshold *idl.Nat               `ic:"freezing_threshold,omitempty" json:"freezing_threshold,omitempty"`
}

type CreateCanisterArgs struct {
	Cycles   uint64           `ic:"cycles" json:"cycles"`
	Settings CanisterSettings `ic:"settings" json:"settings"`
}

type CreateCanisterArgs128 struct {
	Cycles   idl.Nat          `ic:"cycles" json:"cycles"`
	Settings CanisterSettings `ic:"settings" json:"settings"`
}

type Event struct {
	Id        uint32    `ic:"id" json:"id"`
	Timestamp uint64    `ic:"timestamp" json:"timestamp"`
	Kind      EventKind `ic:"kind" json:"kind"`
}

type Event128 struct {
	Id        uint32       `ic:"id" json:"id"`
	Timestamp uint64       `ic:"timestamp" json:"timestamp"`
	Kind      EventKind128 `ic:"kind" json:"kind"`
}

type EventKind struct {
	CyclesSent *struct {
		To     principal.Principal `ic:"to" json:"to"`
		Amount uint64              `ic:"amount" json:"amount"`
		Refund uint64              `ic:"refund" json:"refund"`
	} `ic:"CyclesSent,variant"`
	CyclesReceived *struct {
		From   principal.Principal `ic:"from" json:"from"`
		Amount uint64              `ic:"amount" json:"amount"`
		Memo   *string             `ic:"memo,omitempty" json:"memo,omitempty"`
	} `ic:"CyclesReceived,variant"`
	AddressAdded *struct {
		Id   principal.Principal `ic:"id" json:"id"`
		Name *string             `ic:"name,omitempty" json:"name,omitempty"`
		Role Role                `ic:"role" json:"role"`
	} `ic:"AddressAdded,variant"`
	AddressRemoved *struct {
		Id principal.Principal `ic:"id" json:"id"`
	} `ic:"AddressRemoved,variant"`
	CanisterCreated *struct {
		Canister principal.Principal `ic:"canister" json:"canister"`
		Cycles   uint64              `ic:"cycles" json:"cycles"`
	} `ic:"CanisterCreated,variant"`
	CanisterCalled *struct {
		Canister   principal.Principal `ic:"canister" json:"canister"`
		MethodName string              `ic:"method_name" json:"method_name"`
		Cycles     uint64              `ic:"cycles" json:"cycles"`
	} `ic:"CanisterCalled,variant"`
	WalletDeployed *struct {
		Canister principal.Principal `ic:"canister" json:"canister"`
	} `ic:"WalletDeployed,variant"`
}

type EventKind128 struct {
	CyclesSent *struct {
		To     principal.Principal `ic:"to" json:"to"`
		Amount idl.Nat             `ic:"amount" json:"amount"`
		Refund idl.Nat             `ic:"refund" json:"refund"`
	} `ic:"CyclesSent,variant"`
	CyclesReceived *struct {
		From   principal.Principal `ic:"from" json:"from"`
		Amount idl.Nat             `ic:"amount" json:"amount"`
		Memo   *string             `ic:"memo,omitempty" json:"memo,omitempty"`
	} `ic:"CyclesReceived,variant"`
	AddressAdded *struct {
		Id   principal.Principal `ic:"id" json:"id"`
		Name *string             `ic:"name,omitempty" json:"name,omitempty"`
		Role Role                `ic:"role" json:"role"`
	} `ic:"AddressAdded,variant"`
	AddressRemoved *struct {
		Id principal.Principal `ic:"id" json:"id"`
	} `ic:"AddressRemoved,variant"`
	CanisterCreated *struct {
		Canister principal.Principal `ic:"canister" json:"canister"`
		Cycles   idl.Nat             `ic:"cycles" json:"cycles"`
	} `ic:"CanisterCreated,variant"`
	CanisterCalled *struct {
		Canister   principal.Principal `ic:"canister" json:"canister"`
		MethodName string              `ic:"method_name" json:"method_name"`
		Cycles     idl.Nat             `ic:"cycles" json:"cycles"`
	} `ic:"CanisterCalled,variant"`
	WalletDeployed *struct {
		Canister principal.Principal `ic:"canister" json:"canister"`
	} `ic:"WalletDeployed,variant"`
}

type HeaderField struct {
	Field0 string `ic:"0" json:"0"`
	Field1 string `ic:"1" json:"1"`
}

type HttpRequest struct {
	Method  string        `ic:"method" json:"method"`
	Url     string        `ic:"url" json:"url"`
	Headers []HeaderField `ic:"headers" json:"headers"`
	Body    []byte        `ic:"body" json:"body"`
}

type HttpResponse struct {
	StatusCode        uint16             `ic:"status_code" json:"status_code"`
	Headers           []HeaderField      `ic:"headers" json:"headers"`
	Body              []byte             `ic:"body" json:"body"`
	StreamingStrategy *StreamingStrategy `ic:"streaming_strategy,omitempty" json:"streaming_strategy,omitempty"`
}

type Kind struct {
	Unknown  *idl.Null `ic:"Unknown,variant"`
	User     *idl.Null `ic:"User,variant"`
	Canister *idl.Null `ic:"Canister,variant"`
}

type ManagedCanisterEvent struct {
	Id        uint32                   `ic:"id" json:"id"`
	Timestamp uint64                   `ic:"timestamp" json:"timestamp"`
	Kind      ManagedCanisterEventKind `ic:"kind" json:"kind"`
}

type ManagedCanisterEvent128 struct {
	Id        uint32                      `ic:"id" json:"id"`
	Timestamp uint64                      `ic:"timestamp" json:"timestamp"`
	Kind      ManagedCanisterEventKind128 `ic:"kind" json:"kind"`
}

type ManagedCanisterEventKind struct {
	CyclesSent *struct {
		Amount uint64 `ic:"amount" json:"amount"`
		Refund uint64 `ic:"refund" json:"refund"`
	} `ic:"CyclesSent,variant"`
	Called *struct {
		MethodName string `ic:"method_name" json:"method_name"`
		Cycles     uint64 `ic:"cycles" json:"cycles"`
	} `ic:"Called,variant"`
	Created *struct {
		Cycles uint64 `ic:"cycles" json:"cycles"`
	} `ic:"Created,variant"`
}

type ManagedCanisterEventKind128 struct {
	CyclesSent *struct {
		Amount idl.Nat `ic:"amount" json:"amount"`
		Refund idl.Nat `ic:"refund" json:"refund"`
	} `ic:"CyclesSent,variant"`
	Called *struct {
		MethodName string  `ic:"method_name" json:"method_name"`
		Cycles     idl.Nat `ic:"cycles" json:"cycles"`
	} `ic:"Called,variant"`
	Created *struct {
		Cycles idl.Nat `ic:"cycles" json:"cycles"`
	} `ic:"Created,variant"`
}

type ManagedCanisterInfo struct {
	Id        principal.Principal `ic:"id" json:"id"`
	Name      *string             `ic:"name,omitempty" json:"name,omitempty"`
	CreatedAt uint64              `ic:"created_at" json:"created_at"`
}

type ReceiveOptions struct {
	Memo *string `ic:"memo,omitempty" json:"memo,omitempty"`
}

type Role struct {
	Contact    *idl.Null `ic:"Contact,variant"`
	Custodian  *idl.Null `ic:"Custodian,variant"`
	Controller *idl.Null `ic:"Controller,variant"`
}

type StreamingCallbackHttpResponse struct {
	Body  []byte `ic:"body" json:"body"`
	Token *Token `ic:"token,omitempty" json:"token,omitempty"`
}

type StreamingStrategy struct {
	Callback *struct {
		Callback struct { /* NOT SUPPORTED */
		} `ic:"callback" json:"callback"`
		Token Token `ic:"token" json:"token"`
	} `ic:"Callback,variant"`
}

type Token struct {
}

type WalletResult struct {
	Ok  *idl.Null `ic:"Ok,variant"`
	Err *string   `ic:"Err,variant"`
}

type WalletResultCall struct {
	Ok *struct {
		Return []byte `ic:"return" json:"return"`
	} `ic:"Ok,variant"`
	Err *string `ic:"Err,variant"`
}

type WalletResultCallWithMaxCycles struct {
	Ok *struct {
		Return         []byte  `ic:"return" json:"return"`
		AttachedCycles idl.Nat `ic:"attached_cycles" json:"attached_cycles"`
	} `ic:"Ok,variant"`
	Err *string `ic:"Err,variant"`
}

type WalletResultCreate struct {
	Ok *struct {
		CanisterId principal.Principal `ic:"canister_id" json:"canister_id"`
	} `ic:"Ok,variant"`
	Err *string `ic:"Err,variant"`
}
