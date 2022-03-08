# go-ipfs

Anonfiles API wrapper written in Go.

## Usage

### Get information about the file

#### Get full information

You can get the information in two ways. The first is to get full struct with all the fields.

```go
af.GetInfo("https://anonfiles.com/u1C0ebc4b0/file.txt")
```

Will return a structure which completely copying the standard JSON output, or error message.

#### Get only one field

It is possible to get only one field from the JSON responce. To do this, just use one of the functions below.

```go
af.GetStatus("https://anonfiles.com/u1C0ebc4b0/file.txt") // true or false
af.GetFullURL("https://anonfiles.com/u1C0ebc4b0/")
af.GetShortURL("https://anonfiles.com/u1C0ebc4b0/file.txt")
af.GetName("https://anonfiles.com/u1C0ebc4b0/file.txt") // file_txt
af.GetSizeBytes("https://anonfiles.com/u1C0ebc4b0/file.txt") // 6861
af.GetSizeReadable("https://anonfiles.com/u1C0ebc4b0/file.txt") // 6.7 KB
```

From above functions every will return result and error.

### Upload a file

Not ready yet.

## Roadmap

- [x] Parse file id from the link. GetId()
- [ ] Get information about the file. GetInfo()
    - [ ] GetStatus()
    - [ ] GetFullURL()
    - [ ] GetShortURL()
    - [ ] GetName()
    - [ ] GetSizeBytes()
    - [ ] GetSizeReadable()
- [ ] Upload file to the server. Upload()
- [ ] Get temporary download link (from page HTML)