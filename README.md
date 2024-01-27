# Death Timer Controller

**Description:**

Death Timer Controller is an operator created using Operator SDK with Golang to manage resources based on a configured CRD manifest. This controller tracks and deletes resources from the cluster when their specified time limit is reached. The supported resources currently include:

1. Namespaces
2. Deployments
3. Pods

## How to Deploy:

### Step 1: Create a Cluster

Create a Kubernetes cluster using a tool like kind:

```bash
kind create cluster
```

### Step 2: Configure Sample Configuration
Configure the sample configuration located at ./config/samples/api_v1alpha1_deathtimer.yaml according to your requirements.


### Step 3: Generate Manifests
Run the following command to generate manifests:
```bash
make manifests
```

### Step 4: Deploy CRD Base
Deploy the CRD base:
```bash
kubectl apply -f config/crd/bases/api.omer.aplatony_deathtimers.yaml
```

### Step 5: Deploy CRD
Deploy the CRD:
```bash
kubectl apply -f config/samples/api_v1alpha1_deathtimer.yaml
```

### Step 6: Run the Controller
Run the controller:
```bash
make run
```

View the logs on the controller terminal for additional information.
Feel free to customize the configuration and explore the controller's capabilities.