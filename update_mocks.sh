#! /bin/bash -e

MOCKPATH=$GOPATH'src/gofigure_mocks'
PROJECTPATH=$GOPATH'src/github.com/brysgo/gofigure'

old_ifs=${IFS}
IFS=$'\n'

# Create directory structure for mocks
for directory in $( find $PROJECTPATH -type d ); do
  new_directory=$( echo ${directory} | sed "s%$PROJECTPATH/%$MOCKPATH/%" )
	# echo "mkdir -p \"${new_directory}\""
	mkdir -p "${new_directory}"
done


# Generate mocks
for file in $( find $PROJECTPATH -name "*test.go" -prune -o -type f -print | grep '\.go$' ); do
	mock_file=$( echo ${file} | sed "s%$PROJECTPATH/%$MOCKPATH/%" | sed -E "s%(.*)/(.*)\.go%\1/mock_\2.go%" )
	echo "mockgen -source=${file} > ${mock_file}"
	mockgen -source=${file} > ${mock_file}
done


IFS=${old_ifs}

echo >&2 "OK"

