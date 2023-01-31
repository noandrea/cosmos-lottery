package lottery_test

import (
	"testing"

	"lottery/x/lottery"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *lottery.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: lottery.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &lottery.GenesisState{

				// this line is used by starport scaffolding # lottery/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # lottery/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
