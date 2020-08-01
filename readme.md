# Distributed systems Reliability, Causation and Group Delivery

The finality of this project is to achieve the objectives of practice number one in subject Net and Distributed System at Zaragoza University. 

## The project structure is:
 
```
  reports     -->   This folder contains project specification and requirements.
  src         -->   This folder contains all code about the project.
        chandylamport       -->   
        communication       -->     
        config  functions   -->    
        logs                -->   
        main.go             -->   
        multicast           -->   
        test                -->   
        vclock              -->   
  .gitignore  -->   File indicate files or folder to ignore
  readme.md   -->   Describe all require information you let to know about the project
```

## The objectives of the project are learn and understand:
 
* How events hold communication in distributed systems.
* Ventages and disadvantages that protocols like tcp and udp in distributed application.
* Synchronization and recovery  protocols in distributed systems.  
* When to use ventages of multicast  communication.   

# Installation 
This project requires:
```
go (>= 1.13)
```

Other library used:
* [vclock]("http://labix.org/vclock")
* [go-multicast]("https://github.com/dmichael/go-multicast")

# Source code
You can check the latest sources with the command:
> git clone https://github.com/smmanrrique/sd_paxos.git

**It's very important set correct path to run project or clone repository in folder "/home/userName/go/src/"**

# Examples executing main:

For execute main go program yo must use follow flag:
* name  --> Insert name like machine# (# is a number 1-3) 
* mode  --> Mode to execute [tcp, udp, chandy] | default tcp
* log   --> With true Send output to log file otherwise print on terminal | default false 


## TCP Communication 
You need to open one terminal by every machine and execute go script in this order.

### machina3
>go run main.go -name "machine3" -mode "tcp" -log true 

### machina2
>go run main.go -name "machine2"  -mode "tcp" -log true   

### machina1
>go run main.go -name "machine1" -mode "tcp" -log true   





cd /home/shamuel/go/sd_paxos/src/app

go run main.go -m "machine1"  -e "tcp" 
go run main.go -m "machine2"  -e "tcp" 
go run main.go -m "machine3"  -e "tcp" 


go run main.go -m "machine3"  -e "tcp" 
  167  go run main.go -m "machine3"  -e "udp" 
  168  go run main.go -m "machine3"  -e "chandy" 

  161  go run main.go -m "machine2"  -e "tcp" 
  162  go run main.go -m "machine2"  -e "udp" 
  163  go run main.go -m "machine2"  -e "chandy" 

  go run main.go -m "machine1"  -e "tcp" 
  163  go run main.go -m "machine1"  -e "udp" 
  164  go run main.go -m "machine1"  -e "chandy"

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
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="tcp" 
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="tcp"
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="tcp"
<!-- por shel udp -->
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="udp" 
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="udp"
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="udp"
<!-- por shel chandy -->
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="chandy" 
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="chandy"
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="chandy"

<!-- por shel tcp -->
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.208" -e="tcp" 
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.209" -e="tcp"
/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/main.go -c="155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400" -n=3 -p=":1400" -i="155.210.154.199" -t "155.210.154.209:1400" -d "5s"  -m=true   -e="tcp"













"/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.208\" -e=\"tcp\"" 
"/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.209\" -e=\"tcp\""
"/usr/local/go/bin/go run /home/a802400/go/sd_paxos/src/app/main.go -c=\"155.210.154.199:1400,155.210.154.209:1400,155.210.154.208:1400\" -n=3 -p=\":1400\" -i=\"155.210.154.199\" -t \"155.210.154.209:1400\" -d \"5s\"  -m=true   -e=\"tcp\""
