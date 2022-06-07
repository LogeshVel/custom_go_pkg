### custom_go_pkg
Working with custom Golang package in our Project with Go Module initialized.


The custom Go package locates in our Project directory itself.

PATH:

Need to add the Project's root directory path to the GOPATH

go mod

```
go env -w GO111MODULE=on
go mod init <project_root_foldername>
```

For more details checkout my [Medium article](https://medium.com/@LogeshSakthivel/custom-package-in-golang-with-without-go-module-go-mod-c1fe67560680)