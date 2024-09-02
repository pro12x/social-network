import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {AuthService} from "../../../service/auth.service";
import {Router, RouterLink} from "@angular/router";
import {UtilsService} from "../../../service/utils.service";

@Component({
    selector: 'app-register',
    standalone: true,
    imports: [
        ReactiveFormsModule,
        RouterLink
    ],
    templateUrl: './register.component.html',
    styleUrl: './register.component.scss'
})
export class RegisterComponent implements OnInit {
    title: string = 'Register';
    btnText: string = 'Register';
    age!: number
    avatar: File | null = null

    user: FormGroup = this.fb.group({
        firstname: [null, [Validators.required, Validators.minLength(3)]],
        lastname: [null, [Validators.required, Validators.minLength(3)]],
        email: [null, [Validators.required, Validators.email]],
        password: [null, [Validators.required, Validators.minLength(6)]],
        password_confirmation: [null, [Validators.required, Validators.minLength(6)]],
        date_of_birth: [null, [Validators.required]],
        avatar: [null],
        nickname: [null],
        about_me: [null],
    })

    constructor(
        private fb: FormBuilder,
        private authService: AuthService,
        private utilsService: UtilsService,
        private router: Router
    ) {
    }

    onRegister() {
        const data = {
            ...this.user.value,
            avatar: 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png'
        }

        this.age = this.checkAge(data.date_of_birth)

        if (this.age < 12 || this.age > 120) {
            this.utilsService.onSnackBar('You must be between 12 and 120 years old to register', 'warning');
            return
        }

        if (data.password !== data.password_confirmation) {
            this.utilsService.onSnackBar('Password and password confirmation do not match', 'warning');
            return;
        }

        if (this.user.invalid) {
            this.utilsService.onSnackBar('Please fill in the form correctly', 'warning');
            return;
        } else {
            this.authService.register(data).subscribe((response: any) => {
                if (response.status !== 201) {
                    this.utilsService.onSnackBar(response.message, 'error');
                    return;
                }
                this.utilsService.onSnackBar('You are now registered', 'success');
                this.router.navigateByUrl('/login').then();
            }, (error) => {
                this.utilsService.onSnackBar(error.error.message, 'error');
            })
        }
    }

    checkAge(data: Date): number {
        return Math.floor(Math.abs(Date.now() - new Date(data).getTime()) / (1000 * 3600 * 24 * 365))
    }

    ngOnInit(): void {
        this.utilsService.setTitle(this.title);
        if (this.authService.getToken()) {
            this.router.navigate(['/home']).then();
            this.utilsService.onSnackBar('You are already logged in', 'info');
            return;
        }
    }
}
