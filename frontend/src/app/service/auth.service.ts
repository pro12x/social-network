import {Injectable} from '@angular/core';
import {environment} from "../../environments/environment.development";
import {HttpClient} from "@angular/common/http";
import {map, Observable, of, tap} from "rxjs";

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    api: string = environment.api

    constructor(private http: HttpClient) {
    }

    login(credentials: any): Observable<any> {
        return this.http.post(`${this.api}/login`, credentials)
    }

    register(user: any): Observable<any> {
        return this.http.post(`${this.api}/register`, user)
    }

    logout(token: any): Observable<any> {
        return this.http.post(`${this.api}/logout`, token)
    }

    checkOnlineStatus(token: any): Observable<any> {
        return this.http.post(`${this.api}/is_online`, token)
    }

    isLoggedIn(): Observable<boolean> {
        const data = {
            token: localStorage.getItem('token')
        }

        if (!localStorage.getItem('token') || !localStorage.getItem('userID')) {
            console.log('No token or user id')
            return of(false)
        } else {
            return this.checkOnlineStatus(data).pipe(
                map(response => response.is_online && response.status == 'online')
            )
        }
    }

    getToken(): string | null {
        return localStorage.getItem('token')
    }

    getUserID(): number | null {
        return localStorage.getItem('userID') ? parseInt(localStorage.getItem('userID')!) : null
    }

    getAll() {
        return this.http.get(`${this.api}/users`)
    }

    getUser(id: any) {
        return this.http.get(`${this.api}/profile/${id}`)
    }

    updateUser(id: any, user: any) {
        return this.http.put(`${this.api}/update-profile/${id}`, user)
    }
}
