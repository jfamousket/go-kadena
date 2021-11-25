package helpers

import (
	"time"

	"github.com/jfamousket/go-kadena/common"
)

type PrepareCommand struct {
	KeyPairs  []common.KeyPair
	CmdType   common.CmdType
	Nonce     string
	Proof     string
	Rollback  bool
	Step      uint64
	PactId    string
	EnvData   interface{}
	Meta      *common.Meta
	NetworkId string
	PactCode  string
}

func PrepareExecCommand(cmd PrepareCommand) common.Command {
	EnforceType(cmd.Nonce, "string", "nonce")
	EnforceType(cmd.PactCode, "string", "pactCode")

	_ = getSigners(cmd)
	cmd.Meta = getMeta(cmd)

	exec := common.Exec{
		Data: cmd.EnvData,
		Code: cmd.PactCode,
	}

	cmdToSend, err := MarshalBody(common.CommandField{
		Signers: []common.Signer{},
		Meta:    *cmd.Meta,
		Nonce:   cmd.Nonce,
		Payload: common.Payload{
			Exec: exec,
		},
		NetworkId: cmd.NetworkId,
	})
	EnforceNoError(err)

	return makeSingleCmd(cmdToSend.Bytes())
}

func PrepareContCmd(cmd PrepareCommand) common.Command {
	EnforceType(cmd.Nonce, "string", "nonce")

	cmd.Nonce = getCmdNonce(cmd)

	_ = getSigners(cmd)

	cont := common.Cont{
		Proof:    cmd.Proof,
		PactId:   cmd.PactId,
		Rollback: cmd.Rollback,
		Step:     cmd.Step,
		Data:     cmd.EnvData,
	}

	cmdToSend, err := MarshalBody(common.CommandField{
		Nonce:   cmd.Nonce,
		Meta:    *cmd.Meta,
		Signers: []common.Signer{},
		Payload: common.Payload{
			Cont: cont,
		},
		NetworkId: cmd.NetworkId,
	})

	EnforceNoError(err)

	return makeSingleCmd(cmdToSend.Bytes())
}

func getCmdNonce(cmd PrepareCommand) string {
	if cmd.Nonce == "" {
		return time.Now().Format(time.RFC3339)
	}
	return cmd.Nonce
}

func getSigners(cmd PrepareCommand) []common.Signer {
	var signers []common.Signer
	for _, kp := range cmd.KeyPairs {
		signers = append(signers, MakeSigner(kp))
	}
	return signers
}

func getMeta(cmd PrepareCommand) *common.Meta {
	if cmd.Meta == nil {
		return MakeMeta("", "", 0, 0, 0, 0)
	}
	return cmd.Meta
}

func makeSingleCmd(cmd []byte) common.Command {

	_, hash := CreateBlake2Hash(cmd)

	return common.Command{
		Cmd:  string(cmd),
		Hash: hash,
		Sigs: []common.Sig{},
	}
}
