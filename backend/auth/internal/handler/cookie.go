package handler

import (
    "net/http"
)

const refreshCookieName = "refresh_token"

func setRefreshCookie(w http.ResponseWriter, token string) {
    http.SetCookie(w, &http.Cookie{
        Name:     refreshCookieName,
        Value:    token,
        HttpOnly: true,
        Secure:   false,
        SameSite: http.SameSiteLaxMode,
        Path:     "/",
        MaxAge:   7 * 24 * 60 * 60,
    })
}

func clearRefreshCookie(w http.ResponseWriter) {
    http.SetCookie(w, &http.Cookie{
        Name:     refreshCookieName,
        Value:    "",
        HttpOnly: true,
        Path:     "/",
        MaxAge:   -1,
    })
}

func getRefreshCookie(r *http.Request) (string, error) {
    cookie, err := r.Cookie(refreshCookieName)
    if err != nil {
        return "", err
    }
    return cookie.Value, nil
}