import { Component } from "@angular/core";
import { ActivatedRoute } from "@angular/router";
import { PageRoute } from "nativescript-angular/router";
import { switchMap } from "rxjs/operators";
import { BarcodeScanner } from "nativescript-barcodescanner";

import { HttpClient } from "~/shared/http-client";
import { DevicesService } from "~/shared/devices-service";

@Component({
    moduleId: module.id,
    templateUrl: "./menu.component.html"
})

export class MenuComponent {

    public user: string;
    constructor(
        private route: ActivatedRoute,
        private _devicesService: DevicesService) { }

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
        this._devicesService.getAllDevices(this.user)
            .then(res => alert("Result is " + JSON.stringify(res, null, 2)))
            .catch(err => alert("error while getting devices: " + err));

        console.log("@@@@@@@@@@@@ LS");
        alert("TAPPED!!!!!");
    }

}
