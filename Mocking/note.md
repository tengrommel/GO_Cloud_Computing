# Mocking

Test Driven Development

## Testing the Server/Invoked entity

- Approach: Writing tests against server functionality
- Server = Model/Implementation
- Client = Test(Independent/dependent vars)

## Stub vs. Mock

- Stub
    - Emulate responses of production system
        - Stupid coopertator
    - Hardcodes response behaviour
    - Client tests (state)
    - Generally doesn't require specific framework(beyond test framework, of course)
    
- Mock
    - Can be used for test inversion -> test of client-side behaviour
    - Check function/method invocations(arguments)
    - Tests (interaction/behavour)

## Test Specification
- Arrange
    - Organise resources set up and teardown
    - Environment

- Act
    - Invoke desired functionality
    - Run developed tests (Independent variables)
    
- Assert
    - Check the results
    - Individual logical statements(easier to debug)
        - Assertions more readable than conditional statement
    - Evaluate results(Dependent variables)
 
 #GoMock
 
- Install packages
    - go get github.com/golang/mock/gomock
    - go get github.com/golang/mock/mockgen
- Specify interface (needs to match production service)
- Generate mock
    - mockgen -source=mockObject.go -package=main -destination=mockGen.go
- Write Test
    - Initialise MockController
    - Inject controller into mocked interface
    - Declare expectations
    - Invoke mock with client functionality