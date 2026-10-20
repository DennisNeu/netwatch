# netwatch

A go application to check if web hosted services are online.

## How it works

The application needs a "services.json" file in it's source directory.

Example:
```json
[
  {
  "name": "Tagesschau",
  "url": "https://tagesschau.de"
  },
  {
  "name": "Youtube",
  "url": "https://youtube.com"
  }
]
```
The application will make a HTTP request to each of these services to check if they are up and running.

## How to install

Clone the repo and execute the following command in the source directory:
```bash
go run netwatch.go
```

Go needs to be installed.

I haven't released a proper binary yet and since this is just a learning project, I probably never will. 
