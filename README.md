# `yeah`: `yes` but cooler

## Intrduction

`yeah` is something like command `yes` in Linux-based systems, but instead of printing  
indefinately to terminal, it generates fixed size files.

## How to run

```sh
Usage of bin/yeah:
  -f string
    	Size of file(s) in bytes, comma separated (default "100")
  -s	Synchron mode
```

Very simple. Just open a new terminal and locate the binary. Then run:
```sh
yeah -f 104,512,10k
```

Results:
```sh
-rw-rw-r-- 1 ubuntu ubuntu  104 Apr 15 15:55 yeah_0-104.txt
-rw-rw-r-- 1 ubuntu ubuntu  512 Apr 15 15:55 yeah_1-512.txt
-rw-rw-r-- 1 ubuntu ubuntu 1024 Apr 15 15:55 yeah_2-1k.txt
```

## Performance

With current level of optimization in v0, result for synchron and multithread  
runs are as follows:

### Single file
```sh
ubuntu@uvm08:~/yeah$ ./yeah -f 25M
Elapsed Time (ms): 6108
```

### Multithread (non go-routine)
```sh
ubuntu@uvm08:~/yeah$ ./yeah -f 25M,25M,25M,25M,25M,25M,25M 
Elapsed Time (ms): 8357
```

### Simple (no go-routine)
```sh
ubuntu@uvm08:~/yeah$ ./yeah -f 25M,25M,25M,25M,25M,25M,25M  -s
Elapsed Time (ms): 39512
```
