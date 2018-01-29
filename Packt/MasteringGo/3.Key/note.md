# Methods

- Go does not have classes
- A method is a function with a special receiver argument
    - func (recv <type>)fn()
    - The type to which the method belongs is known as the receiver
- <type> has to be declared in the same package
- Pointer receivers are used when the method needs to change the values to which the receive points
- Pointers only save memory because they only pass the memory address around

# Interfaces

- A set of method signatures
- Any type that implements interface methods is considered implicity be a child of the interface
- A value of interface type can hold any value that implements these methods
- Pointers and functions implement interfaces too since they are all types
- Go expects you to use built-in order to expand functionality 