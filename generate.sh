read_dir(){
    for file in `ls -a $1`
    do
        if [ -d $1"/"$file ]
        then
            if [[ $file != '.' && $file != '..' ]]
            then
                read_dir $1"/"$file
            fi
        else
            if [ "${file##*.}"x = "proto"x ]
            then
              eval "protoc -I ../ -I ./ --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $1"/"$file && protoc-go-inject-tag -input=\"$1/*.pb.go\""
            fi
        fi
    done
}

for file in ./*;
do
  if [ -d "$file" ]; then
    read_dir $file
  fi
done