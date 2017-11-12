# perusseslack
Little web API for Pérusse quotes which can easily be deployed to Heroku.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```powershell
 go get -u github.com/fredroger/perusseslack
 cd $env:GOPATH/src/github.com/fredroger/perusseslack
 heroku local
```

Should now be running on [localhost:5000](http://localhost:5000/).

You should also install [govendor](https://github.com/kardianos/govendor) if you are going to add any dependencies to this app.

## Deploying to Heroku

```powershell
 heroku create
 git push heroku master
 heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
