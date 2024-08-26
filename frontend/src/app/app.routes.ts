import {Routes} from '@angular/router';
import {LoginComponent} from "./pages/auth/login/login.component";
import {HomeComponent} from "./pages/home/home.component";
import {RegisterComponent} from "./pages/auth/register/register.component";
import {ProfileComponent} from "./pages/auth/profile/profile.component";

export const routes: Routes = [
    {path: '', component: HomeComponent},
    {path: 'login', component: LoginComponent},
    {path: 'register', component: RegisterComponent},
    {path: 'profile/:id', component: ProfileComponent},
    //{path: 'error', component: ErrorComponent},
    {path: '**', redirectTo: '', pathMatch: 'full'},
];
