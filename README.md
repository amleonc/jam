# JAM
### Jwt Auth Middleware for Go


This project proposes a JWT middleware for Go applications, with the core functionality to verify requests with JWT objects. The underlying JOSE implementation comes from the [`jwx`](https://github.com/lestrrat-go/jwx) library. It is just a simpler, less flexible version of the [`jwauth`](https://github.com/go-chi/jwtauth) middleware, brought to us by the [`chi`](https://go-chi.io/) team.
