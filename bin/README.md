Everything you need to run the server is here in in the bin directory
You dont have to install go and its dependencies
The server is build in a single binary file to run on linux

## Steps to run the server

### Development Use only

#### METHOD 1

All you need to do is

- navigate to the bin directory
- create a `.env` file and fill it up with what's needed as higlighted in the `env.example` file
- then run `dotenv -f .env run ./hms`

dotenv here is a program to load env variable from file. if you dont have it, install it using
`sudo apt install dotenv`

you may be asked to make hms an executable, do that with the ccommand below
`chmod +x hms`

#### METHOD 2

simply run the bash script ./run_hmslike this

```bash
./run_hms
```

\*\*\* make sure you're in the bin directory when doing all these.
