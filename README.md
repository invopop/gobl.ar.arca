# GOBL ➡️ Argentina ARCA

Argentina ARCA (Agencia de Recaudación y Control Aduanero) e-invoicing addon for [GOBL](https://github.com/invopop/gobl).

Copyright [Invopop S.L.](https://invopop.com) 2026. Released publicly under the [GNU Affero General Public License v3.0](LICENSE). For commercial licenses please contact the [dev team at invopop](mailto:dev@invopop.com).

[![Lint](https://github.com/invopop/gobl.ar.arca/actions/workflows/lint.yaml/badge.svg)](https://github.com/invopop/gobl.ar.arca/actions/workflows/lint.yaml)
[![Test Go](https://github.com/invopop/gobl.ar.arca/actions/workflows/test.yaml/badge.svg)](https://github.com/invopop/gobl.ar.arca/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/invopop/gobl.ar.arca)](https://goreportcard.com/report/github.com/invopop/gobl.ar.arca)
[![codecov](https://codecov.io/gh/invopop/gobl.ar.arca/graph/badge.svg)](https://codecov.io/gh/invopop/gobl.ar.arca)
[![GoDoc](https://godoc.org/github.com/invopop/gobl.ar.arca?status.svg)](https://godoc.org/github.com/invopop/gobl.ar.arca)
![Latest Tag](https://img.shields.io/github/v/tag/invopop/gobl.ar.arca)

This module implements the Argentina ARCA v4 e-invoicing requirements as a GOBL
tax addon (`ar-arca-v4`), based on the ARCA web services for electronic
invoices (WSFE) and tourism invoices (WSCT). It covers invoices, credit notes,
and debit notes of types A, B, C, and T, with validation rules registered under
the `AR-ARCA` namespace.

The addon lives in the `addon/` subpackage and registers into GOBL's global
registry on import; the module root is reserved for future ARCA tooling. The
Argentine tax regime (`regimes/ar`) stays in GOBL core.

## Usage

Add a blank import of the addon, then declare it on documents:

```go
import (
	_ "github.com/invopop/gobl.ar.arca/addon"
)
```

```json
{
	"$schema": "https://gobl.org/draft-0/bill/invoice",
	"$regime": "AR",
	"$addons": ["ar-arca-v4"]
}
```

> **Note**: `ar-arca-v4` is an approved external addon key in GOBL core, but
> documents declaring it fail validation unless this module is imported.

## Extensions

| Key | Description |
| --- | --- |
| `ar-arca-doc-type` | Document type (A, B, C, or T variants). Normalized from the customer VAT status and invoice type when missing. |
| `ar-arca-concept` | Concept: goods, services, or both. Normalized from the line items. |
| `ar-arca-identity-type` | Customer identity document type for customers without a tax ID. |
| `ar-arca-vat-status` | Customer VAT status. Normalized from the customer tax ID. |
| `ar-arca-vat-rate` | VAT rate code for tax combos. Normalized from the GOBL rate. |
| `ar-arca-tax-type` | Tax type for charges reported as other taxes. |
| `ar-arca-tourism-type` | Issuer-receiver relationship for tourism (type T) invoices. |
| `ar-arca-tourism-item` | Tourism item code for type T lines and discounts. |

## Development

`examples/` holds sample documents with golden envelopes under `examples/out/`;
regenerate them with:

```sh
go test . -run TestExamples -update
```

## Sources

- [ARCA WSFE developer manual (COMPG v4.0, PDF)](https://www.afip.gob.ar/ws/documentacion/manuales/manual-desarrollador-ARCA-COMPG-v4-0.pdf)
- [ARCA WSCT developer manual v1.6.4 (PDF)](https://www.afip.gob.ar/ws/documentacion/manuales/Manual_Desarrollador_WSCT_v1.6.4.pdf)
