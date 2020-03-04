## Go Crash Course 📝💥

### Workspace 🏗

Run below command to spit out your go enviroment variables.

```
go env
```

In the output find the `GOPATH` variable which will tell you where your workspaces is based, by default this will be a go directory in your home folder.

Go is very opionionated and has a default way of how you should structure your workspace. 

```
|- go
    |- bin
    |- src
        |- github.com
            |- rahman95
                |- go-crash-course
    |- pkg
```

#### src
The above shows an example of the default directory layout, all your projects should live under a `src` directory which will contain a versioning system directory which is `github` in my case, then under a directory named after your username which is `rahman95`. In that directory is where all your projects will sit.

#### bin
In the `bin` folder is where all your compiled code will live, once you run the compile command on your project it will be found here.

#### pkg
This is the directory where all third-party packages will sit, whenever you will add a new dependency it should be added there.

### Packages ✨

To add packages to your projects simply use;

```
go get github.com/aws/aws-sdk-go/aws
```

This command will fetch the package and place in under the `pkg` directory. For a list of packages avaialable you can use [pkg.go](https://pkg.go.dev/)

### Run 🏃🏽‍♂️

To execute a go file you can use the below command

```
go run <file.go>
```

### Compile 🚀

To compile your project, navigate to your project and run the below

```
go install
```

This will try compile and place in under the `bin` folder in your workspace.