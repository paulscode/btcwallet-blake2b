package wallet

import "github.com/btcsuite/btcd/chaincfg"

// coinbaseMaturity returns the depth a coinbase output mined at height must
// reach before the wallet offers it for spending or counts it as spendable.
//
// Upstream reads chaincfg.Params.CoinbaseMaturity, one depth for every
// coinbase. On the Bitcoin BLAKE2b chain the depth depends on where the
// coinbase was mined (Bitcoin Knots 29.4.3: 6480 blocks before 979919, 25920
// from it), so this asks the chain parameters of the btcd fork the wallet is
// built against, and never answers less than CoinbaseMaturity.
func coinbaseMaturity(params *chaincfg.Params, height int32) int32 {
	maturity := params.RelayCoinbaseMaturityAt(height)
	if base := int32(params.CoinbaseMaturity); base > maturity {
		maturity = base
	}

	return maturity
}
