enum OperationType {
    RUN = 0,
    GET = 1,
    ALL = 2,
    EACH = 3,
    EXEC = 4,
    PREPARE = 5,
}

enum DBPrivilege {
    OPEN_READONLY = 0,
    OPEN_READWRITE = 1,
    OPEN_CREATE = 2,
    OPEN_SHAREDCACHE = 3,
    OPEN_PRIVATECACHE = 4,
    OPEN_URI = 5,
}

struct AuthenticationMessage {
    1: optional DBPrivilege privilege_level;
    2: required string auth_token;
}

union ArgType {
    // number: <thrift primitive type> <corresponding Go primitive type>
    1: i64 int64;
    2: double float64;
    3: bool bool;
    4: binary byte;
    5: string string;
    // 6: i64 time.Time
}

struct OperationMessage {
    1: optional i64 operation_id; // if present, must be unique among all pending operations of a particular client
    2: required OperationType operation_type;
    3: required string sql;
    4: optional list<ArgType> args; // not sure if I'm doing this right (Thrift representation of https://pkg.go.dev/database/sql/driver#Value)
}

struct ResponseMessage {
    1: optional i64 operation_id;
    2: required string sql;
}