# Observer architectural refactoring proposal
================================================

## Goals:
- Create a top down structure to allow for easy navigation and code readability
- Reduce amount of code repetition and add separation of concerns (http/db/bl/dal/areas) 
- Make the code testable by reducing mocking requirements and providing contextual support to eliminate static state


## Proposed package structure

* app\
    * main.go - The initialization of the end points and application the context
    * context\ - App context implementation that needs to be propagated as the first parameter for most methods
    * endpoints\ - End point implementations which map end points to api code and initialize the call/session context
        * rest.go
        * grpc.go
    * api\ - Api code separated into business areas which interacts with the dal and returns/accepts api models
        * model\ - Api related models
            * charts\ 
            * branches\
            * stonks\
        * charts\ 
        * branches\
        * stonks\
    * dal\ - Dal code which interacts with the databases and returns/accepts dal models
        * model\ - Dal related models
            * charts\ 
            * branches\
            * stonks\
        * bigquery\
            * charts.go 
            * branches.go
            * stonks.go
        * postgres\
            * charts.go
            * branches.go
            * stonks.go

