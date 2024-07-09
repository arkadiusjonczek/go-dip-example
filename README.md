# go-dip-example

**Example of the Dependency Inversion Principle (DIP) from SOLID.**

This example uses a `Switch` with a `Switchable` interface.

The `Switch` uses a high level abstraction, and you can connect anything implementing the `Switchable` interface.
The concrete implementation will be injected using the constructor.

In this example you will find a `Lamp` and `CoffeeMaschine` as concrete examples.
