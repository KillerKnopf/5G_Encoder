# 5G_Encoder
 
## Description

This GO module is used to encode an 5GS update type information element following the technical specification 3GPP TS 24.501 v16.9.0 (2021-06) heading 9.11.3.9A.

It uses structex to declare bitfields and encodes them using it's methods.

## Testing

### Run the unit test

To run the provided unit test run following command from a command line:

```
go test
```

### Test coverage

The test coverage of the provided unit test is 42.9%
For a visualization open the file titled "5G_Encoder: Go Coverage Report.html" within this directory.

To create test coverage report on your own run following two statements in your command line:

first

```
go test -coverprofile=cover.out
```

and then

```
go tool cover -html=cover.out
```