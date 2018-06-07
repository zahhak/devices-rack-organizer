import { Component } from "@angular/core";
import { ActivatedRoute } from "@angular/router";
import { PageRoute } from "nativescript-angular/router";
import { switchMap } from "rxjs/operators";

@Component({
    moduleId: module.id,
    templateUrl: "./menu.component.html"
})

export class MenuComponent {

    public user: string;
    constructor(
        private route: ActivatedRoute) { }

    ngOnInit() {
        this.route.queryParams
            .subscribe(params => {
                this.user = params.user;
            });
    }

    public onListTapped() {
        console.log("@@@@@@@@@@@@ LS");
        alert("TAPPED!!!!!");
    }

}
