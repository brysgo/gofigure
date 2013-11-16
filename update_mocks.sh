#! /bin/bash -e

mkdir -p mocks/mock_input_log/
mockgen github.com/brysgo/gofigure/input_log Interface \
  > mocks/mock_input_log/mock_input_log.go
gofmt -w mocks/mock_input_log/mock_input_log.go

mkdir -p mocks/mock_decision_tree/
mockgen github.com/brysgo/gofigure/decision_tree Interface \
  > mocks/mock_decision_tree/mock_decision_tree.go
gofmt -w mocks/mock_decision_tree/mock_decision_tree.go

echo >&2 "OK"

