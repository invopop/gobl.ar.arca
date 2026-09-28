package arca_test

import (
	"testing"

	arca "github.com/invopop/gobl.ar.arca/addon"
	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/cal"
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/num"
	"github.com/invopop/gobl/org"
	"github.com/invopop/gobl/tax"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCorrectWithNormalize(t *testing.T) {
	t.Run("copies tax extensions to preceding", func(t *testing.T) {
		inv := testInvoiceARForCorrection(t)
		require.NoError(t, inv.Calculate())
		assert.Equal(t, cbc.Code("1"), inv.Tax.Ext.Get(arca.ExtKeyDocType))
		assert.Equal(t, cbc.Code("1"), inv.Tax.Ext.Get(arca.ExtKeyConcept))

		err := inv.Correct(
			bill.Credit,
			bill.WithExtension(arca.ExtKeyDocType, "3"), // Credit Note A
		)
		require.NoError(t, err)

		// Original extensions copied to preceding
		pre := inv.Preceding[0]
		assert.Equal(t, cbc.Code("1"), pre.Ext.Get(arca.ExtKeyDocType))
		assert.Equal(t, cbc.Code("1"), pre.Ext.Get(arca.ExtKeyConcept))

		// Invoice doc type set via correction normalizer
		assert.Equal(t, cbc.Code("3"), inv.Tax.Ext.Get(arca.ExtKeyDocType))
	})

}

func testInvoiceARForCorrection(t *testing.T) *bill.Invoice {
	t.Helper()
	return &bill.Invoice{
		Regime:    tax.WithRegime("AR"),
		Addons:    tax.WithAddons(arca.V4),
		Series:    "1",
		Code:      "123",
		IssueDate: cal.MakeDate(2024, 1, 15),
		Supplier: &org.Party{
			TaxID: &tax.Identity{
				Country: "AR",
				Code:    "20345678904",
			},
		},
		Customer: &org.Party{
			TaxID: &tax.Identity{
				Country: "AR",
				Code:    "30500010912",
			},
		},
		Lines: []*bill.Line{
			{
				Quantity: num.MakeAmount(10, 0),
				Item: &org.Item{
					Name:  "Test Item",
					Price: num.NewAmount(10000, 2),
					Key:   org.ItemKeyGoods,
				},
				Taxes: tax.Set{
					{
						Category: "VAT",
						Rate:     "standard",
					},
				},
			},
		},
	}
}
