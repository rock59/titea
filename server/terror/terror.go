package terror

import "errors"

var (
	ErrCommand         error = errors.New("ERR command error")
	ErrCmdParams       error = errors.New("ERR command params error")
	ErrCmdDisable      error = errors.New("ERR command disable")
	ErrKeyEmpty        error = errors.New("ERR key cannot be empty")
	ErrKeyOrFieldEmpty error = errors.New("ERR key or field cannot be empty")
	ErrTypeNotMatch    error = errors.New("ERR WRONGTYPE Operation against a key holding the wrong kind of value")
	ErrTypeTrans       error = errors.New("ERR value is not an integer or out of range")
	ErrCmdInBatch      error = errors.New("ERR some command in batch not supported")
	ErrCmdNumber       error = errors.New("ERR command not enough in batch")
	ErrBackendType     error = errors.New("ERR backend type error")
	ErrTypeAssertion   error = errors.New("ERR interface type assertion failed")
	ErrOutOfIndex      error = errors.New("ERR index out of range")
	ErrInvalidMeta     error = errors.New("ERR invalid key meta")
	ErrUnknownType     error = errors.New("ERR unknown response data type")
	ErrRunWithTxn      error = errors.New("ERR run run with txn")
	ErrAuthNoNeed      error = errors.New("ERR Client sent AUTH, but no password is set")
	ErrAuthFailed      error = errors.New("ERR invalid password")
	ErrAuthRequired    error = errors.New("NOAUTH Authentication required")
	ErrKeyBusy         error = errors.New("BUSYKEY key is deleting, retry later")
	ErrNotInteger      error = errors.New("ERR value is not an integer or out of range")
	ErrTxnFailed       error = errors.New("ERR Txn execute failed")
	ErrKeyNotExist     error = errors.New("ERR Key not exist")
	ErrKeyExist        error = errors.New("ERR Key Exists")
	ErrKeyNotValid     error = errors.New("ERR Key not valide")
	ErrIdentify        error = errors.New("ERR Identify Type Error")
	ErrUnbuildKey      error = errors.New("ERR unbuild key")
	ErrNsInvalied      error = errors.New("ERR namespace invalied group.service")
	ErrNsExist         error = errors.New("ERR namespace already exist")
	ErrNsNotExist      error = errors.New("ERR namespace not exist")
)