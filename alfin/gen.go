package alfin

////go:generate ./meta_generator.py abi [Example]
////go:generate ./meta_generator.py abi [InventoryItemAndDetail]
////go:generate ./meta_generator.py abi [InventoryItem,InventoryItemDetail]

////go:generate rice embed-go

//go:generate mkdir -p pb
//go:generate protoc --go_out=plugins=grpc:pb --go_opt=paths=source_relative outliers.proto
//go:generate make proto
