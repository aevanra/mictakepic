package obj

type User struct {
    ID        string
    Username  string
    PassHash  string
    DataShare string
    Admin bool
}

type DataShare struct {
    ShareName   string
}
