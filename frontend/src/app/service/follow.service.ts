import {Injectable} from '@angular/core';
import {environment} from "../../environments/environment.development";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {AuthService} from "./auth.service";

@Injectable({
    providedIn: 'root'
})
export class FollowService {
    api: string = environment.api
    id!: number

    constructor(private http: HttpClient, private authService: AuthService) {}

    follow(data: any, nature: string): Observable<any> {
        switch (nature) {
            case 'follow':
                return this.http.post(`${this.api}/follow`, data)
            default:
                return this.http.post(`${this.api}/unfollow`, data)
        }
    }

    checkFollow(data: any, nature: string): Observable<any> {
        switch (nature) {
            case 'following':
                return this.http.post(`${this.api}/are-following`, data)
            default:
                return this.http.post(`${this.api}/are-we-friends`, data)
        }
    }

    unfollow(data: any): Observable<any> {
        return this.http.post(`${this.api}/unfollow`, data)
    }

    request(id: any, nature: string): Observable<any> {
        switch (nature) {
            case "decline":
                return this.http.delete(`${this.api}/decline/${id}`, {})
            default:
                return this.http.put(`${this.api}/accept/${id}`, {})
        }
    }

    getList(id: any, nature: string) {
        switch (nature) {
            case 'followings':
                return this.http.get(`${this.api}/followings/${id}`)
            case 'friends':
                return this.http.get(`${this.api}/friends/${id}`)
            default:
                return this.http.get(`${this.api}/followers/${id}`)
        }
    }

    getCount(id: any, nature: string) {
        switch (nature) {
            case 'followings':
                return this.http.get(`${this.api}/following-count/${id}`)
            case 'friends':
                return this.http.get(`${this.api}/friend-count/${id}`)
            case "all":
                return this.http.get(`${this.api}/follow-count`)
            default:
                return this.http.get(`${this.api}/follower-count/${id}`)
        }
    }

    calculate(num: number): any {
        if (num < 1000) {
            return num
        } else if (num >= 1000 && num < 1000000) {
            return `${(num / 1000).toFixed(1)}K`
        } else if (num >= 1000000 && num < 1000000000) {
            return `${(num / 1000000).toFixed(1)}M`
        } else {
            return `${(num / 1000000000).toFixed(1)}B`
        }
    }
}
