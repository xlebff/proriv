package main

import (
    "encoding/json"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
)

type Club struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Icon        string `json:"icon"`
    Schedule    string `json:"schedule"`
    Price       string `json:"price"`
}

func GetClubsHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("postgres", "your-connection-string")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, name, description, icon, schedule, price FROM clubs")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var clubs []Club
    for rows.Next() {
        var club Club
        err := rows.Scan(&club.ID, &club.Name, &club.Description, &club.Icon, &club.Schedule, &club.Price)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        clubs = append(clubs, club)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"clubs": clubs})
}