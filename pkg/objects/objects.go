package obj

type User struct {
    ID        string
    Username  string
    PassHash  string
    DataShare string
}

type DataShare struct {
    ShareName   string
}