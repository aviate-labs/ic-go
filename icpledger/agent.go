// Package icpledger provides a client for the "icpledger" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package icpledger

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

type Account struct {
	Owner      principal.Principal `ic:"owner" json:"owner"`
	Subaccount *SubAccount         `ic:"subaccount,omitempty" json:"subaccount,omitempty"`
}

type AccountBalanceArgs struct {
	Account AccountIdentifier `ic:"account" json:"account"`
}

type AccountBalanceArgsDfx struct {
	Account TextAccountIdentifier `ic:"account" json:"account"`
}

type AccountIdentifier = []byte

// Agent is a client for the "icpledger" canister.
type Agent struct {
	*agent.Agent
	CanisterId principal.Principal
}

// NewAgent creates a new agent for the "icpledger" canister.
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

// AccountBalance calls the "account_balance" method on the "icpledger" canister.
func (a Agent) AccountBalance(arg0 AccountBalanceArgs) (*Tokens, error) {
	var r0 Tokens
	if err := a.Agent.Query(
		a.CanisterId,
		"account_balance",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// AccountBalanceDfx calls the "account_balance_dfx" method on the "icpledger" canister.
func (a Agent) AccountBalanceDfx(arg0 AccountBalanceArgsDfx) (*Tokens, error) {
	var r0 Tokens
	if err := a.Agent.Query(
		a.CanisterId,
		"account_balance_dfx",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// AccountIdentifier calls the "account_identifier" method on the "icpledger" canister.
func (a Agent) AccountIdentifier(arg0 Account) (*AccountIdentifier, error) {
	var r0 AccountIdentifier
	if err := a.Agent.Query(
		a.CanisterId,
		"account_identifier",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Archives calls the "archives" method on the "icpledger" canister.
func (a Agent) Archives() (*Archives, error) {
	var r0 Archives
	if err := a.Agent.Query(
		a.CanisterId,
		"archives",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Decimals calls the "decimals" method on the "icpledger" canister.
func (a Agent) Decimals() (*struct {
	Decimals uint32 `ic:"decimals" json:"decimals"`
}, error) {
	var r0 struct {
		Decimals uint32 `ic:"decimals" json:"decimals"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"decimals",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc10SupportedStandards calls the "icrc10_supported_standards" method on the "icpledger" canister.
func (a Agent) Icrc10SupportedStandards() (*[]struct {
	Name string `ic:"name" json:"name"`
	Url  string `ic:"url" json:"url"`
}, error) {
	var r0 []struct {
		Name string `ic:"name" json:"name"`
		Url  string `ic:"url" json:"url"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc10_supported_standards",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1BalanceOf calls the "icrc1_balance_of" method on the "icpledger" canister.
func (a Agent) Icrc1BalanceOf(arg0 Account) (*Icrc1Tokens, error) {
	var r0 Icrc1Tokens
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_balance_of",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Decimals calls the "icrc1_decimals" method on the "icpledger" canister.
func (a Agent) Icrc1Decimals() (*uint8, error) {
	var r0 uint8
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_decimals",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Fee calls the "icrc1_fee" method on the "icpledger" canister.
func (a Agent) Icrc1Fee() (*Icrc1Tokens, error) {
	var r0 Icrc1Tokens
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_fee",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Metadata calls the "icrc1_metadata" method on the "icpledger" canister.
func (a Agent) Icrc1Metadata() (*[]struct {
	Field0 string `ic:"0" json:"0"`
	Field1 Value  `ic:"1" json:"1"`
}, error) {
	var r0 []struct {
		Field0 string `ic:"0" json:"0"`
		Field1 Value  `ic:"1" json:"1"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_metadata",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1MintingAccount calls the "icrc1_minting_account" method on the "icpledger" canister.
func (a Agent) Icrc1MintingAccount() (**Account, error) {
	var r0 *Account
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_minting_account",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Name calls the "icrc1_name" method on the "icpledger" canister.
func (a Agent) Icrc1Name() (*string, error) {
	var r0 string
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_name",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1SupportedStandards calls the "icrc1_supported_standards" method on the "icpledger" canister.
func (a Agent) Icrc1SupportedStandards() (*[]struct {
	Name string `ic:"name" json:"name"`
	Url  string `ic:"url" json:"url"`
}, error) {
	var r0 []struct {
		Name string `ic:"name" json:"name"`
		Url  string `ic:"url" json:"url"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_supported_standards",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Symbol calls the "icrc1_symbol" method on the "icpledger" canister.
func (a Agent) Icrc1Symbol() (*string, error) {
	var r0 string
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_symbol",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1TotalSupply calls the "icrc1_total_supply" method on the "icpledger" canister.
func (a Agent) Icrc1TotalSupply() (*Icrc1Tokens, error) {
	var r0 Icrc1Tokens
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc1_total_supply",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Transfer calls the "icrc1_transfer" method on the "icpledger" canister.
func (a Agent) Icrc1Transfer(arg0 TransferArg) (*Icrc1TransferResult, error) {
	var r0 Icrc1TransferResult
	if err := a.Agent.Call(
		a.CanisterId,
		"icrc1_transfer",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc21CanisterCallConsentMessage calls the "icrc21_canister_call_consent_message" method on the "icpledger" canister.
func (a Agent) Icrc21CanisterCallConsentMessage(arg0 Icrc21ConsentMessageRequest) (*Icrc21ConsentMessageResponse, error) {
	var r0 Icrc21ConsentMessageResponse
	if err := a.Agent.Call(
		a.CanisterId,
		"icrc21_canister_call_consent_message",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc2Allowance calls the "icrc2_allowance" method on the "icpledger" canister.
func (a Agent) Icrc2Allowance(arg0 AllowanceArgs) (*Allowance, error) {
	var r0 Allowance
	if err := a.Agent.Query(
		a.CanisterId,
		"icrc2_allowance",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc2Approve calls the "icrc2_approve" method on the "icpledger" canister.
func (a Agent) Icrc2Approve(arg0 ApproveArgs) (*ApproveResult, error) {
	var r0 ApproveResult
	if err := a.Agent.Call(
		a.CanisterId,
		"icrc2_approve",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc2TransferFrom calls the "icrc2_transfer_from" method on the "icpledger" canister.
func (a Agent) Icrc2TransferFrom(arg0 TransferFromArgs) (*TransferFromResult, error) {
	var r0 TransferFromResult
	if err := a.Agent.Call(
		a.CanisterId,
		"icrc2_transfer_from",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Name calls the "name" method on the "icpledger" canister.
func (a Agent) Name() (*struct {
	Name string `ic:"name" json:"name"`
}, error) {
	var r0 struct {
		Name string `ic:"name" json:"name"`
	}
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

// QueryBlocks calls the "query_blocks" method on the "icpledger" canister.
func (a Agent) QueryBlocks(arg0 GetBlocksArgs) (*QueryBlocksResponse, error) {
	var r0 QueryBlocksResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"query_blocks",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// QueryEncodedBlocks calls the "query_encoded_blocks" method on the "icpledger" canister.
func (a Agent) QueryEncodedBlocks(arg0 GetBlocksArgs) (*QueryEncodedBlocksResponse, error) {
	var r0 QueryEncodedBlocksResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"query_encoded_blocks",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// SendDfx calls the "send_dfx" method on the "icpledger" canister.
func (a Agent) SendDfx(arg0 SendArgs) (*BlockIndex, error) {
	var r0 BlockIndex
	if err := a.Agent.Call(
		a.CanisterId,
		"send_dfx",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Symbol calls the "symbol" method on the "icpledger" canister.
func (a Agent) Symbol() (*struct {
	Symbol string `ic:"symbol" json:"symbol"`
}, error) {
	var r0 struct {
		Symbol string `ic:"symbol" json:"symbol"`
	}
	if err := a.Agent.Query(
		a.CanisterId,
		"symbol",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Transfer calls the "transfer" method on the "icpledger" canister.
func (a Agent) Transfer(arg0 TransferArgs) (*TransferResult, error) {
	var r0 TransferResult
	if err := a.Agent.Call(
		a.CanisterId,
		"transfer",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// TransferFee calls the "transfer_fee" method on the "icpledger" canister.
func (a Agent) TransferFee(arg0 TransferFeeArg) (*TransferFee, error) {
	var r0 TransferFee
	if err := a.Agent.Query(
		a.CanisterId,
		"transfer_fee",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type Allowance struct {
	Allowance Icrc1Tokens     `ic:"allowance" json:"allowance"`
	ExpiresAt *Icrc1Timestamp `ic:"expires_at,omitempty" json:"expires_at,omitempty"`
}

type AllowanceArgs struct {
	Account Account `ic:"account" json:"account"`
	Spender Account `ic:"spender" json:"spender"`
}

type ApproveArgs struct {
	FromSubaccount    *SubAccount     `ic:"from_subaccount,omitempty" json:"from_subaccount,omitempty"`
	Spender           Account         `ic:"spender" json:"spender"`
	Amount            Icrc1Tokens     `ic:"amount" json:"amount"`
	ExpectedAllowance *Icrc1Tokens    `ic:"expected_allowance,omitempty" json:"expected_allowance,omitempty"`
	ExpiresAt         *Icrc1Timestamp `ic:"expires_at,omitempty" json:"expires_at,omitempty"`
	Fee               *Icrc1Tokens    `ic:"fee,omitempty" json:"fee,omitempty"`
	Memo              *[]byte         `ic:"memo,omitempty" json:"memo,omitempty"`
	CreatedAtTime     *Icrc1Timestamp `ic:"created_at_time,omitempty" json:"created_at_time,omitempty"`
}

type ApproveError struct {
	BadFee *struct {
		ExpectedFee Icrc1Tokens `ic:"expected_fee" json:"expected_fee"`
	} `ic:"BadFee,variant"`
	InsufficientFunds *struct {
		Balance Icrc1Tokens `ic:"balance" json:"balance"`
	} `ic:"InsufficientFunds,variant"`
	AllowanceChanged *struct {
		CurrentAllowance Icrc1Tokens `ic:"current_allowance" json:"current_allowance"`
	} `ic:"AllowanceChanged,variant"`
	Expired *struct {
		LedgerTime uint64 `ic:"ledger_time" json:"ledger_time"`
	} `ic:"Expired,variant"`
	TooOld          *idl.Null `ic:"TooOld,variant"`
	CreatedInFuture *struct {
		LedgerTime uint64 `ic:"ledger_time" json:"ledger_time"`
	} `ic:"CreatedInFuture,variant"`
	Duplicate *struct {
		DuplicateOf Icrc1BlockIndex `ic:"duplicate_of" json:"duplicate_of"`
	} `ic:"Duplicate,variant"`
	TemporarilyUnavailable *idl.Null `ic:"TemporarilyUnavailable,variant"`
	GenericError           *struct {
		ErrorCode idl.Nat `ic:"error_code" json:"error_code"`
		Message   string  `ic:"message" json:"message"`
	} `ic:"GenericError,variant"`
}

type ApproveResult struct {
	Ok  *Icrc1BlockIndex `ic:"Ok,variant"`
	Err *ApproveError    `ic:"Err,variant"`
}

type Archive struct {
	CanisterId principal.Principal `ic:"canister_id" json:"canister_id"`
}

type ArchiveOptions struct {
	TriggerThreshold           uint64                 `ic:"trigger_threshold" json:"trigger_threshold"`
	NumBlocksToArchive         uint64                 `ic:"num_blocks_to_archive" json:"num_blocks_to_archive"`
	NodeMaxMemorySizeBytes     *uint64                `ic:"node_max_memory_size_bytes,omitempty" json:"node_max_memory_size_bytes,omitempty"`
	MaxMessageSizeBytes        *uint64                `ic:"max_message_size_bytes,omitempty" json:"max_message_size_bytes,omitempty"`
	ControllerId               principal.Principal    `ic:"controller_id" json:"controller_id"`
	MoreControllerIds          *[]principal.Principal `ic:"more_controller_ids,omitempty" json:"more_controller_ids,omitempty"`
	CyclesForArchiveCreation   *uint64                `ic:"cycles_for_archive_creation,omitempty" json:"cycles_for_archive_creation,omitempty"`
	MaxTransactionsPerResponse *uint64                `ic:"max_transactions_per_response,omitempty" json:"max_transactions_per_response,omitempty"`
}

type ArchivedBlocksRange struct {
	Start    BlockIndex     `ic:"start" json:"start"`
	Length   uint64         `ic:"length" json:"length"`
	Callback QueryArchiveFn `ic:"callback" json:"callback"`
}

type ArchivedEncodedBlocksRange struct {
	Callback idl.Function `ic:"callback" json:"callback"`
	Start    uint64       `ic:"start" json:"start"`
	Length   uint64       `ic:"length" json:"length"`
}

type Archives struct {
	Archives []Archive `ic:"archives" json:"archives"`
}

type Block struct {
	ParentHash  *[]byte     `ic:"parent_hash,omitempty" json:"parent_hash,omitempty"`
	Transaction Transaction `ic:"transaction" json:"transaction"`
	Timestamp   TimeStamp   `ic:"timestamp" json:"timestamp"`
}

type BlockIndex = uint64

type BlockRange struct {
	Blocks []Block `ic:"blocks" json:"blocks"`
}

type Duration struct {
	Secs  uint64 `ic:"secs" json:"secs"`
	Nanos uint32 `ic:"nanos" json:"nanos"`
}

type FeatureFlags struct {
	Icrc2 bool `ic:"icrc2" json:"icrc2"`
}

type GetBlocksArgs struct {
	Start  BlockIndex `ic:"start" json:"start"`
	Length uint64     `ic:"length" json:"length"`
}

type Icrc1BlockIndex = idl.Nat

type Icrc1Timestamp = uint64

type Icrc1Tokens = idl.Nat

type Icrc1TransferError struct {
	BadFee *struct {
		ExpectedFee Icrc1Tokens `ic:"expected_fee" json:"expected_fee"`
	} `ic:"BadFee,variant"`
	BadBurn *struct {
		MinBurnAmount Icrc1Tokens `ic:"min_burn_amount" json:"min_burn_amount"`
	} `ic:"BadBurn,variant"`
	InsufficientFunds *struct {
		Balance Icrc1Tokens `ic:"balance" json:"balance"`
	} `ic:"InsufficientFunds,variant"`
	TooOld          *idl.Null `ic:"TooOld,variant"`
	CreatedInFuture *struct {
		LedgerTime uint64 `ic:"ledger_time" json:"ledger_time"`
	} `ic:"CreatedInFuture,variant"`
	TemporarilyUnavailable *idl.Null `ic:"TemporarilyUnavailable,variant"`
	Duplicate              *struct {
		DuplicateOf Icrc1BlockIndex `ic:"duplicate_of" json:"duplicate_of"`
	} `ic:"Duplicate,variant"`
	GenericError *struct {
		ErrorCode idl.Nat `ic:"error_code" json:"error_code"`
		Message   string  `ic:"message" json:"message"`
	} `ic:"GenericError,variant"`
}

type Icrc1TransferResult struct {
	Ok  *Icrc1BlockIndex    `ic:"Ok,variant"`
	Err *Icrc1TransferError `ic:"Err,variant"`
}

type Icrc21ConsentInfo struct {
	ConsentMessage Icrc21ConsentMessage         `ic:"consent_message" json:"consent_message"`
	Metadata       Icrc21ConsentMessageMetadata `ic:"metadata" json:"metadata"`
}

type Icrc21ConsentMessage struct {
	GenericDisplayMessage *string `ic:"GenericDisplayMessage,variant"`
	LineDisplayMessage    *struct {
		Pages []struct {
			Lines []string `ic:"lines" json:"lines"`
		} `ic:"pages" json:"pages"`
	} `ic:"LineDisplayMessage,variant"`
}

type Icrc21ConsentMessageMetadata struct {
	Language         string `ic:"language" json:"language"`
	UtcOffsetMinutes *int16 `ic:"utc_offset_minutes,omitempty" json:"utc_offset_minutes,omitempty"`
}

type Icrc21ConsentMessageRequest struct {
	Method          string                   `ic:"method" json:"method"`
	Arg             []byte                   `ic:"arg" json:"arg"`
	UserPreferences Icrc21ConsentMessageSpec `ic:"user_preferences" json:"user_preferences"`
}

type Icrc21ConsentMessageResponse struct {
	Ok  *Icrc21ConsentInfo `ic:"Ok,variant"`
	Err *Icrc21Error       `ic:"Err,variant"`
}

type Icrc21ConsentMessageSpec struct {
	Metadata   Icrc21ConsentMessageMetadata `ic:"metadata" json:"metadata"`
	DeviceSpec *struct {
		GenericDisplay *idl.Null `ic:"GenericDisplay,variant"`
		LineDisplay    *struct {
			CharactersPerLine uint16 `ic:"characters_per_line" json:"characters_per_line"`
			LinesPerPage      uint16 `ic:"lines_per_page" json:"lines_per_page"`
		} `ic:"LineDisplay,variant"`
	} `ic:"device_spec,omitempty" json:"device_spec,omitempty"`
}

type Icrc21Error struct {
	UnsupportedCanisterCall   *Icrc21ErrorInfo `ic:"UnsupportedCanisterCall,variant"`
	ConsentMessageUnavailable *Icrc21ErrorInfo `ic:"ConsentMessageUnavailable,variant"`
	InsufficientPayment       *Icrc21ErrorInfo `ic:"InsufficientPayment,variant"`
	GenericError              *struct {
		ErrorCode   idl.Nat `ic:"error_code" json:"error_code"`
		Description string  `ic:"description" json:"description"`
	} `ic:"GenericError,variant"`
}

type Icrc21ErrorInfo struct {
	Description string `ic:"description" json:"description"`
}

type InitArgs struct {
	MintingAccount      TextAccountIdentifier `ic:"minting_account" json:"minting_account"`
	Icrc1MintingAccount *Account              `ic:"icrc1_minting_account,omitempty" json:"icrc1_minting_account,omitempty"`
	InitialValues       []struct {
		Field0 TextAccountIdentifier `ic:"0" json:"0"`
		Field1 Tokens                `ic:"1" json:"1"`
	} `ic:"initial_values" json:"initial_values"`
	MaxMessageSizeBytes          *uint64               `ic:"max_message_size_bytes,omitempty" json:"max_message_size_bytes,omitempty"`
	TransactionWindow            *Duration             `ic:"transaction_window,omitempty" json:"transaction_window,omitempty"`
	ArchiveOptions               *ArchiveOptions       `ic:"archive_options,omitempty" json:"archive_options,omitempty"`
	SendWhitelist                []principal.Principal `ic:"send_whitelist" json:"send_whitelist"`
	TransferFee                  *Tokens               `ic:"transfer_fee,omitempty" json:"transfer_fee,omitempty"`
	TokenSymbol                  *string               `ic:"token_symbol,omitempty" json:"token_symbol,omitempty"`
	TokenName                    *string               `ic:"token_name,omitempty" json:"token_name,omitempty"`
	FeatureFlags                 *FeatureFlags         `ic:"feature_flags,omitempty" json:"feature_flags,omitempty"`
	MaximumNumberOfAccounts      *uint64               `ic:"maximum_number_of_accounts,omitempty" json:"maximum_number_of_accounts,omitempty"`
	AccountsOverflowTrimQuantity *uint64               `ic:"accounts_overflow_trim_quantity,omitempty" json:"accounts_overflow_trim_quantity,omitempty"`
}

type LedgerCanisterPayload struct {
	Init    *InitArgs     `ic:"Init,variant"`
	Upgrade **UpgradeArgs `ic:"Upgrade,variant"`
}

type Memo = uint64

type Operation struct {
	Mint *struct {
		To     AccountIdentifier `ic:"to" json:"to"`
		Amount Tokens            `ic:"amount" json:"amount"`
	} `ic:"Mint,variant"`
	Burn *struct {
		From    AccountIdentifier  `ic:"from" json:"from"`
		Spender *AccountIdentifier `ic:"spender,omitempty" json:"spender,omitempty"`
		Amount  Tokens             `ic:"amount" json:"amount"`
	} `ic:"Burn,variant"`
	Transfer *struct {
		From    AccountIdentifier `ic:"from" json:"from"`
		To      AccountIdentifier `ic:"to" json:"to"`
		Amount  Tokens            `ic:"amount" json:"amount"`
		Fee     Tokens            `ic:"fee" json:"fee"`
		Spender *[]uint8          `ic:"spender,omitempty" json:"spender,omitempty"`
	} `ic:"Transfer,variant"`
	Approve *struct {
		From              AccountIdentifier `ic:"from" json:"from"`
		Spender           AccountIdentifier `ic:"spender" json:"spender"`
		AllowanceE8s      idl.Int           `ic:"allowance_e8s" json:"allowance_e8s"`
		Allowance         Tokens            `ic:"allowance" json:"allowance"`
		Fee               Tokens            `ic:"fee" json:"fee"`
		ExpiresAt         *TimeStamp        `ic:"expires_at,omitempty" json:"expires_at,omitempty"`
		ExpectedAllowance *Tokens           `ic:"expected_allowance,omitempty" json:"expected_allowance,omitempty"`
	} `ic:"Approve,variant"`
}

type QueryArchiveError struct {
	BadFirstBlockIndex *struct {
		RequestedIndex  BlockIndex `ic:"requested_index" json:"requested_index"`
		FirstValidIndex BlockIndex `ic:"first_valid_index" json:"first_valid_index"`
	} `ic:"BadFirstBlockIndex,variant"`
	Other *struct {
		ErrorCode    uint64 `ic:"error_code" json:"error_code"`
		ErrorMessage string `ic:"error_message" json:"error_message"`
	} `ic:"Other,variant"`
}

type QueryArchiveFn = idl.Function

type QueryArchiveResult struct {
	Ok  *BlockRange        `ic:"Ok,variant"`
	Err *QueryArchiveError `ic:"Err,variant"`
}

type QueryBlocksResponse struct {
	ChainLength     uint64                `ic:"chain_length" json:"chain_length"`
	Certificate     *[]byte               `ic:"certificate,omitempty" json:"certificate,omitempty"`
	Blocks          []Block               `ic:"blocks" json:"blocks"`
	FirstBlockIndex BlockIndex            `ic:"first_block_index" json:"first_block_index"`
	ArchivedBlocks  []ArchivedBlocksRange `ic:"archived_blocks" json:"archived_blocks"`
}

type QueryEncodedBlocksResponse struct {
	Certificate     *[]byte                      `ic:"certificate,omitempty" json:"certificate,omitempty"`
	Blocks          [][]byte                     `ic:"blocks" json:"blocks"`
	ChainLength     uint64                       `ic:"chain_length" json:"chain_length"`
	FirstBlockIndex uint64                       `ic:"first_block_index" json:"first_block_index"`
	ArchivedBlocks  []ArchivedEncodedBlocksRange `ic:"archived_blocks" json:"archived_blocks"`
}

type SendArgs struct {
	Memo           Memo                  `ic:"memo" json:"memo"`
	Amount         Tokens                `ic:"amount" json:"amount"`
	Fee            Tokens                `ic:"fee" json:"fee"`
	FromSubaccount *SubAccount           `ic:"from_subaccount,omitempty" json:"from_subaccount,omitempty"`
	To             TextAccountIdentifier `ic:"to" json:"to"`
	CreatedAtTime  *TimeStamp            `ic:"created_at_time,omitempty" json:"created_at_time,omitempty"`
}

type SubAccount = []byte

type TextAccountIdentifier = string

type TimeStamp struct {
	TimestampNanos uint64 `ic:"timestamp_nanos" json:"timestamp_nanos"`
}

type Tokens struct {
	E8s uint64 `ic:"e8s" json:"e8s"`
}

type Transaction struct {
	Memo          Memo       `ic:"memo" json:"memo"`
	Icrc1Memo     *[]byte    `ic:"icrc1_memo,omitempty" json:"icrc1_memo,omitempty"`
	Operation     *Operation `ic:"operation,omitempty" json:"operation,omitempty"`
	CreatedAtTime TimeStamp  `ic:"created_at_time" json:"created_at_time"`
}

type TransferArg struct {
	FromSubaccount *SubAccount     `ic:"from_subaccount,omitempty" json:"from_subaccount,omitempty"`
	To             Account         `ic:"to" json:"to"`
	Amount         Icrc1Tokens     `ic:"amount" json:"amount"`
	Fee            *Icrc1Tokens    `ic:"fee,omitempty" json:"fee,omitempty"`
	Memo           *[]byte         `ic:"memo,omitempty" json:"memo,omitempty"`
	CreatedAtTime  *Icrc1Timestamp `ic:"created_at_time,omitempty" json:"created_at_time,omitempty"`
}

type TransferArgs struct {
	Memo           Memo              `ic:"memo" json:"memo"`
	Amount         Tokens            `ic:"amount" json:"amount"`
	Fee            Tokens            `ic:"fee" json:"fee"`
	FromSubaccount *SubAccount       `ic:"from_subaccount,omitempty" json:"from_subaccount,omitempty"`
	To             AccountIdentifier `ic:"to" json:"to"`
	CreatedAtTime  *TimeStamp        `ic:"created_at_time,omitempty" json:"created_at_time,omitempty"`
}

type TransferError struct {
	BadFee *struct {
		ExpectedFee Tokens `ic:"expected_fee" json:"expected_fee"`
	} `ic:"BadFee,variant"`
	InsufficientFunds *struct {
		Balance Tokens `ic:"balance" json:"balance"`
	} `ic:"InsufficientFunds,variant"`
	TxTooOld *struct {
		AllowedWindowNanos uint64 `ic:"allowed_window_nanos" json:"allowed_window_nanos"`
	} `ic:"TxTooOld,variant"`
	TxCreatedInFuture *idl.Null `ic:"TxCreatedInFuture,variant"`
	TxDuplicate       *struct {
		DuplicateOf BlockIndex `ic:"duplicate_of" json:"duplicate_of"`
	} `ic:"TxDuplicate,variant"`
}

type TransferFee struct {
	TransferFee Tokens `ic:"transfer_fee" json:"transfer_fee"`
}

type TransferFeeArg struct {
}

type TransferFromArgs struct {
	SpenderSubaccount *SubAccount     `ic:"spender_subaccount,omitempty" json:"spender_subaccount,omitempty"`
	From              Account         `ic:"from" json:"from"`
	To                Account         `ic:"to" json:"to"`
	Amount            Icrc1Tokens     `ic:"amount" json:"amount"`
	Fee               *Icrc1Tokens    `ic:"fee,omitempty" json:"fee,omitempty"`
	Memo              *[]byte         `ic:"memo,omitempty" json:"memo,omitempty"`
	CreatedAtTime     *Icrc1Timestamp `ic:"created_at_time,omitempty" json:"created_at_time,omitempty"`
}

type TransferFromError struct {
	BadFee *struct {
		ExpectedFee Icrc1Tokens `ic:"expected_fee" json:"expected_fee"`
	} `ic:"BadFee,variant"`
	BadBurn *struct {
		MinBurnAmount Icrc1Tokens `ic:"min_burn_amount" json:"min_burn_amount"`
	} `ic:"BadBurn,variant"`
	InsufficientFunds *struct {
		Balance Icrc1Tokens `ic:"balance" json:"balance"`
	} `ic:"InsufficientFunds,variant"`
	InsufficientAllowance *struct {
		Allowance Icrc1Tokens `ic:"allowance" json:"allowance"`
	} `ic:"InsufficientAllowance,variant"`
	TooOld          *idl.Null `ic:"TooOld,variant"`
	CreatedInFuture *struct {
		LedgerTime Icrc1Timestamp `ic:"ledger_time" json:"ledger_time"`
	} `ic:"CreatedInFuture,variant"`
	Duplicate *struct {
		DuplicateOf Icrc1BlockIndex `ic:"duplicate_of" json:"duplicate_of"`
	} `ic:"Duplicate,variant"`
	TemporarilyUnavailable *idl.Null `ic:"TemporarilyUnavailable,variant"`
	GenericError           *struct {
		ErrorCode idl.Nat `ic:"error_code" json:"error_code"`
		Message   string  `ic:"message" json:"message"`
	} `ic:"GenericError,variant"`
}

type TransferFromResult struct {
	Ok  *Icrc1BlockIndex   `ic:"Ok,variant"`
	Err *TransferFromError `ic:"Err,variant"`
}

type TransferResult struct {
	Ok  *BlockIndex    `ic:"Ok,variant"`
	Err *TransferError `ic:"Err,variant"`
}

type UpgradeArgs struct {
	Icrc1MintingAccount *Account      `ic:"icrc1_minting_account,omitempty" json:"icrc1_minting_account,omitempty"`
	FeatureFlags        *FeatureFlags `ic:"feature_flags,omitempty" json:"feature_flags,omitempty"`
}

type Value struct {
	Nat  *idl.Nat `ic:"Nat,variant"`
	Int  *idl.Int `ic:"Int,variant"`
	Text *string  `ic:"Text,variant"`
	Blob *[]byte  `ic:"Blob,variant"`
}
