
# NHSValidator

A NHSValidator which validates existing NHS NUMBERS and is capable of generating new valid ones.
    

## Documentation

The main struct is NHSManager{}.
Two main functions are bundled to it, Validate() and Generate().

Generate() - Written in such way that it should always succeed, returns a valid NHS number as string.

Validate(string) - Takes a string (NHS number) and validates it using the official NHS specification and restrictions of the NHS NUMBER.




## Running Tests

To run tests, run the following command

```bash
  go test
```

