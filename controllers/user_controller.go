package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"my-api/config"
	"my-api/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Jangan lupa set header
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query(`SELECT id, username, password, created_at FROM users`)

	if err != nil {
		log.Println("Gagal mengambil data semua user >>", err)
		http.Error(w, "Gagal mengambil data semua user", http.StatusInternalServerError)
		return
	}

	// Tutup rows setelah digunakan, karena rows terbentuk dari *sql.DB
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt) // cek error

		if err != nil {
			log.Println("Gagal scan baris data >>", err)
			// Lanjutkan ke baris berikutnya
			continue
		}

		users = append(users, u)
	}

	// Cek error setelah loop selesai
	if err := rows.Err(); err != nil {
		//cek error
		log.Println("Error iterator rows saat mengambil data semua user >>", err)
		http.Error(w, "Gagal memproses data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	var u models.User

	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := config.DB.QueryRow(query, userId).Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)

	if err != nil {
		// Cek apakah errornya itu karena data tidak ketemu
		if err == sql.ErrNoRows {
			http.Error(w, "User tidak ditemukan", http.StatusNotFound)
			return
		}

		log.Println("Gagal mengambil data user >>", err)
		http.Error(w, "Gagal mengambil data user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
	result, err := config.DB.Exec(query, u.Username, u.Password, time.Now())

	if err != nil {
		log.Println("Gagal insert user baru >>", err)
		http.Error(w, "Gagal menyimpan data user", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Gagal mengambil last insert id >>", err)
		http.Error(w, "User berhasil dibuat tapi gagal mengambil id", http.StatusInternalServerError)
		return
	}

	u.Id = int(id)
	u.CreatedAt = time.Now()

	json.NewEncoder(w).Encode(u)
}

func EditUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET username = ? WHERE id = ?`
	result, err := config.DB.Exec(query, u.Username, userId)
	if err != nil {
		log.Println("Gagal mengubah data >>", err)
		http.Error(w, "Data user gagal diperbarui", http.StatusInternalServerError)
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Gagal mengambil info rows >>", err)
		http.Error(w, "Kesalahan server", http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		http.Error(w, "User tidak ditemukan atau data tidak berubah", http.StatusBadRequest)
		return
	}

	resp := map[string]string{
		"message": "Username berhasil diperbarui",
		"id":      userId,
	}

	json.NewEncoder(w).Encode(resp)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userId := vars["id"]

	query := `DELETE from users WHERE id = ?`
	result, err := config.DB.Exec(query, userId)

	if err != nil {
		log.Println("Gagal menghapus data >>", err)
		http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Gagal menghapus data >>", err)
		http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		http.Error(w, "Data tidak ditemukan atau gagal dihapus", http.StatusBadRequest)
		return
	}

	resp := map[string]string{
		"message": "User berhasil dihapus",
		"id":      userId,
	}

	json.NewEncoder(w).Encode(resp)
}
