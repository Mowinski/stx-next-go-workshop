[![GoDoc](https://godoc.org/github.com/Mowinski/stx-next-go-workshop?status.svg)](https://godoc.org/github.com/Mowinski/stx-next-go-workshop)

# STX Next - Go Workshop

This application verify VAT number in Europe [VIES](https://ec.europa.eu/taxation_customs/business/vat/eu-vat-rules-topic/vat-identification-numbers_en).

## Example output for valid VAT number
```
➜  stx-next-go-workshop git:(master) ✗ ./stx-next-go-workshop
GB
<VALIDNUMBER>
Error:
ErrorCode:
XMLName: {http://schemas.xmlsoap.org/soap/envelope/ Envelope}
CountryCode: GB
VatNumber: <VALIDNUMBER>
RequestDate: 2018-06-18+02:00
Valid: true
Name: TEST COMPANY NAME
Address: MULTILINE TEST ADDRESS
STREET CITY
POSTALCODE
```


## Example output for invalid VAT number
```
➜  stx-next-go-workshop git:(master) ✗ ./stx-next-go-workshop
GB
<INVALIDNUMBER>
Error:
ErrorCode:
XMLName: {http://schemas.xmlsoap.org/soap/envelope/ Envelope}
CountryCode: GB
VatNumber: <INVALIDNUMBER>
RequestDate: 2018-06-18+02:00
Valid: false
Name: ---
Address: ---
```

# Full version
Full code version is available on branch `final` in this repo: [Link](https://github.com/Mowinski/stx-next-go-workshop/tree/final)
