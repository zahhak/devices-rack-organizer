import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { LoginComponent } from "~/login/login.component";
import { NativeScriptRouterModule } from "nativescript-angular/router";
import { NativeScriptFormsModule } from "nativescript-angular/forms";

const routes: Routes = [
    { path: "", component: LoginComponent }
];

@NgModule({
    imports: [
        NativeScriptRouterModule.forChild(routes),
        NativeScriptFormsModule
    ],
    declarations: [
        LoginComponent
    ],
    bootstrap: [LoginComponent]
})
export class LoginModule { }
