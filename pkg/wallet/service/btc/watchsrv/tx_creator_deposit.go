package watchsrv

import (
	"go.uber.org/zap"

	"github.com/hiromaily/go-crypto-wallet/pkg/account"
	"github.com/hiromaily/go-crypto-wallet/pkg/action"
)

// CreateDepositTx create unsigned tx if client accounts have coins
// - sender: client, receiver: deposit
// - receiver account covers fee, but is should be flexible
func (t *TxCreate) CreateDepositTx(adjustmentFee float64) (string, string, error) {
	sender := account.AccountTypeClient
	receiver := t.depositReceiver
	targetAction := action.ActionTypeDeposit
	t.logger.Debug("account",
		zap.String("sender", sender.String()),
		zap.String("receiver", receiver.String()),
	)

	requiredAmount, err := t.btc.FloatToAmount(0)
	if err != nil {
		return "", "", err
	}

	// create deposit transaction
	return t.createTx(sender, receiver, targetAction, requiredAmount, adjustmentFee, nil, nil)
}
