# go_bandit

This is app for deep understanding of bandit algorithm.
![demo](https://github.com/RottenFruits/go_bandit/blob/master/demo.gif)

# Install

```bash
git
```

# Usage





docker build -t go_bandit -f Dockerfile/Dockerfile .

docker run -it --name go_bandit go_bandit /bin/bash

docker stop go_bandit
docker rm go_bandit

docker run -v /Users/ogawashouhei/Documents/project/go_bandit/src:/go/src -it --name go_bandit go_bandit /bin/bash

docker run -p 8080:8080 -v /Users/ogawashouhei/Documents/project/go_bandit/src:/go/src -it --name go_bandit go_bandit /bin/bash

go run src/go_bandit/*.go