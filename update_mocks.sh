#! /bin/bash -e

mkdir -p $GOPATH/src/gofigure_mocks/mock_input_log/
mockgen github.com/brysgo/gofigure/input_log Interface \
  > $GOPATH/src/gofigure_mocks/mock_input_log/mock_input_log.go
gofmt -w $GOPATH/src/gofigure_mocks/mock_input_log/mock_input_log.go

mkdir -p $GOPATH/src/gofigure_mocks/mock_binary_tree/
mockgen github.com/brysgo/gofigure/binary_tree Interface \
  > $GOPATH/src/gofigure_mocks/mock_binary_tree/mock_binary_tree.go
gofmt -w $GOPATH/src/gofigure_mocks/mock_binary_tree/mock_binary_tree.go

mkdir -p $GOPATH/src/gofigure_mocks/mock_decision_tree/
mockgen github.com/brysgo/gofigure/decision_tree Interface \
  > $GOPATH/src/gofigure_mocks/mock_decision_tree/mock_decision_tree.go
gofmt -w $GOPATH/src/gofigure_mocks/mock_decision_tree/mock_decision_tree.go

echo >&2 "OK"

