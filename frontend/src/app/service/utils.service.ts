import {Injectable} from '@angular/core';
import {MatSnackBar} from "@angular/material/snack-bar";
import {BehaviorSubject, Observable} from "rxjs";
import {Router} from "@angular/router";

@Injectable({
    providedIn: 'root'
})
export class UtilsService {
    private _title = new BehaviorSubject<string>('');

    constructor(private snackBar: MatSnackBar, private router: Router) {
    }

    onSnackBar(message: string, color: string) {
        this.snackBar.open(message, 'Close', {
            duration: 5000,
            horizontalPosition: 'center',
            verticalPosition: 'top',
            panelClass: ['custom-snackbar', `snackbar-${color}`]
        });
    }

    setTitle(title: string): void {
        this._title.next(title);
    }

    getTitle() {
        return this._title.asObservable().pipe(res => {
            return res;
        });
    }
}
