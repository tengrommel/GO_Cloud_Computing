#A closer look at interfaces and methods in Go 1
- Conversions

- The empty interface

#A closer look at interfaces and methods in Go 2

- Embedding
    - Embedding is Go's answer to subclassing
    - When you embed a type inside another, the embedded type is called the inner type while the other type is called the outer type
    - The outer type has access to all the exported fields of the inner type
    - The outer and inner types has to be properly initialized
- Readers and writers
*Outer type stays as a different type than the inner type, meaning that a value of the inner type can not be assigned to a value of the outer type or vice versa*

# Factory pattern
  - Creational pattern
  - The caller does not create objects directly
  - The factory function or object handles the creation of new objects based on some conditions
  
*Why Use Factory Pattern*
- Make your code easily expandable
- Separation of concerns, only troubleshoot in one place
- Efficient team work on large projects 

*How to Implement in Go*
- Create an interface 
- Create concrete types
- Create factory function or method

# Singleton pattern
> When the object needs to keep state throughout our program

# Builder pattern
