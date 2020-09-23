# golang-postgres-jsonb
this library is a helper for processing type jsonb postgresql

how to use:

-first, you have execute to query from postgres
-then, you need to specify struct

example

type SomeStruct struct {
    ID int64 `json:"id" db:"id"`
    Metadata jsonb.JSONB `json:"metadata" db:"metadata"`
}

-after that, you have to struct scan from query fetch
-then, you can check function:
GetValue: if you want to get stringify jsonb
IsValid: if you want to check if jsonb type is valid
IsNull: check if jsonb is null
UnmarshalJSONB: if you want to unmarshal jsonb
MarshalJSONB: if you want to marshal jsonb
