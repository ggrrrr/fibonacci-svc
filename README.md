# [Coding challenge](https://gist.github.com/stepanbujnak/7fa18e2e97de2fd3f593c00b09c445c2)

# Sr. Go Engineer challenge

Please design and implement a web based API that steps through the Fibonacci sequence.

The API must expose 3 endpoints that can be called via HTTP requests:

* `current` - returns the current number in the sequence
* `next` - returns the next number in the sequence
* `previous` - returns the previous number in the sequence

Example:

    current -> 0
    next -> 1
    next -> 1
    next -> 2
    previous -> 1

Requirements:

* The API must be able to handle high throughput (~1k requests per second)
* The API should also be able to recover and restart if it unexpectedly crashes
* Use Go and any framework of your choice for the backend
* The submission should be sent in a GitHub repo


----
# What is Fibonacci sequence

 `In mathematics, the Fibonacci sequence is a sequence in which each number is the sum of the two preceding ones` from [source](https://en.wikipedia.org/wiki/Fibonacci_sequence)

   * Each number is >= 0 e.g. only positive numbers   
   * So we need to keep track two numbers
    
      * `Current` number -> last calculated sequence number e.g.: `N`
      * `Previous` number ->  previous calculated sequence number e.g.: `N-1`
  
   * Also we need to have initial values of these numbers (`Current` and `Previous`). Based on the same source, the initial numbers (previous and current) will be set to 0, and our first few iterations of `Next` operation will be as follows:
     * `init` value Fi(P: 0, C: 0)
     * `next` -> FI(P: 0, C: 1)
     * `next` -> FI(P: n-1.C, C: n-1.C + n-1.P ) => (P: 0, C: 0 + 1=>1)
     * `next` -> FI(P: n-1.C, C: n-1.C + n-1.P ) => (P: 1, C: 1 + 1=>2)
     * ...


# Design considerations

1. Based on the requirement `... steps through the Fibonacci sequence ...`: I understand that the API service will need to store its current/last state, also that the service will need to store only one state e.g. only one *Fibonacci counter*.

2. We will need design the system with ACID operations in mind .e.g thread safe operations

3. Endpoints
   1. `next` -> will fetch the current internal state and calculate next value, store that value in the internal store and return the calculated number.
   2. `current` -> will fetch current internal state and will return the current number.
   3. `previous` -> to my understanding this method will need to behave same as `next`, but in reverse -> FI(P: n-1.C - n-1.P, P: n-1.P, ).
   
4. Also based on the understanding that we need to have only one *Fibonacci counter*, we can have single instance service e.g. we will not be able to scale this service horizontally, in general this is not ideal, but for this case it will cover the requirements, and will allow us to have much faster ACID operations.

    __In case requirements change, we will need to re-design the solution.__

1. `The API must be able to handle high throughput (~1k requests per second)`

     * Here we need to have very fast persistent storage (repository).
     
     * Also in some `database` the insert operations are faster then select.
         
         We can save OP time by only storing the changes into external storage, as well as keep the current value in the memory of the service, so will have the following flow
         1. Next call
            1. READ last FI from RAM -> CPU OP with multi-threading safety(mutex)
            2. Calc next FI -> CPU OP
            3. Store FI in RAM ->CPU OP with multi-threading safety(mutex)
            4. Store FI in persistent store -> Network OP 
            5. Return val -> Network OP
            6. Exit
     * In some cases, it is possible to make the Store FI operation faster, by making it asynchronous, e.g. some kind of MessageBus/EventBus publishing, such solution could be suitable for some much more complex systems, and also very hard to implement, so it will not be used in here.

2. `The API should also be able to recover and restart if it unexpectedly crashes`


     * By storing the last calculated FI number in external storage, each time the service starts, it will read that number and use it as `Current` value.
     
     * We can select any number of external storage solutions. Even we can support multiple, so depending on the infrastructure we can configure it to use different one.
  
    I have implemented Redis, PostgreSQL and in memory, so I can showcase the difference in performance between different storage solutions. 

# Implementation and source tree organization

   * `/common` -> some common packages, this folder can be shared between other golang projects.
     * `/log` -> logging based on `zerolog` package
     * `/api` -> models which are shared between services, in our case only a base model for response.
     * `/system` -> Here I have put tooling for http listener, shutdown process, http error response, these package can be shared by other projects/services.
   * `/internal` -> Business logic and packages related only to this service, this folder is protected by the GoLang package management and can not be imported in external projects.
     * `/fi` -> Methods and data models only related to the definition of `Fibonacci sequence`
     * `/api` -> HTTP handler methods
     * `/app` -> Application layer -> fetching initial data, keeping the in memory storage thread safe for each operation, storing data and external error handling.
     * `/repo` -> interface/implementations related to storing and fetch data from external repository/storage, which are needed for the App layer.
       * `/repo/ramrepo` -> implementation for in memory repository, added only for some testing and POC.
       * `/repo/pgrepo` -> implementation for PostgreSQL.
       * `/repo/redisrepo` -> implementation for Redis.
   * `/sql` -> files which are needed to initialize Postgres databases for automation testing and/or local running.
   * `/docker-compose.yaml` -> all containers needed for test and/or run the service locally or on Ci/Cd pipeline.
   * `/main.go` -> entry  point of the service.
   * `/Makefile` -> help file for local and Ci/Cd pipeline tooling

# Running/Testing

## Requirements
  1. golang version > 1.21
  2. gomock for interface mocks (unit testing)
     
    go install go.uber.org/mock/mockgen@latest
  
  3. make tool
  4. Docker and docker-compose
  5. Load test of the API `ab` [apache benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html)
 
