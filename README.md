# Word of Wisdom - PoW Protected TCP Server

Word of Wisdom is a TCP-based server that provides wisdom quotes, protected against DDoS attacks using a Proof of Work (PoW) mechanism.

The PoW method used in this test task is a simple SHA-256-based proof-of-work algorithm.

## How it works
* **Challenge Generation:**  
  The server generates a random challenge (a string) along with a difficulty parameter, called complexity. In this context, complexity denotes the required number of zeros at the beginning of the SHA-256 hash.   
  Complexity could be changed via env variable.
* **Nonce Finding:**  
  The client iterates through possible nonce values, concatenating each one with the challenge. It then computes the SHA-256 hash of the combined string.
* **Target Comparison:**  
  The client checks whether the resulting hash starts with the number of zeros specified by the complexity. For example, if complexity is 4, the hash must start with at least “0000”.
* **Validation:**
Once the client finds a nonce that yields a hash with the required number of leading zeros, it sends that nonce back to the server. The server independently verifies the proof by recalculating the hash and confirming that it meets the condition.

This method, while conceptually similar to Bitcoin’s proof-of-work mechanism, is simplified and tailored for this test task to protect the TCP server from DDoS attacks by requiring clients to perform a computationally intensive task before receiving a quote.

## Why SHA-256-base implementation?
1. Simple to implement on client
2. Simple to validate on server
3. Proof can be easily checked manually (leading 0 from HEX representation)
4. Complexity is variable depends on needs (potentially can be changed dynamically)
5. Because this implementation is well-known in the community, it reduces the development effort required on the client side.

## How to launch solution?

### I. Run Server
**1. Clone repository**

**2. Build docker image**
```shell
docker build -f Dockerfile.server -t word-of-wisdom-server .
```

**3. Create common docker network**
```shell
docker network create word-net
```

**4. Run**
```shell
docker run --env-file ./cmd/server/.env -p 8080:8080 --network word-net --name word-server word-of-wisdom-server
```

### II. Run Client

**1. Build docker image**
```shell
docker build -f Dockerfile.client -t word-of-wisdom-client .
```

**2. Run client**
```shell
docker run --env-file ./cmd/client/.env --network word-net word-of-wisdom-client
```

