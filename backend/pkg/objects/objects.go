package obj

type User struct {
    ID        string
    Username  string
    PassHash  string
    DefaultDataShare string
    AllDatashares []string
    Admin bool
}

type DataShare struct {
    ShareName   string
}
