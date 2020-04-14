package coldwallet

import (
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/address"
	"github.com/hiromaily/go-bitcoin/pkg/model/rdb/coldrepo"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/types"
)

// ImportPubKey import pubKey from csv file for sign wallet
//  only multisig account is available
func (w *ColdWallet) ImportPubKey(fileName string, accountType account.AccountType) error {
	if w.wtype != types.WalletTypeSignature {
		return errors.New("it's available on sign wallet")
	}

	//validate account, onl
	if !account.AccountTypeMultisig[accountType] {
		w.logger.Info("multisig address can be imported, but this account is not")
		return nil
	}

	//TODO:ImportするファイルのaccountTypeもチェックしたほうがBetter
	//e.g. ./data/pubkey/receipt/
	//receipt_1_1586831083436291000.csv
	//receipt_1_1586831083436291000.csv
	tmp := strings.Split(strings.Split(fileName, "_")[0], "/")
	if tmp[len(tmp)-1] != string(accountType) {
		return errors.Errorf("mismatching between accountType(%s) and file prefix [%s]", accountType, tmp[0])
	}

	//ファイル読み込み(full public key)
	pubKeys, err := w.addrFileStorager.ImportPubKey(fileName)
	if err != nil {
		return errors.Errorf("key.ImportPubKey() error: %s", err)
	}

	//added_pubkey_history_receiptテーブルにInsert
	addedPubkeyHistorys := make([]coldrepo.AddedPubkeyHistoryTable, len(pubKeys))
	for i, key := range pubKeys {
		inner := strings.Split(key, ",")

		//ここでは、FullPublicKeyをセットする必要がある
		addedPubkeyHistorys[i] = coldrepo.AddedPubkeyHistoryTable{
			FullPublicKey:         inner[2],
			AuthAddress1:          "",
			AuthAddress2:          "",
			WalletMultisigAddress: "",
			RedeemScript:          "",
		}
	}
	//TODO:Upsertに変えたほうがいいか？Insert済の場合、エラーが出る
	err = w.storager.InsertAddedPubkeyHistoryTable(accountType, addedPubkeyHistorys, nil, true)
	if err != nil {
		return errors.Errorf("DB.InsertAddedPubkeyHistoryTable() error: %s", err)
	}

	return nil
}

// ImportMultisigAddrForColdWallet1 coldwallet2でexportされたmultisigアドレス情報をimportする for Cold Wallet1
func (w *ColdWallet) ImportMultisigAddrForColdWallet1(fileName string, accountType account.AccountType) error {
	if w.wtype != types.WalletTypeKeyGen {
		return errors.New("it's available on Coldwallet1")
	}

	//accountチェック
	//multisigであればこのチェックはOK
	//if accountType != ctype.AccountTypeReceipt && accountType != ctype.AccountTypePayment {
	//	logger.Info("AccountType should be AccountTypeReceipt or AccountTypePayment")
	//	return nil
	//}
	if !account.AccountTypeMultisig[accountType] {
		w.logger.Info("This func is for only account witch uses multiaddress")
		return nil
	}

	//TODO:ImportするファイルのaccountTypeもチェックしたほうがBetter
	//e.g. ./data/pubkey/receipt
	tmp := strings.Split(strings.Split(fileName, "_")[0], "/")
	if tmp[len(tmp)-1] != string(accountType) {
		return errors.Errorf("mismatching between accountType(%s) and file prefix [%s]", accountType, tmp[0])
	}

	//ファイル読み込み(full public key)
	pubKeys, err := w.addrFileStorager.ImportPubKey(fileName)
	if err != nil {
		return errors.Errorf("key.ImportPubKey() error: %s", err)
	}

	//added_pubkey_history_receiptテーブルにInsert
	accountKeyTable := make([]coldrepo.AccountKeyTable, len(pubKeys))

	tm := time.Now()
	for i, pubkey := range pubKeys {
		//TODO:とりあえず、1カラムのデータを前提でコーディングしておく
		inner := strings.Split(pubkey, ",")
		//csvファイル
		//tmpData := []string{
		//	record.FullPublicKey,
		//	record.AuthAddress1,
		//	record.AuthAddress2,
		//	record.WalletMultisigAddress,
		//	record.RedeemScript,
		//}

		//Upsertをかけるには情報が不足しているので、とりあえず1行ずつUpdateする
		accountKeyTable[i] = coldrepo.AccountKeyTable{
			FullPublicKey:         inner[0],
			WalletMultisigAddress: inner[3],
			RedeemScript:          inner[4],
			AddressStatus:         address.AddressStatusValue[address.AddressStatusMultiAddressImported],
			UpdatedAt:             &tm,
		}
	}
	//Update
	err = w.storager.UpdateMultisigAddrOnAccountKeyTableByFullPubKey(accountType, accountKeyTable, nil, true)
	if err != nil {
		return errors.Errorf("DB.UpdateMultisigAddrOnAccountKeyTableByFullPubKey() error: %s", err)
	}

	return nil
}
