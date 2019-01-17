# go_bandit

This is app for deep understanding of bandit algorithm.
![demo](https://github.com/RottenFruits/go_bandit/blob/master/demo.gif)

# Install

```bash
git clone https://github.com/RottenFruits/go_bandit.git
```

# Usage

- build docker
```bash
cd go_bandit
docker build -t go_bandit -f Dockerfile/Dockerfile .
```
- container login
```bash
docker run -p 8080:8080 -it --name go_bandit go_bandit /bin/bash
```

- app start
```bash
cd src/go_bandit
go run *.go
```

- access server
http://localhost:8080/