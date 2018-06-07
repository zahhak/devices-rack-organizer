import { Component } from "@angular/core";
import { ActivatedRoute } from "@angular/router";
import { PageRoute } from "nativescript-angular/router";
import { switchMap } from "rxjs/operators";
import { BarcodeScanner } from "nativescript-barcodescanner";

@Component({
    moduleId: module.id,
    templateUrl: "./menu.component.html"
})

export class MenuComponent {

    public user: string;
    constructor(
        private route: ActivatedRoute) { }

    public getDevice() {
        const scanner = new BarcodeScanner();
        scanner.scan({
            cancelLabel: "Stop scanning",
            message: "Go scan something",
            preferFrontCamera: false,
            showFlipCameraButton: true
        }).then((result) => {
            const deviceId = result.text;
        });
    }

    public onScanResult(evnt) {
        console.log(evnt)
    }

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
