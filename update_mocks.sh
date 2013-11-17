#! /bin/bash -e

mkdir -p ../gofigure_mocks/mock_input_log/
mockgen github.com/brysgo/gofigure/input_log Interface \
  > ../gofigure_mocks/mock_input_log/mock_input_log.go
gofmt -w ../gofigure_mocks/mock_input_log/mock_input_log.go

mkdir -p ../gofigure_mocks/mock_decision_tree/
mockgen github.com/brysgo/gofigure/decision_tree Interface \
  > ../gofigure_mocks/mock_decision_tree/mock_decision_tree.go
gofmt -w ../gofigure_mocks/mock_decision_tree/mock_decision_tree.go

echo >&2 "OK"

