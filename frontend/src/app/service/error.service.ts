import {Injectable} from '@angular/core';
import {BehaviorSubject} from "rxjs";

@Injectable({
    providedIn: 'root'
})
export class ErrorService {
    private errorSubject = new BehaviorSubject<boolean>(false)
    error$ = this.errorSubject.asObservable()

    constructor() {}

    setError(error: boolean) {
        this.errorSubject.next(error)
    }

    getError() {
        return this.errorSubject.getValue()
    }
}
