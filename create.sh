for i in `seq 1 10`
do
    mkdir part$i
    touch part$i/part$i.go
    echo "package main \n\n\nfunc main(){\n}" >> part$i/part$i.go
done