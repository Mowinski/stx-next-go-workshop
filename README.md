# Taks 6

Implement `main` function in `main.go` file. 

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

PL
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

PL
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
