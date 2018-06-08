import { Component } from "@angular/core";
import { RouterExtensions } from "nativescript-angular/router";

@Component({
    moduleId: module.id,
    templateUrl: "./login.component.html"
})

export class LoginComponent {
    constructor(
        private _routerExtensions: RouterExtensions
    ) { }

    public username: string = "username";

    public submit() {
        alert("Text: " + this.username);
        this._routerExtensions.navigate(["/menu"], {
            clearHistory: true,
            animated: true,
            transition: {
                name: "slideBottom",
                duration: 200,
                curve: "ease"
            },
            queryParams: {
                user: this.username
            }
        })
    }
}
