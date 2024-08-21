import {Routes} from '@angular/router';
import {ErrorComponent} from "./pages/error/error.component";
import {HomeComponent} from "./pages/home/home.component";

export const routes: Routes = [
    {path: '', component: HomeComponent},
    {path: 'error', component: ErrorComponent},
    {path: '**', redirectTo: 'error', pathMatch: 'full'},
];
