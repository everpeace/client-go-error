```
# This follows the official installation manual: https://github.com/kubernetes/client-go/blob/master/INSTALL.md#glide
$ git clone git@github.com:everpeace/client-go-error.git

$ cd client-go-error

$ cat glide.yaml

$ glide --version
glide version 0.13.1

# probably related to https://github.com/Masterminds/glide/issues/945
$  $ glide update --strip-vendor
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching k8s.io/client-go
[INFO]	--> Setting version for k8s.io/client-go to v6.0.0.
[INFO]	Resolving imports
[INFO]	Found Godeps.json file in /Users/everpeace/.glide/cache/src/https-k8s.io-client-go
[INFO]	--> Parsing Godeps metadata...
[ERROR]	Error scanning k8s.io/client-go/pkg/api: cannot find package "." in:
	/Users/everpeace/.glide/cache/src/https-k8s.io-client-go/pkg/api
[ERROR]	Failed to retrieve a list of dependencies: Error resolving imports
```
