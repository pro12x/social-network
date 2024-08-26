import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {AuthService} from "../../../service/auth.service";
import {tap} from "rxjs";
import {Router, RouterLink} from "@angular/router";

@Component({
    selector: 'app-login',
    standalone: true,
    imports: [
        ReactiveFormsModule,
        RouterLink
    ],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss'
})
export class LoginComponent implements OnInit {
    title = 'Login';
    btnText = 'Login';

    user: FormGroup = this.fb.group({
        email: [null, [Validators.required, Validators.email]],
        password: [null, [Validators.required, Validators.minLength(6)]]
    })

    constructor(
        private fb: FormBuilder,
        private authService: AuthService,
        private router: Router
    ) {}

    onSubmit() {
        if (this.user.invalid) {
            alert('Please fill in the form correctly');
            return;
        }
        console.log(this.user.value);
        this.login(this.user.value).subscribe(() => {
            console.log('Logged in');
            this.router.navigateByUrl('/home').then();
            // this.router.navigate(['/home']);
        })
        console.log('Form submitted');
    }

    login(credentials: {email: string, password: string}) {
        return this.authService.login(credentials).pipe(
            tap((res: any) => {
                console.log("Login response", res);
                if (!res.status || res.status !== 'success') {
                    alert(res.message);
                    return;
                }
                localStorage.setItem("token", res.token)
                localStorage.setItem("userID", res.user.id)
            })
        )
    }

    ngOnInit(): void {
        if (this.authService.getToken()) {
            this.router.navigate(['/home']).then();
            console.log('You are already logged in');
            return;
        }

        // this.isLoggedIn().subscribe((res: any) => console.log(res))
        console.log('Login component is running');
    }
}
