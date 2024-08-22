This is custom download manager written from scratch in Golang.
I always had issues with downloading media on my Macbook, as there is no good download manager that exists for Macs like we have for Windows.
So I decided to write one by myself.

**HOW TO BUILD AND RUN**
 - This project contains a _Makefile_ so all you have to do it, open the root of the project, and write **make build**
 - This will generate an executable binary in the _bin_ folder in the root called ***gdm***
 - Now just run ./gdm -u "_url of the resource you want to download_"
 - It will download the file on that url using 8 threads, within seconds.
