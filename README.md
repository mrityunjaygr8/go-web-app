# Go Web App
An opinionated boilerplate template for creating web apps written in golang.
Hugely inspired by the [Lets Go Further](https://lets-go-further.alexedwards.net/) and [this talk](https://www.youtube.com/watch?v=rWBSMsLG8po&pp=ygUaZ29sYW5nIHdlYiBzZXJ2aWNlIDcgeWVhcnM%3D)
- Using [Viper](github.com/spf13/viper) for reading env vars and `.env` files.
- Uses [Chi](github.com/go-chi/chi) as the router, and for many helpful middlewares
- Uses [Health](https://github.com/alexliesenfeld/health) for checking health of components
- Uses [Logrus](https://github.com/sirupsen/logrus) for structured logging
- Uses the [standard Postgres driver](github.com/lib/pq)
- Uses [Air](https://github.com/cosmtrek/air) for hot reload
- Uses [Task](https://github.com/go-task/task) as a make replacement

Feedback welcome.
