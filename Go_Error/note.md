# Error

## Introduction

> Errors are just values, conforming to an interface:
    
    type error interface{
        Error() string
    }

> Package 'errors' implements utility for constructing errors:

    // errorString is a trivial implementation of error.
    type errorString struct {
        s string
    }
    
    func (e *errorString) Error() string {
        return e.s
    }

