# how to use

## examples


### parse string
```go
j := ini.Loads("[main]\nkey = 'value==val;lues'\n; key3 = 'value3\nkey432' =     \"423890fsda=432fds=f=\"")

// get a key
j["section"]["key"]
```

### convert map to ini
```go
j2 := make(map[SectionKey]Section)

j2["main"] = make(Section)
j2["main"]["key"] = "value"

j2["debug"] = make(Section)
j2["debug"]["keyfsafdsa"] = "value::==1=23=:??!#?/'\"a\"'"

j3 := ini.Dumps(j2)
```

### parse file
```go
j, _ := ini.Load("file.ini")

// get a key
j["section"]["key"]
```

### convert map to ini and write it to a file
```go
j2 := make(map[SectionKey]Section)

j2["main"] = make(Section)
j2["main"]["key"] = "value"

j2["debug"] = make(Section)
j2["debug"]["keyfsafdsa"] = "value::==1=23=:??!#?/'\"a\"'"

j3 := ini.Dump(j2, "file.ini")
```