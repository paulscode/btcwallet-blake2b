package wallet

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/require"
)

// TestCoinbaseMaturityByHeight checks that the wallet holds a coinbase for the
// depth its own height calls for: on mainnet 6480 blocks before 979919 and
// 25920 from it (Bitcoin Knots 29.4.3), and CoinbaseMaturity on a network with
// no such rule.
func TestCoinbaseMaturityByHeight(t *testing.T) {
	t.Parallel()

	main := &chaincfg.MainNetParams
	require.EqualValues(t, 6480, coinbaseMaturity(main, 961640))
	require.EqualValues(t, 6480, coinbaseMaturity(main, 979918))
	require.EqualValues(t, 25920, coinbaseMaturity(main, 979919))

	require.EqualValues(t, 100, coinbaseMaturity(
		&chaincfg.RegressionNetParams, 979919,
	))

	// Never less than CoinbaseMaturity, however the params were set.
	raised := chaincfg.RegressionNetParams
	raised.CoinbaseMaturity = 500
	require.EqualValues(t, 500, coinbaseMaturity(&raised, 1))
}
