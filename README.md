# fso

firestore operator command

## Usage


```
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/your/credential.json
```

write

```shell-session
$ cat write.json | fso set projectId collectionId documentId
```

read

```shell-session
$ fso get projectId collectionId documentId
```


## Test

```
$ cp .env.sample .env
# edit to fill .env

$ ./script/test.sh
```


