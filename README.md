# ddd-clean-architect-go-template
a personal template for ddd-clean-architect


## Command

A command describes an **action** that should be performed; it's always named in the imperative tense such as `UpdateUser` or `CreateUser`.


## Event

An event is the notification that something happened in the past. You can view an event as the representation of the reaction to **a command after being executed**. 
All events should be represented as verbs in the past tense such as `UserUpdated`, `UserCreated`.
We create the `UserUpdated` event; it's a pure go struct, and it's the past equivalent to the previous command `UserCreated`:

## Aggregate

The aggregate is a logical boundary for things that can change in a business transaction of a given context. 
In this project context, it simplifies the process the commands and produce events.