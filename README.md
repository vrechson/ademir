# Ademir
Ademir is a bot designed to simulate administration behavior in CTF XSS challenges.

It only opens pages using a chrome web driver and renders any javascript in it.

There are two modes of operation:

* env `USEWEBHOOK=true`
    In this mode of operation, Ademir setup's a webhook at `http://localhost/ademir/dashboard` where it navigate into each url it receives via `address POST parâmeter`.
* env `USEWEBHOOK=false`
    In this mode, just run it as `./ademir url`, it will open your url and end its execution.

`.env` example:

```
PORT=80
FLAG="CTF-BR{BelevieThisIsTheRealFlag}"
INSTANCES=10
USEWEBHOOK=true
DEBUG=true
```

# Running it
`docker build -t ademir .`

`docker run -p 6666:80 -it ademir`

environment will be running at http://localhost:6666

# TODO
- [ ] task 1: Improve cookies management
- [ ] task 2: Code Refactoring

## Improve Cookies management
To navigate to the hooked page with cookies defined, currently, the bot first has to open the hostname page, set cookies and then navigate into the hooked page with cookies already set

At the end of this task, it won't be necessary to navigate any page before set cookies
## Code Refactoring
This version is still with many problems, so it's essential to refactor its code and turn it mo clear and concise.

At the end of this task, the code must be more readable.