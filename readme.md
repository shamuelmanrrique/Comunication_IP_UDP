
cd /home/shamuel/go/src/sd_paxos/app

go run main.go -r "local" -n 3 -p=":5003" -i="127.0.1.1" -e="tcp" 
go run main.go -r "local"  -n 3 -p=":5002" -i="127.0.1.1" -e="tcp"
go run main.go -r "local" -t "127.0.1.1:5002" -d "5s" -n 3 -m=true -p=":5001" -i="127.0.1.1" -e="tcp"


go run main.go -r "local" -n 3 -p=":5003" -i="127.0.1.1" -e="udp" 
go run main.go -r "local"  -n 3 -p=":5002" -i="127.0.1.1" -e="udp"
go run main.go -r "local" -t "127.0.1.1:5002" -d "5s" -n 3 -m=true -p=":5001" -i="127.0.1.1" -e="udp"


go run main.go -r "local" -n 3 -p=":5003" -i="127.0.1.1" -e="chandy" 
go run main.go -r "local"  -n 3 -p=":5002" -i="127.0.1.1" -e="chandy"
go run main.go -r "local" -t "127.0.1.1:5002" -d "5s" -n 3 -m=true -p=":5001" -i="127.0.1.1" -e="chandy"


c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400"
go run main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -r "local" -n 3 -p=":5003" -i="127.0.1.1" -e="tcp" 
go run main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -r "local"  -n 3 -p=":5002" -i="127.0.1.1" -e="tcp"
go run main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -r "local" -t "127.0.1.1:5002" -d "5s" -n 3 -m=true -p=":5001" -i="127.0.1.1" -e="tcp"

/usr/local/go/bin/go run /home/a802400/go/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="tcp" 
/usr/local/go/bin/go run main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="tcp" 
/usr/local/go/bin/go run main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="tcp"
/usr/local/go/bin/go run main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="tcp"


export PATH=$PATH:/usr/local/go/bin;export GOPATH=/home/a802400/go;export GOROOT=/usr/local/go;
<!-- por shel tcp -->
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="tcp" 
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="tcp"
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="tcp"
<!-- por shel udp -->
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="udp" 
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="udp"
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="udp"
<!-- por shel chandy -->
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="chandy" 
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="chandy"
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="chandy"

<!-- por shel tcp -->
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="tcp" 
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="tcp"
/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="tcp"













"/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.208\" -e=\"tcp\"" 
"/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.209\" -e=\"tcp\""
"/usr/local/go/bin/go run /home/a802400/go/src/sd_paxos/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.199\" -t \"155.210.154.209:1400\" -d \"5s\"  -m=true   -e=\"tcp\""
