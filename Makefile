default: build

build: bin
	go build -o bin/cli main.go

cross-build:
	GOOS=linux GOARCH=amd64 go build -o bin/cli-linux main.go

bin:
	@make -p bin

sync:
	rsync -avz -e "ssh -i aws-keypair.pem" --exclude=aws-keypair.pem  \
		--exclude=.git \
		--exclude=./bin/* \
		--exclude=env \
		--delete --progress . \
		ubuntu@54.174.191.193:~/chatchat

run:
	ssh -i aws-keypair.pem ubuntu@54.174.191.193 'tmux kill-session -t mysession' || true
	ssh -i aws-keypair.pem ubuntu@54.174.191.193 'tmux new-session -d -s mysession'
	ssh -i aws-keypair.pem ubuntu@54.174.191.193 'tmux send-keys -t mysession "cd ~/chatchat && make build &&TOKEN=${TOKEN} ./bin/cli" Enter'

list-sessions:
	ssh -i aws-keypair.pem ubuntu@54.174.191.193 'tmux list-sessions'

ls-root:
	ssh -i aws-keypair.pem ubuntu@54.174.191.193 'ls -l ~'

login:
	ssh -i aws-keypair.pem ubuntu@54.174.191.193

