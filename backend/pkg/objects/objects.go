package obj

import (
    "sort"
)

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

type Image struct {
    Filename    string
    Height      int
    Width       int
}

type ImageList struct {
    Images      []Image
}

func (il ImageList) Sort() ImageList {
    sort.Slice(il.Images, func(i, j int) bool {
        return il.Images[i].Filename < il.Images[j].Filename
    })
    return il
}
