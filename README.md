# belajar-golang-dependency-injection
Dependency Injection
    - https://github.com/google/wire
    - https://github.com/uber-go/fx
    - https://github.com/golobby/container
go get github.com/google/wire
go install github.com/google/wire/cmd/wire@latest

Provider
    - Create Constructor like OOP
Injector
    - wire.Build(constructorA,constructorB)
    
Dependency Innjection
    - generate command
    wire gen name_package
    //go:build wireinject
    // +build   wireinject
    - setting.json
    "gopls": {
        "buildFlags": ["-tags=wireinject"]
      },
Error
Parameter Injector
    - isError bool
Multiple Binding
    - Google Wire not support provider with same data type
    - use alias to init same data type