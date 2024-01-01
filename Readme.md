## Running your client-go application as a pod in a cluster

- `docker build -t clientgo:0.1.0 .`

- `docker tag clientgo:0.1.0 yashpimple/clientgo:0.1.0`

- `docker push yashpimple/clientgo:0.1.0`

- `kubectl apply -f clientgo.yaml`

<!-- Create role and bind it to default service account  -->

- `kubectl create role poddepl --resource pods,depployments --verb list`

- `kubectl create rolebinding poddepl --role poddepl --serviceaccount default:default`