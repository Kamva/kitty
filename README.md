#### Hexa contain services kit.

#### Requirements:
go : minimum version `1.13`

#### Install
```
go get github.com/Kamva/hexa
```

#### Available Services:
- Config: config service.
- Logger: logger service
- Translator: translator service.

#### How to use:
example:
```go
// Assume we want to use viper as config service.
v := viper.New()

// tune your viper.
config := hexaconfig.NewViperDriver(v)

// Use config service in app.
```

#### Proposal
- [ ] We can implement `WithUser` method on the context to change its user.  
e.g: in each cron job maybe we need to get a user, so after get the user, we can call to this
method to set the user as context's user.

- [ ] We can implement `WithPermission` method on `User` interface to add a permission to the user's 
permissions temporary.  
e.g: In cron jobs we need to add something that needs to root permission, so we can add to our cron job's guest user
root permission temporary (with specific middleware that skip guests check) to do that job. 

- [ ] We should implement Distributed tracing by using zipkin and open tracing and also use it in gRPC,... . but I think this
should be in the service mesh not in the business logic.

#### Todo
- [ ] Write Tests
- [ ] Add badges to readme.
- [ ] CI

#### Client conventions:
- [ ] on occure error and want to say to user that "some error occured",  give the requestID to the user, then user can give it back to the support team, and support team can see logs relative to that request id.