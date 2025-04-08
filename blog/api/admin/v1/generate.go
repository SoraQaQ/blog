package v1

//go:generate protoc --proto_path=. --proto_path=../../../third_party --openapiv2_out . --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=false ./*.proto
