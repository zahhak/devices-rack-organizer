import { Component } from "@angular/core";

@Component({
    moduleId: module.id,
    templateUrl: "./login.component.html"
})

export class LoginComponent {

    public username: string = "username";

    public submit() {
        alert("Text: " + this.username);
    }
}
